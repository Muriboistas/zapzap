package config

import (
	"log"
	"time"

	"github.com/jinzhu/configor"
)

// Configuration implements all config data
type Configuration struct {
	Command struct {
		Prefix string `required:"true"`
	}
	Whatsapp struct {
		TimeOutDuration time.Duration `default:"5"`
		SessionPath     string        `default:"./session"`
	}
	Qrcode struct {
		FileName    string `default:"./session"`
		Quality     string `default:"medium"`
		Size        uint   `default:"256"`
		GeneratePNG bool   `default:"true"`
		PrintOnCLI  bool   `default:"true"`
	}
}

// Get configs
var Get = loadConfig()

func loadConfig() Configuration {
	var conf Configuration
	if err := configor.Load(&conf, "./config/config.json"); err != nil {
		log.Println(err)
	}
	return conf
}
