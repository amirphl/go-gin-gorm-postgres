package helper

import "strings"

// DefaultErrMsg ...
var DefaultErrMsg = "Failed to process the request"

// Response representation
type Response struct {
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

// BuildResp ...
func BuildResp(message string, data interface{}) Response {
	return Response{
		Message: message,
		Errors:  nil,
		Data:    data,
	}
}

// BuildErrResp ...
func BuildErrResp(message string, errs string, data interface{}) Response {
	return Response{
		Message: message,
		Errors:  strings.Split(errs, "\n"),
		Data:    data,
	}
}
