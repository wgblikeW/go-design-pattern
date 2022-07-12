package abstractFactory

import "fmt"

type Person interface {
	Greet()
}

type jack struct{}

func (j *jack) Greet() {
	fmt.Println("Hi! I am Jack.")
}

type peter struct{}

func (p *peter) Greet() {
	fmt.Println("Hi! I am Peter.")
}

func NewPersonJack() Person {
	return &jack{}
}

func NewPersonPeter() Person {
	return &peter{}
}
