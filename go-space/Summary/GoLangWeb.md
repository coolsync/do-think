# GoLang Web



# Handle request



```go
type myHandler struct {}

func (mh *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello web"))
}

func main() {

	mh := myHandler{}

	server  := http.Server{
		Addr: "localhost:8081",
		Handler: &mh,
	}
	server.ListenAndServe()
	// http.ListenAndServe("localhost:8081", nil)
}
```



# Handle request

```go
package main

import "net/http"

type helloHandler struct{}

func (h *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello!"))
}

type aboutHandler struct{}

func (a *aboutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About！"))
}

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome！"))
}

func main() {

	h := helloHandler{}
	a := aboutHandler{}

	server := http.Server{
		Addr:    "localhost:8081",
		Handler: nil, // DefaultServeMux
	}

	/*
		type Handler interface {
			ServeHTTP(ResponseWriter, *Request)
		}
	*/
	http.Handle("/hello", &h)
	http.Handle("/about", &a)

	http.HandleFunc("/home", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Home!"))
	})

	http.Handle("/welcome", http.HandlerFunc(welcome))

	server.ListenAndServe()
	// http.ListenAndServe("localhost:8081", nil)
}
```

