package core

var UtilityRegistrar func(u *Utility)

var NewBaseFeatureFunc func() Feature

var NewTestFeatureFunc func() Feature

var NewAuthorEntityFunc func(client *PoetrydbSDK, entopts map[string]any) PoetrydbEntity

var NewAuthorabEntityFunc func(client *PoetrydbSDK, entopts map[string]any) PoetrydbEntity

var NewCombinedSearchEntityFunc func(client *PoetrydbSDK, entopts map[string]any) PoetrydbEntity

var NewCombinedSearchWithFieldEntityFunc func(client *PoetrydbSDK, entopts map[string]any) PoetrydbEntity

var NewLineEntityFunc func(client *PoetrydbSDK, entopts map[string]any) PoetrydbEntity

var NewLinecountEntityFunc func(client *PoetrydbSDK, entopts map[string]any) PoetrydbEntity

var NewPoemcountEntityFunc func(client *PoetrydbSDK, entopts map[string]any) PoetrydbEntity

var NewRandomEntityFunc func(client *PoetrydbSDK, entopts map[string]any) PoetrydbEntity

var NewTitleEntityFunc func(client *PoetrydbSDK, entopts map[string]any) PoetrydbEntity

var NewTitleabEntityFunc func(client *PoetrydbSDK, entopts map[string]any) PoetrydbEntity

