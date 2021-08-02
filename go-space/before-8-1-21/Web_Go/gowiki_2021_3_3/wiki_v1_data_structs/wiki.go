package main

import (
	"fmt"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

// save page
func (p *Page) save() error {
	filename := p.Title + ".txt"

	return os.WriteFile(filename, p.Body, 0600)
}

// loadPage
func (p *Page) loadPage(title string) (*Page, error) {
	filename := title + ".txt"

	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return &Page{Title: title, Body: body}, nil
}

func main() {
	p1 := &Page{Title: "openPage", Body: []byte("this is a simple page")}

	p1.save()

	p2, _ := p1.loadPage("openPage")

	fmt.Println(string(p2.Body))
}
