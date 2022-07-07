package main

import (
	"fmt"
	"postS/config"
)

func main() {
	config := config.NewConfig()
	fmt.Println(config.PostDBHost)
	server := NewServer(config)
	server.Start()
}
