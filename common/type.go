package common

type HTMLRenderData struct {
	Name string
	Data *HTMLData
}

type HTMLData struct {
	Status int
	RenderData interface{}
	Message string
	Err *ErrorData
}

type ErrorData struct {
	ErrorString string
	ErrorStack []string
	Status int
}

func (e *ErrorData) Error() string {
	return e.ErrorString
}
