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
	unaryExprEval(t, NewNot(NewBoolLit(true)), NewBoolLit(false), nil)
}

func TestPos(t *testing.T) {
	unaryExprEval(t, NewPos(NewIntLit(-10)), NewIntLit(10), nil)
}

func TestNeg(t *testing.T) {
	unaryExprEval(t, NewNeg(NewIntLit(10)), NewIntLit(-10), nil)
}

func TestPosNeg(t *testing.T) {
	unaryExprEval(t, NewPos(NewNeg(NewIntLit(-10))), NewIntLit(10), nil)
}

func TestNegPos(t *testing.T) {
	unaryExprEval(t, NewNeg(NewPos(NewIntLit(10))), NewIntLit(-10), nil)
}

func TestVarExpr(t *testing.T) {
	exampleVarId := vari.NewVarId("TestIdentifier")

	symbols := symbols.NewVarIdValueSymbols()
	symbols.SetExprForVarId(NewIntLit(2), exampleVarId)

	unaryExprEval(t, NewVarExpr(exampleVarId), NewIntLit(2), symbols)
}
