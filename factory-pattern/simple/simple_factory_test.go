package simpleFactory

import "testing"

func TestSimpleFactory(t *testing.T) {
	peter := NewPerson("Peter", 16)
	peter.Greet()
}
