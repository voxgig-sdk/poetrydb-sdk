<?php
declare(strict_types=1);

// Poetrydb SDK

require_once __DIR__ . '/utility/struct/Struct.php';
require_once __DIR__ . '/core/UtilityType.php';
require_once __DIR__ . '/core/Spec.php';
require_once __DIR__ . '/core/Helpers.php';

// Load utility registration
require_once __DIR__ . '/utility/Register.php';

// Load config and features
require_once __DIR__ . '/config.php';
require_once __DIR__ . '/feature/BaseFeature.php';
require_once __DIR__ . '/features.php';

use Voxgig\Struct\Struct;

class PoetrydbSDK
{
    public string $mode;
    public array $features;
    public ?array $options;

    private $_utility;
    private $_rootctx;

    public function __construct(array $options = [])
    {
        $this->mode = "live";
        $this->features = [];
        $this->options = null;

        $utility = new PoetrydbUtility();
        $this->_utility = $utility;

        $config = PoetrydbConfig::make_config();

        $this->_rootctx = ($utility->make_context)([
            "client" => $this,
            "utility" => $utility,
            "config" => $config,
            "options" => $options ?? [],
            "shared" => [],
        ], null);

        $this->options = ($utility->make_options)($this->_rootctx);

        if (Struct::getpath($this->options, "feature.test.active") === true) {
            $this->mode = "test";
        }

        $this->_rootctx->options = $this->options;

        // Add features from config.
        $feature_opts = PoetrydbHelpers::to_map(Struct::getprop($this->options, "feature"));
        if ($feature_opts) {
            $items = Struct::items($feature_opts);
            if ($items) {
                foreach ($items as $item) {
                    $fname = $item[0];
                    $fopts = PoetrydbHelpers::to_map($item[1]);
                    if ($fopts && isset($fopts["active"]) && $fopts["active"] === true) {
                        ($utility->feature_add)($this->_rootctx, PoetrydbFeatures::make_feature($fname));
                    }
                }
            }
        }

        // Add extension features.
        $extend_val = Struct::getprop($this->options, "extend");
        if (is_array($extend_val)) {
            foreach ($extend_val as $f) {
                if (is_object($f) && method_exists($f, 'get_name')) {
                    ($utility->feature_add)($this->_rootctx, $f);
                }
            }
        }

        // Initialize features.
        foreach ($this->features as $f) {
            ($utility->feature_init)($this->_rootctx, $f);
        }

        ($utility->feature_hook)($this->_rootctx, "PostConstruct");
    }

    public function options_map(): array
    {
        $out = Struct::clone($this->options);
        return is_array($out) ? $out : [];
    }

    public function get_utility()
    {
        return PoetrydbUtility::copy($this->_utility);
    }

    public function get_root_ctx()
    {
        return $this->_rootctx;
    }

    public function prepare(array $fetchargs = []): mixed
    {
        $utility = $this->_utility;
        $fetchargs = $fetchargs ?? [];

        $ctrl = PoetrydbHelpers::to_map(Struct::getprop($fetchargs, "ctrl")) ?? [];

        $ctx = ($utility->make_context)([
            "opname" => "prepare",
            "ctrl" => $ctrl,
        ], $this->_rootctx);

        $opts = $this->options;
        $path = Struct::getprop($fetchargs, "path") ?? "";
        $path = is_string($path) ? $path : "";
        $method_val = Struct::getprop($fetchargs, "method") ?? "GET";
        $method_val = is_string($method_val) ? $method_val : "GET";
        $params = PoetrydbHelpers::to_map(Struct::getprop($fetchargs, "params")) ?? [];
        $query = PoetrydbHelpers::to_map(Struct::getprop($fetchargs, "query")) ?? [];
        $headers = ($utility->prepare_headers)($ctx);

        $base = Struct::getprop($opts, "base") ?? "";
        $base = is_string($base) ? $base : "";
        $prefix = Struct::getprop($opts, "prefix") ?? "";
        $prefix = is_string($prefix) ? $prefix : "";
        $suffix = Struct::getprop($opts, "suffix") ?? "";
        $suffix = is_string($suffix) ? $suffix : "";

        $ctx->spec = new PoetrydbSpec([
            "base" => $base, "prefix" => $prefix, "suffix" => $suffix,
            "path" => $path, "method" => $method_val,
            "params" => $params, "query" => $query, "headers" => $headers,
            "body" => Struct::getprop($fetchargs, "body"),
            "step" => "start",
        ]);

        // Merge user-provided headers.
        $uh = Struct::getprop($fetchargs, "headers");
        if (is_array($uh)) {
            foreach ($uh as $k => $v) {
                $ctx->spec->headers[$k] = $v;
            }
        }

        [$_, $err] = ($utility->prepare_auth)($ctx);
        if ($err) {
            return ($utility->make_error)($ctx, $err);
        }

        [$fetchdef, $fd_err] = ($utility->make_fetch_def)($ctx);
        if ($fd_err) {
            return ($utility->make_error)($ctx, $fd_err);
        }
        return $fetchdef;
    }

    public function direct(array $fetchargs = []): mixed
    {
        $utility = $this->_utility;

        // direct() is the raw-HTTP escape hatch: it never throws, it returns
        // an {ok, err, ...} dict. prepare() now raises on error, so catch it
        // and surface the failure through the dict instead.
        try {
            $fetchdef = $this->prepare($fetchargs);
        } catch (\Throwable $err) {
            return ["ok" => false, "err" => $err];
        }

        $fetchargs = $fetchargs ?? [];
        $ctrl = PoetrydbHelpers::to_map(Struct::getprop($fetchargs, "ctrl")) ?? [];

        $ctx = ($utility->make_context)([
            "opname" => "direct",
            "ctrl" => $ctrl,
        ], $this->_rootctx);

        $url = $fetchdef["url"] ?? "";
        [$fetched, $fetch_err] = ($utility->fetcher)($ctx, $url, $fetchdef);

        if ($fetch_err) {
            return ["ok" => false, "err" => $fetch_err];
        }

        if ($fetched === null) {
            return [
                "ok" => false,
                "err" => $ctx->make_error("direct_no_response", "response: undefined"),
            ];
        }

        if (is_array($fetched)) {
            $status = PoetrydbHelpers::to_int(Struct::getprop($fetched, "status"));
            $headers = Struct::getprop($fetched, "headers") ?? [];

            // No-body responses (204, 304) and explicit zero content-length
            // must skip JSON parsing — calling json() on an empty body errors.
            $content_length = is_array($headers) ? ($headers["content-length"] ?? null) : null;
            $no_body = $status === 204 || $status === 304 || (string)$content_length === "0";

            $json_data = null;
            if (!$no_body) {
                $jf = Struct::getprop($fetched, "json");
                if (is_callable($jf)) {
                    try {
                        $json_data = $jf();
                    } catch (\Throwable $e) {
                        // Non-JSON body — leave data null but keep status/ok.
                        $json_data = null;
                    }
                }
            }

            return [
                "ok" => $status >= 200 && $status < 300,
                "status" => $status,
                "headers" => Struct::getprop($fetched, "headers"),
                "data" => $json_data,
            ];
        }

        return [
            "ok" => false,
            "err" => $ctx->make_error("direct_invalid", "invalid response type"),
        ];
    }


    private $_author = null;

    // Canonical facade: $client->Author()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->author()
    // resolves here too.
    public function Author($data = null)
    {
        require_once __DIR__ . '/entity/author_entity.php';
        if ($data === null) {
            if ($this->_author === null) {
                $this->_author = new AuthorEntity($this, null);
            }
            return $this->_author;
        }
        return new AuthorEntity($this, $data);
    }


    private $_authorab = null;

    // Canonical facade: $client->Authorab()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->authorab()
    // resolves here too.
    public function Authorab($data = null)
    {
        require_once __DIR__ . '/entity/authorab_entity.php';
        if ($data === null) {
            if ($this->_authorab === null) {
                $this->_authorab = new AuthorabEntity($this, null);
            }
            return $this->_authorab;
        }
        return new AuthorabEntity($this, $data);
    }


    private $_combined_search = null;

    // Canonical facade: $client->CombinedSearch()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->combined_search()
    // resolves here too.
    public function CombinedSearch($data = null)
    {
        require_once __DIR__ . '/entity/combined_search_entity.php';
        if ($data === null) {
            if ($this->_combined_search === null) {
                $this->_combined_search = new CombinedSearchEntity($this, null);
            }
            return $this->_combined_search;
        }
        return new CombinedSearchEntity($this, $data);
    }


    private $_combined_search_with_field = null;

    // Canonical facade: $client->CombinedSearchWithField()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->combined_search_with_field()
    // resolves here too.
    public function CombinedSearchWithField($data = null)
    {
        require_once __DIR__ . '/entity/combined_search_with_field_entity.php';
        if ($data === null) {
            if ($this->_combined_search_with_field === null) {
                $this->_combined_search_with_field = new CombinedSearchWithFieldEntity($this, null);
            }
            return $this->_combined_search_with_field;
        }
        return new CombinedSearchWithFieldEntity($this, $data);
    }


    private $_line = null;

    // Canonical facade: $client->Line()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->line()
    // resolves here too.
    public function Line($data = null)
    {
        require_once __DIR__ . '/entity/line_entity.php';
        if ($data === null) {
            if ($this->_line === null) {
                $this->_line = new LineEntity($this, null);
            }
            return $this->_line;
        }
        return new LineEntity($this, $data);
    }


    private $_linecount = null;

    // Canonical facade: $client->Linecount()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->linecount()
    // resolves here too.
    public function Linecount($data = null)
    {
        require_once __DIR__ . '/entity/linecount_entity.php';
        if ($data === null) {
            if ($this->_linecount === null) {
                $this->_linecount = new LinecountEntity($this, null);
            }
            return $this->_linecount;
        }
        return new LinecountEntity($this, $data);
    }


    private $_poemcount = null;

    // Canonical facade: $client->Poemcount()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->poemcount()
    // resolves here too.
    public function Poemcount($data = null)
    {
        require_once __DIR__ . '/entity/poemcount_entity.php';
        if ($data === null) {
            if ($this->_poemcount === null) {
                $this->_poemcount = new PoemcountEntity($this, null);
            }
            return $this->_poemcount;
        }
        return new PoemcountEntity($this, $data);
    }


    private $_random = null;

    // Canonical facade: $client->Random()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->random()
    // resolves here too.
    public function Random($data = null)
    {
        require_once __DIR__ . '/entity/random_entity.php';
        if ($data === null) {
            if ($this->_random === null) {
                $this->_random = new RandomEntity($this, null);
            }
            return $this->_random;
        }
        return new RandomEntity($this, $data);
    }


    private $_title = null;

    // Canonical facade: $client->Title()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->title()
    // resolves here too.
    public function Title($data = null)
    {
        require_once __DIR__ . '/entity/title_entity.php';
        if ($data === null) {
            if ($this->_title === null) {
                $this->_title = new TitleEntity($this, null);
            }
            return $this->_title;
        }
        return new TitleEntity($this, $data);
    }


    private $_titleab = null;

    // Canonical facade: $client->Titleab()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->titleab()
    // resolves here too.
    public function Titleab($data = null)
    {
        require_once __DIR__ . '/entity/titleab_entity.php';
        if ($data === null) {
            if ($this->_titleab === null) {
                $this->_titleab = new TitleabEntity($this, null);
            }
            return $this->_titleab;
        }
        return new TitleabEntity($this, $data);
    }



    public static function test(?array $testopts = null, ?array $sdkopts = null): self
    {
        $sdkopts = $sdkopts ?? [];
        $sdkopts = Struct::clone($sdkopts);
        $sdkopts = is_array($sdkopts) ? $sdkopts : [];

        $testopts = $testopts ?? [];
        $testopts = Struct::clone($testopts);
        $testopts = is_array($testopts) ? $testopts : [];
        $testopts["active"] = true;

        if (!isset($sdkopts["feature"])) {
            $sdkopts["feature"] = [];
        }
        $sdkopts["feature"]["test"] = $testopts;

        $sdk = new PoetrydbSDK($sdkopts);
        $sdk->mode = "test";
        return $sdk;
    }
}
