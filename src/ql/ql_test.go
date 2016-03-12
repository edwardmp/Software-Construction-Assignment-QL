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

	for i, questionA := range questionsA {
		if questionA != questionsB[i] {
			fmt.Print(questionA)
			fmt.Print(questionsB[i])

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
			return false
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
			if f.GetIdentifierAsString() != e.GetIdentifierAsString() {
				t.Errorf("Form identifiers not equal %v %v", f.Identifier, e.Identifier)
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

	exampleOutputForm := stmt.NewFormNoSourceInfo(vari.NewVarIdNoSourceInfo("TestForm"), stmt.NewEmptyStmtListNoSourceInfo())
	testStmtParse(t, exampleEmptyForm, exampleOutputForm)
}

func TestFormQuestion(t *testing.T) {
	exampleFormInput := "form TestForm { \"Did you sell a house in 2010?\" hasSoldHouse: boolean \"Did you enter a loan?\" hasMaintLoan: boolean }"

	firstQuestionOutput := stmt.NewInputQuestionNoSourceInfo(expr.NewStrLitNoSourceInfo("Did you sell a house in 2010?"), vari.NewVarDeclNoSourceInfo(vari.NewVarIdNoSourceInfo("hasSoldHouse"), vari.NewBoolTypeNoSourceInfo()))
	secondQuestionOutput := stmt.NewInputQuestionNoSourceInfo(expr.NewStrLitNoSourceInfo("Did you enter a loan?"), vari.NewVarDeclNoSourceInfo(vari.NewVarIdNoSourceInfo("hasMaintLoan"), vari.NewBoolTypeNoSourceInfo()))
	exampleBodyOutput := stmt.NewStmtListNoSourceInfo([]interfaces.Question{firstQuestionOutput, secondQuestionOutput}, []interfaces.Conditional{})
	exampleOutputForm := stmt.NewFormNoSourceInfo(vari.NewVarIdNoSourceInfo("TestForm"), exampleBodyOutput)

	testStmtParse(t, exampleFormInput, exampleOutputForm)
}

func TestFormComputedQuestion(t *testing.T) {
	exampleFormInput := "form TestForm { \"Did you sell a house in 2010?\" hasSoldHouse: integer \"Did you enter a loan?\" hasMaintLoan: integer \"Value residue:\" valueResidue: integer = (hasSoldHouse - hasMaintLoan) }"

	firstQuestionOutput := stmt.NewInputQuestionNoSourceInfo(expr.NewStrLitNoSourceInfo("Did you sell a house in 2010?"), vari.NewVarDeclNoSourceInfo(vari.NewVarIdNoSourceInfo("hasSoldHouse"), vari.NewIntTypeNoSourceInfo()))
	secondQuestionOutput := stmt.NewInputQuestionNoSourceInfo(expr.NewStrLitNoSourceInfo("Did you enter a loan?"), vari.NewVarDeclNoSourceInfo(vari.NewVarIdNoSourceInfo("hasMaintLoan"), vari.NewIntTypeNoSourceInfo()))
	computedQuestion := stmt.NewComputedQuestionNoSourceInfo(expr.NewStrLitNoSourceInfo("Value residue:"), vari.NewVarDeclNoSourceInfo(vari.NewVarIdNoSourceInfo("valueResidue"), vari.NewIntTypeNoSourceInfo()), expr.NewSubNoSourceInfo(expr.NewVarExprNoSourceInfo(vari.NewVarIdNoSourceInfo("hasSoldHouse")), expr.NewVarExprNoSourceInfo(vari.NewVarIdNoSourceInfo("hasMaintLoan"))))
	exampleBodyOutput := stmt.NewStmtListNoSourceInfo([]interfaces.Question{firstQuestionOutput, secondQuestionOutput, computedQuestion}, []interfaces.Conditional{})
	exampleOutputForm := stmt.NewFormNoSourceInfo(vari.NewVarIdNoSourceInfo("TestForm"), exampleBodyOutput)

	testStmtParse(t, exampleFormInput, exampleOutputForm)
}

func TestFormIf(t *testing.T) {
	exampleFormInput := "form TestForm { \"Did you sell a house in 2010?\" hasSoldHouse: boolean if (true) { \"What was the selling price?\" sellingPrice: integer } }"

	firstQuestionOutput := stmt.NewInputQuestionNoSourceInfo(expr.NewStrLitNoSourceInfo("Did you sell a house in 2010?"), vari.NewVarDeclNoSourceInfo(vari.NewVarIdNoSourceInfo("hasSoldHouse"), vari.NewBoolTypeNoSourceInfo()))
	firstQuestionBodyInput := stmt.NewInputQuestionNoSourceInfo(expr.NewStrLitNoSourceInfo("What was the selling price?"), vari.NewVarDeclNoSourceInfo(vari.NewVarIdNoSourceInfo("sellingPrice"), vari.NewIntTypeNoSourceInfo()))
	ifBodyOutput := stmt.NewStmtListNoSourceInfo([]interfaces.Question{firstQuestionBodyInput}, []interfaces.Conditional{})
	ifOutput := stmt.NewIfNoSourceInfo(expr.NewBoolLitNoSourceInfo(true), ifBodyOutput)
	exampleBodyOutput := stmt.NewStmtListNoSourceInfo([]interfaces.Question{firstQuestionOutput}, []interfaces.Conditional{ifOutput})
	exampleOutputForm := stmt.NewFormNoSourceInfo(vari.NewVarIdNoSourceInfo("TestForm"), exampleBodyOutput)

	testStmtParse(t, exampleFormInput, exampleOutputForm)
}

func TestFormIfElse(t *testing.T) {
	exampleFormInput := "form TestForm { \"Did you sell a house in 2010?\" hasSoldHouse: boolean if (true) { \"What was the selling price?\" sellingPrice: integer } else { \"What was the selling price?\" sellingPrice: integer } }"

	firstQuestionOutput := stmt.NewInputQuestionNoSourceInfo(expr.NewStrLitNoSourceInfo("Did you sell a house in 2010?"), vari.NewVarDeclNoSourceInfo(vari.NewVarIdNoSourceInfo("hasSoldHouse"), vari.NewBoolTypeNoSourceInfo()))
	firstQuestionBodyInput := stmt.NewInputQuestionNoSourceInfo(expr.NewStrLitNoSourceInfo("What was the selling price?"), vari.NewVarDeclNoSourceInfo(vari.NewVarIdNoSourceInfo("sellingPrice"), vari.NewIntTypeNoSourceInfo()))
	ifBodyOutput := stmt.NewStmtListNoSourceInfo([]interfaces.Question{firstQuestionBodyInput}, []interfaces.Conditional{})
	elseBodyOutput := stmt.NewStmtListNoSourceInfo([]interfaces.Question{firstQuestionBodyInput}, []interfaces.Conditional{})
	ifOutput := stmt.NewIfElseNoSourceInfo(expr.NewBoolLitNoSourceInfo(true), ifBodyOutput, elseBodyOutput)
	exampleBodyOutput := stmt.NewStmtListNoSourceInfo([]interfaces.Question{firstQuestionOutput}, []interfaces.Conditional{ifOutput})
	exampleOutputForm := stmt.NewFormNoSourceInfo(vari.NewVarIdNoSourceInfo("TestForm"), exampleBodyOutput)

	testStmtParse(t, exampleFormInput, exampleOutputForm)
}
