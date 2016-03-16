package stmt

import (
	log "github.com/Sirupsen/logrus"
	"ql/interfaces"
)

func (this Form) Accept(v interfaces.Visitor, s interfaces.Symbols) interface{} {
	log.Debug("Accept form")
	v.VisitForm(this, s)

	this.Identifier.Accept(v, s)
	this.Content.Accept(v, s)

	return nil
}

func (this If) Accept(v interfaces.Visitor, s interfaces.Symbols) interface{} {
	log.Debug("Accept If")
	v.VisitIf(this, s)

	this.Cond.Accept(v, s)
	this.Body.Accept(v, s)

	return nil
}

func (this IfElse) Accept(v interfaces.Visitor, s interfaces.Symbols) interface{} {
	log.Debug("Accept IfElse")
	v.VisitIfElse(this, s)

	this.Cond.Accept(v, s)
	this.IfBody.Accept(v, s)
	this.ElseBody.Accept(v, s)

	return nil
}

func (this InputQuestion) Accept(v interfaces.Visitor, s interfaces.Symbols) interface{} {
	log.Debug("Accept InputQuestion")
	v.VisitInputQuestion(this, s)

	this.Label.Accept(v, s)
	this.VarDecl.Accept(v, s)

	return nil
}

func (this ComputedQuestion) Accept(v interfaces.Visitor, s interfaces.Symbols) interface{} {
	log.Debug("Accept ComputedQuestion")
	v.VisitComputedQuestion(this, s)

	this.Label.Accept(v, s)
	this.VarDecl.Accept(v, s)
	this.Computation.Accept(v, s)

	return nil
}

func (this StmtList) Accept(v interfaces.Visitor, s interfaces.Symbols) interface{} {
	log.Debug("Accept StmtList")
	v.VisitStmtList(this, s)

	for _, question := range this.Questions {
		question.Accept(v, s)
	}

	for _, conditional := range this.Conditionals {
		conditional.Accept(v, s)
	}

	return nil
}
