package visitor

import "ql/interfaces"

type BaseVisitor struct {
	interfaces.Visitor
}

func (b *BaseVisitor) VisitAdd(a interfaces.Add, sy interface{}) {

}

func (b *BaseVisitor) VisitAnd(a interfaces.And, sy interface{}) {

}

func (b *BaseVisitor) VisitDiv(d interfaces.Div, sy interface{}) {

}

func (b *BaseVisitor) VisitEq(e interfaces.Eq, sy interface{}) {

}

func (b *BaseVisitor) VisitGEq(g interfaces.GEq, sy interface{}) {

}

func (b *BaseVisitor) VisitGT(g interfaces.GT, sy interface{}) {

}

func (b *BaseVisitor) VisitLEq(l interfaces.LEq, sy interface{}) {

}

func (b *BaseVisitor) VisitLT(l interfaces.LT, sy interface{}) {

}

func (b *BaseVisitor) VisitMul(m interfaces.Mul, sy interface{}) {

}

func (b *BaseVisitor) VisitNEq(n interfaces.NEq, sy interface{}) {

}

func (b *BaseVisitor) VisitOr(o interfaces.Or, sy interface{}) {

}

func (b *BaseVisitor) VisitSub(su interfaces.Sub, sy interface{}) {

}

func (b *BaseVisitor) VisitBoolLit(bo interfaces.BoolLit, sy interface{}) {

}

func (b *BaseVisitor) VisitIntLit(i interfaces.IntLit, sy interface{}) {

}

func (b *BaseVisitor) VisitStrLit(st interfaces.StrLit, sy interface{}) {

}

func (b *BaseVisitor) VisitNeg(n interfaces.Neg, sy interface{}) {

}

func (b *BaseVisitor) VisitNot(n interfaces.Not, sy interface{}) {

}

func (b *BaseVisitor) VisitPos(p interfaces.Pos, sy interface{}) {

}

func (b *BaseVisitor) VisitVarExpr(va interfaces.VarExpr, sy interface{}) {

}

func (b *BaseVisitor) VisitVarDecl(va interfaces.VarDecl, sy interface{}) {

}

func (b *BaseVisitor) VisitVarId(va interfaces.VarId, sy interface{}) {

}

func (b *BaseVisitor) VisitIntType(i interfaces.IntType, sy interface{}) {

}

func (b *BaseVisitor) VisitBoolType(bo interfaces.BoolType, sy interface{}) {

}

func (b *BaseVisitor) VisitStringType(st interfaces.StringType, sy interface{}) {

}

func (b *BaseVisitor) VisitForm(f interfaces.Form, sy interface{}) {

}

func (b *BaseVisitor) VisitComputedQuestion(c interfaces.ComputedQuestion, sy interface{}) {

}

func (b *BaseVisitor) VisitInputQuestion(i interfaces.InputQuestion, sy interface{}) {

}

func (b *BaseVisitor) VisitIf(i interfaces.If, sy interface{}) {

}

func (b *BaseVisitor) VisitIfElse(i interfaces.IfElse, sy interface{}) {

}

func (b *BaseVisitor) VisitStmtList(st interfaces.StmtList, sy interface{}) {

}
