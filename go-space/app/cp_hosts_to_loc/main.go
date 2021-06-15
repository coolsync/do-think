package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	// Get url content
	resp, err := http.Get("https://cdn.jsdelivr.net/gh/521xueweihan/GitHub520@main/hosts")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bs, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// read file 

	fname := "./hosts"
	f, err := os.OpenFile(fname, os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("f: ", err)
	}
	defer f.Close()
}
