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
	addition := Add{litexpr.IntLit{1}, litexpr.IntLit{2}}
	testExprEval(t, addition, litexpr.IntLit{3})
}

func TestMul(t *testing.T) {
	testExprEval(t, Mul{litexpr.IntLit{3}, litexpr.IntLit{2}}, litexpr.IntLit{6})
}

func TestMulAddPrecedence(t *testing.T) {
	testExprEval(t, Add{Mul{litexpr.IntLit{3}, litexpr.IntLit{2}}, litexpr.IntLit{1}}, litexpr.IntLit{7})
}

func TestSub(t *testing.T) {
	testExprEval(t, Sub{litexpr.IntLit{1}, litexpr.IntLit{2}}, litexpr.IntLit{-1})
}

func TestDiv(t *testing.T) {
	testExprEval(t, Div{litexpr.IntLit{9}, litexpr.IntLit{3}}, litexpr.IntLit{3})
}

func TestGT(t *testing.T) {
	testExprEval(t, GT{litexpr.IntLit{3}, litexpr.IntLit{2}}, litexpr.BoolLit{true})
}

func TestLT(t *testing.T) {
	testExprEval(t, LT{litexpr.IntLit{3}, litexpr.IntLit{2}}, litexpr.BoolLit{false})
}

func TestGEq(t *testing.T) {
	testExprEval(t, GEq{litexpr.IntLit{3}, litexpr.IntLit{3}}, litexpr.BoolLit{true})
}

func TestLEq(t *testing.T) {
	testExprEval(t, LEq{litexpr.IntLit{3}, litexpr.IntLit{3}}, litexpr.BoolLit{true})
}

func TestAnd(t *testing.T) {
	testExprEval(t, And{litexpr.BoolLit{true}, litexpr.BoolLit{false}}, litexpr.BoolLit{false})
}

func TestOr(t *testing.T) {
	testExprEval(t, Or{litexpr.BoolLit{true}, litexpr.BoolLit{false}}, litexpr.BoolLit{true})
}

func TestAndOr(t *testing.T) {
	testExprEval(t, And{Or{litexpr.BoolLit{true}, litexpr.BoolLit{false}}, And{litexpr.BoolLit{true}, litexpr.BoolLit{false}}}, litexpr.BoolLit{false})
}

func TestEq(t *testing.T) {
	testExprEval(t, Eq{litexpr.BoolLit{true}, litexpr.BoolLit{false}}, litexpr.BoolLit{false})
}

func TestNEq(t *testing.T) {
	testExprEval(t, NEq{litexpr.BoolLit{true}, litexpr.BoolLit{false}}, litexpr.BoolLit{true})
}
