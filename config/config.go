package config

import (
	"github.com/spf13/viper"
	"log"
)

var (
	C          = initialize()
	configName = "config"
	configType = "yaml"
	configPath = "./"
)

type config struct {
	App    App    `mapstructure:"app" json:"app" yaml:"app"`
	Server Server `mapstructure:"server" json:"server" yaml:"server"`
	Mysql  Mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Redis  Redis  `mapstructure:"redis" json:"redis" yaml:"redis"`
	Jwt    Jwt    `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
}

func initialize() *config {
	var c config
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	viper.AddConfigPath(configPath)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	if err := viper.Unmarshal(&c); err != nil {
		log.Fatal(err)
	}
	return &c
}
