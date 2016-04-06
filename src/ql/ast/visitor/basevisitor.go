package visitor

import "ql/interfaces"

type BaseVisitor struct {
}

func NewBaseVisitor() *BaseVisitor {
	return new(BaseVisitor)
}

func (this *BaseVisitor) VisitAdd(a interfaces.Add, context interface{}) {

}

func (this *BaseVisitor) VisitAnd(a interfaces.And, context interface{}) {

}

func (this *BaseVisitor) VisitDiv(d interfaces.Div, context interface{}) {

}

func (this *BaseVisitor) VisitEq(e interfaces.Eq, context interface{}) {

}

func (this *BaseVisitor) VisitGEq(g interfaces.GEq, context interface{}) {

}

func (this *BaseVisitor) VisitGT(g interfaces.GT, context interface{}) {

}

func (this *BaseVisitor) VisitLEq(l interfaces.LEq, context interface{}) {

}

func (this *BaseVisitor) VisitLT(l interfaces.LT, context interface{}) {

}

func (this *BaseVisitor) VisitMul(m interfaces.Mul, context interface{}) {

}

func (this *BaseVisitor) VisitNEq(n interfaces.NEq, context interface{}) {

}

func (this *BaseVisitor) VisitOr(o interfaces.Or, context interface{}) {

}

func (this *BaseVisitor) VisitSub(su interfaces.Sub, context interface{}) {

}

func (this *BaseVisitor) VisitBoolLiteral(bo interfaces.BoolLiteral, context interface{}) {

}

func (this *BaseVisitor) VisitIntegerLiteral(i interfaces.IntegerLiteral, context interface{}) {

}

func (this *BaseVisitor) VisitStringLiteral(st interfaces.StringLiteral, context interface{}) {

}

func (this *BaseVisitor) VisitNeg(n interfaces.Neg, context interface{}) {

}

func (this *BaseVisitor) VisitNot(n interfaces.Not, context interface{}) {

}

func (this *BaseVisitor) VisitPos(p interfaces.Pos, context interface{}) {

}

func (this *BaseVisitor) VisitVarExpr(va interfaces.VarExpr, context interface{}) {

}

func (this *BaseVisitor) VisitVarDecl(va interfaces.VarDecl, context interface{}) {

}

func (this *BaseVisitor) VisitVarId(va interfaces.VarId, context interface{}) {

}

func (this *BaseVisitor) VisitIntType(i interfaces.IntType, context interface{}) {

}

func (this *BaseVisitor) VisitBoolType(bo interfaces.BoolType, context interface{}) {

}

func (this *BaseVisitor) VisitStringType(st interfaces.StringType, context interface{}) {

}

func (this *BaseVisitor) VisitForm(f interfaces.Form, context interface{}) {

}

func (this *BaseVisitor) VisitComputedQuestion(c interfaces.ComputedQuestion, context interface{}) {

}

func (this *BaseVisitor) VisitInputQuestion(i interfaces.InputQuestion, context interface{}) {

}

func (this *BaseVisitor) VisitIf(i interfaces.If, context interface{}) {

}

func (this *BaseVisitor) VisitIfElse(i interfaces.IfElse, context interface{}) {

}

func (this *BaseVisitor) VisitStmtList(st interfaces.StmtList, context interface{}) {

}
