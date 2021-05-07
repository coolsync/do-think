package main

import (
	"context"
	"fmt"

	"github.com/olivere/elastic/v7"
)

// operate elastic example

type Employee struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Worked bool   `json:"worked"`
}

func main() {
	// 1. init connect, get client obj
	// cli, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))
	cli, err := elastic.NewClient(elastic.SetURL("http://192.168.0.107:9200"))

	if err != nil {
		panic(err)
	}
	fmt.Println("Connect to es success")
	// 2. instace Emp
	e1 := Employee{Name: "Rich", Age: 9000, Worked: true}

	// 3. chain operate, insert obj to es
	put1, err := cli.Index().
		Index("users").
		// Id("1").
		BodyJson(e1).
		Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Indexed users %q to index %q, type %q\n", put1.Id, put1.Index, put1.Type)
	// 4. print info
}
