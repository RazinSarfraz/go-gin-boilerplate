package service

import (
	"user-backend/conf"
	"user-backend/logger"
	"user-backend/redis"
	"user-backend/repository"
	postres "user-backend/repository/postgres"
	"user-backend/utils"

	"go.uber.org/zap"
)

type Container struct {
	ConfigService ConfigService
	Store         repository.Store
	Logger        *zap.Logger
	CacheService  CacheService
	JwtService    JwtService
}

func NewServiceContainer() *Container {
	configService := NewConfigService()
	config := conf.GetConfig()
	utils := utils.NewUtils()
	logger := logger.LoggerInit(config.Logger)
	store := postres.SharedStore()
	redisClient := redis.NewClient(config)
	cacheService := NewCacheService(redisClient, utils)
	jwtService := NewJWTService(cacheService, config)
	return &Container{
		ConfigService: configService,
		Logger:        logger,
		Store:         store,
		CacheService:  cacheService,
		JwtService:    jwtService,
	}
}
