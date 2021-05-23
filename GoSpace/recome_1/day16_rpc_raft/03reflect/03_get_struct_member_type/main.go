package main

import (
	"fmt"
	"reflect"
)

func main() {

	type dog struct {
		Name string
		Type int `json:"type" id="100"` // struct field tag `json:"type" id="100"` not compatible with reflect.StructTag.Get: bad syntax for struct tag pair
		// Type int `json:"type"`
	}

	// get struct instance reflect obj
	d := dog{Name: "hello", Type: 1}
	type_of_dog := reflect.TypeOf(d)

	// traverse struct, get filed name, tag, kind
	for i := 0; i < type_of_dog.NumField(); i++ {
		filed_type := type_of_dog.Field(i)

		fmt.Printf("filed name: %v, tag: %v, kind: %v\n", filed_type.Name, filed_type.Tag, filed_type.Type.Kind())
	}

	// due to FieldByName, get tag
	if filed_type, ok := type_of_dog.FieldByName("Type"); ok {
		fmt.Printf("name:%v, tag:%v\n", filed_type.Name, filed_type.Tag)
	}

	// dog struct tag Type
	if dog_Type, ok := type_of_dog.FieldByName("Type"); ok {
		fmt.Printf("json tag :%v, id tag: %v\n", dog_Type.Tag.Get("json"), dog_Type.Tag.Get("id"))
	}
}
