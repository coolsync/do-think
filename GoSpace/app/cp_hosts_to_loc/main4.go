package main

import (
	"bufio"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	// f, err := os.OpenFile("./etc/hosts_tmp", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	// if err != nil {
	// 	log.Fatal("os.OpenFile err: ", err)
	// }
	// defer f.Close()

	target_file := "/etc/hosts"

	tmp_file := "/etc/hosts_tmp"

	bs := GetUrl()

	append_str := ReadFile2()

	all_bs := []byte(append_str + "\n\n")
	all_bs = append(all_bs, bs...)

	WriteFile2(tmp_file, all_bs, 0664)

	err := os.Rename(tmp_file, target_file)

	if err != nil {
		log.Fatal(err)
	}
}

func GetUrl() []byte {
	resp, err := http.Get("https://cdn.jsdelivr.net/gh/521xueweihan/GitHub520@main/hosts")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bs, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return bs
}

func ReadFile2() string {
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
	return lines
}

func WriteFile2(name string, data []byte, perm os.FileMode) error {
	f, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return err
	}
	_, err = f.Write(data)
	if err1 := f.Close(); err1 != nil && err == nil {
		err = err1
	}
	return err
}
