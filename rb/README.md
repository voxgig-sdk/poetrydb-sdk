# Poetrydb Ruby SDK



The Ruby SDK for the Poetrydb API — an entity-oriented client using idiomatic Ruby conventions.

The SDK exposes the API as capitalised, semantic **Entities** — for example `client.Author` — with named operations (`list`/`load`) instead of raw URL paths and query strings. Working with resources and verbs keeps call sites self-describing and reduces cognitive load.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to RubyGems. Install it from the
GitHub release tag (`rb/vX.Y.Z`):

- Releases: [https://github.com/voxgig-sdk/poetrydb-sdk/releases](https://github.com/voxgig-sdk/poetrydb-sdk/releases)


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```ruby
require_relative "Poetrydb_sdk"

client = PoetrydbSDK.new
```

### 2. List author records

```ruby
begin
  # list returns an Array of Author records — iterate directly.
  authors = client.Author.list
  authors.each do |item|
    puts "#{item["id"]} #{item["author"]}"
  end
rescue => err
  warn "list failed: #{err}"
end
```

### 3. Load an author

```ruby
begin
  # load returns the ENTITY — call data_get for the Author record (raises on error).
  author = client.Author.load({ "id" => "example_id" })
  puts author
rescue => err
  warn "load failed: #{err}"
end
```


## Error handling

Entity operations raise on failure, so rescue them:

```ruby
begin
  linecounts = client.Linecount.list()
rescue => err
  warn "list failed: #{err}"
end
```

`direct` does **not** raise — it returns the result hash. Branch on
`ok`; on failure `status` holds the HTTP status (for error responses) and
`err` holds a transport error, so read both defensively:

```ruby
result = client.direct({
  "path" => "/api/resource/{id}",
  "method" => "GET",
  "params" => { "id" => "example_id" },
})

warn "request failed: #{result["err"] || "HTTP #{result["status"]}"}" unless result["ok"]
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```ruby
result = client.direct({
  "path" => "/api/resource/{id}",
  "method" => "GET",
  "params" => { "id" => "example" },
})

if result["ok"]
  puts result["status"]  # 200
  puts result["data"]    # response body
else
  # On an HTTP error status there is no err (only a transport failure sets
  # it), so fall back to the status code.
  warn(result["err"] || "HTTP #{result["status"]}")
end
```

### Prepare a request without sending it

```ruby
begin
  fetchdef = client.prepare({
    "path" => "/api/resource/{id}",
    "method" => "DELETE",
    "params" => { "id" => "example" },
  })
  puts fetchdef["url"]
  puts fetchdef["method"]
  puts fetchdef["headers"]
rescue => err
  warn "prepare failed: #{err}"
end
```

### Use test mode

Create a mock client for unit testing — no server required. Seed fixture
data via the `entity` option so offline calls resolve without a live server:

```ruby
client = PoetrydbSDK.test({
  "entity" => { "linecount" => { "test01" => { "id" => "test01" } } },
})

# Entity ops return the ENTITY (raises on error);
# call data_get for the mock record.
linecount = client.Linecount.list()
puts linecount
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```ruby
mock_fetch = ->(url, init) {
  return {
    "status" => 200,
    "statusText" => "OK",
    "headers" => {},
    "json" => ->() { { "id" => "mock01" } },
  }, nil
}

client = PoetrydbSDK.new({
  "base" => "http://localhost:8080",
  "system" => {
    "fetch" => mock_fetch,
  },
})
```

### Run live tests

Create a `.env.local` file at the project root:

```
POETRYDB_TEST_LIVE=TRUE
```

Then run:

```bash
cd rb && ruby -Itest -e "Dir['test/*_test.rb'].each { |f| require_relative f }"
```


## Reference

### PoetrydbSDK

```ruby
require_relative "Poetrydb_sdk"
client = PoetrydbSDK.new(options)
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `base` | `String` | Base URL of the API server. |
| `prefix` | `String` | URL path prefix prepended to all requests. |
| `suffix` | `String` | URL path suffix appended to all requests. |
| `feature` | `Hash` | Feature activation flags. |
| `extend` | `Hash` | Additional Feature instances to load. |
| `system` | `Hash` | System overrides (e.g. custom `fetch` lambda). |

### test

```ruby
client = PoetrydbSDK.test(testopts, sdkopts)
```

Creates a test-mode client with mock transport. Both arguments may be `nil`.

### PoetrydbSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `() -> Hash` | Deep copy of current SDK options. |
| `get_utility` | `() -> Utility` | Copy of the SDK utility object. |
| `prepare` | `(fetchargs) -> Hash` | Build an HTTP request definition without sending. Raises on error. |
| `direct` | `(fetchargs) -> Hash` | Build and send an HTTP request. Returns a result hash (`result["ok"]`); does not raise. |
| `Author` | `(data) -> AuthorEntity` | Create an Author entity instance. |
| `Authorab` | `(data) -> AuthorabEntity` | Create an Authorab entity instance. |
| `CombinedSearch` | `(data) -> CombinedSearchEntity` | Create a CombinedSearch entity instance. |
| `CombinedSearchWithField` | `(data) -> CombinedSearchWithFieldEntity` | Create a CombinedSearchWithField entity instance. |
| `Line` | `(data) -> LineEntity` | Create a Line entity instance. |
| `Linecount` | `(data) -> LinecountEntity` | Create a Linecount entity instance. |
| `Poemcount` | `(data) -> PoemcountEntity` | Create a Poemcount entity instance. |
| `Random` | `(data) -> RandomEntity` | Create a Random entity instance. |
| `Title` | `(data) -> TitleEntity` | Create a Title entity instance. |
| `Titleab` | `(data) -> TitleabEntity` | Create a Titleab entity instance. |

### Entity interface

All entities share the same interface.

| Method | Signature | Description |
| --- | --- | --- |
| `load` | `(reqmatch, ctrl) -> any` | Load a single entity by match criteria. Raises on error. |
| `list` | `(reqmatch = nil, ctrl) -> Array` | List entities matching the criteria (call with no argument to list all). Raises on error. |
| `data_get` | `() -> Hash` | Get entity data. |
| `data_set` | `(data)` | Set entity data. |
| `match_get` | `() -> Hash` | Get entity match criteria. |
| `match_set` | `(match)` | Set entity match criteria. |
| `make` | `() -> Entity` | Create a new instance with the same options. |
| `get_name` | `() -> String` | Return the entity name. |

### Result shape

Entity operations return the result data directly. On failure they
raise a `PoetrydbError` (a `StandardError` subclass), so wrap
calls in `begin`/`rescue` where you need to handle errors.

The `direct` escape hatch is the exception: it never raises and instead
returns a result `Hash` with these keys:

| Key | Type | Description |
| --- | --- | --- |
| `ok` | `Boolean` | `true` if the HTTP status is 2xx. |
| `status` | `Integer` | HTTP status code. |
| `headers` | `Hash` | Response headers. |
| `data` | `any` | Parsed JSON response body. |
| `err` | `Error` | Present when `ok` is `false`. |

### Entities

#### Author

| Field | Description |
| --- | --- |
| `author` | The author of the poem |
| `authors` |  |
| `id` |  |
| `linecount` | The number of lines in the poem (including section headings, excluding empty lines) |
| `lines` | The lines of the poem |
| `title` | The title of the poem |

Operations: List, Load.

API path: `/author/{author}/{outputFields}.{format}`

#### Authorab

| Field | Description |
| --- | --- |
| `author` | The author of the poem |
| `linecount` | The number of lines in the poem (including section headings, excluding empty lines) |
| `lines` | The lines of the poem |
| `title` | The title of the poem |

Operations: List.

API path: `/author/{author}:abs`

#### CombinedSearch

| Field | Description |
| --- | --- |
| `author` | The author of the poem |
| `linecount` | The number of lines in the poem (including section headings, excluding empty lines) |
| `lines` | The lines of the poem |
| `title` | The title of the poem |

Operations: List.

API path: `/{inputField1},{inputField2}/{searchTerm1};{searchTerm2}`

#### CombinedSearchWithField

| Field | Description |
| --- | --- |

Operations: List.

API path: `/{inputField1},{inputField2}/{searchTerm1};{searchTerm2}/{outputFields}`

#### Line

| Field | Description |
| --- | --- |
| `author` | The author of the poem |
| `id` |  |
| `linecount` | The number of lines in the poem (including section headings, excluding empty lines) |
| `lines` | The lines of the poem |
| `title` | The title of the poem |

Operations: List, Load.

API path: `/lines/{lines}/{outputFields}.{format}`

#### Linecount

| Field | Description |
| --- | --- |
| `author` | The author of the poem |
| `id` |  |
| `linecount` | The number of lines in the poem (including section headings, excluding empty lines) |
| `lines` | The lines of the poem |
| `title` | The title of the poem |

Operations: List, Load.

API path: `/linecount/{linecount}/{outputFields}.{format}`

#### Poemcount

| Field | Description |
| --- | --- |
| `author` | The author of the poem |
| `id` |  |
| `linecount` | The number of lines in the poem (including section headings, excluding empty lines) |
| `lines` | The lines of the poem |
| `title` | The title of the poem |

Operations: Load.

API path: `/poemcount/{count}`

#### Random

| Field | Description |
| --- | --- |
| `author` | The author of the poem |
| `id` |  |
| `linecount` | The number of lines in the poem (including section headings, excluding empty lines) |
| `lines` | The lines of the poem |
| `title` | The title of the poem |

Operations: List, Load.

API path: `/random/{count}/{outputFields}`

#### Title

| Field | Description |
| --- | --- |
| `author` | The author of the poem |
| `id` |  |
| `linecount` | The number of lines in the poem (including section headings, excluding empty lines) |
| `lines` | The lines of the poem |
| `title` | The title of the poem |
| `titles` |  |

Operations: List, Load.

API path: `/title/{title}/{outputFields}.{format}`

#### Titleab

| Field | Description |
| --- | --- |
| `author` | The author of the poem |
| `linecount` | The number of lines in the poem (including section headings, excluding empty lines) |
| `lines` | The lines of the poem |
| `title` | The title of the poem |

Operations: List.

API path: `/title/{title}:abs`



## Entities


### Author

Create an instance: `author = client.Author`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `author` | `String` | The author of the poem |
| `authors` | `Array` |  |
| `id` | `String` |  |
| `linecount` | `Integer` | The number of lines in the poem (including section headings, excluding empty lines) |
| `lines` | `Array` | The lines of the poem |
| `title` | `String` | The title of the poem |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the Author record (raises on error).
author = client.Author.load({ "id" => "author_id" })
```

#### Example: List

```ruby
# list returns an Array of Author records (raises on error).
authors = client.Author.list
```


### Authorab

Create an instance: `authorab = client.Authorab`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `author` | `String` | The author of the poem |
| `linecount` | `Integer` | The number of lines in the poem (including section headings, excluding empty lines) |
| `lines` | `Array` | The lines of the poem |
| `title` | `String` | The title of the poem |

#### Example: List

```ruby
# list returns an Array of Authorab records (raises on error).
authorabs = client.Authorab.list
```


### CombinedSearch

Create an instance: `combined_search = client.CombinedSearch`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `author` | `String` | The author of the poem |
| `linecount` | `Integer` | The number of lines in the poem (including section headings, excluding empty lines) |
| `lines` | `Array` | The lines of the poem |
| `title` | `String` | The title of the poem |

#### Example: List

```ruby
# list returns an Array of CombinedSearch records (raises on error).
combined_searchs = client.CombinedSearch.list
```


### CombinedSearchWithField

Create an instance: `combined_search_with_field = client.CombinedSearchWithField`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Example: List

```ruby
# list returns an Array of CombinedSearchWithField records (raises on error).
combined_search_with_fields = client.CombinedSearchWithField.list
```


### Line

Create an instance: `line = client.Line`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `author` | `String` | The author of the poem |
| `id` | `String` |  |
| `linecount` | `Integer` | The number of lines in the poem (including section headings, excluding empty lines) |
| `lines` | `Array` | The lines of the poem |
| `title` | `String` | The title of the poem |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the Line record (raises on error).
line = client.Line.load({ "id" => "line_id" })
```

#### Example: List

```ruby
# list returns an Array of Line records (raises on error).
lines = client.Line.list
```


### Linecount

Create an instance: `linecount = client.Linecount`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `author` | `String` | The author of the poem |
| `id` | `String` |  |
| `linecount` | `Integer` | The number of lines in the poem (including section headings, excluding empty lines) |
| `lines` | `Array` | The lines of the poem |
| `title` | `String` | The title of the poem |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the Linecount record (raises on error).
linecount = client.Linecount.load({ "id" => 1 })
```

#### Example: List

```ruby
# list returns an Array of Linecount records (raises on error).
linecounts = client.Linecount.list
```


### Poemcount

Create an instance: `poemcount = client.Poemcount`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `author` | `String` | The author of the poem |
| `id` | `String` |  |
| `linecount` | `Integer` | The number of lines in the poem (including section headings, excluding empty lines) |
| `lines` | `Array` | The lines of the poem |
| `title` | `String` | The title of the poem |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the Poemcount record (raises on error).
poemcount = client.Poemcount.load({ "id" => 1 })
```


### Random

Create an instance: `random = client.Random`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `author` | `String` | The author of the poem |
| `id` | `String` |  |
| `linecount` | `Integer` | The number of lines in the poem (including section headings, excluding empty lines) |
| `lines` | `Array` | The lines of the poem |
| `title` | `String` | The title of the poem |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the Random record (raises on error).
random = client.Random.load({ "id" => 1 })
```

#### Example: List

```ruby
# list returns an Array of Random records (raises on error).
randoms = client.Random.list
```


### Title

Create an instance: `title = client.Title`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `author` | `String` | The author of the poem |
| `id` | `String` |  |
| `linecount` | `Integer` | The number of lines in the poem (including section headings, excluding empty lines) |
| `lines` | `Array` | The lines of the poem |
| `title` | `String` | The title of the poem |
| `titles` | `Array` |  |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the Title record (raises on error).
title = client.Title.load({ "id" => "title_id" })
```

#### Example: List

```ruby
# list returns an Array of Title records (raises on error).
titles = client.Title.list
```


### Titleab

Create an instance: `titleab = client.Titleab`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `author` | `String` | The author of the poem |
| `linecount` | `Integer` | The number of lines in the poem (including section headings, excluding empty lines) |
| `lines` | `Array` | The lines of the poem |
| `title` | `String` | The title of the poem |

#### Example: List

```ruby
# list returns an Array of Titleab records (raises on error).
titleabs = client.Titleab.list
```

## Features

This SDK ships 4 optional features. Each is **inactive until you
switch it on**, so an SDK you have not configured behaves exactly as if none of
them existed — no retries, no cache, no logging, no measurable overhead.

Activate a feature by name in the client options, alongside the options shown
above:

| Feature | What it does |
|---|---|
| [`ratelimit`](#ratelimit) | Client-side rate limiting via a token bucket |
| [`retry`](#retry) | Automatic retry of transient failures with exponential backoff |
| [`test`](#test) | In-memory mock transport for testing without a live server |
| [`timeout`](#timeout) | Per-request timeout with transport abort |

> **Order matters for `ratelimit`, `retry`, `timeout`.** These wrap the
> transport, so each one wraps whatever is already installed: the order you
> activate them in IS the nesting order. Activating them as an ordered list
> rather than a map is what fixes that order.

### ratelimit

Client-side rate limiting via a token bucket.

| Option | Default |
|---|---|
| `active` | `false` |
| `burst` | `5` |
| `rate` | `5` |

Set `feature.ratelimit.active` to enable it, then override any of the options above.

`ratelimit` wraps the transport, so its position among the other
transport features decides what it sees. A feature activated later wraps one
activated earlier.

### retry

Automatic retry of transient failures with exponential backoff.

| Option | Default |
|---|---|
| `active` | `false` |
| `factor` | `2` |
| `maxDelay` | `2000` |
| `minDelay` | `50` |
| `retries` | `2` |
| `statuses` | `[408, 425, 429, 500, 502, 503, 504]` |

Set `feature.retry.active` to enable it, then override any of the options above.

`retry` wraps the transport, so its position among the other
transport features decides what it sees. A feature activated later wraps one
activated earlier.

### test

In-memory mock transport for testing without a live server.

| Option | Default |
|---|---|
| `active` | `false` |

Set `feature.test.active` to enable it, then override any of the options above.

### timeout

Per-request timeout with transport abort.

| Option | Default |
|---|---|
| `active` | `false` |
| `ms` | `30000` |

Set `feature.timeout.active` to enable it, then override any of the options above.

`timeout` wraps the transport, so its position among the other
transport features decides what it sees. A feature activated later wraps one
activated earlier.


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

Features are the extension mechanism. A feature is a Ruby class
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **RatelimitFeature**: Client-side rate limiting via a token bucket
- **RetryFeature**: Automatic retry of transient failures with exponential backoff
- **TestFeature**: In-memory mock transport for testing without a live server
- **TimeoutFeature**: Per-request timeout with transport abort

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as hashes

The Ruby SDK uses plain Ruby hashes throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `Helpers.to_map()` to safely validate that a value is a hash.

### Module structure

```
rb/
├── Poetrydb_sdk.rb       -- Main SDK module
├── config.rb                  -- Configuration
├── schema.rb                  -- Generated option + entity specs
├── features.rb                -- Feature factory
├── core/                      -- Core types and context
├── entity/                    -- Entity implementations
├── feature/                   -- Built-in features (Base, Test, Log)
├── utility/                   -- Utility functions and struct library
└── test/                      -- Test suites
```

The main module (`Poetrydb_sdk`) exports the SDK class
and test helper. Import entity or utility modules directly only
when needed.

### Entity state

Entity instances are stateful. After a successful `list`, the entity
stores the returned data and match criteria internally.

```ruby
linecount = client.Linecount
linecount.list()

# linecount.data_get now returns the linecount data from the last list
# linecount.match_get returns the last match criteria
```

Call `make` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

`direct` gives full control over the HTTP request. Use it for
non-standard endpoints, bulk operations, or any path not modelled as
an entity. `prepare` builds the request without sending it — useful
for debugging or custom transport.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
