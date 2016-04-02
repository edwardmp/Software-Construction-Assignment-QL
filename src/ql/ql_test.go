package main

import (
	"ql/ast/expr"
	"ql/ast/stmt"
	"ql/ast/vari"
	"ql/interfaces"
	"ql/lexer"
	"ql/parser"
	"testing"
)

func testStmtParse(t *testing.T, stmtAsString string, expectedOutput interfaces.Form) stmt.Form {
	lex := lexer.NewLexer([]byte(stmtAsString))
	parser := parser.NewParser()
	parseResult, err := parser.Parse(lex)

	if err != nil {
		t.Fatalf("Encountered fatal error during parse: %s", err)
	}

	if f, fOk := parseResult.(stmt.Form); fOk {
		if e, eOk := expectedOutput.(stmt.Form); eOk {
			if firstFormIdentifier, secondFormIdentifier := f.Identifier(), e.Identifier(); firstFormIdentifier != secondFormIdentifier {
				t.Errorf("Form identifiers not equal: %s and %s", firstFormIdentifier, secondFormIdentifier)
			}

			if !stmt.SlicesEqual(f.Content(), e.Content()) {
				t.Errorf("Form content not equal: %v and %v", f, e)
			}
		}
	}

	return parseResult.(stmt.Form)
}

/* Tests for statements */

func TestFormIdentifierExtraction(t *testing.T) {
	exampleEmptyForm := "form TestForm {}"

	exampleOutputForm := stmt.NewForm(vari.NewVarId("TestForm"), stmt.NewEmptyStmtList())
	testStmtParse(t, exampleEmptyForm, exampleOutputForm)
}

func TestFormQuestion(t *testing.T) {
	exampleFormInput := "form TestForm { \"Did you sell a house in 2010?\" hasSoldHouse: boolean \"Did you enter a loan?\" hasMaintLoan: boolean }"

	firstQuestionOutput := stmt.NewInputQuestion(expr.NewStrLit("Did you sell a house in 2010?"), vari.NewVarDecl(vari.NewVarId("hasSoldHouse"), expr.NewBoolType()))
	secondQuestionOutput := stmt.NewInputQuestion(expr.NewStrLit("Did you enter a loan?"), vari.NewVarDecl(vari.NewVarId("hasMaintLoan"), expr.NewBoolType()))
	exampleBodyOutput := stmt.NewStmtList([]interfaces.Question{firstQuestionOutput, secondQuestionOutput}, []interfaces.Conditional{})
	exampleOutputForm := stmt.NewForm(vari.NewVarId("TestForm"), exampleBodyOutput)

	testStmtParse(t, exampleFormInput, exampleOutputForm)
}

func TestFormComputedQuestion(t *testing.T) {
	exampleFormInput := "form TestForm { \"Did you sell a house in 2010?\" hasSoldHouse: integer \"Did you enter a loan?\" hasMaintLoan: integer \"Value residue:\" valueResidue: integer = (hasSoldHouse - hasMaintLoan) }"

	firstQuestionOutput := stmt.NewInputQuestion(expr.NewStrLit("Did you sell a house in 2010?"), vari.NewVarDecl(vari.NewVarId("hasSoldHouse"), expr.NewIntType()))
	secondQuestionOutput := stmt.NewInputQuestion(expr.NewStrLit("Did you enter a loan?"), vari.NewVarDecl(vari.NewVarId("hasMaintLoan"), expr.NewIntType()))
	computedQuestion := stmt.NewComputedQuestion(expr.NewStrLit("Value residue:"), vari.NewVarDecl(vari.NewVarId("valueResidue"), expr.NewIntType()), expr.NewSub(expr.NewVarExpr(vari.NewVarId("hasSoldHouse")), expr.NewVarExpr(vari.NewVarId("hasMaintLoan"))))
	exampleBodyOutput := stmt.NewStmtList([]interfaces.Question{firstQuestionOutput, secondQuestionOutput, computedQuestion}, []interfaces.Conditional{})
	exampleOutputForm := stmt.NewForm(vari.NewVarId("TestForm"), exampleBodyOutput)

	testStmtParse(t, exampleFormInput, exampleOutputForm)
}

func TestFormIf(t *testing.T) {
	exampleFormInput := "form TestForm { \"Did you sell a house in 2010?\" hasSoldHouse: boolean if (true) { \"What was the selling price?\" sellingPrice: integer } }"

	firstQuestionOutput := stmt.NewInputQuestion(expr.NewStrLit("Did you sell a house in 2010?"), vari.NewVarDecl(vari.NewVarId("hasSoldHouse"), expr.NewBoolType()))
	firstQuestionBodyInput := stmt.NewInputQuestion(expr.NewStrLit("What was the selling price?"), vari.NewVarDecl(vari.NewVarId("sellingPrice"), expr.NewIntType()))
	ifBodyOutput := stmt.NewStmtList([]interfaces.Question{firstQuestionBodyInput}, []interfaces.Conditional{})
	ifOutput := stmt.NewIf(expr.NewBoolLit(true), ifBodyOutput)
	exampleBodyOutput := stmt.NewStmtList([]interfaces.Question{firstQuestionOutput}, []interfaces.Conditional{ifOutput})
	exampleOutputForm := stmt.NewForm(vari.NewVarId("TestForm"), exampleBodyOutput)

	testStmtParse(t, exampleFormInput, exampleOutputForm)
}

func TestFormIfElse(t *testing.T) {
	exampleFormInput := "form TestForm { \"Did you sell a house in 2010?\" hasSoldHouse: boolean if (true) { \"What was the selling price?\" sellingPrice: integer } else { \"What was the selling price?\" sellingPrice: integer } }"

	firstQuestionOutput := stmt.NewInputQuestion(expr.NewStrLit("Did you sell a house in 2010?"), vari.NewVarDecl(vari.NewVarId("hasSoldHouse"), expr.NewBoolType()))
	firstQuestionBodyInput := stmt.NewInputQuestion(expr.NewStrLit("What was the selling price?"), vari.NewVarDecl(vari.NewVarId("sellingPrice"), expr.NewIntType()))
	ifBodyOutput := stmt.NewStmtList([]interfaces.Question{firstQuestionBodyInput}, []interfaces.Conditional{})
	elseBodyOutput := stmt.NewStmtList([]interfaces.Question{firstQuestionBodyInput}, []interfaces.Conditional{})
	ifOutput := stmt.NewIfElse(expr.NewBoolLit(true), ifBodyOutput, elseBodyOutput)
	exampleBodyOutput := stmt.NewStmtList([]interfaces.Question{firstQuestionOutput}, []interfaces.Conditional{ifOutput})
	exampleOutputForm := stmt.NewForm(vari.NewVarId("TestForm"), exampleBodyOutput)

	testStmtParse(t, exampleFormInput, exampleOutputForm)
}
