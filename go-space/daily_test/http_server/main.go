package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// 定义结构体 和 method 实现 Handler 接口
type Handelrs struct {
}

// func (h *Handelrs) handler(w http.ResponseWriter, r *http.Request) {
func (h *Handelrs) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)

	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}

	fmt.Fprintf(w, "Host = %q\n", r.Host)

	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

func main() {
	// http.HandleFunc("/", handler)

	// http.ListenAndServe(":8081", nil)

	s := http.Server{
		Addr:           ":8081",
		Handler:        &Handelrs{},
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
