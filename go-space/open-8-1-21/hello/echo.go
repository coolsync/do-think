package main

import (
	"fmt"
	"os"
)

func main() {
	var out []*int

	// var v *int
	for i := 0; i < 3; i++ {
		// fmt.Println(&i)
		// v = &i
		// fmt.Println(*v)
		i := i
		out = append(out, &i)
	}

	fmt.Fprintln(os.Stdout, "values: ", *out[0], *out[1], *out[2])
}
