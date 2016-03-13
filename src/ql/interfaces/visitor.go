package interfaces

type Visitor interface {
	/* Expr */
	VisitAdd(Add, Symbols)
	VisitAnd(And, Symbols)
	VisitDiv(Div, Symbols)
	VisitEq(Eq, Symbols)
	VisitGEq(GEq, Symbols)
	VisitGT(GT, Symbols)
	VisitLEq(LEq, Symbols)
	VisitLT(LT, Symbols)
	VisitMul(Mul, Symbols)
	VisitNEq(NEq, Symbols)
	VisitOr(Or, Symbols)
	VisitSub(Sub, Symbols)
	VisitBoolLit(BoolLit, Symbols)
	VisitIntLit(IntLit, Symbols)
	VisitStrLit(StrLit, Symbols)
	VisitNeg(Neg, Symbols)
	VisitNot(Not, Symbols)
	VisitPos(Pos, Symbols)
	VisitVarExpr(VarExpr, Symbols)
	/* Vari */
	VisitVarDecl(VarDecl, Symbols)
	VisitVarId(VarId, Symbols)
	VisitIntType(IntType, Symbols)
	VisitBoolType(BoolType, Symbols)
	VisitStringType(StringType, Symbols)
	/* Stmt */
	VisitForm(Form, Symbols)
	VisitComputedQuestion(ComputedQuestion, Symbols)
	VisitInputQuestion(InputQuestion, Symbols)
	VisitIf(If, Symbols)
	VisitIfElse(IfElse, Symbols)
	VisitStmtList(StmtList, Symbols)
}
