package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

// reader read data from stdio
type HelloReader struct{}

func (h *HelloReader) Read(b []byte) (n int, err error) {
	fmt.Print("in > ")
	return os.Stdin.Read(b)
}

// writer write data to stdout
type HelloWriter struct{}

func (h *HelloWriter) Write(b []byte) (n int, err error) {
	fmt.Print("out > ")
	return os.Stdout.Write(b)
}

func main() {
	var (
		reader HelloReader
		writer HelloWriter
	)

	// // create buffer
	// input := make([]byte, 1024)

	// n, err := reader.Read(input)
	// if err != nil {
	// 	log.Fatalln("read data failed")
	// }
	// fmt.Printf("read from stdio: %d\n", n)

	// n, err = writer.write(input)
	// if err != nil {
	// 	log.Fatalln("write data failed")
	// }

	// fmt.Printf("wrote to stdout: %d\n", n)
	
	_ , err := io.Copy(&writer, &reader)
	if err != nil {
		log.Fatalln("Unable read/write data")
	}
}
