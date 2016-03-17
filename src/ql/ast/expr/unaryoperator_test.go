package expr

import (
	"ql/ast/vari"
	"ql/interfaces"
	"ql/symbols"
	"testing"
)

func unaryExprEval(t *testing.T, exampleInput interfaces.Expr, expectedOutput interfaces.Expr, symbols interfaces.VarIdValueSymbols) {
	if eval, expectedOutputEval := exampleInput.Eval(symbols), expectedOutput.(interfaces.Expr).Eval(symbols); eval != expectedOutputEval {
		t.Errorf("UnaryOperator test error: should be %v (%T) for %v but is %v (%T)", expectedOutputEval, expectedOutputEval, eval, eval)
	}
}

/* Test for unary expressions */

func TestNot(t *testing.T) {
	unaryExprEval(t, NewNotNoSourceInfo(NewBoolLitNoSourceInfo(true)), NewBoolLitNoSourceInfo(false), nil)
}

func TestPos(t *testing.T) {
	unaryExprEval(t, NewPosNoSourceInfo(NewIntLitNoSourceInfo(-10)), NewIntLitNoSourceInfo(10), nil)
}

func TestNeg(t *testing.T) {
	unaryExprEval(t, NewNegNoSourceInfo(NewIntLitNoSourceInfo(10)), NewIntLitNoSourceInfo(-10), nil)
}

func TestPosNeg(t *testing.T) {
	unaryExprEval(t, NewPosNoSourceInfo(NewNegNoSourceInfo(NewIntLitNoSourceInfo(-10))), NewIntLitNoSourceInfo(10), nil)
}

func TestNegPos(t *testing.T) {
	unaryExprEval(t, NewNegNoSourceInfo(NewPosNoSourceInfo(NewIntLitNoSourceInfo(10))), NewIntLitNoSourceInfo(-10), nil)
}

func TestVarExpr(t *testing.T) {
	exampleVarId := vari.NewVarIdNoSourceInfo("TestIdentifier")

	symbols := symbols.NewVarIdValueSymbols()
	symbols.SetExprForVarId(NewIntLitNoSourceInfo(2), exampleVarId)

	unaryExprEval(t, NewVarExprNoSourceInfo(exampleVarId), NewIntLitNoSourceInfo(2), symbols)
}
