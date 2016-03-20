package expr

import (
	"ql/interfaces"
	"testing"
)

func binaryExprEval(t *testing.T, exampleInput interfaces.Expr, expectedOutput interfaces.Expr) {
	if eval, expectedOutputEval := exampleInput.Eval(nil), expectedOutput.(interfaces.Expr).Eval(nil); eval != expectedOutputEval {
		t.Errorf("BinaryOperator test error: should be %v (%T) for %v but is %v (%T)", expectedOutputEval, expectedOutputEval, eval, eval)
	}
}

/* Tests for binary expression evaluation */

func TestAdd(t *testing.T) {
	binaryExprEval(t, NewAdd(NewIntLit(1), NewIntLit(2)), NewIntLit(3))
}

func TestMul(t *testing.T) {
	binaryExprEval(t, NewMul(NewIntLit(3), NewIntLit(2)), NewIntLit(6))
}

func TestMulAddPrecedence(t *testing.T) {
	binaryExprEval(t, NewAdd(NewMul(NewIntLit(3), NewIntLit(2)), NewIntLit(1)), NewIntLit(7))
}

func TestSub(t *testing.T) {
	binaryExprEval(t, NewSub(NewIntLit(1), NewIntLit(2)), NewIntLit(-1))
}

func TestDiv(t *testing.T) {
	binaryExprEval(t, NewDiv(NewIntLit(9), NewIntLit(3)), NewIntLit(3))
}

func TestGT(t *testing.T) {
	binaryExprEval(t, NewGT(NewIntLit(3), NewIntLit(2)), NewBoolLit(true))
}

func TestLT(t *testing.T) {
	binaryExprEval(t, NewLT(NewIntLit(3), NewIntLit(2)), NewBoolLit(false))
}

func TestGEq(t *testing.T) {
	binaryExprEval(t, NewGEq(NewIntLit(3), NewIntLit(3)), NewBoolLit(true))
}

func TestLEq(t *testing.T) {
	binaryExprEval(t, NewLEq(NewIntLit(3), NewIntLit(3)), NewBoolLit(true))
}

func TestAnd(t *testing.T) {
	binaryExprEval(t, NewAnd(NewBoolLit(true), NewBoolLit(false)), NewBoolLit(false))
}

func TestOr(t *testing.T) {
	binaryExprEval(t, NewOr(NewBoolLit(true), NewBoolLit(false)), NewBoolLit(true))
}

func TestAndOr(t *testing.T) {
	binaryExprEval(t, NewAnd(NewOr(NewBoolLit(true), NewBoolLit(false)), NewAnd(NewBoolLit(true), NewBoolLit(false))), NewBoolLit(false))
}

func TestEq(t *testing.T) {
	binaryExprEval(t, NewEq(NewBoolLit(true), NewBoolLit(false)), NewBoolLit(false))
}

func TestNEq(t *testing.T) {
	binaryExprEval(t, NewNEq(NewBoolLit(true), NewBoolLit(false)), NewBoolLit(true))
}
