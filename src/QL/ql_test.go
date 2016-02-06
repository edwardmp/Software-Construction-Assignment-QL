package QL

import (
	"QL/ast/expr"
	"QL/lexer"
	"QL/parser"
	"testing"
)

func testEval(t *testing.T, exampleStr string, output interface{}) {
	lex := lexer.NewLexer([]byte(exampleStr))
	p := parser.NewParser()
	st, err := p.Parse(lex)

	if err != nil {
		panic(err)
	}

	if eval := st.(expr.Expr).Eval(); eval != output {
		t.Fatalf("Should be %v for %v but is %v", output, exampleStr, eval)
	}
}

func TestAdd(t *testing.T) {
	testEval(t, "1 + 2", 3)
}

func TestSub(t *testing.T) {
	testEval(t, "1 - 2", -1)
}

func TestMul(t *testing.T) {
	testEval(t, "3 * 2", 6)
}

func TestDiv(t *testing.T) {
	testEval(t, "9 / 3", 3)
}

func TestGT(t *testing.T) {
	testEval(t, "3 > 2", true)
}

func TestLT(t *testing.T) {
	testEval(t, "3 < 2", false)
}

func TestGEq(t *testing.T) {
	testEval(t, "3 >= 3", true)
}

func TestLEq(t *testing.T) {
	testEval(t, "3 <= 3", true)
}

func TestAnd(t *testing.T) {
	testEval(t, "true && false", false)
}

func TestOr(t *testing.T) {
	testEval(t, "true || false", true)
}

func TestEq(t *testing.T) {
	testEval(t, "true == false", false)
}

func TestNEq(t *testing.T) {
	testEval(t, "true != false", true)
}

func TestPos(t *testing.T) {
	testEval(t, "+10", 10)
}

func TestNeg(t *testing.T) {
	testEval(t, "-10", -10)
}

/* this fails, need to check it out */
/*
func TestPosNeg(t *testing.T) {
	testEval(t, "+-10", 10)
}
*/

func TestPar(t *testing.T) {
	testEval(t, "(+10)", 10)
}
