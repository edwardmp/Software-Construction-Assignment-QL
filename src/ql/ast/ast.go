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
func NewPos(value interface{}, sourcePosInfo interface{}) (interfaces.Expr, error) {
	expr := expr.NewPos(value.(interfaces.Expr))
	expr.SetSourceInfo(sourcePosInfo.(token.Pos))
	return expr, nil
}

func NewNeg(value interface{}, sourcePosInfo interface{}) (interfaces.Expr, error) {
	expr := expr.NewNeg(value.(interfaces.Expr))
	expr.SetSourceInfo(sourcePosInfo.(token.Pos))
	return expr, nil
}

func NewNot(value interface{}, sourcePosInfo interface{}) (interfaces.Expr, error) {
	expr := expr.NewNot(value.(interfaces.Expr))
	expr.SetSourceInfo(sourcePosInfo.(token.Pos))
	return expr, nil
}

func NewVarExpr(identifier interface{}) (interfaces.Expr, error) {
	varId := identifier.(vari.VarId)
	expr := expr.NewVarExpr(varId)
	expr.SetSourceInfo(varId.GetSourceInfo())
	return expr, nil
}

/* binary operator expressins */
func NewMul(lhs interface{}, rhs interface{}, sourcePosInfo interface{}) (interfaces.Expr, error) {
	expr := expr.NewMul(lhs.(interfaces.Expr), rhs.(interfaces.Expr))
	expr.SetSourceInfo(sourcePosInfo.(token.Pos))
	return expr, nil
}

func NewDiv(lhs interface{}, rhs interface{}, sourcePosInfo interface{}) (interfaces.Expr, error) {
	expr := expr.NewDiv(lhs.(interfaces.Expr), rhs.(interfaces.Expr))
	expr.SetSourceInfo(sourcePosInfo.(token.Pos))
	return expr, nil
}

func NewAdd(lhs interface{}, rhs interface{}, sourcePosInfo interface{}) (interfaces.Expr, error) {
	expr := expr.NewAdd(lhs.(interfaces.Expr), rhs.(interfaces.Expr))
	expr.SetSourceInfo(sourcePosInfo.(token.Pos))
	return expr, nil
}

func NewSub(lhs interface{}, rhs interface{}, sourcePosInfo interface{}) (interfaces.Expr, error) {
	expr := expr.NewSub(lhs.(interfaces.Expr), rhs.(interfaces.Expr))
	expr.SetSourceInfo(sourcePosInfo.(token.Pos))
	return expr, nil
}

func NewEq(lhs interface{}, rhs interface{}, sourcePosInfo interface{}) (interfaces.Expr, error) {
	expr := expr.NewEq(lhs.(interfaces.Expr), rhs.(interfaces.Expr))
	expr.SetSourceInfo(sourcePosInfo.(token.Pos))
	return expr, nil
}

func NewNEq(lhs interface{}, rhs interface{}, sourcePosInfo interface{}) (interfaces.Expr, error) {
	expr := expr.NewNEq(lhs.(interfaces.Expr), rhs.(interfaces.Expr))
	expr.SetSourceInfo(sourcePosInfo.(token.Pos))
	return expr, nil
}

func NewGT(lhs interface{}, rhs interface{}, sourcePosInfo interface{}) (interfaces.Expr, error) {
	expr := expr.NewGT(lhs.(interfaces.Expr).(interfaces.Expr), rhs.(interfaces.Expr))
	expr.SetSourceInfo(sourcePosInfo.(token.Pos))
	return expr, nil
}

func NewLT(lhs interface{}, rhs interface{}, sourcePosInfo interface{}) (interfaces.Expr, error) {
	expr := expr.NewLT(lhs.(interfaces.Expr), rhs.(interfaces.Expr))
	expr.SetSourceInfo(sourcePosInfo.(token.Pos))
	return expr, nil
}

func NewGEq(lhs interface{}, rhs interface{}, sourcePosInfo interface{}) (interfaces.Expr, error) {
	expr := expr.NewGEq(lhs.(interfaces.Expr), rhs.(interfaces.Expr))
	expr.SetSourceInfo(sourcePosInfo.(token.Pos))
	return expr, nil
}

func NewLEq(lhs interface{}, rhs interface{}, sourcePosInfo interface{}) (interfaces.Expr, error) {
	expr := expr.NewLEq(lhs.(interfaces.Expr), rhs.(interfaces.Expr))
	expr.SetSourceInfo(sourcePosInfo.(token.Pos))
	return expr, nil
}

func NewAnd(lhs interface{}, rhs interface{}, sourcePosInfo interface{}) (interfaces.Expr, error) {
	expr := expr.NewAnd(lhs.(interfaces.Expr), rhs.(interfaces.Expr))
	expr.SetSourceInfo(sourcePosInfo.(token.Pos))
	return expr, nil
}

func NewOr(lhs interface{}, rhs interface{}, sourcePosInfo interface{}) (interfaces.Expr, error) {
	expr := expr.NewOr(lhs.(interfaces.Expr), rhs.(interfaces.Expr))
	expr.SetSourceInfo(sourcePosInfo.(token.Pos))
	return expr, nil
}

/* literals */
func NewIntLit(litValueToken interface{}) (interfaces.Expr, error) {
	sourcePosInfo := litValueToken.(*token.Token).Pos
	value, err := util.IntValue(litValueToken.(*token.Token).Lit)
	expr := expr.NewIntLit(int(value))
	expr.SetSourceInfo(sourcePosInfo)
	return expr, err
}

func NewBoolLit(value bool, sourcePosInfo interface{}) (interfaces.Expr, error) {
	expr := expr.NewBoolLit(value)
	expr.SetSourceInfo(sourcePosInfo.(token.Pos))
	return expr, nil
}

func NewStrLit(valueToken interface{}) (interfaces.Expr, error) {
	sourcePosInfo := valueToken.(*token.Token).Pos
	literalString := stringLiteralTokensToString(valueToken.(*token.Token))
	expr := expr.NewStrLit(literalString)
	expr.SetSourceInfo(sourcePosInfo)
	return expr, nil
}

/** Vari **/

func NewVarDecl(ident interface{}, typeIdent interface{}, sourcePosInfo interface{}) (interfaces.VarDecl, error) {
	vari := vari.NewVarDecl(ident.(interfaces.VarId), typeIdent.(interfaces.ValueType))
	vari.SetSourceInfo(sourcePosInfo.(token.Pos))
	return vari, nil
}

func NewVarId(identToken interface{}) (vari.VarId, error) {
	sourcePosInfo := identToken.(*token.Token).Pos
	identifierString := string(identToken.(*token.Token).Lit)
	vari := vari.NewVarId(identifierString)
	vari.SetSourceInfo(sourcePosInfo)
	return vari, nil
}

func NewIntType(typeTokenLit interface{}) (interfaces.IntType, error) {
	token := typeTokenLit.(*token.Token)
	expr := expr.NewIntType()
	expr.SetSourceInfo(token.Pos)

	return expr, nil
}

func NewBoolType(typeTokenLit interface{}) (interfaces.BoolType, error) {
	token := typeTokenLit.(*token.Token)
	expr := expr.NewBoolType()
	expr.SetSourceInfo(token.Pos)
	return expr, nil
}

func NewStringType(typeTokenLit interface{}) (interfaces.StringType, error) {
	token := typeTokenLit.(*token.Token)
	expr := expr.NewStringType()
	expr.SetSourceInfo(token.Pos)
	return expr, nil
}

/** Statements **/

func NewForm(identifier interface{}, body interface{}, sourcePosInfo interface{}) (interfaces.Form, error) {
	stmt := stmt.NewForm(identifier.(vari.VarId), body.(stmt.StmtList))
	stmt.SetSourceInfo(sourcePosInfo.(token.Pos))
	return stmt, nil
}

func NewInputQuestion(label interface{}, varDecl interface{}) (interfaces.InputQuestion, error) {
	labelStrLit := label.(expr.StrLit)
	stmt := stmt.NewInputQuestion(labelStrLit, varDecl.(vari.VarDecl))
	stmt.SetSourceInfo(labelStrLit.GetSourceInfo())
	return stmt, nil
}

func NewComputedQuestion(label interface{}, varDecl interface{}, computation interface{}, sourcePosInfo interface{}) (interfaces.ComputedQuestion, error) {
	stmt := stmt.NewComputedQuestion(label.(expr.StrLit), varDecl.(vari.VarDecl), computation.(interfaces.Expr))
	stmt.SetSourceInfo(sourcePosInfo.(token.Pos))
	return stmt, nil
}

func NewStmtList(stmtElt interface{}) (interfaces.StmtList, error) {
	stmtEltTypeAsserted := stmtElt.(interfaces.Stmt)
	stmt := stmt.NewEmptyStmtList()
	stmt.SetSourceInfo(stmtEltTypeAsserted.GetSourceInfo())
	return stmt.AddToCorrectSlice(stmtElt), nil
}

func NewEmptyStmtList(sourcePosInfo interface{}) (interfaces.StmtList, error) {
	stmt := stmt.NewEmptyStmtList()
	stmt.SetSourceInfo(sourcePosInfo.(token.Pos))
	return stmt, nil
}

func AppendStmt(stmtList, stmtElt interface{}) (interfaces.StmtList, error) {
	stmt := stmtList.(stmt.StmtList).AddToCorrectSlice(stmtElt)
	return stmt, nil
}

func NewIf(cond interface{}, body interface{}, sourcePosInfo interface{}) (interfaces.If, error) {
	stmt := stmt.NewIf(cond.(interfaces.Expr), body.(stmt.StmtList))
	stmt.SetSourceInfo(sourcePosInfo.(token.Pos))
	return stmt, nil
}

func NewIfElse(cond interface{}, ifBody interface{}, elseBody interface{}, sourcePosInfo interface{}) (interfaces.IfElse, error) {
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
