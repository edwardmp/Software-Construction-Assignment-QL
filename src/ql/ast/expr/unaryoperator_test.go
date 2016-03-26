package expr

import (
	"github.com/stretchr/testify/assert"
	"ql/ast/vari"
	"ql/symbols"
	"testing"
)

/* Tests for unary expressions */

func TestNot(t *testing.T) {
	assert.Equal(t, NewNot(NewBoolLit(true)).Eval(nil), false)
}

func TestPos(t *testing.T) {
	assert.Equal(t, NewPos(NewIntLit(-10)).Eval(nil), 10)
}

func TestNeg(t *testing.T) {
	assert.Equal(t, NewNeg(NewIntLit(10)).Eval(nil), -10)
}

func TestPosNeg(t *testing.T) {
	assert.Equal(t, NewPos(NewNeg(NewIntLit(-10))).Eval(nil), 10)
}

func TestNegPos(t *testing.T) {
	assert.Equal(t, NewNeg(NewPos(NewIntLit(10))).Eval(nil), -10)
}

func TestVarExpr(t *testing.T) {
	exampleVarId := vari.NewVarId("TestIdentifier")

	symbols := symbols.NewVarIdValueSymbols()
	symbols.SetExprForVarId(NewIntLit(2), exampleVarId)

	assert.Equal(t, NewVarExpr(exampleVarId).Eval(symbols), 2)
}
