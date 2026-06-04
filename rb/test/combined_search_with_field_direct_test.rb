# CombinedSearchWithField direct test

require "minitest/autorun"
require "json"
require_relative "../Poetrydb_sdk"
require_relative "runner"

class CombinedSearchWithFieldDirectTest < Minitest::Test
  def test_direct_list_combined_search_with_field
    setup = combined_search_with_field_direct_setup([
      { "id" => "direct01" },
      { "id" => "direct02" },
    ])
    _should_skip, _reason = Runner.is_control_skipped("direct", "direct-list-combined_search_with_field", setup[:live] ? "live" : "unit")
    if _should_skip
      skip(_reason || "skipped via sdk-test-control.json")
      return
    end
    if setup[:live]
      ["input_field101", "input_field201", "output_field01", "search_term101", "search_term201"].each do |_live_key|
        if setup[:idmap][_live_key].nil?
          skip "live test needs #{_live_key} via *_ENTID env var (synthetic IDs only)"
          return
        end
      end
    end
    client = setup[:client]

    params = {}
    if setup[:live]
      params["input_field1"] = setup[:idmap]["input_field101"]
    else
      params["input_field1"] = "direct01"
    end
    if setup[:live]
      params["input_field2"] = setup[:idmap]["input_field201"]
    else
      params["input_field2"] = "direct01"
    end
    if setup[:live]
      params["output_field"] = setup[:idmap]["output_field01"]
    else
      params["output_field"] = "direct01"
    end
    if setup[:live]
      params["search_term1"] = setup[:idmap]["search_term101"]
    else
      params["search_term1"] = "direct01"
    end
    if setup[:live]
      params["search_term2"] = setup[:idmap]["search_term201"]
    else
      params["search_term2"] = "direct01"
    end

    result, err = client.direct({
      "path" => "{input_field1},{input_field2}/{search_term1};{search_term2}/{output_field}",
      "method" => "GET",
      "params" => params,
    })
    if setup[:live]
      # Live mode is lenient: synthetic IDs frequently 4xx and the list-
      # response shape varies wildly across public APIs. Skip rather than
      # fail when the call doesn't return a usable list.
      if !err.nil?
        skip("list call failed (likely synthetic IDs against live API): #{err}")
        return
      end
      unless result["ok"]
        skip("list call not ok (likely synthetic IDs against live API)")
        return
      end
      status = Helpers.to_int(result["status"])
      if status < 200 || status >= 300
        skip("expected 2xx status, got #{status}")
        return
      end
    else
      assert_nil err
      assert result["ok"]
      assert_equal 200, Helpers.to_int(result["status"])
      assert result["data"].is_a?(Array)
      assert_equal 2, result["data"].length
      assert_equal 1, setup[:calls].length
    end
  end

end


def combined_search_with_field_direct_setup(mockres)
  Runner.load_env_local

  calls = []

  env = Runner.env_override({
    "POETRYDB_TEST_COMBINED_SEARCH_WITH_FIELD_ENTID" => {},
    "POETRYDB_TEST_LIVE" => "FALSE",
  })

  live = env["POETRYDB_TEST_LIVE"] == "TRUE"

  if live
    merged_opts = {
    }
    client = PoetrydbSDK.new(merged_opts)
    return {
      client: client,
      calls: calls,
      live: true,
      idmap: {},
    }
  end

  mock_fetch = ->(url, init) {
    calls.push({ "url" => url, "init" => init })
    return {
      "status" => 200,
      "statusText" => "OK",
      "headers" => {},
      "json" => ->() {
        if !mockres.nil?
          return mockres
        end
        return { "id" => "direct01" }
      },
      "body" => "mock",
    }, nil
  }

  client = PoetrydbSDK.new({
    "base" => "http://localhost:8080",
    "system" => {
      "fetch" => mock_fetch,
    },
  })

  {
    client: client,
    calls: calls,
    live: false,
    idmap: {},
  }
end
