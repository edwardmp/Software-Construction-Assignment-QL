package QL

import (
	"fmt"
	"ql/ast/expr"
	"ql/ast/stmt"
	"ql/ast/vari"
	"ql/lexer"
	"ql/parser"
	"ql/token"
	//"io/ioutil"
	"testing"
)

// slices don't support equality checking, so have to do it manually
func slicesEqual(a stmt.StmtList, b stmt.StmtList) bool {
	questionsA := a.Questions
	questionsB := b.Questions

	if len(questionsA) != len(questionsB) {
		return false
	}

	for i, _ := range questionsA {
		if questionsA[i] != questionsB[i] {
			return false
		}
	}

	conditionalsA := a.Conditionals
	conditionalsB := b.Conditionals

	for i, _ := range conditionalsA {
		if conditionalsA[i].(stmt.Conditional).EvalCondition() != conditionalsB[i].(stmt.Conditional).EvalCondition() {
			return false
		}

		if !slicesEqualConditional(conditionalsA[i].(stmt.Conditional), conditionalsB[i].(stmt.Conditional)) {
			return true
		}
	}

	return true
}

func slicesEqualConditional(ifA, ifB stmt.Conditional) bool {
	if fmt.Sprintf("%T", ifA) != fmt.Sprintf("%T", ifB) {
		panic("Types not equal") // TODO replace with assert
	}

	switch t := ifA.(type) {
	default:
		panic(fmt.Sprintf("unexpected Conditional type %T\n", t))
	case stmt.If:
		bodyA := ifA.(stmt.If).Body
		bodyB := ifB.(stmt.If).Body
		return slicesEqual(bodyA, bodyB)
	case stmt.IfElse:
		return slicesEqual(ifA.(stmt.IfElse).IfBody, ifB.(stmt.IfElse).IfBody) && slicesEqual(ifA.(stmt.IfElse).ElseBody, ifB.(stmt.IfElse).ElseBody)
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

func testExprEval(t *testing.T, exampleExpr interface{}, expectedOutput interface{}) {
	if eval, expectedOutputEval := exampleExpr.(expr.Expr).Eval(), expectedOutput.(expr.Expr).Eval(); eval != expectedOutputEval {
		t.Fatalf("Should be %v (%T) for %v but is %v (%T)", expectedOutputEval, expectedOutputEval, exampleExpr, eval, eval)
	}
}

func testStmtParse(t *testing.T, stmtAsString string, expectedOutput interface{}) stmt.Form {
	lex := lexer.NewLexer([]byte(stmtAsString))
	//printLexerTokens(lex)
	p := parser.NewParser()
	r, err := p.Parse(lex)

	if err != nil {
		panic(err)
	}

	if f, fOk := r.(stmt.Form); fOk {
		if e, eOk := expectedOutput.(stmt.Form); eOk {
			if f.Identifier != e.Identifier {
				t.Errorf("Form identifiers not equal")
			}

			if !slicesEqual(f.Content, e.Content) {
				t.Errorf("Form content not equal %v %v", f.Content, e.Content)
			}

			fCond := f.Content.Conditionals
			eCond := e.Content.Conditionals
			if len(fCond) != 0 && len(fCond) == len(eCond) {
				for i, _ := range fCond {
					if !slicesEqualConditional(fCond[i].(stmt.Conditional), eCond[i].(stmt.Conditional)) {
						t.Errorf("parse test failed conditionals not equal: %v %v", fCond, eCond)
					}
				}
			}
		}
	}

	return r.(stmt.Form)
}

/* Tests for statements */

func TestFormIdentifierExtraction(t *testing.T) {
	exampleEmptyForm := "form TestForm {}"

	var exampleOutputForm stmt.Form = stmt.Form{"TestForm", stmt.StmtList{}}
	testStmtParse(t, exampleEmptyForm, exampleOutputForm)
}

func TestFormQuestion(t *testing.T) {
	exampleFormInput := "form TestForm { \"Did you sell a house in 2010?\" hasSoldHouse: boolean \"Did you enter a loan?\" hasMaintLoan: boolean }"

	firstQuestionOutput := stmt.InputQuestion{expr.StrLit{"Did you sell a house in 2010?"}, vari.VarDecl{vari.VarId{"hasSoldHouse"}, vari.BOOLEAN}}
	secondQuestionOutput := stmt.InputQuestion{expr.StrLit{"Did you enter a loan?"}, vari.VarDecl{vari.VarId{"hasMaintLoan"}, vari.BOOLEAN}}
	exampleBodyOutput := stmt.StmtList{[]stmt.Question{firstQuestionOutput, secondQuestionOutput}, []stmt.Conditional{}}
	exampleOutputForm := stmt.Form{"TestForm", exampleBodyOutput}

	testStmtParse(t, exampleFormInput, exampleOutputForm)
}

/*
func TestFormComputedQuestion(t *testing.T) {
	exampleFormInput := "form TestForm { \"Did you sell a house in 2010?\" hasSoldHouse: integer \"Did you enter a loan?\" hasMaintLoan: integer \"Value residue:\" valueResidue: int = (hasSoldHouse - hasMaintLoan) }"

	firstQuestionOutput := stmt.InputQuestion{expr.StrLit{"Did you sell a house in 2010?"}, vari.VarDecl{vari.VarId{"hasSoldHouse"}, vari.INT}}
	secondQuestionOutput := stmt.InputQuestion{expr.StrLit{"Did you enter a loan?"}, vari.VarDecl{vari.VarId{"hasMaintLoan"}, vari.INT}}
	computedQuestion := stmt.ComputedQuestion{expr.StrLit{"Value residue:"}, vari.VarDecl{vari.VarId{"valueResidue"}, vari.INT}, expr.Sub{vari.VarId{"hasSoldHouse"}, vari.Varid{"hasMaintLoan"}}}
	exampleBodyOutput := stmt.StmtList{[]stmt.Question{firstQuestionOutput, secondQuestionOutput, computedQuestion}, []stmt.Conditional{}}
	exampleOutputForm := stmt.Form{"TestForm", exampleBodyOutput}

	testStmtParse(t, exampleFormInput, exampleOutputForm)
}
*/

func TestFormIf(t *testing.T) {
	exampleFormInput := "form TestForm { \"Did you sell a house in 2010?\" hasSoldHouse: boolean if (true) { \"What was the selling price?\" sellingPrice: money } }"

	firstQuestionOutput := stmt.InputQuestion{expr.StrLit{"Did you sell a house in 2010?"}, vari.VarDecl{vari.VarId{"hasSoldHouse"}, vari.BOOLEAN}}
	firstQuestionBodyInput := stmt.InputQuestion{expr.StrLit{"What was the selling price?"}, vari.VarDecl{vari.VarId{"sellingPrice"}, vari.MONEY}}
	ifBodyOutput := stmt.StmtList{[]stmt.Question{firstQuestionBodyInput}, []stmt.Conditional{}}
	ifOutput := stmt.If{expr.BoolLit{true}, ifBodyOutput}
	exampleBodyOutput := stmt.StmtList{[]stmt.Question{firstQuestionOutput}, []stmt.Conditional{ifOutput}}
	exampleOutputForm := stmt.Form{"TestForm", exampleBodyOutput}

	testStmtParse(t, exampleFormInput, exampleOutputForm)
}

func TestFormIfElse(t *testing.T) {
	exampleFormInput := "form TestForm { \"Did you sell a house in 2010?\" hasSoldHouse: boolean if (true) { \"What was the selling price?\" sellingPrice: money } else { \"What was the selling price?\" sellingPrice: money } }"

	firstQuestionOutput := stmt.InputQuestion{expr.StrLit{"Did you sell a house in 2010?"}, vari.VarDecl{vari.VarId{"hasSoldHouse"}, vari.BOOLEAN}}
	firstQuestionBodyInput := stmt.InputQuestion{expr.StrLit{"What was the selling price?"}, vari.VarDecl{vari.VarId{"sellingPrice"}, vari.MONEY}}
	ifBodyOutput := stmt.StmtList{[]stmt.Question{firstQuestionBodyInput}, []stmt.Conditional{}}
	elseBodyOutput := stmt.StmtList{[]stmt.Question{firstQuestionBodyInput}, []stmt.Conditional{}}
	ifOutput := stmt.IfElse{expr.BoolLit{true}, ifBodyOutput, elseBodyOutput}
	exampleBodyOutput := stmt.StmtList{[]stmt.Question{firstQuestionOutput}, []stmt.Conditional{ifOutput}}
	exampleOutputForm := stmt.Form{"TestForm", exampleBodyOutput}

	testStmtParse(t, exampleFormInput, exampleOutputForm)
}

/* Tests for expressions */

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

func TestNot(t *testing.T) {
	testExprEval(t, expr.Not{expr.BoolLit{true}}, expr.BoolLit{false})
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
