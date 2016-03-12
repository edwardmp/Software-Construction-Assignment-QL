package stmt

import (
	"fmt"
	"ql/ast/expr"
	"ql/ast/typechecker"
	"ql/ast/vari"
	"ql/interfaces"
	"ql/symboltable"
	"testing"
)

// TODO PUT IN HELPER FILE
// slices don't support equality checking, so have to do it manually
func slicesEqual(a StmtList, b StmtList) bool {
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
	case If:
		bodyA := ifA.(If).Body
		bodyB := ifB.(If).Body
		return slicesEqual(bodyA, bodyB)
	case IfElse:
		return slicesEqual(ifA.(IfElse).IfBody, ifB.(IfElse).IfBody) && slicesEqual(ifA.(IfElse).ElseBody, ifB.(IfElse).ElseBody)
	}
}

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

	if !slicesEqual(exampleForm.Content, stmtListExample) {
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

	if !slicesEqual(ifExample.Body, ifBodyExample) {
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

	if !slicesEqual(ifElseExample.IfBody, ifBodyExample) {
		t.Errorf("IfElse else body is not set correctly")
	}

	if !slicesEqual(ifElseExample.ElseBody, elseBodyExample) {
		t.Errorf("IfElse if body is not set correctly")
	}

	if ifElseExample.Cond != ifCondExample {
		t.Errorf("If cond is not set correctly")
	}
}

func TestStmtList(t *testing.T) {
	questionExample := NewInputQuestionNoSourceInfo(expr.NewStrLitNoSourceInfo("Did you sell a house in 2010?"), vari.NewVarDeclNoSourceInfo(vari.NewVarIdNoSourceInfo("hasSoldHouse"), vari.NewBoolTypeNoSourceInfo()))
	questionListExample := []interfaces.Question{questionExample}

	ifExample := NewIfNoSourceInfo(expr.NewBoolLitNoSourceInfo(true), StmtList{})
	conditionalListExample := []interfaces.Conditional{ifExample}

	stmtListExample := NewStmtListNoSourceInfo(questionListExample, conditionalListExample)

	if len(stmtListExample.Questions) != len(questionListExample) {
		t.Errorf("Stmtlist questions list is not set correctly")
	}

	if len(stmtListExample.Conditionals) != len(conditionalListExample) {
		t.Errorf("Stmtlist conditionals list is not set correctly")
	}
}

func TestNonBoolConditionalChecker(t *testing.T) {
	exampleQuestion := NewInputQuestionNoSourceInfo(expr.NewStrLitNoSourceInfo("Did you sell a house in 2010?"), vari.NewVarDeclNoSourceInfo(vari.NewVarIdNoSourceInfo("hasSoldHouse"), vari.NewBoolTypeNoSourceInfo()))
	exampleIf := NewIfNoSourceInfo(expr.NewIntLitNoSourceInfo(10), NewStmtListNoSourceInfo([]interfaces.Question{exampleQuestion}, []interfaces.Conditional{}))
	exampleBody := NewStmtListNoSourceInfo([]interfaces.Question{}, []interfaces.Conditional{exampleIf})
	exampleForm := NewFormNoSourceInfo(vari.NewVarIdNoSourceInfo("TestForm"), exampleBody)

	typeChecker := typechecker.NewTypeChecker()
	exampleForm.TypeCheck(&typeChecker, symboltable.NewSymbolTable())
	errorsReported := typeChecker.GetEncountedErrorsForCheckType("NonBoolConditionals")

	if len(errorsReported) != 1 || fmt.Sprintf("%v", errorsReported[0]) != fmt.Sprintf("%v", fmt.Errorf("Non-boolean type used as condition: int")) {
		t.Errorf("Non bool condition type checker did not correctly report condition of invalid type %v", errorsReported)
	}
}

func TestDuplicateLabelChecker(t *testing.T) {
	firstQuestion := NewInputQuestionNoSourceInfo(expr.NewStrLitNoSourceInfo("Did you sell a house in 2010?"), vari.NewVarDeclNoSourceInfo(vari.NewVarIdNoSourceInfo("hasSoldHouse"), vari.NewBoolTypeNoSourceInfo()))
	secondQuestion := NewInputQuestionNoSourceInfo(expr.NewStrLitNoSourceInfo("Did you sell a house in 2010?"), vari.NewVarDeclNoSourceInfo(vari.NewVarIdNoSourceInfo("hasMaintLoan"), vari.NewBoolTypeNoSourceInfo()))
	exampleBody := NewStmtListNoSourceInfo([]interfaces.Question{firstQuestion, secondQuestion}, []interfaces.Conditional{})
	exampleForm := NewFormNoSourceInfo(vari.NewVarIdNoSourceInfo("TestForm"), exampleBody)

	typeChecker := typechecker.NewTypeChecker()
	exampleForm.TypeCheck(&typeChecker, symboltable.NewSymbolTable())
	warningsReported := typeChecker.GetEncountedErrorsForCheckType("DuplicateLabels")

	if len(warningsReported) != 1 || fmt.Sprintf("%v", warningsReported[0]) != fmt.Sprintf("%v", fmt.Errorf("Label \"Did you sell a house in 2010?\" already used for question with identifier hasSoldHouse, using again for question with identifier hasMaintLoan")) {
		t.Errorf("Duplicate label not reported correctly by type checker")
	}
}

func TestDuplicateVarDeclChecker(t *testing.T) {
	firstQuestion := NewInputQuestionNoSourceInfo(expr.NewStrLitNoSourceInfo("Did you sell a house in 2010?"), vari.NewVarDeclNoSourceInfo(vari.NewVarIdNoSourceInfo("hasSoldHouse"), vari.NewBoolTypeNoSourceInfo()))
	secondQuestion := NewInputQuestionNoSourceInfo(expr.NewStrLitNoSourceInfo("Did you sell a house in 2010?"), vari.NewVarDeclNoSourceInfo(vari.NewVarIdNoSourceInfo("hasSoldHouse"), vari.NewIntTypeNoSourceInfo()))
	exampleBody := NewStmtListNoSourceInfo([]interfaces.Question{firstQuestion, secondQuestion}, []interfaces.Conditional{})
	exampleForm := NewFormNoSourceInfo(vari.NewVarIdNoSourceInfo("TestForm"), exampleBody)

	typeChecker := typechecker.NewTypeChecker()
	exampleForm.TypeCheck(&typeChecker, symboltable.NewSymbolTable())
	errorsReported := typeChecker.GetEncountedErrorsForCheckType("DuplicateVarDeclarations")

	if len(errorsReported) != 1 || fmt.Sprintf("%v", errorsReported[0]) != fmt.Sprintf("%v", fmt.Errorf("Question redeclared with different types: vari.IntType and vari.BoolType")) {
		t.Errorf("Duplicate var decl not reported correctly by type checker")
	}
}
