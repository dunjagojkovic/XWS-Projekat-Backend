package main

func main() {
	config := NewConfig()
	server := NewServer(config)
	server.Start()
}
