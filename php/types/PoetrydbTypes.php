<?php
declare(strict_types=1);

// Typed models for the Poetrydb SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
//
// These are documentation-grade value objects (PHP 8 typed properties),
// registered on the composer classmap autoload. The SDK boundary exchanges
// assoc-arrays; these classes name the shapes for tooling and typed callers.

/** Author entity data model. */
class Author
{
    public ?string $author = null;
    public ?array $line = null;
    public ?int $linecount = null;
    public ?string $title = null;
}

/** Request payload for Author#load. */
class AuthorLoadMatch
{
    public string $id;
}

/** Request payload for Author#list. */
class AuthorListMatch
{
    public string $author;
    public string $format;
    public string $output_field;
}

/** Authorab entity data model. */
class Authorab
{
    public ?string $author = null;
    public ?array $line = null;
    public ?int $linecount = null;
    public ?string $title = null;
}

/** Request payload for Authorab#list. */
class AuthorabListMatch
{
    public string $author;
}

/** CombinedSearch entity data model. */
class CombinedSearch
{
    public ?string $author = null;
    public ?array $line = null;
    public ?int $linecount = null;
    public ?string $title = null;
}

/** Request payload for CombinedSearch#list. */
class CombinedSearchListMatch
{
    public string $input_field1;
    public string $input_field2;
    public string $search_term1;
    public string $search_term2;
}

/** CombinedSearchWithField entity data model. */
class CombinedSearchWithField
{
}

/** Request payload for CombinedSearchWithField#list. */
class CombinedSearchWithFieldListMatch
{
    public string $input_field1;
    public string $input_field2;
    public string $output_field;
    public string $search_term1;
    public string $search_term2;
}

/** Line entity data model. */
class Line
{
    public ?string $author = null;
    public ?array $line = null;
    public ?int $linecount = null;
    public ?string $title = null;
}

/** Request payload for Line#load. */
class LineLoadMatch
{
    public string $id;
}

/** Request payload for Line#list. */
class LineListMatch
{
    public string $format;
    public string $line;
    public string $output_field;
}

/** Linecount entity data model. */
class Linecount
{
    public ?string $author = null;
    public ?array $line = null;
    public ?int $linecount = null;
    public ?string $title = null;
}

/** Request payload for Linecount#load. */
class LinecountLoadMatch
{
    public int $id;
}

/** Request payload for Linecount#list. */
class LinecountListMatch
{
    public string $format;
    public int $linecount;
    public string $output_field;
}

/** Poemcount entity data model. */
class Poemcount
{
    public ?string $author = null;
    public ?array $line = null;
    public ?int $linecount = null;
    public ?string $title = null;
}

/** Request payload for Poemcount#load. */
class PoemcountLoadMatch
{
    public int $id;
}

/** Random entity data model. */
class Random
{
    public ?string $author = null;
    public ?array $line = null;
    public ?int $linecount = null;
    public ?string $title = null;
}

/** Request payload for Random#load. */
class RandomLoadMatch
{
    public int $id;
}

/** Request payload for Random#list. */
class RandomListMatch
{
    public int $count;
    public string $output_field;
}

/** Title entity data model. */
class Title
{
    public ?string $author = null;
    public ?array $line = null;
    public ?int $linecount = null;
    public ?string $title = null;
}

/** Request payload for Title#load. */
class TitleLoadMatch
{
    public string $id;
}

/** Request payload for Title#list. */
class TitleListMatch
{
    public string $format;
    public string $output_field;
    public string $title;
}

/** Titleab entity data model. */
class Titleab
{
    public ?string $author = null;
    public ?array $line = null;
    public ?int $linecount = null;
    public ?string $title = null;
}

/** Request payload for Titleab#list. */
class TitleabListMatch
{
    public string $title;
}

