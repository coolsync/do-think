package main

func main() {
	server := NewServer("localhost", 8081)
	// function main is undeclared in the main package
	server.Start()
}
