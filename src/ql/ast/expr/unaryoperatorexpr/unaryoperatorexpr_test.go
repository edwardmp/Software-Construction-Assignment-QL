package unaryoperatorexpr

import (
	"ql/ast/expr"
	"ql/ast/expr/litexpr"
	"ql/ast/vari"
	"ql/symboltable"
	"testing"
)

func testExprEval(t *testing.T, exampleExpr interface{}, expectedOutput interface{}, symbolTable interface{}) {
	if eval, expectedOutputEval := exampleExpr.(expr.Expr).Eval(symbolTable), expectedOutput.(expr.Expr).Eval(symbolTable); eval != expectedOutputEval {
		t.Errorf("Expr test error: should be %v (%T) for %v but is %v (%T)", expectedOutputEval, expectedOutputEval, exampleExpr, eval, eval)
	}
}

/* Test for unary expressions */

func TestNot(t *testing.T) {
	testExprEval(t, Not{litexpr.BoolLit{true}}, litexpr.BoolLit{false}, nil)
}

func TestPos(t *testing.T) {
	testExprEval(t, Pos{litexpr.IntLit{-10}}, litexpr.IntLit{10}, nil)
}

func TestNeg(t *testing.T) {
	testExprEval(t, Neg{litexpr.IntLit{10}}, litexpr.IntLit{-10}, nil)
}

func TestPosNeg(t *testing.T) {
	testExprEval(t, Pos{Neg{litexpr.IntLit{-10}}}, litexpr.IntLit{10}, nil)
}

func TestNegPos(t *testing.T) {
	testExprEval(t, Neg{Pos{litexpr.IntLit{10}}}, litexpr.IntLit{-10}, nil)
}

func TestVarExpr(t *testing.T) {
	symbolTable := symboltable.NewSymbolTable()
	symbolTable.SetNodeForIdentifier(litexpr.IntLit{2}, vari.VarId{"TestIdentifier"})

	testExprEval(t, VarExpr{vari.VarId{"TestIdentifier"}}, litexpr.IntLit{2}, symbolTable)
}
