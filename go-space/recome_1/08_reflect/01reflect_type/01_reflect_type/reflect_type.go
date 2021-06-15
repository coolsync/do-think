package main

import (
	"fmt"
	"reflect"
)

// def Enum type
type Enum int

const (
	Zero Enum = 0
)

func main() {
	type dog struct{}

	type_of_dog := reflect.TypeOf(dog{})

	fmt.Println(type_of_dog.Name(), type_of_dog.Kind()) // dog struct

	type_of_enum := reflect.TypeOf(Zero)

	fmt.Println(type_of_enum.Name(), type_of_enum.Kind()) // Enum int

}
