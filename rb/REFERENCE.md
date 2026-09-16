# Poetrydb Ruby SDK Reference

Complete API reference for the Poetrydb Ruby SDK.


## PoetrydbSDK

### Constructor

```ruby
require_relative 'Poetrydb_sdk'

client = PoetrydbSDK.new(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `Hash` | SDK configuration options. |
| `options["base"]` | `String` | Base URL for API requests. |
| `options["prefix"]` | `String` | URL prefix appended after base. |
| `options["suffix"]` | `String` | URL suffix appended after path. |
| `options["headers"]` | `Hash` | Custom headers for all requests. |
| `options["feature"]` | `Hash` | Feature configuration. |
| `options["system"]` | `Hash` | System overrides (e.g. custom fetch). |


### Static Methods

#### `PoetrydbSDK.test(testopts = nil, sdkopts = nil)`

Create a test client with mock features active. Both arguments may be `nil`.

```ruby
client = PoetrydbSDK.test
```


### Instance Methods

#### `Author(data = nil)`

Create a new `Author` entity instance. Pass `nil` for no initial data.

#### `Authorab(data = nil)`

Create a new `Authorab` entity instance. Pass `nil` for no initial data.

#### `CombinedSearch(data = nil)`

Create a new `CombinedSearch` entity instance. Pass `nil` for no initial data.

#### `CombinedSearchWithField(data = nil)`

Create a new `CombinedSearchWithField` entity instance. Pass `nil` for no initial data.

#### `Line(data = nil)`

Create a new `Line` entity instance. Pass `nil` for no initial data.

#### `Linecount(data = nil)`

Create a new `Linecount` entity instance. Pass `nil` for no initial data.

#### `Poemcount(data = nil)`

Create a new `Poemcount` entity instance. Pass `nil` for no initial data.

#### `Random(data = nil)`

Create a new `Random` entity instance. Pass `nil` for no initial data.

#### `Title(data = nil)`

Create a new `Title` entity instance. Pass `nil` for no initial data.

#### `Titleab(data = nil)`

Create a new `Titleab` entity instance. Pass `nil` for no initial data.

#### `options_map -> Hash`

Return a deep copy of the current SDK options.

#### `get_utility -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs = {}) -> Hash`

Make a direct HTTP request to any API endpoint. Returns a result hash
(`{ "ok" => ..., "status" => ..., "data" => ..., "err" => ... }`); it
does not raise — inspect `result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `String` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `String` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `Hash` | Path parameter values for `{param}` substitution. |
| `fetchargs["query"]` | `Hash` | Query string parameters. |
| `fetchargs["headers"]` | `Hash` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (hashes are JSON-serialized). |
| `fetchargs["ctrl"]` | `Hash` | Control options (e.g. `{ "explain" => true }`). |

**Returns:** `Hash`

#### `prepare(fetchargs = {}) -> Hash`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`. Raises on error.

**Returns:** `Hash` (the fetch definition; raises on error)


---

## AuthorEntity

```ruby
author = client.Author
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `author` | `String` | No | The author of the poem |
| `authors` | `Array` | No |  |
| `id` | `String` | No |  |
| `linecount` | `Integer` | No | The number of lines in the poem (including section headings, excluding empty lines) |
| `lines` | `Array` | No | The lines of the poem |
| `title` | `String` | No | The title of the poem |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Author.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Author.load({ "id" => "author_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `AuthorEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## AuthorabEntity

```ruby
authorab = client.Authorab
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `author` | `String` | No | The author of the poem |
| `linecount` | `Integer` | No | The number of lines in the poem (including section headings, excluding empty lines) |
| `lines` | `Array` | No | The lines of the poem |
| `title` | `String` | No | The title of the poem |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Authorab.list
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `AuthorabEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## CombinedSearchEntity

```ruby
combined_search = client.CombinedSearch
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `author` | `String` | No | The author of the poem |
| `linecount` | `Integer` | No | The number of lines in the poem (including section headings, excluding empty lines) |
| `lines` | `Array` | No | The lines of the poem |
| `title` | `String` | No | The title of the poem |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.CombinedSearch.list
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `CombinedSearchEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## CombinedSearchWithFieldEntity

```ruby
combined_search_with_field = client.CombinedSearchWithField
```

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.CombinedSearchWithField.list
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `CombinedSearchWithFieldEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## LineEntity

```ruby
line = client.Line
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `author` | `String` | No | The author of the poem |
| `id` | `String` | No |  |
| `linecount` | `Integer` | No | The number of lines in the poem (including section headings, excluding empty lines) |
| `lines` | `Array` | No | The lines of the poem |
| `title` | `String` | No | The title of the poem |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Line.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Line.load({ "id" => "line_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `LineEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## LinecountEntity

```ruby
linecount = client.Linecount
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `author` | `String` | No | The author of the poem |
| `id` | `String` | No |  |
| `linecount` | `Integer` | No | The number of lines in the poem (including section headings, excluding empty lines) |
| `lines` | `Array` | No | The lines of the poem |
| `title` | `String` | No | The title of the poem |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Linecount.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Linecount.load({ "id" => 1 })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `LinecountEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## PoemcountEntity

```ruby
poemcount = client.Poemcount
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `author` | `String` | No | The author of the poem |
| `id` | `String` | No |  |
| `linecount` | `Integer` | No | The number of lines in the poem (including section headings, excluding empty lines) |
| `lines` | `Array` | No | The lines of the poem |
| `title` | `String` | No | The title of the poem |

### Operations

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Poemcount.load({ "id" => 1 })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `PoemcountEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## RandomEntity

```ruby
random = client.Random
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `author` | `String` | No | The author of the poem |
| `id` | `String` | No |  |
| `linecount` | `Integer` | No | The number of lines in the poem (including section headings, excluding empty lines) |
| `lines` | `Array` | No | The lines of the poem |
| `title` | `String` | No | The title of the poem |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Random.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Random.load({ "id" => 1 })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `RandomEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## TitleEntity

```ruby
title = client.Title
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `author` | `String` | No | The author of the poem |
| `id` | `String` | No |  |
| `linecount` | `Integer` | No | The number of lines in the poem (including section headings, excluding empty lines) |
| `lines` | `Array` | No | The lines of the poem |
| `title` | `String` | No | The title of the poem |
| `titles` | `Array` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Title.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Title.load({ "id" => "title_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `TitleEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## TitleabEntity

```ruby
titleab = client.Titleab
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `author` | `String` | No | The author of the poem |
| `linecount` | `Integer` | No | The number of lines in the poem (including section headings, excluding empty lines) |
| `lines` | `Array` | No | The lines of the poem |
| `title` | `String` | No | The title of the poem |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Titleab.list
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `TitleabEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `ratelimit` | 0.0.1 | Client-side rate limiting via a token bucket |
| `retry` | 0.0.1 | Automatic retry of transient failures with exponential backoff |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |
| `timeout` | 0.0.1 | Per-request timeout with transport abort |


Features are activated via the `feature` option:

```ruby
client = PoetrydbSDK.new({
  "feature" => {
    "ratelimit" => { "active" => true },
    "retry" => { "active" => true },
    "test" => { "active" => true },
    "timeout" => { "active" => true },
  },
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

