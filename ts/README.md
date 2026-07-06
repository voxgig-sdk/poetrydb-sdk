# Poetrydb TypeScript SDK



The TypeScript SDK for the Poetrydb API — a type-safe, entity-oriented client with full async/await support.

The API is exposed as capitalised, semantic **Entities** — e.g.
`client.Author()` — each with a small set of operations (`list`, `load`)
instead of raw URL paths and query parameters. This keeps the surface
predictable and low-friction for both humans and AI agents.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to npm. Install it from the GitHub
release tag (`ts/vX.Y.Z`):

- Releases: [https://github.com/voxgig-sdk/poetrydb-sdk/releases](https://github.com/voxgig-sdk/poetrydb-sdk/releases)


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```ts
import { PoetrydbSDK } from '@voxgig-sdk/poetrydb'

const client = new PoetrydbSDK()
```

### 2. List author records

`list()` resolves to an array of Author objects — iterate it directly:

```ts
const authors = await client.Author().list()

for (const author of authors) {
  console.log(author)
}
```

### 3. Load an author

`load()` returns the entity directly and throws on failure:

```ts
try {
  const author = await client.Author().load({ id: 'example_id' })
  console.log(author)
} catch (err) {
  console.error('load failed:', err)
}
```


## Error handling

Entity operations reject on failure, so wrap them in `try` / `catch`:

```ts
try {
  const authors = await client.Author().list()
  console.log(authors)
} catch (err) {
  console.error('list failed:', err)
}
```

The low-level `direct()` method does **not** throw — it returns the
value or an `Error`, so check the result before using it:

```ts
const result = await client.direct({
  path: '/api/resource/{id}',
  method: 'GET',
  params: { id: 'example_id' },
})

if (result instanceof Error) {
  throw result
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

if (result instanceof Error) {
  throw result
}
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

const author = await client.Author().list()
// author is a bare entity populated with mock response data
console.log(author)
```

You can also use the instance method:

```ts
const client = new PoetrydbSDK()
const testClient = client.tester()
```

### Retain entity state across calls

Entity instances remember their last match and data:

```ts
const entity = client.Author()

// First call runs the operation and stores its result
await entity.list()

// Subsequent calls reuse the stored state
const data = entity.data()
console.log(data)
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
| `Author(data?)` | `AuthorEntity` | Create an Author entity instance. |
| `Authorab(data?)` | `AuthorabEntity` | Create an Authorab entity instance. |
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
| `load` | `load(reqmatch?, ctrl?): Promise<Entity>` | Load a single entity by match criteria. |
| `list` | `list(reqmatch?, ctrl?): Promise<Entity[]>` | List entities matching the criteria. |
| `data` | `data(data?: Partial<Entity>): Entity` | Get or set entity data. |
| `match` | `match(match?: Partial<Entity>): Partial<Entity>` | Get or set entity match criteria. |
| `make` | `make(): Entity` | Create a new instance with the same options. |
| `client` | `client(): PoetrydbSDK` | Return the parent SDK client. |
| `entopts` | `entopts(): object` | Return a copy of the entity options. |

#### Return values

Entity operations resolve to the entity data directly — there is no
result envelope:

- `load` resolves to a single entity object.
- `list` resolves to an **array** of entity objects (iterate it directly;
  there is no `.data` and no `.ok`).

On a failed request these methods **throw**, so wrap calls in
`try`/`catch` to handle errors. Only `direct()` returns the result
envelope described below.

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
| `author` | `string` |  |
| `line` | `any[]` |  |
| `linecount` | `number` |  |
| `title` | `string` |  |

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
| `author` | `string` |  |
| `line` | `any[]` |  |
| `linecount` | `number` |  |
| `title` | `string` |  |

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
| `author` | `string` |  |
| `line` | `any[]` |  |
| `linecount` | `number` |  |
| `title` | `string` |  |

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
| `author` | `string` |  |
| `line` | `any[]` |  |
| `linecount` | `number` |  |
| `title` | `string` |  |

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
| `author` | `string` |  |
| `line` | `any[]` |  |
| `linecount` | `number` |  |
| `title` | `string` |  |

#### Example: Load

```ts
const linecount = await client.Linecount().load({ id: 1 })
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
| `author` | `string` |  |
| `line` | `any[]` |  |
| `linecount` | `number` |  |
| `title` | `string` |  |

#### Example: Load

```ts
const poemcount = await client.Poemcount().load({ id: 1 })
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
| `author` | `string` |  |
| `line` | `any[]` |  |
| `linecount` | `number` |  |
| `title` | `string` |  |

#### Example: Load

```ts
const random = await client.Random().load({ id: 1 })
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
| `author` | `string` |  |
| `line` | `any[]` |  |
| `linecount` | `number` |  |
| `title` | `string` |  |

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
| `author` | `string` |  |
| `line` | `any[]` |  |
| `linecount` | `number` |  |
| `title` | `string` |  |

#### Example: List

```ts
const titleabs = await client.Titleab().list()
```


## Advanced

> The sections above cover everyday use. The material below explains the
> SDK's internals — useful when extending it with custom features, but not
> needed for normal use.

### The operation pipeline

Every entity operation follows a six-stage pipeline. Each stage fires a
feature hook before executing:

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

If any stage errors, the pipeline short-circuits and the error surfaces
to the caller — see [Error handling](#error-handling) for how that looks
in this language.

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
import { PoetrydbSDK } from '@voxgig-sdk/poetrydb'
```

### Entity state

Entity instances are stateful. After a successful `list`, the entity
stores the returned data and match criteria internally. Subsequent
calls on the same instance can rely on this state.

```ts
const author = client.Author()
await author.list()

// author.data() now returns the author data from the last `list`
// author.match() returns the last match criteria
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
