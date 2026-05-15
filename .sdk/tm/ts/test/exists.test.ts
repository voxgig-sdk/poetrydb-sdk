
import { test, describe } from 'node:test'
import { equal } from 'node:assert'


import { PoetrydbSDK } from '..'


describe('exists', async () => {

  test('test-mode', async () => {
    const testsdk = await PoetrydbSDK.test()
    equal(null !== testsdk, true)
  })

})
