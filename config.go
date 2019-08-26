package main

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type config struct {
	Test  string `yaml:"test"`
}

func (c *config) getConfig() {

	viper.SetConfigName("config")
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
		fmt.Printf("%+v", c)
	})
	viper.Unmarshal(c)
}

func main() {
	done := make(chan bool)
	var Config config
	Config.getConfig()

	fmt.Printf("%+v", Config)
	<-done
}
