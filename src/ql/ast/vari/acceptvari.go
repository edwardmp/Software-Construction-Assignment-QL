package vari

import "ql/interfaces"

func (va VarDecl) Accept(v interfaces.Visitor, s interface{}) interface{} {
	v.VisitVarDecl(va, s)

	va.Ident.Accept(v, s)

	return nil
}

func (va VarId) Accept(v interfaces.Visitor, s interface{}) interface{} {
	v.VisitVarId(va, s)
	return nil
}

func (i IntType) Accept(v interfaces.Visitor, s interface{}) interface{} {
	v.VisitIntType(i, s)
	return nil
}

func (b BoolType) Accept(v interfaces.Visitor, s interface{}) interface{} {
	v.VisitBoolType(b, s)
	return nil
}

func (s StringType) Accept(v interfaces.Visitor, sy interface{}) interface{} {
	v.VisitStringType(s, sy)
	return nil
}
