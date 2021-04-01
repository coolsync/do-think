# HTTP



http server

```go
func listHandler(w http.ResponseWriter, r *http.Request) {
	bs, err := os.ReadFile("./list.html")
	if err != nil {
		log.Println(err)
	}

	w.Write(bs)
}

func main() {
	http.HandleFunc("/list", listHandler)
	http.ListenAndServe("localhost:8090", nil)
}
```



go 1.6 ioutil package change

deprecation of ioutil:

- [`Discard`](https://docs.studygolang.com/pkg/io/ioutil/#Discard)      => [`io.Discard`](https://docs.studygolang.com/pkg/io/#Discard)
- [`NopCloser`](https://docs.studygolang.com/pkg/io/ioutil/#NopCloser)      => [`io.NopCloser`](https://docs.studygolang.com/pkg/io/#NopCloser)
- [`ReadAll`](https://docs.studygolang.com/pkg/io/ioutil/#ReadAll)      => [`io.ReadAll`](https://docs.studygolang.com/pkg/io/#ReadAll)
- [`ReadDir`](https://docs.studygolang.com/pkg/io/ioutil/#ReadDir)      => [`os.ReadDir`](https://docs.studygolang.com/pkg/os/#ReadDir)      (note: returns a slice of      [`os.DirEntry`](https://docs.studygolang.com/pkg/os/#DirEntry)      rather than a slice of      [`fs.FileInfo`](https://docs.studygolang.com/pkg/fs/#FileInfo))    
- [`ReadFile`](https://docs.studygolang.com/pkg/io/ioutil/#ReadFile)      => [`os.ReadFile`](https://docs.studygolang.com/pkg/os/#ReadFile)
- [`TempDir`](https://docs.studygolang.com/pkg/io/ioutil/#TempDir)      => [`os.MkdirTemp`](https://docs.studygolang.com/pkg/os/#MkdirTemp)
- [`TempFile`](https://docs.studygolang.com/pkg/io/ioutil/#TempFile)      => [`os.CreateTemp`](https://docs.studygolang.com/pkg/os/#CreateTemp)
- [`WriteFile`](https://docs.studygolang.com/pkg/io/ioutil/#WriteFile)      => [`os.WriteFile`](https://docs.studygolang.com/pkg/os/#WriteFile)



# Handle request 1



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



# Handle request 2

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



# Built-in Handlers



http.ServeFile

```go
http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		http.ServeFile(rw, r, "./template_html"+r.URL.Path)
	})

http.ListenAndServe("localhost:8081", nil)
```

http.FileServer

```go
	http.ListenAndServe("localhost:8081", http.FileServer(http.Dir("./template_html")))
```



