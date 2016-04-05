package expr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/* Tests for binary expression evaluation */

func TestAdd(t *testing.T) {
	assert.Equal(t, NewAdd(NewIntLit(1), NewIntLit(2)).Eval(nil), NewIntValue(3))
}

func TestMul(t *testing.T) {
	assert.Equal(t, NewMul(NewIntLit(3), NewIntLit(2)).Eval(nil), NewIntValue(6))
}

func TestMulAddPrecedence(t *testing.T) {
	assert.Equal(t, NewAdd(NewMul(NewIntLit(3), NewIntLit(2)), NewIntLit(1)).Eval(nil), NewIntValue(7))
}

func TestSub(t *testing.T) {
	assert.Equal(t, NewSub(NewIntLit(1), NewIntLit(2)).Eval(nil), NewIntValue(-1))
}

func TestDiv(t *testing.T) {
	assert.Equal(t, NewDiv(NewIntLit(9), NewIntLit(3)).Eval(nil), NewIntValue(3))
}

func TestGT(t *testing.T) {
	assert.Equal(t, NewGT(NewIntLit(3), NewIntLit(2)).Eval(nil), NewBoolValue(true))
}

func TestLT(t *testing.T) {
	assert.Equal(t, NewLT(NewIntLit(3), NewIntLit(2)).Eval(nil), NewBoolValue(false))
}

func TestGEq(t *testing.T) {
	assert.Equal(t, NewGEq(NewIntLit(3), NewIntLit(3)).Eval(nil), NewBoolValue(true))
}

func TestLEq(t *testing.T) {
	assert.Equal(t, NewLEq(NewIntLit(3), NewIntLit(3)).Eval(nil), NewBoolValue(true))
}

func TestAnd(t *testing.T) {
	assert.Equal(t, NewAnd(NewBoolLit(true), NewBoolLit(false)).Eval(nil), NewBoolValue(false))
}

func TestOr(t *testing.T) {
	assert.Equal(t, NewOr(NewBoolLit(true), NewBoolLit(false)).Eval(nil), NewBoolValue(true))
}

func TestAndOr(t *testing.T) {
	assert.Equal(t, NewAnd(NewOr(NewBoolLit(true), NewBoolLit(false)), NewBoolLit(false)).Eval(nil), NewBoolValue(false))
}

func TestEq(t *testing.T) {
	assert.Equal(t, NewEq(NewBoolLit(true), NewBoolLit(false)).Eval(nil), NewBoolValue(false))
}

func TestNEq(t *testing.T) {
	assert.Equal(t, NewNEq(NewBoolLit(true), NewBoolLit(false)).Eval(nil), NewBoolValue(true))
}
