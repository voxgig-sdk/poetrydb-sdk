# Poetrydb SDK configuration


# The sekreto plugin DEFINITIONS the model selected per feature, imported
# above by name from the modules the catalogue's active `plugin.def`
# entries declare. Handed to each feature (secrets builds its Sekreto
# with them): a provider kind not listed here is unknown to that SDK.
FEATURE_PLUGINS = {
}


_shared_config = None


def shared_config():
    """Return the process-wide config, built once on first use.

    The SDK reads the config on every request and never writes to it, so one
    instance is shared by every client rather than rebuilt per client.

    The returned dict is shared: treat it as read-only. Callers that need to
    mutate should use make_config, which always returns a fresh copy.
    """
    global _shared_config
    if _shared_config is None:
        _shared_config = make_config()
    return _shared_config


def make_config():
    """Build a fresh, fully materialised config dict.

    Every call rebuilds the whole structure, so prefer shared_config unless
    you need a private copy you intend to mutate.
    """
    return {
        "main": {
            "name": "Poetrydb",
            "slug": "poetrydb",
            "version": "0.0.1",
            "target": "py",
        },
        "feature": {
            "test": {
        "options": {
          "active": False,
        },
        "transport": "base",
      },
        },
        "options": {
            "base": "https://poetrydb.org",
            "headers": {
        "content-type": "application/json",
      },
            "entity": {
                "author": {},
                "authorab": {},
                "combined_search": {},
                "combined_search_with_field": {},
                "line": {},
                "linecount": {},
                "poemcount": {},
                "random": {},
                "title": {},
                "titleab": {},
            },
        },
        "entity": {
      "author": {
        "fields": [
          {
            "name": "author",
            "short": "The author of the poem",
            "type": "`$STRING`",
          },
          {
            "name": "authors",
            "type": "`$ARRAY`",
          },
          {
            "name": "id",
            "type": "`$STRING`",
          },
          {
            "name": "linecount",
            "short": "The number of lines in the poem (including section headings, excluding empty lines)",
            "type": "`$INTEGER`",
          },
          {
            "name": "lines",
            "short": "The lines of the poem",
            "type": "`$ARRAY`",
          },
          {
            "name": "title",
            "short": "The title of the poem",
            "type": "`$STRING`",
          },
        ],
        "id": {
          "field": "id",
          "name": "id",
        },
        "name": "author",
        "op": {
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "args": {
                  "params": [
                    {
                      "example": "Ernest Dowson",
                      "kind": "param",
                      "name": "author",
                      "orig": "author",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                    {
                      "example": "text",
                      "kind": "param",
                      "name": "format",
                      "orig": "format",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                    {
                      "example": "author,title,linecount",
                      "kind": "param",
                      "name": "output_field",
                      "orig": "output_field",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/author/{author}/{outputFields}.{format}",
                "rename": {
                  "param": {
                    "outputFields}.{format": "output_fields}_{format",
                  },
                },
                "segments": [
                  {
                    "lit": "author",
                  },
                  {
                    "var": "author",
                  },
                  {
                    "lit": "{outputFields}.{format}",
                  },
                ],
                "select": {
                  "exist": [
                    "author",
                    "format",
                    "output_field",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "parts": [
                  "author",
                  "{author}",
                  "{outputFields}.{format}",
                ],
              },
              {
                "args": {
                  "params": [
                    {
                      "example": "Ernest Dowson",
                      "kind": "param",
                      "name": "author",
                      "orig": "author",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                    {
                      "example": "author,title,linecount",
                      "kind": "param",
                      "name": "output_field",
                      "orig": "output_field",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/author/{author}/{outputFields}",
                "rename": {
                  "param": {
                    "outputFields": "output_field",
                  },
                },
                "segments": [
                  {
                    "lit": "author",
                  },
                  {
                    "var": "author",
                  },
                  {
                    "var": "output_field",
                  },
                ],
                "select": {
                  "exist": [
                    "author",
                    "output_field",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "parts": [
                  "author",
                  "{author}",
                  "{output_field}",
                ],
              },
              {
                "args": {},
                "kind": "http",
                "method": "GET",
                "orig": "/author",
                "segments": [
                  {
                    "lit": "author",
                  },
                ],
                "select": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body.authors`",
                },
                "parts": [
                  "author",
                ],
              },
            ],
          },
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "args": {
                  "params": [
                    {
                      "example": "Ernest Dowson",
                      "kind": "param",
                      "name": "id",
                      "orig": "author",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/author/{author}",
                "rename": {
                  "param": {
                    "author": "id",
                  },
                },
                "segments": [
                  {
                    "lit": "author",
                  },
                  {
                    "var": "id",
                  },
                ],
                "select": {
                  "exist": [
                    "id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "parts": [
                  "author",
                  "{id}",
                ],
              },
            ],
          },
        },
        "relations": {
          "ancestors": [
            [
              "author",
            ],
          ],
        },
      },
      "authorab": {
        "fields": [
          {
            "name": "author",
            "short": "The author of the poem",
            "type": "`$STRING`",
          },
          {
            "name": "linecount",
            "short": "The number of lines in the poem (including section headings, excluding empty lines)",
            "type": "`$INTEGER`",
          },
          {
            "name": "lines",
            "short": "The lines of the poem",
            "type": "`$ARRAY`",
          },
          {
            "name": "title",
            "short": "The title of the poem",
            "type": "`$STRING`",
          },
        ],
        "name": "authorab",
        "op": {
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "args": {
                  "params": [
                    {
                      "example": "Ernest Dowson",
                      "kind": "param",
                      "name": "author",
                      "orig": "author",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/author/{author}:abs",
                "segments": [
                  {
                    "lit": "author",
                  },
                  {
                    "lit": "{author}:abs",
                  },
                ],
                "select": {
                  "exist": [
                    "author",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "parts": [
                  "author",
                  "{author}:abs",
                ],
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "combined_search": {
        "fields": [
          {
            "name": "author",
            "short": "The author of the poem",
            "type": "`$STRING`",
          },
          {
            "name": "linecount",
            "short": "The number of lines in the poem (including section headings, excluding empty lines)",
            "type": "`$INTEGER`",
          },
          {
            "name": "lines",
            "short": "The lines of the poem",
            "type": "`$ARRAY`",
          },
          {
            "name": "title",
            "short": "The title of the poem",
            "type": "`$STRING`",
          },
        ],
        "name": "combined_search",
        "op": {
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "args": {
                  "params": [
                    {
                      "example": "author",
                      "kind": "param",
                      "name": "input_field1",
                      "orig": "input_field1",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                    {
                      "example": "linecount",
                      "kind": "param",
                      "name": "input_field2",
                      "orig": "input_field2",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                    {
                      "example": "Shakespeare",
                      "kind": "param",
                      "name": "search_term1",
                      "orig": "search_term1",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                    {
                      "example": "14",
                      "kind": "param",
                      "name": "search_term2",
                      "orig": "search_term2",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/{inputField1},{inputField2}/{searchTerm1};{searchTerm2}",
                "rename": {
                  "param": {
                    "inputField1},{inputField2": "input_field1},{input_field2",
                    "searchTerm1};{searchTerm2": "search_term1};{search_term2",
                  },
                },
                "segments": [
                  {
                    "lit": "{inputField1},{inputField2}",
                  },
                  {
                    "lit": "{searchTerm1};{searchTerm2}",
                  },
                ],
                "select": {
                  "exist": [
                    "input_field1",
                    "input_field2",
                    "search_term1",
                    "search_term2",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "parts": [
                  "{inputField1},{inputField2}",
                  "{searchTerm1};{searchTerm2}",
                ],
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "combined_search_with_field": {
        "fields": [],
        "name": "combined_search_with_field",
        "op": {
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "args": {
                  "params": [
                    {
                      "example": "author",
                      "kind": "param",
                      "name": "input_field1",
                      "orig": "input_field1",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                    {
                      "example": "linecount",
                      "kind": "param",
                      "name": "input_field2",
                      "orig": "input_field2",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                    {
                      "example": "lines",
                      "kind": "param",
                      "name": "output_field",
                      "orig": "output_field",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                    {
                      "example": "Shakespeare",
                      "kind": "param",
                      "name": "search_term1",
                      "orig": "search_term1",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                    {
                      "example": "14",
                      "kind": "param",
                      "name": "search_term2",
                      "orig": "search_term2",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/{inputField1},{inputField2}/{searchTerm1};{searchTerm2}/{outputFields}",
                "rename": {
                  "param": {
                    "inputField1},{inputField2": "input_field1},{input_field2",
                    "outputFields": "output_field",
                    "searchTerm1};{searchTerm2": "search_term1};{search_term2",
                  },
                },
                "segments": [
                  {
                    "lit": "{inputField1},{inputField2}",
                  },
                  {
                    "lit": "{searchTerm1};{searchTerm2}",
                  },
                  {
                    "var": "output_field",
                  },
                ],
                "select": {
                  "exist": [
                    "input_field1",
                    "input_field2",
                    "output_field",
                    "search_term1",
                    "search_term2",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "parts": [
                  "{inputField1},{inputField2}",
                  "{searchTerm1};{searchTerm2}",
                  "{output_field}",
                ],
              },
            ],
          },
        },
        "relations": {
          "ancestors": [
            [
              "{search_term1};{search_term2}",
            ],
          ],
        },
      },
      "line": {
        "fields": [
          {
            "name": "author",
            "short": "The author of the poem",
            "type": "`$STRING`",
          },
          {
            "name": "id",
            "type": "`$STRING`",
          },
          {
            "name": "linecount",
            "short": "The number of lines in the poem (including section headings, excluding empty lines)",
            "type": "`$INTEGER`",
          },
          {
            "name": "lines",
            "short": "The lines of the poem",
            "type": "`$ARRAY`",
          },
          {
            "name": "title",
            "short": "The title of the poem",
            "type": "`$STRING`",
          },
        ],
        "id": {
          "field": "id",
          "name": "id",
        },
        "name": "line",
        "op": {
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "args": {
                  "params": [
                    {
                      "example": "text",
                      "kind": "param",
                      "name": "format",
                      "orig": "format",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                    {
                      "example": "Latitudeless Place",
                      "kind": "param",
                      "name": "line",
                      "orig": "line",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                    {
                      "example": "author,title,linecount",
                      "kind": "param",
                      "name": "output_field",
                      "orig": "output_field",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/lines/{lines}/{outputFields}.{format}",
                "rename": {
                  "param": {
                    "lines": "line",
                    "outputFields}.{format": "output_fields}_{format",
                  },
                },
                "segments": [
                  {
                    "lit": "lines",
                  },
                  {
                    "var": "line",
                  },
                  {
                    "lit": "{outputFields}.{format}",
                  },
                ],
                "select": {
                  "exist": [
                    "format",
                    "line",
                    "output_field",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "parts": [
                  "lines",
                  "{line}",
                  "{outputFields}.{format}",
                ],
              },
              {
                "args": {
                  "params": [
                    {
                      "example": "Latitudeless Place",
                      "kind": "param",
                      "name": "line",
                      "orig": "line",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                    {
                      "example": "author,title,linecount",
                      "kind": "param",
                      "name": "output_field",
                      "orig": "output_field",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/lines/{lines}/{outputFields}",
                "rename": {
                  "param": {
                    "lines": "line",
                    "outputFields": "output_field",
                  },
                },
                "segments": [
                  {
                    "lit": "lines",
                  },
                  {
                    "var": "line",
                  },
                  {
                    "var": "output_field",
                  },
                ],
                "select": {
                  "exist": [
                    "line",
                    "output_field",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "parts": [
                  "lines",
                  "{line}",
                  "{output_field}",
                ],
              },
            ],
          },
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "args": {
                  "params": [
                    {
                      "example": "Latitudeless Place",
                      "kind": "param",
                      "name": "id",
                      "orig": "line",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/lines/{lines}",
                "rename": {
                  "param": {
                    "lines": "id",
                  },
                },
                "segments": [
                  {
                    "lit": "lines",
                  },
                  {
                    "var": "id",
                  },
                ],
                "select": {
                  "exist": [
                    "id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "parts": [
                  "lines",
                  "{id}",
                ],
              },
            ],
          },
        },
        "relations": {
          "ancestors": [
            [
              "line",
            ],
          ],
        },
      },
      "linecount": {
        "fields": [
          {
            "name": "author",
            "short": "The author of the poem",
            "type": "`$STRING`",
          },
          {
            "name": "id",
            "type": "`$STRING`",
          },
          {
            "name": "linecount",
            "short": "The number of lines in the poem (including section headings, excluding empty lines)",
            "type": "`$INTEGER`",
          },
          {
            "name": "lines",
            "short": "The lines of the poem",
            "type": "`$ARRAY`",
          },
          {
            "name": "title",
            "short": "The title of the poem",
            "type": "`$STRING`",
          },
        ],
        "id": {
          "field": "id",
          "name": "id",
        },
        "name": "linecount",
        "op": {
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "args": {
                  "params": [
                    {
                      "example": "text",
                      "kind": "param",
                      "name": "format",
                      "orig": "format",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                    {
                      "example": 14,
                      "kind": "param",
                      "name": "linecount",
                      "orig": "linecount",
                      "reqd": True,
                      "type": "`$INTEGER`",
                    },
                    {
                      "example": "author,title",
                      "kind": "param",
                      "name": "output_field",
                      "orig": "output_field",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/linecount/{linecount}/{outputFields}.{format}",
                "rename": {
                  "param": {
                    "outputFields}.{format": "output_fields}_{format",
                  },
                },
                "segments": [
                  {
                    "lit": "linecount",
                  },
                  {
                    "var": "linecount",
                  },
                  {
                    "lit": "{outputFields}.{format}",
                  },
                ],
                "select": {
                  "exist": [
                    "format",
                    "linecount",
                    "output_field",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "parts": [
                  "linecount",
                  "{linecount}",
                  "{outputFields}.{format}",
                ],
              },
              {
                "args": {
                  "params": [
                    {
                      "example": 14,
                      "kind": "param",
                      "name": "linecount",
                      "orig": "linecount",
                      "reqd": True,
                      "type": "`$INTEGER`",
                    },
                    {
                      "example": "author,title",
                      "kind": "param",
                      "name": "output_field",
                      "orig": "output_field",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/linecount/{linecount}/{outputFields}",
                "rename": {
                  "param": {
                    "outputFields": "output_field",
                  },
                },
                "segments": [
                  {
                    "lit": "linecount",
                  },
                  {
                    "var": "linecount",
                  },
                  {
                    "var": "output_field",
                  },
                ],
                "select": {
                  "exist": [
                    "linecount",
                    "output_field",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "parts": [
                  "linecount",
                  "{linecount}",
                  "{output_field}",
                ],
              },
            ],
          },
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "args": {
                  "params": [
                    {
                      "example": 14,
                      "kind": "param",
                      "name": "id",
                      "orig": "linecount",
                      "reqd": True,
                      "type": "`$INTEGER`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/linecount/{linecount}",
                "rename": {
                  "param": {
                    "linecount": "id",
                  },
                },
                "segments": [
                  {
                    "lit": "linecount",
                  },
                  {
                    "var": "id",
                  },
                ],
                "select": {
                  "exist": [
                    "id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "parts": [
                  "linecount",
                  "{id}",
                ],
              },
            ],
          },
        },
        "relations": {
          "ancestors": [
            [
              "linecount",
            ],
          ],
        },
      },
      "poemcount": {
        "fields": [
          {
            "name": "author",
            "short": "The author of the poem",
            "type": "`$STRING`",
          },
          {
            "name": "id",
            "type": "`$STRING`",
          },
          {
            "name": "linecount",
            "short": "The number of lines in the poem (including section headings, excluding empty lines)",
            "type": "`$INTEGER`",
          },
          {
            "name": "lines",
            "short": "The lines of the poem",
            "type": "`$ARRAY`",
          },
          {
            "name": "title",
            "short": "The title of the poem",
            "type": "`$STRING`",
          },
        ],
        "id": {
          "field": "id",
          "name": "id",
        },
        "name": "poemcount",
        "op": {
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "args": {
                  "params": [
                    {
                      "example": 10,
                      "kind": "param",
                      "name": "id",
                      "orig": "count",
                      "reqd": True,
                      "type": "`$INTEGER`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/poemcount/{count}",
                "rename": {
                  "param": {
                    "count": "id",
                  },
                },
                "segments": [
                  {
                    "lit": "poemcount",
                  },
                  {
                    "var": "id",
                  },
                ],
                "select": {
                  "exist": [
                    "id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "parts": [
                  "poemcount",
                  "{id}",
                ],
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "random": {
        "fields": [
          {
            "name": "author",
            "short": "The author of the poem",
            "type": "`$STRING`",
          },
          {
            "name": "id",
            "type": "`$STRING`",
          },
          {
            "name": "linecount",
            "short": "The number of lines in the poem (including section headings, excluding empty lines)",
            "type": "`$INTEGER`",
          },
          {
            "name": "lines",
            "short": "The lines of the poem",
            "type": "`$ARRAY`",
          },
          {
            "name": "title",
            "short": "The title of the poem",
            "type": "`$STRING`",
          },
        ],
        "id": {
          "field": "id",
          "name": "id",
        },
        "name": "random",
        "op": {
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "args": {
                  "params": [
                    {
                      "example": 5,
                      "kind": "param",
                      "name": "count",
                      "orig": "count",
                      "reqd": True,
                      "type": "`$INTEGER`",
                    },
                    {
                      "example": "author,title",
                      "kind": "param",
                      "name": "output_field",
                      "orig": "output_field",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/random/{count}/{outputFields}",
                "rename": {
                  "param": {
                    "outputFields": "output_field",
                  },
                },
                "segments": [
                  {
                    "lit": "random",
                  },
                  {
                    "var": "count",
                  },
                  {
                    "var": "output_field",
                  },
                ],
                "select": {
                  "exist": [
                    "count",
                    "output_field",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "parts": [
                  "random",
                  "{count}",
                  "{output_field}",
                ],
              },
            ],
          },
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "args": {
                  "params": [
                    {
                      "example": 5,
                      "kind": "param",
                      "name": "id",
                      "orig": "count",
                      "reqd": True,
                      "type": "`$INTEGER`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/random/{count}",
                "rename": {
                  "param": {
                    "count": "id",
                  },
                },
                "segments": [
                  {
                    "lit": "random",
                  },
                  {
                    "var": "id",
                  },
                ],
                "select": {
                  "exist": [
                    "id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "parts": [
                  "random",
                  "{id}",
                ],
              },
            ],
          },
        },
        "relations": {
          "ancestors": [
            [
              "random",
            ],
          ],
        },
      },
      "title": {
        "fields": [
          {
            "name": "author",
            "short": "The author of the poem",
            "type": "`$STRING`",
          },
          {
            "name": "id",
            "type": "`$STRING`",
          },
          {
            "name": "linecount",
            "short": "The number of lines in the poem (including section headings, excluding empty lines)",
            "type": "`$INTEGER`",
          },
          {
            "name": "lines",
            "short": "The lines of the poem",
            "type": "`$ARRAY`",
          },
          {
            "name": "title",
            "short": "The title of the poem",
            "type": "`$STRING`",
          },
          {
            "name": "titles",
            "type": "`$ARRAY`",
          },
        ],
        "id": {
          "field": "id",
          "name": "id",
        },
        "name": "title",
        "op": {
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "args": {
                  "params": [
                    {
                      "example": "text",
                      "kind": "param",
                      "name": "format",
                      "orig": "format",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                    {
                      "example": "title,lines",
                      "kind": "param",
                      "name": "output_field",
                      "orig": "output_field",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                    {
                      "example": "Ozymandias",
                      "kind": "param",
                      "name": "title",
                      "orig": "title",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/title/{title}/{outputFields}.{format}",
                "rename": {
                  "param": {
                    "outputFields}.{format": "output_fields}_{format",
                  },
                },
                "segments": [
                  {
                    "lit": "title",
                  },
                  {
                    "var": "title",
                  },
                  {
                    "lit": "{outputFields}.{format}",
                  },
                ],
                "select": {
                  "exist": [
                    "format",
                    "output_field",
                    "title",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "parts": [
                  "title",
                  "{title}",
                  "{outputFields}.{format}",
                ],
              },
              {
                "args": {
                  "params": [
                    {
                      "example": "author,title,linecount",
                      "kind": "param",
                      "name": "output_field",
                      "orig": "output_field",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                    {
                      "example": "Ozymandias",
                      "kind": "param",
                      "name": "title",
                      "orig": "title",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/title/{title}/{outputFields}",
                "rename": {
                  "param": {
                    "outputFields": "output_field",
                  },
                },
                "segments": [
                  {
                    "lit": "title",
                  },
                  {
                    "var": "title",
                  },
                  {
                    "var": "output_field",
                  },
                ],
                "select": {
                  "exist": [
                    "output_field",
                    "title",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "parts": [
                  "title",
                  "{title}",
                  "{output_field}",
                ],
              },
              {
                "args": {},
                "kind": "http",
                "method": "GET",
                "orig": "/title",
                "segments": [
                  {
                    "lit": "title",
                  },
                ],
                "select": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body.titles`",
                },
                "parts": [
                  "title",
                ],
              },
            ],
          },
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "args": {
                  "params": [
                    {
                      "example": "Ozymandias",
                      "kind": "param",
                      "name": "id",
                      "orig": "title",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/title/{title}",
                "rename": {
                  "param": {
                    "title": "id",
                  },
                },
                "segments": [
                  {
                    "lit": "title",
                  },
                  {
                    "var": "id",
                  },
                ],
                "select": {
                  "exist": [
                    "id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "parts": [
                  "title",
                  "{id}",
                ],
              },
            ],
          },
        },
        "relations": {
          "ancestors": [
            [
              "title",
            ],
          ],
        },
      },
      "titleab": {
        "fields": [
          {
            "name": "author",
            "short": "The author of the poem",
            "type": "`$STRING`",
          },
          {
            "name": "linecount",
            "short": "The number of lines in the poem (including section headings, excluding empty lines)",
            "type": "`$INTEGER`",
          },
          {
            "name": "lines",
            "short": "The lines of the poem",
            "type": "`$ARRAY`",
          },
          {
            "name": "title",
            "short": "The title of the poem",
            "type": "`$STRING`",
          },
        ],
        "name": "titleab",
        "op": {
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "args": {
                  "params": [
                    {
                      "example": "Ozymandias",
                      "kind": "param",
                      "name": "title",
                      "orig": "title",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/title/{title}:abs",
                "segments": [
                  {
                    "lit": "title",
                  },
                  {
                    "lit": "{title}:abs",
                  },
                ],
                "select": {
                  "exist": [
                    "title",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "parts": [
                  "title",
                  "{title}:abs",
                ],
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
    },
    }
