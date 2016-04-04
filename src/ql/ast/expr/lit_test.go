package expr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntLit(t *testing.T) {
	exampleIntLit := NewIntLit(10)
	assert.Equal(t, NewIntLit(10), exampleIntLit)
}

func TestIntLitGetValue(t *testing.T) {
	assert.Equal(t, NewIntLit(10).Value(), 10)
}

func TestBoolLit(t *testing.T) {
	exampleBoolLit := NewBoolLit(true)
	assert.Equal(t, NewBoolLit(true), exampleBoolLit)
}

func TestBoolLitGetValue(t *testing.T) {
	assert.True(t, NewBoolLit(true).Value())
}

func TestStrLit(t *testing.T) {
	exampleStrLit := NewStrLit("Test")
	assert.Equal(t, NewStrLit("Test"), exampleStrLit)
}

func TestStrLitGetValue(t *testing.T) {
	assert.Equal(t, NewStrLit("Test").Value(), "Test")
}
