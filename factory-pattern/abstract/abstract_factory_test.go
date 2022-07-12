package abstractFactory_test

import (
	"testing"

	abstractFactory "github.com/p1nant0m/strategy-pattern/factory-pattern/abstract"
)

func TestAbstractFactory(t *testing.T) {
	jack := abstractFactory.NewPersonJack()
	sayHello(jack)

	peter := abstractFactory.NewPersonPeter()
	sayHello(peter)
}

func sayHello(person abstractFactory.Person) {
	person.Greet()
}
