package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type EnvConfig struct {
	PORT string `mapstructure:"PORT"`
}

func NewEnvConfig(log *logrus.Logger) (*EnvConfig, error) {
	var config EnvConfig
	viper.AddConfigPath("../../")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Errorf("error while reading config: %s", err.Error())
		return nil, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Errorf("error while unmarshall config: %s", err.Error())
		return nil, err
	}

	return &config, nil
}
