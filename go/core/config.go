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
			"ratelimit": map[string]any{
				"options": map[string]any{
					"active": false,
					"burst": 5,
					"rate": 5,
				},
				"optspec": map[string]any{
					"now": "`$FUNCTION`",
					"sleep": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "wrap",
			},
			"retry": map[string]any{
				"options": map[string]any{
					"active": false,
					"factor": 2,
					"maxDelay": 2000,
					"minDelay": 50,
					"retries": 2,
					"statuses": []any{
						408,
						425,
						429,
						500,
						502,
						503,
						504,
					},
				},
				"optspec": map[string]any{
					"jitter": "`$BOOLEAN`",
					"sleep": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "wrap",
			},
			"test": map[string]any{
				"options": map[string]any{
					"active": false,
				},
				"optspec": map[string]any{
					"entity": "`$MAP`",
					"net": "`$MAP`",
				},
				"strict": false,
				"transport": "base",
			},
			"timeout": map[string]any{
				"options": map[string]any{
					"active": false,
					"ms": 30000,
				},
				"optspec": map[string]any{
					"clearTimer": "`$FUNCTION`",
					"setTimer": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "wrap",
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
				"id": map[string]any{
					"field": "id",
					"name": "id",
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
								"rename": map[string]any{
									"param": map[string]any{
										"outputFields}.{format": "output_fields}_{format",
									},
								},
								"segments": []any{
									map[string]any{
										"lit": "author",
									},
									map[string]any{
										"var": "author",
									},
									map[string]any{
										"lit": "{outputFields}.{format}",
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
								"parts": []any{
									"author",
									"{author}",
									"{outputFields}.{format}",
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
								"rename": map[string]any{
									"param": map[string]any{
										"outputFields": "output_field",
									},
								},
								"segments": []any{
									map[string]any{
										"lit": "author",
									},
									map[string]any{
										"var": "author",
									},
									map[string]any{
										"var": "output_field",
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
								"parts": []any{
									"author",
									"{author}",
									"{output_field}",
								},
							},
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/author",
								"segments": []any{
									map[string]any{
										"lit": "author",
									},
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.authors`",
								},
								"parts": []any{
									"author",
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
								"rename": map[string]any{
									"param": map[string]any{
										"author": "id",
									},
								},
								"segments": []any{
									map[string]any{
										"lit": "author",
									},
									map[string]any{
										"var": "id",
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
								"parts": []any{
									"author",
									"{id}",
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
								"segments": []any{
									map[string]any{
										"lit": "author",
									},
									map[string]any{
										"lit": "{author}:abs",
									},
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
								"parts": []any{
									"author",
									"{author}:abs",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
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
								"rename": map[string]any{
									"param": map[string]any{
										"inputField1},{inputField2": "input_field1},{input_field2",
										"searchTerm1};{searchTerm2": "search_term1};{search_term2",
									},
								},
								"segments": []any{
									map[string]any{
										"lit": "{inputField1},{inputField2}",
									},
									map[string]any{
										"lit": "{searchTerm1};{searchTerm2}",
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
								"parts": []any{
									"{inputField1},{inputField2}",
									"{searchTerm1};{searchTerm2}",
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
								"rename": map[string]any{
									"param": map[string]any{
										"inputField1},{inputField2": "input_field1},{input_field2",
										"outputFields": "output_field",
										"searchTerm1};{searchTerm2": "search_term1};{search_term2",
									},
								},
								"segments": []any{
									map[string]any{
										"lit": "{inputField1},{inputField2}",
									},
									map[string]any{
										"lit": "{searchTerm1};{searchTerm2}",
									},
									map[string]any{
										"var": "output_field",
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
								"parts": []any{
									"{inputField1},{inputField2}",
									"{searchTerm1};{searchTerm2}",
									"{output_field}",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"{search_term1};{search_term2}",
						},
					},
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
				"id": map[string]any{
					"field": "id",
					"name": "id",
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
								"rename": map[string]any{
									"param": map[string]any{
										"lines": "line",
										"outputFields}.{format": "output_fields}_{format",
									},
								},
								"segments": []any{
									map[string]any{
										"lit": "lines",
									},
									map[string]any{
										"var": "line",
									},
									map[string]any{
										"lit": "{outputFields}.{format}",
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
								"parts": []any{
									"lines",
									"{line}",
									"{outputFields}.{format}",
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
								"rename": map[string]any{
									"param": map[string]any{
										"lines": "line",
										"outputFields": "output_field",
									},
								},
								"segments": []any{
									map[string]any{
										"lit": "lines",
									},
									map[string]any{
										"var": "line",
									},
									map[string]any{
										"var": "output_field",
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
								"parts": []any{
									"lines",
									"{line}",
									"{output_field}",
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
								"rename": map[string]any{
									"param": map[string]any{
										"lines": "id",
									},
								},
								"segments": []any{
									map[string]any{
										"lit": "lines",
									},
									map[string]any{
										"var": "id",
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
								"parts": []any{
									"lines",
									"{id}",
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
				"id": map[string]any{
					"field": "id",
					"name": "id",
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
								"rename": map[string]any{
									"param": map[string]any{
										"outputFields}.{format": "output_fields}_{format",
									},
								},
								"segments": []any{
									map[string]any{
										"lit": "linecount",
									},
									map[string]any{
										"var": "linecount",
									},
									map[string]any{
										"lit": "{outputFields}.{format}",
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
								"parts": []any{
									"linecount",
									"{linecount}",
									"{outputFields}.{format}",
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
								"rename": map[string]any{
									"param": map[string]any{
										"outputFields": "output_field",
									},
								},
								"segments": []any{
									map[string]any{
										"lit": "linecount",
									},
									map[string]any{
										"var": "linecount",
									},
									map[string]any{
										"var": "output_field",
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
								"parts": []any{
									"linecount",
									"{linecount}",
									"{output_field}",
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
								"rename": map[string]any{
									"param": map[string]any{
										"linecount": "id",
									},
								},
								"segments": []any{
									map[string]any{
										"lit": "linecount",
									},
									map[string]any{
										"var": "id",
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
								"parts": []any{
									"linecount",
									"{id}",
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
				"id": map[string]any{
					"field": "id",
					"name": "id",
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
								"rename": map[string]any{
									"param": map[string]any{
										"count": "id",
									},
								},
								"segments": []any{
									map[string]any{
										"lit": "poemcount",
									},
									map[string]any{
										"var": "id",
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
								"parts": []any{
									"poemcount",
									"{id}",
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
				"id": map[string]any{
					"field": "id",
					"name": "id",
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
								"rename": map[string]any{
									"param": map[string]any{
										"outputFields": "output_field",
									},
								},
								"segments": []any{
									map[string]any{
										"lit": "random",
									},
									map[string]any{
										"var": "count",
									},
									map[string]any{
										"var": "output_field",
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
								"parts": []any{
									"random",
									"{count}",
									"{output_field}",
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
								"rename": map[string]any{
									"param": map[string]any{
										"count": "id",
									},
								},
								"segments": []any{
									map[string]any{
										"lit": "random",
									},
									map[string]any{
										"var": "id",
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
								"parts": []any{
									"random",
									"{id}",
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
				"id": map[string]any{
					"field": "id",
					"name": "id",
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
								"rename": map[string]any{
									"param": map[string]any{
										"outputFields}.{format": "output_fields}_{format",
									},
								},
								"segments": []any{
									map[string]any{
										"lit": "title",
									},
									map[string]any{
										"var": "title",
									},
									map[string]any{
										"lit": "{outputFields}.{format}",
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
								"parts": []any{
									"title",
									"{title}",
									"{outputFields}.{format}",
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
								"rename": map[string]any{
									"param": map[string]any{
										"outputFields": "output_field",
									},
								},
								"segments": []any{
									map[string]any{
										"lit": "title",
									},
									map[string]any{
										"var": "title",
									},
									map[string]any{
										"var": "output_field",
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
								"parts": []any{
									"title",
									"{title}",
									"{output_field}",
								},
							},
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/title",
								"segments": []any{
									map[string]any{
										"lit": "title",
									},
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.titles`",
								},
								"parts": []any{
									"title",
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
								"rename": map[string]any{
									"param": map[string]any{
										"title": "id",
									},
								},
								"segments": []any{
									map[string]any{
										"lit": "title",
									},
									map[string]any{
										"var": "id",
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
								"parts": []any{
									"title",
									"{id}",
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
								"segments": []any{
									map[string]any{
										"lit": "title",
									},
									map[string]any{
										"lit": "{title}:abs",
									},
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
								"parts": []any{
									"title",
									"{title}:abs",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
		},
	}
}

// The plugin definitions the model selected per feature, as []any so a
// feature package can consume them without core naming its types. Empty
// when no active feature declares active plugin groups for this target.
var featurePlugins = map[string][]any{
}

// FeaturePlugins is the definitions list for one feature's chain.
func FeaturePlugins(name string) []any {
	return featurePlugins[name]
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
	case "ratelimit":
		if NewRatelimitFeatureFunc != nil {
			return NewRatelimitFeatureFunc()
		}
	case "retry":
		if NewRetryFeatureFunc != nil {
			return NewRetryFeatureFunc()
		}
	case "test":
		if NewTestFeatureFunc != nil {
			return NewTestFeatureFunc()
		}
	case "timeout":
		if NewTimeoutFeatureFunc != nil {
			return NewTimeoutFeatureFunc()
		}
	default:
		if NewBaseFeatureFunc != nil {
			return NewBaseFeatureFunc()
		}
	}
	return nil
}
