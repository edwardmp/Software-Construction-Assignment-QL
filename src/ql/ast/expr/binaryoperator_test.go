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

func testAdd(t *testing.T) {
	binaryExprEval(t, NewAdd(IntLit{1}, IntLit{2}), IntLit{3})
}

func testMul(t *testing.T) {
	binaryExprEval(t, NewMul(IntLit{3}, IntLit{2}), IntLit{6})
}

func testMulAddPrecedence(t *testing.T) {
	binaryExprEval(t, NewAdd(NewMul(IntLit{3}, IntLit{2}), IntLit{1}), IntLit{7})
}

func testSub(t *testing.T) {
	binaryExprEval(t, NewSub(IntLit{1}, IntLit{2}), IntLit{-1})
}

func testDiv(t *testing.T) {
	binaryExprEval(t, NewDiv(IntLit{9}, IntLit{3}), IntLit{3})
}

func testGT(t *testing.T) {
	binaryExprEval(t, NewGT(IntLit{3}, IntLit{2}), BoolLit{true})
}

func testLT(t *testing.T) {
	binaryExprEval(t, NewLT(IntLit{3}, IntLit{2}), BoolLit{false})
}

func testGEq(t *testing.T) {
	binaryExprEval(t, NewGEq(IntLit{3}, IntLit{3}), BoolLit{true})
}

func testLEq(t *testing.T) {
	binaryExprEval(t, NewLEq(IntLit{3}, IntLit{3}), BoolLit{true})
}

func testAnd(t *testing.T) {
	binaryExprEval(t, NewAnd(BoolLit{true}, BoolLit{false}), BoolLit{false})
}

func testOr(t *testing.T) {
	binaryExprEval(t, NewOr(BoolLit{true}, BoolLit{false}), BoolLit{true})
}

func testAndOr(t *testing.T) {
	binaryExprEval(t, NewAnd(NewOr(BoolLit{true}, BoolLit{false}), NewAnd(BoolLit{true}, BoolLit{false})), BoolLit{false})
}

func testEq(t *testing.T) {
	binaryExprEval(t, NewEq(BoolLit{true}, BoolLit{false}), BoolLit{false})
}

func testNEq(t *testing.T) {
	binaryExprEval(t, NewNEq(BoolLit{true}, BoolLit{false}), BoolLit{true})
}
