package core

func MakeConfig() map[string]any {
	return map[string]any{
		"main": map[string]any{
			"name": "Poetrydb",
		},
		"feature": map[string]any{
			"test": map[string]any{
				"options": map[string]any{
					"active": false,
				},
			},
		},
		"options": map[string]any{
			"base": "https://poetrydb.org",
			"auth": map[string]any{
				"prefix": "Bearer",
			},
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
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 0,
					},
					map[string]any{
						"name": "line",
						"req": false,
						"type": "`$ARRAY`",
						"active": true,
						"index$": 1,
					},
					map[string]any{
						"name": "linecount",
						"req": false,
						"type": "`$INTEGER`",
						"active": true,
						"index$": 2,
					},
					map[string]any{
						"name": "title",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 3,
					},
				},
				"name": "author",
				"op": map[string]any{
					"list": map[string]any{
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
											"active": true,
										},
										map[string]any{
											"example": "text",
											"kind": "param",
											"name": "format",
											"orig": "format",
											"reqd": true,
											"type": "`$STRING`",
											"active": true,
										},
										map[string]any{
											"example": "author,title,linecount",
											"kind": "param",
											"name": "output_field",
											"orig": "output_field",
											"reqd": true,
											"type": "`$STRING`",
											"active": true,
										},
									},
								},
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
								"active": true,
								"index$": 0,
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
											"active": true,
										},
										map[string]any{
											"example": "author,title,linecount",
											"kind": "param",
											"name": "output_field",
											"orig": "output_field",
											"reqd": true,
											"type": "`$STRING`",
											"active": true,
										},
									},
								},
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
								"active": true,
								"index$": 1,
							},
							map[string]any{
								"method": "GET",
								"orig": "/author",
								"parts": []any{
									"author",
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"active": true,
								"args": map[string]any{},
								"select": map[string]any{},
								"index$": 2,
							},
						},
						"input": "data",
						"key$": "list",
					},
					"load": map[string]any{
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
											"active": true,
										},
									},
								},
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
								"active": true,
								"index$": 0,
							},
						},
						"input": "data",
						"key$": "load",
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
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 0,
					},
					map[string]any{
						"name": "line",
						"req": false,
						"type": "`$ARRAY`",
						"active": true,
						"index$": 1,
					},
					map[string]any{
						"name": "linecount",
						"req": false,
						"type": "`$INTEGER`",
						"active": true,
						"index$": 2,
					},
					map[string]any{
						"name": "title",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 3,
					},
				},
				"name": "authorab",
				"op": map[string]any{
					"list": map[string]any{
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
											"active": true,
										},
									},
								},
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
								"active": true,
								"index$": 0,
							},
						},
						"input": "data",
						"key$": "list",
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
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 0,
					},
					map[string]any{
						"name": "line",
						"req": false,
						"type": "`$ARRAY`",
						"active": true,
						"index$": 1,
					},
					map[string]any{
						"name": "linecount",
						"req": false,
						"type": "`$INTEGER`",
						"active": true,
						"index$": 2,
					},
					map[string]any{
						"name": "title",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 3,
					},
				},
				"name": "combined_search",
				"op": map[string]any{
					"list": map[string]any{
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
											"active": true,
										},
										map[string]any{
											"example": "linecount",
											"kind": "param",
											"name": "input_field2",
											"orig": "input_field2",
											"reqd": true,
											"type": "`$STRING`",
											"active": true,
										},
										map[string]any{
											"example": "Shakespeare",
											"kind": "param",
											"name": "search_term1",
											"orig": "search_term1",
											"reqd": true,
											"type": "`$STRING`",
											"active": true,
										},
										map[string]any{
											"example": "14",
											"kind": "param",
											"name": "search_term2",
											"orig": "search_term2",
											"reqd": true,
											"type": "`$STRING`",
											"active": true,
										},
									},
								},
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
								"active": true,
								"index$": 0,
							},
						},
						"input": "data",
						"key$": "list",
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
											"active": true,
										},
										map[string]any{
											"example": "linecount",
											"kind": "param",
											"name": "input_field2",
											"orig": "input_field2",
											"reqd": true,
											"type": "`$STRING`",
											"active": true,
										},
										map[string]any{
											"example": "lines",
											"kind": "param",
											"name": "output_field",
											"orig": "output_field",
											"reqd": true,
											"type": "`$STRING`",
											"active": true,
										},
										map[string]any{
											"example": "Shakespeare",
											"kind": "param",
											"name": "search_term1",
											"orig": "search_term1",
											"reqd": true,
											"type": "`$STRING`",
											"active": true,
										},
										map[string]any{
											"example": "14",
											"kind": "param",
											"name": "search_term2",
											"orig": "search_term2",
											"reqd": true,
											"type": "`$STRING`",
											"active": true,
										},
									},
								},
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
								"active": true,
								"index$": 0,
							},
						},
						"input": "data",
						"key$": "list",
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
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 0,
					},
					map[string]any{
						"name": "line",
						"req": false,
						"type": "`$ARRAY`",
						"active": true,
						"index$": 1,
					},
					map[string]any{
						"name": "linecount",
						"req": false,
						"type": "`$INTEGER`",
						"active": true,
						"index$": 2,
					},
					map[string]any{
						"name": "title",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 3,
					},
				},
				"name": "line",
				"op": map[string]any{
					"list": map[string]any{
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
											"active": true,
										},
										map[string]any{
											"example": "Latitudeless Place",
											"kind": "param",
											"name": "line",
											"orig": "line",
											"reqd": true,
											"type": "`$STRING`",
											"active": true,
										},
										map[string]any{
											"example": "author,title,linecount",
											"kind": "param",
											"name": "output_field",
											"orig": "output_field",
											"reqd": true,
											"type": "`$STRING`",
											"active": true,
										},
									},
								},
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
								"active": true,
								"index$": 0,
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
											"active": true,
										},
										map[string]any{
											"example": "author,title,linecount",
											"kind": "param",
											"name": "output_field",
											"orig": "output_field",
											"reqd": true,
											"type": "`$STRING`",
											"active": true,
										},
									},
								},
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
								"active": true,
								"index$": 1,
							},
						},
						"input": "data",
						"key$": "list",
					},
					"load": map[string]any{
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
											"active": true,
										},
									},
								},
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
								"active": true,
								"index$": 0,
							},
						},
						"input": "data",
						"key$": "load",
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
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 0,
					},
					map[string]any{
						"name": "line",
						"req": false,
						"type": "`$ARRAY`",
						"active": true,
						"index$": 1,
					},
					map[string]any{
						"name": "linecount",
						"req": false,
						"type": "`$INTEGER`",
						"active": true,
						"index$": 2,
					},
					map[string]any{
						"name": "title",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 3,
					},
				},
				"name": "linecount",
				"op": map[string]any{
					"list": map[string]any{
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
											"active": true,
										},
										map[string]any{
											"example": 14,
											"kind": "param",
											"name": "linecount",
											"orig": "linecount",
											"reqd": true,
											"type": "`$INTEGER`",
											"active": true,
										},
										map[string]any{
											"example": "author,title",
											"kind": "param",
											"name": "output_field",
											"orig": "output_field",
											"reqd": true,
											"type": "`$STRING`",
											"active": true,
										},
									},
								},
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
								"active": true,
								"index$": 0,
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
											"active": true,
										},
										map[string]any{
											"example": "author,title",
											"kind": "param",
											"name": "output_field",
											"orig": "output_field",
											"reqd": true,
											"type": "`$STRING`",
											"active": true,
										},
									},
								},
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
								"active": true,
								"index$": 1,
							},
						},
						"input": "data",
						"key$": "list",
					},
					"load": map[string]any{
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
											"active": true,
										},
									},
								},
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
								"active": true,
								"index$": 0,
							},
						},
						"input": "data",
						"key$": "load",
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
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 0,
					},
					map[string]any{
						"name": "line",
						"req": false,
						"type": "`$ARRAY`",
						"active": true,
						"index$": 1,
					},
					map[string]any{
						"name": "linecount",
						"req": false,
						"type": "`$INTEGER`",
						"active": true,
						"index$": 2,
					},
					map[string]any{
						"name": "title",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 3,
					},
				},
				"name": "poemcount",
				"op": map[string]any{
					"load": map[string]any{
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
											"active": true,
										},
									},
								},
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
								"active": true,
								"index$": 0,
							},
						},
						"input": "data",
						"key$": "load",
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
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 0,
					},
					map[string]any{
						"name": "line",
						"req": false,
						"type": "`$ARRAY`",
						"active": true,
						"index$": 1,
					},
					map[string]any{
						"name": "linecount",
						"req": false,
						"type": "`$INTEGER`",
						"active": true,
						"index$": 2,
					},
					map[string]any{
						"name": "title",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 3,
					},
				},
				"name": "random",
				"op": map[string]any{
					"list": map[string]any{
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
											"active": true,
										},
										map[string]any{
											"example": "author,title",
											"kind": "param",
											"name": "output_field",
											"orig": "output_field",
											"reqd": true,
											"type": "`$STRING`",
											"active": true,
										},
									},
								},
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
								"active": true,
								"index$": 0,
							},
						},
						"input": "data",
						"key$": "list",
					},
					"load": map[string]any{
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
											"active": true,
										},
									},
								},
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
								"active": true,
								"index$": 0,
							},
						},
						"input": "data",
						"key$": "load",
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
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 0,
					},
					map[string]any{
						"name": "line",
						"req": false,
						"type": "`$ARRAY`",
						"active": true,
						"index$": 1,
					},
					map[string]any{
						"name": "linecount",
						"req": false,
						"type": "`$INTEGER`",
						"active": true,
						"index$": 2,
					},
					map[string]any{
						"name": "title",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 3,
					},
				},
				"name": "title",
				"op": map[string]any{
					"list": map[string]any{
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
											"active": true,
										},
										map[string]any{
											"example": "title,lines",
											"kind": "param",
											"name": "output_field",
											"orig": "output_field",
											"reqd": true,
											"type": "`$STRING`",
											"active": true,
										},
										map[string]any{
											"example": "Ozymandias",
											"kind": "param",
											"name": "title",
											"orig": "title",
											"reqd": true,
											"type": "`$STRING`",
											"active": true,
										},
									},
								},
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
								"active": true,
								"index$": 0,
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
											"active": true,
										},
										map[string]any{
											"example": "Ozymandias",
											"kind": "param",
											"name": "title",
											"orig": "title",
											"reqd": true,
											"type": "`$STRING`",
											"active": true,
										},
									},
								},
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
								"active": true,
								"index$": 1,
							},
							map[string]any{
								"method": "GET",
								"orig": "/title",
								"parts": []any{
									"title",
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"active": true,
								"args": map[string]any{},
								"select": map[string]any{},
								"index$": 2,
							},
						},
						"input": "data",
						"key$": "list",
					},
					"load": map[string]any{
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
											"active": true,
										},
									},
								},
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
								"active": true,
								"index$": 0,
							},
						},
						"input": "data",
						"key$": "load",
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
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 0,
					},
					map[string]any{
						"name": "line",
						"req": false,
						"type": "`$ARRAY`",
						"active": true,
						"index$": 1,
					},
					map[string]any{
						"name": "linecount",
						"req": false,
						"type": "`$INTEGER`",
						"active": true,
						"index$": 2,
					},
					map[string]any{
						"name": "title",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 3,
					},
				},
				"name": "titleab",
				"op": map[string]any{
					"list": map[string]any{
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
											"active": true,
										},
									},
								},
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
								"active": true,
								"index$": 0,
							},
						},
						"input": "data",
						"key$": "list",
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
