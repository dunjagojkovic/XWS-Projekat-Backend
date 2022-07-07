package config

import "os"

type Config struct {
	Port            string
	PostDBHost      string
	PostDBPort      string
	LogsFolder      string
	InfoLogsFile    string
	DebugLogsFile   string
	ErrorLogsFile   string
	SuccessLogsFile string
	WarningLogsFile string
}

func NewConfig() *Config {
	return &Config{
		Port:/*  "8000",   */ os.Getenv("POST_SERVICE_PORT"),
		PostDBHost:/*"localhost", */ os.Getenv("POST_DB_HOST"),
		PostDBPort:/*"27017",    */ os.Getenv("POST_DB_PORT"),
		LogsFolder:      "logs",
		InfoLogsFile:    "/info.log",
		DebugLogsFile:   "/debug.log",
		ErrorLogsFile:   "/error.log",
		SuccessLogsFile: "/success.log",
		WarningLogsFile: "/warning.log",
	}
}
