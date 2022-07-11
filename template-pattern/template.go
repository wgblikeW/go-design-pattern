package template

import "fmt"

// Template Pattern

type Cooker interface {
	fire()
	cooke()
	outfire()
}

type CookMenu struct{}

func (CookMenu) fire() {
	fmt.Println("Fire!")
}

func (CookMenu) cooke() {}

func (CookMenu) outfire() {
	fmt.Println("Outfire!")
}

func doCook(cook Cooker) {
	cook.fire()
	cook.cooke()
	cook.outfire()
}

type Tomatoes struct {
	CookMenu
}

func (*Tomatoes) cooke() {
	fmt.Println("Cooking Tomatoes !")
}

type Eggs struct {
	CookMenu
}

func (*Eggs) cooke() {
	fmt.Println("Cooking Eggs !")
}
