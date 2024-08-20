package models

type DecodeJWTClaims struct {
	ExpiredAt float64 `json:"exp"`
	UserId    string  `json:"uid"`
	Phone     string  `json:"phoneNumber"`
	TokenType string  `json:"tokenType"`
	LoginType string  `json:"loginType"`
}
