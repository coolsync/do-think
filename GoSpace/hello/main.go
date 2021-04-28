package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)a

func main() {
	fname := "/etc/hosts"
	f, err := os.OpenFile(fname, os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// -rw-r--r--. 1 root root 2049 Apr 26 15:03 /etc/hosts
	// 接受io.Reader类型参数 返回一个bufio.Scanner实例
	scanner := bufio.NewScanner(f)

	var count int

	for scanner.Scan() {

		count++

		// 读取当前行内容 "# GitHub520 Host Start"
		line := scanner.Text()
		if count == 5 {

		}
		fmt.Printf("%d %s\n", count, line)
	}
}
