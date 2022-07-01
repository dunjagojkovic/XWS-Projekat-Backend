package main

import (
	"os"
)

type Config struct {
	Port             string
	FollowDBHost     string
	FollowDBUsername string
	FollowDBPassword string
	FollowDatabase   string
	DBNeo4jVersion   string
	FollowDBPort     string
	UserHost         string
	UserPort         string
}

func NewConfig() *Config {
	return &Config{
		Port:/*  "8000",   */ os.Getenv("FOLLOW_SERVICE_PORT"),
		FollowDBHost:/*"localhost", */ os.Getenv("FOLLOW_DB_HOST"),
		FollowDBPort:/*"27017",    */ os.Getenv("FOLLOW_DB_PORT"),
		FollowDBUsername: os.Getenv("FOLLOW_DB_USERNAME"),
		FollowDBPassword: os.Getenv("FOLLOW_DB_PASSWORD"),
		FollowDatabase:   os.Getenv("FOLLOW_DATABASE"),
		DBNeo4jVersion:   os.Getenv("DB_NEO4J_VERSION"),
		UserHost:         os.Getenv("USER_SERVICE_HOST"),
		UserPort:         os.Getenv("USER_SERVICE_PORT"),
	}
}
