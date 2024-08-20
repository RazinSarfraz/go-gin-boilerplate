package rest

import "user-backend/logger"

type StandardResponse struct {
	Result  bool        `json:"result"`
	Code    uint        `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewStandardResponse(session string, result bool, code uint, msg string, data interface{}) *StandardResponse {

	// Nil data would be set to empty
	if data == nil {
		data = ""
	}
	res := &StandardResponse{
		Result:  result,
		Code:    code,
		Message: msg,
		Data:    data,
	}
	logger.LogDebug2("Response:", session, res)
	return res
}
