
import { Context } from './Context'


class PoetrydbError extends Error {

  isPoetrydbError = true

  sdk = 'Poetrydb'

  code: string
  ctx: Context

  constructor(code: string, msg: string, ctx: Context) {
    super(msg)
    this.code = code
    this.ctx = ctx
  }

}

export {
  PoetrydbError
}

