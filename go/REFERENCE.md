# Poetrydb Golang SDK Reference

Complete API reference for the Poetrydb Golang SDK.


## PoetrydbSDK

### Constructor

```go
func NewPoetrydbSDK(options map[string]any) *PoetrydbSDK
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `map[string]any` | SDK configuration options. |
| `options["base"]` | `string` | Base URL for API requests. |
| `options["prefix"]` | `string` | URL prefix appended after base. |
| `options["suffix"]` | `string` | URL suffix appended after path. |
| `options["headers"]` | `map[string]any` | Custom headers for all requests. |
| `options["feature"]` | `map[string]any` | Feature configuration. |
| `options["system"]` | `map[string]any` | System overrides (e.g. custom fetch). |


### Static Methods

#### `Test() *PoetrydbSDK`

No-arg convenience constructor for the common no-options test case.

```go
client := sdk.Test()
```

#### `TestSDK(testopts, sdkopts map[string]any) *PoetrydbSDK`

Test client with options. Both arguments may be `nil`.

```go
client := sdk.TestSDK(testopts, sdkopts)
```


### Instance Methods

#### `Author(data map[string]any) PoetrydbEntity`

Create a new `Author` entity instance. Pass `nil` for no initial data.

#### `Authorab(data map[string]any) PoetrydbEntity`

Create a new `Authorab` entity instance. Pass `nil` for no initial data.

#### `CombinedSearch(data map[string]any) PoetrydbEntity`

Create a new `CombinedSearch` entity instance. Pass `nil` for no initial data.

#### `CombinedSearchWithField(data map[string]any) PoetrydbEntity`

Create a new `CombinedSearchWithField` entity instance. Pass `nil` for no initial data.

#### `Line(data map[string]any) PoetrydbEntity`

Create a new `Line` entity instance. Pass `nil` for no initial data.

#### `Linecount(data map[string]any) PoetrydbEntity`

Create a new `Linecount` entity instance. Pass `nil` for no initial data.

#### `Poemcount(data map[string]any) PoetrydbEntity`

Create a new `Poemcount` entity instance. Pass `nil` for no initial data.

#### `Random(data map[string]any) PoetrydbEntity`

Create a new `Random` entity instance. Pass `nil` for no initial data.

#### `Title(data map[string]any) PoetrydbEntity`

Create a new `Title` entity instance. Pass `nil` for no initial data.

#### `Titleab(data map[string]any) PoetrydbEntity`

Create a new `Titleab` entity instance. Pass `nil` for no initial data.

#### `OptionsMap() map[string]any`

Return a deep copy of the current SDK options.

#### `GetUtility() *Utility`

Return a copy of the SDK utility object.

#### `Direct(fetchargs map[string]any) (map[string]any, error)`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `string` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `map[string]any` | Path parameter values for `{param}` substitution. |
| `fetchargs["query"]` | `map[string]any` | Query string parameters. |
| `fetchargs["headers"]` | `map[string]any` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (maps are JSON-serialized). |
| `fetchargs["ctrl"]` | `map[string]any` | Control options (e.g. `map[string]any{"explain": true}`). |

**Returns:** `(map[string]any, error)`

#### `Prepare(fetchargs map[string]any) (map[string]any, error)`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `Direct()`.

**Returns:** `(map[string]any, error)`


---

## AuthorEntity

```go
author := client.Author(nil)
fmt.Println(author.GetName()) // "author"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `author` | `string` | No | The author of the poem |
| `authors` | `[]any` | No |  |
| `id` | `string` | No |  |
| `linecount` | `int` | No | The number of lines in the poem (including section headings, excluding empty lines) |
| `lines` | `[]any` | No | The lines of the poem |
| `title` | `string` | No | The title of the poem |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Author(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Author(nil).Load(map[string]any{"id": "author_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `AuthorEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## AuthorabEntity

```go
authorab := client.Authorab(nil)
fmt.Println(authorab.GetName()) // "authorab"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `author` | `string` | No | The author of the poem |
| `linecount` | `int` | No | The number of lines in the poem (including section headings, excluding empty lines) |
| `lines` | `[]any` | No | The lines of the poem |
| `title` | `string` | No | The title of the poem |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Authorab(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `AuthorabEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## CombinedSearchEntity

```go
combinedSearch := client.CombinedSearch(nil)
fmt.Println(combinedSearch.GetName()) // "combined_search"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `author` | `string` | No | The author of the poem |
| `linecount` | `int` | No | The number of lines in the poem (including section headings, excluding empty lines) |
| `lines` | `[]any` | No | The lines of the poem |
| `title` | `string` | No | The title of the poem |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.CombinedSearch(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `CombinedSearchEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## CombinedSearchWithFieldEntity

```go
combinedSearchWithField := client.CombinedSearchWithField(nil)
fmt.Println(combinedSearchWithField.GetName()) // "combined_search_with_field"
```

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.CombinedSearchWithField(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `CombinedSearchWithFieldEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## LineEntity

```go
line := client.Line(nil)
fmt.Println(line.GetName()) // "line"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `author` | `string` | No | The author of the poem |
| `id` | `string` | No |  |
| `linecount` | `int` | No | The number of lines in the poem (including section headings, excluding empty lines) |
| `lines` | `[]any` | No | The lines of the poem |
| `title` | `string` | No | The title of the poem |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Line(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Line(nil).Load(map[string]any{"id": "line_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `LineEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## LinecountEntity

```go
linecount := client.Linecount(nil)
fmt.Println(linecount.GetName()) // "linecount"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `author` | `string` | No | The author of the poem |
| `id` | `string` | No |  |
| `linecount` | `int` | No | The number of lines in the poem (including section headings, excluding empty lines) |
| `lines` | `[]any` | No | The lines of the poem |
| `title` | `string` | No | The title of the poem |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Linecount(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Linecount(nil).Load(map[string]any{"id": 1}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `LinecountEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## PoemcountEntity

```go
poemcount := client.Poemcount(nil)
fmt.Println(poemcount.GetName()) // "poemcount"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `author` | `string` | No | The author of the poem |
| `id` | `string` | No |  |
| `linecount` | `int` | No | The number of lines in the poem (including section headings, excluding empty lines) |
| `lines` | `[]any` | No | The lines of the poem |
| `title` | `string` | No | The title of the poem |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Poemcount(nil).Load(map[string]any{"id": 1}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `PoemcountEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## RandomEntity

```go
random := client.Random(nil)
fmt.Println(random.GetName()) // "random"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `author` | `string` | No | The author of the poem |
| `id` | `string` | No |  |
| `linecount` | `int` | No | The number of lines in the poem (including section headings, excluding empty lines) |
| `lines` | `[]any` | No | The lines of the poem |
| `title` | `string` | No | The title of the poem |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Random(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Random(nil).Load(map[string]any{"id": 1}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `RandomEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## TitleEntity

```go
title := client.Title(nil)
fmt.Println(title.GetName()) // "title"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `author` | `string` | No | The author of the poem |
| `id` | `string` | No |  |
| `linecount` | `int` | No | The number of lines in the poem (including section headings, excluding empty lines) |
| `lines` | `[]any` | No | The lines of the poem |
| `title` | `string` | No | The title of the poem |
| `titles` | `[]any` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Title(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Title(nil).Load(map[string]any{"id": "title_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `TitleEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## TitleabEntity

```go
titleab := client.Titleab(nil)
fmt.Println(titleab.GetName()) // "titleab"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `author` | `string` | No | The author of the poem |
| `linecount` | `int` | No | The number of lines in the poem (including section headings, excluding empty lines) |
| `lines` | `[]any` | No | The lines of the poem |
| `title` | `string` | No | The title of the poem |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Titleab(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `TitleabEntity` instance with the same client and
options.

#### `GetName() string`

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

```go
client := sdk.NewPoetrydbSDK(map[string]any{
    "feature": map[string]any{
        "ratelimit": map[string]any{"active": true},
        "retry": map[string]any{"active": true},
        "test": map[string]any{"active": true},
        "timeout": map[string]any{"active": true},
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

