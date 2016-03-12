package expr

import (
	"ql/interfaces"
	"testing"
)

func binaryExprEval(t *testing.T, exampleInput interfaces.Expr, expectedOutput interfaces.Expr) {
	if eval, expectedOutputEval := exampleInput.Eval(nil), expectedOutput.(interfaces.Expr).Eval(nil); eval != expectedOutputEval {
		t.Errorf("interfaces.Expr test error: should be %v (%T) for %v but is %v (%T)", expectedOutputEval, expectedOutputEval, eval, eval)
	}
}

/* Tests for binary expression evaluation */

func TestAdd(t *testing.T) {
	binaryExprEval(t, NewAddNoSourceInfo(NewIntLitNoSourceInfo(1), NewIntLitNoSourceInfo(2)), NewIntLitNoSourceInfo(3))
}

func TestMul(t *testing.T) {
	binaryExprEval(t, NewMulNoSourceInfo(NewIntLitNoSourceInfo(3), NewIntLitNoSourceInfo(2)), NewIntLitNoSourceInfo(6))
}

func TestMulAddPrecedence(t *testing.T) {
	binaryExprEval(t, NewAddNoSourceInfo(NewMulNoSourceInfo(NewIntLitNoSourceInfo(3), NewIntLitNoSourceInfo(2)), NewIntLitNoSourceInfo(1)), NewIntLitNoSourceInfo(7))
}

func TestSub(t *testing.T) {
	binaryExprEval(t, NewSubNoSourceInfo(NewIntLitNoSourceInfo(1), NewIntLitNoSourceInfo(2)), NewIntLitNoSourceInfo(-1))
}

func TestDiv(t *testing.T) {
	binaryExprEval(t, NewDivNoSourceInfo(NewIntLitNoSourceInfo(9), NewIntLitNoSourceInfo(3)), NewIntLitNoSourceInfo(3))
}

func TestGT(t *testing.T) {
	binaryExprEval(t, NewGTNoSourceInfo(NewIntLitNoSourceInfo(3), NewIntLitNoSourceInfo(2)), NewBoolLitNoSourceInfo(true))
}

func TestLT(t *testing.T) {
	binaryExprEval(t, NewLTNoSourceInfo(NewIntLitNoSourceInfo(3), NewIntLitNoSourceInfo(2)), NewBoolLitNoSourceInfo(false))
}

func TestGEq(t *testing.T) {
	binaryExprEval(t, NewGEqNoSourceInfo(NewIntLitNoSourceInfo(3), NewIntLitNoSourceInfo(3)), NewBoolLitNoSourceInfo(true))
}

func TestLEq(t *testing.T) {
	binaryExprEval(t, NewLEqNoSourceInfo(NewIntLitNoSourceInfo(3), NewIntLitNoSourceInfo(3)), NewBoolLitNoSourceInfo(true))
}

func TestAnd(t *testing.T) {
	binaryExprEval(t, NewAndNoSourceInfo(NewBoolLitNoSourceInfo(true), NewBoolLitNoSourceInfo(false)), NewBoolLitNoSourceInfo(false))
}

func TestOr(t *testing.T) {
	binaryExprEval(t, NewOrNoSourceInfo(NewBoolLitNoSourceInfo(true), NewBoolLitNoSourceInfo(false)), NewBoolLitNoSourceInfo(true))
}

func TestAndOr(t *testing.T) {
	binaryExprEval(t, NewAndNoSourceInfo(NewOrNoSourceInfo(NewBoolLitNoSourceInfo(true), NewBoolLitNoSourceInfo(false)), NewAndNoSourceInfo(NewBoolLitNoSourceInfo(true), NewBoolLitNoSourceInfo(false))), NewBoolLitNoSourceInfo(false))
}

func TestEq(t *testing.T) {
	binaryExprEval(t, NewEqNoSourceInfo(NewBoolLitNoSourceInfo(true), NewBoolLitNoSourceInfo(false)), NewBoolLitNoSourceInfo(false))
}

func TestNEq(t *testing.T) {
	binaryExprEval(t, NewNEqNoSourceInfo(NewBoolLitNoSourceInfo(true), NewBoolLitNoSourceInfo(false)), NewBoolLitNoSourceInfo(true))
}
