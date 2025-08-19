package utils

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DbHost     string `mapstructure:"DB_HOST"`
	DbPort     int32  `mapstructure:"DB_PORT"`
	DbUser     string `mapstructure:"DB_USER"`
	DbPassword string `mapstructure:"DB_PASSWORD"`
	DbName     string `mapstructure:"DB_NAME"`
	TestDbHost string `mapstructure:"TEST_DB_HOST"`
	TestDbPort int32  `mapstructure:"TEST_DB_PORT"`
	DbSslmode  string `mapstructure:"DB_SSLMODE"`

	GinMode       string `mapstructure:"GIN_MODE"`
	Port          int32  `mapstructure:"PORT"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`

	DbDriver            string        `mapstructure:"DB_DRIVER"`
	DbSource            string        `mapstructure:"DB_SOURCE"`
	TestDbSource        string        `mapstructure:"TEST_DB_SOURCE"`
	TokenSymmetricKey   string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
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
	config.DbSource = fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=%s",
		config.DbUser, config.DbPassword, config.DbHost, config.DbPort, config.DbName, config.DbSslmode)

	config.TestDbSource = fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=%s",
		config.DbUser, config.DbPassword, config.TestDbHost, config.TestDbPort, config.DbName, config.DbSslmode)
	config.ServerAddress = fmt.Sprintf("%s:%d", "0.0.0.0", config.Port)

	return
}
