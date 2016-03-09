package stmt

import (
	"fmt"
	"ql/ast/expr/binaryoperatorexpr"
	"ql/ast/expr/litexpr"
	"ql/ast/vari"
	"ql/ast/vari/vartype"
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
		if conditionalsA[i].(Conditional).EvalCondition() != conditionalsB[i].(Conditional).EvalCondition() {
			return false
		}

		if !slicesEqualConditional(conditionalsA[i].(Conditional), conditionalsB[i].(Conditional)) {
			return true
		}
	}

	return true
}

func slicesEqualConditional(ifA, ifB Conditional) bool {
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
	identifier := vari.VarId{"TestForm"}
	exampleForm := Form{identifier, StmtList{}}

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
	identifier := vari.VarId{"TestForm"}
	questionExample := InputQuestion{litexpr.StrLit{"What was the selling price?"}, vari.VarDecl{vari.VarId{"sellingPrice"}, vartype.IntType{}}}
	questionsListExample := []Question{questionExample}
	stmtListExample := StmtList{Questions: questionsListExample}
	exampleForm := Form{identifier, stmtListExample}

	if len(exampleForm.Content.Questions) != 1 {
		t.Errorf("Form content questions does not have 1 question while it should")
	}

	if !slicesEqual(exampleForm.Content, stmtListExample) {
		t.Errorf("Form content not set correctly")
	}
}

func TestInputQuestion(t *testing.T) {
	exampleLabel := litexpr.StrLit{"Did you sell a house in 2010?"}
	exampleVarDecl := vari.VarDecl{vari.VarId{"hasSoldHouse"}, vartype.BoolType{}}

	exampleQuestion := InputQuestion{exampleLabel, exampleVarDecl}

	if exampleQuestion.Label != exampleLabel {
		t.Errorf("Question label is not set correctly")
	}
}

func TestComputedQuestion(t *testing.T) {
	exampleLabel := litexpr.StrLit{"Value residue"}
	exampleVarDecl := vari.VarDecl{vari.VarId{"hasSoldHouse"}, vartype.IntType{}}
	exampleComputation := binaryoperatorexpr.Sub{litexpr.IntLit{10}, litexpr.IntLit{5}}

	exampleQuestion := ComputedQuestion{exampleLabel, exampleVarDecl, exampleComputation}

	if exampleQuestion.Label != exampleLabel {
		t.Errorf("Computed question label is not set correctly")
	}

	if exampleQuestion.Computation != exampleComputation {
		t.Errorf("Computed question computation is not set correctly")
	}
}

func TestIf(t *testing.T) {
	questionExample := InputQuestion{litexpr.StrLit{"What was the selling price?"}, vari.VarDecl{vari.VarId{"sellingPrice"}, vartype.IntType{}}}
	ifBodyExample := StmtList{[]Question{questionExample}, []Conditional{}}
	ifCondExample := litexpr.BoolLit{true}
	ifExample := If{ifCondExample, ifBodyExample}

	if !slicesEqual(ifExample.Body, ifBodyExample) {
		t.Errorf("If body is not set correctly")
	}

	if ifExample.Cond != ifCondExample {
		t.Errorf("If cond is not set correctly")
	}
}

func TestIfElse(t *testing.T) {
	ifQuestionExample := InputQuestion{litexpr.StrLit{"Did you sell a house in 2010?"}, vari.VarDecl{vari.VarId{"hasSoldHouse"}, vartype.BoolType{}}}
	ifBodyExample := StmtList{[]Question{ifQuestionExample}, []Conditional{}}
	ifCondExample := litexpr.BoolLit{true}

	elseQuestionExample := InputQuestion{litexpr.StrLit{"What was the selling price?"}, vari.VarDecl{vari.VarId{"sellingPrice"}, vartype.IntType{}}}
	elseBodyExample := StmtList{[]Question{elseQuestionExample}, []Conditional{}}

	ifElseExample := IfElse{ifCondExample, ifBodyExample, elseBodyExample}

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
	questionExample := InputQuestion{litexpr.StrLit{"Did you sell a house in 2010?"}, vari.VarDecl{vari.VarId{"hasSoldHouse"}, vartype.BoolType{}}}
	questionListExample := []Question{questionExample}

	ifExample := If{litexpr.BoolLit{true}, StmtList{}}
	conditionalListExample := []Conditional{ifExample}

	stmtListExample := StmtList{questionListExample, conditionalListExample}

	if len(stmtListExample.Questions) != len(questionListExample) {
		t.Errorf("Stmtlist questions list is not set correctly")
	}

	if len(stmtListExample.Conditionals) != len(conditionalListExample) {
		t.Errorf("Stmtlist conditionals list is not set correctly")
	}
}
