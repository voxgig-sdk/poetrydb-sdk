<?php
declare(strict_types=1);

// Poemcount entity test

require_once __DIR__ . '/../poetrydb_sdk.php';
require_once __DIR__ . '/Runner.php';

use PHPUnit\Framework\TestCase;
use Voxgig\Struct\Struct as Vs;

class PoemcountEntityTest extends TestCase
{
    public function test_create_instance(): void
    {
        $testsdk = PoetrydbSDK::test(null, null);
        $ent = $testsdk->Poemcount(null);
        $this->assertNotNull($ent);
    }

    public function test_basic_flow(): void
    {
        $setup = poemcount_basic_setup(null);
        // Per-op sdk-test-control.json skip.
        $_live = !empty($setup["live"]);
        foreach (["load"] as $_op) {
            [$_shouldSkip, $_reason] = Runner::is_control_skipped("entityOp", "poemcount." . $_op, $_live ? "live" : "unit");
            if ($_shouldSkip) {
                $this->markTestSkipped($_reason ?? "skipped via sdk-test-control.json");
                return;
            }
        }
        // The basic flow consumes synthetic IDs from the fixture. In live mode
        // without an *_ENTID env override, those IDs hit the live API and 4xx.
        if (!empty($setup["synthetic_only"])) {
            $this->markTestSkipped("live entity test uses synthetic IDs from fixture — set POETRYDB_TEST_POEMCOUNT_ENTID JSON to run live");
            return;
        }
        $client = $setup["client"];

        // Bootstrap entity data from existing test data.
        $poemcount_ref01_data_raw = Vs::items(Helpers::to_map(
            Vs::getpath($setup["data"], "existing.poemcount")));
        $poemcount_ref01_data = null;
        if (count($poemcount_ref01_data_raw) > 0) {
            $poemcount_ref01_data = Helpers::to_map($poemcount_ref01_data_raw[0][1]);
        }

        // LOAD
        $poemcount_ref01_ent = $client->Poemcount(null);
        $poemcount_ref01_match_dt0 = [];
        [$poemcount_ref01_data_dt0_loaded, $err] = $poemcount_ref01_ent->load($poemcount_ref01_match_dt0, null);
        $this->assertNull($err);
        $this->assertNotNull($poemcount_ref01_data_dt0_loaded);

    }
}

function poemcount_basic_setup($extra)
{
    Runner::load_env_local();

    $entity_data_file = __DIR__ . '/../../.sdk/test/entity/poemcount/PoemcountTestData.json';
    $entity_data_source = file_get_contents($entity_data_file);
    $entity_data = json_decode($entity_data_source, true);

    $options = [];
    $options["entity"] = $entity_data["existing"];

    $client = PoetrydbSDK::test($options, $extra);

    // Generate idmap.
    $idmap = [];
    foreach (["poemcount01", "poemcount02", "poemcount03"] as $k) {
        $idmap[$k] = strtoupper($k);
    }

    // Detect ENTID env override before envOverride consumes it. When live
    // mode is on without a real override, the basic test runs against synthetic
    // IDs from the fixture and 4xx's. Surface this so the test can skip.
    $entid_env_raw = getenv("POETRYDB_TEST_POEMCOUNT_ENTID");
    $idmap_overridden = $entid_env_raw !== false && str_starts_with(trim($entid_env_raw), "{");

    $env = Runner::env_override([
        "POETRYDB_TEST_POEMCOUNT_ENTID" => $idmap,
        "POETRYDB_TEST_LIVE" => "FALSE",
        "POETRYDB_TEST_EXPLAIN" => "FALSE",
    ]);

    $idmap_resolved = Helpers::to_map(
        $env["POETRYDB_TEST_POEMCOUNT_ENTID"]);
    if ($idmap_resolved === null) {
        $idmap_resolved = Helpers::to_map($idmap);
    }

    if ($env["POETRYDB_TEST_LIVE"] === "TRUE") {
        $merged_opts = Vs::merge([
            [
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
