package config

import "os"

type Config struct {
	Port            string
	MessageDBHost   string
	MessageDBPort   string
	LogsFolder      string
	InfoLogsFile    string
	DebugLogsFile   string
	ErrorLogsFile   string
	SuccessLogsFile string
	WarningLogsFile string
}

func NewConfig() *Config {
	return &Config{
		Port:            os.Getenv("MESSAGE_SERVICE_PORT"),
		MessageDBHost:   os.Getenv("MESSAGE_DB_HOST"),
		MessageDBPort:   os.Getenv("MESSAGE_DB_PORT"),
		LogsFolder:      "logs",
		InfoLogsFile:    "/info.log",
		DebugLogsFile:   "/debug.log",
		ErrorLogsFile:   "/error.log",
		SuccessLogsFile: "/success.log",
		WarningLogsFile: "/warning.log",
	}
}
