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

func renderPage(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]

	p, _ := loadPage(title)

	// fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)

	// t, _ := template.ParseFiles("view.html") // display view page

	// t.Execute(w, p) // p struct data output to view page {{}}
	renderPage(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]

	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}

	// htmlStr := `<h1>Editing %s</h1>
	// <form action="/save/%s" method="POST">
	//     <textarea name="body">%s</textarea><br>
	//     <input type="submit" value="Save">
	// </form>`

	// fmt.Fprintf(w, htmlStr, p.Title, p.Title, p.Body)

	// t, err := template.ParseFiles("edit.html") // display edit page
	// if err != nil {
	// 	log.Fatal("ParseFiles err:", err)
	// }

	// t.Execute(w, p)
	renderPage(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]

	body := r.PostFormValue("body")

	p := &Page{Title: title, Body: []byte(body)}

	p.save()
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	// http.HandleFunc("/save/", saveHandler)

	log.Fatal(http.ListenAndServe(":8081", nil))
}
