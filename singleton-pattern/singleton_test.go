package singleton

import (
	"testing"

	"github.com/likexian/gokit/assert"
)

func TestSingleton(t *testing.T) {
	singleton := GetInsOr()
	assert.NotNil(t, singleton)
}
