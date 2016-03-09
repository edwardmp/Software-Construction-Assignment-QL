package unaryoperatorexpr

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

/* Test for unary expressions */

func TestNot(t *testing.T) {
	testExprEval(t, Not{lit.BoolLit{true}}, lit.BoolLit{false})
}

func TestPos(t *testing.T) {
	testExprEval(t, Pos{lit.IntLit{-10}}, lit.IntLit{10})
}

func TestNeg(t *testing.T) {
	testExprEval(t, Neg{lit.IntLit{10}}, lit.IntLit{-10})
}

func TestPosNeg(t *testing.T) {
	testExprEval(t, Pos{Neg{lit.IntLit{-10}}}, lit.IntLit{10})
}

func TestNegPos(t *testing.T) {
	testExprEval(t, Neg{Pos{lit.IntLit{10}}}, lit.IntLit{-10})
}
