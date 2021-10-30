package config

import "github.com/spf13/viper"

type config struct {
	DatabaseURL string `mapstructure:"database_url"`
	LogLevel    string `mapstructure:"log_level"`
}

var Config config

func Load() error {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.SetEnvPrefix("SEARCH_API")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return viper.Unmarshal(&Config)
}
