# Poetrydb TypeScript SDK Reference

Complete API reference for the Poetrydb TypeScript SDK.


## PoetrydbSDK

### Constructor

```ts
new PoetrydbSDK(options?: object)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `object` | SDK configuration options. |
| `options.base` | `string` | Base URL for API requests. |
| `options.prefix` | `string` | URL prefix appended after base. |
| `options.suffix` | `string` | URL suffix appended after path. |
| `options.headers` | `object` | Custom headers for all requests. |
| `options.feature` | `object` | Feature configuration. |
| `options.system` | `object` | System overrides (e.g. custom fetch). |


### Static Methods

#### `PoetrydbSDK.test(testopts?, sdkopts?)`

Create a test client with mock features active.

```ts
const client = PoetrydbSDK.test()
```

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `testopts` | `object` | Test feature options. |
| `sdkopts` | `object` | Additional SDK options merged with test defaults. |

**Returns:** `PoetrydbSDK` instance in test mode.


### Instance Methods

#### `Author(data?: object)`

Create a new `Author` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `AuthorEntity` instance.

#### `Authorab(data?: object)`

Create a new `Authorab` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `AuthorabEntity` instance.

#### `CombinedSearch(data?: object)`

Create a new `CombinedSearch` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `CombinedSearchEntity` instance.

#### `CombinedSearchWithField(data?: object)`

Create a new `CombinedSearchWithField` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `CombinedSearchWithFieldEntity` instance.

#### `Line(data?: object)`

Create a new `Line` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `LineEntity` instance.

#### `Linecount(data?: object)`

Create a new `Linecount` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `LinecountEntity` instance.

#### `Poemcount(data?: object)`

Create a new `Poemcount` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `PoemcountEntity` instance.

#### `Random(data?: object)`

Create a new `Random` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `RandomEntity` instance.

#### `Title(data?: object)`

Create a new `Title` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `TitleEntity` instance.

#### `Titleab(data?: object)`

Create a new `Titleab` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `TitleabEntity` instance.

#### `options()`

Return a deep copy of the current SDK options.

**Returns:** `object`

#### `utility()`

Return a copy of the SDK utility object.

**Returns:** `object`

#### `direct(fetchargs?: object)`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs.path` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs.method` | `string` | HTTP method (default: `GET`). |
| `fetchargs.params` | `object` | Path parameter values for `{param}` substitution. |
| `fetchargs.query` | `object` | Query string parameters. |
| `fetchargs.headers` | `object` | Request headers (merged with defaults). |
| `fetchargs.body` | `any` | Request body (objects are JSON-serialized). |
| `fetchargs.ctrl` | `object` | Control options (e.g. `{ explain: true }`). |

**Returns:** `Promise<{ ok, status, headers, data } | Error>`

#### `prepare(fetchargs?: object)`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`.

**Returns:** `Promise<{ url, method, headers, body } | Error>`

#### `tester(testopts?, sdkopts?)`

Alias for `PoetrydbSDK.test()`.

**Returns:** `PoetrydbSDK` instance in test mode.


---

## AuthorEntity

```ts
const author = client.Author()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `author` | `string` | No | The author of the poem |
| `authors` | `any[]` | No |  |
| `id` | `string` | No |  |
| `linecount` | `number` | No | The number of lines in the poem (including section headings, excluding empty lines) |
| `lines` | `any[]` | No | The lines of the poem |
| `title` | `string` | No | The title of the poem |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Author().list({ author: "example", format: "example", output_field: "example" })
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Author().load({ id: 'author_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `AuthorEntity` instance with the same client and
options.

#### `client()`

Return the parent `PoetrydbSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## AuthorabEntity

```ts
const authorab = client.Authorab()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `author` | `string` | No | The author of the poem |
| `linecount` | `number` | No | The number of lines in the poem (including section headings, excluding empty lines) |
| `lines` | `any[]` | No | The lines of the poem |
| `title` | `string` | No | The title of the poem |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Authorab().list({ author: "example" })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `AuthorabEntity` instance with the same client and
options.

#### `client()`

Return the parent `PoetrydbSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## CombinedSearchEntity

```ts
const combined_search = client.CombinedSearch()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `author` | `string` | No | The author of the poem |
| `linecount` | `number` | No | The number of lines in the poem (including section headings, excluding empty lines) |
| `lines` | `any[]` | No | The lines of the poem |
| `title` | `string` | No | The title of the poem |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.CombinedSearch().list({ input_field1: "example", input_field2: "example", search_term1: "example", search_term2: "example" })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `CombinedSearchEntity` instance with the same client and
options.

#### `client()`

Return the parent `PoetrydbSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## CombinedSearchWithFieldEntity

```ts
const combined_search_with_field = client.CombinedSearchWithField()
```

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.CombinedSearchWithField().list({ input_field1: "example", input_field2: "example", output_field: "example", search_term1: "example", search_term2: "example" })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `CombinedSearchWithFieldEntity` instance with the same client and
options.

#### `client()`

Return the parent `PoetrydbSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## LineEntity

```ts
const line = client.Line()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `author` | `string` | No | The author of the poem |
| `id` | `string` | No |  |
| `linecount` | `number` | No | The number of lines in the poem (including section headings, excluding empty lines) |
| `lines` | `any[]` | No | The lines of the poem |
| `title` | `string` | No | The title of the poem |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Line().list({ line: "example", output_field: "example" })
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Line().load({ id: 'line_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `LineEntity` instance with the same client and
options.

#### `client()`

Return the parent `PoetrydbSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## LinecountEntity

```ts
const linecount = client.Linecount()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `author` | `string` | No | The author of the poem |
| `id` | `string` | No |  |
| `linecount` | `number` | No | The number of lines in the poem (including section headings, excluding empty lines) |
| `lines` | `any[]` | No | The lines of the poem |
| `title` | `string` | No | The title of the poem |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Linecount().list({ linecount: 1, output_field: "example" })
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Linecount().load({ id: 1 })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `LinecountEntity` instance with the same client and
options.

#### `client()`

Return the parent `PoetrydbSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## PoemcountEntity

```ts
const poemcount = client.Poemcount()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `author` | `string` | No | The author of the poem |
| `id` | `string` | No |  |
| `linecount` | `number` | No | The number of lines in the poem (including section headings, excluding empty lines) |
| `lines` | `any[]` | No | The lines of the poem |
| `title` | `string` | No | The title of the poem |

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Poemcount().load({ id: 1 })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `PoemcountEntity` instance with the same client and
options.

#### `client()`

Return the parent `PoetrydbSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## RandomEntity

```ts
const random = client.Random()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `author` | `string` | No | The author of the poem |
| `id` | `string` | No |  |
| `linecount` | `number` | No | The number of lines in the poem (including section headings, excluding empty lines) |
| `lines` | `any[]` | No | The lines of the poem |
| `title` | `string` | No | The title of the poem |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Random().list({ count: 1, output_field: "example" })
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Random().load({ id: 1 })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `RandomEntity` instance with the same client and
options.

#### `client()`

Return the parent `PoetrydbSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## TitleEntity

```ts
const title = client.Title()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `author` | `string` | No | The author of the poem |
| `id` | `string` | No |  |
| `linecount` | `number` | No | The number of lines in the poem (including section headings, excluding empty lines) |
| `lines` | `any[]` | No | The lines of the poem |
| `title` | `string` | No | The title of the poem |
| `titles` | `any[]` | No |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Title().list({ format: "example", output_field: "example", title: "example" })
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Title().load({ id: 'title_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `TitleEntity` instance with the same client and
options.

#### `client()`

Return the parent `PoetrydbSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## TitleabEntity

```ts
const titleab = client.Titleab()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `author` | `string` | No | The author of the poem |
| `linecount` | `number` | No | The number of lines in the poem (including section headings, excluding empty lines) |
| `lines` | `any[]` | No | The lines of the poem |
| `title` | `string` | No | The title of the poem |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Titleab().list({ title: "example" })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `TitleabEntity` instance with the same client and
options.

#### `client()`

Return the parent `PoetrydbSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `ratelimit` | 0.0.1 | Client-side rate limiting via a token bucket |
| `retry` | 0.0.1 | Automatic retry of transient failures with exponential backoff |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |
| `timeout` | 0.0.1 | Per-request timeout with transport abort |


Features are activated via the `feature` option:

```ts
const client = new PoetrydbSDK({
  feature: {
    ratelimit: { active: true },
    retry: { active: true },
    test: { active: true },
    timeout: { active: true },
  }
})
```


### Configuring features

Each feature is inactive until switched on, and an SDK with no feature
configured does no feature work at all. Every option below keeps its default
unless you name it.

The array form of \`feature\` is significant: several features wrap the
transport, and the order you list them in is the order they nest.

#### Ordering

`ratelimit`, `retry`, `timeout` wrap the transport. Each
wraps whatever is already installed, so **activation order is nesting order**:
a feature activated later sits OUTSIDE one activated earlier, and sees the call
first.

That decides behaviour, not just sequence: a feature that short-circuits the
call, such as a cache serving a hit, stops every feature nested inside it from
ever seeing that call.

`test` attach to pipeline hooks
rather than the transport, so their order does not affect what they observe.

#### `ratelimit`

Client-side rate limiting via a token bucket.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |
| `burst` | `5` |
| `rate` | `5` |

| Option | Type |
|---|---|
| `now` | function |
| `sleep` | function |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.ratelimit.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Wraps the transport: its place in the activation order decides what it
  sees. See [Ordering](#ordering) above.
- Inactive by default: leaving it out costs nothing at runtime.

#### `retry`

Automatic retry of transient failures with exponential backoff.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |
| `factor` | `2` |
| `maxDelay` | `2000` |
| `minDelay` | `50` |
| `retries` | `2` |
| `statuses` | `[408, 425, 429, 500, 502, 503, 504]` |

| Option | Type |
|---|---|
| `jitter` | boolean |
| `sleep` | function |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.retry.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Wraps the transport: its place in the activation order decides what it
  sees. See [Ordering](#ordering) above.
- Inactive by default: leaving it out costs nothing at runtime.

#### `test`

In-memory mock transport for testing without a live server.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |

| Option | Type |
|---|---|
| `entity` | map |
| `net` | map |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.test.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Attaches to pipeline hooks, not the transport, so activation order does
  not change what it observes.
- Installs the BASE transport that the wrapping features wrap, so it must be
  activated before them.
- Inactive by default: leaving it out costs nothing at runtime.

#### `timeout`

Per-request timeout with transport abort.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |
| `ms` | `30000` |

| Option | Type |
|---|---|
| `clearTimer` | function |
| `setTimer` | function |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.timeout.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Wraps the transport: its place in the activation order decides what it
  sees. See [Ordering](#ordering) above.
- Inactive by default: leaving it out costs nothing at runtime.

