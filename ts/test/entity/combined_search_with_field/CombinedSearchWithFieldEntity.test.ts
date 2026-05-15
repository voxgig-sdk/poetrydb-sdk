
const envlocal = __dirname + '/../../../.env.local'
require('dotenv').config({ quiet: true, path: [envlocal] })

import Path from 'node:path'
import * as Fs from 'node:fs'

import { test, describe, afterEach } from 'node:test'
import assert from 'node:assert'


import { PoetrydbSDK, BaseFeature, stdutil } from '../../..'

import {
  envOverride,
  liveDelay,
  makeCtrl,
  makeMatch,
  makeReqdata,
  makeStepData,
  makeValid,
  maybeSkipControl,
} from '../../utility'


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
      if (maybeSkipControl(t, 'entityOp', 'combined_search_with_field.' + op, live)) return
    }

    const setup = basicSetup()
    // The basic flow consumes synthetic IDs and field values from the
    // fixture (entity TestData.json). Those don't exist on the live API.
    // Skip live runs unless the user provided a real ENTID env override.
    if (setup.syntheticOnly) {
      t.skip('live entity test uses synthetic IDs from fixture — set POETRYDB_TEST_COMBINED_SEARCH_WITH_FIELD_ENTID JSON to run live')
      return
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

    const combined_search_with_field_ref01_list = await combined_search_with_field_ref01_ent.list(combined_search_with_field_ref01_match)


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
    ['combined_search_with_field01','combined_search_with_field02','combined_search_with_field03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  // Detect whether the user provided a real ENTID JSON via env var. The
  // basic flow consumes synthetic IDs from the fixture file; without an
  // override those synthetic IDs reach the live API and 4xx. Surface this
  // to the test so it can skip rather than fail.
  const idmapEnvVal = process.env['POETRYDB_TEST_COMBINED_SEARCH_WITH_FIELD_ENTID']
  const idmapOverridden = null != idmapEnvVal && idmapEnvVal.trim().startsWith('{')

  const env = envOverride({
    'POETRYDB_TEST_COMBINED_SEARCH_WITH_FIELD_ENTID': idmap,
    'POETRYDB_TEST_LIVE': 'FALSE',
    'POETRYDB_TEST_EXPLAIN': 'FALSE',
    'POETRYDB_APIKEY': 'NONE',
  })

  idmap = env['POETRYDB_TEST_COMBINED_SEARCH_WITH_FIELD_ENTID']

  const live = 'TRUE' === env.POETRYDB_TEST_LIVE

  if (live) {
    client = new PoetrydbSDK(merge([
      {
        apikey: env.POETRYDB_APIKEY,
      },
      extra
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
    syntheticOnly: live && !idmapOverridden,
    now: Date.now(),
  }

  return setup
}
  
