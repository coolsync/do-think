package main

import (
	"cphosts/middle"
	"fmt"
	"log"
	"os"
)

func main() {
	url := "https://cdn.jsdelivr.net/gh/521xueweihan/GitHub520@main/hosts"
	bs, err := middle.GetResp(url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bs))

	// write to hosts file
	file_path := "./hosts"
	if err = os.WriteFile(file_path, bs, 0644); err != nil {
		log.Fatal(err)
	}
	
	// f, err := os.OpenFile(file_path, os.O_CREATE|os.O_WRONLY, 0644)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// bufio.NewWriter(f)


	// resp, err := http.Get("https://cdn.jsdelivr.net/gh/521xueweihan/GitHub520@main/hosts")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer resp.Body.Close()

	// bs, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(string(bs))

}
