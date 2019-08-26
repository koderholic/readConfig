# readConfig
This is a simple program to read values from a config.yaml file built in Golang.

This program reads config data from a conf.yaml file and loads it into a defined struct, using viper and fsNotify.

fsNotify watches the config.yaml file for changes in real time

Set up : 
make sure to have go installed and setup on your computer, then run "go build" to build the executable, afterwhich you can run the executable. This will print to screen the value you have set in your config.yaml
