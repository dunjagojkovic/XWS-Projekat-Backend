package main

import (
	"postS/config"
)

func main() {
	config := config.NewConfig()
	server := NewServer(config)
	server.Start()
}
