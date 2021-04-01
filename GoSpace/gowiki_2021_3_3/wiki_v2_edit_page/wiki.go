package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

// save page file
func (p *Page) save() error {
	filename := p.Title + ".txt"

	return os.WriteFile(filename, p.Body, 0600)
}

// load page file
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"

	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return &Page{Title: title, Body: body}, nil
}

// invalid memory address or nil pointer dereference
func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]

	p1, _ := loadPage(title)

	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p1.Title, p1.Body)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	html_str := `<h1>Editing %s</h1>
    <form action="/save/%s" method="POST">
        <textarea name="body">%s</textarea><br>
        <input type="submit" value="Save">
    </form>`

	title := r.URL.Path[len("/edit/"):]

	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}

	fmt.Fprintf(w, html_str, p.Title, p.Title, p.Body)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]

	body := []byte(r.FormValue("body"))

	p := &Page{Title: title, Body: body}

	err := p.save()

	if err != nil {
		log.Fatal("save err:", err)
	}

	fmt.Fprintf(w, "save %s ok", p.Title)
}

func main() {

	// registe callback function
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)

	log.Fatal(http.ListenAndServe(":8081", nil))
}
