package main

import "followS/config"

func main() {
	config := config.NewConfig()
	server := NewServer(config)
	server.Start()
}
