package helper

import "strings"

//Response struct json
type Response struct {
	Code    int         `json:"coce"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
	Data    interface{} `json:"data"`
}

//Struct to give Empty Object
type EmptyObj struct {
}

//function Build Response
func BuildResponse(code int, message string, data interface{}) Response {
	res := Response{
		Code:    code,
		Message: message,
		Error:   nil,
		Data:    data,
	}
	return res
}

//function if the response is error
func BuildErrorResponse(code int, message string, err string, data interface{}) Response {
	splittedError := strings.Split(err, "\n")
	res := Response{
		Code:    code,
		Message: message,
		Error:   splittedError,
		Data:    data,
	}
	return res
}
