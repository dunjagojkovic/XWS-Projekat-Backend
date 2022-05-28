package config

import "os"

type Config struct {
	Port     string
	PostHost string
	PostPort string
}

func NewConfig() *Config {
	return &Config{
		Port:/* "8000",     */ os.Getenv("GATEWAY_PORT"),
		PostHost:/*"post_service",*/ os.Getenv("POST_SERVICE_HOST"),
		PostPort:/*"8000",       */ os.Getenv("POST_SERVICE_PORT"),
	}
}
