<?php
declare(strict_types=1);

// Poetrydb SDK configuration

class PoetrydbConfig
{
    public static function make_config(): array
    {
        return [
            "main" => [
                "name" => "Poetrydb",
            ],
            "feature" => [
                "test" => [
          'options' => [
            'active' => false,
          ],
        ],
            ],
            "options" => [
                "base" => "https://poetrydb.org",
                "auth" => [
                    "prefix" => "Bearer",
                ],
                "headers" => [
          'content-type' => 'application/json',
        ],
                "entity" => [
                    "author" => [],
                    "authorab" => [],
                    "combined_search" => [],
                    "combined_search_with_field" => [],
                    "line" => [],
                    "linecount" => [],
                    "poemcount" => [],
                    "random" => [],
                    "title" => [],
                    "titleab" => [],
                ],
            ],
            "entity" => [
        'author' => [
          'fields' => [
            [
              'name' => 'author',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 0,
            ],
            [
              'name' => 'line',
              'req' => false,
              'type' => '`$ARRAY`',
              'active' => true,
              'index$' => 1,
            ],
            [
              'name' => 'linecount',
              'req' => false,
              'type' => '`$INTEGER`',
              'active' => true,
              'index$' => 2,
            ],
            [
              'name' => 'title',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 3,
            ],
          ],
          'name' => 'author',
          'op' => [
            'list' => [
              'name' => 'list',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'example' => 'Ernest Dowson',
                        'kind' => 'param',
                        'name' => 'author',
                        'orig' => 'author',
                        'reqd' => true,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                      [
                        'example' => 'text',
                        'kind' => 'param',
                        'name' => 'format',
                        'orig' => 'format',
                        'reqd' => true,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                      [
                        'example' => 'author,title,linecount',
                        'kind' => 'param',
                        'name' => 'output_field',
                        'orig' => 'output_field',
                        'reqd' => true,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                    ],
                  ],
                  'method' => 'GET',
                  'orig' => '/author/{author}/{outputFields}.{format}',
                  'parts' => [
                    'author',
                    '{author}',
                    '{output_fields}_{format}',
                  ],
                  'rename' => [
                    'param' => [
                      'outputFields}.{format' => 'output_fields}_{format',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'author',
                      'format',
                      'output_field',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'active' => true,
                  'index$' => 0,
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'example' => 'Ernest Dowson',
                        'kind' => 'param',
                        'name' => 'author',
                        'orig' => 'author',
                        'reqd' => true,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                      [
                        'example' => 'author,title,linecount',
                        'kind' => 'param',
                        'name' => 'output_field',
                        'orig' => 'output_field',
                        'reqd' => true,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                    ],
                  ],
                  'method' => 'GET',
                  'orig' => '/author/{author}/{outputFields}',
                  'parts' => [
                    'author',
                    '{author}',
                    '{output_field}',
                  ],
                  'rename' => [
                    'param' => [
                      'outputFields' => 'output_field',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'author',
                      'output_field',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'active' => true,
                  'index$' => 1,
                ],
                [
                  'method' => 'GET',
                  'orig' => '/author',
                  'parts' => [
                    'author',
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'active' => true,
                  'args' => [],
                  'select' => [],
                  'index$' => 2,
                ],
              ],
              'input' => 'data',
              'key$' => 'list',
            ],
            'load' => [
              'name' => 'load',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'example' => 'Ernest Dowson',
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'author',
                        'reqd' => true,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                    ],
                  ],
                  'method' => 'GET',
                  'orig' => '/author/{author}',
                  'parts' => [
                    'author',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'author' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'active' => true,
                  'index$' => 0,
                ],
              ],
              'input' => 'data',
              'key$' => 'load',
            ],
          ],
          'relations' => [
            'ancestors' => [
              [
                'author',
              ],
            ],
          ],
        ],
        'authorab' => [
          'fields' => [
            [
              'name' => 'author',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 0,
            ],
            [
              'name' => 'line',
              'req' => false,
              'type' => '`$ARRAY`',
              'active' => true,
              'index$' => 1,
            ],
            [
              'name' => 'linecount',
              'req' => false,
              'type' => '`$INTEGER`',
              'active' => true,
              'index$' => 2,
            ],
            [
              'name' => 'title',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 3,
            ],
          ],
          'name' => 'authorab',
          'op' => [
            'list' => [
              'name' => 'list',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'example' => 'Ernest Dowson',
                        'kind' => 'param',
                        'name' => 'author',
                        'orig' => 'author',
                        'reqd' => true,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                    ],
                  ],
                  'method' => 'GET',
                  'orig' => '/author/{author}:abs',
                  'parts' => [
                    'author',
                    '{author}:abs',
                  ],
                  'select' => [
                    'exist' => [
                      'author',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'active' => true,
                  'index$' => 0,
                ],
              ],
              'input' => 'data',
              'key$' => 'list',
            ],
          ],
          'relations' => [
            'ancestors' => [
              [
                'author',
              ],
            ],
          ],
        ],
        'combined_search' => [
          'fields' => [
            [
              'name' => 'author',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 0,
            ],
            [
              'name' => 'line',
              'req' => false,
              'type' => '`$ARRAY`',
              'active' => true,
              'index$' => 1,
            ],
            [
              'name' => 'linecount',
              'req' => false,
              'type' => '`$INTEGER`',
              'active' => true,
              'index$' => 2,
            ],
            [
              'name' => 'title',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 3,
            ],
          ],
          'name' => 'combined_search',
          'op' => [
            'list' => [
              'name' => 'list',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'example' => 'author',
                        'kind' => 'param',
                        'name' => 'input_field1',
                        'orig' => 'input_field1',
                        'reqd' => true,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                      [
                        'example' => 'linecount',
                        'kind' => 'param',
                        'name' => 'input_field2',
                        'orig' => 'input_field2',
                        'reqd' => true,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                      [
                        'example' => 'Shakespeare',
                        'kind' => 'param',
                        'name' => 'search_term1',
                        'orig' => 'search_term1',
                        'reqd' => true,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                      [
                        'example' => '14',
                        'kind' => 'param',
                        'name' => 'search_term2',
                        'orig' => 'search_term2',
                        'reqd' => true,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                    ],
                  ],
                  'method' => 'GET',
                  'orig' => '/{inputField1},{inputField2}/{searchTerm1};{searchTerm2}',
                  'parts' => [
                    '{input_field1},{input_field2}',
                    '{search_term1};{search_term2}',
                  ],
                  'rename' => [
                    'param' => [
                      'inputField1},{inputField2' => 'input_field1},{input_field2',
                      'searchTerm1};{searchTerm2' => 'search_term1};{search_term2',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'input_field1',
                      'input_field2',
                      'search_term1',
                      'search_term2',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'active' => true,
                  'index$' => 0,
                ],
              ],
              'input' => 'data',
              'key$' => 'list',
            ],
          ],
          'relations' => [
            'ancestors' => [],
          ],
        ],
        'combined_search_with_field' => [
          'fields' => [],
          'name' => 'combined_search_with_field',
          'op' => [
            'list' => [
              'name' => 'list',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'example' => 'author',
                        'kind' => 'param',
                        'name' => 'input_field1',
                        'orig' => 'input_field1',
                        'reqd' => true,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                      [
                        'example' => 'linecount',
                        'kind' => 'param',
                        'name' => 'input_field2',
                        'orig' => 'input_field2',
                        'reqd' => true,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                      [
                        'example' => 'lines',
                        'kind' => 'param',
                        'name' => 'output_field',
                        'orig' => 'output_field',
                        'reqd' => true,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                      [
                        'example' => 'Shakespeare',
                        'kind' => 'param',
                        'name' => 'search_term1',
                        'orig' => 'search_term1',
                        'reqd' => true,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                      [
                        'example' => '14',
                        'kind' => 'param',
                        'name' => 'search_term2',
                        'orig' => 'search_term2',
                        'reqd' => true,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                    ],
                  ],
                  'method' => 'GET',
                  'orig' => '/{inputField1},{inputField2}/{searchTerm1};{searchTerm2}/{outputFields}',
                  'parts' => [
                    '{input_field1},{input_field2}',
                    '{search_term1};{search_term2}',
                    '{output_field}',
                  ],
                  'rename' => [
                    'param' => [
                      'inputField1},{inputField2' => 'input_field1},{input_field2',
                      'outputFields' => 'output_field',
                      'searchTerm1};{searchTerm2' => 'search_term1};{search_term2',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'input_field1',
                      'input_field2',
                      'output_field',
                      'search_term1',
                      'search_term2',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'active' => true,
                  'index$' => 0,
                ],
              ],
              'input' => 'data',
              'key$' => 'list',
            ],
          ],
          'relations' => [
            'ancestors' => [],
          ],
        ],
        'line' => [
          'fields' => [
            [
              'name' => 'author',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 0,
            ],
            [
              'name' => 'line',
              'req' => false,
              'type' => '`$ARRAY`',
              'active' => true,
              'index$' => 1,
            ],
            [
              'name' => 'linecount',
              'req' => false,
              'type' => '`$INTEGER`',
              'active' => true,
              'index$' => 2,
            ],
            [
              'name' => 'title',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 3,
            ],
          ],
          'name' => 'line',
          'op' => [
            'list' => [
              'name' => 'list',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'example' => 'text',
                        'kind' => 'param',
                        'name' => 'format',
                        'orig' => 'format',
                        'reqd' => true,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                      [
                        'example' => 'Latitudeless Place',
                        'kind' => 'param',
                        'name' => 'line',
                        'orig' => 'line',
                        'reqd' => true,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                      [
                        'example' => 'author,title,linecount',
                        'kind' => 'param',
                        'name' => 'output_field',
                        'orig' => 'output_field',
                        'reqd' => true,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                    ],
                  ],
                  'method' => 'GET',
                  'orig' => '/lines/{lines}/{outputFields}.{format}',
                  'parts' => [
                    'lines',
                    '{line}',
                    '{output_fields}_{format}',
                  ],
                  'rename' => [
                    'param' => [
                      'lines' => 'line',
                      'outputFields}.{format' => 'output_fields}_{format',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'format',
                      'line',
                      'output_field',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'active' => true,
                  'index$' => 0,
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'example' => 'Latitudeless Place',
                        'kind' => 'param',
                        'name' => 'line',
                        'orig' => 'line',
                        'reqd' => true,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                      [
                        'example' => 'author,title,linecount',
                        'kind' => 'param',
                        'name' => 'output_field',
                        'orig' => 'output_field',
                        'reqd' => true,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                    ],
                  ],
                  'method' => 'GET',
                  'orig' => '/lines/{lines}/{outputFields}',
                  'parts' => [
                    'lines',
                    '{line}',
                    '{output_field}',
                  ],
                  'rename' => [
                    'param' => [
                      'lines' => 'line',
                      'outputFields' => 'output_field',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'line',
                      'output_field',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'active' => true,
                  'index$' => 1,
                ],
              ],
              'input' => 'data',
              'key$' => 'list',
            ],
            'load' => [
              'name' => 'load',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'example' => 'Latitudeless Place',
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'line',
                        'reqd' => true,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                    ],
                  ],
                  'method' => 'GET',
                  'orig' => '/lines/{lines}',
                  'parts' => [
                    'lines',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'lines' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'active' => true,
                  'index$' => 0,
                ],
              ],
              'input' => 'data',
              'key$' => 'load',
            ],
          ],
          'relations' => [
            'ancestors' => [
              [
                'line',
              ],
            ],
          ],
        ],
        'linecount' => [
          'fields' => [
            [
              'name' => 'author',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 0,
            ],
            [
              'name' => 'line',
              'req' => false,
              'type' => '`$ARRAY`',
              'active' => true,
              'index$' => 1,
            ],
            [
              'name' => 'linecount',
              'req' => false,
              'type' => '`$INTEGER`',
              'active' => true,
              'index$' => 2,
            ],
            [
              'name' => 'title',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 3,
            ],
          ],
          'name' => 'linecount',
          'op' => [
            'list' => [
              'name' => 'list',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'example' => 'text',
                        'kind' => 'param',
                        'name' => 'format',
                        'orig' => 'format',
                        'reqd' => true,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                      [
                        'example' => 14,
                        'kind' => 'param',
                        'name' => 'linecount',
                        'orig' => 'linecount',
                        'reqd' => true,
                        'type' => '`$INTEGER`',
                        'active' => true,
                      ],
                      [
                        'example' => 'author,title',
                        'kind' => 'param',
                        'name' => 'output_field',
                        'orig' => 'output_field',
                        'reqd' => true,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                    ],
                  ],
                  'method' => 'GET',
                  'orig' => '/linecount/{linecount}/{outputFields}.{format}',
                  'parts' => [
                    'linecount',
                    '{linecount}',
                    '{output_fields}_{format}',
                  ],
                  'rename' => [
                    'param' => [
                      'outputFields}.{format' => 'output_fields}_{format',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'format',
                      'linecount',
                      'output_field',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'active' => true,
                  'index$' => 0,
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'example' => 14,
                        'kind' => 'param',
                        'name' => 'linecount',
                        'orig' => 'linecount',
                        'reqd' => true,
                        'type' => '`$INTEGER`',
                        'active' => true,
                      ],
                      [
                        'example' => 'author,title',
                        'kind' => 'param',
                        'name' => 'output_field',
                        'orig' => 'output_field',
                        'reqd' => true,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                    ],
                  ],
                  'method' => 'GET',
                  'orig' => '/linecount/{linecount}/{outputFields}',
                  'parts' => [
                    'linecount',
                    '{linecount}',
                    '{output_field}',
                  ],
                  'rename' => [
                    'param' => [
                      'outputFields' => 'output_field',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'linecount',
                      'output_field',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'active' => true,
                  'index$' => 1,
                ],
              ],
              'input' => 'data',
              'key$' => 'list',
            ],
            'load' => [
              'name' => 'load',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'example' => 14,
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'linecount',
                        'reqd' => true,
                        'type' => '`$INTEGER`',
                        'active' => true,
                      ],
                    ],
                  ],
                  'method' => 'GET',
                  'orig' => '/linecount/{linecount}',
                  'parts' => [
                    'linecount',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'linecount' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'active' => true,
                  'index$' => 0,
                ],
              ],
              'input' => 'data',
              'key$' => 'load',
            ],
          ],
          'relations' => [
            'ancestors' => [
              [
                'linecount',
              ],
            ],
          ],
        ],
        'poemcount' => [
          'fields' => [
            [
              'name' => 'author',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 0,
            ],
            [
              'name' => 'line',
              'req' => false,
              'type' => '`$ARRAY`',
              'active' => true,
              'index$' => 1,
            ],
            [
              'name' => 'linecount',
              'req' => false,
              'type' => '`$INTEGER`',
              'active' => true,
              'index$' => 2,
            ],
            [
              'name' => 'title',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 3,
            ],
          ],
          'name' => 'poemcount',
          'op' => [
            'load' => [
              'name' => 'load',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'example' => 10,
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'count',
                        'reqd' => true,
                        'type' => '`$INTEGER`',
                        'active' => true,
                      ],
                    ],
                  ],
                  'method' => 'GET',
                  'orig' => '/poemcount/{count}',
                  'parts' => [
                    'poemcount',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'count' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'active' => true,
                  'index$' => 0,
                ],
              ],
              'input' => 'data',
              'key$' => 'load',
            ],
          ],
          'relations' => [
            'ancestors' => [],
          ],
        ],
        'random' => [
          'fields' => [
            [
              'name' => 'author',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 0,
            ],
            [
              'name' => 'line',
              'req' => false,
              'type' => '`$ARRAY`',
              'active' => true,
              'index$' => 1,
            ],
            [
              'name' => 'linecount',
              'req' => false,
              'type' => '`$INTEGER`',
              'active' => true,
              'index$' => 2,
            ],
            [
              'name' => 'title',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 3,
            ],
          ],
          'name' => 'random',
          'op' => [
            'list' => [
              'name' => 'list',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'example' => 5,
                        'kind' => 'param',
                        'name' => 'count',
                        'orig' => 'count',
                        'reqd' => true,
                        'type' => '`$INTEGER`',
                        'active' => true,
                      ],
                      [
                        'example' => 'author,title',
                        'kind' => 'param',
                        'name' => 'output_field',
                        'orig' => 'output_field',
                        'reqd' => true,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                    ],
                  ],
                  'method' => 'GET',
                  'orig' => '/random/{count}/{outputFields}',
                  'parts' => [
                    'random',
                    '{count}',
                    '{output_field}',
                  ],
                  'rename' => [
                    'param' => [
                      'outputFields' => 'output_field',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'count',
                      'output_field',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'active' => true,
                  'index$' => 0,
                ],
              ],
              'input' => 'data',
              'key$' => 'list',
            ],
            'load' => [
              'name' => 'load',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'example' => 5,
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'count',
                        'reqd' => true,
                        'type' => '`$INTEGER`',
                        'active' => true,
                      ],
                    ],
                  ],
                  'method' => 'GET',
                  'orig' => '/random/{count}',
                  'parts' => [
                    'random',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'count' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'active' => true,
                  'index$' => 0,
                ],
              ],
              'input' => 'data',
              'key$' => 'load',
            ],
          ],
          'relations' => [
            'ancestors' => [
              [
                'random',
              ],
            ],
          ],
        ],
        'title' => [
          'fields' => [
            [
              'name' => 'author',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 0,
            ],
            [
              'name' => 'line',
              'req' => false,
              'type' => '`$ARRAY`',
              'active' => true,
              'index$' => 1,
            ],
            [
              'name' => 'linecount',
              'req' => false,
              'type' => '`$INTEGER`',
              'active' => true,
              'index$' => 2,
            ],
            [
              'name' => 'title',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 3,
            ],
          ],
          'name' => 'title',
          'op' => [
            'list' => [
              'name' => 'list',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'example' => 'text',
                        'kind' => 'param',
                        'name' => 'format',
                        'orig' => 'format',
                        'reqd' => true,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                      [
                        'example' => 'title,lines',
                        'kind' => 'param',
                        'name' => 'output_field',
                        'orig' => 'output_field',
                        'reqd' => true,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                      [
                        'example' => 'Ozymandias',
                        'kind' => 'param',
                        'name' => 'title',
                        'orig' => 'title',
                        'reqd' => true,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                    ],
                  ],
                  'method' => 'GET',
                  'orig' => '/title/{title}/{outputFields}.{format}',
                  'parts' => [
                    'title',
                    '{title}',
                    '{output_fields}_{format}',
                  ],
                  'rename' => [
                    'param' => [
                      'outputFields}.{format' => 'output_fields}_{format',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'format',
                      'output_field',
                      'title',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'active' => true,
                  'index$' => 0,
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'example' => 'author,title,linecount',
                        'kind' => 'param',
                        'name' => 'output_field',
                        'orig' => 'output_field',
                        'reqd' => true,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                      [
                        'example' => 'Ozymandias',
                        'kind' => 'param',
                        'name' => 'title',
                        'orig' => 'title',
                        'reqd' => true,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                    ],
                  ],
                  'method' => 'GET',
                  'orig' => '/title/{title}/{outputFields}',
                  'parts' => [
                    'title',
                    '{title}',
                    '{output_field}',
                  ],
                  'rename' => [
                    'param' => [
                      'outputFields' => 'output_field',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'output_field',
                      'title',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'active' => true,
                  'index$' => 1,
                ],
                [
                  'method' => 'GET',
                  'orig' => '/title',
                  'parts' => [
                    'title',
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'active' => true,
                  'args' => [],
                  'select' => [],
                  'index$' => 2,
                ],
              ],
              'input' => 'data',
              'key$' => 'list',
            ],
            'load' => [
              'name' => 'load',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'example' => 'Ozymandias',
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'title',
                        'reqd' => true,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                    ],
                  ],
                  'method' => 'GET',
                  'orig' => '/title/{title}',
                  'parts' => [
                    'title',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'title' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'active' => true,
                  'index$' => 0,
                ],
              ],
              'input' => 'data',
              'key$' => 'load',
            ],
          ],
          'relations' => [
            'ancestors' => [
              [
                'title',
              ],
            ],
          ],
        ],
        'titleab' => [
          'fields' => [
            [
              'name' => 'author',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 0,
            ],
            [
              'name' => 'line',
              'req' => false,
              'type' => '`$ARRAY`',
              'active' => true,
              'index$' => 1,
            ],
            [
              'name' => 'linecount',
              'req' => false,
              'type' => '`$INTEGER`',
              'active' => true,
              'index$' => 2,
            ],
            [
              'name' => 'title',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 3,
            ],
          ],
          'name' => 'titleab',
          'op' => [
            'list' => [
              'name' => 'list',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'example' => 'Ozymandias',
                        'kind' => 'param',
                        'name' => 'title',
                        'orig' => 'title',
                        'reqd' => true,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                    ],
                  ],
                  'method' => 'GET',
                  'orig' => '/title/{title}:abs',
                  'parts' => [
                    'title',
                    '{title}:abs',
                  ],
                  'select' => [
                    'exist' => [
                      'title',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'active' => true,
                  'index$' => 0,
                ],
              ],
              'input' => 'data',
              'key$' => 'list',
            ],
          ],
          'relations' => [
            'ancestors' => [
              [
                'title',
              ],
            ],
          ],
        ],
      ],
        ];
    }


    public static function make_feature(string $name)
    {
        require_once __DIR__ . '/features.php';
        return PoetrydbFeatures::make_feature($name);
    }
}
