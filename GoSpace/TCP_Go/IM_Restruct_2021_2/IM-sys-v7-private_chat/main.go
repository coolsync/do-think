package main

func main() {
	server := NewServer("localhost", 8081)

	server.Start()
}