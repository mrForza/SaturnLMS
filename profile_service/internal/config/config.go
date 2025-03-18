package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	HttpServer HttpServerConfig `yaml:"http_server"`
	Database   DatabaseConfig   `yaml:"database"`
	Logging    LoggingConfig    `yaml:"logging"`
}

type HttpServerConfig struct {
	Host         string        `yaml:"host"`
	Port         uint16        `yaml:"port"`
	Timeout      time.Duration `yaml:"timeout"`
	IddleTimeout time.Duration `yaml:"iddle_timeout"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     uint16 `yaml:"port"`
	Database string `yaml:"database"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type LoggingConfig struct {
	Level string `yaml:"level"`
}

func Load() *Config {
	configPath := os.Getenv("CONFIG_PATH")

	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file '%s' is not exists", configPath)
	}

	var config Config
	if err := cleanenv.ReadConfig(configPath, &config); err != nil {
		log.Fatalf("cannot read config due to error: %s", err)
	}

	return &config
}

func (config *DatabaseConfig) ToPostgresUrl() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Europe/Moscow",
		config.Host, config.User, config.Password, config.Database, config.Port,
	)
}
