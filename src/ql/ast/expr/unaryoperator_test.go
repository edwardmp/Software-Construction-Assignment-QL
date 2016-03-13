package expr

import (
	"ql/interfaces"
	"ql/symbols"
	"ql/token"
	"testing"
)

func unaryExprEval(t *testing.T, exampleInput interfaces.Expr, expectedOutput interfaces.Expr, symbols interfaces.Symbols) {
	if eval, expectedOutputEval := exampleInput.Eval(symbols), expectedOutput.(interfaces.Expr).Eval(symbols); eval != expectedOutputEval {
		t.Errorf("interfaces.Expr test error: should be %v (%T) for %v but is %v (%T)", expectedOutputEval, expectedOutputEval, eval, eval)
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

// FIXME is this a good idea? If not leads to import cycle..
type VarIdentifier struct {
	Ident string
}

func (v VarIdentifier) GetIdent() string {
	return v.Ident
}

func (v VarIdentifier) String() string {
	return v.Ident
}

func (v VarIdentifier) Accept(va interfaces.Visitor, s interfaces.Symbols) interface{} {
	return nil
}

func (v VarIdentifier) TypeCheck(typeChecker interfaces.TypeChecker, symbols interfaces.Symbols) {

}

func (v VarIdentifier) GetSourceInfo() token.Pos {
	return token.Pos{}
}

func (v VarIdentifier) ResetSourceInfo() {
}

func TestVarExpr(t *testing.T) {
	symbols := symbols.NewSymbols()
	symbols.SetNodeForIdentifier(NewIntLitNoSourceInfo(2), VarIdentifier{"TestIdentifier"})

	unaryExprEval(t, NewVarExprNoSourceInfo(VarIdentifier{"TestIdentifier"}), NewIntLitNoSourceInfo(2), symbols)
}
