package unaryoperatorexpr

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

/* Test for unary expressions */

func TestNot(t *testing.T) {
	testExprEval(t, Not{litexpr.BoolLit{true}}, litexpr.BoolLit{false})
}

func TestPos(t *testing.T) {
	testExprEval(t, Pos{litexpr.IntLit{-10}}, litexpr.IntLit{10})
}

func TestNeg(t *testing.T) {
	testExprEval(t, Neg{litexpr.IntLit{10}}, litexpr.IntLit{-10})
}

func TestPosNeg(t *testing.T) {
	testExprEval(t, Pos{Neg{litexpr.IntLit{-10}}}, litexpr.IntLit{10})
}

func TestNegPos(t *testing.T) {
	testExprEval(t, Neg{Pos{litexpr.IntLit{10}}}, litexpr.IntLit{-10})
}

func TestVarExpr(t *testing.T) {
}
