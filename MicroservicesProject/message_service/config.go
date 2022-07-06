package main

import "os"

type Config struct {
	Port                        string
	MessageDBHost               string
	MessageDBPort               string
	NatsHost                    string
	NatsPort                    string
	NatsUser                    string
	NatsPass                    string
	CreateMessageCommandSubject string
	CreateMessageReplySubject   string
}

func NewConfig() *Config {
	return &Config{
		Port:                        os.Getenv("MESSAGE_SERVICE_PORT"),
		MessageDBHost:               os.Getenv("MESSAGE_DB_HOST"),
		MessageDBPort:               os.Getenv("MESSAGE_DB_PORT"),
		NatsHost:                    os.Getenv("NATS_HOST"),
		NatsPort:                    os.Getenv("NATS_PORT"),
		NatsUser:                    os.Getenv("NATS_USER"),
		NatsPass:                    os.Getenv("NATS_PASS"),
		CreateMessageCommandSubject: os.Getenv("CREATE_MESSAGE_COMMAND_SUBJECT"),
		CreateMessageReplySubject:   os.Getenv("CREATE_MESSAGE_REPLY_SUBJECT"),
	}
}
