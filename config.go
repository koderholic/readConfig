package main

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type conf struct {
	Suffix  string `yaml:"suffix"`
	Keyword string `yaml:"keyword"`
}

func (c *conf) getConf() {

	viper.SetConfigName("conf")
	viper.AddConfigPath(".")
	viper.WatchConfig()
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		err := viper.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
		viper.Unmarshal(c)
	})
	viper.Unmarshal(c)
}

func main() {
	done := make(chan bool)
	var config conf
	go config.getConf()

	fmt.Println(config.Keyword)
	<-done
}
