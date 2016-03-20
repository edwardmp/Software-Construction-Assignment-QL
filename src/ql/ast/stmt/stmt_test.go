package stmt

import (
	"ql/ast/expr"
	"ql/ast/vari"
	"ql/interfaces"
	"testing"
)

/* Tests for statements */

func TestFormWithEmptyContent(t *testing.T) {
	identifier := vari.NewVarId("TestForm")
	exampleForm := NewForm(identifier, NewEmptyStmtList())

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
	identifier := vari.NewVarId("TestForm")
	questionExample := NewInputQuestion(expr.NewStrLit("What was the selling price?"), vari.NewVarDecl(vari.NewVarId("sellingPrice"), expr.NewIntType()))
	questionsListExample := []interfaces.Question{questionExample}
	stmtListExample := NewStmtList(questionsListExample, []interfaces.Conditional{})
	exampleForm := NewForm(identifier, stmtListExample)

	if len(exampleForm.Content.Questions) != 1 {
		t.Errorf("Form content questions does not have 1 question while it should")
	}

	if !SlicesEqual(exampleForm.Content, stmtListExample) {
		t.Errorf("Form content not set correctly")
	}
}

func TestInputQuestion(t *testing.T) {
	exampleLabel := expr.NewStrLit("Did you sell a house in 2010?")
	exampleVarDecl := vari.NewVarDecl(vari.NewVarId("hasSoldHouse"), expr.NewBoolType())

	exampleQuestion := NewInputQuestion(exampleLabel, exampleVarDecl)

	if exampleQuestion.Label != exampleLabel {
		t.Errorf("Question label is not set correctly")
	}
}

func TestComputedQuestion(t *testing.T) {
	exampleLabel := expr.NewStrLit("Value residue")
	exampleVarDecl := vari.NewVarDecl(vari.NewVarId("hasSoldHouse"), expr.NewIntType())
	exampleComputation := expr.NewSub(expr.NewIntLit(10), expr.NewIntLit(5))

	exampleQuestion := NewComputedQuestion(exampleLabel, exampleVarDecl, exampleComputation)

	if exampleQuestion.Label != exampleLabel {
		t.Errorf("Computed question label is not set correctly")
	}

	if exampleQuestion.Computation != exampleComputation {
		t.Errorf("Computed question computation is not set correctly")
	}
}

func TestIf(t *testing.T) {
	questionExample := NewInputQuestion(expr.NewStrLit("What was the selling price?"), vari.NewVarDecl(vari.NewVarId("sellingPrice"), expr.NewIntType()))
	ifBodyExample := NewStmtList([]interfaces.Question{questionExample}, []interfaces.Conditional{})
	ifCondExample := expr.NewBoolLit(true)
	ifExample := NewIf(ifCondExample, ifBodyExample)

	if !SlicesEqual(ifExample.Body, ifBodyExample) {
		t.Errorf("If body is not set correctly")
	}

	if ifExample.Cond != ifCondExample {
		t.Errorf("If cond is not set correctly")
	}
}

func TestIfElse(t *testing.T) {
	ifQuestionExample := NewInputQuestion(expr.NewStrLit("Did you sell a house in 2010?"), vari.NewVarDecl(vari.NewVarId("hasSoldHouse"), expr.NewBoolType()))
	ifBodyExample := NewStmtList([]interfaces.Question{ifQuestionExample}, []interfaces.Conditional{})
	ifCondExample := expr.NewBoolLit(true)

	elseQuestionExample := NewInputQuestion(expr.NewStrLit("What was the selling price?"), vari.NewVarDecl(vari.NewVarId("sellingPrice"), expr.NewIntType()))
	elseBodyExample := NewStmtList([]interfaces.Question{elseQuestionExample}, []interfaces.Conditional{})

	ifElseExample := NewIfElse(ifCondExample, ifBodyExample, elseBodyExample)

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
	questionExample := NewInputQuestion(expr.NewStrLit("Did you sell a house in 2010?"), vari.NewVarDecl(vari.NewVarId("hasSoldHouse"), expr.NewBoolType()))
	questionListExample := []interfaces.Question{questionExample}

	ifExample := NewIf(expr.NewBoolLit(true), NewEmptyStmtList())
	conditionalListExample := []interfaces.Conditional{ifExample}

	stmtListExample := NewStmtList(questionListExample, conditionalListExample)

	if len(stmtListExample.Questions) != len(questionListExample) {
		t.Errorf("Stmtlist questions list is not set correctly")
	}

	if len(stmtListExample.Conditionals) != len(conditionalListExample) {
		t.Errorf("Stmtlist conditionals list is not set correctly")
	}
}
