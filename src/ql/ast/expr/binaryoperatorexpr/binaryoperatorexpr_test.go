package binaryoperatorexpr

import (
	"ql/ast/expr"
	"ql/ast/expr/lit"
	"testing"
)

func testExprEval(t *testing.T, exampleExpr interface{}, expectedOutput interface{}) {
	if eval, expectedOutputEval := exampleExpr.(expr.Expr).Eval(nil), expectedOutput.(expr.Expr).Eval(nil); eval != expectedOutputEval {
		t.Errorf("Expr test error: should be %v (%T) for %v but is %v (%T)", expectedOutputEval, expectedOutputEval, exampleExpr, eval, eval)
	}
}

/* Tests for binary expressions */

func TestAdd(t *testing.T) {
	addition := Add{lit.IntLit{1}, lit.IntLit{2}}
	testExprEval(t, addition, lit.IntLit{3})
}

func TestMul(t *testing.T) {
	testExprEval(t, Mul{lit.IntLit{3}, lit.IntLit{2}}, lit.IntLit{6})
}

func TestMulAddPrecedence(t *testing.T) {
	testExprEval(t, Add{Mul{lit.IntLit{3}, lit.IntLit{2}}, lit.IntLit{1}}, lit.IntLit{7})
}

func TestSub(t *testing.T) {
	testExprEval(t, Sub{lit.IntLit{1}, lit.IntLit{2}}, lit.IntLit{-1})
}

func TestDiv(t *testing.T) {
	testExprEval(t, Div{lit.IntLit{9}, lit.IntLit{3}}, lit.IntLit{3})
}

func TestGT(t *testing.T) {
	testExprEval(t, GT{lit.IntLit{3}, lit.IntLit{2}}, lit.BoolLit{true})
}

func TestLT(t *testing.T) {
	testExprEval(t, LT{lit.IntLit{3}, lit.IntLit{2}}, lit.BoolLit{false})
}

func TestGEq(t *testing.T) {
	testExprEval(t, GEq{lit.IntLit{3}, lit.IntLit{3}}, lit.BoolLit{true})
}

func TestLEq(t *testing.T) {
	testExprEval(t, LEq{lit.IntLit{3}, lit.IntLit{3}}, lit.BoolLit{true})
}

func TestAnd(t *testing.T) {
	testExprEval(t, And{lit.BoolLit{true}, lit.BoolLit{false}}, lit.BoolLit{false})
}

func TestOr(t *testing.T) {
	testExprEval(t, Or{lit.BoolLit{true}, lit.BoolLit{false}}, lit.BoolLit{true})
}

func TestAndOr(t *testing.T) {
	testExprEval(t, And{Or{lit.BoolLit{true}, lit.BoolLit{false}}, And{lit.BoolLit{true}, lit.BoolLit{false}}}, lit.BoolLit{false})
}

func TestEq(t *testing.T) {
	testExprEval(t, Eq{lit.BoolLit{true}, lit.BoolLit{false}}, lit.BoolLit{false})
}

func TestNEq(t *testing.T) {
	testExprEval(t, NEq{lit.BoolLit{true}, lit.BoolLit{false}}, lit.BoolLit{true})
}
