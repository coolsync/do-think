package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"regexp"
	"text/template"
)

type Page struct {
	Title string
	Body  []byte
}

var templates = template.Must(template.ParseFiles("view.html", "edit.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	// t, err := template.ParseFiles(tmpl + ".html")
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	// err = t.Execute(w, p)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }

	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

// -------------------
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r) // NotFound replies to the request with an HTTP 404 not found error.
		return "", errors.New("invalid title")
	}

	return m[2], nil // m[2] title name
}

// -------------------

func (p *Page) save() error {
	filename := p.Title + ".txt" // save data to local txt, can looking json data,

	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"

	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	// -------------------
	// title := r.URL.Path[len("/view/"):]
	// title, err := getTitle(w, r)
	// if err != nil {
	// 	return
	// }

	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound) // 302
		return
	}

	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	// title := r.URL.Path[len("/edit/"):]
	// title, err := getTitle(w, r)
	// if err != nil {
	// 	return
	// }

	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}

	renderTemplate(w, "edit", p)
}
func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	// title := r.URL.Path[len("/save/"):]
	// title, err := getTitle(w, r)
	// if err != nil {
	// 	return
	// }

	body := r.FormValue("body")

	p := &Page{Title: title, Body: []byte(body)}

	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func makeHandler(fn func(w http.ResponseWriter, r *http.Request, title string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		title, err := getTitle(w, r) // vilidation title
		if err != nil {
			return
		}

		fn(w, r, title)
	}
}

func main() {
	// http.HandleFunc("/view/", viewHandler)
	// http.HandleFunc("/edit/", editHandler)
	// http.HandleFunc("/save/", saveHandler)

	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))

	log.Fatal(http.ListenAndServe(":8081", nil))
}