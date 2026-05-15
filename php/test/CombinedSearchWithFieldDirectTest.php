<?php
declare(strict_types=1);

// CombinedSearchWithField direct test

require_once __DIR__ . '/../poetrydb_sdk.php';
require_once __DIR__ . '/Runner.php';

use PHPUnit\Framework\TestCase;

class CombinedSearchWithFieldDirectTest extends TestCase
{
    public function test_direct_list_combined_search_with_field(): void
    {
        $setup = combined_search_with_field_direct_setup([
            ["id" => "direct01"],
            ["id" => "direct02"],
        ]);
        [$_shouldSkip, $_reason] = Runner::is_control_skipped("direct", "direct-list-combined_search_with_field", $setup["live"] ? "live" : "unit");
        if ($_shouldSkip) {
            $this->markTestSkipped($_reason ?? "skipped via sdk-test-control.json");
            return;
        }
        if ($setup["live"]) {
            foreach (["input_field101", "input_field201", "output_field01", "search_term101", "search_term201"] as $_liveKey) {
                if (!isset($setup["idmap"][$_liveKey]) || $setup["idmap"][$_liveKey] === null) {
                    $this->markTestSkipped("live test needs $_liveKey via *_ENTID env var (synthetic IDs only)");
                    return;
                }
            }
        }
        $client = $setup["client"];

        $params = [];
        if ($setup["live"]) {
            $params["input_field1"] = $setup["idmap"]["input_field101"];
        } else {
            $params["input_field1"] = "direct01";
        }
        if ($setup["live"]) {
            $params["input_field2"] = $setup["idmap"]["input_field201"];
        } else {
            $params["input_field2"] = "direct01";
        }
        if ($setup["live"]) {
            $params["output_field"] = $setup["idmap"]["output_field01"];
        } else {
            $params["output_field"] = "direct01";
        }
        if ($setup["live"]) {
            $params["search_term1"] = $setup["idmap"]["search_term101"];
        } else {
            $params["search_term1"] = "direct01";
        }
        if ($setup["live"]) {
            $params["search_term2"] = $setup["idmap"]["search_term201"];
        } else {
            $params["search_term2"] = "direct01";
        }

        [$result, $err] = $client->direct([
            "path" => "{input_field1},{input_field2}/{search_term1};{search_term2}/{output_field}",
            "method" => "GET",
            "params" => $params,
        ]);
        if ($setup["live"]) {
            // Live mode is lenient: synthetic IDs frequently 4xx and the
            // list-response shape varies wildly across public APIs. Skip
            // rather than fail when the call doesn't return a usable list.
            if ($err !== null) {
                $this->markTestSkipped("list call failed (likely synthetic IDs against live API): " . (string)$err);
                return;
            }
            if (empty($result["ok"])) {
                $this->markTestSkipped("list call not ok (likely synthetic IDs against live API)");
                return;
            }
            $status = Helpers::to_int($result["status"]);
            if ($status < 200 || $status >= 300) {
                $this->markTestSkipped("expected 2xx status, got " . $status);
                return;
            }
        } else {
            $this->assertNull($err);
            $this->assertTrue($result["ok"]);
            $this->assertEquals(200, Helpers::to_int($result["status"]));
            $this->assertIsArray($result["data"]);
            $this->assertCount(2, $result["data"]);
            $this->assertCount(1, $setup["calls"]);
        }
    }

}


function combined_search_with_field_direct_setup($mockres)
{
    Runner::load_env_local();

    $calls = new \ArrayObject();

    $env = Runner::env_override([
        "POETRYDB_TEST_COMBINED_SEARCH_WITH_FIELD_ENTID" => [],
        "POETRYDB_TEST_LIVE" => "FALSE",
        "POETRYDB_APIKEY" => "NONE",
    ]);

    $live = $env["POETRYDB_TEST_LIVE"] === "TRUE";

    if ($live) {
        $merged_opts = [
            "apikey" => $env["POETRYDB_APIKEY"],
        ];
        $client = new PoetrydbSDK($merged_opts);
        return [
            "client" => $client,
            "calls" => $calls,
            "live" => true,
            "idmap" => [],
        ];
    }

    $mock_fetch = function ($url, $init) use ($calls, $mockres) {
        $calls[] = ["url" => $url, "init" => $init];
        return [
            [
                "status" => 200,
                "statusText" => "OK",
                "headers" => [],
                "json" => function () use ($mockres) {
                    if ($mockres !== null) {
                        return $mockres;
                    }
                    return ["id" => "direct01"];
                },
                "body" => "mock",
            ],
            null,
        ];
    };

    $client = new PoetrydbSDK([
        "base" => "http://localhost:8080",
        "system" => [
            "fetch" => $mock_fetch,
        ],
    ]);

    return [
        "client" => $client,
        "calls" => $calls,
        "live" => false,
        "idmap" => [],
    ];
}
