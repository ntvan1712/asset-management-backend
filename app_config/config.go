package app_config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

const (
	configFolder = "./app_config"
)

var appConfig *AppConfig

type AppConfig struct {
	FiberServerConfig    FiberServerConfig    `mapstructure:"fiber_server"`
	RabbitMqConfig       RabbitMQConfig       `mapstructure:"rabbitmq"`
	LoggerConfig         LoggerConfig         `mapstructure:"logger"`
	MinioConfig          MinioConfig          `mapstructure:"minio"`
	GRPCConnectionConfig GRPCConnectionConfig `mapstructure:"grpc_connection"`
	PostgresConfig       PostgresConfig       `mapstructure:"postgres"`
	AuthConfig           AuthConfig           `mapstructure:"auth"`
}

func GetAppConfig() *AppConfig {
	if appConfig == nil {
		log.Fatal("[AppConfig] appConfig == nil")
	}
	return appConfig
}

func LoadAppConfig(configFileName string) {
	configFilePath := fmt.Sprintf("%s/%s", configFolder, configFileName)
	viper.SetConfigFile(configFilePath)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Failed to read config file", err)
	}
	err = viper.Unmarshal(&appConfig)
	if err != nil {
		log.Fatal("Failed to load config", err)
	}
}
