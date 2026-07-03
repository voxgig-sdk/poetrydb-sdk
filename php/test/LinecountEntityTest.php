<?php
declare(strict_types=1);

// Linecount entity test

require_once __DIR__ . '/../poetrydb_sdk.php';
require_once __DIR__ . '/Runner.php';

use PHPUnit\Framework\TestCase;
use Voxgig\Struct\Struct as Vs;

class LinecountEntityTest extends TestCase
{
    public function test_create_instance(): void
    {
        $testsdk = PoetrydbSDK::test(null, null);
        $ent = $testsdk->Linecount(null);
        $this->assertNotNull($ent);
    }

    public function test_basic_flow(): void
    {
        $setup = linecount_basic_setup(null);
        // Per-op sdk-test-control.json skip.
        $_live = !empty($setup["live"]);
        foreach (["list", "load"] as $_op) {
            [$_shouldSkip, $_reason] = Runner::is_control_skipped("entityOp", "linecount." . $_op, $_live ? "live" : "unit");
            if ($_shouldSkip) {
                $this->markTestSkipped($_reason ?? "skipped via sdk-test-control.json");
                return;
            }
        }
        // The basic flow consumes synthetic IDs from the fixture. In live mode
        // without an *_ENTID env override, those IDs hit the live API and 4xx.
        if (!empty($setup["synthetic_only"])) {
            $this->markTestSkipped("live entity test uses synthetic IDs from fixture — set POETRYDB_TEST_LINECOUNT_ENTID JSON to run live");
            return;
        }
        $client = $setup["client"];

        // Bootstrap entity data from existing test data.
        $linecount_ref01_data_raw = Vs::items(Helpers::to_map(
            Vs::getpath($setup["data"], "existing.linecount")));
        $linecount_ref01_data = null;
        if (count($linecount_ref01_data_raw) > 0) {
            $linecount_ref01_data = Helpers::to_map($linecount_ref01_data_raw[0][1]);
        }

        // LIST
        $linecount_ref01_ent = $client->Linecount(null);
        $linecount_ref01_match = [
            "linecount" => $setup["idmap"]["linecount01"],
            "output_field" => $setup["idmap"]["output_field01"],
        ];

        [$linecount_ref01_list_result, $err] = $linecount_ref01_ent->list($linecount_ref01_match, null);
        $this->assertNull($err);
        $this->assertIsArray($linecount_ref01_list_result);

        // LOAD
        $linecount_ref01_match_dt0 = [];
        [$linecount_ref01_data_dt0_loaded, $err] = $linecount_ref01_ent->load($linecount_ref01_match_dt0, null);
        $this->assertNull($err);
        $this->assertNotNull($linecount_ref01_data_dt0_loaded);

    }
}

function linecount_basic_setup($extra)
{
    Runner::load_env_local();

    $entity_data_file = __DIR__ . '/../../.sdk/test/entity/linecount/LinecountTestData.json';
    $entity_data_source = file_get_contents($entity_data_file);
    $entity_data = json_decode($entity_data_source, true);

    $options = [];
    $options["entity"] = $entity_data["existing"];

    $client = PoetrydbSDK::test($options, $extra);

    // Generate idmap.
    $idmap = [];
    foreach (["linecount01", "linecount02", "linecount03", "output_field01"] as $k) {
        $idmap[$k] = strtoupper($k);
    }

    // Detect ENTID env override before envOverride consumes it. When live
    // mode is on without a real override, the basic test runs against synthetic
    // IDs from the fixture and 4xx's. Surface this so the test can skip.
    $entid_env_raw = getenv("POETRYDB_TEST_LINECOUNT_ENTID");
    $idmap_overridden = $entid_env_raw !== false && str_starts_with(trim($entid_env_raw), "{");

    $env = Runner::env_override([
        "POETRYDB_TEST_LINECOUNT_ENTID" => $idmap,
        "POETRYDB_TEST_LIVE" => "FALSE",
        "POETRYDB_TEST_EXPLAIN" => "FALSE",
        "POETRYDB_APIKEY" => "NONE",
    ]);

    $idmap_resolved = Helpers::to_map(
        $env["POETRYDB_TEST_LINECOUNT_ENTID"]);
    if ($idmap_resolved === null) {
        $idmap_resolved = Helpers::to_map($idmap);
    }

    if ($env["POETRYDB_TEST_LIVE"] === "TRUE") {
        $merged_opts = Vs::merge([
            [
                "apikey" => $env["POETRYDB_APIKEY"],
            ],
            $extra ?? [],
        ]);
        $client = new PoetrydbSDK(Helpers::to_map($merged_opts));
    }

    $live = $env["POETRYDB_TEST_LIVE"] === "TRUE";
    return [
        "client" => $client,
        "data" => $entity_data,
        "idmap" => $idmap_resolved,
        "env" => $env,
        "explain" => $env["POETRYDB_TEST_EXPLAIN"] === "TRUE",
        "live" => $live,
        "synthetic_only" => $live && !$idmap_overridden,
        "now" => (int)(microtime(true) * 1000),
    ];
}
