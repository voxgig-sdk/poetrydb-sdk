# Typed models for the Poetrydb SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
# params (op.<name>.points[].args.params[]). Field/param types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Do not edit by hand.

from __future__ import annotations

from dataclasses import dataclass
from typing import Optional, Any


@dataclass
class Author:
    author: Optional[str] = None
    line: Optional[list] = None
    linecount: Optional[int] = None
    title: Optional[str] = None


@dataclass
class AuthorLoadMatch:
    id: str


@dataclass
class AuthorListMatch:
    author: str
    format: str
    output_field: str


@dataclass
class Authorab:
    author: Optional[str] = None
    line: Optional[list] = None
    linecount: Optional[int] = None
    title: Optional[str] = None


@dataclass
class AuthorabListMatch:
    author: str


@dataclass
class CombinedSearch:
    author: Optional[str] = None
    line: Optional[list] = None
    linecount: Optional[int] = None
    title: Optional[str] = None


@dataclass
class CombinedSearchListMatch:
    input_field1: str
    input_field2: str
    search_term1: str
    search_term2: str


@dataclass
class CombinedSearchWithField:
    pass


@dataclass
class CombinedSearchWithFieldListMatch:
    input_field1: str
    input_field2: str
    output_field: str
    search_term1: str
    search_term2: str


@dataclass
class Line:
    author: Optional[str] = None
    line: Optional[list] = None
    linecount: Optional[int] = None
    title: Optional[str] = None


@dataclass
class LineLoadMatch:
    id: str


@dataclass
class LineListMatch:
    format: str
    line: str
    output_field: str


@dataclass
class Linecount:
    author: Optional[str] = None
    line: Optional[list] = None
    linecount: Optional[int] = None
    title: Optional[str] = None


@dataclass
class LinecountLoadMatch:
    id: int


@dataclass
class LinecountListMatch:
    format: str
    linecount: int
    output_field: str


@dataclass
class Poemcount:
    author: Optional[str] = None
    line: Optional[list] = None
    linecount: Optional[int] = None
    title: Optional[str] = None


@dataclass
class PoemcountLoadMatch:
    id: int


@dataclass
class Random:
    author: Optional[str] = None
    line: Optional[list] = None
    linecount: Optional[int] = None
    title: Optional[str] = None


@dataclass
class RandomLoadMatch:
    id: int


@dataclass
class RandomListMatch:
    count: int
    output_field: str


@dataclass
class Title:
    author: Optional[str] = None
    line: Optional[list] = None
    linecount: Optional[int] = None
    title: Optional[str] = None


@dataclass
class TitleLoadMatch:
    id: str


@dataclass
class TitleListMatch:
    format: str
    output_field: str
    title: str


@dataclass
class Titleab:
    author: Optional[str] = None
    line: Optional[list] = None
    linecount: Optional[int] = None
    title: Optional[str] = None


@dataclass
class TitleabListMatch:
    title: str

