package core

type PoetrydbError struct {
	IsPoetrydbError bool
	Sdk              string
	Code             string
	Msg              string
	Ctx              *Context
	Result           any
	Spec             any
}

func NewPoetrydbError(code string, msg string, ctx *Context) *PoetrydbError {
	return &PoetrydbError{
		IsPoetrydbError: true,
		Sdk:              "Poetrydb",
		Code:             code,
		Msg:              msg,
		Ctx:              ctx,
	}
}

func (e *PoetrydbError) Error() string {
	return e.Msg
}
