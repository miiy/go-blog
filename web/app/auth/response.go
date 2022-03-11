package auth

import "github.com/miiy/go-web/pkg/response"

func SuccessResponse(v interface{}) *response.Response {
	resp := &response.Response{
		Code:       response.CodeSuccess,
		Message:    response.MsgSuccess,
		Data:       v,
	}
	return resp
}

func BadRequestResponse(v interface{}) *response.Response {
	resp := &response.Response{
		Code:       response.CodeBadRequest,
		Message:    response.MsgBadRequest,
		Data:       v,
	}
	return resp
}

func ErrorResponse(v interface{}) *response.Response {
	resp := &response.Response{
		Code:       response.CodeError,
		Message:    response.MsgError,
		Data:       v,
	}
	return resp
}

