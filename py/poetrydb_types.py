# Typed models for the Poetrydb SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
# params (op.<name>.points[].args.params[]). Field/param types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Do not edit by hand.
#
# These are TypedDicts, not dataclasses: the SDK ops return/accept plain dicts
# at runtime, and a TypedDict IS a dict shape, so the types match the runtime.
# Optional (req:false) keys are modelled as TypedDict key-optionality
# (total=False), split into a required base + total=False subclass when a type
# has both required and optional keys.

from __future__ import annotations

from typing import TypedDict, Any


class Author(TypedDict, total=False):
    author: str
    line: list
    linecount: int
    title: str


class AuthorLoadMatch(TypedDict):
    id: str


class AuthorListMatch(TypedDict, total=False):
    author: str
    format: str
    output_field: str


class Authorab(TypedDict, total=False):
    author: str
    line: list
    linecount: int
    title: str


class AuthorabListMatch(TypedDict):
    author: str


class CombinedSearch(TypedDict, total=False):
    author: str
    line: list
    linecount: int
    title: str


class CombinedSearchListMatch(TypedDict):
    input_field1: str
    input_field2: str
    search_term1: str
    search_term2: str


class CombinedSearchWithField(TypedDict):
    pass


class CombinedSearchWithFieldListMatch(TypedDict):
    input_field1: str
    input_field2: str
    output_field: str
    search_term1: str
    search_term2: str


class Line(TypedDict, total=False):
    author: str
    line: list
    linecount: int
    title: str


class LineLoadMatch(TypedDict):
    id: str


class LineListMatchRequired(TypedDict):
    line: str
    output_field: str


class LineListMatch(LineListMatchRequired, total=False):
    format: str


class Linecount(TypedDict, total=False):
    author: str
    line: list
    linecount: int
    title: str


class LinecountLoadMatch(TypedDict):
    id: int


class LinecountListMatchRequired(TypedDict):
    linecount: int
    output_field: str


class LinecountListMatch(LinecountListMatchRequired, total=False):
    format: str


class Poemcount(TypedDict, total=False):
    author: str
    line: list
    linecount: int
    title: str


class PoemcountLoadMatch(TypedDict):
    id: int


class Random(TypedDict, total=False):
    author: str
    line: list
    linecount: int
    title: str


class RandomLoadMatch(TypedDict):
    id: int


class RandomListMatch(TypedDict):
    count: int
    output_field: str


class Title(TypedDict, total=False):
    author: str
    line: list
    linecount: int
    title: str


class TitleLoadMatch(TypedDict):
    id: str


class TitleListMatch(TypedDict, total=False):
    format: str
    output_field: str
    title: str


class Titleab(TypedDict, total=False):
    author: str
    line: list
    linecount: int
    title: str


class TitleabListMatch(TypedDict):
    title: str
