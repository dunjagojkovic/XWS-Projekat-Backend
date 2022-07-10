package main

import (
	"notificationS/config"
)

func main() {
	config := config.NewConfig()
	server := NewServer(config)
	server.Start()
}
