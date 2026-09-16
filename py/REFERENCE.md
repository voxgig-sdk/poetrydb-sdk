# Poetrydb Python SDK Reference

Complete API reference for the Poetrydb Python SDK.


## PoetrydbSDK

### Constructor

```python
from poetrydb_sdk import PoetrydbSDK

client = PoetrydbSDK(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `dict` | SDK configuration options. |
| `options["base"]` | `str` | Base URL for API requests. |
| `options["prefix"]` | `str` | URL prefix appended after base. |
| `options["suffix"]` | `str` | URL suffix appended after path. |
| `options["headers"]` | `dict` | Custom headers for all requests. |
| `options["feature"]` | `dict` | Feature configuration. |
| `options["system"]` | `dict` | System overrides (e.g. custom fetch). |


### Static Methods

#### `PoetrydbSDK.test(testopts=None, sdkopts=None)`

Create a test client with mock features active. Both arguments may be `None`.

```python
client = PoetrydbSDK.test()
```


### Instance Methods

#### `Author(data=None)`

Create a new `AuthorEntity` instance. Pass `None` for no initial data.

#### `Authorab(data=None)`

Create a new `AuthorabEntity` instance. Pass `None` for no initial data.

#### `CombinedSearch(data=None)`

Create a new `CombinedSearchEntity` instance. Pass `None` for no initial data.

#### `CombinedSearchWithField(data=None)`

Create a new `CombinedSearchWithFieldEntity` instance. Pass `None` for no initial data.

#### `Line(data=None)`

Create a new `LineEntity` instance. Pass `None` for no initial data.

#### `Linecount(data=None)`

Create a new `LinecountEntity` instance. Pass `None` for no initial data.

#### `Poemcount(data=None)`

Create a new `PoemcountEntity` instance. Pass `None` for no initial data.

#### `Random(data=None)`

Create a new `RandomEntity` instance. Pass `None` for no initial data.

#### `Title(data=None)`

Create a new `TitleEntity` instance. Pass `None` for no initial data.

#### `Titleab(data=None)`

Create a new `TitleabEntity` instance. Pass `None` for no initial data.

#### `options_map() -> dict`

Return a deep copy of the current SDK options.

#### `get_utility() -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs=None) -> dict`

Make a direct HTTP request to any API endpoint. Returns a result `dict` with `ok`, `status`, `headers`, and `data` (or `err` on failure). This escape hatch never raises — branch on `result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `str` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `str` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `dict` | Path parameter values. |
| `fetchargs["query"]` | `dict` | Query string parameters. |
| `fetchargs["headers"]` | `dict` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (dicts are JSON-serialized). |

**Returns:** `result_dict`

#### `prepare(fetchargs=None) -> dict`

Prepare a fetch definition without sending. Returns the `fetchdef` and raises on error.


---

## AuthorEntity

```python
author = client.Author()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `author` | `str` | No | The author of the poem |
| `authors` | `list` | No |  |
| `id` | `str` | No |  |
| `linecount` | `int` | No | The number of lines in the poem (including section headings, excluding empty lines) |
| `lines` | `list` | No | The lines of the poem |
| `title` | `str` | No | The title of the poem |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Author().list({"author": "example", "format": "example", "output_field": "example"})
for author in results:
    print(author)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Author().load({"id": "author_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `AuthorEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## AuthorabEntity

```python
authorab = client.Authorab()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `author` | `str` | No | The author of the poem |
| `linecount` | `int` | No | The number of lines in the poem (including section headings, excluding empty lines) |
| `lines` | `list` | No | The lines of the poem |
| `title` | `str` | No | The title of the poem |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Authorab().list({"author": "example"})
for authorab in results:
    print(authorab)
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `AuthorabEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## CombinedSearchEntity

```python
combined_search = client.CombinedSearch()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `author` | `str` | No | The author of the poem |
| `linecount` | `int` | No | The number of lines in the poem (including section headings, excluding empty lines) |
| `lines` | `list` | No | The lines of the poem |
| `title` | `str` | No | The title of the poem |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.CombinedSearch().list({"input_field1": "example", "input_field2": "example", "search_term1": "example", "search_term2": "example"})
for combined_search in results:
    print(combined_search)
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `CombinedSearchEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## CombinedSearchWithFieldEntity

```python
combined_search_with_field = client.CombinedSearchWithField()
```

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.CombinedSearchWithField().list({"input_field1": "example", "input_field2": "example", "output_field": "example", "search_term1": "example", "search_term2": "example"})
for combined_search_with_field in results:
    print(combined_search_with_field)
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `CombinedSearchWithFieldEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## LineEntity

```python
line = client.Line()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `author` | `str` | No | The author of the poem |
| `id` | `str` | No |  |
| `linecount` | `int` | No | The number of lines in the poem (including section headings, excluding empty lines) |
| `lines` | `list` | No | The lines of the poem |
| `title` | `str` | No | The title of the poem |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Line().list({"line": "example", "output_field": "example"})
for line in results:
    print(line)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Line().load({"id": "line_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `LineEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## LinecountEntity

```python
linecount = client.Linecount()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `author` | `str` | No | The author of the poem |
| `id` | `str` | No |  |
| `linecount` | `int` | No | The number of lines in the poem (including section headings, excluding empty lines) |
| `lines` | `list` | No | The lines of the poem |
| `title` | `str` | No | The title of the poem |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Linecount().list({"linecount": 1, "output_field": "example"})
for linecount in results:
    print(linecount)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Linecount().load({"id": 1})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `LinecountEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## PoemcountEntity

```python
poemcount = client.Poemcount()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `author` | `str` | No | The author of the poem |
| `id` | `str` | No |  |
| `linecount` | `int` | No | The number of lines in the poem (including section headings, excluding empty lines) |
| `lines` | `list` | No | The lines of the poem |
| `title` | `str` | No | The title of the poem |

### Operations

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Poemcount().load({"id": 1})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PoemcountEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## RandomEntity

```python
random = client.Random()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `author` | `str` | No | The author of the poem |
| `id` | `str` | No |  |
| `linecount` | `int` | No | The number of lines in the poem (including section headings, excluding empty lines) |
| `lines` | `list` | No | The lines of the poem |
| `title` | `str` | No | The title of the poem |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Random().list({"count": 1, "output_field": "example"})
for random in results:
    print(random)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Random().load({"id": 1})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `RandomEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## TitleEntity

```python
title = client.Title()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `author` | `str` | No | The author of the poem |
| `id` | `str` | No |  |
| `linecount` | `int` | No | The number of lines in the poem (including section headings, excluding empty lines) |
| `lines` | `list` | No | The lines of the poem |
| `title` | `str` | No | The title of the poem |
| `titles` | `list` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Title().list({"format": "example", "output_field": "example", "title": "example"})
for title in results:
    print(title)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Title().load({"id": "title_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `TitleEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## TitleabEntity

```python
titleab = client.Titleab()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `author` | `str` | No | The author of the poem |
| `linecount` | `int` | No | The number of lines in the poem (including section headings, excluding empty lines) |
| `lines` | `list` | No | The lines of the poem |
| `title` | `str` | No | The title of the poem |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Titleab().list({"title": "example"})
for titleab in results:
    print(titleab)
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `TitleabEntity` instance with the same options.

#### `get_name() -> str`

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

```python
client = PoetrydbSDK({
    "feature": {
        "ratelimit": {"active": True},
        "retry": {"active": True},
        "test": {"active": True},
        "timeout": {"active": True},
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

