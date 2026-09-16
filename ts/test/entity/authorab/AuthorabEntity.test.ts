

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


describe('AuthorabEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when POETRYDB_TEST_LIVE=TRUE.
  afterEach(liveDelay('POETRYDB_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = PoetrydbSDK.test()
    const ent = testsdk.Authorab()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.POETRYDB_TEST_LIVE
    for (const op of ['list']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'authorab.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":[{"active":true,"name":"author","req":false,"short":"The author of the poem","type":"`$STRING`","index$":0},{"active":true,"name":"linecount","req":false,"short":"The number of lines in the poem (including section headings, excluding empty lines)","type":"`$INTEGER`","index$":1},{"active":true,"name":"lines","req":false,"short":"The lines of the poem","type":"`$ARRAY`","index$":2},{"active":true,"name":"title","req":false,"short":"The title of the poem","type":"`$STRING`","index$":3}],"name":"authorab","op":{"list":{"input":"data","name":"list","points":[{"active":true,"args":{"params":[{"active":true,"example":"Ernest Dowson","kind":"param","name":"author","orig":"author","reqd":true,"type":"`$STRING`","index$":0}]},"contract":{"id":"GET /author/{author}:abs","json":"{\"operationId\":\"searchByAuthorExact\",\"parameters\":[{\"description\":\"The exact name of the author\",\"example\":\"Ernest Dowson\",\"in\":\"path\",\"name\":\"author\",\"required\":true,\"schema\":{\"type\":\"string\"}}],\"protocol\":\"http\",\"responses\":{\"200\":{\"content\":{\"application/json\":{\"schema\":{\"items\":{\"properties\":{\"author\":{\"description\":\"The author of the poem\",\"example\":\"Percy Bysshe Shelley\",\"type\":\"string\"},\"linecount\":{\"description\":\"The number of lines in the poem (including section headings, excluding empty lines)\",\"example\":14,\"type\":\"integer\"},\"lines\":{\"description\":\"The lines of the poem\",\"example\":[\"I met a traveller from an antique land\",\"Who said: \\\"Two vast and trunkless legs of stone\",\"Stand in the desert. Near them on the sand,\"],\"items\":{\"type\":\"string\"},\"type\":\"array\"},\"title\":{\"description\":\"The title of the poem\",\"example\":\"Ozymandias\",\"type\":\"string\"}},\"type\":\"object\"},\"type\":\"array\"}}},\"description\":\"Successful response\"},\"404\":{\"content\":{\"application/json\":{\"schema\":{\"properties\":{\"reason\":{\"example\":\"Not found\",\"type\":\"string\"},\"status\":{\"example\":404,\"type\":\"integer\"}},\"type\":\"object\"}}},\"description\":\"No poems found\"}},\"securitySource\":\"unspecified\"}","source":"openapi3","version":1},"kind":"http","method":"GET","orig":"/author/{author}:abs","segments":[{"lit":"author"},{"lit":"{author}:abs"}],"select":{"exist":["author"]},"transform":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"list"}},"relations":{"ancestors":[]},"key$":"authorab","name__orig":"authorab","Name":"Authorab","name_":"authorab","name-":"authorab","NAME":"AUTHORAB","index$":1}, {"active":true,"entity":"authorab","key$":"BasicAuthorabFlow","kind":"basic","name":"BasicAuthorabFlow","param":{},"step":[{"active":true,"data":{},"input":{},"match":{"author":"author01"},"op":"list","spec":[],"valid":[{"apply":"ItemExists","def":{"ref":"authorab_ref01"}}],"index$":0}]}, 'Authorab')
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select

    let authorab_ref01_data = Object.values(setup.data.existing.authorab)[0] as any

    // LIST
    const authorab_ref01_ent = client.Authorab()
    const authorab_ref01_match: any = {}
    authorab_ref01_match['author'] = setup.idmap['author01']

    const authorab_ref01_list = (await authorab_ref01_ent.list(authorab_ref01_match)).map((e: any) => e.data())


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/authorab/AuthorabTestData.json')

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
    ['authorab01','authorab02','authorab03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'POETRYDB_TEST_AUTHORAB_ENTID': idmap,
    'POETRYDB_TEST_LIVE': 'FALSE',
    'POETRYDB_TEST_EXPLAIN': 'FALSE',
  })

  idmap = env['POETRYDB_TEST_AUTHORAB_ENTID']

  const live = 'TRUE' === env.POETRYDB_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['POETRYDB_TEST_AUTHORAB_ENTID']
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
  
