package models

type APILimiterDto struct {
	UserIp   string
	Api      string
	Tries    int64
	MaxTries int
}

type SmsOTP struct {
	Phone string `json:"phone"`
	OTP   string `json:"otp"`
}
