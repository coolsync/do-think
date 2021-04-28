package main

import (
	"log"
	"net/http"
	"os"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	bs, err := os.ReadFile("./list.html")
	if err != nil {
		log.Println(err)
	}

	w.Write(bs)
}

func main() {
	http.HandleFunc("/list", homeHandler)
	http.ListenAndServe("localhost:8090", nil)
}
