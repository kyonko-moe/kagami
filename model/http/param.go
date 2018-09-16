package http

// ParamRegister is.
type ParamRegister struct {
	Name string `form:"name"`
}

type ParamLocate struct {
}

type ParamConnect struct {
	Name string `form:"name"`
}
