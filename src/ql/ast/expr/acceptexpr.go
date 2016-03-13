package expr

import (
	log "github.com/Sirupsen/logrus"
	"ql/interfaces"
)

func (this Add) Accept(v interfaces.Visitor, s interfaces.Symbols) interface{} {
	log.Debug("Accept Add")
	v.VisitAdd(this, s)
	acceptLhsAndRhs(this, v, s)

	return nil
}

func (this And) Accept(v interfaces.Visitor, s interfaces.Symbols) interface{} {
	log.Debug("Accept And")
	v.VisitAnd(this, s)
	acceptLhsAndRhs(this, v, s)

	return nil
}

func (this Div) Accept(v interfaces.Visitor, s interfaces.Symbols) interface{} {
	log.Debug("Accept Div")
	v.VisitDiv(this, s)
	acceptLhsAndRhs(this, v, s)

	return nil
}

func (this Eq) Accept(v interfaces.Visitor, s interfaces.Symbols) interface{} {
	log.Debug("Accept Eq")
	v.VisitEq(this, s)
	acceptLhsAndRhs(this, v, s)

	return nil
}

func (this GEq) Accept(v interfaces.Visitor, s interfaces.Symbols) interface{} {
	log.Debug("Accept GEq")
	v.VisitEq(this, s)
	acceptLhsAndRhs(this, v, s)

	return nil
}

func (this GT) Accept(v interfaces.Visitor, s interfaces.Symbols) interface{} {
	log.Debug("Accept GT")
	v.VisitGT(this, s)
	acceptLhsAndRhs(this, v, s)

	return nil
}

func (this LEq) Accept(v interfaces.Visitor, s interfaces.Symbols) interface{} {
	log.Debug("Accept LEq")
	v.VisitLEq(this, s)

	acceptLhsAndRhs(this, v, s)

	return nil
}

func (this LT) Accept(v interfaces.Visitor, s interfaces.Symbols) interface{} {
	log.Debug("Accept LT")
	v.VisitLT(this, s)
	acceptLhsAndRhs(this, v, s)

	return nil
}

func (this Mul) Accept(v interfaces.Visitor, s interfaces.Symbols) interface{} {
	log.Debug("Accept Mul")
	v.VisitMul(this, s)
	acceptLhsAndRhs(this, v, s)

	return nil
}

func (this NEq) Accept(v interfaces.Visitor, s interfaces.Symbols) interface{} {
	log.Debug("Accept NEq")
	v.VisitNEq(this, s)
	acceptLhsAndRhs(this, v, s)

	return nil
}

func (this Or) Accept(v interfaces.Visitor, s interfaces.Symbols) interface{} {
	log.Debug("Accept Or")
	v.VisitOr(this, s)
	acceptLhsAndRhs(this, v, s)

	return nil
}

func (this Sub) Accept(v interfaces.Visitor, s interfaces.Symbols) interface{} {
	log.Debug("Accept Sub")
	v.VisitSub(this, s)
	acceptLhsAndRhs(this, v, s)

	return nil
}

func (this IntLit) Accept(v interfaces.Visitor, s interfaces.Symbols) interface{} {
	log.Debug("Accept IntLit")
	v.VisitIntLit(this, s)

	return nil
}

func (this StrLit) Accept(v interfaces.Visitor, s interfaces.Symbols) interface{} {
	log.Debug("Accept StrLit")
	v.VisitStrLit(this, s)

	return nil
}

func (this Not) Accept(v interfaces.Visitor, s interfaces.Symbols) interface{} {
	log.Debug("Accept Not")
	v.VisitNot(this, s)
	this.Value.Accept(v, s)

	return nil
}

func (this Pos) Accept(v interfaces.Visitor, s interfaces.Symbols) interface{} {
	log.Debug("Accept Pos")
	v.VisitPos(this, s)
	this.Value.Accept(v, s)

	return nil
}

func (this VarExpr) Accept(v interfaces.Visitor, s interfaces.Symbols) interface{} {
	log.Debug("Accept VarExpr")
	v.VisitVarExpr(this, s)
	this.Identifier.Accept(v, s)

	return nil
}

func (this Expr) Accept(v interfaces.Visitor, s interfaces.Symbols) interface{} {
	panic("Expr Accept method not overridden")

	return nil
}

func acceptLhsAndRhs(binaryExpr interfaces.BinaryOperatorExpr, v interfaces.Visitor, s interfaces.Symbols) {
	binaryExpr.GetLhs().Accept(v, s)
	binaryExpr.GetRhs().Accept(v, s)
}
