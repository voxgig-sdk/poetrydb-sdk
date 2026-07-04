-- Typed models for the Poetrydb SDK (LuaLS annotations).
--
-- GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
-- params (op.<name>.points[].args.params[]). Field/param types come from the
-- canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
-- @voxgig/apidef VALID_CANON). Annotations only — no runtime effect. Do not
-- edit by hand.

---@class Author
---@field author? string
---@field line? table
---@field linecount? number
---@field title? string

---@class AuthorLoadMatch
---@field id string

---@class AuthorListMatch
---@field author string
---@field format string
---@field output_field string

---@class Authorab
---@field author? string
---@field line? table
---@field linecount? number
---@field title? string

---@class AuthorabListMatch
---@field author string

---@class CombinedSearch
---@field author? string
---@field line? table
---@field linecount? number
---@field title? string

---@class CombinedSearchListMatch
---@field input_field1 string
---@field input_field2 string
---@field search_term1 string
---@field search_term2 string

---@class CombinedSearchWithField

---@class CombinedSearchWithFieldListMatch
---@field input_field1 string
---@field input_field2 string
---@field output_field string
---@field search_term1 string
---@field search_term2 string

---@class Line
---@field author? string
---@field line? table
---@field linecount? number
---@field title? string

---@class LineLoadMatch
---@field id string

---@class LineListMatch
---@field format string
---@field line string
---@field output_field string

---@class Linecount
---@field author? string
---@field line? table
---@field linecount? number
---@field title? string

---@class LinecountLoadMatch
---@field id number

---@class LinecountListMatch
---@field format string
---@field linecount number
---@field output_field string

---@class Poemcount
---@field author? string
---@field line? table
---@field linecount? number
---@field title? string

---@class PoemcountLoadMatch
---@field id number

---@class Random
---@field author? string
---@field line? table
---@field linecount? number
---@field title? string

---@class RandomLoadMatch
---@field id number

---@class RandomListMatch
---@field count number
---@field output_field string

---@class Title
---@field author? string
---@field line? table
---@field linecount? number
---@field title? string

---@class TitleLoadMatch
---@field id string

---@class TitleListMatch
---@field format string
---@field output_field string
---@field title string

---@class Titleab
---@field author? string
---@field line? table
---@field linecount? number
---@field title? string

---@class TitleabListMatch
---@field title string

local M = {}

return M
