package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Env      string
	HTTPAddr string

	DB DBConfig

	RedisAddr string
}

type DBConfig struct {
	Host     string
	Port     int
	Name     string
	User     string
	Password string
}

func Load() (Config, error) {
	cfg := Config{
		Env:       getenv("APP_ENV", "dev"),
		HTTPAddr:  getenv("HTTP_ADDR", ":8080"),
		RedisAddr: getenv("REDIS_ADDR", "localhost:6379"),
		DB: DBConfig{
			Host:     getenv("DB_HOST", "localhost"),
			Name:     getenv("DB_NAME", "budgetpilot"),
			User:     getenv("DB_USER", "budgetpilot"),
			Password: getenv("DB_PASSWORD", "budgetpilot"),
		},
	}

	portStr := getenv("DB_PORT", "5432")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return Config{}, fmt.Errorf("invalid DB_PORT %q: %w", portStr, err)
	}
	cfg.DB.Port = port

	if strings.ToLower(cfg.Env) == "prod" {
		if cfg.DB.Password == "" {
			return Config{}, fmt.Errorf("DB_PASSWORD must be set in prod")
		}
	}

	return cfg, nil
}

func getenv(k, fallback string) string {
	v := os.Getenv(k)
	if v == "" {
		return fallback
	}
	return v
}
