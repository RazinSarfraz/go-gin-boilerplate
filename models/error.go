package models

import (
	"fmt"
	"user-backend/logger"
)

type StandardError struct {
	Code        uint
	ActualError error
	Line        string
	Message     string
}

func (s StandardError) Error() string {
	errStr := fmt.Sprintf("Code : %v Line:%v \n Error:%v \n Message: %v", s.Code, s.Line, s.ActualError, s.Message)
	return errStr
}

func GetStandardError(session string, code uint, message string) *StandardError {
	err := &StandardError{
		Code:    code,
		Message: message,
	}
	logger.LogError(session, err)

	return err
}
