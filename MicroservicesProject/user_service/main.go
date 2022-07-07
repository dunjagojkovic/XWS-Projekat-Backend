package main

import (
	"fmt"
	"userS/config"
)

func main() {

	config := config.NewConfig()
	fmt.Println(config.UserDBHost)
	server := NewServer(config)
	server.Start()

}
