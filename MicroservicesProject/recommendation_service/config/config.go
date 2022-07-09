package config

import (
	"os"
)

type Config struct {
	Port                     string
	RecommendationDBHost     string
	RecommendationDBUsername string
	RecommendationDBPassword string
	RecommendationDatabase   string
	DBNeo4jVersion           string
	RecommendationDBPort     string
}

func NewConfig() *Config {
	return &Config{
		Port:/*  "8000",   */ os.Getenv("RECOMMENDATION_SERVICE_PORT"),
		RecommendationDBHost:/*"localhost", */ os.Getenv("RECOMMENDATION_DB_HOST"),
		RecommendationDBPort:/*"27017",    */ os.Getenv("RECOMMENDATION_DB_PORT"),
		RecommendationDBUsername: os.Getenv("RECOMMENDATION_DB_USERNAME"),
		RecommendationDBPassword: os.Getenv("RECOMMENDATION_DB_PASSWORD"),
		RecommendationDatabase:   os.Getenv("RECOMMENDATION_DATABASE"),
		DBNeo4jVersion:           os.Getenv("DB_NEO4J_VERSION"),
	}
}
