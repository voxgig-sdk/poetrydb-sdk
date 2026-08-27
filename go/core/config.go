package core

import (
	"sync"
)

// MakeConfig builds a fresh, fully materialised config map. Every call
// rebuilds the whole structure, so prefer SharedConfig unless you need a
// private copy you intend to mutate.
func MakeConfig() map[string]any {
	return map[string]any{
		"main": map[string]any{
			"name": "Poetrydb",
			"slug": "poetrydb",
			"version": "0.0.1",
			"target": "go",
		},
		"feature": map[string]any{
			"test": map[string]any{
				"options": map[string]any{
					"active": false,
				},
				"transport": "base",
			},
		},
		"options": map[string]any{
			"base": "https://poetrydb.org",
			"headers": map[string]any{
				"content-type": "application/json",
			},
			"entity": map[string]any{
				"author": map[string]any{},
				"authorab": map[string]any{},
				"combined_search": map[string]any{},
				"combined_search_with_field": map[string]any{},
				"line": map[string]any{},
				"linecount": map[string]any{},
				"poemcount": map[string]any{},
				"random": map[string]any{},
				"title": map[string]any{},
				"titleab": map[string]any{},
			},
		},
		"entity": map[string]any{
			"author": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "author",
						"short": "The author of the poem",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "authors",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "linecount",
						"short": "The number of lines in the poem (including section headings, excluding empty lines)",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "lines",
						"short": "The lines of the poem",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "title",
						"short": "The title of the poem",
						"type": "`$STRING`",
					},
				},
				"name": "author",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"example": "Ernest Dowson",
											"kind": "param",
											"name": "author",
											"orig": "author",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"example": "text",
											"kind": "param",
											"name": "format",
											"orig": "format",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"example": "author,title,linecount",
											"kind": "param",
											"name": "output_field",
											"orig": "output_field",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/author/{author}/{outputFields}.{format}",
								"parts": []any{
									"author",
									"{author}",
									"{output_fields}_{format}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"outputFields}.{format": "output_fields}_{format",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"author",
										"format",
										"output_field",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"example": "Ernest Dowson",
											"kind": "param",
											"name": "author",
											"orig": "author",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"example": "author,title,linecount",
											"kind": "param",
											"name": "output_field",
											"orig": "output_field",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/author/{author}/{outputFields}",
								"parts": []any{
									"author",
									"{author}",
									"{output_field}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"outputFields": "output_field",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"author",
										"output_field",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/author",
								"parts": []any{
									"author",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.authors`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"example": "Ernest Dowson",
											"kind": "param",
											"name": "id",
											"orig": "author",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/author/{author}",
								"parts": []any{
									"author",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"author": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"author",
						},
					},
				},
			},
			"authorab": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "author",
						"short": "The author of the poem",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "linecount",
						"short": "The number of lines in the poem (including section headings, excluding empty lines)",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "lines",
						"short": "The lines of the poem",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "title",
						"short": "The title of the poem",
						"type": "`$STRING`",
					},
				},
				"name": "authorab",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"example": "Ernest Dowson",
											"kind": "param",
											"name": "author",
											"orig": "author",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/author/{author}:abs",
								"parts": []any{
									"author",
									"{author}:abs",
								},
								"select": map[string]any{
									"exist": []any{
										"author",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"author",
						},
					},
				},
			},
			"combined_search": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "author",
						"short": "The author of the poem",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "linecount",
						"short": "The number of lines in the poem (including section headings, excluding empty lines)",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "lines",
						"short": "The lines of the poem",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "title",
						"short": "The title of the poem",
						"type": "`$STRING`",
					},
				},
				"name": "combined_search",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"example": "author",
											"kind": "param",
											"name": "input_field1",
											"orig": "input_field1",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"example": "linecount",
											"kind": "param",
											"name": "input_field2",
											"orig": "input_field2",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"example": "Shakespeare",
											"kind": "param",
											"name": "search_term1",
											"orig": "search_term1",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"example": "14",
											"kind": "param",
											"name": "search_term2",
											"orig": "search_term2",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/{inputField1},{inputField2}/{searchTerm1};{searchTerm2}",
								"parts": []any{
									"{input_field1},{input_field2}",
									"{search_term1};{search_term2}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"inputField1},{inputField2": "input_field1},{input_field2",
										"searchTerm1};{searchTerm2": "search_term1};{search_term2",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"input_field1",
										"input_field2",
										"search_term1",
										"search_term2",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"combined_search_with_field": map[string]any{
				"fields": []any{},
				"name": "combined_search_with_field",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"example": "author",
											"kind": "param",
											"name": "input_field1",
											"orig": "input_field1",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"example": "linecount",
											"kind": "param",
											"name": "input_field2",
											"orig": "input_field2",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"example": "lines",
											"kind": "param",
											"name": "output_field",
											"orig": "output_field",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"example": "Shakespeare",
											"kind": "param",
											"name": "search_term1",
											"orig": "search_term1",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"example": "14",
											"kind": "param",
											"name": "search_term2",
											"orig": "search_term2",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/{inputField1},{inputField2}/{searchTerm1};{searchTerm2}/{outputFields}",
								"parts": []any{
									"{input_field1},{input_field2}",
									"{search_term1};{search_term2}",
									"{output_field}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"inputField1},{inputField2": "input_field1},{input_field2",
										"outputFields": "output_field",
										"searchTerm1};{searchTerm2": "search_term1};{search_term2",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"input_field1",
										"input_field2",
										"output_field",
										"search_term1",
										"search_term2",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"line": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "author",
						"short": "The author of the poem",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "linecount",
						"short": "The number of lines in the poem (including section headings, excluding empty lines)",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "lines",
						"short": "The lines of the poem",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "title",
						"short": "The title of the poem",
						"type": "`$STRING`",
					},
				},
				"name": "line",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"example": "text",
											"kind": "param",
											"name": "format",
											"orig": "format",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"example": "Latitudeless Place",
											"kind": "param",
											"name": "line",
											"orig": "line",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"example": "author,title,linecount",
											"kind": "param",
											"name": "output_field",
											"orig": "output_field",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/lines/{lines}/{outputFields}.{format}",
								"parts": []any{
									"lines",
									"{line}",
									"{output_fields}_{format}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"lines": "line",
										"outputFields}.{format": "output_fields}_{format",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"format",
										"line",
										"output_field",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"example": "Latitudeless Place",
											"kind": "param",
											"name": "line",
											"orig": "line",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"example": "author,title,linecount",
											"kind": "param",
											"name": "output_field",
											"orig": "output_field",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/lines/{lines}/{outputFields}",
								"parts": []any{
									"lines",
									"{line}",
									"{output_field}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"lines": "line",
										"outputFields": "output_field",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"line",
										"output_field",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"example": "Latitudeless Place",
											"kind": "param",
											"name": "id",
											"orig": "line",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/lines/{lines}",
								"parts": []any{
									"lines",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"lines": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"line",
						},
					},
				},
			},
			"linecount": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "author",
						"short": "The author of the poem",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "linecount",
						"short": "The number of lines in the poem (including section headings, excluding empty lines)",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "lines",
						"short": "The lines of the poem",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "title",
						"short": "The title of the poem",
						"type": "`$STRING`",
					},
				},
				"name": "linecount",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"example": "text",
											"kind": "param",
											"name": "format",
											"orig": "format",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"example": 14,
											"kind": "param",
											"name": "linecount",
											"orig": "linecount",
											"reqd": true,
											"type": "`$INTEGER`",
										},
										map[string]any{
											"example": "author,title",
											"kind": "param",
											"name": "output_field",
											"orig": "output_field",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/linecount/{linecount}/{outputFields}.{format}",
								"parts": []any{
									"linecount",
									"{linecount}",
									"{output_fields}_{format}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"outputFields}.{format": "output_fields}_{format",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"format",
										"linecount",
										"output_field",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"example": 14,
											"kind": "param",
											"name": "linecount",
											"orig": "linecount",
											"reqd": true,
											"type": "`$INTEGER`",
										},
										map[string]any{
											"example": "author,title",
											"kind": "param",
											"name": "output_field",
											"orig": "output_field",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/linecount/{linecount}/{outputFields}",
								"parts": []any{
									"linecount",
									"{linecount}",
									"{output_field}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"outputFields": "output_field",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"linecount",
										"output_field",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"example": 14,
											"kind": "param",
											"name": "id",
											"orig": "linecount",
											"reqd": true,
											"type": "`$INTEGER`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/linecount/{linecount}",
								"parts": []any{
									"linecount",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"linecount": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"linecount",
						},
					},
				},
			},
			"poemcount": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "author",
						"short": "The author of the poem",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "linecount",
						"short": "The number of lines in the poem (including section headings, excluding empty lines)",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "lines",
						"short": "The lines of the poem",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "title",
						"short": "The title of the poem",
						"type": "`$STRING`",
					},
				},
				"name": "poemcount",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"example": 10,
											"kind": "param",
											"name": "id",
											"orig": "count",
											"reqd": true,
											"type": "`$INTEGER`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/poemcount/{count}",
								"parts": []any{
									"poemcount",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"count": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"random": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "author",
						"short": "The author of the poem",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "linecount",
						"short": "The number of lines in the poem (including section headings, excluding empty lines)",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "lines",
						"short": "The lines of the poem",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "title",
						"short": "The title of the poem",
						"type": "`$STRING`",
					},
				},
				"name": "random",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"example": 5,
											"kind": "param",
											"name": "count",
											"orig": "count",
											"reqd": true,
											"type": "`$INTEGER`",
										},
										map[string]any{
											"example": "author,title",
											"kind": "param",
											"name": "output_field",
											"orig": "output_field",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/random/{count}/{outputFields}",
								"parts": []any{
									"random",
									"{count}",
									"{output_field}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"outputFields": "output_field",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"count",
										"output_field",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"example": 5,
											"kind": "param",
											"name": "id",
											"orig": "count",
											"reqd": true,
											"type": "`$INTEGER`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/random/{count}",
								"parts": []any{
									"random",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"count": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"random",
						},
					},
				},
			},
			"title": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "author",
						"short": "The author of the poem",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "linecount",
						"short": "The number of lines in the poem (including section headings, excluding empty lines)",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "lines",
						"short": "The lines of the poem",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "title",
						"short": "The title of the poem",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "titles",
						"type": "`$ARRAY`",
					},
				},
				"name": "title",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"example": "text",
											"kind": "param",
											"name": "format",
											"orig": "format",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"example": "title,lines",
											"kind": "param",
											"name": "output_field",
											"orig": "output_field",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"example": "Ozymandias",
											"kind": "param",
											"name": "title",
											"orig": "title",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/title/{title}/{outputFields}.{format}",
								"parts": []any{
									"title",
									"{title}",
									"{output_fields}_{format}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"outputFields}.{format": "output_fields}_{format",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"format",
										"output_field",
										"title",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"example": "author,title,linecount",
											"kind": "param",
											"name": "output_field",
											"orig": "output_field",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"example": "Ozymandias",
											"kind": "param",
											"name": "title",
											"orig": "title",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/title/{title}/{outputFields}",
								"parts": []any{
									"title",
									"{title}",
									"{output_field}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"outputFields": "output_field",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"output_field",
										"title",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/title",
								"parts": []any{
									"title",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.titles`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"example": "Ozymandias",
											"kind": "param",
											"name": "id",
											"orig": "title",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/title/{title}",
								"parts": []any{
									"title",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"title": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"title",
						},
					},
				},
			},
			"titleab": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "author",
						"short": "The author of the poem",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "linecount",
						"short": "The number of lines in the poem (including section headings, excluding empty lines)",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "lines",
						"short": "The lines of the poem",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "title",
						"short": "The title of the poem",
						"type": "`$STRING`",
					},
				},
				"name": "titleab",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"example": "Ozymandias",
											"kind": "param",
											"name": "title",
											"orig": "title",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/title/{title}:abs",
								"parts": []any{
									"title",
									"{title}:abs",
								},
								"select": map[string]any{
									"exist": []any{
										"title",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"title",
						},
					},
				},
			},
		},
	}
}

var (
	sharedConfigOnce sync.Once
	sharedConfigVal  map[string]any
)

// SharedConfig returns the process-wide config, built once on first use.
// The SDK reads the config on every request and never writes to it, so one
// instance is shared by every client rather than rebuilt per client.
//
// The returned map is shared: treat it as read-only. Callers that need to
// mutate should use MakeConfig, which always returns a fresh copy.
func SharedConfig() map[string]any {
	sharedConfigOnce.Do(func() {
		sharedConfigVal = MakeConfig()
	})
	return sharedConfigVal
}

func makeFeature(name string) Feature {
	switch name {
	case "test":
		if NewTestFeatureFunc != nil {
			return NewTestFeatureFunc()
		}
	default:
		if NewBaseFeatureFunc != nil {
			return NewBaseFeatureFunc()
		}
	}
	return nil
}
