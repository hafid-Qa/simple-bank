package utils

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DB_HOST      string `mapstructure:"DB_HOST"`
	DB_PORT      int32  `mapstructure:"DB_PORT"`
	DB_USER      string `mapstructure:"DB_USER"`
	DB_PASSWORD  string `mapstructure:"DB_PASSWORD"`
	DB_NAME      string `mapstructure:"DB_NAME"`
	TEST_DB_HOST string `mapstructure:"TEST_DB_HOST"`
	TEST_DB_PORT int32  `mapstructure:"TEST_DB_PORT"`
	DB_SSLMODE   string `mapstructure:"DB_SSLMODE"`

	GIN_MODE      string `mapstructure:"GIN_MODE"`
	PORT          int32  `mapstructure:"PORT"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`

	DB_DRIVER      string `mapstructure:"DB_DRIVER"`
	DB_SOURCE      string `mapstructure:"DB_SOURCE"`
	TEST_DB_SOURCE string `mapstructure:"TEST_DB_SOURCE"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	if err != nil {
		return
	}

	// Set db
	config.DB_SOURCE = fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=%s",
		config.DB_USER, config.DB_PASSWORD, config.DB_HOST, config.DB_PORT, config.DB_NAME, config.DB_SSLMODE)

	config.TEST_DB_SOURCE = fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=%s",
		config.DB_USER, config.DB_PASSWORD, config.TEST_DB_HOST, config.TEST_DB_PORT, config.DB_NAME, config.DB_SSLMODE)
	config.ServerAddress = fmt.Sprintf("%s:%d", config.DB_HOST, config.PORT)

	return
}
