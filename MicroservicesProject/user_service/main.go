package main

import "fmt"

func main() {

	config := NewConfig()
	fmt.Println(config.UserDBHost)
	server := NewServer(config)
	server.Start()

}
