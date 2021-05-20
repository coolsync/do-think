package main

import (
	"log"
	"strings"

	"github.com/beego/beego/v2/core/validation"
)

// Set validation function in "valid" tag
// Use ";" as the separator of multiple functions. Spaces accept after ";"
// Wrap parameters with "()" and separate parameter with ",". Spaces accept after ","
// Wrap regex match with "//"
//
// 各个函数的结果的key值为字段名.验证函数名
type user struct {
	Id int
	// Name   string `valid:"Required;Match(/^Bee.*/)"` // Name can't be empty or start with Bee
	Name   string `valid:"Required"`
	Age    int    `valid:"Range(1, 140)"`       // 1 <= Age <= 140, only valid in this range
	Email  string `valid:"Email; MaxSize(100)"` // Need to be a valid Email address and no more than 100 characters.
	// Mobile string `valid:"Mobile"`              // Must be a valid mobile number
	// IP     string `valid:"IP"`                  // Must be a valid IPv4 address
}

// If your struct implemented interface `validation.ValidFormer`
// When all tests in StructTag succeed, it will execute Valid function for custom validation
func (u *user) Valid(v *validation.Validation) {
	if strings.Index(u.Name, "admin") != -1 {
		// Set error messages of Name by SetError and HasErrors will return true
		v.SetError("Name", "Can't contain 'admin' in Name")
	}
}

func main() {
	valid := validation.Validation{}
	u := user{Name: "admin", Age: 2, Email: "dev@beego.me"}
	b, err := valid.Valid(&u)
	if err != nil {
		// handle error
		log.Println(err)
		return
	}
	if !b {
		// validation does not pass
		// blabla...
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}
}
