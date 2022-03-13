package helper

import "strings"

// Response representation
type Response struct {
	Status  uint16      `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

// Builds a HTTP response
// TODO Why not returning a pointer?
func BuildResp(status uint16, message string, data interface{}) Response {
	return Response{
		Status:  status,
		Message: message,
		Errors:  nil,
		Data:    data,
	}
}

// Builds a HTTP response indicates some errors
func BuildErrResp(status uint16, message string, errs string, data interface{}) Response {
	return Response{
		Status:  status,
		Message: message,
		Errors:  strings.Split(errs, "\n"),
		Data:    data,
	}
}
