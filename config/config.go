package config

import (
	"log"

	"github.com/jinzhu/configor"
)

// Get configs
var Get = loadConfig()

// Command store the commands data
type Command struct {
	Prefix string `required:"true"`
}

// Configuration implements all configurations
type Configuration struct {
	Command Command
}

func loadConfig() Configuration {
	var conf Configuration
	if err := configor.Load(&conf, "./config/config.json"); err != nil {
		log.Println(err)
	}
	return conf
}
