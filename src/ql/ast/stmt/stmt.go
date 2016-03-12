package stmt

import (
	"fmt"
	"ql/ast/node"
	"ql/interfaces"
)

type Stmt struct {
	node.Node
}

func NewStmt(sourceInfo interface{}) Stmt {
	return Stmt{node.NewNode(sourceInfo)}
}

func NewStmtNoSourceInfo() Stmt {
	return NewStmt(nil)
}

// slices don't support equality checking, so have to do it manually
func SlicesEqual(a StmtList, b StmtList) bool {
	questionsA := a.Questions
	questionsB := b.Questions

	if len(questionsA) != len(questionsB) {
		return false
	}

	for i, questionA := range questionsA {
		//questionA = questionA.ResetSourceInfo().(interfaces.Question)
		questionB := questionsB[i]
		//questionB = questionB.ResetSourceInfo().(interfaces.Question)

		if questionA != questionB {
			return false
		}
	}

	conditionalsA := a.Conditionals
	conditionalsB := b.Conditionals

	for i := range conditionalsA {
		if conditionalsA[i].(interfaces.Conditional).EvalCondition() != conditionalsB[i].(interfaces.Conditional).EvalCondition() {
			return false
		}

		if !SlicesEqualConditional(conditionalsA[i].(interfaces.Conditional), conditionalsB[i].(interfaces.Conditional)) {
			return false
		}
	}

	return true
}

func SlicesEqualConditional(ifA, ifB interfaces.Conditional) bool {
	if fmt.Sprintf("%T", ifA) != fmt.Sprintf("%T", ifB) {
		panic("Types not equal") // TODO replace with assert
	}

	switch t := ifA.(type) {
	default:
		panic(fmt.Sprintf("unexpected Conditional type %T\n", t))
	case If:
		bodyA := ifA.(If).Body
		bodyB := ifB.(If).Body
		return SlicesEqual(bodyA, bodyB)
	case IfElse:
		return SlicesEqual(ifA.(IfElse).IfBody, ifB.(IfElse).IfBody) && SlicesEqual(ifA.(IfElse).ElseBody, ifB.(IfElse).ElseBody)
	}
}
