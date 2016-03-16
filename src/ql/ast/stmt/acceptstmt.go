package stmt

import (
	log "github.com/Sirupsen/logrus"
	"ql/interfaces"
)

func (this Form) Accept(visitor interfaces.Visitor, context interface{}) interface{} {
	log.Debug("Accept form")
	visitor.VisitForm(this, context)

	this.Identifier.Accept(visitor, context)
	this.Content.Accept(visitor, context)

	return nil
}

func (this If) Accept(visitor interfaces.Visitor, context interface{}) interface{} {
	log.Debug("Accept If")
	visitor.VisitIf(this, context)

	this.Cond.Accept(visitor, context)
	this.Body.Accept(visitor, context)

	return nil
}

func (this IfElse) Accept(visitor interfaces.Visitor, context interface{}) interface{} {
	log.Debug("Accept IfElse")
	visitor.VisitIfElse(this, context)

	this.Cond.Accept(visitor, context)
	this.IfBody.Accept(visitor, context)
	this.ElseBody.Accept(visitor, context)

	return nil
}

func (this InputQuestion) Accept(visitor interfaces.Visitor, context interface{}) interface{} {
	log.Debug("Accept InputQuestion")
	visitor.VisitInputQuestion(this, context)

	this.Label.Accept(visitor, context)
	this.VarDecl.Accept(visitor, context)

	return nil
}

func (this ComputedQuestion) Accept(visitor interfaces.Visitor, context interface{}) interface{} {
	log.Debug("Accept ComputedQuestion")
	visitor.VisitComputedQuestion(this, context)

	this.Label.Accept(visitor, context)
	this.VarDecl.Accept(visitor, context)
	this.Computation.Accept(visitor, context)

	return nil
}

func (this StmtList) Accept(visitor interfaces.Visitor, context interface{}) interface{} {
	log.Debug("Accept StmtList")
	visitor.VisitStmtList(this, context)

	for _, question := range this.Questions {
		question.Accept(visitor, context)
	}

	for _, conditional := range this.Conditionals {
		conditional.Accept(visitor, context)
	}

	return nil
}
