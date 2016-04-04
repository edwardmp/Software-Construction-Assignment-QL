package vari

import (
	"github.com/stretchr/testify/assert"
	"ql/ast/expr"
	"testing"
)

func TestNewVarId(t *testing.T) {
	varId := NewVarId("testIdentifier")

	assert.Equal(t, varId.Identifier(), "testIdentifier")
	assert.Equal(t, varId.String(), "testIdentifier")
}

func TestNewVarDecl(t *testing.T) {
	varId := NewVarId("testVarId")
	valueType := expr.NewIntType()
	varDecl := NewVarDecl(varId, valueType)

	assert.Equal(t, varDecl.VariableIdentifier(), varId)
	assert.Equal(t, varDecl.Type(), valueType)
}
