package sdktest

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	sdk "github.com/voxgig-sdk/poetrydb-sdk/go"
	"github.com/voxgig-sdk/poetrydb-sdk/go/core"

	vs "github.com/voxgig-sdk/poetrydb-sdk/go/utility/struct"
)

func TestCombinedSearchEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.CombinedSearch(nil)
		if ent == nil {
			t.Fatal("expected non-nil CombinedSearchEntity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := combined_searchBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"list"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "combined_search." + _op, _mode); _shouldSkip {
				if _reason == "" {
					_reason = "skipped via sdk-test-control.json"
				}
				t.Skip(_reason)
				return
			}
		}
		// The basic flow consumes synthetic IDs from the fixture. In live mode
		// without an *_ENTID env override, those IDs hit the live API and 4xx.
		if setup.syntheticOnly {
			t.Skip("live entity test uses synthetic IDs from fixture — set POETRYDB_TEST_COMBINED_SEARCH_ENTID JSON to run live")
			return
		}
		client := setup.client

		// Bootstrap entity data from existing test data (no create step in flow).
		combinedSearchRef01DataRaw := vs.Items(core.ToMapAny(vs.GetPath("existing.combined_search", setup.data)))
		var combinedSearchRef01Data map[string]any
		if len(combinedSearchRef01DataRaw) > 0 {
			combinedSearchRef01Data = core.ToMapAny(combinedSearchRef01DataRaw[0][1])
		}
		// Discard guards against Go's unused-var check when the flow's steps
		// happen not to consume the bootstrap data (e.g. list-only flows).
		_ = combinedSearchRef01Data

		// LIST
		combinedSearchRef01Ent := client.CombinedSearch(nil)
		combinedSearchRef01Match := map[string]any{
			"input_field1": setup.idmap["input_field101"],
			"input_field2": setup.idmap["input_field201"],
			"search_term1": setup.idmap["search_term101"],
			"search_term2": setup.idmap["search_term201"],
		}

		combinedSearchRef01ListResult, err := combinedSearchRef01Ent.List(combinedSearchRef01Match, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		_, combinedSearchRef01ListOk := combinedSearchRef01ListResult.([]any)
		if !combinedSearchRef01ListOk {
			t.Fatalf("expected list result to be an array, got %T", combinedSearchRef01ListResult)
		}

	})
}

func combined_searchBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "combined_search", "CombinedSearchTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read combined_search test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse combined_search test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"combined_search01", "combined_search02", "combined_search03", "input_field101", "input_field201", "search_term101", "search_term201"},
		map[string]any{
			"`$PACK`": []any{"", map[string]any{
				"`$KEY`": "`$COPY`",
				"`$VAL`": []any{"`$FORMAT`", "upper", "`$COPY`"},
			}},
		},
	)

	// Detect ENTID env override before envOverride consumes it. When live
	// mode is on without a real override, the basic test runs against synthetic
	// IDs from the fixture and 4xx's. Surface this so the test can skip.
	entidEnvRaw := os.Getenv("POETRYDB_TEST_COMBINED_SEARCH_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"POETRYDB_TEST_COMBINED_SEARCH_ENTID": idmap,
		"POETRYDB_TEST_LIVE":      "FALSE",
		"POETRYDB_TEST_EXPLAIN":   "FALSE",
		"POETRYDB_APIKEY":         "NONE",
	})

	idmapResolved := core.ToMapAny(env["POETRYDB_TEST_COMBINED_SEARCH_ENTID"])
	if idmapResolved == nil {
		idmapResolved = core.ToMapAny(idmap)
	}

	if env["POETRYDB_TEST_LIVE"] == "TRUE" {
		mergedOpts := vs.Merge([]any{
			map[string]any{
				"apikey": env["POETRYDB_APIKEY"],
			},
			extra,
		})
		client = sdk.NewPoetrydbSDK(core.ToMapAny(mergedOpts))
	}

	live := env["POETRYDB_TEST_LIVE"] == "TRUE"
	return &entityTestSetup{
		client:        client,
		data:          entityData,
		idmap:         idmapResolved,
		env:           env,
		explain:       env["POETRYDB_TEST_EXPLAIN"] == "TRUE",
		live:          live,
		syntheticOnly: live && !idmapOverridden,
		now:           time.Now().UnixMilli(),
	}
}
