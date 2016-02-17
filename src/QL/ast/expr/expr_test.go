package expr

import (
	"ql/ast/expr"
	"ql/ast/expr/binaryoperatorexpr"
	"ql/ast/expr/lit"
	"ql/ast/expr/unaryoperatorexpr"
	"testing"
)

func testExprEval(t *testing.T, exampleExpr interface{}, expectedOutput interface{}) {
	if eval, expectedOutputEval := exampleExpr.(expr.Expr).Eval(), expectedOutput.(expr.Expr).Eval(); eval != expectedOutputEval {
		t.Fatalf("Should be %v (%T) for %v but is %v (%T)", expectedOutputEval, expectedOutputEval, exampleExpr, eval, eval)
	}
}

/* Tests for binary expressions */

func TestAdd(t *testing.T) {
	addition := binaryoperatorexpr.Add{lit.IntLit{1}, lit.IntLit{2}}
	testExprEval(t, addition, lit.IntLit{3})
}

func TestMul(t *testing.T) {
	testExprEval(t, binaryoperatorexpr.Mul{lit.IntLit{3}, lit.IntLit{2}}, lit.IntLit{6})
}

func TestMulAddPrecedence(t *testing.T) {
	testExprEval(t, binaryoperatorexpr.Add{binaryoperatorexpr.Mul{lit.IntLit{3}, lit.IntLit{2}}, lit.IntLit{1}}, lit.IntLit{7})
}

func TestSub(t *testing.T) {
	testExprEval(t, binaryoperatorexpr.Sub{lit.IntLit{1}, lit.IntLit{2}}, lit.IntLit{-1})
}

func TestDiv(t *testing.T) {
	testExprEval(t, binaryoperatorexpr.Div{lit.IntLit{9}, lit.IntLit{3}}, lit.IntLit{3})
}

func TestGT(t *testing.T) {
	testExprEval(t, binaryoperatorexpr.GT{lit.IntLit{3}, lit.IntLit{2}}, lit.BoolLit{true})
}

func TestLT(t *testing.T) {
	testExprEval(t, binaryoperatorexpr.LT{lit.IntLit{3}, lit.IntLit{2}}, lit.BoolLit{false})
}

func TestGEq(t *testing.T) {
	testExprEval(t, binaryoperatorexpr.GEq{lit.IntLit{3}, lit.IntLit{3}}, lit.BoolLit{true})
}

func TestLEq(t *testing.T) {
	testExprEval(t, binaryoperatorexpr.LEq{lit.IntLit{3}, lit.IntLit{3}}, lit.BoolLit{true})
}

func TestAnd(t *testing.T) {
	testExprEval(t, binaryoperatorexpr.And{lit.BoolLit{true}, lit.BoolLit{false}}, lit.BoolLit{false})
}

func TestOr(t *testing.T) {
	testExprEval(t, binaryoperatorexpr.Or{lit.BoolLit{true}, lit.BoolLit{false}}, lit.BoolLit{true})
}

func TestAndOr(t *testing.T) {
	testExprEval(t, binaryoperatorexpr.And{binaryoperatorexpr.Or{lit.BoolLit{true}, lit.BoolLit{false}}, binaryoperatorexpr.And{lit.BoolLit{true}, lit.BoolLit{false}}}, lit.BoolLit{false})
}

func TestEq(t *testing.T) {
	testExprEval(t, binaryoperatorexpr.Eq{lit.BoolLit{true}, lit.BoolLit{false}}, lit.BoolLit{false})
}

func TestNEq(t *testing.T) {
	testExprEval(t, binaryoperatorexpr.NEq{lit.BoolLit{true}, lit.BoolLit{false}}, lit.BoolLit{true})
}

/* Test for unary expressions */

func TestNot(t *testing.T) {
	testExprEval(t, unaryoperatorexpr.Not{lit.BoolLit{true}}, lit.BoolLit{false})
}

func TestPos(t *testing.T) {
	testExprEval(t, unaryoperatorexpr.Pos{lit.IntLit{-10}}, lit.IntLit{10})
}

func TestNeg(t *testing.T) {
	testExprEval(t, unaryoperatorexpr.Neg{lit.IntLit{10}}, lit.IntLit{-10})
}

func TestPosNeg(t *testing.T) {
	testExprEval(t, unaryoperatorexpr.Pos{unaryoperatorexpr.Neg{lit.IntLit{-10}}}, lit.IntLit{10})
}

func TestNegPos(t *testing.T) {
	testExprEval(t, unaryoperatorexpr.Neg{unaryoperatorexpr.Pos{lit.IntLit{10}}}, lit.IntLit{-10})
}
