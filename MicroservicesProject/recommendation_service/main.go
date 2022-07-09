package main

import "recommendationS/config"

func main() {

	config := config.NewConfig()
	server := NewServer(config)
	server.Start()
}
