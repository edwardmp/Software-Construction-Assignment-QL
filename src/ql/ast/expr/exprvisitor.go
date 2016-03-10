package expr

import "ql/interfaces"

func (a Add) Accept(v interfaces.Visitor, s interface{}) interface{} {
	v.VisitAdd(a, s)

	a.Lhs.Accept(v, s)
	a.Rhs.Accept(v, s)

	return nil
}

func (a And) Accept(v interfaces.Visitor, s interface{}) interface{} {
	v.VisitAnd(a, s)

	a.Lhs.Accept(v, s)
	a.Rhs.Accept(v, s)

	return nil
}

func (d Div) Accept(v interfaces.Visitor, s interface{}) interface{} {
	v.VisitDiv(d, s)

	d.Lhs.Accept(v, s)
	d.Rhs.Accept(v, s)

	return nil
}

func (e Eq) Accept(v interfaces.Visitor, s interface{}) interface{} {
	v.VisitEq(e, s)

	e.Lhs.Accept(v, s)
	e.Rhs.Accept(v, s)

	return nil
}

func (g GEq) Accept(v interfaces.Visitor, s interface{}) interface{} {
	v.VisitEq(g, s)

	g.Lhs.Accept(v, s)
	g.Rhs.Accept(v, s)

	return nil
}

func (g GT) Accept(v interfaces.Visitor, s interface{}) interface{} {
	v.VisitGT(g, s)

	g.Lhs.Accept(v, s)
	g.Rhs.Accept(v, s)

	return nil
}

func (l LEq) Accept(v interfaces.Visitor, s interface{}) interface{} {
	v.VisitLEq(l, s)

	l.Lhs.Accept(v, s)
	l.Rhs.Accept(v, s)

	return nil
}

func (l LT) Accept(v interfaces.Visitor, s interface{}) interface{} {
	v.VisitLT(l, s)

	l.Lhs.Accept(v, s)
	l.Rhs.Accept(v, s)

	return nil
}

func (m Mul) Accept(v interfaces.Visitor, s interface{}) interface{} {
	v.VisitMul(m, s)

	m.Lhs.Accept(v, s)
	m.Rhs.Accept(v, s)

	return nil
}

func (n NEq) Accept(v interfaces.Visitor, s interface{}) interface{} {
	v.VisitNEq(n, s)

	n.Lhs.Accept(v, s)
	n.Rhs.Accept(v, s)

	return nil
}

func (o Or) Accept(v interfaces.Visitor, s interface{}) interface{} {
	v.VisitOr(o, s)

	o.Lhs.Accept(v, s)
	o.Rhs.Accept(v, s)

	return nil
}

func (s Sub) Accept(v interfaces.Visitor, sym interface{}) interface{} {
	v.VisitSub(s, sym)

	s.Lhs.Accept(v, s)
	s.Rhs.Accept(v, s)

	return nil
}

func (b BoolLit) Accept(v interfaces.Visitor, s interface{}) interface{} {
	v.VisitBoolLit(b, s)
	return nil
}

func (i IntLit) Accept(v interfaces.Visitor, s interface{}) interface{} {
	v.VisitIntLit(i, s)
	return nil
}

func (s StrLit) Accept(v interfaces.Visitor, sy interface{}) interface{} {
	v.VisitStrLit(s, sy)
	return nil
}

func (n Neg) Accept(v interfaces.Visitor, s interface{}) interface{} {
	v.VisitNeg(n, s)
	n.Value.Accept(v, s)
	return nil
}

func (n Not) Accept(v interfaces.Visitor, s interface{}) interface{} {
	v.VisitNot(n, s)
	n.Value.Accept(v, s)
	return nil
}

func (p Pos) Accept(v interfaces.Visitor, s interface{}) interface{} {
	v.VisitPos(p, s)
	p.Value.Accept(v, s)
	return nil
}

func (va VarExpr) Accept(v interfaces.Visitor, s interface{}) interface{} {
	v.VisitVarExpr(va, s)
	va.Identifier.Accept(v, s)
	return nil
}
