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
	binaryExprEval(t, NewAdd(IntLit{1}, IntLit{2}), IntLit{3})
}

func TestMul(t *testing.T) {
	binaryExprEval(t, NewMul(IntLit{3}, IntLit{2}), IntLit{6})
}

func TestMulAddPrecedence(t *testing.T) {
	binaryExprEval(t, NewAdd(NewMul(IntLit{3}, IntLit{2}), IntLit{1}), IntLit{7})
}

func TestSub(t *testing.T) {
	binaryExprEval(t, NewSub(IntLit{1}, IntLit{2}), IntLit{-1})
}

func TestDiv(t *testing.T) {
	binaryExprEval(t, NewDiv(IntLit{9}, IntLit{3}), IntLit{3})
}

func TestGT(t *testing.T) {
	binaryExprEval(t, NewGT(IntLit{3}, IntLit{2}), BoolLit{true})
}

func TestLT(t *testing.T) {
	binaryExprEval(t, NewLT(IntLit{3}, IntLit{2}), BoolLit{false})
}

func TestGEq(t *testing.T) {
	binaryExprEval(t, NewGEq(IntLit{3}, IntLit{3}), BoolLit{true})
}

func TestLEq(t *testing.T) {
	binaryExprEval(t, NewLEq(IntLit{3}, IntLit{3}), BoolLit{true})
}

func TestAnd(t *testing.T) {
	binaryExprEval(t, NewAnd(BoolLit{true}, BoolLit{false}), BoolLit{false})
}

func TestOr(t *testing.T) {
	binaryExprEval(t, NewOr(BoolLit{true}, BoolLit{false}), BoolLit{true})
}

func TestAndOr(t *testing.T) {
	binaryExprEval(t, NewAnd(NewOr(BoolLit{true}, BoolLit{false}), NewAnd(BoolLit{true}, BoolLit{false})), BoolLit{false})
}

func TestEq(t *testing.T) {
	binaryExprEval(t, NewEq(BoolLit{true}, BoolLit{false}), BoolLit{false})
}

func TestNEq(t *testing.T) {
	binaryExprEval(t, NewNEq(BoolLit{true}, BoolLit{false}), BoolLit{true})
}
