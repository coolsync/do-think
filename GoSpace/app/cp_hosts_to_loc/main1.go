package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://cdn.jsdelivr.net/gh/521xueweihan/GitHub520@main/hosts")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bs, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// write to hosts file
	fname := "./etc/hosts"
	if err = os.WriteFile(fname, bs, 0644); err != nil {
		log.Fatal(err)
	}
}
