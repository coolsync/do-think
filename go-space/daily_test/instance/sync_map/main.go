package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	var m sync.Map

	// set key value
	// Store sets the value for a key.
	m.Store("bob", 15)
	m.Store("paul", 30)

	// get key
	// Load returns the value stored in the map for a key, or nil if no
	// value is present.
	// The ok result indicates whether value was found in the map.

	// age, ok := m.Load("jerry")
	age, ok := m.Load("bob")
	if !ok {
		fmt.Fprintf(os.Stdout, "%v\n", "not store")

	} else {
		fmt.Fprintf(os.Stdout, "%T, %d\n", age, age)
	}

	// traverse map key, value
	// Range calls f sequentially for each key and value present in the map.
	// If f returns false, range stops the iteration.
	m.Range(func(key, value interface{}) bool {
		name := key.(string)
		age := value.(int)

		fmt.Println(name, age)

		return true
	})

	// delete
	// Delete deletes the value for a key.
	m.Delete("paul")
	age, ok = m.Load("paul")
	fmt.Println(age, ok)

	// read or write
	// LoadOrStore returns the existing value for the key if present.
	// Otherwise, it stores and returns the given value.
	// The loaded result is true if the value was loaded, false if stored.

	// m.LoadOrStore("jerry", 25)
	// age, _ = m.Load("jerry")
	// fmt.Println(age)

	actual, ok := m.LoadOrStore("bob", 30)	// 15 true
	fmt.Println(actual, ok)

	age, _ = m.Load("bob")	// 15 true
	fmt.Println(age)
}
