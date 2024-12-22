package app_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// dummy test just to get it up & running.
func TestFoo(t *testing.T) {
	expected := 1
	actual := 1
	assert.Equal(t, expected, actual)

	expected = 2
	assert.NotEqual(t, expected, actual)
}
