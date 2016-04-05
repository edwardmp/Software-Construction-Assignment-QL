package expr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntLitGetValue(t *testing.T) {
	assert.Equal(t, NewIntLit(10).Value(), NewIntValue(10))
}

func TestBoolLitGetValue(t *testing.T) {
	assert.Equal(t, NewBoolLit(true).Value(), NewBoolValue(true))
}

func TestStrLitGetValue(t *testing.T) {
	assert.Equal(t, NewStrLit("Test").Value(), NewStringValue("Test"))
}
