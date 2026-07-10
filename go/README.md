# Poetrydb Golang SDK



The Golang SDK for the Poetrydb API — an entity-oriented client using standard Go conventions. No generics required; data flows as `map[string]any`.

It exposes the API as capitalised, semantic **Entities** — e.g. `client.Author(nil)` — each with the same small set of operations (`List`, `Load`) instead of raw URL paths and query strings. You call meaning, not endpoints, which keeps the cognitive load low.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
```bash
go get github.com/voxgig-sdk/poetrydb-sdk/go@latest
```

The Go module proxy resolves the version from the `go/vX.Y.Z` GitHub
release tag — see [Releases](https://github.com/voxgig-sdk/poetrydb-sdk/releases) for the available versions.

To vendor from a local checkout instead, clone this repo alongside your
project and add a `replace` directive pointing at the checked-out
`go/` directory:

```bash
go mod edit -replace github.com/voxgig-sdk/poetrydb-sdk/go=../poetrydb-sdk/go
```


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### Quickstart

A complete program: create a client, then call the entity operations.
Each operation returns `(value, error)` — the value is the data itself
(there is no `{ok, data}` wrapper), so check `err` and use the value
directly.

```go
package main

import (
    "fmt"
    sdk "github.com/voxgig-sdk/poetrydb-sdk/go"
)

func main() {
    client := sdk.New()

    // List author records — the value is the array of records itself.
    authors, err := client.Author(nil).List(nil, nil)
    if err != nil {
        panic(err)
    }
    for _, item := range authors.([]any) {
        fmt.Println(item)
    }

    // Load a single author — the value is the loaded record.
    author, err := client.Author(nil).Load(map[string]any{"id": "example_id"}, nil)
    if err != nil {
        panic(err)
    }
    fmt.Println(author)
}
```


## Error handling

Every entity operation returns `(value, error)`. Check `err` before
using the value — there is no exception to catch:

```go
authors, err := client.Author(nil).List(nil, nil)
if err != nil {
    // handle err
    return
}
_ = authors
```

`Direct` follows the same `(value, error)` convention:

```go
result, err := client.Direct(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "GET",
    "params": map[string]any{"id": "example_id"},
})
if err != nil {
    // handle err
}
_ = result
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```go
result, err := client.Direct(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "GET",
    "params": map[string]any{"id": "example"},
})
if err != nil {
    panic(err)
}

if result["ok"] == true {
    fmt.Println(result["status"]) // 200
    fmt.Println(result["data"])   // response body
}
```

### Prepare a request without sending it

```go
fetchdef, err := client.Prepare(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "DELETE",
    "params": map[string]any{"id": "example"},
})
if err != nil {
    panic(err)
}

fmt.Println(fetchdef["url"])
fmt.Println(fetchdef["method"])
fmt.Println(fetchdef["headers"])
```

### Use test mode

Create a mock client for unit testing — no server required:

```go
client := sdk.Test()

author, err := client.Author(nil).List(
    nil, nil,
)
if err != nil {
    panic(err)
}
fmt.Println(author) // the returned mock data
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```go
mockFetch := func(url string, init map[string]any) (map[string]any, error) {
    return map[string]any{
        "status":     200,
        "statusText": "OK",
        "headers":    map[string]any{},
        "json": (func() any)(func() any {
            return map[string]any{"id": "mock01"}
        }),
    }, nil
}

client := sdk.NewPoetrydbSDK(map[string]any{
    "base": "http://localhost:8080",
    "system": map[string]any{
        "fetch": (func(string, map[string]any) (map[string]any, error))(mockFetch),
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
cd go && go test ./test/...
```


## Reference

### NewPoetrydbSDK

```go
func NewPoetrydbSDK(options map[string]any) *PoetrydbSDK
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `"base"` | `string` | Base URL of the API server. |
| `"prefix"` | `string` | URL path prefix prepended to all requests. |
| `"suffix"` | `string` | URL path suffix appended to all requests. |
| `"feature"` | `map[string]any` | Feature activation flags. |
| `"extend"` | `[]any` | Additional Feature instances to load. |
| `"system"` | `map[string]any` | System overrides (e.g. custom `"fetch"` function). |

### TestSDK

```go
func TestSDK(testopts map[string]any, sdkopts map[string]any) *PoetrydbSDK
```

Creates a test-mode client with mock transport. Both arguments may be `nil`.

### PoetrydbSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `OptionsMap` | `() map[string]any` | Deep copy of current SDK options. |
| `GetUtility` | `() *Utility` | Copy of the SDK utility object. |
| `Prepare` | `(fetchargs map[string]any) (map[string]any, error)` | Build an HTTP request definition without sending. |
| `Direct` | `(fetchargs map[string]any) (map[string]any, error)` | Build and send an HTTP request. |
| `Author` | `(data map[string]any) PoetrydbEntity` | Create an Author entity instance. |
| `Authorab` | `(data map[string]any) PoetrydbEntity` | Create an Authorab entity instance. |
| `CombinedSearch` | `(data map[string]any) PoetrydbEntity` | Create a CombinedSearch entity instance. |
| `CombinedSearchWithField` | `(data map[string]any) PoetrydbEntity` | Create a CombinedSearchWithField entity instance. |
| `Line` | `(data map[string]any) PoetrydbEntity` | Create a Line entity instance. |
| `Linecount` | `(data map[string]any) PoetrydbEntity` | Create a Linecount entity instance. |
| `Poemcount` | `(data map[string]any) PoetrydbEntity` | Create a Poemcount entity instance. |
| `Random` | `(data map[string]any) PoetrydbEntity` | Create a Random entity instance. |
| `Title` | `(data map[string]any) PoetrydbEntity` | Create a Title entity instance. |
| `Titleab` | `(data map[string]any) PoetrydbEntity` | Create a Titleab entity instance. |

### Entity interface (PoetrydbEntity)

All entities implement the `PoetrydbEntity` interface.

| Method | Signature | Description |
| --- | --- | --- |
| `Load` | `(reqmatch, ctrl map[string]any) (any, error)` | Load a single entity by match criteria. |
| `List` | `(reqmatch, ctrl map[string]any) (any, error)` | List entities matching the criteria. |
| `Data` | `(args ...any) any` | Get or set entity data. |
| `Match` | `(args ...any) any` | Get or set entity match criteria. |
| `Make` | `() Entity` | Create a new instance with the same options. |
| `GetName` | `() string` | Return the entity name. |

### Result shape

Entity operations return `(value, error)`. The `value` is the
operation's data **directly** — there is no wrapper:

| Operation | `value` |
| --- | --- |
| `Load` | the entity record (`map[string]any`) |
| `List` | a `[]any` of entity records |

Check `err` first, then use the value directly (or the typed
`...Typed` variants, which return the entity's model struct and a typed
slice):

    author, err := client.Author(nil).List(map[string]any{/* fields */}, nil)
    if err != nil { /* handle */ }
    // author is the returned record

Only `Direct()` returns a response envelope — a `map[string]any` with
`"ok"`, `"status"`, `"headers"`, and `"data"` keys.

### Entities

#### Author

| Field | Description |
| --- | --- |
| `"author"` |  |
| `"line"` |  |
| `"linecount"` |  |
| `"title"` |  |

Operations: List, Load.

API path: `/author/{author}/{outputFields}.{format}`

#### Authorab

| Field | Description |
| --- | --- |
| `"author"` |  |
| `"line"` |  |
| `"linecount"` |  |
| `"title"` |  |

Operations: List.

API path: `/author/{author}:abs`

#### CombinedSearch

| Field | Description |
| --- | --- |
| `"author"` |  |
| `"line"` |  |
| `"linecount"` |  |
| `"title"` |  |

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
| `"author"` |  |
| `"line"` |  |
| `"linecount"` |  |
| `"title"` |  |

Operations: List, Load.

API path: `/lines/{lines}/{outputFields}.{format}`

#### Linecount

| Field | Description |
| --- | --- |
| `"author"` |  |
| `"line"` |  |
| `"linecount"` |  |
| `"title"` |  |

Operations: List, Load.

API path: `/linecount/{linecount}/{outputFields}.{format}`

#### Poemcount

| Field | Description |
| --- | --- |
| `"author"` |  |
| `"line"` |  |
| `"linecount"` |  |
| `"title"` |  |

Operations: Load.

API path: `/poemcount/{count}`

#### Random

| Field | Description |
| --- | --- |
| `"author"` |  |
| `"line"` |  |
| `"linecount"` |  |
| `"title"` |  |

Operations: List, Load.

API path: `/random/{count}/{outputFields}`

#### Title

| Field | Description |
| --- | --- |
| `"author"` |  |
| `"line"` |  |
| `"linecount"` |  |
| `"title"` |  |

Operations: List, Load.

API path: `/title/{title}/{outputFields}.{format}`

#### Titleab

| Field | Description |
| --- | --- |
| `"author"` |  |
| `"line"` |  |
| `"linecount"` |  |
| `"title"` |  |

Operations: List.

API path: `/title/{title}:abs`



## Entities


### Author

Create an instance: `author := client.Author(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `author` | `string` |  |
| `line` | `[]any` |  |
| `linecount` | `int` |  |
| `title` | `string` |  |

#### Example: Load

```go
author, err := client.Author(nil).Load(map[string]any{"id": "author_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(author) // the loaded record
```

#### Example: List

```go
authors, err := client.Author(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(authors) // the array of records
```


### Authorab

Create an instance: `authorab := client.Authorab(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `author` | `string` |  |
| `line` | `[]any` |  |
| `linecount` | `int` |  |
| `title` | `string` |  |

#### Example: List

```go
authorabs, err := client.Authorab(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(authorabs) // the array of records
```


### CombinedSearch

Create an instance: `combinedSearch := client.CombinedSearch(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `author` | `string` |  |
| `line` | `[]any` |  |
| `linecount` | `int` |  |
| `title` | `string` |  |

#### Example: List

```go
combinedSearchs, err := client.CombinedSearch(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(combinedSearchs) // the array of records
```


### CombinedSearchWithField

Create an instance: `combinedSearchWithField := client.CombinedSearchWithField(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Example: List

```go
combinedSearchWithFields, err := client.CombinedSearchWithField(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(combinedSearchWithFields) // the array of records
```


### Line

Create an instance: `line := client.Line(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `author` | `string` |  |
| `line` | `[]any` |  |
| `linecount` | `int` |  |
| `title` | `string` |  |

#### Example: Load

```go
line, err := client.Line(nil).Load(map[string]any{"id": "line_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(line) // the loaded record
```

#### Example: List

```go
lines, err := client.Line(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(lines) // the array of records
```


### Linecount

Create an instance: `linecount := client.Linecount(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `author` | `string` |  |
| `line` | `[]any` |  |
| `linecount` | `int` |  |
| `title` | `string` |  |

#### Example: Load

```go
linecount, err := client.Linecount(nil).Load(map[string]any{"id": 1}, nil)
if err != nil {
    panic(err)
}
fmt.Println(linecount) // the loaded record
```

#### Example: List

```go
linecounts, err := client.Linecount(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(linecounts) // the array of records
```


### Poemcount

Create an instance: `poemcount := client.Poemcount(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `author` | `string` |  |
| `line` | `[]any` |  |
| `linecount` | `int` |  |
| `title` | `string` |  |

#### Example: Load

```go
poemcount, err := client.Poemcount(nil).Load(map[string]any{"id": 1}, nil)
if err != nil {
    panic(err)
}
fmt.Println(poemcount) // the loaded record
```


### Random

Create an instance: `random := client.Random(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `author` | `string` |  |
| `line` | `[]any` |  |
| `linecount` | `int` |  |
| `title` | `string` |  |

#### Example: Load

```go
random, err := client.Random(nil).Load(map[string]any{"id": 1}, nil)
if err != nil {
    panic(err)
}
fmt.Println(random) // the loaded record
```

#### Example: List

```go
randoms, err := client.Random(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(randoms) // the array of records
```


### Title

Create an instance: `title := client.Title(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `author` | `string` |  |
| `line` | `[]any` |  |
| `linecount` | `int` |  |
| `title` | `string` |  |

#### Example: Load

```go
title, err := client.Title(nil).Load(map[string]any{"id": "title_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(title) // the loaded record
```

#### Example: List

```go
titles, err := client.Title(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(titles) // the array of records
```


### Titleab

Create an instance: `titleab := client.Titleab(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `author` | `string` |  |
| `line` | `[]any` |  |
| `linecount` | `int` |  |
| `title` | `string` |  |

#### Example: List

```go
titleabs, err := client.Titleab(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(titleabs) // the array of records
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

Features are the extension mechanism. A feature implements the
`Feature` interface and provides hooks — functions keyed by pipeline
stage names.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as maps

The Go SDK uses `map[string]any` throughout rather than typed structs.
This mirrors the dynamic nature of the API and keeps the SDK
flexible — no code generation is needed when the API schema changes.

Use `core.ToMapAny()` to safely cast results and nested data.

### Package structure

```
github.com/voxgig-sdk/poetrydb-sdk/go/
├── poetrydb.go        # Root package — type aliases and constructors
├── core/               # SDK core — client, types, pipeline
├── entity/             # Entity implementations
├── feature/            # Built-in features (Base, Test, Log)
├── utility/            # Utility functions and struct library
└── test/               # Test suites
```

The root package (`github.com/voxgig-sdk/poetrydb-sdk/go`) re-exports everything needed
for normal use. Import sub-packages only when you need specific types
like `core.ToMapAny`.

### Entity state

Entity instances are stateful. After a successful `List`, the entity
stores the returned data and match criteria internally.

```go
author := client.Author(nil)
author.List(nil, nil)

// author.Data() now returns the author data from the last list
// author.Match() returns the last match criteria
```

Call `Make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

`Direct()` gives full control over the HTTP request. Use it for
non-standard endpoints, bulk operations, or any path not modelled as
an entity. `Prepare()` builds the request without sending it — useful
for debugging or custom transport.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
