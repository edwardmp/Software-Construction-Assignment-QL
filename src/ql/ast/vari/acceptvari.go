package vari

import "ql/interfaces"

func (this VarDecl) Accept(visitor interfaces.Visitor, context interface{}) interface{} {
	visitor.VisitVarDecl(this, context)

	this.VariableIdentifier().Accept(visitor, context)

	return nil
}

func (this VarID) Accept(visitor interfaces.Visitor, context interface{}) interface{} {
	visitor.VisitVarID(this, context)

	return nil
}
