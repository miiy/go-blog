package response

var (
	CodeSuccess    = 0
	CodeError      = 1
	CodeBadRequest = 3

	MsgSuccess    = "Success."
	MsgError      = "Error."
	MsgBadRequest = "Bad request."
)

type Response struct{
	Code       int         `json:"code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

func NewSuccess(v interface{}) *Response {
	resp := &Response{
		Code:       CodeSuccess,
		Message:    MsgSuccess,
		Data:       v,
	}
	return resp
}

func NewError(v interface{}) *Response {
	resp := &Response{
		Code:       CodeError,
		Message:    MsgError,
		Data:       v,
	}
	return resp
}
