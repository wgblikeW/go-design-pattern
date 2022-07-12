package simpleFactory

import "fmt"

// Factory  Pattern

type Person struct {
	Name string
	Age  int
}

func (p *Person) Greet() {
	fmt.Printf("Hi! My name is %s", p.Name)
}

// Simple Factory Pattern
// NewPerson return the instance of Person on given arguments
func NewPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}
