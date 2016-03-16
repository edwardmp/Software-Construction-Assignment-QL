package visitor

import "ql/interfaces"

type BaseVisitor struct {
}

func (b *BaseVisitor) VisitAdd(a interfaces.Add, context interface{}) {

}

func (b *BaseVisitor) VisitAnd(a interfaces.And, context interface{}) {

}

func (b *BaseVisitor) VisitDiv(d interfaces.Div, context interface{}) {

}

func (b *BaseVisitor) VisitEq(e interfaces.Eq, context interface{}) {

}

func (b *BaseVisitor) VisitGEq(g interfaces.GEq, context interface{}) {

}

func (b *BaseVisitor) VisitGT(g interfaces.GT, context interface{}) {

}

func (b *BaseVisitor) VisitLEq(l interfaces.LEq, context interface{}) {

}

func (b *BaseVisitor) VisitLT(l interfaces.LT, context interface{}) {

}

func (b *BaseVisitor) VisitMul(m interfaces.Mul, context interface{}) {

}

func (b *BaseVisitor) VisitNEq(n interfaces.NEq, context interface{}) {

}

func (b *BaseVisitor) VisitOr(o interfaces.Or, context interface{}) {

}

func (b *BaseVisitor) VisitSub(su interfaces.Sub, context interface{}) {

}

func (b *BaseVisitor) VisitBoolLit(bo interfaces.BoolLit, context interface{}) {

}

func (b *BaseVisitor) VisitIntLit(i interfaces.IntLit, context interface{}) {

}

func (b *BaseVisitor) VisitStrLit(st interfaces.StrLit, context interface{}) {

}

func (b *BaseVisitor) VisitNeg(n interfaces.Neg, context interface{}) {

}

func (b *BaseVisitor) VisitNot(n interfaces.Not, context interface{}) {

}

func (b *BaseVisitor) VisitPos(p interfaces.Pos, context interface{}) {

}

func (b *BaseVisitor) VisitVarExpr(va interfaces.VarExpr, context interface{}) {

}

func (b *BaseVisitor) VisitVarDecl(va interfaces.VarDecl, context interface{}) {

}

func (b *BaseVisitor) VisitVarId(va interfaces.VarId, context interface{}) {

}

func (b *BaseVisitor) VisitIntType(i interfaces.IntType, context interface{}) {

}

func (b *BaseVisitor) VisitBoolType(bo interfaces.BoolType, context interface{}) {

}

func (b *BaseVisitor) VisitStringType(st interfaces.StringType, context interface{}) {

}

func (b *BaseVisitor) VisitForm(f interfaces.Form, context interface{}) {

}

func (b *BaseVisitor) VisitComputedQuestion(c interfaces.ComputedQuestion, context interface{}) {

}

func (b *BaseVisitor) VisitInputQuestion(i interfaces.InputQuestion, context interface{}) {

}

func (b *BaseVisitor) VisitIf(i interfaces.If, context interface{}) {

}

func (b *BaseVisitor) VisitIfElse(i interfaces.IfElse, context interface{}) {

}

func (b *BaseVisitor) VisitStmtList(st interfaces.StmtList, context interface{}) {

}
