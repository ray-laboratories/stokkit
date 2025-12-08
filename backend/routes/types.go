package routes

type Response struct {
	Data any `json:"data"`
}

func NewResponse(data any) Response {
	return Response{
		Data: data,
	}
}

func NewResponseFromError(err error) Response {
	return Response{
		Data: err.Error(),
	}
}
