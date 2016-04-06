package interfaces

type Visitor interface {
	/* Expr */
	VisitAdd(Add, interface{})
	VisitAnd(And, interface{})
	VisitDiv(Div, interface{})
	VisitEq(Eq, interface{})
	VisitGEq(GEq, interface{})
	VisitGT(GT, interface{})
	VisitLEq(LEq, interface{})
	VisitLT(LT, interface{})
	VisitMul(Mul, interface{})
	VisitNEq(NEq, interface{})
	VisitOr(Or, interface{})
	VisitSub(Sub, interface{})
	VisitBoolLiteral(BoolLiteral, interface{})
	VisitIntegerLiteral(IntegerLiteral, interface{})
	VisitStringLiteral(StringLiteral, interface{})
	VisitNeg(Neg, interface{})
	VisitNot(Not, interface{})
	VisitPos(Pos, interface{})
	VisitVarExpr(VarExpr, interface{})
	/* Vari */
	VisitVarDecl(VarDecl, interface{})
	VisitVarID(VarID, interface{})
	VisitIntType(IntType, interface{})
	VisitBoolType(BoolType, interface{})
	VisitStringType(StringType, interface{})
	/* Stmt */
	VisitForm(Form, interface{})
	VisitComputedQuestion(ComputedQuestion, interface{})
	VisitInputQuestion(InputQuestion, interface{})
	VisitIf(If, interface{})
	VisitIfElse(IfElse, interface{})
	VisitStmtList(StmtList, interface{})
}
