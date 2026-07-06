# Poetrydb Lua SDK



The Lua SDK for the Poetrydb API — an entity-oriented client using Lua conventions.

It exposes the API as capitalised, semantic **Entities** — e.g. `client:Author()` — each with the same small set of operations (`list`, `load`) instead of raw URL paths and query strings. You call meaning, not endpoints, which keeps the cognitive load low.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to LuaRocks. Install it from the
GitHub release tag (`lua/vX.Y.Z`, see [Releases](https://github.com/voxgig-sdk/poetrydb-sdk/releases)),
or add the source directory to your `LUA_PATH`:

```bash
export LUA_PATH="path/to/lua/?.lua;path/to/lua/?/init.lua;;"
```


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```lua
local sdk = require("poetrydb_sdk")

local client = sdk.new()
```

### 2. List author records

Entity operations return `(value, err)`. For `list`, `value` is the
array of records itself — iterate it directly (there is no wrapper).

```lua
local authors, err = client:Author():list()
if err then error(err) end

for _, item in ipairs(authors) do
  print(item["author"])
end
```

### 3. Load an author

```lua
local author, err = client:Author():load({ id = "example_id" })
if err then error(err) end
print(author)
```


## Error handling

Entity operations return `(value, err)`. Check `err` before using
the value:

```lua
local authors, err = client:Author():list()
if err then error(err) end
```

`direct` follows the same `(value, err)` convention:

```lua
local result, err = client:direct({
  path = "/api/resource/{id}",
  method = "GET",
  params = { id = "example_id" },
})
if err then error(err) end
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```lua
local result, err = client:direct({
  path = "/api/resource/{id}",
  method = "GET",
  params = { id = "example" },
})
if err then error(err) end

if result["ok"] then
  print(result["status"])  -- 200
  print(result["data"])    -- response body
end
```

### Prepare a request without sending it

```lua
local fetchdef, err = client:prepare({
  path = "/api/resource/{id}",
  method = "DELETE",
  params = { id = "example" },
})
if err then error(err) end

print(fetchdef["url"])
print(fetchdef["method"])
print(fetchdef["headers"])
```

### Use test mode

Create a mock client for unit testing — no server required:

```lua
local client = sdk.test()

local result, err = client:Author():list()
-- result is the returned data; err is set on failure
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```lua
local function mock_fetch(url, init)
  return {
    status = 200,
    statusText = "OK",
    headers = {},
    json = function()
      return { id = "mock01" }
    end,
  }, nil
end

local client = sdk.new({
  base = "http://localhost:8080",
  system = {
    fetch = mock_fetch,
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
cd lua && busted test/
```


## Reference

### PoetrydbSDK

```lua
local sdk = require("poetrydb_sdk")
local client = sdk.new(options)
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `base` | `string` | Base URL of the API server. |
| `prefix` | `string` | URL path prefix prepended to all requests. |
| `suffix` | `string` | URL path suffix appended to all requests. |
| `feature` | `table` | Feature activation flags. |
| `extend` | `table` | Additional Feature instances to load. |
| `system` | `table` | System overrides (e.g. custom `fetch` function). |

### test

```lua
local client = sdk.test(testopts, sdkopts)
```

Creates a test-mode client with mock transport. Both arguments may be `nil`.

### PoetrydbSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `() -> table` | Deep copy of current SDK options. |
| `get_utility` | `() -> Utility` | Copy of the SDK utility object. |
| `prepare` | `(fetchargs) -> table, err` | Build an HTTP request definition without sending. |
| `direct` | `(fetchargs) -> table, err` | Build and send an HTTP request. |
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
| `load` | `(reqmatch, ctrl) -> any, err` | Load a single entity by match criteria. |
| `list` | `(reqmatch, ctrl) -> any, err` | List entities matching the criteria. |
| `data_get` | `() -> table` | Get entity data. |
| `data_set` | `(data)` | Set entity data. |
| `match_get` | `() -> table` | Get entity match criteria. |
| `match_set` | `(match)` | Set entity match criteria. |
| `make` | `() -> Entity` | Create a new instance with the same options. |
| `get_name` | `() -> string` | Return the entity name. |

### Result shape

Entity operations return `(value, err)`. The `value` is the operation's
data **directly** — there is no wrapper:

| Operation | `value` |
| --- | --- |
| `load` | the entity record (a `table`) |
| `list` | an array (`table`) of entity records |

Check `err` first (it is non-`nil` on failure), then use `value`:

    local author, err = client:Author():load({ id = "example_id" })
    if err then error(err) end
    -- author is the loaded record

Only `direct()` returns a response envelope — a `table` with `ok`,
`status`, `headers`, and `data` keys.

### Entities

#### Author

| Field | Description |
| --- | --- |
| `author` |  |
| `line` |  |
| `linecount` |  |
| `title` |  |

Operations: List, Load.

API path: `/author/{author}/{outputFields}.{format}`

#### Authorab

| Field | Description |
| --- | --- |
| `author` |  |
| `line` |  |
| `linecount` |  |
| `title` |  |

Operations: List.

API path: `/author/{author}:abs`

#### CombinedSearch

| Field | Description |
| --- | --- |
| `author` |  |
| `line` |  |
| `linecount` |  |
| `title` |  |

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
| `author` |  |
| `line` |  |
| `linecount` |  |
| `title` |  |

Operations: List, Load.

API path: `/lines/{lines}/{outputFields}.{format}`

#### Linecount

| Field | Description |
| --- | --- |
| `author` |  |
| `line` |  |
| `linecount` |  |
| `title` |  |

Operations: List, Load.

API path: `/linecount/{linecount}/{outputFields}.{format}`

#### Poemcount

| Field | Description |
| --- | --- |
| `author` |  |
| `line` |  |
| `linecount` |  |
| `title` |  |

Operations: Load.

API path: `/poemcount/{count}`

#### Random

| Field | Description |
| --- | --- |
| `author` |  |
| `line` |  |
| `linecount` |  |
| `title` |  |

Operations: List, Load.

API path: `/random/{count}/{outputFields}`

#### Title

| Field | Description |
| --- | --- |
| `author` |  |
| `line` |  |
| `linecount` |  |
| `title` |  |

Operations: List, Load.

API path: `/title/{title}/{outputFields}.{format}`

#### Titleab

| Field | Description |
| --- | --- |
| `author` |  |
| `line` |  |
| `linecount` |  |
| `title` |  |

Operations: List.

API path: `/title/{title}:abs`



## Entities


### Author

Create an instance: `local author = client:Author(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `author` | `string` |  |
| `line` | `table` |  |
| `linecount` | `number` |  |
| `title` | `string` |  |

#### Example: Load

```lua
local author, err = client:Author():load({ id = "author_id" })
```

#### Example: List

```lua
local authors, err = client:Author():list()
```


### Authorab

Create an instance: `local authorab = client:Authorab(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `author` | `string` |  |
| `line` | `table` |  |
| `linecount` | `number` |  |
| `title` | `string` |  |

#### Example: List

```lua
local authorabs, err = client:Authorab():list()
```


### CombinedSearch

Create an instance: `local combined_search = client:CombinedSearch(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `author` | `string` |  |
| `line` | `table` |  |
| `linecount` | `number` |  |
| `title` | `string` |  |

#### Example: List

```lua
local combined_searchs, err = client:CombinedSearch():list()
```


### CombinedSearchWithField

Create an instance: `local combined_search_with_field = client:CombinedSearchWithField(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Example: List

```lua
local combined_search_with_fields, err = client:CombinedSearchWithField():list()
```


### Line

Create an instance: `local line = client:Line(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `author` | `string` |  |
| `line` | `table` |  |
| `linecount` | `number` |  |
| `title` | `string` |  |

#### Example: Load

```lua
local line, err = client:Line():load({ id = "line_id" })
```

#### Example: List

```lua
local lines, err = client:Line():list()
```


### Linecount

Create an instance: `local linecount = client:Linecount(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `author` | `string` |  |
| `line` | `table` |  |
| `linecount` | `number` |  |
| `title` | `string` |  |

#### Example: Load

```lua
local linecount, err = client:Linecount():load({ id = "linecount_id" })
```

#### Example: List

```lua
local linecounts, err = client:Linecount():list()
```


### Poemcount

Create an instance: `local poemcount = client:Poemcount(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `author` | `string` |  |
| `line` | `table` |  |
| `linecount` | `number` |  |
| `title` | `string` |  |

#### Example: Load

```lua
local poemcount, err = client:Poemcount():load({ id = "poemcount_id" })
```


### Random

Create an instance: `local random = client:Random(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `author` | `string` |  |
| `line` | `table` |  |
| `linecount` | `number` |  |
| `title` | `string` |  |

#### Example: Load

```lua
local random, err = client:Random():load({ id = "random_id" })
```

#### Example: List

```lua
local randoms, err = client:Random():list()
```


### Title

Create an instance: `local title = client:Title(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `author` | `string` |  |
| `line` | `table` |  |
| `linecount` | `number` |  |
| `title` | `string` |  |

#### Example: Load

```lua
local title, err = client:Title():load({ id = "title_id" })
```

#### Example: List

```lua
local titles, err = client:Title():list()
```


### Titleab

Create an instance: `local titleab = client:Titleab(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `author` | `string` |  |
| `line` | `table` |  |
| `linecount` | `number` |  |
| `title` | `string` |  |

#### Example: List

```lua
local titleabs, err = client:Titleab():list()
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

Features are the extension mechanism. A feature is a Lua table
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as tables

The Lua SDK uses plain Lua tables throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `helpers.to_map()` to safely validate that a value is a table.

### Module structure

```
lua/
├── poetrydb_sdk.lua    -- Main SDK module
├── config.lua               -- Configuration
├── features.lua             -- Feature factory
├── core/                    -- Core types and context
├── entity/                  -- Entity implementations
├── feature/                 -- Built-in features (Base, Test, Log)
├── utility/                 -- Utility functions and struct library
└── test/                    -- Test suites
```

The main module (`poetrydb_sdk`) exports the SDK constructor
and test helper. Import entity or utility modules directly only
when needed.

### Entity state

Entity instances are stateful. After a successful `list`, the entity
stores the returned data and match criteria internally.

```lua
local author = client:Author()
author:list()

-- author:data_get() now returns the author data from the last list
-- author:match_get() returns the last match criteria
```

Call `make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

`direct()` gives full control over the HTTP request. Use it for
non-standard endpoints, bulk operations, or any path not modelled as
an entity. `prepare()` builds the request without sending it — useful
for debugging or custom transport.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
