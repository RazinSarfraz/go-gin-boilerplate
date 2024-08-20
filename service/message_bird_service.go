package service

import (
	"user-backend/conf"
	"user-backend/logger"
	messagebirdModels "user-backend/models/messageBird"
	"user-backend/repository"

	messagebird "github.com/messagebird/go-rest-api"
	"github.com/messagebird/go-rest-api/sms"
)

type MessageBirdService interface {
	SendSms(sms *messagebirdModels.SMS) error
}

type messageBirdService struct {
	store  repository.Store
	config conf.Config
}

func NewMessageBirdService(store repository.Store, config conf.Config) MessageBirdService {
	return &messageBirdService{
		store:  store,
		config: config,
	}
}

func (s *messageBirdService) SendSms(authSms *messagebirdModels.SMS) error {
	accessKey := s.config.MessageBird
	params := &sms.Params{
		Type:       "sms",
		DataCoding: "unicode",
	}

	client := messagebird.New(accessKey)

	message, err := sms.Create(
		client,
		"SalaamBank",
		[]string{authSms.Phone},
		authSms.Body,
		params,
	)
	if err != nil {
		logger.LogError(err.Error(), err)
	}
	logger.LogInfo(message.Body, "")
	return nil
}
