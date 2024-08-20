package models

import "time"

type User struct {
	UserID      string `gorm:"primaryKey;type:uuid;DEFAULT uuid_generate_v4()" json:"user_id"`
	Phone       string `json:"phone,omitempty"`
	FirstName   string `json:"firstName,omitempty"`
	LastName    string `json:"lastName,omitempty"`
	Status      string `json:"status,omitempty"`
	OtpVerified bool   `json:"otpVerified,omitempty"`
	IBAN        string `json:"iban,omitempty"`
	Password    string `json:"password,omitempty"`
	Salt        string `json:"salt,omitempty"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
