package main

import "fmt"

// def inter
type GetInfo interface {
	GetName() string
	GetAge() uint32
}

type CompanyName interface {
	GetCompanyName() string
}

type Employee struct {
	Name string
	Age  uint32
}

func (e *Employee) GetName() string {
	return e.Name
}

func (e *Employee) GetAge() uint32 {
	return e.Age
}

type Manager struct {
	Name       string
	Age        uint32
	Department string
}

func (m Manager) GetName() string {
	return m.Name
}

func (m Manager) GetAge() uint32 {
	return m.Age
}

func printInfo(item GetInfo) {
	fmt.Printf("%s, %d\n", item.GetName(), item.GetAge())
}

func main() {
	// var info GetInfo
	e := &Employee{Name: "hello", Age: 30}
	m := Manager{Name: "bob", Age: 30}
	// info = e

	printInfo(e)
	printInfo(m)
}
