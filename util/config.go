package config

import (
	"github.com/spf13/viper"
)

type DbConfig struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Dbname   string `mapstructure:"dbname"`
}
type ServerConfig struct {
	Port string `mapstructure:"port"`
}
type Config struct {
	Db     DbConfig     `mapstructure:"db"`
	Server ServerConfig `mapstructure:"server"`
}

func LoadConfig() (Config, error) {
	vp := viper.New()
	var config Config
	vp.SetConfigName("config")
	vp.SetConfigType("json")
	vp.AddConfigPath(".")
	vp.AddConfigPath("./util")
	err := vp.ReadInConfig()
	if err != nil {
		return Config{}, err
	}
	err = vp.Unmarshal(&config)
	if err != nil {
		return Config{}, err
	}

	return config, err
}
