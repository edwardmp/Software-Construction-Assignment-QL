package expr

import (
	"ql/interfaces"
	"ql/symboltable"
	"testing"
)

func unaryExprEval(t *testing.T, exampleInput interfaces.Expr, expectedOutput interfaces.Expr, symbolTable interface{}) {
	if eval, expectedOutputEval := exampleInput.Eval(symbolTable), expectedOutput.(interfaces.Expr).Eval(symbolTable); eval != expectedOutputEval {
		t.Errorf("interfaces.Expr test error: should be %v (%T) for %v but is %v (%T)", expectedOutputEval, expectedOutputEval, eval, eval)
	}
}

/* Test for unary expressions */

func testNot(t *testing.T) {
	unaryExprEval(t, NewNot(NewBoolLit(true)), NewBoolLit(false), nil)
}

func testPos(t *testing.T) {
	unaryExprEval(t, NewPos(NewIntLit(-10)), NewIntLit(10), nil)
}

func testNeg(t *testing.T) {
	unaryExprEval(t, NewNeg(NewIntLit(10)), NewIntLit(-10), nil)
}

func testPosNeg(t *testing.T) {
	unaryExprEval(t, NewPos(NewNeg(NewIntLit(-10))), NewIntLit(10), nil)
}

func testNegPos(t *testing.T) {
	unaryExprEval(t, NewNeg(NewPos(NewIntLit(10))), NewIntLit(-10), nil)
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

func (v VarIdentifier) Accept(va interfaces.Visitor, s interface{}) interface{} {
	return nil
}

func (v VarIdentifier) TypeCheck(typeChecker interfaces.TypeChecker, symbolTable interfaces.SymbolTable) {

}

func TestVarExpr(t *testing.T) {
	symbolTable := symboltable.NewSymbolTable()
	symbolTable.SetNodeForIdentifier(NewIntLit(2), VarIdentifier{"TestIdentifier"})

	unaryExprEval(t, VarExpr{VarIdentifier{"TestIdentifier"}}, NewIntLit(2), symbolTable)
}
