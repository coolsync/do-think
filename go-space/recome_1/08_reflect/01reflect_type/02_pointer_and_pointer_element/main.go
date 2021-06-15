package main

import (
	"fmt"
	"reflect"
)

func main() {

	type dog struct{}

	// Get dog intance pointer obj
	type_of_dog := reflect.TypeOf(&dog{})

	fmt.Printf("Name: %q, Kind: %v\n", type_of_dog.Name(), type_of_dog.Kind()) // Name: "", Kind: ptr

	// Get element
	type_dog_elem := type_of_dog.Elem()

	fmt.Printf("Name: %q, Kind: %v\n", type_dog_elem.Name(), type_dog_elem.Kind()) // Name: "dog", Kind: struct

}
