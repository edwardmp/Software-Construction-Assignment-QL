package stmt

import (
	log "github.com/Sirupsen/logrus"
	"ql/interfaces"
)

func (f Form) Accept(v interfaces.Visitor, s interface{}) interface{} {
	log.Debug("Accept form")
	v.VisitForm(f, s)

	f.Identifier.Accept(v, s)
	f.Content.Accept(v, s)

	return nil
}

func (i If) Accept(v interfaces.Visitor, s interface{}) interface{} {
	log.Debug("Accept If")
	v.VisitIf(i, s)

	i.Cond.Accept(v, s)
	i.Body.Accept(v, s)

	return nil
}

func (i IfElse) Accept(v interfaces.Visitor, s interface{}) interface{} {
	log.Debug("Accept IfElse")
	v.VisitIfElse(i, s)

	i.Cond.Accept(v, s)
	i.IfBody.Accept(v, s)
	i.ElseBody.Accept(v, s)

	return nil
}

func (i InputQuestion) Accept(v interfaces.Visitor, s interface{}) interface{} {
	log.Debug("Accept InputQuestion")
	v.VisitInputQuestion(i, s)

	i.Label.Accept(v, s)
	i.VarDecl.Accept(v, s)

	return nil
}

func (c ComputedQuestion) Accept(v interfaces.Visitor, s interface{}) interface{} {
	log.Debug("Accept ComputedQuestion")
	v.VisitComputedQuestion(c, s)

	c.Label.Accept(v, s)
	c.VarDecl.Accept(v, s)
	c.Computation.Accept(v, s)

	return nil
}

func (s StmtList) Accept(v interfaces.Visitor, sy interface{}) interface{} {
	log.Debug("Accept StmtList")
	v.VisitStmtList(s, sy)

	for _, question := range s.Questions {
		question.Accept(v, sy)
	}

	for _, conditional := range s.Conditionals {
		conditional.Accept(v, sy)
	}

	return nil
}
