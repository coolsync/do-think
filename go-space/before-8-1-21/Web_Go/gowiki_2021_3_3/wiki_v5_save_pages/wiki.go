package main

import (
	"log"
	"net/http"
	"os"
	"text/template"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"

	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	// filename := title + ".html"

	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]

	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound) // 302
		return
	}

	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]

	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	// runtime error: invalid memory address or nil pointer dereference

	renderTemplate(w, "edit", p)
}
func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]

	body := r.FormValue("body")

	p := &Page{Title: title, Body: []byte(body)}

	p.save()

	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)

	log.Fatal(http.ListenAndServe(":8081", nil))
}
