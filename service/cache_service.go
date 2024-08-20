package service

import (
	"encoding/json"
	"user-backend/models"
	"user-backend/redis"
	"user-backend/utils"
)

type CacheService interface {
	AddSmsOTP(phone string, type_ string) (string, error)
	GetSmsOTP(phone string) (*models.SmsOTP, error)
	DelKey(key string) error
}

type cacheService struct {
	client redis.Client
	utils  utils.Utils
}

func NewCacheService(client redis.Client, utils utils.Utils) CacheService {
	return &cacheService{
		client: client,
		utils:  utils,
	}
}

func (c *cacheService) AddSmsOTP(phone string, type_ string) (string, error) {

	otp := c.utils.GenerateRandomNumber()

	if len(otp) != 6 {
		stdErr := models.GetStandardError("", models.INTERNAL_SERVER_ERROR, models.INTERNAL_SERVER_ERROR_MESSAGE)
		return "", stdErr
	}

	otpValue := models.SmsOTP{
		Phone: phone,
		OTP:   otp,
	}
	data, err := json.Marshal(otpValue)
	if err != nil {
		return "", err
	}

	err = c.client.Set(type_+phone, data)
	if err != nil {
		return "", err
	}
	return otp, nil
}

func (c *cacheService) GetSmsOTP(phone string) (*models.SmsOTP, error) {
	cacheOtp, err := c.client.Get(phone)
	if err != nil {
		return nil, err
	}

	if len(cacheOtp) == 0 {
		return nil, nil
	}
	var otpValue models.SmsOTP
	err = json.Unmarshal([]byte(cacheOtp), &otpValue)
	if err != nil {
		stdErr := models.GetStandardError("", models.INTERNAL_SERVER_ERROR, err.Error())
		return nil, stdErr
	}
	return &otpValue, nil
}

func (c *cacheService) DelKey(key string) error {
	err := c.client.Del(key)
	if err != nil {
		return err
	}
	return nil
}
