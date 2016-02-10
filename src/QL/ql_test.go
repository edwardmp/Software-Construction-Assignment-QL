package QL

import (
	"QL/ast/expr"
	"QL/lexer"
	"QL/parser"
	"QL/token"
	"fmt"
	//"io/ioutil"
	"testing"
)

func testExprEval(t *testing.T, exampleExpr interface{}, expectedOutput interface{}) {
	if eval, expectedOutputEval := exampleExpr.(expr.Expr).Eval(), expectedOutput.(expr.Expr).Eval(); eval != expectedOutputEval {
		t.Fatalf("Should be %v (%T) for %v but is %v (%T)", expectedOutputEval, expectedOutputEval, exampleExpr, eval, eval)
	}
}

func testStmt(t *testing.T, exampleStmt string, expectedOutput interface{}) {
	lex := lexer.NewLexer([]byte(exampleStmt))
	//printLexerTokens(lex)
	p := parser.NewParser()
	_, err := p.Parse(lex)

	if err != nil {
		panic(err)
	}
}

func printLexerTokens(lexer *lexer.Lexer) {
	for {

		d := lexer.Scan()
		if d.Type == 1 {
			break
		}

		fmt.Println(token.TokMap.Id(d.Type))
	}
}

func TestAdd(t *testing.T) {
	addition := expr.Add{expr.IntLit{1}, expr.IntLit{2}}
	testExprEval(t, addition, expr.IntLit{3})
}

func TestMul(t *testing.T) {
	testExprEval(t, expr.Mul{expr.IntLit{3}, expr.IntLit{2}}, expr.IntLit{6})
}

func TestMulAddPrecedence(t *testing.T) {
	testExprEval(t, expr.Add{expr.Mul{expr.IntLit{3}, expr.IntLit{2}}, expr.IntLit{1}}, expr.IntLit{7})
}

func TestSub(t *testing.T) {
	testExprEval(t, expr.Sub{expr.IntLit{1}, expr.IntLit{2}}, expr.IntLit{-1})
}

func TestDiv(t *testing.T) {
	testExprEval(t, expr.Div{expr.IntLit{9}, expr.IntLit{3}}, expr.IntLit{3})
}

func TestGT(t *testing.T) {
	testExprEval(t, expr.GT{expr.IntLit{3}, expr.IntLit{2}}, expr.BoolLit{true})
}

func TestLT(t *testing.T) {
	testExprEval(t, expr.LT{expr.IntLit{3}, expr.IntLit{2}}, expr.BoolLit{false})
}

func TestGEq(t *testing.T) {
	testExprEval(t, expr.GEq{expr.IntLit{3}, expr.IntLit{3}}, expr.BoolLit{true})
}

func TestLEq(t *testing.T) {
	testExprEval(t, expr.LEq{expr.IntLit{3}, expr.IntLit{3}}, expr.BoolLit{true})
}

func TestAnd(t *testing.T) {
	testExprEval(t, expr.And{expr.BoolLit{true}, expr.BoolLit{false}}, expr.BoolLit{false})
}

func TestOr(t *testing.T) {
	testExprEval(t, expr.Or{expr.BoolLit{true}, expr.BoolLit{false}}, expr.BoolLit{true})
}

func TestAndOr(t *testing.T) {
	testExprEval(t, expr.And{expr.Or{expr.BoolLit{true}, expr.BoolLit{false}}, expr.And{expr.BoolLit{true}, expr.BoolLit{false}}}, expr.BoolLit{false})
}

func TestEq(t *testing.T) {
	testExprEval(t, expr.Eq{expr.BoolLit{true}, expr.BoolLit{false}}, expr.BoolLit{false})
}

func TestNEq(t *testing.T) {
	testExprEval(t, expr.NEq{expr.BoolLit{true}, expr.BoolLit{false}}, expr.BoolLit{true})
}

func TestPos(t *testing.T) {
	testExprEval(t, expr.Pos{expr.IntLit{-10}}, expr.IntLit{10})
}

func TestNeg(t *testing.T) {
	testExprEval(t, expr.Neg{expr.IntLit{10}}, expr.IntLit{-10})
}

func TestPosNeg(t *testing.T) {
	testExprEval(t, expr.Pos{expr.Neg{expr.IntLit{-10}}}, expr.IntLit{10})
}

func TestNegPos(t *testing.T) {
	testExprEval(t, expr.Neg{expr.Pos{expr.IntLit{10}}}, expr.IntLit{-10})
}
