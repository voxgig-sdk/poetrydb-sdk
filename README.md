# Poetrydb SDK

Search and retrieve classic poetry by author, title, lines, or line count, with JSON results

> TypeScript, Python, PHP, Golang, Ruby, Lua SDKs, a CLI, an interactive REPL, and an MCP server for AI agents — all generated from one OpenAPI spec by [@voxgig/sdkgen](https://github.com/voxgig/sdkgen).

## About PoetryDB API

[PoetryDB](https://poetrydb.org) is a free, public API that exposes a curated collection of classic English-language poetry as JSON. The project's concept and code are by [@thundercomb](https://twitter.com/thundercomb), with the poems stored in a MongoDB database and served by an open-source Ruby application.

What you get from the API:
- Search poems by `author`, `title`, `lines` (full-text match against the body), or exact `linecount`.
- Combine multiple search fields in a single request, with optional `:abs` modifier for exact matches.
- Retrieve random poems via `/random[/<count>]`.
- For each poem, the API returns the fields `title`, `author`, `lines` (array of strings) and `linecount`.
- Results are available in `.json` (default) or `.text` format.

The service has CORS enabled and requires no authentication or API key. The README does not document a rate limit, but as a free community service it should be used courteously.

## Try it

**TypeScript**
```bash
npm install poetrydb
```

**Python**
```bash
pip install poetrydb-sdk
```

**PHP**
```bash
composer require voxgig/poetrydb-sdk
```

**Golang**
```bash
go get github.com/voxgig-sdk/poetrydb-sdk/go
```

**Ruby**
```bash
gem install poetrydb-sdk
```

**Lua**
```bash
luarocks install poetrydb-sdk
```

## 30-second quickstart

### TypeScript

```ts
import { PoetrydbSDK } from 'poetrydb'

const client = new PoetrydbSDK({})

// List all authors
const authors = await client.Author().list()
```

See the [TypeScript README](ts/README.md) for the
full guide, or scroll down for the same example in other languages.

## What's in the box

| Surface | Use it for | Path |
| --- | --- | --- |
| **SDK** (TypeScript, Python, PHP, Golang, Ruby, Lua) | App integration | `ts/` `py/` `php/` `go/` `rb/` `lua/` |
| **CLI** | Scripts, CI, ops, one-off API calls | `go-cli/` |
| **MCP server** | AI agents (Claude, Cursor, Cline) | `go-mcp/` |

## Use it from an AI agent (MCP)

The generated MCP server exposes every operation in this SDK as an
[MCP](https://modelcontextprotocol.io) tool that Claude, Cursor or Cline
can call directly. Build and register it:

```bash
cd go-mcp && go build -o poetrydb-mcp .
```

Then add it to your agent's MCP config (Claude Desktop, Cursor, etc.):

```json
{
  "mcpServers": {
    "poetrydb": {
      "command": "/abs/path/to/poetrydb-mcp"
    }
  }
}
```

## Entities

The API exposes 10 entities:

| Entity | Description | API path |
| --- | --- | --- |
| **Author** | Search poems by author name; path `/author/<author>[/<output field>][.<format>]`. | `/author/{author}/{outputFields}.{format}` |
| **Authorab** | Exact-match (`:abs`) variant for author searches; path `/author/<author>:abs[/<output field>][.<format>]`. | `/author/{author}:abs` |
| **CombinedSearch** | Combine multiple input fields (e.g. author and title) in one query, separated by semicolons. | `/{inputField1},{inputField2}/{searchTerm1};{searchTerm2}` |
| **CombinedSearchWithField** | Combined search that also selects which output field(s) to return per poem. | `/{inputField1},{inputField2}/{searchTerm1};{searchTerm2}/{outputFields}` |
| **Line** | Full-text search against the lines of poems; path `/lines/<lines>[:abs][/<output field>][.<format>]`. | `/lines/{lines}/{outputFields}.{format}` |
| **Linecount** | Find poems with an exact number of lines; path `/linecount/<linecount>[/<output field>][.<format>]`. | `/linecount/{linecount}/{outputFields}.{format}` |
| **Poemcount** | Limit or count results when combined with other search fields. | `/poemcount/{count}` |
| **Random** | Return one or more randomly chosen poems; path `/random[/<count>][/<output field>][.<format>]`. | `/random/{count}/{outputFields}` |
| **Title** | Search poems by title; path `/title/<title>[:abs][/<output field>][.<format>]`. | `/title/{title}/{outputFields}.{format}` |
| **Titleab** | Exact-match (`:abs`) variant for title searches; path `/title/<title>:abs[/<output field>][.<format>]`. | `/title/{title}:abs` |

Each entity supports the following operations where available: **load**,
**list**, **create**, **update**, and **remove**.

## Quickstart in other languages

### Python

```python
from poetrydb_sdk import PoetrydbSDK

client = PoetrydbSDK({})

# List all authors
authors, err = client.Author(None).list(None, None)

# Load a specific author
author, err = client.Author(None).load(
    {"id": "example_id"}, None
)
```

### PHP

```php
<?php
require_once 'poetrydb_sdk.php';

$client = new PoetrydbSDK([]);

// List all authors
[$authors, $err] = $client->Author(null)->list(null, null);

// Load a specific author
[$author, $err] = $client->Author(null)->load(
    ["id" => "example_id"], null
);
```

### Golang

```go
import sdk "github.com/voxgig-sdk/poetrydb-sdk/go"

client := sdk.NewPoetrydbSDK(map[string]any{})

// List all authors
authors, err := client.Author(nil).List(nil, nil)
```

### Ruby

```ruby
require_relative "Poetrydb_sdk"

client = PoetrydbSDK.new({})

# List all authors
authors, err = client.Author(nil).list(nil, nil)

# Load a specific author
author, err = client.Author(nil).load(
  { "id" => "example_id" }, nil
)
```

### Lua

```lua
local sdk = require("poetrydb_sdk")

local client = sdk.new({})

-- List all authors
local authors, err = client:Author(nil):list(nil, nil)

-- Load a specific author
local author, err = client:Author(nil):load(
  { id = "example_id" }, nil
)
```

## Unit testing in offline mode

Every SDK ships a test mode that swaps the HTTP transport for an
in-memory mock, so unit tests run offline.

### TypeScript

```ts
const client = PoetrydbSDK.test()
const result = await client.Author().load({ id: 'test01' })
// result.ok === true, result.data contains mock data
```

### Python

```python
client = PoetrydbSDK.test(None, None)
result, err = client.Author(None).load(
    {"id": "test01"}, None
)
```

### PHP

```php
$client = PoetrydbSDK::test(null, null);
[$result, $err] = $client->Author(null)->load(
    ["id" => "test01"], null
);
```

### Golang

```go
client := sdk.TestSDK(nil, nil)
result, err := client.Author(nil).Load(
    map[string]any{"id": "test01"}, nil,
)
```

### Ruby

```ruby
client = PoetrydbSDK.test(nil, nil)
result, err = client.Author(nil).load(
  { "id" => "test01" }, nil
)
```

### Lua

```lua
local client = sdk.test(nil, nil)
local result, err = client:Author(nil):load(
  { id = "test01" }, nil
)
```

## How it works

Every SDK call runs the same five-stage pipeline:

1. **Point** — resolve the API endpoint from the operation definition.
2. **Spec** — build the HTTP specification (URL, method, headers, body).
3. **Request** — send the HTTP request.
4. **Response** — receive and parse the response.
5. **Result** — extract the result data for the caller.

A feature hook fires at each stage (e.g. `PrePoint`, `PreSpec`,
`PreRequest`), so features can inspect or modify the pipeline without
forking the SDK.

### Features

| Feature | Purpose |
| --- | --- |
| **TestFeature** | In-memory mock transport for testing without a live server |

Pass custom features via the `extend` option at construction time.

### Direct and Prepare

For endpoints the entity model doesn't cover, use the low-level methods:

- **`direct(fetchargs)`** — build and send an HTTP request in one step.
- **`prepare(fetchargs)`** — build the request without sending it.

Both accept a map with `path`, `method`, `params`, `query`,
`headers`, and `body`. See the [How-to guides](#how-to-guides) below.

## How-to guides

### Make a direct API call

When the entity interface does not cover an endpoint, use `direct`:

**TypeScript:**
```ts
const result = await client.direct({
  path: '/api/resource/{id}',
  method: 'GET',
  params: { id: 'example' },
})
console.log(result.data)
```

**Python:**
```python
result, err = client.direct({
    "path": "/api/resource/{id}",
    "method": "GET",
    "params": {"id": "example"},
})
```

**PHP:**
```php
[$result, $err] = $client->direct([
    "path" => "/api/resource/{id}",
    "method" => "GET",
    "params" => ["id" => "example"],
]);
```

**Go:**
```go
result, err := client.Direct(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "GET",
    "params": map[string]any{"id": "example"},
})
```

**Ruby:**
```ruby
result, err = client.direct({
  "path" => "/api/resource/{id}",
  "method" => "GET",
  "params" => { "id" => "example" },
})
```

**Lua:**
```lua
local result, err = client:direct({
  path = "/api/resource/{id}",
  method = "GET",
  params = { id = "example" },
})
```

## Per-language documentation

- [TypeScript](ts/README.md)
- [Python](py/README.md)
- [PHP](php/README.md)
- [Golang](go/README.md)
- [Ruby](rb/README.md)
- [Lua](lua/README.md)

## Using the PoetryDB API

- Upstream: [https://poetrydb.org](https://poetrydb.org)
- API docs: [https://github.com/thundercomb/poetrydb/blob/master/README.md](https://github.com/thundercomb/poetrydb/blob/master/README.md)

- Released under the [GNU General Public License v2](https://www.gnu.org/licenses/old-licenses/gpl-2.0.html).
- The Ruby source code for the API is open source and available on [GitHub](https://github.com/thundercomb/poetrydb).
- Attribution to the PoetryDB project (concept and code by [@thundercomb](https://twitter.com/thundercomb)) is appreciated when redistributing data.

---

Generated from the PoetryDB API OpenAPI spec by [@voxgig/sdkgen](https://github.com/voxgig/sdkgen).
