# frozen_string_literal: true

# Typed models for the Poetrydb SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
# params (op.<name>.points[].args.params[]). Member types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Ruby types are unenforced; these YARD
# annotations document the shapes. Do not edit by hand.

# Author entity data model.
#
# @!attribute [rw] author
#   @return [String, nil]
#
# @!attribute [rw] line
#   @return [Array, nil]
#
# @!attribute [rw] linecount
#   @return [Integer, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
Author = Struct.new(
  :author,
  :line,
  :linecount,
  :title,
  keyword_init: true
)

# Request payload for Author#load.
#
# @!attribute [rw] id
#   @return [String]
AuthorLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Request payload for Author#list.
#
# @!attribute [rw] author
#   @return [String, nil]
#
# @!attribute [rw] format
#   @return [String, nil]
#
# @!attribute [rw] output_field
#   @return [String, nil]
AuthorListMatch = Struct.new(
  :author,
  :format,
  :output_field,
  keyword_init: true
)

# Authorab entity data model.
#
# @!attribute [rw] author
#   @return [String, nil]
#
# @!attribute [rw] line
#   @return [Array, nil]
#
# @!attribute [rw] linecount
#   @return [Integer, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
Authorab = Struct.new(
  :author,
  :line,
  :linecount,
  :title,
  keyword_init: true
)

# Request payload for Authorab#list.
#
# @!attribute [rw] author
#   @return [String]
AuthorabListMatch = Struct.new(
  :author,
  keyword_init: true
)

# CombinedSearch entity data model.
#
# @!attribute [rw] author
#   @return [String, nil]
#
# @!attribute [rw] line
#   @return [Array, nil]
#
# @!attribute [rw] linecount
#   @return [Integer, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
CombinedSearch = Struct.new(
  :author,
  :line,
  :linecount,
  :title,
  keyword_init: true
)

# Request payload for CombinedSearch#list.
#
# @!attribute [rw] input_field1
#   @return [String]
#
# @!attribute [rw] input_field2
#   @return [String]
#
# @!attribute [rw] search_term1
#   @return [String]
#
# @!attribute [rw] search_term2
#   @return [String]
CombinedSearchListMatch = Struct.new(
  :input_field1,
  :input_field2,
  :search_term1,
  :search_term2,
  keyword_init: true
)

# CombinedSearchWithField entity data model.
class CombinedSearchWithField
end

# Request payload for CombinedSearchWithField#list.
#
# @!attribute [rw] input_field1
#   @return [String]
#
# @!attribute [rw] input_field2
#   @return [String]
#
# @!attribute [rw] output_field
#   @return [String]
#
# @!attribute [rw] search_term1
#   @return [String]
#
# @!attribute [rw] search_term2
#   @return [String]
CombinedSearchWithFieldListMatch = Struct.new(
  :input_field1,
  :input_field2,
  :output_field,
  :search_term1,
  :search_term2,
  keyword_init: true
)

# Line entity data model.
#
# @!attribute [rw] author
#   @return [String, nil]
#
# @!attribute [rw] line
#   @return [Array, nil]
#
# @!attribute [rw] linecount
#   @return [Integer, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
Line = Struct.new(
  :author,
  :line,
  :linecount,
  :title,
  keyword_init: true
)

# Request payload for Line#load.
#
# @!attribute [rw] id
#   @return [String]
LineLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Request payload for Line#list.
#
# @!attribute [rw] format
#   @return [String, nil]
#
# @!attribute [rw] line
#   @return [String]
#
# @!attribute [rw] output_field
#   @return [String]
LineListMatch = Struct.new(
  :format,
  :line,
  :output_field,
  keyword_init: true
)

# Linecount entity data model.
#
# @!attribute [rw] author
#   @return [String, nil]
#
# @!attribute [rw] line
#   @return [Array, nil]
#
# @!attribute [rw] linecount
#   @return [Integer, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
Linecount = Struct.new(
  :author,
  :line,
  :linecount,
  :title,
  keyword_init: true
)

# Request payload for Linecount#load.
#
# @!attribute [rw] id
#   @return [Integer]
LinecountLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Request payload for Linecount#list.
#
# @!attribute [rw] format
#   @return [String, nil]
#
# @!attribute [rw] linecount
#   @return [Integer]
#
# @!attribute [rw] output_field
#   @return [String]
LinecountListMatch = Struct.new(
  :format,
  :linecount,
  :output_field,
  keyword_init: true
)

# Poemcount entity data model.
#
# @!attribute [rw] author
#   @return [String, nil]
#
# @!attribute [rw] line
#   @return [Array, nil]
#
# @!attribute [rw] linecount
#   @return [Integer, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
Poemcount = Struct.new(
  :author,
  :line,
  :linecount,
  :title,
  keyword_init: true
)

# Request payload for Poemcount#load.
#
# @!attribute [rw] id
#   @return [Integer]
PoemcountLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Random entity data model.
#
# @!attribute [rw] author
#   @return [String, nil]
#
# @!attribute [rw] line
#   @return [Array, nil]
#
# @!attribute [rw] linecount
#   @return [Integer, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
Random = Struct.new(
  :author,
  :line,
  :linecount,
  :title,
  keyword_init: true
)

# Request payload for Random#load.
#
# @!attribute [rw] id
#   @return [Integer]
RandomLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Request payload for Random#list.
#
# @!attribute [rw] count
#   @return [Integer]
#
# @!attribute [rw] output_field
#   @return [String]
RandomListMatch = Struct.new(
  :count,
  :output_field,
  keyword_init: true
)

# Title entity data model.
#
# @!attribute [rw] author
#   @return [String, nil]
#
# @!attribute [rw] line
#   @return [Array, nil]
#
# @!attribute [rw] linecount
#   @return [Integer, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
Title = Struct.new(
  :author,
  :line,
  :linecount,
  :title,
  keyword_init: true
)

# Request payload for Title#load.
#
# @!attribute [rw] id
#   @return [String]
TitleLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Request payload for Title#list.
#
# @!attribute [rw] format
#   @return [String, nil]
#
# @!attribute [rw] output_field
#   @return [String, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
TitleListMatch = Struct.new(
  :format,
  :output_field,
  :title,
  keyword_init: true
)

# Titleab entity data model.
#
# @!attribute [rw] author
#   @return [String, nil]
#
# @!attribute [rw] line
#   @return [Array, nil]
#
# @!attribute [rw] linecount
#   @return [Integer, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
Titleab = Struct.new(
  :author,
  :line,
  :linecount,
  :title,
  keyword_init: true
)

# Request payload for Titleab#list.
#
# @!attribute [rw] title
#   @return [String]
TitleabListMatch = Struct.new(
  :title,
  keyword_init: true
)

