package util

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port     string `mapstructure:"PORT"`
	HostUrl  string `mapstructure:"HOST_URL"`
	UrlPath  string `mapstructure:"URL_PATH"`
	BotToken string `mapstructure:"TELEGRAM_BOT_TOKEN"`
	MongoURI string `mapstructure:"MONGO_URI"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
