// Typed models for the Poetrydb SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
package entity

import "encoding/json"

// Author is the typed data model for the author entity.
type Author struct {
	Author *string `json:"author,omitempty"`
	Line *[]any `json:"line,omitempty"`
	Linecount *int `json:"linecount,omitempty"`
	Title *string `json:"title,omitempty"`
}

// AuthorLoadMatch is the typed request payload for Author.LoadTyped.
type AuthorLoadMatch struct {
	Id string `json:"id"`
}

// AuthorListMatch is the typed request payload for Author.ListTyped.
type AuthorListMatch struct {
	Author string `json:"author"`
	Format string `json:"format"`
	OutputField string `json:"output_field"`
}

// Authorab is the typed data model for the authorab entity.
type Authorab struct {
	Author *string `json:"author,omitempty"`
	Line *[]any `json:"line,omitempty"`
	Linecount *int `json:"linecount,omitempty"`
	Title *string `json:"title,omitempty"`
}

// AuthorabListMatch is the typed request payload for Authorab.ListTyped.
type AuthorabListMatch struct {
	Author string `json:"author"`
}

// CombinedSearch is the typed data model for the combined_search entity.
type CombinedSearch struct {
	Author *string `json:"author,omitempty"`
	Line *[]any `json:"line,omitempty"`
	Linecount *int `json:"linecount,omitempty"`
	Title *string `json:"title,omitempty"`
}

// CombinedSearchListMatch is the typed request payload for CombinedSearch.ListTyped.
type CombinedSearchListMatch struct {
	InputField1 string `json:"input_field1"`
	InputField2 string `json:"input_field2"`
	SearchTerm1 string `json:"search_term1"`
	SearchTerm2 string `json:"search_term2"`
}

// CombinedSearchWithField is the typed data model for the combined_search_with_field entity.
type CombinedSearchWithField struct {
}

// CombinedSearchWithFieldListMatch is the typed request payload for CombinedSearchWithField.ListTyped.
type CombinedSearchWithFieldListMatch struct {
	InputField1 string `json:"input_field1"`
	InputField2 string `json:"input_field2"`
	OutputField string `json:"output_field"`
	SearchTerm1 string `json:"search_term1"`
	SearchTerm2 string `json:"search_term2"`
}

// Line is the typed data model for the line entity.
type Line struct {
	Author *string `json:"author,omitempty"`
	Line *[]any `json:"line,omitempty"`
	Linecount *int `json:"linecount,omitempty"`
	Title *string `json:"title,omitempty"`
}

// LineLoadMatch is the typed request payload for Line.LoadTyped.
type LineLoadMatch struct {
	Id string `json:"id"`
}

// LineListMatch is the typed request payload for Line.ListTyped.
type LineListMatch struct {
	Format string `json:"format"`
	Line string `json:"line"`
	OutputField string `json:"output_field"`
}

// Linecount is the typed data model for the linecount entity.
type Linecount struct {
	Author *string `json:"author,omitempty"`
	Line *[]any `json:"line,omitempty"`
	Linecount *int `json:"linecount,omitempty"`
	Title *string `json:"title,omitempty"`
}

// LinecountLoadMatch is the typed request payload for Linecount.LoadTyped.
type LinecountLoadMatch struct {
	Id int `json:"id"`
}

// LinecountListMatch is the typed request payload for Linecount.ListTyped.
type LinecountListMatch struct {
	Format string `json:"format"`
	Linecount int `json:"linecount"`
	OutputField string `json:"output_field"`
}

// Poemcount is the typed data model for the poemcount entity.
type Poemcount struct {
	Author *string `json:"author,omitempty"`
	Line *[]any `json:"line,omitempty"`
	Linecount *int `json:"linecount,omitempty"`
	Title *string `json:"title,omitempty"`
}

// PoemcountLoadMatch is the typed request payload for Poemcount.LoadTyped.
type PoemcountLoadMatch struct {
	Id int `json:"id"`
}

// Random is the typed data model for the random entity.
type Random struct {
	Author *string `json:"author,omitempty"`
	Line *[]any `json:"line,omitempty"`
	Linecount *int `json:"linecount,omitempty"`
	Title *string `json:"title,omitempty"`
}

// RandomLoadMatch is the typed request payload for Random.LoadTyped.
type RandomLoadMatch struct {
	Id int `json:"id"`
}

// RandomListMatch is the typed request payload for Random.ListTyped.
type RandomListMatch struct {
	Count int `json:"count"`
	OutputField string `json:"output_field"`
}

// Title is the typed data model for the title entity.
type Title struct {
	Author *string `json:"author,omitempty"`
	Line *[]any `json:"line,omitempty"`
	Linecount *int `json:"linecount,omitempty"`
	Title *string `json:"title,omitempty"`
}

// TitleLoadMatch is the typed request payload for Title.LoadTyped.
type TitleLoadMatch struct {
	Id string `json:"id"`
}

// TitleListMatch is the typed request payload for Title.ListTyped.
type TitleListMatch struct {
	Format string `json:"format"`
	OutputField string `json:"output_field"`
	Title string `json:"title"`
}

// Titleab is the typed data model for the titleab entity.
type Titleab struct {
	Author *string `json:"author,omitempty"`
	Line *[]any `json:"line,omitempty"`
	Linecount *int `json:"linecount,omitempty"`
	Title *string `json:"title,omitempty"`
}

// TitleabListMatch is the typed request payload for Titleab.ListTyped.
type TitleabListMatch struct {
	Title string `json:"title"`
}

// asMap turns a typed request/data struct into the map[string]any the
// runtime op pipeline consumes, honouring the json tags above.
func asMap(v any) map[string]any {
	out := map[string]any{}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

// typedFrom decodes a runtime value (a map[string]any produced by the op
// pipeline) into a typed model T via a JSON round-trip. On any error it
// returns the zero value of T; the op's own (value, error) tuple carries the
// real error.
func typedFrom[T any](v any) T {
	var out T
	if v == nil {
		return out
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

// typedSliceFrom decodes a runtime list value ([]any of maps) into a typed
// slice []T via a JSON round-trip, for list ops.
func typedSliceFrom[T any](v any) []T {
	var out []T
	if v == nil {
		return out
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}
