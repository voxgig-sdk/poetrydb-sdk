# CombinedSearchWithField entity test

import json
import os
import time

import pytest

from utility.voxgig_struct import voxgig_struct as vs
from poetrydb_sdk import PoetrydbSDK
from core import helpers

_TEST_DIR = os.path.dirname(os.path.abspath(__file__))
from test import runner


class TestCombinedSearchWithFieldEntity:

    def test_should_create_instance(self):
        testsdk = PoetrydbSDK.test(None, None)
        ent = testsdk.CombinedSearchWithField(None)
        assert ent is not None

    def test_should_run_basic_flow(self):
        setup = _combined_search_with_field_basic_setup(None)
        # Per-op sdk-test-control.json skip — basic test exercises a flow with
        # multiple ops; skipping any one skips the whole flow (steps depend
        # on each other).
        _live = setup.get("live", False)
        for _op in ["list"]:
            _skip, _reason = runner.is_control_skipped("entityOp", "combined_search_with_field." + _op, "live" if _live else "unit")
            if _skip:
                pytest.skip(_reason or "skipped via sdk-test-control.json")
                return
        # The basic flow consumes synthetic IDs from the fixture. In live mode
        # without an *_ENTID env override, those IDs hit the live API and 4xx.
        if setup.get("synthetic_only"):
            pytest.skip("live entity test uses synthetic IDs from fixture — "
                        "set POETRYDB_TEST_COMBINED_SEARCH_WITH_FIELD_ENTID JSON to run live")
        client = setup["client"]

        # Bootstrap entity data from existing test data.
        combined_search_with_field_ref01_data_raw = vs.items(helpers.to_map(
            vs.getpath(setup["data"], "existing.combined_search_with_field")))
        combined_search_with_field_ref01_data = None
        if len(combined_search_with_field_ref01_data_raw) > 0:
            combined_search_with_field_ref01_data = helpers.to_map(combined_search_with_field_ref01_data_raw[0][1])

        # LIST
        combined_search_with_field_ref01_ent = client.CombinedSearchWithField(None)
        combined_search_with_field_ref01_match = {
            "input_field1": setup["idmap"]["input_field101"],
            "input_field2": setup["idmap"]["input_field201"],
            "output_field": setup["idmap"]["output_field01"],
            "search_term1": setup["idmap"]["search_term101"],
            "search_term2": setup["idmap"]["search_term201"],
        }

        combined_search_with_field_ref01_list_result = combined_search_with_field_ref01_ent.list(combined_search_with_field_ref01_match, None)
        assert isinstance(combined_search_with_field_ref01_list_result, list)



def _combined_search_with_field_basic_setup(extra):
    runner.load_env_local()

    entity_data_file = os.path.join(_TEST_DIR, "../../.sdk/test/entity/combined_search_with_field/CombinedSearchWithFieldTestData.json")
    with open(entity_data_file, "r") as f:
        entity_data_source = f.read()

    entity_data = json.loads(entity_data_source)

    options = {}
    options["entity"] = entity_data.get("existing")

    client = PoetrydbSDK.test(options, extra)

    # Generate idmap via transform.
    idmap = vs.transform(
        ["combined_search_with_field01", "combined_search_with_field02", "combined_search_with_field03", "input_field101", "input_field201", "output_field01", "search_term101", "search_term201"],
        {
            "`$PACK`": ["", {
                "`$KEY`": "`$COPY`",
                "`$VAL`": ["`$FORMAT`", "upper", "`$COPY`"],
            }],
        }
    )

    # Detect ENTID env override before envOverride consumes it. When live
    # mode is on without a real override, the basic test runs against synthetic
    # IDs from the fixture and 4xx's. We surface this so the test can skip.
    _entid_env_raw = os.environ.get(
        "POETRYDB_TEST_COMBINED_SEARCH_WITH_FIELD_ENTID")
    _idmap_overridden = _entid_env_raw is not None and _entid_env_raw.strip().startswith("{")

    env = runner.env_override({
        "POETRYDB_TEST_COMBINED_SEARCH_WITH_FIELD_ENTID": idmap,
        "POETRYDB_TEST_LIVE": "FALSE",
        "POETRYDB_TEST_EXPLAIN": "FALSE",
    })

    idmap_resolved = helpers.to_map(
        env.get("POETRYDB_TEST_COMBINED_SEARCH_WITH_FIELD_ENTID"))
    if idmap_resolved is None:
        idmap_resolved = helpers.to_map(idmap)

    if env.get("POETRYDB_TEST_LIVE") == "TRUE":
        merged_opts = vs.merge([
            {
            },
            extra or {},
        ])
        client = PoetrydbSDK(helpers.to_map(merged_opts))

    _live = env.get("POETRYDB_TEST_LIVE") == "TRUE"
    return {
        "client": client,
        "data": entity_data,
        "idmap": idmap_resolved,
        "env": env,
        "explain": env.get("POETRYDB_TEST_EXPLAIN") == "TRUE",
        "live": _live,
        "synthetic_only": _live and not _idmap_overridden,
        "now": int(time.time() * 1000),
    }
