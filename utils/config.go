package utils

import "github.com/spf13/viper"

type Config struct {
	DB_HOST      string `mapstructure:"DB_HOST"`
	DB_PORT      int32  `mapstructure:"DB_PORT"`
	DB_USER      string `mapstructure:"DB_USER"`
	DB_PASSWORD  string `mapstructure:"DB_PASSWORD"`
	DB_NAME      string `mapstructure:"DB_NAME"`
	TEST_DB_NAME string `mapstructure:"TEST_DB_NAME"`
	TEST_DB_HOST string `mapstructure:"TEST_DB_HOST"`
	TEST_DB_PORT int32  `mapstructure:"TEST_DB_PORT"`
	DB_SSLMODE   string `mapstructure:"DB_SSLMODE"`

	GIN_MODE string `mapstructure:"GIN_MODE"`
	PORT     int32  `mapstructure:"PORT"`

	DB_DRIVER string `mapstructure:"DB_DRIVER"`
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
	return
}
