package vari

import "ql/interfaces"

func (this VarDecl) Accept(v interfaces.Visitor, s interfaces.Symbols) interface{} {
	v.VisitVarDecl(this, s)

	this.Ident.Accept(v, s)

	return nil
}

func (this VarId) Accept(v interfaces.Visitor, s interfaces.Symbols) interface{} {
	v.VisitVarId(this, s)
	return nil
}

func (this IntType) Accept(v interfaces.Visitor, s interfaces.Symbols) interface{} {
	v.VisitIntType(this, s)
	return nil
}

func (this BoolType) Accept(v interfaces.Visitor, s interfaces.Symbols) interface{} {
	v.VisitBoolType(this, s)
	return nil
}

func (this StringType) Accept(v interfaces.Visitor, s interfaces.Symbols) interface{} {
	v.VisitStringType(this, s)
	return nil
}
