package expr

import (
	"ql/ast/vari"
	"ql/interfaces"
	"ql/symbols"
	"testing"
)

func unaryExprEval(t *testing.T, testExpr interfaces.Expr, expectedExpr interfaces.Expr, symbols interfaces.VarIdValueSymbols) {
	if actualEvalValue, expectedEvalValue := testExpr.Eval(symbols), expectedExpr.Eval(symbols); actualEvalValue != expectedEvalValue {
		t.Errorf("UnaryOperator test error: should be %v (%T) for %v but is %v (%T)", expectedEvalValue, expectedEvalValue, actualEvalValue, actualEvalValue)
	}
}

/* Tests for unary expressions */

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
