package main

func main() {
	server, _ := NewServer("localhost", 8081)

	server.Start()
}