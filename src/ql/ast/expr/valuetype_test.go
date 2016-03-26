package expr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVarTypeDefaultValueForIntType(t *testing.T) {
	varTypeExample := NewIntType()
	correctDefaultValue := NewIntLit(0)
	assert.Equal(t, varTypeExample.DefaultValue(), correctDefaultValue)
}

func TestVarTypeDefaultValueForBoolType(t *testing.T) {
	varTypeExample := NewBoolType()
	correctDefaultValue := NewBoolLit(false)
	assert.Equal(t, varTypeExample.DefaultValue(), correctDefaultValue)
}

func TestVarTypeDefaultValueForStringType(t *testing.T) {
	varTypeExample := NewStringType()
	correctDefaultValue := NewStrLit("")
	assert.Equal(t, varTypeExample.DefaultValue(), correctDefaultValue)
}
