# Poetrydb SDK configuration

module PoetrydbConfig
  # Return the process-wide config, built once on first use. The SDK reads
  # the config on every request and never writes to it, so one instance is
  # shared by every client rather than rebuilt per client.
  #
  # The returned hash is shared: treat it as read-only. Callers that need to
  # mutate should use make_config, which always returns a fresh copy.
  def self.shared_config
    @shared_config ||= make_config
  end


  # Build a fresh, fully materialised config hash. Every call rebuilds the
  # whole structure, so prefer shared_config unless you need a private copy
  # you intend to mutate.
  def self.make_config
    {
      "main" => {
        "name" => "Poetrydb",
      },
      "feature" => {
        "test" => {
          "options" => {
            "active" => false,
          },
        },
      },
      "options" => {
        "base" => "https://poetrydb.org",
        "headers" => {
          "content-type" => "application/json",
        },
        "entity" => {
          "author" => {},
          "authorab" => {},
          "combined_search" => {},
          "combined_search_with_field" => {},
          "line" => {},
          "linecount" => {},
          "poemcount" => {},
          "random" => {},
          "title" => {},
          "titleab" => {},
        },
      },
      "entity" => {
        "author" => {
          "fields" => [
            {
              "name" => "author",
              "type" => "`$STRING`",
            },
            {
              "name" => "authors",
              "type" => "`$ARRAY`",
            },
            {
              "name" => "linecount",
              "type" => "`$INTEGER`",
            },
            {
              "name" => "lines",
              "type" => "`$ARRAY`",
            },
            {
              "name" => "title",
              "type" => "`$STRING`",
            },
          ],
          "name" => "author",
          "op" => {
            "list" => {
              "input" => "data",
              "name" => "list",
              "points" => [
                {
                  "args" => {
                    "params" => [
                      {
                        "example" => "Ernest Dowson",
                        "kind" => "param",
                        "name" => "author",
                        "orig" => "author",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                      {
                        "example" => "text",
                        "kind" => "param",
                        "name" => "format",
                        "orig" => "format",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                      {
                        "example" => "author,title,linecount",
                        "kind" => "param",
                        "name" => "output_field",
                        "orig" => "output_field",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                    ],
                  },
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/author/{author}/{outputFields}.{format}",
                  "parts" => [
                    "author",
                    "{author}",
                    "{output_fields}_{format}",
                  ],
                  "rename" => {
                    "param" => {
                      "outputFields}.{format" => "output_fields}_{format",
                    },
                  },
                  "select" => {
                    "exist" => [
                      "author",
                      "format",
                      "output_field",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                },
                {
                  "args" => {
                    "params" => [
                      {
                        "example" => "Ernest Dowson",
                        "kind" => "param",
                        "name" => "author",
                        "orig" => "author",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                      {
                        "example" => "author,title,linecount",
                        "kind" => "param",
                        "name" => "output_field",
                        "orig" => "output_field",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                    ],
                  },
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/author/{author}/{outputFields}",
                  "parts" => [
                    "author",
                    "{author}",
                    "{output_field}",
                  ],
                  "rename" => {
                    "param" => {
                      "outputFields" => "output_field",
                    },
                  },
                  "select" => {
                    "exist" => [
                      "author",
                      "output_field",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                },
                {
                  "args" => {},
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/author",
                  "parts" => [
                    "author",
                  ],
                  "select" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body.authors`",
                  },
                },
              ],
            },
            "load" => {
              "input" => "data",
              "name" => "load",
              "points" => [
                {
                  "args" => {
                    "params" => [
                      {
                        "example" => "Ernest Dowson",
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "author",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                    ],
                  },
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/author/{author}",
                  "parts" => [
                    "author",
                    "{id}",
                  ],
                  "rename" => {
                    "param" => {
                      "author" => "id",
                    },
                  },
                  "select" => {
                    "exist" => [
                      "id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                },
              ],
            },
          },
          "relations" => {
            "ancestors" => [
              [
                "author",
              ],
            ],
          },
        },
        "authorab" => {
          "fields" => [
            {
              "name" => "author",
              "type" => "`$STRING`",
            },
            {
              "name" => "linecount",
              "type" => "`$INTEGER`",
            },
            {
              "name" => "lines",
              "type" => "`$ARRAY`",
            },
            {
              "name" => "title",
              "type" => "`$STRING`",
            },
          ],
          "name" => "authorab",
          "op" => {
            "list" => {
              "input" => "data",
              "name" => "list",
              "points" => [
                {
                  "args" => {
                    "params" => [
                      {
                        "example" => "Ernest Dowson",
                        "kind" => "param",
                        "name" => "author",
                        "orig" => "author",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                    ],
                  },
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/author/{author}:abs",
                  "parts" => [
                    "author",
                    "{author}:abs",
                  ],
                  "select" => {
                    "exist" => [
                      "author",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                },
              ],
            },
          },
          "relations" => {
            "ancestors" => [
              [
                "author",
              ],
            ],
          },
        },
        "combined_search" => {
          "fields" => [
            {
              "name" => "author",
              "type" => "`$STRING`",
            },
            {
              "name" => "linecount",
              "type" => "`$INTEGER`",
            },
            {
              "name" => "lines",
              "type" => "`$ARRAY`",
            },
            {
              "name" => "title",
              "type" => "`$STRING`",
            },
          ],
          "name" => "combined_search",
          "op" => {
            "list" => {
              "input" => "data",
              "name" => "list",
              "points" => [
                {
                  "args" => {
                    "params" => [
                      {
                        "example" => "author",
                        "kind" => "param",
                        "name" => "input_field1",
                        "orig" => "input_field1",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                      {
                        "example" => "linecount",
                        "kind" => "param",
                        "name" => "input_field2",
                        "orig" => "input_field2",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                      {
                        "example" => "Shakespeare",
                        "kind" => "param",
                        "name" => "search_term1",
                        "orig" => "search_term1",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                      {
                        "example" => "14",
                        "kind" => "param",
                        "name" => "search_term2",
                        "orig" => "search_term2",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                    ],
                  },
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/{inputField1},{inputField2}/{searchTerm1};{searchTerm2}",
                  "parts" => [
                    "{input_field1},{input_field2}",
                    "{search_term1};{search_term2}",
                  ],
                  "rename" => {
                    "param" => {
                      "inputField1},{inputField2" => "input_field1},{input_field2",
                      "searchTerm1};{searchTerm2" => "search_term1};{search_term2",
                    },
                  },
                  "select" => {
                    "exist" => [
                      "input_field1",
                      "input_field2",
                      "search_term1",
                      "search_term2",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                },
              ],
            },
          },
          "relations" => {
            "ancestors" => [],
          },
        },
        "combined_search_with_field" => {
          "fields" => [],
          "name" => "combined_search_with_field",
          "op" => {
            "list" => {
              "input" => "data",
              "name" => "list",
              "points" => [
                {
                  "args" => {
                    "params" => [
                      {
                        "example" => "author",
                        "kind" => "param",
                        "name" => "input_field1",
                        "orig" => "input_field1",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                      {
                        "example" => "linecount",
                        "kind" => "param",
                        "name" => "input_field2",
                        "orig" => "input_field2",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                      {
                        "example" => "lines",
                        "kind" => "param",
                        "name" => "output_field",
                        "orig" => "output_field",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                      {
                        "example" => "Shakespeare",
                        "kind" => "param",
                        "name" => "search_term1",
                        "orig" => "search_term1",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                      {
                        "example" => "14",
                        "kind" => "param",
                        "name" => "search_term2",
                        "orig" => "search_term2",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                    ],
                  },
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/{inputField1},{inputField2}/{searchTerm1};{searchTerm2}/{outputFields}",
                  "parts" => [
                    "{input_field1},{input_field2}",
                    "{search_term1};{search_term2}",
                    "{output_field}",
                  ],
                  "rename" => {
                    "param" => {
                      "inputField1},{inputField2" => "input_field1},{input_field2",
                      "outputFields" => "output_field",
                      "searchTerm1};{searchTerm2" => "search_term1};{search_term2",
                    },
                  },
                  "select" => {
                    "exist" => [
                      "input_field1",
                      "input_field2",
                      "output_field",
                      "search_term1",
                      "search_term2",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                },
              ],
            },
          },
          "relations" => {
            "ancestors" => [],
          },
        },
        "line" => {
          "fields" => [
            {
              "name" => "author",
              "type" => "`$STRING`",
            },
            {
              "name" => "linecount",
              "type" => "`$INTEGER`",
            },
            {
              "name" => "lines",
              "type" => "`$ARRAY`",
            },
            {
              "name" => "title",
              "type" => "`$STRING`",
            },
          ],
          "name" => "line",
          "op" => {
            "list" => {
              "input" => "data",
              "name" => "list",
              "points" => [
                {
                  "args" => {
                    "params" => [
                      {
                        "example" => "text",
                        "kind" => "param",
                        "name" => "format",
                        "orig" => "format",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                      {
                        "example" => "Latitudeless Place",
                        "kind" => "param",
                        "name" => "line",
                        "orig" => "line",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                      {
                        "example" => "author,title,linecount",
                        "kind" => "param",
                        "name" => "output_field",
                        "orig" => "output_field",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                    ],
                  },
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/lines/{lines}/{outputFields}.{format}",
                  "parts" => [
                    "lines",
                    "{line}",
                    "{output_fields}_{format}",
                  ],
                  "rename" => {
                    "param" => {
                      "lines" => "line",
                      "outputFields}.{format" => "output_fields}_{format",
                    },
                  },
                  "select" => {
                    "exist" => [
                      "format",
                      "line",
                      "output_field",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                },
                {
                  "args" => {
                    "params" => [
                      {
                        "example" => "Latitudeless Place",
                        "kind" => "param",
                        "name" => "line",
                        "orig" => "line",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                      {
                        "example" => "author,title,linecount",
                        "kind" => "param",
                        "name" => "output_field",
                        "orig" => "output_field",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                    ],
                  },
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/lines/{lines}/{outputFields}",
                  "parts" => [
                    "lines",
                    "{line}",
                    "{output_field}",
                  ],
                  "rename" => {
                    "param" => {
                      "lines" => "line",
                      "outputFields" => "output_field",
                    },
                  },
                  "select" => {
                    "exist" => [
                      "line",
                      "output_field",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                },
              ],
            },
            "load" => {
              "input" => "data",
              "name" => "load",
              "points" => [
                {
                  "args" => {
                    "params" => [
                      {
                        "example" => "Latitudeless Place",
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "line",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                    ],
                  },
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/lines/{lines}",
                  "parts" => [
                    "lines",
                    "{id}",
                  ],
                  "rename" => {
                    "param" => {
                      "lines" => "id",
                    },
                  },
                  "select" => {
                    "exist" => [
                      "id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                },
              ],
            },
          },
          "relations" => {
            "ancestors" => [
              [
                "line",
              ],
            ],
          },
        },
        "linecount" => {
          "fields" => [
            {
              "name" => "author",
              "type" => "`$STRING`",
            },
            {
              "name" => "linecount",
              "type" => "`$INTEGER`",
            },
            {
              "name" => "lines",
              "type" => "`$ARRAY`",
            },
            {
              "name" => "title",
              "type" => "`$STRING`",
            },
          ],
          "name" => "linecount",
          "op" => {
            "list" => {
              "input" => "data",
              "name" => "list",
              "points" => [
                {
                  "args" => {
                    "params" => [
                      {
                        "example" => "text",
                        "kind" => "param",
                        "name" => "format",
                        "orig" => "format",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                      {
                        "example" => 14,
                        "kind" => "param",
                        "name" => "linecount",
                        "orig" => "linecount",
                        "reqd" => true,
                        "type" => "`$INTEGER`",
                      },
                      {
                        "example" => "author,title",
                        "kind" => "param",
                        "name" => "output_field",
                        "orig" => "output_field",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                    ],
                  },
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/linecount/{linecount}/{outputFields}.{format}",
                  "parts" => [
                    "linecount",
                    "{linecount}",
                    "{output_fields}_{format}",
                  ],
                  "rename" => {
                    "param" => {
                      "outputFields}.{format" => "output_fields}_{format",
                    },
                  },
                  "select" => {
                    "exist" => [
                      "format",
                      "linecount",
                      "output_field",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                },
                {
                  "args" => {
                    "params" => [
                      {
                        "example" => 14,
                        "kind" => "param",
                        "name" => "linecount",
                        "orig" => "linecount",
                        "reqd" => true,
                        "type" => "`$INTEGER`",
                      },
                      {
                        "example" => "author,title",
                        "kind" => "param",
                        "name" => "output_field",
                        "orig" => "output_field",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                    ],
                  },
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/linecount/{linecount}/{outputFields}",
                  "parts" => [
                    "linecount",
                    "{linecount}",
                    "{output_field}",
                  ],
                  "rename" => {
                    "param" => {
                      "outputFields" => "output_field",
                    },
                  },
                  "select" => {
                    "exist" => [
                      "linecount",
                      "output_field",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                },
              ],
            },
            "load" => {
              "input" => "data",
              "name" => "load",
              "points" => [
                {
                  "args" => {
                    "params" => [
                      {
                        "example" => 14,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "linecount",
                        "reqd" => true,
                        "type" => "`$INTEGER`",
                      },
                    ],
                  },
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/linecount/{linecount}",
                  "parts" => [
                    "linecount",
                    "{id}",
                  ],
                  "rename" => {
                    "param" => {
                      "linecount" => "id",
                    },
                  },
                  "select" => {
                    "exist" => [
                      "id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                },
              ],
            },
          },
          "relations" => {
            "ancestors" => [
              [
                "linecount",
              ],
            ],
          },
        },
        "poemcount" => {
          "fields" => [
            {
              "name" => "author",
              "type" => "`$STRING`",
            },
            {
              "name" => "linecount",
              "type" => "`$INTEGER`",
            },
            {
              "name" => "lines",
              "type" => "`$ARRAY`",
            },
            {
              "name" => "title",
              "type" => "`$STRING`",
            },
          ],
          "name" => "poemcount",
          "op" => {
            "load" => {
              "input" => "data",
              "name" => "load",
              "points" => [
                {
                  "args" => {
                    "params" => [
                      {
                        "example" => 10,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "count",
                        "reqd" => true,
                        "type" => "`$INTEGER`",
                      },
                    ],
                  },
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/poemcount/{count}",
                  "parts" => [
                    "poemcount",
                    "{id}",
                  ],
                  "rename" => {
                    "param" => {
                      "count" => "id",
                    },
                  },
                  "select" => {
                    "exist" => [
                      "id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                },
              ],
            },
          },
          "relations" => {
            "ancestors" => [],
          },
        },
        "random" => {
          "fields" => [
            {
              "name" => "author",
              "type" => "`$STRING`",
            },
            {
              "name" => "linecount",
              "type" => "`$INTEGER`",
            },
            {
              "name" => "lines",
              "type" => "`$ARRAY`",
            },
            {
              "name" => "title",
              "type" => "`$STRING`",
            },
          ],
          "name" => "random",
          "op" => {
            "list" => {
              "input" => "data",
              "name" => "list",
              "points" => [
                {
                  "args" => {
                    "params" => [
                      {
                        "example" => 5,
                        "kind" => "param",
                        "name" => "count",
                        "orig" => "count",
                        "reqd" => true,
                        "type" => "`$INTEGER`",
                      },
                      {
                        "example" => "author,title",
                        "kind" => "param",
                        "name" => "output_field",
                        "orig" => "output_field",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                    ],
                  },
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/random/{count}/{outputFields}",
                  "parts" => [
                    "random",
                    "{count}",
                    "{output_field}",
                  ],
                  "rename" => {
                    "param" => {
                      "outputFields" => "output_field",
                    },
                  },
                  "select" => {
                    "exist" => [
                      "count",
                      "output_field",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                },
              ],
            },
            "load" => {
              "input" => "data",
              "name" => "load",
              "points" => [
                {
                  "args" => {
                    "params" => [
                      {
                        "example" => 5,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "count",
                        "reqd" => true,
                        "type" => "`$INTEGER`",
                      },
                    ],
                  },
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/random/{count}",
                  "parts" => [
                    "random",
                    "{id}",
                  ],
                  "rename" => {
                    "param" => {
                      "count" => "id",
                    },
                  },
                  "select" => {
                    "exist" => [
                      "id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                },
              ],
            },
          },
          "relations" => {
            "ancestors" => [
              [
                "random",
              ],
            ],
          },
        },
        "title" => {
          "fields" => [
            {
              "name" => "author",
              "type" => "`$STRING`",
            },
            {
              "name" => "linecount",
              "type" => "`$INTEGER`",
            },
            {
              "name" => "lines",
              "type" => "`$ARRAY`",
            },
            {
              "name" => "title",
              "type" => "`$STRING`",
            },
            {
              "name" => "titles",
              "type" => "`$ARRAY`",
            },
          ],
          "name" => "title",
          "op" => {
            "list" => {
              "input" => "data",
              "name" => "list",
              "points" => [
                {
                  "args" => {
                    "params" => [
                      {
                        "example" => "text",
                        "kind" => "param",
                        "name" => "format",
                        "orig" => "format",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                      {
                        "example" => "title,lines",
                        "kind" => "param",
                        "name" => "output_field",
                        "orig" => "output_field",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                      {
                        "example" => "Ozymandias",
                        "kind" => "param",
                        "name" => "title",
                        "orig" => "title",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                    ],
                  },
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/title/{title}/{outputFields}.{format}",
                  "parts" => [
                    "title",
                    "{title}",
                    "{output_fields}_{format}",
                  ],
                  "rename" => {
                    "param" => {
                      "outputFields}.{format" => "output_fields}_{format",
                    },
                  },
                  "select" => {
                    "exist" => [
                      "format",
                      "output_field",
                      "title",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                },
                {
                  "args" => {
                    "params" => [
                      {
                        "example" => "author,title,linecount",
                        "kind" => "param",
                        "name" => "output_field",
                        "orig" => "output_field",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                      {
                        "example" => "Ozymandias",
                        "kind" => "param",
                        "name" => "title",
                        "orig" => "title",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                    ],
                  },
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/title/{title}/{outputFields}",
                  "parts" => [
                    "title",
                    "{title}",
                    "{output_field}",
                  ],
                  "rename" => {
                    "param" => {
                      "outputFields" => "output_field",
                    },
                  },
                  "select" => {
                    "exist" => [
                      "output_field",
                      "title",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                },
                {
                  "args" => {},
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/title",
                  "parts" => [
                    "title",
                  ],
                  "select" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body.titles`",
                  },
                },
              ],
            },
            "load" => {
              "input" => "data",
              "name" => "load",
              "points" => [
                {
                  "args" => {
                    "params" => [
                      {
                        "example" => "Ozymandias",
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "title",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                    ],
                  },
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/title/{title}",
                  "parts" => [
                    "title",
                    "{id}",
                  ],
                  "rename" => {
                    "param" => {
                      "title" => "id",
                    },
                  },
                  "select" => {
                    "exist" => [
                      "id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                },
              ],
            },
          },
          "relations" => {
            "ancestors" => [
              [
                "title",
              ],
            ],
          },
        },
        "titleab" => {
          "fields" => [
            {
              "name" => "author",
              "type" => "`$STRING`",
            },
            {
              "name" => "linecount",
              "type" => "`$INTEGER`",
            },
            {
              "name" => "lines",
              "type" => "`$ARRAY`",
            },
            {
              "name" => "title",
              "type" => "`$STRING`",
            },
          ],
          "name" => "titleab",
          "op" => {
            "list" => {
              "input" => "data",
              "name" => "list",
              "points" => [
                {
                  "args" => {
                    "params" => [
                      {
                        "example" => "Ozymandias",
                        "kind" => "param",
                        "name" => "title",
                        "orig" => "title",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                    ],
                  },
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/title/{title}:abs",
                  "parts" => [
                    "title",
                    "{title}:abs",
                  ],
                  "select" => {
                    "exist" => [
                      "title",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                },
              ],
            },
          },
          "relations" => {
            "ancestors" => [
              [
                "title",
              ],
            ],
          },
        },
      },
    }
  end


  def self.make_feature(name)
    require_relative 'features'
    PoetrydbFeatures.make_feature(name)
  end
end
