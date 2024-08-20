package conf

import (
	"encoding/json"
	"os"
	"sync"
	"user-backend/logger"
)

var (
	configPath     = "."
	configFileName = "config.json"
)

type Config struct {
	Logger         logger.LoggerConfig     `json:"logger"`
	PostDataSource PostgreDataSourceConfig `json:"postgreDataSource"`
	MessageBird    string                  `json:"messageBird"`
	Redis          RedisConfig             `json:"redis"`
	RestServer     RestServerConfig        `json:"restServer"`
	JwtSecret      string                  `json:"jwtSecret"`
}

type PostgreDataSourceConfig struct {
	DriverName        string `json:"driverName"`
	Addr              string `json:"addr"`
	Port              string `json:"port"`
	Database          string `json:"database"`
	User              string `json:"user"`
	Password          string `json:"password"`
	EnableAutoMigrate bool   `json:"enableAutoMigrate"`
}

type RestServerConfig struct {
	Addr string `json:"addr"`
}

type RedisConfig struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
}

var config Config
var configOnce sync.Once

func GetConfig() *Config {
	configOnce.Do(func() {
		bytes, err := os.ReadFile(configPath + "/" + configFileName)
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(bytes, &config)
		if err != nil {
			panic(err)
		}
	})
	return &config
}
