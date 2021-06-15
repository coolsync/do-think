package main

import (
	"log"

	"github.com/astaxie/beego/validation"
)

type User struct {
	Name string
	Age  int
}

func main() {
	// u := User{"mark", 40}
	u := User{"mark2", 12}

	// Init validator
	valid := validation.Validation{}

	valid.Required(u.Name, "name")
	valid.MaxSize(u.Name, 15, "nameMax")
	valid.Range(u.Age, 0, 18, "age")

	// customize error messages
	minAge := 18
	valid.Min(u.Age, minAge, "age").Message("18+ only!!")
	// Format error messages
	valid.Min(u.Age, minAge, "age").Message("%d+", minAge)


	if valid.HasErrors() {
		// If there are error messages it means the validation didn't pass
		// Print error message
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}

	// or user like this
	// if v := valid.Max(u.Age, 18, "age"); !v.Ok {
	if v := valid.Max(u.Age, 80, "age"); !v.Ok {
		log.Println(v.Error.Key, v.Error.Message)
	}
}
