package voxgigpoetrydbsdk

import (
	"github.com/voxgig-sdk/poetrydb-sdk/go/core"
	"github.com/voxgig-sdk/poetrydb-sdk/go/entity"
	"github.com/voxgig-sdk/poetrydb-sdk/go/feature"
	_ "github.com/voxgig-sdk/poetrydb-sdk/go/utility"
)

// Type aliases preserve external API.
type PoetrydbSDK = core.PoetrydbSDK
type Context = core.Context
type Utility = core.Utility
type Feature = core.Feature
type Entity = core.Entity
type PoetrydbEntity = core.PoetrydbEntity
type FetcherFunc = core.FetcherFunc
type Spec = core.Spec
type Result = core.Result
type Response = core.Response
type Operation = core.Operation
type Control = core.Control
type PoetrydbError = core.PoetrydbError

// BaseFeature from feature package.
type BaseFeature = feature.BaseFeature

func init() {
	core.NewBaseFeatureFunc = func() core.Feature {
		return feature.NewBaseFeature()
	}
	core.NewTestFeatureFunc = func() core.Feature {
		return feature.NewTestFeature()
	}
	core.NewAuthorEntityFunc = func(client *core.PoetrydbSDK, entopts map[string]any) core.PoetrydbEntity {
		return entity.NewAuthorEntity(client, entopts)
	}
	core.NewAuthorabEntityFunc = func(client *core.PoetrydbSDK, entopts map[string]any) core.PoetrydbEntity {
		return entity.NewAuthorabEntity(client, entopts)
	}
	core.NewCombinedSearchEntityFunc = func(client *core.PoetrydbSDK, entopts map[string]any) core.PoetrydbEntity {
		return entity.NewCombinedSearchEntity(client, entopts)
	}
	core.NewCombinedSearchWithFieldEntityFunc = func(client *core.PoetrydbSDK, entopts map[string]any) core.PoetrydbEntity {
		return entity.NewCombinedSearchWithFieldEntity(client, entopts)
	}
	core.NewLineEntityFunc = func(client *core.PoetrydbSDK, entopts map[string]any) core.PoetrydbEntity {
		return entity.NewLineEntity(client, entopts)
	}
	core.NewLinecountEntityFunc = func(client *core.PoetrydbSDK, entopts map[string]any) core.PoetrydbEntity {
		return entity.NewLinecountEntity(client, entopts)
	}
	core.NewPoemcountEntityFunc = func(client *core.PoetrydbSDK, entopts map[string]any) core.PoetrydbEntity {
		return entity.NewPoemcountEntity(client, entopts)
	}
	core.NewRandomEntityFunc = func(client *core.PoetrydbSDK, entopts map[string]any) core.PoetrydbEntity {
		return entity.NewRandomEntity(client, entopts)
	}
	core.NewTitleEntityFunc = func(client *core.PoetrydbSDK, entopts map[string]any) core.PoetrydbEntity {
		return entity.NewTitleEntity(client, entopts)
	}
	core.NewTitleabEntityFunc = func(client *core.PoetrydbSDK, entopts map[string]any) core.PoetrydbEntity {
		return entity.NewTitleabEntity(client, entopts)
	}
}

// Constructor re-exports.
var NewPoetrydbSDK = core.NewPoetrydbSDK
var TestSDK = core.TestSDK
var NewContext = core.NewContext
var NewSpec = core.NewSpec
var NewResult = core.NewResult
var NewResponse = core.NewResponse
var NewOperation = core.NewOperation
var MakeConfig = core.MakeConfig

// No-arg convenience constructors. Go has no default-argument syntax,
// so these aliases let callers write `sdk.New()` / `sdk.Test()`
// instead of `sdk.NewPoetrydbSDK(nil)` / `sdk.TestSDK(nil, nil)`
// for the common no-options case.
func New() *PoetrydbSDK  { return NewPoetrydbSDK(nil) }
func Test() *PoetrydbSDK { return TestSDK(nil, nil) }
var NewBaseFeature = feature.NewBaseFeature
var NewTestFeature = feature.NewTestFeature
