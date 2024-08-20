package service

import "user-backend/conf"

type ConfigService interface {
	GetConfig() *conf.Config
}
type configService struct {
}

func NewConfigService() ConfigService {
	return &configService{}
}

func (c *configService) GetConfig() *conf.Config {

	return conf.GetConfig()

}
