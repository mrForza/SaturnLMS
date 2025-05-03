package configs

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type DatabaseConfig struct {
	Host               string        `yaml:"host"`
	Port               uint16        `yaml:"port"`
	Database           string        `yaml:"database"`
	User               string        `yaml:"user"`
	Password           string        `yaml:"password"`
	MaxIdleConnections int           `yaml:"max_idle_connections"`
	MaxOprnConnections int           `yaml:"max_oprn_connections"`
	MaxIdleTime        time.Duration `yaml:"max_idle_time"`
	MaxLifeTime        time.Duration `yaml:"max_life_time"`
}

func (cfg DatabaseConfig) GetShortDescription() string {
	return fmt.Sprintf(
		"DatabaseConfig:\n\t"+"Host: %s\n\tPort: %d",
		cfg.Host, cfg.Port,
	)
}

func (config *DatabaseConfig) ToDataSourceName() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Europe/Moscow",
		config.Host, config.User, config.Password, config.Database, config.Port,
	)
}

func LoadDatabaseConfig(configPath string) *Configurable {
	var config DatabaseConfig
	var configurable Configurable

	if err := cleanenv.ReadConfig(configPath, &config); err != nil {
		log.Fatalf("Cannot read config due to error: %s", err)
	}

	configurable = config
	return &configurable
}

var DbConfig = LoadConfig(os.Getenv("CONFIG_PATH")+"/db.yaml", "DatabaseConfig", LoadDatabaseConfig)
