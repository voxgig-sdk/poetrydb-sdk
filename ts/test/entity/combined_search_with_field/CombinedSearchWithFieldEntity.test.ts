

import Path from 'node:path'
import * as Fs from 'node:fs'

import { test, describe, afterEach } from 'node:test'
import assert from 'node:assert'
import { createLiveTransport } from '../../live-runner'
import { runLiveEntity } from '../../live-entity'


import { PoetrydbSDK, BaseFeature, stdutil } from '../../..'

import {
  envOverride,
  liveClientOptions,
  liveDelay,
  loadEnvLocal,
  makeCtrl,
  makeMatch,
  makeReqdata,
  makeStepData,
  makeValid,
  maybeSkipControl,
} from '../../utility'


// AFTER the imports on purpose: TypeScript hoists `import` above any
// statement in the emitted CommonJS, so a loader placed above them would
// run only after every imported module had already been evaluated - and
// anything reading process.env at module scope would miss these values.
loadEnvLocal(__dirname + '/../../../.env.local')


describe('CombinedSearchWithFieldEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when POETRYDB_TEST_LIVE=TRUE.
  afterEach(liveDelay('POETRYDB_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = PoetrydbSDK.test()
    const ent = testsdk.CombinedSearchWithField()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.POETRYDB_TEST_LIVE
    for (const op of ['list']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'combined_search_with_field.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":[],"name":"combined_search_with_field","op":{"list":{"input":"data","name":"list","points":[{"active":true,"args":{"params":[{"active":true,"example":"author","kind":"param","name":"input_field1","orig":"input_field1","reqd":true,"type":"`$STRING`","index$":0},{"active":true,"example":"linecount","kind":"param","name":"input_field2","orig":"input_field2","reqd":true,"type":"`$STRING`","index$":1},{"active":true,"example":"lines","kind":"param","name":"output_field","orig":"output_field","reqd":true,"type":"`$STRING`","index$":2},{"active":true,"example":"Shakespeare","kind":"param","name":"search_term1","orig":"search_term1","reqd":true,"type":"`$STRING`","index$":3},{"active":true,"example":"14","kind":"param","name":"search_term2","orig":"search_term2","reqd":true,"type":"`$STRING`","index$":4}]},"contract":{"id":"GET /{inputField1},{inputField2}/{searchTerm1};{searchTerm2}/{outputFields}","json":"{\"operationId\":\"combinedSearchWithFields\",\"parameters\":[{\"description\":\"First input field (author, title, lines, linecount)\",\"example\":\"author\",\"in\":\"path\",\"name\":\"inputField1\",\"required\":true,\"schema\":{\"enum\":[\"author\",\"title\",\"lines\",\"linecount\"],\"type\":\"string\"}},{\"description\":\"Second input field (author, title, lines, linecount)\",\"example\":\"linecount\",\"in\":\"path\",\"name\":\"inputField2\",\"required\":true,\"schema\":{\"enum\":[\"author\",\"title\",\"lines\",\"linecount\"],\"type\":\"string\"}},{\"description\":\"Search term for first input field\",\"example\":\"Shakespeare\",\"in\":\"path\",\"name\":\"searchTerm1\",\"required\":true,\"schema\":{\"type\":\"string\"}},{\"description\":\"Search term for second input field\",\"example\":\"14\",\"in\":\"path\",\"name\":\"searchTerm2\",\"required\":true,\"schema\":{\"type\":\"string\"}},{\"description\":\"Comma-separated list of output fields (author, title, lines, linecount, all)\",\"example\":\"lines\",\"in\":\"path\",\"name\":\"outputFields\",\"required\":true,\"schema\":{\"type\":\"string\"}}],\"protocol\":\"http\",\"responses\":{\"200\":{\"content\":{\"application/json\":{\"schema\":{\"items\":{\"type\":\"object\"},\"type\":\"array\"}}},\"description\":\"Successful response\"},\"404\":{\"content\":{\"application/json\":{\"schema\":{\"properties\":{\"reason\":{\"example\":\"Not found\",\"type\":\"string\"},\"status\":{\"example\":404,\"type\":\"integer\"}},\"type\":\"object\"}}},\"description\":\"No poems found\"}},\"securitySource\":\"unspecified\"}","source":"openapi3","version":1},"kind":"http","method":"GET","orig":"/{inputField1},{inputField2}/{searchTerm1};{searchTerm2}/{outputFields}","rename":{"param":{"inputField1},{inputField2":"input_field1},{input_field2","outputFields":"output_field","searchTerm1};{searchTerm2":"search_term1};{search_term2"}},"segments":[{"lit":"{inputField1},{inputField2}"},{"lit":"{searchTerm1};{searchTerm2}"},{"var":"output_field"}],"select":{"exist":["input_field1","input_field2","output_field","search_term1","search_term2"]},"transform":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"list"}},"relations":{"ancestors":[["{search_term1};{search_term2}"]]},"key$":"combined_search_with_field","name__orig":"combined_search_with_field","Name":"CombinedSearchWithField","name_":"combined_search_with_field","name-":"combined-search-with-field","NAME":"COMBINED_SEARCH_WITH_FIELD","index$":3}, {"active":true,"entity":"combined_search_with_field","key$":"BasicCombinedSearchWithFieldFlow","kind":"basic","name":"BasicCombinedSearchWithFieldFlow","param":{},"step":[{"active":true,"data":{},"input":{},"match":{"input_field1":"input_field101","input_field2":"input_field201","output_field":"output_field01","search_term1":"search_term101","search_term2":"search_term201"},"op":"list","spec":[],"valid":[{"apply":"ItemExists","def":{"ref":"combined_search_with_field_ref01"}}],"index$":0}]}, 'CombinedSearchWithField')
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select

    let combined_search_with_field_ref01_data = Object.values(setup.data.existing.combined_search_with_field)[0] as any

    // LIST
    const combined_search_with_field_ref01_ent = client.CombinedSearchWithField()
    const combined_search_with_field_ref01_match: any = {}
    combined_search_with_field_ref01_match['input_field1'] = setup.idmap['input_field101']
    combined_search_with_field_ref01_match['input_field2'] = setup.idmap['input_field201']
    combined_search_with_field_ref01_match['output_field'] = setup.idmap['output_field01']
    combined_search_with_field_ref01_match['search_term1'] = setup.idmap['search_term101']
    combined_search_with_field_ref01_match['search_term2'] = setup.idmap['search_term201']

    const combined_search_with_field_ref01_list = (await combined_search_with_field_ref01_ent.list(combined_search_with_field_ref01_match)).map((e: any) => e.data())


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/combined_search_with_field/CombinedSearchWithFieldTestData.json')

  // TODO: file ready util needed?
  const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8')

  // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
  const entityData = JSON.parse(entityDataSource)

  options.entity = entityData.existing

  let client = PoetrydbSDK.test(options, extra)
  const struct = client.utility().struct
  const merge = struct.merge
  const transform = struct.transform

  let idmap = transform(
    ['combined_search_with_field01','combined_search_with_field02','combined_search_with_field03','{search_term1};{search_term2}01','{search_term1};{search_term2}02','{search_term1};{search_term2}03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'POETRYDB_TEST_COMBINED_SEARCH_WITH_FIELD_ENTID': idmap,
    'POETRYDB_TEST_LIVE': 'FALSE',
    'POETRYDB_TEST_EXPLAIN': 'FALSE',
  })

  idmap = env['POETRYDB_TEST_COMBINED_SEARCH_WITH_FIELD_ENTID']

  const live = 'TRUE' === env.POETRYDB_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['POETRYDB_TEST_COMBINED_SEARCH_WITH_FIELD_ENTID']
    idmap = rawIds && rawIds.trim() ? JSON.parse(rawIds) : {}
    if (!idmap || Array.isArray(idmap) || typeof idmap !== 'object') {
      throw new Error('Live ENTID must be a JSON object')
    }
    client = new PoetrydbSDK(merge([
      // FIRST, so the generated fields below win: sdk-test-control.json's
      // test.client.options adds to the live client, it does not redirect it.
      liveClientOptions(),
      {
      },
      // 'extra || {}', not a bare 'extra': struct.merge returns UNDEFINED when the
      // last entry is undefined, and basicSetup is normally called with no
      // argument at all - so a bare 'extra' silently discarded the apikey
      // and server values above and handed the SDK undefined. Harmless
      // while there was nothing in that object; not harmless now.
      extra || {},
      { system: { fetch: transport.fetch } }
    ]))
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
  }

  return setup
}
  
