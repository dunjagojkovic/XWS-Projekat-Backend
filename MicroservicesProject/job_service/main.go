package main

import "fmt"

func main() {

	config := NewConfig()
	fmt.Println(config.JobDBHost)
	server := NewServer(config)
	server.Start()

}
