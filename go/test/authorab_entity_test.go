package sdktest

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	sdk "github.com/voxgig-sdk/poetrydb-sdk"
	"github.com/voxgig-sdk/poetrydb-sdk/core"

	vs "github.com/voxgig/struct"
)

func TestAuthorabEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.Authorab(nil)
		if ent == nil {
			t.Fatal("expected non-nil AuthorabEntity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := authorabBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"list"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "authorab." + _op, _mode); _shouldSkip {
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
			t.Skip("live entity test uses synthetic IDs from fixture — set POETRYDB_TEST_AUTHORAB_ENTID JSON to run live")
			return
		}
		client := setup.client

		// Bootstrap entity data from existing test data (no create step in flow).
		authorabRef01DataRaw := vs.Items(core.ToMapAny(vs.GetPath("existing.authorab", setup.data)))
		var authorabRef01Data map[string]any
		if len(authorabRef01DataRaw) > 0 {
			authorabRef01Data = core.ToMapAny(authorabRef01DataRaw[0][1])
		}
		// Discard guards against Go's unused-var check when the flow's steps
		// happen not to consume the bootstrap data (e.g. list-only flows).
		_ = authorabRef01Data

		// LIST
		authorabRef01Ent := client.Authorab(nil)
		authorabRef01Match := map[string]any{
			"author": setup.idmap["author01"],
		}

		authorabRef01ListResult, err := authorabRef01Ent.List(authorabRef01Match, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		_, authorabRef01ListOk := authorabRef01ListResult.([]any)
		if !authorabRef01ListOk {
			t.Fatalf("expected list result to be an array, got %T", authorabRef01ListResult)
		}

	})
}

func authorabBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "authorab", "AuthorabTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read authorab test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse authorab test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"authorab01", "authorab02", "authorab03", "author01", "author02", "author03"},
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
	entidEnvRaw := os.Getenv("POETRYDB_TEST_AUTHORAB_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"POETRYDB_TEST_AUTHORAB_ENTID": idmap,
		"POETRYDB_TEST_LIVE":      "FALSE",
		"POETRYDB_TEST_EXPLAIN":   "FALSE",
		"POETRYDB_APIKEY":         "NONE",
	})

	idmapResolved := core.ToMapAny(env["POETRYDB_TEST_AUTHORAB_ENTID"])
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
