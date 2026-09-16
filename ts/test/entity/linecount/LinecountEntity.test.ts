

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


describe('LinecountEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when POETRYDB_TEST_LIVE=TRUE.
  afterEach(liveDelay('POETRYDB_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = PoetrydbSDK.test()
    const ent = testsdk.Linecount()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.POETRYDB_TEST_LIVE
    for (const op of ['list', 'load']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'linecount.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":[{"active":true,"name":"author","req":false,"short":"The author of the poem","type":"`$STRING`","index$":0},{"active":true,"name":"id","req":false,"type":"`$STRING`","index$":1},{"active":true,"name":"linecount","req":false,"short":"The number of lines in the poem (including section headings, excluding empty lines)","type":"`$INTEGER`","index$":2},{"active":true,"name":"lines","req":false,"short":"The lines of the poem","type":"`$ARRAY`","index$":3},{"active":true,"name":"title","req":false,"short":"The title of the poem","type":"`$STRING`","index$":4}],"id":{"field":"id","name":"id"},"name":"linecount","op":{"list":{"input":"data","name":"list","points":[{"active":true,"args":{"params":[{"active":true,"example":"text","kind":"param","name":"format","orig":"format","reqd":true,"type":"`$STRING`","index$":0},{"active":true,"example":14,"kind":"param","name":"linecount","orig":"linecount","reqd":true,"type":"`$INTEGER`","index$":1},{"active":true,"example":"author,title","kind":"param","name":"output_field","orig":"output_field","reqd":true,"type":"`$STRING`","index$":2}]},"contract":{"id":"GET /linecount/{linecount}/{outputFields}.{format}","json":"{\"operationId\":\"searchByLinecountWithFieldsAndFormat\",\"parameters\":[{\"description\":\"The exact number of lines in the poem\",\"example\":14,\"in\":\"path\",\"name\":\"linecount\",\"required\":true,\"schema\":{\"type\":\"integer\"}},{\"description\":\"Comma-separated list of output fields (author, title, lines, linecount, all)\",\"example\":\"author,title\",\"in\":\"path\",\"name\":\"outputFields\",\"required\":true,\"schema\":{\"type\":\"string\"}},{\"description\":\"Output format (json or text)\",\"example\":\"text\",\"in\":\"path\",\"name\":\"format\",\"required\":true,\"schema\":{\"enum\":[\"json\",\"text\"],\"type\":\"string\"}}],\"protocol\":\"http\",\"responses\":{\"200\":{\"content\":{\"application/json\":{\"schema\":{\"items\":{\"type\":\"object\"},\"type\":\"array\"}},\"text/plain\":{\"schema\":{\"type\":\"string\"}}},\"description\":\"Successful response\"},\"404\":{\"content\":{\"application/json\":{\"schema\":{\"properties\":{\"reason\":{\"example\":\"Not found\",\"type\":\"string\"},\"status\":{\"example\":404,\"type\":\"integer\"}},\"type\":\"object\"}}},\"description\":\"No poems found\"}},\"securitySource\":\"unspecified\"}","source":"openapi3","version":1},"kind":"http","method":"GET","orig":"/linecount/{linecount}/{outputFields}.{format}","rename":{"param":{"outputFields}.{format":"output_fields}_{format"}},"segments":[{"lit":"linecount"},{"var":"linecount"},{"lit":"{outputFields}.{format}"}],"select":{"exist":["format","linecount","output_field"]},"transform":{"req":"`reqdata`","res":"`body`"},"index$":0},{"active":true,"args":{"params":[{"active":true,"example":14,"kind":"param","name":"linecount","orig":"linecount","reqd":true,"type":"`$INTEGER`","index$":0},{"active":true,"example":"author,title","kind":"param","name":"output_field","orig":"output_field","reqd":true,"type":"`$STRING`","index$":1}]},"contract":{"id":"GET /linecount/{linecount}/{outputFields}","json":"{\"operationId\":\"searchByLinecountWithFields\",\"parameters\":[{\"description\":\"The exact number of lines in the poem\",\"example\":14,\"in\":\"path\",\"name\":\"linecount\",\"required\":true,\"schema\":{\"type\":\"integer\"}},{\"description\":\"Comma-separated list of output fields (author, title, lines, linecount, all)\",\"example\":\"author,title\",\"in\":\"path\",\"name\":\"outputFields\",\"required\":true,\"schema\":{\"type\":\"string\"}}],\"protocol\":\"http\",\"responses\":{\"200\":{\"content\":{\"application/json\":{\"schema\":{\"items\":{\"type\":\"object\"},\"type\":\"array\"}}},\"description\":\"Successful response\"},\"404\":{\"content\":{\"application/json\":{\"schema\":{\"properties\":{\"reason\":{\"example\":\"Not found\",\"type\":\"string\"},\"status\":{\"example\":404,\"type\":\"integer\"}},\"type\":\"object\"}}},\"description\":\"No poems found\"}},\"securitySource\":\"unspecified\"}","source":"openapi3","version":1},"kind":"http","method":"GET","orig":"/linecount/{linecount}/{outputFields}","rename":{"param":{"outputFields":"output_field"}},"segments":[{"lit":"linecount"},{"var":"linecount"},{"var":"output_field"}],"select":{"exist":["linecount","output_field"]},"transform":{"req":"`reqdata`","res":"`body`"},"index$":1}],"key$":"list"},"load":{"input":"data","name":"load","points":[{"active":true,"args":{"params":[{"active":true,"example":14,"kind":"param","name":"id","orig":"linecount","reqd":true,"type":"`$INTEGER`","index$":0}]},"contract":{"id":"GET /linecount/{linecount}","json":"{\"operationId\":\"searchByLinecount\",\"parameters\":[{\"description\":\"The exact number of lines in the poem\",\"example\":14,\"in\":\"path\",\"name\":\"linecount\",\"required\":true,\"schema\":{\"type\":\"integer\"}}],\"protocol\":\"http\",\"responses\":{\"200\":{\"content\":{\"application/json\":{\"schema\":{\"items\":{\"properties\":{\"author\":{\"description\":\"The author of the poem\",\"example\":\"Percy Bysshe Shelley\",\"type\":\"string\"},\"linecount\":{\"description\":\"The number of lines in the poem (including section headings, excluding empty lines)\",\"example\":14,\"type\":\"integer\"},\"lines\":{\"description\":\"The lines of the poem\",\"example\":[\"I met a traveller from an antique land\",\"Who said: \\\"Two vast and trunkless legs of stone\",\"Stand in the desert. Near them on the sand,\"],\"items\":{\"type\":\"string\"},\"type\":\"array\"},\"title\":{\"description\":\"The title of the poem\",\"example\":\"Ozymandias\",\"type\":\"string\"}},\"type\":\"object\"},\"type\":\"array\"}}},\"description\":\"Successful response\"},\"404\":{\"content\":{\"application/json\":{\"schema\":{\"properties\":{\"reason\":{\"example\":\"Not found\",\"type\":\"string\"},\"status\":{\"example\":404,\"type\":\"integer\"}},\"type\":\"object\"}}},\"description\":\"No poems found\"}},\"securitySource\":\"unspecified\"}","source":"openapi3","version":1},"kind":"http","method":"GET","orig":"/linecount/{linecount}","rename":{"param":{"linecount":"id"}},"segments":[{"lit":"linecount"},{"var":"id"}],"select":{"exist":["id"]},"transform":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"load"}},"relations":{"ancestors":[["linecount"]]},"key$":"linecount","name__orig":"linecount","Name":"Linecount","name_":"linecount","name-":"linecount","NAME":"LINECOUNT","index$":5}, {"active":true,"entity":"linecount","key$":"BasicLinecountFlow","kind":"basic","name":"BasicLinecountFlow","param":{},"step":[{"active":true,"data":{},"input":{},"match":{"linecount":"linecount01","output_field":"output_field01"},"op":"list","spec":[],"valid":[{"apply":"ItemExists","def":{"ref":"linecount_ref01"}}],"index$":0},{"active":true,"data":{},"input":{"ref":"linecount_ref01","srcdatavar":"linecount_ref01_data","suffix":"_dt0"},"match":{"id":"linecount01"},"op":"load","spec":[],"valid":[{"apply":"TextFieldMark","def":{"mark":"Mark01-linecount_ref01"}}],"index$":1}]}, 'Linecount')
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select

    let linecount_ref01_data = Object.values(setup.data.existing.linecount)[0] as any

    // LIST
    const linecount_ref01_ent = client.Linecount()
    const linecount_ref01_match: any = {}
    linecount_ref01_match['linecount'] = setup.idmap['linecount01']
    linecount_ref01_match['output_field'] = setup.idmap['output_field01']

    const linecount_ref01_list = (await linecount_ref01_ent.list(linecount_ref01_match)).map((e: any) => e.data())


    // LOAD
    const linecount_ref01_match_dt0: any = {}
    linecount_ref01_match_dt0.id = linecount_ref01_data.id
    const linecount_ref01_data_dt0 = (await linecount_ref01_ent.load(linecount_ref01_match_dt0)).data()
    assert(linecount_ref01_data_dt0.id === linecount_ref01_data.id)


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/linecount/LinecountTestData.json')

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
    ['linecount01','linecount02','linecount03','linecount01','linecount02','linecount03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'POETRYDB_TEST_LINECOUNT_ENTID': idmap,
    'POETRYDB_TEST_LIVE': 'FALSE',
    'POETRYDB_TEST_EXPLAIN': 'FALSE',
  })

  idmap = env['POETRYDB_TEST_LINECOUNT_ENTID']

  const live = 'TRUE' === env.POETRYDB_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['POETRYDB_TEST_LINECOUNT_ENTID']
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
  
