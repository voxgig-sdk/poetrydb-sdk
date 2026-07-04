// Poetrydb Ts SDK

import { AuthorEntity } from './entity/AuthorEntity'
import { AuthorabEntity } from './entity/AuthorabEntity'
import { CombinedSearchEntity } from './entity/CombinedSearchEntity'
import { CombinedSearchWithFieldEntity } from './entity/CombinedSearchWithFieldEntity'
import { LineEntity } from './entity/LineEntity'
import { LinecountEntity } from './entity/LinecountEntity'
import { PoemcountEntity } from './entity/PoemcountEntity'
import { RandomEntity } from './entity/RandomEntity'
import { TitleEntity } from './entity/TitleEntity'
import { TitleabEntity } from './entity/TitleabEntity'

export type * from './PoetrydbTypes'


import { inspect } from 'node:util'

import type { Context, Feature } from './types'

import { config } from './Config'
import { PoetrydbEntityBase } from './PoetrydbEntityBase'
import { Utility } from './utility/Utility'


import { BaseFeature } from './feature/base/BaseFeature'


const stdutil = new Utility()


class PoetrydbSDK {
  _mode: string = 'live'
  _options: any
  _utility = new Utility()
  _features: Feature[]
  _rootctx: Context

  constructor(options?: any) {

    this._rootctx = this._utility.makeContext({
      client: this,
      utility: this._utility,
      config,
      options,
      shared: new WeakMap()
    })

    this._options = this._utility.makeOptions(this._rootctx)

    const struct = this._utility.struct
    const getpath = struct.getpath
    const items = struct.items

    if (true === getpath(this._options.feature, 'test.active')) {
      this._mode = 'test'
    }

    this._rootctx.options = this._options

    this._features = []

    const featureAdd = this._utility.featureAdd
    const featureInit = this._utility.featureInit

    items(this._options.feature, (fitem: [string, any]) => {
      const fname = fitem[0]
      const fopts = fitem[1]
      if (fopts.active) {
        featureAdd(this._rootctx, this._rootctx.config.makeFeature(fname))
      }
    })

    if (null != this._options.extend) {
      for (let f of this._options.extend) {
        featureAdd(this._rootctx, f)
      }
    }

    for (let f of this._features) {
      featureInit(this._rootctx, f)
    }

    const featureHook = this._utility.featureHook
    featureHook(this._rootctx, 'PostConstruct')
  }


  options() {
    return this._utility.struct.clone(this._options)
  }


  utility() {
    return this._utility.struct.clone(this._utility)
  }


  async prepare(fetchargs?: any) {
    const utility = this._utility
    const struct = utility.struct
    const clone = struct.clone

    const {
      makeContext,
      makeFetchDef,
      prepareHeaders,
      prepareAuth,
    } = utility

    fetchargs = fetchargs || {}

    let ctx: Context = makeContext({
      opname: 'prepare',
      ctrl: fetchargs.ctrl || {},
    }, this._rootctx)

    const options = this._options

    // Build spec directly from SDK options + user-provided fetch args.
    const spec: any = {
      base: options.base,
      prefix: options.prefix,
      suffix: options.suffix,
      path: fetchargs.path || '',
      method: fetchargs.method || 'GET',
      params: fetchargs.params || {},
      query: fetchargs.query || {},
      headers: prepareHeaders(ctx),
      body: fetchargs.body,
      step: 'start',
    }

    ctx.spec = spec

    // Merge user-provided headers over SDK defaults.
    if (fetchargs.headers) {
      const uheaders = fetchargs.headers
      for (let key in uheaders) {
        spec.headers[key] = uheaders[key]
      }
    }

    // Apply SDK auth (apikey, auth prefix, etc.)
    const authResult = prepareAuth(ctx)
    if (authResult instanceof Error) {
      return authResult
    }

    return makeFetchDef(ctx)
  }


  async direct(fetchargs?: any) {
    const utility = this._utility
    const fetcher = utility.fetcher
    const makeContext = utility.makeContext

    const fetchdef = await this.prepare(fetchargs)
    if (fetchdef instanceof Error) {
      return fetchdef
    }

    let ctx: Context = makeContext({
      opname: 'direct',
      ctrl: (fetchargs || {}).ctrl || {},
    }, this._rootctx)

    try {
      const fetched = await fetcher(ctx, fetchdef.url, fetchdef)

      if (null == fetched) {
        return { ok: false, err: ctx.error('direct_no_response', 'response: undefined') }
      }
      else if (fetched instanceof Error) {
        return { ok: false, err: fetched }
      }

      const status = fetched.status

      // No body responses (204 No Content, 304 Not Modified) and explicit
      // zero content-length must skip JSON parsing — fetched.json() would
      // throw `Unexpected end of JSON input` on an empty body.
      const headers = fetched.headers
      const contentLength = headers && 'function' === typeof headers.get
        ? headers.get('content-length')
        : (headers || {})['content-length']
      const noBody = 204 === status || 304 === status || '0' === String(contentLength)

      let json: any = undefined
      if (!noBody) {
        try {
          json = 'function' === typeof fetched.json ? await fetched.json() : fetched.json
        }
        catch (parseErr) {
          // Body wasn't valid JSON — surface the raw response rather than
          // throwing. data stays undefined; callers can inspect status/headers.
          json = undefined
        }
      }

      return {
        ok: status >= 200 && status < 300,
        status,
        headers: fetched.headers,
        data: json,
      }
    }
    catch (err: any) {
      return { ok: false, err }
    }
  }



  _author?: AuthorEntity

  // Idiomatic facade: `client.author.list()` / `client.author.load({ id })`.
  get author(): AuthorEntity {
    return (this._author ??= new AuthorEntity(this, undefined))
  }

  /** @deprecated Use `client.author` instead. */
  Author(data?: any) {
    const self = this
    return new AuthorEntity(self,data)
  }


  _authorab?: AuthorabEntity

  // Idiomatic facade: `client.authorab.list()` / `client.authorab.load({ id })`.
  get authorab(): AuthorabEntity {
    return (this._authorab ??= new AuthorabEntity(this, undefined))
  }

  /** @deprecated Use `client.authorab` instead. */
  Authorab(data?: any) {
    const self = this
    return new AuthorabEntity(self,data)
  }


  _combined_search?: CombinedSearchEntity

  // Idiomatic facade: `client.combined_search.list()` / `client.combined_search.load({ id })`.
  get combined_search(): CombinedSearchEntity {
    return (this._combined_search ??= new CombinedSearchEntity(this, undefined))
  }

  /** @deprecated Use `client.combined_search` instead. */
  CombinedSearch(data?: any) {
    const self = this
    return new CombinedSearchEntity(self,data)
  }


  _combined_search_with_field?: CombinedSearchWithFieldEntity

  // Idiomatic facade: `client.combined_search_with_field.list()` / `client.combined_search_with_field.load({ id })`.
  get combined_search_with_field(): CombinedSearchWithFieldEntity {
    return (this._combined_search_with_field ??= new CombinedSearchWithFieldEntity(this, undefined))
  }

  /** @deprecated Use `client.combined_search_with_field` instead. */
  CombinedSearchWithField(data?: any) {
    const self = this
    return new CombinedSearchWithFieldEntity(self,data)
  }


  _line?: LineEntity

  // Idiomatic facade: `client.line.list()` / `client.line.load({ id })`.
  get line(): LineEntity {
    return (this._line ??= new LineEntity(this, undefined))
  }

  /** @deprecated Use `client.line` instead. */
  Line(data?: any) {
    const self = this
    return new LineEntity(self,data)
  }


  _linecount?: LinecountEntity

  // Idiomatic facade: `client.linecount.list()` / `client.linecount.load({ id })`.
  get linecount(): LinecountEntity {
    return (this._linecount ??= new LinecountEntity(this, undefined))
  }

  /** @deprecated Use `client.linecount` instead. */
  Linecount(data?: any) {
    const self = this
    return new LinecountEntity(self,data)
  }


  _poemcount?: PoemcountEntity

  // Idiomatic facade: `client.poemcount.list()` / `client.poemcount.load({ id })`.
  get poemcount(): PoemcountEntity {
    return (this._poemcount ??= new PoemcountEntity(this, undefined))
  }

  /** @deprecated Use `client.poemcount` instead. */
  Poemcount(data?: any) {
    const self = this
    return new PoemcountEntity(self,data)
  }


  _random?: RandomEntity

  // Idiomatic facade: `client.random.list()` / `client.random.load({ id })`.
  get random(): RandomEntity {
    return (this._random ??= new RandomEntity(this, undefined))
  }

  /** @deprecated Use `client.random` instead. */
  Random(data?: any) {
    const self = this
    return new RandomEntity(self,data)
  }


  _title?: TitleEntity

  // Idiomatic facade: `client.title.list()` / `client.title.load({ id })`.
  get title(): TitleEntity {
    return (this._title ??= new TitleEntity(this, undefined))
  }

  /** @deprecated Use `client.title` instead. */
  Title(data?: any) {
    const self = this
    return new TitleEntity(self,data)
  }


  _titleab?: TitleabEntity

  // Idiomatic facade: `client.titleab.list()` / `client.titleab.load({ id })`.
  get titleab(): TitleabEntity {
    return (this._titleab ??= new TitleabEntity(this, undefined))
  }

  /** @deprecated Use `client.titleab` instead. */
  Titleab(data?: any) {
    const self = this
    return new TitleabEntity(self,data)
  }




  static test(testoptsarg?: any, sdkoptsarg?: any) {
    const struct = stdutil.struct
    const setpath = struct.setpath
    const getdef = struct.getdef
    const clone = struct.clone
    const setprop = struct.setprop

    const sdkopts = getdef(clone(sdkoptsarg), {})
    const testopts = getdef(clone(testoptsarg), {})
    setprop(testopts, 'active', true)
    setpath(sdkopts, 'feature.test', testopts)

    const testsdk = new PoetrydbSDK(sdkopts)
    testsdk._mode = 'test'

    return testsdk
  }


  tester(testopts?: any, sdkopts?: any) {
    return PoetrydbSDK.test(testopts, sdkopts)
  }


  toJSON() {
    return { name: 'Poetrydb' }
  }

  toString() {
    return 'Poetrydb ' + this._utility.struct.jsonify(this.toJSON())
  }

  [inspect.custom]() {
    return this.toString()
  }

}




const SDK = PoetrydbSDK


export {
  stdutil,

  BaseFeature,
  PoetrydbEntityBase,

  PoetrydbSDK,
  SDK,
}


