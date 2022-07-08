package main

import (
	"messageS/config"
)

func main() {

	config := config.NewConfig()
	server := NewServer(config)
	server.Start()

}
