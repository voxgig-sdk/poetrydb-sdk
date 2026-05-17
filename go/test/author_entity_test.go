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

func TestAuthorEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.Author(nil)
		if ent == nil {
			t.Fatal("expected non-nil AuthorEntity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := authorBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"list", "load"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "author." + _op, _mode); _shouldSkip {
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
			t.Skip("live entity test uses synthetic IDs from fixture — set POETRYDB_TEST_AUTHOR_ENTID JSON to run live")
			return
		}
		client := setup.client

		// Bootstrap entity data from existing test data (no create step in flow).
		authorRef01DataRaw := vs.Items(core.ToMapAny(vs.GetPath("existing.author", setup.data)))
		var authorRef01Data map[string]any
		if len(authorRef01DataRaw) > 0 {
			authorRef01Data = core.ToMapAny(authorRef01DataRaw[0][1])
		}
		// Discard guards against Go's unused-var check when the flow's steps
		// happen not to consume the bootstrap data (e.g. list-only flows).
		_ = authorRef01Data

		// LIST
		authorRef01Ent := client.Author(nil)
		authorRef01Match := map[string]any{}

		authorRef01ListResult, err := authorRef01Ent.List(authorRef01Match, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		_, authorRef01ListOk := authorRef01ListResult.([]any)
		if !authorRef01ListOk {
			t.Fatalf("expected list result to be an array, got %T", authorRef01ListResult)
		}

		// LOAD
		authorRef01MatchDt0 := map[string]any{}
		authorRef01DataDt0Loaded, err := authorRef01Ent.Load(authorRef01MatchDt0, nil)
		if err != nil {
			t.Fatalf("load failed: %v", err)
		}
		if authorRef01DataDt0Loaded == nil {
			t.Fatal("expected load result to be non-nil")
		}

	})
}

func authorBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "author", "AuthorTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read author test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse author test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"author01", "author02", "author03"},
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
	entidEnvRaw := os.Getenv("POETRYDB_TEST_AUTHOR_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"POETRYDB_TEST_AUTHOR_ENTID": idmap,
		"POETRYDB_TEST_LIVE":      "FALSE",
		"POETRYDB_TEST_EXPLAIN":   "FALSE",
		"POETRYDB_APIKEY":         "NONE",
	})

	idmapResolved := core.ToMapAny(env["POETRYDB_TEST_AUTHOR_ENTID"])
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
