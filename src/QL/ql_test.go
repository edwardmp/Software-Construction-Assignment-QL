package QL

import (
	"fmt"
	"ql/ast/expr"
	"ql/ast/expr/binaryoperatorexpr"
	"ql/ast/expr/lit"
	"ql/ast/expr/unaryoperatorexpr"
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

	for i := range questionsA {
		if questionsA[i] != questionsB[i] {
			return false
		}
	}

	conditionalsA := a.Conditionals
	conditionalsB := b.Conditionals

	for i := range conditionalsA {
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

type VisitorAdapter struct {
}

func (v VisitorAdapter) Visit(t interface{}) interface{} {
	switch t.(type) {
	default:
		panic(fmt.Sprintf("Unexpected node type %T\n", t))
	case stmt.Form:
		t.(stmt.Form).Identifier.Accept(v)
		t.(stmt.Form).Content.Accept(v)
	case vari.VarId:
		fmt.Printf("Varid") // TODO handle
	case vari.VarDecl:
		fmt.Printf("VarDecl") // TODO handle
		t.(vari.VarDecl).Ident.Accept(v)
	case stmt.StmtList:
		for i := range t.(stmt.StmtList).Questions {
			t.(stmt.StmtList).Questions[i].(stmt.Question).Accept(v)
		}
		for i := range t.(stmt.StmtList).Conditionals {
			t.(stmt.StmtList).Conditionals[i].(stmt.Conditional).Accept(v)
		}
	case stmt.InputQuestion:
		t.(stmt.InputQuestion).Label.Accept(v)
		t.(stmt.InputQuestion).VarDecl.Accept(v)
	case stmt.ComputedQuestion:
		t.(stmt.ComputedQuestion).Label.Accept(v)
		t.(stmt.ComputedQuestion).VarDecl.Accept(v)
		t.(stmt.ComputedQuestion).Computation.Accept(v)
	case stmt.If:
		t.(stmt.If).Cond.Accept(v)
		t.(stmt.If).Body.Accept(v)
	case stmt.IfElse:
		t.(stmt.IfElse).Cond.Accept(v)
		t.(stmt.IfElse).IfBody.Accept(v)
		t.(stmt.IfElse).ElseBody.Accept(v)
	case lit.StrLit:
		fmt.Printf("StrLit") // TODO handle
	case lit.BoolLit:
		fmt.Printf("BoolLit") // TODO handle
	case lit.IntLit:
		fmt.Printf("IntLit") // TODO handle
	case binaryoperatorexpr.BinaryOperatorExpr:
		t.(binaryoperatorexpr.BinaryOperatorExpr).GetLhs().(expr.Expr).Accept(v)
		t.(binaryoperatorexpr.BinaryOperatorExpr).GetRhs().(expr.Expr).Accept(v)
	case unaryoperatorexpr.UnaryOperatorExpr:
		t.(unaryoperatorexpr.UnaryOperatorExpr).GetValue().(expr.Expr).Accept(v)
	}

	return false
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
		f.Accept(VisitorAdapter{})
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
				for i := range fCond {
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

	exampleOutputForm := stmt.Form{vari.VarId{"TestForm"}, stmt.StmtList{}}
	testStmtParse(t, exampleEmptyForm, exampleOutputForm)
}

func TestFormQuestion(t *testing.T) {
	exampleFormInput := "form TestForm { \"Did you sell a house in 2010?\" hasSoldHouse: boolean \"Did you enter a loan?\" hasMaintLoan: boolean }"

	firstQuestionOutput := stmt.InputQuestion{lit.StrLit{"Did you sell a house in 2010?"}, vari.VarDecl{vari.VarId{"hasSoldHouse"}, vari.BOOLEAN}}
	secondQuestionOutput := stmt.InputQuestion{lit.StrLit{"Did you enter a loan?"}, vari.VarDecl{vari.VarId{"hasMaintLoan"}, vari.BOOLEAN}}
	exampleBodyOutput := stmt.StmtList{[]stmt.Question{firstQuestionOutput, secondQuestionOutput}, []stmt.Conditional{}}
	exampleOutputForm := stmt.Form{vari.VarId{"TestForm"}, exampleBodyOutput}

	testStmtParse(t, exampleFormInput, exampleOutputForm)
}

/*
func TestFormComputedQuestion(t *testing.T) {
	exampleFormInput := "form TestForm { \"Did you sell a house in 2010?\" hasSoldHouse: integer \"Did you enter a loan?\" hasMaintLoan: integer \"Value residue:\" valueResidue: int = (hasSoldHouse - hasMaintLoan) }"

	firstQuestionOutput := stmt.InputQuestion{lit.StrLit{"Did you sell a house in 2010?"}, vari.VarDecl{vari.VarId{"hasSoldHouse"}, vari.INT}}
	secondQuestionOutput := stmt.InputQuestion{lit.StrLit{"Did you enter a loan?"}, vari.VarDecl{vari.VarId{"hasMaintLoan"}, vari.INT}}
	computedQuestion := stmt.ComputedQuestion{lit.StrLit{"Value residue:"}, vari.VarDecl{vari.VarId{"valueResidue"}, vari.INT}, expr.Sub{vari.VarId{"hasSoldHouse"}, vari.Varid{"hasMaintLoan"}}}
	exampleBodyOutput := stmt.StmtList{[]stmt.Question{firstQuestionOutput, secondQuestionOutput, computedQuestion}, []stmt.Conditional{}}
	exampleOutputForm := stmt.Form{vari.VarId{"TestForm"}, exampleBodyOutput}

	testStmtParse(t, exampleFormInput, exampleOutputForm)
}
*/

func TestFormIf(t *testing.T) {
	exampleFormInput := "form TestForm { \"Did you sell a house in 2010?\" hasSoldHouse: boolean if (true) { \"What was the selling price?\" sellingPrice: money } }"

	firstQuestionOutput := stmt.InputQuestion{lit.StrLit{"Did you sell a house in 2010?"}, vari.VarDecl{vari.VarId{"hasSoldHouse"}, vari.BOOLEAN}}
	firstQuestionBodyInput := stmt.InputQuestion{lit.StrLit{"What was the selling price?"}, vari.VarDecl{vari.VarId{"sellingPrice"}, vari.MONEY}}
	ifBodyOutput := stmt.StmtList{[]stmt.Question{firstQuestionBodyInput}, []stmt.Conditional{}}
	ifOutput := stmt.If{lit.BoolLit{true}, ifBodyOutput}
	exampleBodyOutput := stmt.StmtList{[]stmt.Question{firstQuestionOutput}, []stmt.Conditional{ifOutput}}
	exampleOutputForm := stmt.Form{vari.VarId{"TestForm"}, exampleBodyOutput}

	testStmtParse(t, exampleFormInput, exampleOutputForm)
}

func TestFormIfElse(t *testing.T) {
	exampleFormInput := "form TestForm { \"Did you sell a house in 2010?\" hasSoldHouse: boolean if (true) { \"What was the selling price?\" sellingPrice: money } else { \"What was the selling price?\" sellingPrice: money } }"

	firstQuestionOutput := stmt.InputQuestion{lit.StrLit{"Did you sell a house in 2010?"}, vari.VarDecl{vari.VarId{"hasSoldHouse"}, vari.BOOLEAN}}
	firstQuestionBodyInput := stmt.InputQuestion{lit.StrLit{"What was the selling price?"}, vari.VarDecl{vari.VarId{"sellingPrice"}, vari.MONEY}}
	ifBodyOutput := stmt.StmtList{[]stmt.Question{firstQuestionBodyInput}, []stmt.Conditional{}}
	elseBodyOutput := stmt.StmtList{[]stmt.Question{firstQuestionBodyInput}, []stmt.Conditional{}}
	ifOutput := stmt.IfElse{lit.BoolLit{true}, ifBodyOutput, elseBodyOutput}
	exampleBodyOutput := stmt.StmtList{[]stmt.Question{firstQuestionOutput}, []stmt.Conditional{ifOutput}}
	exampleOutputForm := stmt.Form{vari.VarId{"TestForm"}, exampleBodyOutput}

	testStmtParse(t, exampleFormInput, exampleOutputForm)
}
