package binaryoperatorexpr

import (
	"ql/ast/expr"
	"ql/ast/expr/litexpr"
	"testing"
)

func testExprEval(t *testing.T, exampleExpr interface{}, expectedOutput interface{}) {
	if eval, expectedOutputEval := exampleExpr.(expr.Expr).Eval(nil), expectedOutput.(expr.Expr).Eval(nil); eval != expectedOutputEval {
		t.Errorf("Expr test error: should be %v (%T) for %v but is %v (%T)", expectedOutputEval, expectedOutputEval, exampleExpr, eval, eval)
	}
}

/* Tests for binary expressions */

func TestAdd(t *testing.T) {
	testExprEval(t, NewAdd(litexpr.IntLit{1}, litexpr.IntLit{2}), litexpr.IntLit{3})
}

func TestMul(t *testing.T) {
	testExprEval(t, NewMul(litexpr.IntLit{3}, litexpr.IntLit{2}), litexpr.IntLit{6})
}

func TestMulAddPrecedence(t *testing.T) {
	testExprEval(t, NewAdd(NewMul(litexpr.IntLit{3}, litexpr.IntLit{2}), litexpr.IntLit{1}), litexpr.IntLit{7})
}

func TestSub(t *testing.T) {
	testExprEval(t, NewSub(litexpr.IntLit{1}, litexpr.IntLit{2}), litexpr.IntLit{-1})
}

func TestDiv(t *testing.T) {
	testExprEval(t, NewDiv(litexpr.IntLit{9}, litexpr.IntLit{3}), litexpr.IntLit{3})
}

func TestGT(t *testing.T) {
	testExprEval(t, NewGT(litexpr.IntLit{3}, litexpr.IntLit{2}), litexpr.BoolLit{true})
}

func TestLT(t *testing.T) {
	testExprEval(t, NewLT(litexpr.IntLit{3}, litexpr.IntLit{2}), litexpr.BoolLit{false})
}

func TestGEq(t *testing.T) {
	testExprEval(t, NewGEq(litexpr.IntLit{3}, litexpr.IntLit{3}), litexpr.BoolLit{true})
}

func TestLEq(t *testing.T) {
	testExprEval(t, NewLEq(litexpr.IntLit{3}, litexpr.IntLit{3}), litexpr.BoolLit{true})
}

func TestAnd(t *testing.T) {
	testExprEval(t, NewAnd(litexpr.BoolLit{true}, litexpr.BoolLit{false}), litexpr.BoolLit{false})
}

func TestOr(t *testing.T) {
	testExprEval(t, NewOr(litexpr.BoolLit{true}, litexpr.BoolLit{false}), litexpr.BoolLit{true})
}

func TestAndOr(t *testing.T) {
	testExprEval(t, NewAnd(NewOr(litexpr.BoolLit{true}, litexpr.BoolLit{false}), NewAnd(litexpr.BoolLit{true}, litexpr.BoolLit{false})), litexpr.BoolLit{false})
}

func TestEq(t *testing.T) {
	testExprEval(t, NewEq(litexpr.BoolLit{true}, litexpr.BoolLit{false}), litexpr.BoolLit{false})
}

func TestNEq(t *testing.T) {
	testExprEval(t, NewNEq(litexpr.BoolLit{true}, litexpr.BoolLit{false}), litexpr.BoolLit{true})
}
