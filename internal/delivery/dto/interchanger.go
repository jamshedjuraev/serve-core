package dto

import "encoding/json"

type Response struct {
	StatusCode int         `json:"status_code"`
	Error      string      `json:"error"`
	Data       interface{} `json:"data"`
}

func NewSuccessResponse(statusCode int, data interface{}) []byte {
	v, err := json.Marshal(&data)
	if err != nil {
		v, _ := json.Marshal(&Response{StatusCode: 500, Error: "Cannot convert struct to json. Error: " + err.Error()})
		return v
	}
	return v
}

func NewErrorResponse(statusCode int, errMsg string) []byte {
	v, _ := json.Marshal(&Response{
		StatusCode: statusCode,
		Error:      errMsg,
	})
	return v
}
