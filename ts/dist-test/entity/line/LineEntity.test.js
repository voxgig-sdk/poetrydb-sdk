"use strict";
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    var desc = Object.getOwnPropertyDescriptor(m, k);
    if (!desc || ("get" in desc ? !m.__esModule : desc.writable || desc.configurable)) {
      desc = { enumerable: true, get: function() { return m[k]; } };
    }
    Object.defineProperty(o, k2, desc);
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __setModuleDefault = (this && this.__setModuleDefault) || (Object.create ? (function(o, v) {
    Object.defineProperty(o, "default", { enumerable: true, value: v });
}) : function(o, v) {
    o["default"] = v;
});
var __importStar = (this && this.__importStar) || (function () {
    var ownKeys = function(o) {
        ownKeys = Object.getOwnPropertyNames || function (o) {
            var ar = [];
            for (var k in o) if (Object.prototype.hasOwnProperty.call(o, k)) ar[ar.length] = k;
            return ar;
        };
        return ownKeys(o);
    };
    return function (mod) {
        if (mod && mod.__esModule) return mod;
        var result = {};
        if (mod != null) for (var k = ownKeys(mod), i = 0; i < k.length; i++) if (k[i] !== "default") __createBinding(result, mod, k[i]);
        __setModuleDefault(result, mod);
        return result;
    };
})();
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const node_path_1 = __importDefault(require("node:path"));
const Fs = __importStar(require("node:fs"));
const node_test_1 = require("node:test");
const node_assert_1 = __importDefault(require("node:assert"));
const live_runner_1 = require("../../live-runner");
const live_entity_1 = require("../../live-entity");
const __1 = require("../../..");
const utility_1 = require("../../utility");
// AFTER the imports on purpose: TypeScript hoists `import` above any
// statement in the emitted CommonJS, so a loader placed above them would
// run only after every imported module had already been evaluated - and
// anything reading process.env at module scope would miss these values.
(0, utility_1.loadEnvLocal)(__dirname + '/../../../.env.local');
(0, node_test_1.describe)('LineEntity', async () => {
    // Per-test live pacing. Delay is read from sdk-test-control.json's
    // `test.live.delayMs`; only sleeps when POETRYDB_TEST_LIVE=TRUE.
    (0, node_test_1.afterEach)((0, utility_1.liveDelay)('POETRYDB_TEST_LIVE'));
    (0, node_test_1.test)('instance', async () => {
        const testsdk = __1.PoetrydbSDK.test();
        const ent = testsdk.Line();
        (0, node_assert_1.default)(null != ent);
    });
    (0, node_test_1.test)('basic', async (t) => {
        const live = 'TRUE' === process.env.POETRYDB_TEST_LIVE;
        for (const op of ['list', 'load']) {
            if (!live && (0, utility_1.maybeSkipControl)(t, 'entityOp', 'line.' + op, live))
                return;
        }
        const setup = basicSetup();
        if (setup.live) {
            return (0, live_entity_1.runLiveEntity)(setup, { "active": true, "alias": { "field": {} }, "fields": [{ "active": true, "name": "author", "req": false, "short": "The author of the poem", "type": "`$STRING`", "index$": 0 }, { "active": true, "name": "id", "req": false, "type": "`$STRING`", "index$": 1 }, { "active": true, "name": "linecount", "req": false, "short": "The number of lines in the poem (including section headings, excluding empty lines)", "type": "`$INTEGER`", "index$": 2 }, { "active": true, "name": "lines", "req": false, "short": "The lines of the poem", "type": "`$ARRAY`", "index$": 3 }, { "active": true, "name": "title", "req": false, "short": "The title of the poem", "type": "`$STRING`", "index$": 4 }], "id": { "field": "id", "name": "id" }, "name": "line", "op": { "list": { "input": "data", "name": "list", "points": [{ "active": true, "args": { "params": [{ "active": true, "example": "text", "kind": "param", "name": "format", "orig": "format", "reqd": true, "type": "`$STRING`", "index$": 0 }, { "active": true, "example": "Latitudeless Place", "kind": "param", "name": "line", "orig": "line", "reqd": true, "type": "`$STRING`", "index$": 1 }, { "active": true, "example": "author,title,linecount", "kind": "param", "name": "output_field", "orig": "output_field", "reqd": true, "type": "`$STRING`", "index$": 2 }] }, "contract": { "id": "GET /lines/{lines}/{outputFields}.{format}", "json": "{\"operationId\":\"searchByLinesWithFieldsAndFormat\",\"parameters\":[{\"description\":\"Part of a line or lines of a poem\",\"example\":\"Latitudeless Place\",\"in\":\"path\",\"name\":\"lines\",\"required\":true,\"schema\":{\"type\":\"string\"}},{\"description\":\"Comma-separated list of output fields (author, title, lines, linecount, all)\",\"example\":\"author,title,linecount\",\"in\":\"path\",\"name\":\"outputFields\",\"required\":true,\"schema\":{\"type\":\"string\"}},{\"description\":\"Output format (json or text)\",\"example\":\"text\",\"in\":\"path\",\"name\":\"format\",\"required\":true,\"schema\":{\"enum\":[\"json\",\"text\"],\"type\":\"string\"}}],\"protocol\":\"http\",\"responses\":{\"200\":{\"content\":{\"application/json\":{\"schema\":{\"items\":{\"type\":\"object\"},\"type\":\"array\"}},\"text/plain\":{\"schema\":{\"type\":\"string\"}}},\"description\":\"Successful response\"},\"404\":{\"content\":{\"application/json\":{\"schema\":{\"properties\":{\"reason\":{\"example\":\"Not found\",\"type\":\"string\"},\"status\":{\"example\":404,\"type\":\"integer\"}},\"type\":\"object\"}}},\"description\":\"No poems found\"}},\"securitySource\":\"unspecified\"}", "source": "openapi3", "version": 1 }, "kind": "http", "method": "GET", "orig": "/lines/{lines}/{outputFields}.{format}", "rename": { "param": { "lines": "line", "outputFields}.{format": "output_fields}_{format" } }, "segments": [{ "lit": "lines" }, { "var": "line" }, { "lit": "{outputFields}.{format}" }], "select": { "exist": ["format", "line", "output_field"] }, "transform": { "req": "`reqdata`", "res": "`body`" }, "index$": 0 }, { "active": true, "args": { "params": [{ "active": true, "example": "Latitudeless Place", "kind": "param", "name": "line", "orig": "line", "reqd": true, "type": "`$STRING`", "index$": 0 }, { "active": true, "example": "author,title,linecount", "kind": "param", "name": "output_field", "orig": "output_field", "reqd": true, "type": "`$STRING`", "index$": 1 }] }, "contract": { "id": "GET /lines/{lines}/{outputFields}", "json": "{\"operationId\":\"searchByLinesWithFields\",\"parameters\":[{\"description\":\"Part of a line or lines of a poem\",\"example\":\"Latitudeless Place\",\"in\":\"path\",\"name\":\"lines\",\"required\":true,\"schema\":{\"type\":\"string\"}},{\"description\":\"Comma-separated list of output fields (author, title, lines, linecount, all)\",\"example\":\"author,title,linecount\",\"in\":\"path\",\"name\":\"outputFields\",\"required\":true,\"schema\":{\"type\":\"string\"}}],\"protocol\":\"http\",\"responses\":{\"200\":{\"content\":{\"application/json\":{\"schema\":{\"items\":{\"type\":\"object\"},\"type\":\"array\"}}},\"description\":\"Successful response\"},\"404\":{\"content\":{\"application/json\":{\"schema\":{\"properties\":{\"reason\":{\"example\":\"Not found\",\"type\":\"string\"},\"status\":{\"example\":404,\"type\":\"integer\"}},\"type\":\"object\"}}},\"description\":\"No poems found\"}},\"securitySource\":\"unspecified\"}", "source": "openapi3", "version": 1 }, "kind": "http", "method": "GET", "orig": "/lines/{lines}/{outputFields}", "rename": { "param": { "lines": "line", "outputFields": "output_field" } }, "segments": [{ "lit": "lines" }, { "var": "line" }, { "var": "output_field" }], "select": { "exist": ["line", "output_field"] }, "transform": { "req": "`reqdata`", "res": "`body`" }, "index$": 1 }], "key$": "list" }, "load": { "input": "data", "name": "load", "points": [{ "active": true, "args": { "params": [{ "active": true, "example": "Latitudeless Place", "kind": "param", "name": "id", "orig": "line", "reqd": true, "type": "`$STRING`", "index$": 0 }] }, "contract": { "id": "GET /lines/{lines}", "json": "{\"operationId\":\"searchByLines\",\"parameters\":[{\"description\":\"Part of a line or lines of a poem\",\"example\":\"Latitudeless Place\",\"in\":\"path\",\"name\":\"lines\",\"required\":true,\"schema\":{\"type\":\"string\"}}],\"protocol\":\"http\",\"responses\":{\"200\":{\"content\":{\"application/json\":{\"schema\":{\"items\":{\"properties\":{\"author\":{\"description\":\"The author of the poem\",\"example\":\"Percy Bysshe Shelley\",\"type\":\"string\"},\"linecount\":{\"description\":\"The number of lines in the poem (including section headings, excluding empty lines)\",\"example\":14,\"type\":\"integer\"},\"lines\":{\"description\":\"The lines of the poem\",\"example\":[\"I met a traveller from an antique land\",\"Who said: \\\"Two vast and trunkless legs of stone\",\"Stand in the desert. Near them on the sand,\"],\"items\":{\"type\":\"string\"},\"type\":\"array\"},\"title\":{\"description\":\"The title of the poem\",\"example\":\"Ozymandias\",\"type\":\"string\"}},\"type\":\"object\"},\"type\":\"array\"}}},\"description\":\"Successful response\"},\"404\":{\"content\":{\"application/json\":{\"schema\":{\"properties\":{\"reason\":{\"example\":\"Not found\",\"type\":\"string\"},\"status\":{\"example\":404,\"type\":\"integer\"}},\"type\":\"object\"}}},\"description\":\"No poems found\"}},\"securitySource\":\"unspecified\"}", "source": "openapi3", "version": 1 }, "kind": "http", "method": "GET", "orig": "/lines/{lines}", "rename": { "param": { "lines": "id" } }, "segments": [{ "lit": "lines" }, { "var": "id" }], "select": { "exist": ["id"] }, "transform": { "req": "`reqdata`", "res": "`body`" }, "index$": 0 }], "key$": "load" } }, "relations": { "ancestors": [["line"]] }, "key$": "line", "name__orig": "line", "Name": "Line", "name_": "line", "name-": "line", "NAME": "LINE", "index$": 4 }, { "active": true, "entity": "line", "key$": "BasicLineFlow", "kind": "basic", "name": "BasicLineFlow", "param": {}, "step": [{ "active": true, "data": {}, "input": {}, "match": { "line": "line01", "output_field": "output_field01" }, "op": "list", "spec": [], "valid": [{ "apply": "ItemExists", "def": { "ref": "line_ref01" } }], "index$": 0 }, { "active": true, "data": {}, "input": { "ref": "line_ref01", "srcdatavar": "line_ref01_data", "suffix": "_dt0" }, "match": { "id": "line01" }, "op": "load", "spec": [], "valid": [{ "apply": "TextFieldMark", "def": { "mark": "Mark01-line_ref01" } }], "index$": 1 }] }, 'Line');
        }
        const client = setup.client;
        const struct = setup.struct;
        const isempty = struct.isempty;
        const select = struct.select;
        let line_ref01_data = Object.values(setup.data.existing.line)[0];
        // LIST
        const line_ref01_ent = client.Line();
        const line_ref01_match = {};
        line_ref01_match['line'] = setup.idmap['line01'];
        line_ref01_match['output_field'] = setup.idmap['output_field01'];
        const line_ref01_list = (await line_ref01_ent.list(line_ref01_match)).map((e) => e.data());
        // LOAD
        const line_ref01_match_dt0 = {};
        line_ref01_match_dt0.id = line_ref01_data.id;
        const line_ref01_data_dt0 = (await line_ref01_ent.load(line_ref01_match_dt0)).data();
        (0, node_assert_1.default)(line_ref01_data_dt0.id === line_ref01_data.id);
    });
});
function basicSetup(extra) {
    // TODO: fix test def options
    const options = {}; // null
    // TODO: needs test utility to resolve path
    const entityDataFile = node_path_1.default.resolve(__dirname, '../../../../.sdk/test/entity/line/LineTestData.json');
    // TODO: file ready util needed?
    const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8');
    // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
    const entityData = JSON.parse(entityDataSource);
    options.entity = entityData.existing;
    let client = __1.PoetrydbSDK.test(options, extra);
    const struct = client.utility().struct;
    const merge = struct.merge;
    const transform = struct.transform;
    let idmap = transform(['line01', 'line02', 'line03', 'line01', 'line02', 'line03'], {
        '`$PACK`': ['', {
                '`$KEY`': '`$COPY`',
                '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
            }]
    });
    const env = (0, utility_1.envOverride)({
        'POETRYDB_TEST_LINE_ENTID': idmap,
        'POETRYDB_TEST_LIVE': 'FALSE',
        'POETRYDB_TEST_EXPLAIN': 'FALSE',
    });
    idmap = env['POETRYDB_TEST_LINE_ENTID'];
    const live = 'TRUE' === env.POETRYDB_TEST_LIVE;
    const transport = (0, live_runner_1.createLiveTransport)();
    if (live) {
        const rawIds = process.env['POETRYDB_TEST_LINE_ENTID'];
        idmap = rawIds && rawIds.trim() ? JSON.parse(rawIds) : {};
        if (!idmap || Array.isArray(idmap) || typeof idmap !== 'object') {
            throw new Error('Live ENTID must be a JSON object');
        }
        client = new __1.PoetrydbSDK(merge([
            // FIRST, so the generated fields below win: sdk-test-control.json's
            // test.client.options adds to the live client, it does not redirect it.
            (0, utility_1.liveClientOptions)(),
            {},
            // 'extra || {}', not a bare 'extra': struct.merge returns UNDEFINED when the
            // last entry is undefined, and basicSetup is normally called with no
            // argument at all - so a bare 'extra' silently discarded the apikey
            // and server values above and handed the SDK undefined. Harmless
            // while there was nothing in that object; not harmless now.
            extra || {},
            { system: { fetch: transport.fetch } }
        ]));
    }
    const setup = {
        idmap,
        env,
        options,
        client,
        struct,
        data: entityData,
        explain: 'TRUE' === env.POETRYDB_TEST_EXPLAIN,
        live,
        transport,
        now: Date.now(),
    };
    return setup;
}
//# sourceMappingURL=LineEntity.test.js.map