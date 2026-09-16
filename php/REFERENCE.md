# Poetrydb PHP SDK Reference

Complete API reference for the Poetrydb PHP SDK.


## PoetrydbSDK

### Constructor

```php
require_once __DIR__ . '/poetrydb_sdk.php';

$client = new PoetrydbSDK($options);
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `$options` | `array` | SDK configuration options. |
| `$options["base"]` | `string` | Base URL for API requests. |
| `$options["prefix"]` | `string` | URL prefix appended after base. |
| `$options["suffix"]` | `string` | URL suffix appended after path. |
| `$options["headers"]` | `array` | Custom headers for all requests. |
| `$options["feature"]` | `array` | Feature configuration. |
| `$options["system"]` | `array` | System overrides (e.g. custom fetch). |


### Static Methods

#### `PoetrydbSDK::test($testopts = null, $sdkopts = null)`

Create a test client with mock features active. Both arguments may be `null`.

```php
$client = PoetrydbSDK::test();
```


### Instance Methods

#### `Author($data = null)`

Create a new `AuthorEntity` instance. Pass `null` for no initial data.

#### `Authorab($data = null)`

Create a new `AuthorabEntity` instance. Pass `null` for no initial data.

#### `CombinedSearch($data = null)`

Create a new `CombinedSearchEntity` instance. Pass `null` for no initial data.

#### `CombinedSearchWithField($data = null)`

Create a new `CombinedSearchWithFieldEntity` instance. Pass `null` for no initial data.

#### `Line($data = null)`

Create a new `LineEntity` instance. Pass `null` for no initial data.

#### `Linecount($data = null)`

Create a new `LinecountEntity` instance. Pass `null` for no initial data.

#### `Poemcount($data = null)`

Create a new `PoemcountEntity` instance. Pass `null` for no initial data.

#### `Random($data = null)`

Create a new `RandomEntity` instance. Pass `null` for no initial data.

#### `Title($data = null)`

Create a new `TitleEntity` instance. Pass `null` for no initial data.

#### `Titleab($data = null)`

Create a new `TitleabEntity` instance. Pass `null` for no initial data.

#### `options_map(): array`

Return a deep copy of the current SDK options.

#### `get_utility(): PoetrydbUtility`

Return a copy of the SDK utility object.

#### `direct(array $fetchargs = []): array`

Make a direct HTTP request to any API endpoint. This is the raw-HTTP escape
hatch: it does **not** throw. It returns a result array
`["ok" => bool, "status" => int, "headers" => array, "data" => mixed]`, or
`["ok" => false, "err" => \Exception]` on failure. Branch on `$result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `$fetchargs["path"]` | `string` | URL path with optional `{param}` placeholders. |
| `$fetchargs["method"]` | `string` | HTTP method (default: `"GET"`). |
| `$fetchargs["params"]` | `array` | Path parameter values for `{param}` substitution. |
| `$fetchargs["query"]` | `array` | Query string parameters. |
| `$fetchargs["headers"]` | `array` | Request headers (merged with defaults). |
| `$fetchargs["body"]` | `mixed` | Request body (arrays are JSON-serialized). |
| `$fetchargs["ctrl"]` | `array` | Control options. |

**Returns:** `array` — the result dict (see above); never throws.

#### `prepare(array $fetchargs = []): mixed`

Prepare a fetch definition without sending the request. Returns the
`$fetchdef` array. Throws on error.


---

## AuthorEntity

```php
$author = $client->Author();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `author` | `string` | No | The author of the poem |
| `authors` | `array` | No |  |
| `id` | `string` | No |  |
| `linecount` | `int` | No | The number of lines in the poem (including section headings, excluding empty lines) |
| `lines` | `array` | No | The lines of the poem |
| `title` | `string` | No | The title of the poem |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Author()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Author()->load(["id" => "author_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): AuthorEntity`

Create a new `AuthorEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## AuthorabEntity

```php
$authorab = $client->Authorab();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `author` | `string` | No | The author of the poem |
| `linecount` | `int` | No | The number of lines in the poem (including section headings, excluding empty lines) |
| `lines` | `array` | No | The lines of the poem |
| `title` | `string` | No | The title of the poem |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Authorab()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): AuthorabEntity`

Create a new `AuthorabEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## CombinedSearchEntity

```php
$combined_search = $client->CombinedSearch();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `author` | `string` | No | The author of the poem |
| `linecount` | `int` | No | The number of lines in the poem (including section headings, excluding empty lines) |
| `lines` | `array` | No | The lines of the poem |
| `title` | `string` | No | The title of the poem |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->CombinedSearch()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): CombinedSearchEntity`

Create a new `CombinedSearchEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## CombinedSearchWithFieldEntity

```php
$combined_search_with_field = $client->CombinedSearchWithField();
```

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->CombinedSearchWithField()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): CombinedSearchWithFieldEntity`

Create a new `CombinedSearchWithFieldEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## LineEntity

```php
$line = $client->Line();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `author` | `string` | No | The author of the poem |
| `id` | `string` | No |  |
| `linecount` | `int` | No | The number of lines in the poem (including section headings, excluding empty lines) |
| `lines` | `array` | No | The lines of the poem |
| `title` | `string` | No | The title of the poem |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Line()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Line()->load(["id" => "line_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): LineEntity`

Create a new `LineEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## LinecountEntity

```php
$linecount = $client->Linecount();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `author` | `string` | No | The author of the poem |
| `id` | `string` | No |  |
| `linecount` | `int` | No | The number of lines in the poem (including section headings, excluding empty lines) |
| `lines` | `array` | No | The lines of the poem |
| `title` | `string` | No | The title of the poem |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Linecount()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Linecount()->load(["id" => 1]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): LinecountEntity`

Create a new `LinecountEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## PoemcountEntity

```php
$poemcount = $client->Poemcount();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `author` | `string` | No | The author of the poem |
| `id` | `string` | No |  |
| `linecount` | `int` | No | The number of lines in the poem (including section headings, excluding empty lines) |
| `lines` | `array` | No | The lines of the poem |
| `title` | `string` | No | The title of the poem |

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Poemcount()->load(["id" => 1]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): PoemcountEntity`

Create a new `PoemcountEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## RandomEntity

```php
$random = $client->Random();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `author` | `string` | No | The author of the poem |
| `id` | `string` | No |  |
| `linecount` | `int` | No | The number of lines in the poem (including section headings, excluding empty lines) |
| `lines` | `array` | No | The lines of the poem |
| `title` | `string` | No | The title of the poem |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Random()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Random()->load(["id" => 1]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): RandomEntity`

Create a new `RandomEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## TitleEntity

```php
$title = $client->Title();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `author` | `string` | No | The author of the poem |
| `id` | `string` | No |  |
| `linecount` | `int` | No | The number of lines in the poem (including section headings, excluding empty lines) |
| `lines` | `array` | No | The lines of the poem |
| `title` | `string` | No | The title of the poem |
| `titles` | `array` | No |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Title()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Title()->load(["id" => "title_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): TitleEntity`

Create a new `TitleEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## TitleabEntity

```php
$titleab = $client->Titleab();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `author` | `string` | No | The author of the poem |
| `linecount` | `int` | No | The number of lines in the poem (including section headings, excluding empty lines) |
| `lines` | `array` | No | The lines of the poem |
| `title` | `string` | No | The title of the poem |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Titleab()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): TitleabEntity`

Create a new `TitleabEntity` instance with the same client and
options.

#### `get_name(): string`

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

```php
$client = new PoetrydbSDK([
  "feature" => [
    "ratelimit" => ["active" => true],
    "retry" => ["active" => true],
    "test" => ["active" => true],
    "timeout" => ["active" => true],
  ],
]);
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

