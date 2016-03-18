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
			if firstFormIdentifier, secondFormIdentifier := f.GetIdentifier(), e.GetIdentifier(); firstFormIdentifier != secondFormIdentifier {
				t.Errorf("Form identifiers not equal: %s and %s", firstFormIdentifier, secondFormIdentifier)
			}

			if !stmt.SlicesEqual(f.Content, e.Content) {
				t.Errorf("Form content not equal: %v and %v", f, e)
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

	firstQuestionOutput := stmt.NewInputQuestionNoSourceInfo(expr.NewStrLitNoSourceInfo("Did you sell a house in 2010?"), vari.NewVarDeclNoSourceInfo(vari.NewVarIdNoSourceInfo("hasSoldHouse"), expr.NewBoolTypeNoSourceInfo()))
	secondQuestionOutput := stmt.NewInputQuestionNoSourceInfo(expr.NewStrLitNoSourceInfo("Did you enter a loan?"), vari.NewVarDeclNoSourceInfo(vari.NewVarIdNoSourceInfo("hasMaintLoan"), expr.NewBoolTypeNoSourceInfo()))
	exampleBodyOutput := stmt.NewStmtListNoSourceInfo([]interfaces.Question{firstQuestionOutput, secondQuestionOutput}, []interfaces.Conditional{})
	exampleOutputForm := stmt.NewFormNoSourceInfo(vari.NewVarIdNoSourceInfo("TestForm"), exampleBodyOutput)

	testStmtParse(t, exampleFormInput, exampleOutputForm)
}

func TestFormComputedQuestion(t *testing.T) {
	exampleFormInput := "form TestForm { \"Did you sell a house in 2010?\" hasSoldHouse: integer \"Did you enter a loan?\" hasMaintLoan: integer \"Value residue:\" valueResidue: integer = (hasSoldHouse - hasMaintLoan) }"

	firstQuestionOutput := stmt.NewInputQuestionNoSourceInfo(expr.NewStrLitNoSourceInfo("Did you sell a house in 2010?"), vari.NewVarDeclNoSourceInfo(vari.NewVarIdNoSourceInfo("hasSoldHouse"), expr.NewIntTypeNoSourceInfo()))
	secondQuestionOutput := stmt.NewInputQuestionNoSourceInfo(expr.NewStrLitNoSourceInfo("Did you enter a loan?"), vari.NewVarDeclNoSourceInfo(vari.NewVarIdNoSourceInfo("hasMaintLoan"), expr.NewIntTypeNoSourceInfo()))
	computedQuestion := stmt.NewComputedQuestionNoSourceInfo(expr.NewStrLitNoSourceInfo("Value residue:"), vari.NewVarDeclNoSourceInfo(vari.NewVarIdNoSourceInfo("valueResidue"), expr.NewIntTypeNoSourceInfo()), expr.NewSubNoSourceInfo(expr.NewVarExprNoSourceInfo(vari.NewVarIdNoSourceInfo("hasSoldHouse")), expr.NewVarExprNoSourceInfo(vari.NewVarIdNoSourceInfo("hasMaintLoan"))))
	exampleBodyOutput := stmt.NewStmtListNoSourceInfo([]interfaces.Question{firstQuestionOutput, secondQuestionOutput, computedQuestion}, []interfaces.Conditional{})
	exampleOutputForm := stmt.NewFormNoSourceInfo(vari.NewVarIdNoSourceInfo("TestForm"), exampleBodyOutput)

	testStmtParse(t, exampleFormInput, exampleOutputForm)
}

func TestFormIf(t *testing.T) {
	exampleFormInput := "form TestForm { \"Did you sell a house in 2010?\" hasSoldHouse: boolean if (true) { \"What was the selling price?\" sellingPrice: integer } }"

	firstQuestionOutput := stmt.NewInputQuestionNoSourceInfo(expr.NewStrLitNoSourceInfo("Did you sell a house in 2010?"), vari.NewVarDeclNoSourceInfo(vari.NewVarIdNoSourceInfo("hasSoldHouse"), expr.NewBoolTypeNoSourceInfo()))
	firstQuestionBodyInput := stmt.NewInputQuestionNoSourceInfo(expr.NewStrLitNoSourceInfo("What was the selling price?"), vari.NewVarDeclNoSourceInfo(vari.NewVarIdNoSourceInfo("sellingPrice"), expr.NewIntTypeNoSourceInfo()))
	ifBodyOutput := stmt.NewStmtListNoSourceInfo([]interfaces.Question{firstQuestionBodyInput}, []interfaces.Conditional{})
	ifOutput := stmt.NewIfNoSourceInfo(expr.NewBoolLitNoSourceInfo(true), ifBodyOutput)
	exampleBodyOutput := stmt.NewStmtListNoSourceInfo([]interfaces.Question{firstQuestionOutput}, []interfaces.Conditional{ifOutput})
	exampleOutputForm := stmt.NewFormNoSourceInfo(vari.NewVarIdNoSourceInfo("TestForm"), exampleBodyOutput)

	testStmtParse(t, exampleFormInput, exampleOutputForm)
}

func TestFormIfElse(t *testing.T) {
	exampleFormInput := "form TestForm { \"Did you sell a house in 2010?\" hasSoldHouse: boolean if (true) { \"What was the selling price?\" sellingPrice: integer } else { \"What was the selling price?\" sellingPrice: integer } }"

	firstQuestionOutput := stmt.NewInputQuestionNoSourceInfo(expr.NewStrLitNoSourceInfo("Did you sell a house in 2010?"), vari.NewVarDeclNoSourceInfo(vari.NewVarIdNoSourceInfo("hasSoldHouse"), expr.NewBoolTypeNoSourceInfo()))
	firstQuestionBodyInput := stmt.NewInputQuestionNoSourceInfo(expr.NewStrLitNoSourceInfo("What was the selling price?"), vari.NewVarDeclNoSourceInfo(vari.NewVarIdNoSourceInfo("sellingPrice"), expr.NewIntTypeNoSourceInfo()))
	ifBodyOutput := stmt.NewStmtListNoSourceInfo([]interfaces.Question{firstQuestionBodyInput}, []interfaces.Conditional{})
	elseBodyOutput := stmt.NewStmtListNoSourceInfo([]interfaces.Question{firstQuestionBodyInput}, []interfaces.Conditional{})
	ifOutput := stmt.NewIfElseNoSourceInfo(expr.NewBoolLitNoSourceInfo(true), ifBodyOutput, elseBodyOutput)
	exampleBodyOutput := stmt.NewStmtListNoSourceInfo([]interfaces.Question{firstQuestionOutput}, []interfaces.Conditional{ifOutput})
	exampleOutputForm := stmt.NewFormNoSourceInfo(vari.NewVarIdNoSourceInfo("TestForm"), exampleBodyOutput)

	testStmtParse(t, exampleFormInput, exampleOutputForm)
}
