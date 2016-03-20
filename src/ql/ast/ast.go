package ast

import (
	"ql/ast/expr"
	"ql/ast/stmt"
	"ql/ast/vari"
	"ql/interfaces"
	"ql/token"
	"ql/util"
	"strconv"
)

const (
	TRUE  = true
	FALSE = false
)

/** Expressions **/

/* unary operator expressions */
func NewPosNode(value interface{}, sourcePosInfo interface{}) (interfaces.Expr, error) {
	expr := expr.NewPos(value.(interfaces.Expr))
	expr.SetSourceInfo(sourcePosInfo.(token.Pos))
	return expr, nil
}

func NewNegNode(value interface{}, sourcePosInfo interface{}) (interfaces.Expr, error) {
	expr := expr.NewNeg(value.(interfaces.Expr))
	expr.SetSourceInfo(sourcePosInfo.(token.Pos))
	return expr, nil
}

func NewNotNode(value interface{}, sourcePosInfo interface{}) (interfaces.Expr, error) {
	expr := expr.NewNot(value.(interfaces.Expr))
	expr.SetSourceInfo(sourcePosInfo.(token.Pos))
	return expr, nil
}

func NewVarExprNode(identifier interface{}) (interfaces.Expr, error) {
	varId := identifier.(vari.VarId)
	expr := expr.NewVarExpr(varId)
	expr.SetSourceInfo(varId.GetSourceInfo())
	return expr, nil
}

/* binary operator expressins */
func NewMulNode(lhs interface{}, rhs interface{}, sourcePosInfo interface{}) (interfaces.Expr, error) {
	expr := expr.NewMul(lhs.(interfaces.Expr), rhs.(interfaces.Expr))
	expr.SetSourceInfo(sourcePosInfo.(token.Pos))
	return expr, nil
}

func NewDivNode(lhs interface{}, rhs interface{}, sourcePosInfo interface{}) (interfaces.Expr, error) {
	expr := expr.NewDiv(lhs.(interfaces.Expr), rhs.(interfaces.Expr))
	expr.SetSourceInfo(sourcePosInfo.(token.Pos))
	return expr, nil
}

func NewAddNode(lhs interface{}, rhs interface{}, sourcePosInfo interface{}) (interfaces.Expr, error) {
	expr := expr.NewAdd(lhs.(interfaces.Expr), rhs.(interfaces.Expr))
	expr.SetSourceInfo(sourcePosInfo.(token.Pos))
	return expr, nil
}

func NewSubNode(lhs interface{}, rhs interface{}, sourcePosInfo interface{}) (interfaces.Expr, error) {
	expr := expr.NewSub(lhs.(interfaces.Expr), rhs.(interfaces.Expr))
	expr.SetSourceInfo(sourcePosInfo.(token.Pos))
	return expr, nil
}

func NewEqNode(lhs interface{}, rhs interface{}, sourcePosInfo interface{}) (interfaces.Expr, error) {
	expr := expr.NewEq(lhs.(interfaces.Expr), rhs.(interfaces.Expr))
	expr.SetSourceInfo(sourcePosInfo.(token.Pos))
	return expr, nil
}

func NewNEqNode(lhs interface{}, rhs interface{}, sourcePosInfo interface{}) (interfaces.Expr, error) {
	expr := expr.NewNEq(lhs.(interfaces.Expr), rhs.(interfaces.Expr))
	expr.SetSourceInfo(sourcePosInfo.(token.Pos))
	return expr, nil
}

func NewGTNode(lhs interface{}, rhs interface{}, sourcePosInfo interface{}) (interfaces.Expr, error) {
	expr := expr.NewGT(lhs.(interfaces.Expr).(interfaces.Expr), rhs.(interfaces.Expr))
	expr.SetSourceInfo(sourcePosInfo.(token.Pos))
	return expr, nil
}

func NewLTNode(lhs interface{}, rhs interface{}, sourcePosInfo interface{}) (interfaces.Expr, error) {
	expr := expr.NewLT(lhs.(interfaces.Expr), rhs.(interfaces.Expr))
	expr.SetSourceInfo(sourcePosInfo.(token.Pos))
	return expr, nil
}

func NewGEqNode(lhs interface{}, rhs interface{}, sourcePosInfo interface{}) (interfaces.Expr, error) {
	expr := expr.NewGEq(lhs.(interfaces.Expr), rhs.(interfaces.Expr))
	expr.SetSourceInfo(sourcePosInfo.(token.Pos))
	return expr, nil
}

func NewLEqNode(lhs interface{}, rhs interface{}, sourcePosInfo interface{}) (interfaces.Expr, error) {
	expr := expr.NewLEq(lhs.(interfaces.Expr), rhs.(interfaces.Expr))
	expr.SetSourceInfo(sourcePosInfo.(token.Pos))
	return expr, nil
}

func NewAndNode(lhs interface{}, rhs interface{}, sourcePosInfo interface{}) (interfaces.Expr, error) {
	expr := expr.NewAnd(lhs.(interfaces.Expr), rhs.(interfaces.Expr))
	expr.SetSourceInfo(sourcePosInfo.(token.Pos))
	return expr, nil
}

func NewOrNode(lhs interface{}, rhs interface{}, sourcePosInfo interface{}) (interfaces.Expr, error) {
	expr := expr.NewOr(lhs.(interfaces.Expr), rhs.(interfaces.Expr))
	expr.SetSourceInfo(sourcePosInfo.(token.Pos))
	return expr, nil
}

/* literals */
func NewIntLitNode(litValueToken interface{}) (interfaces.Expr, error) {
	sourcePosInfo := litValueToken.(*token.Token).Pos
	value, err := util.IntValue(litValueToken.(*token.Token).Lit)
	expr := expr.NewIntLit(int(value))
	expr.SetSourceInfo(sourcePosInfo)
	return expr, err
}

func NewBoolLitNode(value bool, sourcePosInfo interface{}) (interfaces.Expr, error) {
	expr := expr.NewBoolLit(value)
	expr.SetSourceInfo(sourcePosInfo.(token.Pos))
	return expr, nil
}

func NewStrLitNode(valueToken interface{}) (interfaces.Expr, error) {
	sourcePosInfo := valueToken.(*token.Token).Pos
	literalString := stringLiteralTokensToString(valueToken.(*token.Token))
	expr := expr.NewStrLit(literalString)
	expr.SetSourceInfo(sourcePosInfo)
	return expr, nil
}

/** Vari **/

func NewVarDeclNode(ident interface{}, typeIdent interface{}, sourcePosInfo interface{}) (interfaces.VarDecl, error) {
	vari := vari.NewVarDecl(ident.(interfaces.VarId), typeIdent.(interfaces.ValueType))
	vari.SetSourceInfo(sourcePosInfo.(token.Pos))
	return vari, nil
}

func NewVarIdNode(identToken interface{}) (vari.VarId, error) {
	sourcePosInfo := identToken.(*token.Token).Pos
	identifierString := string(identToken.(*token.Token).Lit)
	vari := vari.NewVarId(identifierString)
	vari.SetSourceInfo(sourcePosInfo)
	return vari, nil
}

func NewIntTypeNode(typeTokenLit interface{}) (interfaces.IntType, error) {
	token := typeTokenLit.(*token.Token)
	expr := expr.NewIntType()
	expr.SetSourceInfo(token.Pos)

	return expr, nil
}

func NewBoolTypeNode(typeTokenLit interface{}) (interfaces.BoolType, error) {
	token := typeTokenLit.(*token.Token)
	expr := expr.NewBoolType()
	expr.SetSourceInfo(token.Pos)
	return expr, nil
}

func NewStringTypeNode(typeTokenLit interface{}) (interfaces.StringType, error) {
	token := typeTokenLit.(*token.Token)
	expr := expr.NewStringType()
	expr.SetSourceInfo(token.Pos)
	return expr, nil
}

/** Statements **/

func NewFormNode(identifier interface{}, body interface{}, sourcePosInfo interface{}) (interfaces.Form, error) {
	stmt := stmt.NewForm(identifier.(vari.VarId), body.(stmt.StmtList))
	stmt.SetSourceInfo(sourcePosInfo.(token.Pos))
	return stmt, nil
}

func NewInputQuestionNode(label interface{}, varDecl interface{}) (interfaces.InputQuestion, error) {
	labelStrLit := label.(expr.StrLit)
	stmt := stmt.NewInputQuestion(labelStrLit, varDecl.(vari.VarDecl))
	stmt.SetSourceInfo(labelStrLit.GetSourceInfo())
	return stmt, nil
}

func NewComputedQuestionNode(label interface{}, varDecl interface{}, computation interface{}, sourcePosInfo interface{}) (interfaces.ComputedQuestion, error) {
	stmt := stmt.NewComputedQuestion(label.(expr.StrLit), varDecl.(vari.VarDecl), computation.(interfaces.Expr))
	stmt.SetSourceInfo(sourcePosInfo.(token.Pos))
	return stmt, nil
}

func NewStmtListNode(stmtElt interface{}) (interfaces.StmtList, error) {
	stmtEltTypeAsserted := stmtElt.(interfaces.Stmt)
	stmt := stmt.NewEmptyStmtList()
	stmt.SetSourceInfo(stmtEltTypeAsserted.GetSourceInfo())
	return stmt.AddToCorrectSlice(stmtElt), nil
}

func NewEmptyStmtListNode(sourcePosInfo interface{}) (interfaces.StmtList, error) {
	stmt := stmt.NewEmptyStmtList()
	stmt.SetSourceInfo(sourcePosInfo.(token.Pos))
	return stmt, nil
}

func AppendStmt(stmtList, stmtElt interface{}) (interfaces.StmtList, error) {
	stmt := stmtList.(stmt.StmtList).AddToCorrectSlice(stmtElt)
	return stmt, nil
}

func NewIfNode(cond interface{}, body interface{}, sourcePosInfo interface{}) (interfaces.If, error) {
	stmt := stmt.NewIf(cond.(interfaces.Expr), body.(stmt.StmtList))
	stmt.SetSourceInfo(sourcePosInfo.(token.Pos))
	return stmt, nil
}

func NewIfElseNode(cond interface{}, ifBody interface{}, elseBody interface{}, sourcePosInfo interface{}) (interfaces.IfElse, error) {
	stmt := stmt.NewIfElse(cond.(interfaces.Expr), ifBody.(stmt.StmtList), elseBody.(stmt.StmtList))
	stmt.SetSourceInfo(sourcePosInfo.(token.Pos))
	return stmt, nil
}

// TODO place in util?
func stringLiteralTokensToString(token *token.Token) (str string) {
	astr, err := strconv.Unquote(string(token.Lit))
	if err != nil {
		return ""
	}

	return astr
}
