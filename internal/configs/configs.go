package configs

import "github.com/spf13/viper"

type Config struct {
	Port string
}

func NewConfig() (*Config, error) {

	if err := fetchConfig(); err != nil {
		return nil, err
	}

	return &Config{
		Port: viper.GetString("server.port"),
	}, nil
}

func fetchConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	return viper.ReadInConfig()
}
