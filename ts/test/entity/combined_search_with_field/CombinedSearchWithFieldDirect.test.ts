
const envlocal = __dirname + '/../../../.env.local'
require('dotenv').config({ quiet: true, path: [envlocal] })

import { test, describe, afterEach } from 'node:test'
import assert from 'node:assert'


import { PoetrydbSDK } from '../../..'

import {
  envOverride,
  liveDelay,
  maybeSkipControl,
  skipIfMissingIds,
} from '../../utility'


describe('CombinedSearchWithFieldDirect', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when POETRYDB_TEST_LIVE=TRUE.
  afterEach(liveDelay('POETRYDB_TEST_LIVE'))

  test('direct-exists', async () => {
    const sdk = new PoetrydbSDK({
      system: { fetch: async () => ({}) }
    })
    assert('function' === typeof sdk.direct)
    assert('function' === typeof sdk.prepare)
  })


  test('direct-list-combined_search_with_field', async (t: any) => {
    const setup = directSetup([{ id: 'direct01' }, { id: 'direct02' }])
    if (maybeSkipControl(t, 'direct', 'direct-list-combined_search_with_field', setup.live)) return
    if (skipIfMissingIds(t, setup, ["input_field101","input_field201","output_field01","search_term101","search_term201"])) return
    const { client, calls } = setup

    const params: any = {}
    const query: any = {}
    if (setup.live) {
      params.input_field1 = setup.idmap['input_field101']
      params.input_field2 = setup.idmap['input_field201']
      params.output_field = setup.idmap['output_field01']
      params.search_term1 = setup.idmap['search_term101']
      params.search_term2 = setup.idmap['search_term201']
    } else {
      params.input_field1 = 'direct01'
      params.input_field2 = 'direct02'
      params.output_field = 'direct03'
      params.search_term1 = 'direct04'
      params.search_term2 = 'direct05'
    }

    const result: any = await client.direct({
      path: '{input_field1},{input_field2}/{search_term1};{search_term2}/{output_field}',
      method: 'GET',
      params,
      query,
    })

    if (setup.live) {
      // Live mode is lenient: synthetic IDs frequently 4xx and the list-
      // response shape varies wildly across public APIs. Skip rather than
      // fail when the call doesn't return a usable list.
      if (!result.ok || result.status < 200 || result.status >= 300) {
        return
      }
      const listArr = unwrapListData(result.data)
      if (!Array.isArray(listArr)) {
        return
      }
    } else {
      assert(result.ok === true)
      assert(result.status === 200)
      assert(null != result.data)
      const listArr = unwrapListData(result.data)
      assert(Array.isArray(listArr))
      assert(listArr!.length === 2)
      assert(calls.length === 1)
      assert(calls[0].init.method === 'GET')
      assert(calls[0].url.includes('direct01'))
      assert(calls[0].url.includes('direct02'))
      assert(calls[0].url.includes('direct03'))
      assert(calls[0].url.includes('direct04'))
      assert(calls[0].url.includes('direct05'))
    }
  })

})



function directSetup(mockres?: any) {
  const calls: any[] = []

  const env = envOverride({
    'POETRYDB_TEST_COMBINED_SEARCH_WITH_FIELD_ENTID': {},
    'POETRYDB_TEST_LIVE': 'FALSE',
    'POETRYDB_APIKEY': 'NONE',
  })

  const live = 'TRUE' === env.POETRYDB_TEST_LIVE

  if (live) {
    const client = new PoetrydbSDK({
      apikey: env.POETRYDB_APIKEY,
    })

    let idmap: any = env['POETRYDB_TEST_COMBINED_SEARCH_WITH_FIELD_ENTID']
    if ('string' === typeof idmap && idmap.startsWith('{')) {
      idmap = JSON.parse(idmap)
    }

    return { client, calls, live, idmap }
  }

  const mockFetch = async (url: string, init: any) => {
    calls.push({ url, init })
    return {
      status: 200,
      statusText: 'OK',
      headers: {},
      json: async () => (null != mockres ? mockres : { id: 'direct01' }),
    }
  }

  const client = new PoetrydbSDK({
    base: 'http://localhost:8080',
    system: { fetch: mockFetch },
  })

  return { client, calls, live, idmap: {} as any }
}

// direct() returns the raw response body. List endpoints often wrap the
// array in an envelope (e.g. { data: [...] }, { entities: [...] },
// { pagination, data: [...] }). The test transforms the raw body to
// extract the first array — either the body itself or the first array
// property of an envelope object.
function unwrapListData(data: any): any[] | null {
  if (Array.isArray(data)) return data
  if (data && 'object' === typeof data) {
    for (const v of Object.values(data)) {
      if (Array.isArray(v)) return v as any[]
    }
  }
  return null
}
  
