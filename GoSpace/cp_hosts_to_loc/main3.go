package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// f, err := os.Open()
	bs, err := os.ReadFile("/etc/hosts")
	if err != nil {
		log.Fatal("os.ReadFile err: ", err)
	}
	// fmt.Println(string(bs))

	scan := bufio.NewScanner(strings.NewReader(string(bs)))

	// var count int
	var lines string

	for scan.Scan() {
		// count++
		lineStr := scan.Text()

		if strings.Contains(lineStr, "# GitHub520 Host Start") {
			break
		}
		lines += lineStr + "\n"
		// fmt.Printf("%d line is: %s\n", count, lineStr)
	}

	lines = strings.TrimRight(lines, "\n")
	// fmt.Println(lines)
}
