package server

import (
	"os"
)

type Config struct {
	Port        string
	JWTSecret   string
	DatabaseUrl string
}

func (config *Config) GetPort() string {
	return config.Port
}

func (config *Config) GetJWTSecret() string {
	return config.JWTSecret
}

func (config *Config) GetDatabaseUrl() string {
	return config.DatabaseUrl
}

func NewConfig() *Config {
	return &Config{
		Port:        os.Getenv("PORT"),
		JWTSecret:   os.Getenv("JWT_SECRET"),
		DatabaseUrl: os.Getenv("DATABASE_URL"),
	}
}

func GetConfig() *Config {
	return NewConfig()
}
