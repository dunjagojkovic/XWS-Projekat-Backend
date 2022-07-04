package main

import "fmt"

func main() {

	config := NewConfig()
	fmt.Println(config.MessageDBHost)
	server := NewServer(config)
	server.Start()

}
