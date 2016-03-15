package stmt

import (
	"ql/ast/expr"
	"ql/ast/vari"
	"ql/interfaces"
	"testing"
)

/* Tests for statements */

func TestFormWithEmptyContent(t *testing.T) {
	identifier := vari.NewVarIdNoSourceInfo("TestForm")
	exampleForm := NewFormNoSourceInfo(identifier, NewEmptyStmtListNoSourceInfo())

	if exampleForm.Identifier != identifier {
		t.Errorf("Form identifier is not set correctly")
	}

	if len(exampleForm.Content.Questions) != 0 {
		t.Errorf("Form content questions are not set correctly")
	}

	if len(exampleForm.Content.Conditionals) != 0 {
		t.Errorf("Form content conditionals are not set correctly")
	}
}

func TestFormWithNonEmptyContent(t *testing.T) {
	identifier := vari.NewVarIdNoSourceInfo("TestForm")
	questionExample := NewInputQuestionNoSourceInfo(expr.NewStrLitNoSourceInfo("What was the selling price?"), vari.NewVarDeclNoSourceInfo(vari.NewVarIdNoSourceInfo("sellingPrice"), vari.NewIntTypeNoSourceInfo()))
	questionsListExample := []interfaces.Question{questionExample}
	stmtListExample := NewStmtListNoSourceInfo(questionsListExample, []interfaces.Conditional{})
	exampleForm := NewFormNoSourceInfo(identifier, stmtListExample)

	if len(exampleForm.Content.Questions) != 1 {
		t.Errorf("Form content questions does not have 1 question while it should")
	}

	if !SlicesEqual(exampleForm.Content, stmtListExample) {
		t.Errorf("Form content not set correctly")
	}
}

func TestInputQuestion(t *testing.T) {
	exampleLabel := expr.NewStrLitNoSourceInfo("Did you sell a house in 2010?")
	exampleVarDecl := vari.NewVarDeclNoSourceInfo(vari.NewVarIdNoSourceInfo("hasSoldHouse"), vari.NewBoolTypeNoSourceInfo())

	exampleQuestion := NewInputQuestionNoSourceInfo(exampleLabel, exampleVarDecl)

	if exampleQuestion.Label != exampleLabel {
		t.Errorf("Question label is not set correctly")
	}
}

func TestComputedQuestion(t *testing.T) {
	exampleLabel := expr.NewStrLitNoSourceInfo("Value residue")
	exampleVarDecl := vari.NewVarDeclNoSourceInfo(vari.NewVarIdNoSourceInfo("hasSoldHouse"), vari.NewIntTypeNoSourceInfo())
	exampleComputation := expr.NewSubNoSourceInfo(expr.NewIntLitNoSourceInfo(10), expr.NewIntLitNoSourceInfo(5))

	exampleQuestion := NewComputedQuestionNoSourceInfo(exampleLabel, exampleVarDecl, exampleComputation)

	if exampleQuestion.Label != exampleLabel {
		t.Errorf("Computed question label is not set correctly")
	}

	if exampleQuestion.Computation != exampleComputation {
		t.Errorf("Computed question computation is not set correctly")
	}
}

func TestIf(t *testing.T) {
	questionExample := NewInputQuestionNoSourceInfo(expr.NewStrLitNoSourceInfo("What was the selling price?"), vari.NewVarDeclNoSourceInfo(vari.NewVarIdNoSourceInfo("sellingPrice"), vari.NewIntTypeNoSourceInfo()))
	ifBodyExample := NewStmtListNoSourceInfo([]interfaces.Question{questionExample}, []interfaces.Conditional{})
	ifCondExample := expr.NewBoolLitNoSourceInfo(true)
	ifExample := NewIfNoSourceInfo(ifCondExample, ifBodyExample)

	if !SlicesEqual(ifExample.Body, ifBodyExample) {
		t.Errorf("If body is not set correctly")
	}

	if ifExample.Cond != ifCondExample {
		t.Errorf("If cond is not set correctly")
	}
}

func TestIfElse(t *testing.T) {
	ifQuestionExample := NewInputQuestionNoSourceInfo(expr.NewStrLitNoSourceInfo("Did you sell a house in 2010?"), vari.NewVarDeclNoSourceInfo(vari.NewVarIdNoSourceInfo("hasSoldHouse"), vari.NewBoolTypeNoSourceInfo()))
	ifBodyExample := NewStmtListNoSourceInfo([]interfaces.Question{ifQuestionExample}, []interfaces.Conditional{})
	ifCondExample := expr.NewBoolLitNoSourceInfo(true)

	elseQuestionExample := NewInputQuestionNoSourceInfo(expr.NewStrLitNoSourceInfo("What was the selling price?"), vari.NewVarDeclNoSourceInfo(vari.NewVarIdNoSourceInfo("sellingPrice"), vari.NewIntTypeNoSourceInfo()))
	elseBodyExample := NewStmtListNoSourceInfo([]interfaces.Question{elseQuestionExample}, []interfaces.Conditional{})

	ifElseExample := NewIfElseNoSourceInfo(ifCondExample, ifBodyExample, elseBodyExample)

	if !SlicesEqual(ifElseExample.IfBody, ifBodyExample) {
		t.Errorf("IfElse else body is not set correctly")
	}

	if !SlicesEqual(ifElseExample.ElseBody, elseBodyExample) {
		t.Errorf("IfElse if body is not set correctly")
	}

	if ifElseExample.Cond != ifCondExample {
		t.Errorf("If cond is not set correctly")
	}
}

func TestStmtList(t *testing.T) {
	questionExample := NewInputQuestionNoSourceInfo(expr.NewStrLitNoSourceInfo("Did you sell a house in 2010?"), vari.NewVarDeclNoSourceInfo(vari.NewVarIdNoSourceInfo("hasSoldHouse"), vari.NewBoolTypeNoSourceInfo()))
	questionListExample := []interfaces.Question{questionExample}

	ifExample := NewIfNoSourceInfo(expr.NewBoolLitNoSourceInfo(true), NewEmptyStmtListNoSourceInfo())
	conditionalListExample := []interfaces.Conditional{ifExample}

	stmtListExample := NewStmtListNoSourceInfo(questionListExample, conditionalListExample)

	if len(stmtListExample.Questions) != len(questionListExample) {
		t.Errorf("Stmtlist questions list is not set correctly")
	}

	if len(stmtListExample.Conditionals) != len(conditionalListExample) {
		t.Errorf("Stmtlist conditionals list is not set correctly")
	}
}
