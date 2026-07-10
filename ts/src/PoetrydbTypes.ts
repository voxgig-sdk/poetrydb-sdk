// Typed models for the Poetrydb SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.

export interface Author {
  author?: string
  line?: any[]
  linecount?: number
  title?: string
}

export interface AuthorLoadMatch {
  id: string
}

export interface AuthorListMatch {
  author?: string
  format?: string
  output_field?: string
}

export interface Authorab {
  author?: string
  line?: any[]
  linecount?: number
  title?: string
}

export interface AuthorabListMatch {
  author: string
}

export interface CombinedSearch {
  author?: string
  line?: any[]
  linecount?: number
  title?: string
}

export interface CombinedSearchListMatch {
  input_field1: string
  input_field2: string
  search_term1: string
  search_term2: string
}

export interface CombinedSearchWithField {
}

export interface CombinedSearchWithFieldListMatch {
  input_field1: string
  input_field2: string
  output_field: string
  search_term1: string
  search_term2: string
}

export interface Line {
  author?: string
  line?: any[]
  linecount?: number
  title?: string
}

export interface LineLoadMatch {
  id: string
}

export interface LineListMatch {
  format?: string
  line: string
  output_field: string
}

export interface Linecount {
  author?: string
  line?: any[]
  linecount?: number
  title?: string
}

export interface LinecountLoadMatch {
  id: number
}

export interface LinecountListMatch {
  format?: string
  linecount: number
  output_field: string
}

export interface Poemcount {
  author?: string
  line?: any[]
  linecount?: number
  title?: string
}

export interface PoemcountLoadMatch {
  id: number
}

export interface Random {
  author?: string
  line?: any[]
  linecount?: number
  title?: string
}

export interface RandomLoadMatch {
  id: number
}

export interface RandomListMatch {
  count: number
  output_field: string
}

export interface Title {
  author?: string
  line?: any[]
  linecount?: number
  title?: string
}

export interface TitleLoadMatch {
  id: string
}

export interface TitleListMatch {
  format?: string
  output_field?: string
  title?: string
}

export interface Titleab {
  author?: string
  line?: any[]
  linecount?: number
  title?: string
}

export interface TitleabListMatch {
  title: string
}

