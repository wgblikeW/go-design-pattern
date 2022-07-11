package template

import "testing"

func TestCooke(t *testing.T) {
	egg := &Eggs{}
	t.Log("This should be Cooking Eggs =======>")
	doCook(egg)

	tomato := &Tomatoes{}
	t.Log("This should be Cooking Tomatoes ======>")
	doCook(tomato)
}
