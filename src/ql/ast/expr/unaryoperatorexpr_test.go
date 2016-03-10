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

func TestNot(t *testing.T) {
	unaryExprEval(t, NewNot(BoolLit{true}), BoolLit{false}, nil)
}

func TestPos(t *testing.T) {
	unaryExprEval(t, NewPos(IntLit{-10}), IntLit{10}, nil)
}

func TestNeg(t *testing.T) {
	unaryExprEval(t, NewNeg(IntLit{10}), IntLit{-10}, nil)
}

func TestPosNeg(t *testing.T) {
	unaryExprEval(t, NewPos(NewNeg(IntLit{-10})), IntLit{10}, nil)
}

func TestNegPos(t *testing.T) {
	unaryExprEval(t, NewNeg(NewPos(IntLit{10})), IntLit{-10}, nil)
}

// TODO is this a good idea?
type VarIdentifier struct {
	Ident string
}

func (v VarIdentifier) GetIdent() string {
	return v.Ident
}

func (v VarIdentifier) String() string {
	return v.Ident
}

func TestVarExpr(t *testing.T) {
	symbolTable := symboltable.NewSymbolTable()
	symbolTable.SetNodeForIdentifier(IntLit{2}, VarIdentifier{"TestIdentifier"})

	unaryExprEval(t, VarExpr{VarIdentifier{"TestIdentifier"}}, IntLit{2}, symbolTable)
}
