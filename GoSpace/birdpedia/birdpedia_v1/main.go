package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Declare a new router
	r := mux.NewRouter()

	// register callback function
	// http.HandleFunc("/", handler)
	r.HandleFunc("/hello", handler).Methods("GET")

	// listen port
	// log.Fatal(http.ListenAndServe(":8080", nil))
	http.ListenAndServe(":8080", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", "hello, bird")
}
