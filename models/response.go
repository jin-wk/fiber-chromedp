package models

// Response : type response struct
type ResponseModel struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Response(code int, message string, data interface{}) ResponseModel {
	return ResponseModel{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
