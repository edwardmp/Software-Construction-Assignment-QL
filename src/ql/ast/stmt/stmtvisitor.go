package stmt

import "ql/interfaces"

func (c ComputedQuestion) Accept(v interfaces.Visitor, s interface{}) interface{} {
	v.VisitComputedQuestion(c, s)
	return nil
}

func (f Form) Accept(v interfaces.Visitor, s interface{}) interface{} {
	v.VisitForm(f, s)
	return nil
}

func (i If) Accept(v interfaces.Visitor, s interface{}) interface{} {
	v.VisitIf(i, s)
	return nil
}

func (i IfElse) Accept(v interfaces.Visitor, s interface{}) interface{} {
	v.VisitIfElse(i, s)
	return nil
}

func (i InputQuestion) Accept(v interfaces.Visitor, s interface{}) interface{} {
	v.VisitInputQuestion(i, s)
	return nil
}

func (s StmtList) Accept(v interfaces.Visitor, sy interface{}) interface{} {
	v.VisitStmtList(s, sy)
	return nil
}
