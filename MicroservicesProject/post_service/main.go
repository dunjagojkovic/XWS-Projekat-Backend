package main

import "fmt"

func main() {
	config := NewConfig()
	fmt.Println(config.PostDBHost)
	server := NewServer(config)
	server.Start()
}
