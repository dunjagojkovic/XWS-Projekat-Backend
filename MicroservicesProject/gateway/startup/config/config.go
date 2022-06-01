package config

import "os"

type Config struct {
	Port     string
	PostHost string
	PostPort string
	JobHost  string
	JobPort  string
}

func NewConfig() *Config {
	return &Config{
		Port:/* "8000",     */ os.Getenv("GATEWAY_PORT"),
		PostHost:/*"post_service",*/ os.Getenv("POST_SERVICE_HOST"),
		PostPort:/*"8000",       */ os.Getenv("POST_SERVICE_PORT"),
		JobHost: os.Getenv("JOB_SERVICE_HOST"),
		JobPort: os.Getenv("JOB_SERVICE_PORT"),
	}
}
