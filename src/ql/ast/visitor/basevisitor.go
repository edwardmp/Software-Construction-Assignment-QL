package visitor

import "ql/interfaces"

type BaseVisitor struct {
	interfaces.Visitor
}

func (b *BaseVisitor) VisitAdd(a interfaces.Add, s interfaces.Symbols) {

}

func (b *BaseVisitor) VisitAnd(a interfaces.And, s interfaces.Symbols) {

}

func (b *BaseVisitor) VisitDiv(d interfaces.Div, s interfaces.Symbols) {

}

func (b *BaseVisitor) VisitEq(e interfaces.Eq, s interfaces.Symbols) {

}

func (b *BaseVisitor) VisitGEq(g interfaces.GEq, s interfaces.Symbols) {

}

func (b *BaseVisitor) VisitGT(g interfaces.GT, s interfaces.Symbols) {

}

func (b *BaseVisitor) VisitLEq(l interfaces.LEq, s interfaces.Symbols) {

}

func (b *BaseVisitor) VisitLT(l interfaces.LT, s interfaces.Symbols) {

}

func (b *BaseVisitor) VisitMul(m interfaces.Mul, s interfaces.Symbols) {

}

func (b *BaseVisitor) VisitNEq(n interfaces.NEq, s interfaces.Symbols) {

}

func (b *BaseVisitor) VisitOr(o interfaces.Or, s interfaces.Symbols) {

}

func (b *BaseVisitor) VisitSub(su interfaces.Sub, s interfaces.Symbols) {

}

func (b *BaseVisitor) VisitBoolLit(bo interfaces.BoolLit, s interfaces.Symbols) {

}

func (b *BaseVisitor) VisitIntLit(i interfaces.IntLit, s interfaces.Symbols) {

}

func (b *BaseVisitor) VisitStrLit(st interfaces.StrLit, s interfaces.Symbols) {

}

func (b *BaseVisitor) VisitNeg(n interfaces.Neg, s interfaces.Symbols) {

}

func (b *BaseVisitor) VisitNot(n interfaces.Not, s interfaces.Symbols) {

}

func (b *BaseVisitor) VisitPos(p interfaces.Pos, s interfaces.Symbols) {

}

func (b *BaseVisitor) VisitVarExpr(va interfaces.VarExpr, s interfaces.Symbols) {

}

func (b *BaseVisitor) VisitVarDecl(va interfaces.VarDecl, s interfaces.Symbols) {

}

func (b *BaseVisitor) VisitVarId(va interfaces.VarId, s interfaces.Symbols) {

}

func (b *BaseVisitor) VisitIntType(i interfaces.IntType, s interfaces.Symbols) {

}

func (b *BaseVisitor) VisitBoolType(bo interfaces.BoolType, s interfaces.Symbols) {

}

func (b *BaseVisitor) VisitStringType(st interfaces.StringType, s interfaces.Symbols) {

}

func (b *BaseVisitor) VisitForm(f interfaces.Form, s interfaces.Symbols) {

}

func (b *BaseVisitor) VisitComputedQuestion(c interfaces.ComputedQuestion, s interfaces.Symbols) {

}

func (b *BaseVisitor) VisitInputQuestion(i interfaces.InputQuestion, s interfaces.Symbols) {

}

func (b *BaseVisitor) VisitIf(i interfaces.If, s interfaces.Symbols) {

}

func (b *BaseVisitor) VisitIfElse(i interfaces.IfElse, s interfaces.Symbols) {

}

func (b *BaseVisitor) VisitStmtList(st interfaces.StmtList, s interfaces.Symbols) {

}
