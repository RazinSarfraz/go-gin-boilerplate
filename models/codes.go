package models

const (
	INTERNAL_SERVER_ERROR = 500

	SUCCESS                  = 1000
	INVALID_INPUT            = 1001
	INVALID_TOKEN            = 1002
	TOKEN_EXPIRED            = 1003
	PHONE_ALREADY_REGISTERED = 1004
)

const (
	INTERNAL_SERVER_ERROR_MESSAGE = "Internal server error"

	SUCCESS_MESSAGE                  = "Success"
	INVALID_INPUT_MESSAGE            = "Invalid input"
	INVALID_TOKEN_MESSAGE            = "Invalid token"
	TOKEN_EXPIRED_MESSAGE            = "Token Expired"
	PHONE_ALREADY_REGISTERED_MESSAGE = "Phone Number already registered"
	PING                             = "Ping Success"
)
