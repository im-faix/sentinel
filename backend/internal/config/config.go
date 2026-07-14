package config

import (
	"os"
)

type Config struct {
	AppName    string
	AppVersion string
	Host       string
	Port       string
	LogLevel   string
}

func Load() *Config {
	return &Config{
		AppName:    getEnv("APP_NAME", "Sentinel"),
		AppVersion: getEnv("APP_VERSION", "0.1.0"),
		Host:       getEnv("HOST", "0.0.0.0"),
		Port:       getEnv("PORT", "8080"),
		LogLevel:   getEnv("LOG_LEVEL", "INFO"),
	}
}

func (c *Config) Address() string {
	return c.Host + ":" + c.Port
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
