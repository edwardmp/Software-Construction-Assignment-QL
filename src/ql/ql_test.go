package main

import (
	"fmt"
	"ql/ast/expr"
	"ql/ast/stmt"
	"ql/ast/vari"
	"ql/interfaces"
	"ql/lexer"
	"ql/parser"
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
		if conditionalsA[i].(interfaces.Conditional).EvalCondition() != conditionalsB[i].(interfaces.Conditional).EvalCondition() {
			return false
		}

		if !slicesEqualConditional(conditionalsA[i].(interfaces.Conditional), conditionalsB[i].(interfaces.Conditional)) {
			return true
		}
	}

	return true
}

func slicesEqualConditional(ifA, ifB interfaces.Conditional) bool {
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

func testStmtParse(t *testing.T, stmtAsString string, expectedOutput interface{}) stmt.Form {
	lex := lexer.NewLexer([]byte(stmtAsString))
	parser := parser.NewParser()
	parseResult, err := parser.Parse(lex)

	if err != nil {
		panic(err)
	}

	if f, fOk := parseResult.(stmt.Form); fOk {
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
					if !slicesEqualConditional(fCond[i].(interfaces.Conditional), eCond[i].(interfaces.Conditional)) {
						t.Errorf("parse test failed conditionals not equal: %v %v", fCond, eCond)
					}
				}
			}
		}
	}

	return parseResult.(stmt.Form)
}

/* Tests for statements */

func TestFormIdentifierExtraction(t *testing.T) {
	exampleEmptyForm := "form TestForm {}"

	exampleOutputForm := stmt.Form{vari.VarId{"TestForm"}, stmt.StmtList{}}
	testStmtParse(t, exampleEmptyForm, exampleOutputForm)
}

func TestFormQuestion(t *testing.T) {
	exampleFormInput := "form TestForm { \"Did you sell a house in 2010?\" hasSoldHouse: boolean \"Did you enter a loan?\" hasMaintLoan: boolean }"

	firstQuestionOutput := stmt.InputQuestion{expr.StrLit{"Did you sell a house in 2010?"}, vari.VarDecl{vari.VarId{"hasSoldHouse"}, vari.BoolType{}}}
	secondQuestionOutput := stmt.InputQuestion{expr.StrLit{"Did you enter a loan?"}, vari.VarDecl{vari.VarId{"hasMaintLoan"}, vari.BoolType{}}}
	exampleBodyOutput := stmt.StmtList{[]interfaces.Question{firstQuestionOutput, secondQuestionOutput}, []interfaces.Conditional{}}
	exampleOutputForm := stmt.Form{vari.VarId{"TestForm"}, exampleBodyOutput}

	testStmtParse(t, exampleFormInput, exampleOutputForm)
}

func TestFormComputedQuestion(t *testing.T) {
	exampleFormInput := "form TestForm { \"Did you sell a house in 2010?\" hasSoldHouse: integer \"Did you enter a loan?\" hasMaintLoan: integer \"Value residue:\" valueResidue: integer = (hasSoldHouse - hasMaintLoan) }"

	firstQuestionOutput := stmt.InputQuestion{expr.StrLit{"Did you sell a house in 2010?"}, vari.VarDecl{vari.VarId{"hasSoldHouse"}, vari.IntType{}}}
	secondQuestionOutput := stmt.InputQuestion{expr.StrLit{"Did you enter a loan?"}, vari.VarDecl{vari.VarId{"hasMaintLoan"}, vari.IntType{}}}
	computedQuestion := stmt.ComputedQuestion{expr.StrLit{"Value residue:"}, vari.VarDecl{vari.VarId{"valueResidue"}, vari.IntType{}}, expr.NewSub(expr.VarExpr{vari.VarId{"hasSoldHouse"}}, expr.VarExpr{vari.VarId{"hasMaintLoan"}})}
	exampleBodyOutput := stmt.StmtList{[]interfaces.Question{firstQuestionOutput, secondQuestionOutput, computedQuestion}, []interfaces.Conditional{}}
	exampleOutputForm := stmt.Form{vari.VarId{"TestForm"}, exampleBodyOutput}

	testStmtParse(t, exampleFormInput, exampleOutputForm)
}

func TestFormIf(t *testing.T) {
	exampleFormInput := "form TestForm { \"Did you sell a house in 2010?\" hasSoldHouse: boolean if (true) { \"What was the selling price?\" sellingPrice: integer } }"

	firstQuestionOutput := stmt.InputQuestion{expr.StrLit{"Did you sell a house in 2010?"}, vari.VarDecl{vari.VarId{"hasSoldHouse"}, vari.BoolType{}}}
	firstQuestionBodyInput := stmt.InputQuestion{expr.StrLit{"What was the selling price?"}, vari.VarDecl{vari.VarId{"sellingPrice"}, vari.IntType{}}}
	ifBodyOutput := stmt.StmtList{[]interfaces.Question{firstQuestionBodyInput}, []interfaces.Conditional{}}
	ifOutput := stmt.If{expr.BoolLit{true}, ifBodyOutput}
	exampleBodyOutput := stmt.StmtList{[]interfaces.Question{firstQuestionOutput}, []interfaces.Conditional{ifOutput}}
	exampleOutputForm := stmt.Form{vari.VarId{"TestForm"}, exampleBodyOutput}

	testStmtParse(t, exampleFormInput, exampleOutputForm)
}

func TestFormIfElse(t *testing.T) {
	exampleFormInput := "form TestForm { \"Did you sell a house in 2010?\" hasSoldHouse: boolean if (true) { \"What was the selling price?\" sellingPrice: integer } else { \"What was the selling price?\" sellingPrice: integer } }"

	firstQuestionOutput := stmt.InputQuestion{expr.StrLit{"Did you sell a house in 2010?"}, vari.VarDecl{vari.VarId{"hasSoldHouse"}, vari.BoolType{}}}
	firstQuestionBodyInput := stmt.InputQuestion{expr.StrLit{"What was the selling price?"}, vari.VarDecl{vari.VarId{"sellingPrice"}, vari.IntType{}}}
	ifBodyOutput := stmt.StmtList{[]interfaces.Question{firstQuestionBodyInput}, []interfaces.Conditional{}}
	elseBodyOutput := stmt.StmtList{[]interfaces.Question{firstQuestionBodyInput}, []interfaces.Conditional{}}
	ifOutput := stmt.IfElse{expr.BoolLit{true}, ifBodyOutput, elseBodyOutput}
	exampleBodyOutput := stmt.StmtList{[]interfaces.Question{firstQuestionOutput}, []interfaces.Conditional{ifOutput}}
	exampleOutputForm := stmt.Form{vari.VarId{"TestForm"}, exampleBodyOutput}

	testStmtParse(t, exampleFormInput, exampleOutputForm)
}
