package response

type Response struct {
	Status string      `json:"status"`
	Code   int         `json:"code"`
	Data   interface{} `json:"data"`
}

func NewSuccessResponse(code int, data interface{}) Response {
	return Response{
		Status: "success",
		Code:code,
		Data: data,
	}
}

func NewErrorResponse(code int, data interface{}) Response {
	return Response{
		Status: "error",
		Code: code,
		Data: data,
	}
}