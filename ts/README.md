# Poetrydb TypeScript SDK

The TypeScript SDK for the Poetrydb API. Provides a type-safe, entity-oriented interface with full async/await support.


## Install
```bash
npm install poetrydb
```
## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```ts
import { PoetrydbSDK } from 'poetrydb'

const client = new PoetrydbSDK({})
```

### 2. List authors

```ts
const result = await client.Author().list()

if (result.ok) {
  for (const item of result.data) {
    console.log(item.id, item.name)
  }
}
```

### 3. Load a author

```ts
const result = await client.Author().load({ id: 'example_id' })

if (result.ok) {
  console.log(result.data)
}
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```ts
const result = await client.direct({
  path: '/api/resource/{id}',
  method: 'GET',
  params: { id: 'example' },
})

if (result.ok) {
  console.log(result.status)  // 200
  console.log(result.data)    // response body
}
```

### Prepare a request without sending it

```ts
const fetchdef = await client.prepare({
  path: '/api/resource/{id}',
  method: 'DELETE',
  params: { id: 'example' },
})

// Inspect before sending
console.log(fetchdef.url)
console.log(fetchdef.method)
console.log(fetchdef.headers)
```

### Use test mode

Create a mock client for unit testing — no server required:

```ts
const client = PoetrydbSDK.test()

const result = await client.Planet().load({ id: 'test01' })
// result.ok === true
// result.data contains mock response data
```

You can also use the instance method:

```ts
const client = new PoetrydbSDK()
const testClient = client.tester()
```

### Retain entity state across calls

Entity instances remember their last match and data:

```ts
const entity = client.Planet()

// First call sets internal match
await entity.load({ id: 'example' })

// Subsequent calls reuse the stored match
const data = entity.data()
console.log(data.id) // 'example'
```

### Add custom middleware

Pass features via the `extend` option:

```ts
const logger = {
  hooks: {
    PreRequest: (ctx: any) => {
      console.log('Requesting:', ctx.spec.method, ctx.spec.path)
    },
    PreResponse: (ctx: any) => {
      console.log('Status:', ctx.out.request?.status)
    },
  },
}

const client = new PoetrydbSDK({
  extend: [logger],
})
```

### Run live tests

Create a `.env.local` file at the project root:

```
POETRYDB_TEST_LIVE=TRUE
```

Then run:

```bash
cd ts && npm test
```


## Reference

### PoetrydbSDK

#### Constructor

```ts
new PoetrydbSDK(options?: {
  base?: string
  prefix?: string
  suffix?: string
  feature?: Record<string, { active: boolean }>
  extend?: Feature[]
})
```

| Option | Type | Description |
| --- | --- | --- |
| `base` | `string` | Base URL of the API server. |
| `prefix` | `string` | URL path prefix prepended to all requests. |
| `suffix` | `string` | URL path suffix appended to all requests. |
| `feature` | `object` | Feature activation flags (e.g. `{ test: { active: true } }`). |
| `extend` | `Feature[]` | Additional feature instances to load. |

#### Methods

| Method | Returns | Description |
| --- | --- | --- |
| `options()` | `object` | Deep copy of current SDK options. |
| `utility()` | `Utility` | Deep copy of the SDK utility object. |
| `prepare(fetchargs?)` | `Promise<FetchDef>` | Build an HTTP request definition without sending it. |
| `direct(fetchargs?)` | `Promise<DirectResult>` | Build and send an HTTP request. |
| `Author(data?)` | `AuthorEntity` | Create a Author entity instance. |
| `Authorab(data?)` | `AuthorabEntity` | Create a Authorab entity instance. |
| `CombinedSearch(data?)` | `CombinedSearchEntity` | Create a CombinedSearch entity instance. |
| `CombinedSearchWithField(data?)` | `CombinedSearchWithFieldEntity` | Create a CombinedSearchWithField entity instance. |
| `Line(data?)` | `LineEntity` | Create a Line entity instance. |
| `Linecount(data?)` | `LinecountEntity` | Create a Linecount entity instance. |
| `Poemcount(data?)` | `PoemcountEntity` | Create a Poemcount entity instance. |
| `Random(data?)` | `RandomEntity` | Create a Random entity instance. |
| `Title(data?)` | `TitleEntity` | Create a Title entity instance. |
| `Titleab(data?)` | `TitleabEntity` | Create a Titleab entity instance. |
| `tester(testopts?, sdkopts?)` | `PoetrydbSDK` | Create a test-mode client instance. |

#### Static methods

| Method | Returns | Description |
| --- | --- | --- |
| `PoetrydbSDK.test(testopts?, sdkopts?)` | `PoetrydbSDK` | Create a test-mode client. |

### Entity interface

All entities share the same interface.

#### Methods

| Method | Signature | Description |
| --- | --- | --- |
| `load` | `load(reqmatch?, ctrl?): Promise<Result>` | Load a single entity by match criteria. |
| `list` | `list(reqmatch?, ctrl?): Promise<Result>` | List entities matching the criteria. |
| `create` | `create(reqdata?, ctrl?): Promise<Result>` | Create a new entity. |
| `update` | `update(reqdata?, ctrl?): Promise<Result>` | Update an existing entity. |
| `remove` | `remove(reqmatch?, ctrl?): Promise<Result>` | Remove an entity. |
| `data` | `data(data?): any` | Get or set entity data. |
| `match` | `match(match?): any` | Get or set entity match criteria. |
| `make` | `make(): Entity` | Create a new instance with the same options. |
| `client` | `client(): PoetrydbSDK` | Return the parent SDK client. |
| `entopts` | `entopts(): object` | Return a copy of the entity options. |

#### Result shape

All entity operations return a Result object:

```ts
{
  ok: boolean      // true if the HTTP status is 2xx
  status: number   // HTTP status code
  headers: object  // response headers
  data: any        // parsed JSON response body
}
```

### DirectResult shape

The `direct()` method returns:

```ts
{
  ok: boolean
  status: number
  headers: object
  data: any
}
```

On error, `ok` is `false` and an `err` property contains the error.

### FetchDef shape

The `prepare()` method returns:

```ts
{
  url: string
  method: string
  headers: Record<string, string>
  body?: any
}
```

### Entities

#### Author

| Field | Description |
| --- | --- |
| `author` |  |
| `line` |  |
| `linecount` |  |
| `title` |  |

Operations: list, load.

API path: `/author/{author}/{outputFields}.{format}`

#### Authorab

| Field | Description |
| --- | --- |
| `author` |  |
| `line` |  |
| `linecount` |  |
| `title` |  |

Operations: list.

API path: `/author/{author}:abs`

#### CombinedSearch

| Field | Description |
| --- | --- |
| `author` |  |
| `line` |  |
| `linecount` |  |
| `title` |  |

Operations: list.

API path: `/{inputField1},{inputField2}/{searchTerm1};{searchTerm2}`

#### CombinedSearchWithField

| Field | Description |
| --- | --- |

Operations: list.

API path: `/{inputField1},{inputField2}/{searchTerm1};{searchTerm2}/{outputFields}`

#### Line

| Field | Description |
| --- | --- |
| `author` |  |
| `line` |  |
| `linecount` |  |
| `title` |  |

Operations: list, load.

API path: `/lines/{lines}/{outputFields}.{format}`

#### Linecount

| Field | Description |
| --- | --- |
| `author` |  |
| `line` |  |
| `linecount` |  |
| `title` |  |

Operations: list, load.

API path: `/linecount/{linecount}/{outputFields}.{format}`

#### Poemcount

| Field | Description |
| --- | --- |
| `author` |  |
| `line` |  |
| `linecount` |  |
| `title` |  |

Operations: load.

API path: `/poemcount/{count}`

#### Random

| Field | Description |
| --- | --- |
| `author` |  |
| `line` |  |
| `linecount` |  |
| `title` |  |

Operations: list, load.

API path: `/random/{count}/{outputFields}`

#### Title

| Field | Description |
| --- | --- |
| `author` |  |
| `line` |  |
| `linecount` |  |
| `title` |  |

Operations: list, load.

API path: `/title/{title}/{outputFields}.{format}`

#### Titleab

| Field | Description |
| --- | --- |
| `author` |  |
| `line` |  |
| `linecount` |  |
| `title` |  |

Operations: list.

API path: `/title/{title}:abs`



## Entities


### Author

Create an instance: `const author = client.Author()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `author` | ``$STRING`` |  |
| `line` | ``$ARRAY`` |  |
| `linecount` | ``$INTEGER`` |  |
| `title` | ``$STRING`` |  |

#### Example: Load

```ts
const author = await client.Author().load({ id: 'author_id' })
```

#### Example: List

```ts
const authors = await client.Author().list()
```


### Authorab

Create an instance: `const authorab = client.Authorab()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `author` | ``$STRING`` |  |
| `line` | ``$ARRAY`` |  |
| `linecount` | ``$INTEGER`` |  |
| `title` | ``$STRING`` |  |

#### Example: List

```ts
const authorabs = await client.Authorab().list()
```


### CombinedSearch

Create an instance: `const combined_search = client.CombinedSearch()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `author` | ``$STRING`` |  |
| `line` | ``$ARRAY`` |  |
| `linecount` | ``$INTEGER`` |  |
| `title` | ``$STRING`` |  |

#### Example: List

```ts
const combined_searchs = await client.CombinedSearch().list()
```


### CombinedSearchWithField

Create an instance: `const combined_search_with_field = client.CombinedSearchWithField()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Example: List

```ts
const combined_search_with_fields = await client.CombinedSearchWithField().list()
```


### Line

Create an instance: `const line = client.Line()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `author` | ``$STRING`` |  |
| `line` | ``$ARRAY`` |  |
| `linecount` | ``$INTEGER`` |  |
| `title` | ``$STRING`` |  |

#### Example: Load

```ts
const line = await client.Line().load({ id: 'line_id' })
```

#### Example: List

```ts
const lines = await client.Line().list()
```


### Linecount

Create an instance: `const linecount = client.Linecount()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `author` | ``$STRING`` |  |
| `line` | ``$ARRAY`` |  |
| `linecount` | ``$INTEGER`` |  |
| `title` | ``$STRING`` |  |

#### Example: Load

```ts
const linecount = await client.Linecount().load({ id: 'linecount_id' })
```

#### Example: List

```ts
const linecounts = await client.Linecount().list()
```


### Poemcount

Create an instance: `const poemcount = client.Poemcount()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `author` | ``$STRING`` |  |
| `line` | ``$ARRAY`` |  |
| `linecount` | ``$INTEGER`` |  |
| `title` | ``$STRING`` |  |

#### Example: Load

```ts
const poemcount = await client.Poemcount().load({ id: 'poemcount_id' })
```


### Random

Create an instance: `const random = client.Random()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `author` | ``$STRING`` |  |
| `line` | ``$ARRAY`` |  |
| `linecount` | ``$INTEGER`` |  |
| `title` | ``$STRING`` |  |

#### Example: Load

```ts
const random = await client.Random().load({ id: 'random_id' })
```

#### Example: List

```ts
const randoms = await client.Random().list()
```


### Title

Create an instance: `const title = client.Title()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `author` | ``$STRING`` |  |
| `line` | ``$ARRAY`` |  |
| `linecount` | ``$INTEGER`` |  |
| `title` | ``$STRING`` |  |

#### Example: Load

```ts
const title = await client.Title().load({ id: 'title_id' })
```

#### Example: List

```ts
const titles = await client.Title().list()
```


### Titleab

Create an instance: `const titleab = client.Titleab()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `author` | ``$STRING`` |  |
| `line` | ``$ARRAY`` |  |
| `linecount` | ``$INTEGER`` |  |
| `title` | ``$STRING`` |  |

#### Example: List

```ts
const titleabs = await client.Titleab().list()
```


## Explanation

### The operation pipeline

Every entity operation (load, list, create, update, remove) follows a
six-stage pipeline. Each stage fires a feature hook before executing:

```
PrePoint → PreSpec → PreRequest → PreResponse → PreResult → PreDone
```

- **PrePoint**: Resolves which API endpoint to call based on the
  operation name and entity configuration.
- **PreSpec**: Builds the HTTP spec — URL, method, headers, body —
  from the resolved point and the caller's parameters.
- **PreRequest**: Sends the HTTP request. Features can intercept here
  to replace the transport (as TestFeature does with mocks).
- **PreResponse**: Parses the raw HTTP response.
- **PreResult**: Extracts the business data from the parsed response.
- **PreDone**: Final stage before returning to the caller. Entity
  state (match, data) is updated here.

If any stage returns an error, the pipeline short-circuits and the
error is returned to the caller.

An unexpected exception triggers the `PreUnexpected` hook before
propagating.

### Features and hooks

Features are the extension mechanism. A feature is an object with a
`hooks` map. Each hook key is a pipeline stage name, and the value is
a function that receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Module structure

```
poetrydb/
├── src/
│   ├── PoetrydbSDK.ts        # Main SDK class
│   ├── entity/             # Entity implementations
│   ├── feature/            # Built-in features (Base, Test, Log)
│   └── utility/            # Utility functions
├── test/                   # Test suites
└── dist/                   # Compiled output
```

Import the SDK from the package root:

```ts
import { PoetrydbSDK } from 'poetrydb'
```

### Entity state

Entity instances are stateful. After a successful `load`, the entity
stores the returned data and match criteria internally. Subsequent
calls on the same instance can rely on this state.

```ts
const moon = client.Moon()
await moon.load({ planet_id: 'earth', id: 'luna' })

// moon.data() now returns the loaded moon data
// moon.match() returns { planet_id: 'earth', id: 'luna' }
```

Call `make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

The `direct` method gives full control over the HTTP request. Use it
for non-standard endpoints, bulk operations, or any path not modelled
as an entity. The `prepare` method is useful for debugging — it
shows exactly what `direct` would send.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
