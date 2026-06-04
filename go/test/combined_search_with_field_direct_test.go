package sdktest

import (
	"encoding/json"
	"os"
	"strings"
	"testing"

	sdk "github.com/voxgig-sdk/poetrydb-sdk/go"
	"github.com/voxgig-sdk/poetrydb-sdk/go/core"
)

func TestCombinedSearchWithFieldDirect(t *testing.T) {
	t.Run("direct-list-combined_search_with_field", func(t *testing.T) {
		setup := combined_search_with_fieldDirectSetup([]any{
			map[string]any{"id": "direct01"},
			map[string]any{"id": "direct02"},
		})
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		if _shouldSkip, _reason := isControlSkipped("direct", "direct-list-combined_search_with_field", _mode); _shouldSkip {
			if _reason == "" {
				_reason = "skipped via sdk-test-control.json"
			}
			t.Skip(_reason)
			return
		}
		if setup.live {
			for _, _liveKey := range []string{"input_field101", "input_field201", "output_field01", "search_term101", "search_term201"} {
				if v := setup.idmap[_liveKey]; v == nil {
					t.Skipf("live test needs %s via *_ENTID env var (synthetic IDs only)", _liveKey)
					return
				}
			}
		}
		client := setup.client

		params := map[string]any{}
		if setup.live {
			params["input_field1"] = setup.idmap["input_field101"]
		} else {
			params["input_field1"] = "direct01"
		}
		if setup.live {
			params["input_field2"] = setup.idmap["input_field201"]
		} else {
			params["input_field2"] = "direct02"
		}
		if setup.live {
			params["output_field"] = setup.idmap["output_field01"]
		} else {
			params["output_field"] = "direct03"
		}
		if setup.live {
			params["search_term1"] = setup.idmap["search_term101"]
		} else {
			params["search_term1"] = "direct04"
		}
		if setup.live {
			params["search_term2"] = setup.idmap["search_term201"]
		} else {
			params["search_term2"] = "direct05"
		}

		result, err := client.Direct(map[string]any{
			"path":   "{input_field1},{input_field2}/{search_term1};{search_term2}/{output_field}",
			"method": "GET",
			"params": params,
		})
		if setup.live {
			// Live mode is lenient: synthetic IDs frequently 4xx and the
			// list-response shape varies wildly across public APIs. Skip
			// rather than fail when the call doesn't return a usable list.
			if err != nil {
				t.Skipf("list call failed (likely synthetic IDs against live API): %v", err)
			}
			if result["ok"] != true {
				t.Skipf("list call not ok (likely synthetic IDs against live API): %v", result)
			}
			status := core.ToInt(result["status"])
			if status < 200 || status >= 300 {
				t.Skipf("expected 2xx status, got %v", result["status"])
			}
		} else {
			if err != nil {
				t.Fatalf("direct failed: %v", err)
			}
			if result["ok"] != true {
				t.Fatalf("expected ok to be true, got %v", result["ok"])
			}
			if core.ToInt(result["status"]) != 200 {
				t.Fatalf("expected status 200, got %v", result["status"])
			}
		}

		if !setup.live {
			if dataList, ok := result["data"].([]any); ok {
				if len(dataList) != 2 {
					t.Fatalf("expected 2 items, got %d", len(dataList))
				}
			} else {
				t.Fatalf("expected data to be an array, got %T", result["data"])
			}

			if len(*setup.calls) != 1 {
				t.Fatalf("expected 1 call, got %d", len(*setup.calls))
			}
			call := (*setup.calls)[0]
			if initMap, ok := call["init"].(map[string]any); ok {
				if initMap["method"] != "GET" {
					t.Fatalf("expected method GET, got %v", initMap["method"])
				}
			}
			if url, ok := call["url"].(string); ok {
				if !strings.Contains(url, "direct01") {
					t.Fatalf("expected url to contain direct01, got %v", url)
				}
				if !strings.Contains(url, "direct02") {
					t.Fatalf("expected url to contain direct02, got %v", url)
				}
				if !strings.Contains(url, "direct03") {
					t.Fatalf("expected url to contain direct03, got %v", url)
				}
				if !strings.Contains(url, "direct04") {
					t.Fatalf("expected url to contain direct04, got %v", url)
				}
				if !strings.Contains(url, "direct05") {
					t.Fatalf("expected url to contain direct05, got %v", url)
				}
			}
		}
	})

}

type combined_search_with_fieldDirectSetupResult struct {
	client *sdk.PoetrydbSDK
	calls  *[]map[string]any
	live   bool
	idmap  map[string]any
}

func combined_search_with_fieldDirectSetup(mockres any) *combined_search_with_fieldDirectSetupResult {
	loadEnvLocal()

	calls := &[]map[string]any{}

	env := envOverride(map[string]any{
		"POETRYDB_TEST_COMBINED_SEARCH_WITH_FIELD_ENTID": map[string]any{},
		"POETRYDB_TEST_LIVE":    "FALSE",
	})

	live := env["POETRYDB_TEST_LIVE"] == "TRUE"

	if live {
		mergedOpts := map[string]any{
		}
		client := sdk.NewPoetrydbSDK(mergedOpts)

		idmap := map[string]any{}
		if entidRaw, ok := env["POETRYDB_TEST_COMBINED_SEARCH_WITH_FIELD_ENTID"]; ok {
			if entidStr, ok := entidRaw.(string); ok && strings.HasPrefix(entidStr, "{") {
				json.Unmarshal([]byte(entidStr), &idmap)
			} else if entidMap, ok := entidRaw.(map[string]any); ok {
				idmap = entidMap
			}
		}

		return &combined_search_with_fieldDirectSetupResult{client: client, calls: calls, live: true, idmap: idmap}
	}

	mockFetch := func(url string, init map[string]any) (map[string]any, error) {
		*calls = append(*calls, map[string]any{"url": url, "init": init})
		return map[string]any{
			"status":     200,
			"statusText": "OK",
			"headers":    map[string]any{},
			"json": (func() any)(func() any {
				if mockres != nil {
					return mockres
				}
				return map[string]any{"id": "direct01"}
			}),
		}, nil
	}

	client := sdk.NewPoetrydbSDK(map[string]any{
		"base": "http://localhost:8080",
		"system": map[string]any{
			"fetch": (func(string, map[string]any) (map[string]any, error))(mockFetch),
		},
	})

	return &combined_search_with_fieldDirectSetupResult{client: client, calls: calls, live: false, idmap: map[string]any{}}
}

var _ = os.Getenv
var _ = json.Unmarshal
