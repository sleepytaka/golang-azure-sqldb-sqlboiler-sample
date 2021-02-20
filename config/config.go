package config

import (
	"github.com/spf13/viper"
	"strings"
)

type Config struct {
	Database Database `yaml:"db"`
}

type Database  struct {
	Dbname     string `yaml:"dbname"`
	Host     string `yaml:"host"`
	Port     int `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	SslMode string `yaml:"sslMode"`
	Schema string `yaml:"schema"`
}

func New()(*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config/")

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig();err != nil {
		return nil, err
	}

	var c Config
	if err := viper.Unmarshal(&c);err != nil {
		return nil, err
	}

	return &c, nil
}

