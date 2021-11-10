package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

var (
	C = initialize()
	//o = sync.Once{}
)

type config struct {
	App    App    `mapstructure:"app" json:"app" yaml:"app"`
	Server Server `mapstructure:"server" json:"server" yaml:"server"`
	Mysql  Mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Redis  Redis  `mapstructure:"redis" json:"redis" yaml:"redis"`
}

func initialize() *config {
	var c config
	fmt.Println(11111)
	//o.Do(func() {
	//
	//})
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	if err := viper.Unmarshal(&c); err != nil {
		log.Fatal(err)
	}
	return &c
}
