package expr

import (
	log "github.com/Sirupsen/logrus"
	"ql/interfaces"
)

func (a Add) Accept(v interfaces.Visitor, s interface{}) interface{} {
	log.Debug("Accept Add")
	v.VisitAdd(a, s)
	acceptLhsAndRhs(a, v, s)

	return nil
}

func (a And) Accept(v interfaces.Visitor, s interface{}) interface{} {
	log.Debug("Accept And")
	v.VisitAnd(a, s)
	acceptLhsAndRhs(a, v, s)

	return nil
}

func (d Div) Accept(v interfaces.Visitor, s interface{}) interface{} {
	log.Debug("Accept Div")
	v.VisitDiv(d, s)
	acceptLhsAndRhs(d, v, s)

	return nil
}

func (e Eq) Accept(v interfaces.Visitor, s interface{}) interface{} {
	log.Debug("Accept Eq")
	v.VisitEq(e, s)
	acceptLhsAndRhs(e, v, s)

	return nil
}

func (g GEq) Accept(v interfaces.Visitor, s interface{}) interface{} {
	log.Debug("Accept GEq")
	v.VisitEq(g, s)
	acceptLhsAndRhs(g, v, s)

	return nil
}

func (g GT) Accept(v interfaces.Visitor, s interface{}) interface{} {
	log.Debug("Accept GT")
	v.VisitGT(g, s)
	acceptLhsAndRhs(g, v, s)

	return nil
}

func (l LEq) Accept(v interfaces.Visitor, s interface{}) interface{} {
	log.Debug("Accept LEq")
	v.VisitLEq(l, s)

	acceptLhsAndRhs(l, v, s)

	return nil
}

func (l LT) Accept(v interfaces.Visitor, s interface{}) interface{} {
	log.Debug("Accept LT")
	v.VisitLT(l, s)
	acceptLhsAndRhs(l, v, s)

	return nil
}

func (m Mul) Accept(v interfaces.Visitor, s interface{}) interface{} {
	log.Debug("Accept MUl")
	v.VisitMul(m, s)
	acceptLhsAndRhs(m, v, s)

	return nil
}

func (n NEq) Accept(v interfaces.Visitor, s interface{}) interface{} {
	log.Debug("Accept NEq")
	v.VisitNEq(n, s)
	acceptLhsAndRhs(n, v, s)

	return nil
}

func (o Or) Accept(v interfaces.Visitor, s interface{}) interface{} {
	log.Debug("Accept Or")
	v.VisitOr(o, s)
	acceptLhsAndRhs(o, v, s)

	return nil
}

func (s Sub) Accept(v interfaces.Visitor, sym interface{}) interface{} {
	log.Debug("Accept Sub")
	v.VisitSub(s, sym)
	acceptLhsAndRhs(s, v, sym)

	return nil
}

func (b BoolLit) Accept(v interfaces.Visitor, s interface{}) interface{} {
	log.Debug("Accept BoolLit")
	v.VisitBoolLit(b, s)

	return nil
}

func (i IntLit) Accept(v interfaces.Visitor, s interface{}) interface{} {
	log.Debug("Accept IntLit")
	v.VisitIntLit(i, s)

	return nil
}

func (s StrLit) Accept(v interfaces.Visitor, sy interface{}) interface{} {
	log.Debug("Accept StrLit")
	v.VisitStrLit(s, sy)

	return nil
}

func (n Neg) Accept(v interfaces.Visitor, s interface{}) interface{} {
	log.Debug("Accept Neg")
	v.VisitNeg(n, s)
	n.Value.Accept(v, s)

	return nil
}

func (n Not) Accept(v interfaces.Visitor, s interface{}) interface{} {
	log.Debug("Accept Not")
	v.VisitNot(n, s)
	n.Value.Accept(v, s)

	return nil
}

func (p Pos) Accept(v interfaces.Visitor, s interface{}) interface{} {
	log.Debug("Accept Pos")
	v.VisitPos(p, s)
	p.Value.Accept(v, s)

	return nil
}

func (va VarExpr) Accept(v interfaces.Visitor, s interface{}) interface{} {
	log.Debug("Accept VarExpr")
	v.VisitVarExpr(va, s)
	va.Identifier.Accept(v, s)

	return nil
}

func (this Expr) Accept(v interfaces.Visitor, s interface{}) interface{} {
	panic("Expr Accept method not overridden")

	return nil
}

func acceptLhsAndRhs(binaryExpr interfaces.BinaryOperatorExpr, v interfaces.Visitor, s interface{}) {
	binaryExpr.GetLhs().Accept(v, s)
	binaryExpr.GetRhs().Accept(v, s)
}
