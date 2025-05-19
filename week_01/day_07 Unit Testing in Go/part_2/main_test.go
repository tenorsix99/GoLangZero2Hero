package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {
	assert := assert.New(t)

	var a string = "Hello"
	var b string = "Hello"

	assert.Equal(a, b, "The two words should be the same.")
}

func TestExample(t *testing.T) {
	result := 2 + 3
	assert.Equal(t, 5, result, "2 + 3 should be 5")
}
