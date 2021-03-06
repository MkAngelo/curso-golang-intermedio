package main

import "fmt"

type Employee struct {
	id   int
	name string
}

func NewEmployee(id int, name string) *Employee {
	return &Employee{
		id:   id,
		name: name,
	}
}

func (e *Employee) SetId(id int) { // Function Arguments
	e.id = id
}

func (e *Employee) SetName(name string) {
	e.name = name
}

func (e *Employee) GetId() int {
	return e.id
}

func (e *Employee) GetName() string {
	return e.name
}

func main() {
	// Constructor f1
	e := Employee{}
	fmt.Printf("%v\n", e)

	// CONSTRUCTOR FORMA 2
	e2 := Employee{
		id:   2,
		name: "Name2",
	}
	fmt.Printf("%v\n", e2)
	//  Forma 3
	e3 := new(Employee)
	fmt.Printf("%v\n", *e3)
	// Forma 4
	e4 := NewEmployee(4, "Name4")
	fmt.Printf("%v\n", *e4)
}
