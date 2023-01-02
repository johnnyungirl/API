package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Dbconfig struct {
	Address  string `mapstructure:"address"`
	Port     string `mapstructure:"port"`
	Password string `mapstructure:"password"`
	User     string `mapstructure:"user"`
}
type ServerConfig struct {
	Port string `mapstructure:"port"`
}
type Config struct {
	DB     Dbconfig     `mapstructure:"db"`
	Server ServerConfig `mapstructure:"server"`
}

var vp *viper.Viper

func LoadConfig() (Config, error) {
	vp = viper.New()
	var config Config
	vp.SetConfigName("config")
	vp.SetConfigType("json")
	vp.AddConfigPath(".")
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
func main() {
	config, err := LoadConfig()
	if err != nil {
		log.Fatalln("config no found")
	}
	fmt.Println(config)
}
