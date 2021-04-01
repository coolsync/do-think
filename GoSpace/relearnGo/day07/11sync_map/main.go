package main

import (
	"fmt"
	"strconv"
	"sync"
)

var m sync.Map

func main() {
	wg := &sync.WaitGroup{}

	for i := 0; i < 20; i++ {
		key := strconv.Itoa(i)

		m.Store(key, i)

		wg.Add(1)
		go func() {
			defer wg.Done()
			value, _ := m.Load(key)

			fmt.Printf("key %#v,value %d\n", key, value)
		}()
	}
	wg.Wait()
}
