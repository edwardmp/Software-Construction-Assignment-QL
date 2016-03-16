package vari

import "ql/interfaces"

func (this VarDecl) Accept(visitor interfaces.Visitor, context interface{}) interface{} {
	visitor.VisitVarDecl(this, context)

	this.Ident.Accept(visitor, context)

	return nil
}

func (this VarId) Accept(visitor interfaces.Visitor, context interface{}) interface{} {
	visitor.VisitVarId(this, context)

	return nil
}
