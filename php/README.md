# Poetrydb PHP SDK



The PHP SDK for the Poetrydb API — an entity-oriented client using PHP conventions.

The SDK exposes the API as capitalised, semantic **Entities** — for example `$client->Author()` — with named operations (`list`/`load`) instead of raw URL paths and query strings. Working with resources and verbs keeps call sites self-describing and reduces cognitive load.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to Packagist. Install it from the
GitHub release tag (`php/vX.Y.Z`):

- Releases: [https://github.com/voxgig-sdk/poetrydb-sdk/releases](https://github.com/voxgig-sdk/poetrydb-sdk/releases)


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```php
<?php
require_once 'poetrydb_sdk.php';

$client = new PoetrydbSDK();
```

### 2. List author records

```php
try {
    // list() returns an array of Author records — iterate directly.
    $authors = $client->Author()->list();
    foreach ($authors as $item) {
        echo $item["author"] . "\n";
    }
} catch (\Throwable $err) {
    echo "Error: " . $err->getMessage();
}
```

### 3. Load an author

```php
try {
    // load() returns the bare Author record (throws on error).
    $author = $client->Author()->load(["id" => "example_id"]);
    print_r($author);
} catch (\Throwable $err) {
    echo "Error: " . $err->getMessage();
}
```


## Error handling

Entity operations throw a `\Throwable` on failure, so wrap them in
`try` / `catch`:

```php
try {
    $authors = $client->Author()->list();
} catch (\Throwable $err) {
    echo "Error: " . $err->getMessage();
}
```

`direct()` does **not** throw — it returns the result array. Branch on
`ok`; on failure `status` holds the HTTP status (for error responses) and
`err` holds a transport error, so read both defensively:

```php
$result = $client->direct([
    "path" => "/api/resource/{id}",
    "method" => "GET",
    "params" => ["id" => "example_id"],
]);

if (! $result["ok"]) {
    $err = $result["err"] ?? null;
    echo "request failed: " . ($err ? $err->getMessage() : "HTTP " . $result["status"]);
}
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```php
// direct() is the raw-HTTP escape hatch: it returns a result array
// (it does not throw). Branch on $result["ok"].
$result = $client->direct([
    "path" => "/api/resource/{id}",
    "method" => "GET",
    "params" => ["id" => "example"],
]);

if ($result["ok"]) {
    echo $result["status"];  // 200
    print_r($result["data"]);  // response body
} else {
    // On an HTTP error status there is no err (only a transport failure sets
    // it), so fall back to the status code.
    $err = $result["err"] ?? null;
    echo "Error: " . ($err ? $err->getMessage() : "HTTP " . $result["status"]);
}
```

### Prepare a request without sending it

```php
// prepare() throws on error and returns the fetch definition.
$fetchdef = $client->prepare([
    "path" => "/api/resource/{id}",
    "method" => "DELETE",
    "params" => ["id" => "example"],
]);

echo $fetchdef["url"];
echo $fetchdef["method"];
print_r($fetchdef["headers"]);
```

### Use test mode

Create a mock client for unit testing — no server required. Seed fixture
data via the `entity` option so offline calls resolve without a live server:

```php
$client = PoetrydbSDK::test([
    "entity" => ["author" => ["test01" => ["id" => "test01"]]],
]);

// Entity ops return the bare mock record (throws on error).
$author = $client->Author()->list();
print_r($author);
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```php
$mock_fetch = function ($url, $init) {
    return [
        [
            "status" => 200,
            "statusText" => "OK",
            "headers" => [],
            "json" => function () { return ["id" => "mock01"]; },
        ],
        null,
    ];
};

$client = new PoetrydbSDK([
    "base" => "http://localhost:8080",
    "system" => [
        "fetch" => $mock_fetch,
    ],
]);
```

### Run live tests

Create a `.env.local` file at the project root:

```
POETRYDB_TEST_LIVE=TRUE
```

Then run:

```bash
cd php && ./vendor/bin/phpunit test/
```


## Reference

### PoetrydbSDK

```php
require_once 'poetrydb_sdk.php';
$client = new PoetrydbSDK($options);
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `base` | `string` | Base URL of the API server. |
| `prefix` | `string` | URL path prefix prepended to all requests. |
| `suffix` | `string` | URL path suffix appended to all requests. |
| `feature` | `array` | Feature activation flags. |
| `extend` | `array` | Additional Feature instances to load. |
| `system` | `array` | System overrides (e.g. custom `fetch` callable). |

### test

```php
$client = PoetrydbSDK::test($testopts, $sdkopts);
```

Creates a test-mode client with mock transport. Both arguments may be `null`.

### PoetrydbSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `(): array` | Deep copy of current SDK options. |
| `get_utility` | `(): Utility` | Copy of the SDK utility object. |
| `prepare` | `(array $fetchargs): array` | Build an HTTP request definition without sending. |
| `direct` | `(array $fetchargs): array` | Build and send an HTTP request. |
| `Author` | `($data): AuthorEntity` | Create an Author entity instance. |
| `Authorab` | `($data): AuthorabEntity` | Create an Authorab entity instance. |
| `CombinedSearch` | `($data): CombinedSearchEntity` | Create a CombinedSearch entity instance. |
| `CombinedSearchWithField` | `($data): CombinedSearchWithFieldEntity` | Create a CombinedSearchWithField entity instance. |
| `Line` | `($data): LineEntity` | Create a Line entity instance. |
| `Linecount` | `($data): LinecountEntity` | Create a Linecount entity instance. |
| `Poemcount` | `($data): PoemcountEntity` | Create a Poemcount entity instance. |
| `Random` | `($data): RandomEntity` | Create a Random entity instance. |
| `Title` | `($data): TitleEntity` | Create a Title entity instance. |
| `Titleab` | `($data): TitleabEntity` | Create a Titleab entity instance. |

### Entity interface

All entities share the same interface.

| Method | Signature | Description |
| --- | --- | --- |
| `load` | `($reqmatch, $ctrl): array` | Load a single entity by match criteria. |
| `list` | `(?array $reqmatch = null, $ctrl): array` | List entities matching the criteria (call with no argument to list all). |
| `data_get` | `(): array` | Get entity data. |
| `data_set` | `($data): void` | Set entity data. |
| `match_get` | `(): array` | Get entity match criteria. |
| `match_set` | `($match): void` | Set entity match criteria. |
| `make` | `(): Entity` | Create a new instance with the same options. |
| `get_name` | `(): string` | Return the entity name. |

### Result shape

Entity operations return the bare result data (an `array` for single-entity
ops, a `list` for `list`) and throw on error. Wrap calls in
`try`/`catch` to handle failures.

The `direct()` escape hatch never throws — it returns a result `array`
you branch on via `$result["ok"]`:

| Key | Type | Description |
| --- | --- | --- |
| `ok` | `bool` | `true` if the HTTP status is 2xx. |
| `status` | `int` | HTTP status code. |
| `headers` | `array` | Response headers. |
| `data` | `mixed` | Parsed JSON response body. |

On error, `ok` is `false` and `$err` contains the error value.

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

Create an instance: `$author = $client->Author();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `author` | `string` |  |
| `line` | `array` |  |
| `linecount` | `int` |  |
| `title` | `string` |  |

#### Example: Load

```php
// load() returns the bare Author record (throws on error).
$author = $client->Author()->load(["id" => "author_id"]);
```

#### Example: List

```php
// list() returns an array of Author records (throws on error).
$authors = $client->Author()->list();
```


### Authorab

Create an instance: `$authorab = $client->Authorab();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `author` | `string` |  |
| `line` | `array` |  |
| `linecount` | `int` |  |
| `title` | `string` |  |

#### Example: List

```php
// list() returns an array of Authorab records (throws on error).
$authorabs = $client->Authorab()->list();
```


### CombinedSearch

Create an instance: `$combined_search = $client->CombinedSearch();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `author` | `string` |  |
| `line` | `array` |  |
| `linecount` | `int` |  |
| `title` | `string` |  |

#### Example: List

```php
// list() returns an array of CombinedSearch records (throws on error).
$combined_searchs = $client->CombinedSearch()->list();
```


### CombinedSearchWithField

Create an instance: `$combined_search_with_field = $client->CombinedSearchWithField();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Example: List

```php
// list() returns an array of CombinedSearchWithField records (throws on error).
$combined_search_with_fields = $client->CombinedSearchWithField()->list();
```


### Line

Create an instance: `$line = $client->Line();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `author` | `string` |  |
| `line` | `array` |  |
| `linecount` | `int` |  |
| `title` | `string` |  |

#### Example: Load

```php
// load() returns the bare Line record (throws on error).
$line = $client->Line()->load(["id" => "line_id"]);
```

#### Example: List

```php
// list() returns an array of Line records (throws on error).
$lines = $client->Line()->list();
```


### Linecount

Create an instance: `$linecount = $client->Linecount();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `author` | `string` |  |
| `line` | `array` |  |
| `linecount` | `int` |  |
| `title` | `string` |  |

#### Example: Load

```php
// load() returns the bare Linecount record (throws on error).
$linecount = $client->Linecount()->load(["id" => "linecount_id"]);
```

#### Example: List

```php
// list() returns an array of Linecount records (throws on error).
$linecounts = $client->Linecount()->list();
```


### Poemcount

Create an instance: `$poemcount = $client->Poemcount();`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `author` | `string` |  |
| `line` | `array` |  |
| `linecount` | `int` |  |
| `title` | `string` |  |

#### Example: Load

```php
// load() returns the bare Poemcount record (throws on error).
$poemcount = $client->Poemcount()->load(["id" => "poemcount_id"]);
```


### Random

Create an instance: `$random = $client->Random();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `author` | `string` |  |
| `line` | `array` |  |
| `linecount` | `int` |  |
| `title` | `string` |  |

#### Example: Load

```php
// load() returns the bare Random record (throws on error).
$random = $client->Random()->load(["id" => "random_id"]);
```

#### Example: List

```php
// list() returns an array of Random records (throws on error).
$randoms = $client->Random()->list();
```


### Title

Create an instance: `$title = $client->Title();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `author` | `string` |  |
| `line` | `array` |  |
| `linecount` | `int` |  |
| `title` | `string` |  |

#### Example: Load

```php
// load() returns the bare Title record (throws on error).
$title = $client->Title()->load(["id" => "title_id"]);
```

#### Example: List

```php
// list() returns an array of Title records (throws on error).
$titles = $client->Title()->list();
```


### Titleab

Create an instance: `$titleab = $client->Titleab();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `author` | `string` |  |
| `line` | `array` |  |
| `linecount` | `int` |  |
| `title` | `string` |  |

#### Example: List

```php
// list() returns an array of Titleab records (throws on error).
$titleabs = $client->Titleab()->list();
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

Features are the extension mechanism. A feature is a PHP class
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as arrays

The PHP SDK uses plain PHP associative arrays throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `Helpers::to_map()` to safely validate that a value is an array.

### Directory structure

```
php/
├── poetrydb_sdk.php          -- Main SDK class
├── config.php                     -- Configuration
├── features.php                   -- Feature factory
├── core/                          -- Core types and context
├── entity/                        -- Entity implementations
├── feature/                       -- Built-in features (Base, Test, Log)
├── utility/                       -- Utility functions and struct library
└── test/                          -- Test suites
```

The main class (`poetrydb_sdk.php`) exports the SDK class
and test helper. Import entity or utility modules directly only
when needed.

### Entity state

Entity instances are stateful. After a successful `list`, the entity
stores the returned data and match criteria internally.

```php
$author = $client->Author();
$author->list();

// $author->data_get() now returns the author data from the last list
// $author->match_get() returns the last match criteria
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
