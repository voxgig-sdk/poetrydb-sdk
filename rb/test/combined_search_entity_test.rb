# CombinedSearch entity test

require "minitest/autorun"
require "json"
require_relative "../Poetrydb_sdk"
require_relative "runner"

class CombinedSearchEntityTest < Minitest::Test
  def test_create_instance
    testsdk = PoetrydbSDK.test(nil, nil)
    ent = testsdk.CombinedSearch(nil)
    assert !ent.nil?
  end

  def test_basic_flow
    setup = combined_search_basic_setup(nil)
    # Per-op sdk-test-control.json skip.
    _live = setup[:live] || false
    ["list"].each do |_op|
      _should_skip, _reason = Runner.is_control_skipped("entityOp", "combined_search." + _op, _live ? "live" : "unit")
      if _should_skip
        skip(_reason || "skipped via sdk-test-control.json")
        return
      end
    end
    # The basic flow consumes synthetic IDs from the fixture. In live mode
    # without an *_ENTID env override, those IDs hit the live API and 4xx.
    if setup[:synthetic_only]
      skip "live entity test uses synthetic IDs from fixture — set POETRYDB_TEST_COMBINED_SEARCH_ENTID JSON to run live"
      return
    end
    client = setup[:client]

    # Bootstrap entity data from existing test data.
    combined_search_ref01_data_raw = Vs.items(Helpers.to_map(
      Vs.getpath(setup[:data], "existing.combined_search")))
    combined_search_ref01_data = nil
    if combined_search_ref01_data_raw.length > 0
      combined_search_ref01_data = Helpers.to_map(combined_search_ref01_data_raw[0][1])
    end

    # LIST
    combined_search_ref01_ent = client.CombinedSearch(nil)
    combined_search_ref01_match = {
      "input_field1" => setup[:idmap]["input_field101"],
      "input_field2" => setup[:idmap]["input_field201"],
      "search_term1" => setup[:idmap]["search_term101"],
      "search_term2" => setup[:idmap]["search_term201"],
    }

    combined_search_ref01_list_result, err = combined_search_ref01_ent.list(combined_search_ref01_match, nil)
    assert_nil err
    assert combined_search_ref01_list_result.is_a?(Array)

  end
end

def combined_search_basic_setup(extra)
  Runner.load_env_local

  entity_data_file = File.join(__dir__, "..", "..", ".sdk", "test", "entity", "combined_search", "CombinedSearchTestData.json")
  entity_data_source = File.read(entity_data_file)
  entity_data = JSON.parse(entity_data_source)

  options = {}
  options["entity"] = entity_data["existing"]

  client = PoetrydbSDK.test(options, extra)

  # Generate idmap via transform.
  idmap = Vs.transform(
    ["combined_search01", "combined_search02", "combined_search03", "input_field101", "input_field201", "search_term101", "search_term201"],
    {
      "`$PACK`" => ["", {
        "`$KEY`" => "`$COPY`",
        "`$VAL`" => ["`$FORMAT`", "upper", "`$COPY`"],
      }],
    }
  )

  # Detect ENTID env override before envOverride consumes it. When live
  # mode is on without a real override, the basic test runs against synthetic
  # IDs from the fixture and 4xx's. Surface this so the test can skip.
  entid_env_raw = ENV["POETRYDB_TEST_COMBINED_SEARCH_ENTID"]
  idmap_overridden = !entid_env_raw.nil? && entid_env_raw.strip.start_with?("{")

  env = Runner.env_override({
    "POETRYDB_TEST_COMBINED_SEARCH_ENTID" => idmap,
    "POETRYDB_TEST_LIVE" => "FALSE",
    "POETRYDB_TEST_EXPLAIN" => "FALSE",
  })

  idmap_resolved = Helpers.to_map(
    env["POETRYDB_TEST_COMBINED_SEARCH_ENTID"])
  if idmap_resolved.nil?
    idmap_resolved = Helpers.to_map(idmap)
  end

  if env["POETRYDB_TEST_LIVE"] == "TRUE"
    merged_opts = Vs.merge([
      {
      },
      extra || {},
    ])
    client = PoetrydbSDK.new(Helpers.to_map(merged_opts))
  end

  live = env["POETRYDB_TEST_LIVE"] == "TRUE"
  {
    client: client,
    data: entity_data,
    idmap: idmap_resolved,
    env: env,
    explain: env["POETRYDB_TEST_EXPLAIN"] == "TRUE",
    live: live,
    synthetic_only: live && !idmap_overridden,
    now: (Time.now.to_f * 1000).to_i,
  }
end
