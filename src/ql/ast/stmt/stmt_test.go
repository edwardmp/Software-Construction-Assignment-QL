package stmt

import (
	"fmt"
	"ql/ast/expr/binaryoperatorexpr"
	"ql/ast/expr/lit"
	"ql/ast/stmt"
	"ql/ast/vari"
	"ql/ast/vari/vartype"
	"testing"
)

// TODO PUT IN HELPER FILE
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

/* Tests for statements */

func TestFormWithEmptyContent(t *testing.T) {
	identifier := vari.VarId{"TestForm"}
	exampleForm := stmt.Form{identifier, stmt.StmtList{}}

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
	questionExample := stmt.InputQuestion{lit.StrLit{"What was the selling price?"}, vari.VarDecl{vari.VarId{"sellingPrice"}, vartype.IntType{}}}
	questionsListExample := []stmt.Question{questionExample}
	stmtListExample := stmt.StmtList{Questions: questionsListExample}
	exampleForm := stmt.Form{identifier, stmtListExample}

	if len(exampleForm.Content.Questions) != 1 {
		t.Errorf("Form content questions does not have 1 question while it should")
	}

	if !slicesEqual(exampleForm.Content, stmtListExample) {
		t.Errorf("Form content not set correctly")
	}
}

func TestInputQuestion(t *testing.T) {
	exampleLabel := lit.StrLit{"Did you sell a house in 2010?"}
	exampleVarDecl := vari.VarDecl{vari.VarId{"hasSoldHouse"}, vartype.BoolType{}}

	exampleQuestion := stmt.InputQuestion{exampleLabel, exampleVarDecl}

	if exampleQuestion.Label != exampleLabel {
		t.Errorf("Question label is not set correctly")
	}
}

func TestComputedQuestion(t *testing.T) {
	exampleLabel := lit.StrLit{"Value residue"}
	exampleVarDecl := vari.VarDecl{vari.VarId{"hasSoldHouse"}, vartype.IntType{}}
	exampleComputation := binaryoperatorexpr.Sub{lit.IntLit{10}, lit.IntLit{5}}

	exampleQuestion := stmt.ComputedQuestion{exampleLabel, exampleVarDecl, exampleComputation}

	if exampleQuestion.Label != exampleLabel {
		t.Errorf("Computed question label is not set correctly")
	}

	if exampleQuestion.Computation != exampleComputation {
		t.Errorf("Computed question computation is not set correctly")
	}
}

func TestIf(t *testing.T) {
	questionExample := stmt.InputQuestion{lit.StrLit{"What was the selling price?"}, vari.VarDecl{vari.VarId{"sellingPrice"}, vartype.IntType{}}}
	ifBodyExample := stmt.StmtList{[]stmt.Question{questionExample}, []stmt.Conditional{}}
	ifCondExample := lit.BoolLit{true}
	ifExample := stmt.If{ifCondExample, ifBodyExample}

	if !slicesEqual(ifExample.Body, ifBodyExample) {
		t.Errorf("If body is not set correctly")
	}

	if ifExample.Cond != ifCondExample {
		t.Errorf("If cond is not set correctly")
	}
}

func TestIfElse(t *testing.T) {
	ifQuestionExample := stmt.InputQuestion{lit.StrLit{"Did you sell a house in 2010?"}, vari.VarDecl{vari.VarId{"hasSoldHouse"}, vartype.BoolType{}}}
	ifBodyExample := stmt.StmtList{[]stmt.Question{ifQuestionExample}, []stmt.Conditional{}}
	ifCondExample := lit.BoolLit{true}

	elseQuestionExample := stmt.InputQuestion{lit.StrLit{"What was the selling price?"}, vari.VarDecl{vari.VarId{"sellingPrice"}, vartype.IntType{}}}
	elseBodyExample := stmt.StmtList{[]stmt.Question{elseQuestionExample}, []stmt.Conditional{}}

	ifElseExample := stmt.IfElse{ifCondExample, ifBodyExample, elseBodyExample}

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
	questionExample := stmt.InputQuestion{lit.StrLit{"Did you sell a house in 2010?"}, vari.VarDecl{vari.VarId{"hasSoldHouse"}, vartype.BoolType{}}}
	questionListExample := []stmt.Question{questionExample}

	ifExample := stmt.If{lit.BoolLit{true}, stmt.StmtList{}}
	conditionalListExample := []stmt.Conditional{ifExample}

	stmtListExample := stmt.StmtList{questionListExample, conditionalListExample}

	if len(stmtListExample.Questions) != len(questionListExample) {
		t.Errorf("Stmtlist questions list is not set correctly")
	}

	if len(stmtListExample.Conditionals) != len(conditionalListExample) {
		t.Errorf("Stmtlist conditionals list is not set correctly")
	}
}
