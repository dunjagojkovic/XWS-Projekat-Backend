package main

import (
	"gateway/startup"
	"gateway/startup/config"
)

func main() {
	config := config.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
