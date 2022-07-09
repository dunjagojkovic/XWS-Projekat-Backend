package config

import "os"

type Config struct {
	Port               string
	PostHost           string
	PostPort           string
	JobHost            string
	JobPort            string
	UserHost           string
	UserPort           string
	FollowHost         string
	FollowPort         string
	MessageHost        string
	MessagePort        string
	RecommendationHost string
	RecommendationPort string
}

func NewConfig() *Config {
	return &Config{
		Port:/* "8000",     */ os.Getenv("GATEWAY_PORT"),
		PostHost:/*"post_service",*/ os.Getenv("POST_SERVICE_HOST"),
		PostPort:/*"8000",       */ os.Getenv("POST_SERVICE_PORT"),
		JobHost:            os.Getenv("JOB_SERVICE_HOST"),
		JobPort:            os.Getenv("JOB_SERVICE_PORT"),
		UserHost:           os.Getenv("USER_SERVICE_HOST"),
		UserPort:           os.Getenv("USER_SERVICE_PORT"),
		FollowHost:         os.Getenv("FOLLOW_SERVICE_HOST"),
		FollowPort:         os.Getenv("FOLLOW_SERVICE_PORT"),
		MessageHost:        os.Getenv("MESSAGE_SERVICE_HOST"),
		MessagePort:        os.Getenv("MESSAGE_SERVICE_PORT"),
		RecommendationHost: os.Getenv("RECOMMENDATION_SERVICE_HOST"),
		RecommendationPort: os.Getenv("RECOMMENDATION_SERVICE_PORT"),
	}
}
