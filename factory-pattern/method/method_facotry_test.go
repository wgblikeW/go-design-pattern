package methodFactory

import (
	"testing"

	"github.com/likexian/gokit/assert"
)

func TestMethodFacotry(t *testing.T) {
	groupAge1 := NewPersonFactory(1)
	babyOne := groupAge1("Bod")
	assert.Equal(t, babyOne.Age, 1)
	assert.Equal(t, babyOne.Name, "Bod")

	groupAge16 := NewPersonFactory(16)
	teenager16 := groupAge16("Jack")
	assert.Equal(t, teenager16.Age, 16)
	assert.Equal(t, teenager16.Name, "Jack")
}
