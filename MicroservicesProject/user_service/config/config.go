package config

import "os"

type Config struct {
	Port            string
	UserDBHost      string
	UserDBPort      string
	LogsFolder      string
	InfoLogsFile    string
	DebugLogsFile   string
	ErrorLogsFile   string
	SuccessLogsFile string
	WarningLogsFile string
}

func NewConfig() *Config {
	return &Config{
		Port:            os.Getenv("USER_SERVICE_PORT"),
		UserDBHost:      os.Getenv("USER_DB_HOST"),
		UserDBPort:      os.Getenv("USER_DB_PORT"),
		LogsFolder:      "logs",
		InfoLogsFile:    "/info.log",
		DebugLogsFile:   "/debug.log",
		ErrorLogsFile:   "/error.log",
		SuccessLogsFile: "/success.log",
		WarningLogsFile: "/warning.log",
	}
}
