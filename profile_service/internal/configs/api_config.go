package configs

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type HttpApiConfig struct {
	Host         string        `yaml:"host"`
	Port         uint16        `yaml:"port"`
	Timeout      time.Duration `yaml:"timeout"`
	IddleTimeout time.Duration `yaml:"iddle_timeout"`
}

func (cfg HttpApiConfig) GetShortDescription() string {
	return fmt.Sprintf(
		"HttpApiConfig:\n\t"+"Host: %s\n\tPort: %d",
		cfg.Host, cfg.Port,
	)
}

func LoadHttpApiConfig(configPath string) *Configurable {
	var config HttpApiConfig
	var configurable Configurable

	if err := cleanenv.ReadConfig(configPath, &config); err != nil {
		log.Fatalf("Cannot read config due to error: %s", err)
	}

	configurable = config
	return &configurable
}

var ApiConfig = LoadConfig(os.Getenv("CONFIG_PATH")+"/api.yaml", "HttpApiConfig", LoadHttpApiConfig)
