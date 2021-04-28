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
	target_file := "/etc/hosts"
	tmp_file := "/etc/hosts_tmp"

	url := "https://cdn.jsdelivr.net/gh/521xueweihan/GitHub520@main/hosts"
	bs := GetUrl(url)            // 后面添加的内容
	str := ReadFile(target_file) // 目标行前面需要提取的内容

	all_bs := []byte(str + "\n\n")
	all_bs = append(all_bs, bs...)

	os.WriteFile(tmp_file, all_bs, 0664)

	err := os.Rename(tmp_file, target_file)

	if err != nil {
		log.Fatal(err)
	}
}

func GetUrl(url string) []byte {
	resp, err := http.Get(url)
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

func ReadFile(target_file string) string {
	bs, err := os.ReadFile(target_file)
	if err != nil {
		log.Fatal("os.ReadFile err: ", err)
	}

	scan := bufio.NewScanner(strings.NewReader(string(bs)))

	var lines string

	for scan.Scan() {
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
