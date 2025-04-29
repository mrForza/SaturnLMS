package configs

import (
	"fmt"
	"log"
	"os"
)

type Configurable interface {
	GetShortDescription() string
}

func LoadConfig(configPath string, configName string, loader func(string) *Configurable) *Configurable {
	fmt.Println(configPath)
	if configPath == "" {
		log.Fatalf("CONFIG_PATH for %s is not set", configName)
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config file for '%s' is not exists", configName)
	}

	return loader(configPath)
}
