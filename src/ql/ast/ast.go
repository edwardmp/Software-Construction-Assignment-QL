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

/** expressions **/

/* unary operator expressions */
func NewPos(value interface{}, sourcePosInfo interface{}) (interfaces.Expr, error) {
	return expr.NewPos(value.(interfaces.Expr), sourcePosInfo), nil
}

func NewNeg(value interface{}, sourcePosInfo interface{}) (interfaces.Expr, error) {
	return expr.NewNeg(value.(interfaces.Expr), sourcePosInfo), nil
}

func NewNot(value interface{}, sourcePosInfo interface{}) (interfaces.Expr, error) {
	return expr.NewNot(value.(interfaces.Expr), sourcePosInfo), nil
}

func NewVarExpr(identifier interface{}) (interfaces.Expr, error) {
	varId := identifier.(vari.VarId)
	return expr.NewVarExpr(varId, varId.GetSourceInfo()), nil
}

/* binary operator expressins */
func NewMul(lhs interface{}, rhs interface{}, sourcePosInfo interface{}) (interfaces.Expr, error) {
	return expr.NewMul(lhs.(interfaces.Expr), rhs.(interfaces.Expr), sourcePosInfo), nil
}

func NewDiv(lhs interface{}, rhs interface{}, sourcePosInfo interface{}) (interfaces.Expr, error) {
	return expr.NewDiv(lhs.(interfaces.Expr), rhs.(interfaces.Expr), sourcePosInfo), nil
}

func NewAdd(lhs interface{}, rhs interface{}, sourcePosInfo interface{}) (interfaces.Expr, error) {
	return expr.NewAdd(lhs.(interfaces.Expr), rhs.(interfaces.Expr), sourcePosInfo), nil
}

func NewSub(lhs interface{}, rhs interface{}, sourcePosInfo interface{}) (interfaces.Expr, error) {
	return expr.NewSub(lhs.(interfaces.Expr), rhs.(interfaces.Expr), sourcePosInfo), nil
}

func NewEq(lhs interface{}, rhs interface{}, sourcePosInfo interface{}) (interfaces.Expr, error) {
	return expr.NewEq(lhs.(interfaces.Expr), rhs.(interfaces.Expr), sourcePosInfo), nil
}

func NewNEq(lhs interface{}, rhs interface{}, sourcePosInfo interface{}) (interfaces.Expr, error) {
	return expr.NewNEq(lhs.(interfaces.Expr), rhs.(interfaces.Expr), sourcePosInfo), nil
}

func NewGT(lhs interface{}, rhs interface{}, sourcePosInfo interface{}) (interfaces.Expr, error) {
	return expr.NewGT(lhs.(interfaces.Expr).(interfaces.Expr), rhs.(interfaces.Expr), sourcePosInfo), nil
}

func NewLT(lhs interface{}, rhs interface{}, sourcePosInfo interface{}) (interfaces.Expr, error) {
	return expr.NewLT(lhs.(interfaces.Expr), rhs.(interfaces.Expr), sourcePosInfo), nil
}

func NewGEq(lhs interface{}, rhs interface{}, sourcePosInfo interface{}) (interfaces.Expr, error) {
	return expr.NewGEq(lhs.(interfaces.Expr), rhs.(interfaces.Expr), sourcePosInfo), nil
}

func NewLEq(lhs interface{}, rhs interface{}, sourcePosInfo interface{}) (interfaces.Expr, error) {
	return expr.NewLEq(lhs.(interfaces.Expr), rhs.(interfaces.Expr), sourcePosInfo), nil
}

func NewAnd(lhs interface{}, rhs interface{}, sourcePosInfo interface{}) (interfaces.Expr, error) {
	return expr.NewAnd(lhs.(interfaces.Expr), rhs.(interfaces.Expr), sourcePosInfo), nil
}

func NewOr(lhs interface{}, rhs interface{}, sourcePosInfo interface{}) (interfaces.Expr, error) {
	return expr.NewOr(lhs.(interfaces.Expr), rhs.(interfaces.Expr), sourcePosInfo), nil
}

/* Literals */

func NewIntLit(litValueToken interface{}) (interfaces.Expr, error) {
	sourcePosInfo := litValueToken.(*token.Token).Pos
	value, err := util.IntValue(litValueToken.(*token.Token).Lit)
	return expr.NewIntLit(int(value), sourcePosInfo), err
}

func NewBoolLit(value bool, sourcePosInfo interface{}) (interfaces.Expr, error) {
	return expr.NewBoolLit(value, sourcePosInfo), nil
}

func NewStrLit(valueToken interface{}) (interfaces.Expr, error) {
	sourcePosInfo := valueToken.(*token.Token).Pos
	literalString := stringLiteralTokensToString(valueToken.(*token.Token))
	return expr.NewStrLit(literalString, sourcePosInfo), nil
}

/** vari **/

func NewIntType(typeTokenLit interface{}) (vari.IntType, error) {
	token := typeTokenLit.(*token.Token)
	return vari.NewIntType(token.Pos), nil
}

func NewBoolType(typeTokenLit interface{}) (vari.BoolType, error) {
	token := typeTokenLit.(*token.Token)
	return vari.NewBoolType(token.Pos), nil
}

func NewStringType(typeTokenLit interface{}) (vari.StringType, error) {
	token := typeTokenLit.(*token.Token)
	return vari.NewStringType(token.Pos), nil
}

/** statements **/

func NewForm(identifier interface{}, body interface{}, sourcePosInfo interface{}) (stmt.Form, error) {
	return stmt.NewForm(identifier.(vari.VarId), body.(stmt.StmtList), sourcePosInfo), nil
}

func NewInputQuestion(label interface{}, varDecl interface{}) (stmt.InputQuestion, error) {
	labelStrLit := label.(expr.StrLit)
	return stmt.NewInputQuestion(labelStrLit, varDecl.(vari.VarDecl), labelStrLit.GetSourceInfo()), nil
}

func NewComputedQuestion(label interface{}, varDecl interface{}, computation interface{}, sourcePosInfo interface{}) (stmt.ComputedQuestion, error) {
	return stmt.NewComputedQuestion(label.(expr.StrLit), varDecl.(vari.VarDecl), computation.(interfaces.Expr), sourcePosInfo), nil
}

func NewStmtList(stmtElt interface{}) (stmt.StmtList, error) {
	stmtEltTypeAsserted := stmtElt.(interfaces.Stmt)
	s := stmt.NewEmptyStmtList(stmtEltTypeAsserted.GetSourceInfo())
	return s.AddToCorrectSlice(stmtElt), nil
}

func NewEmptyStmtList(sourcePosInfo interface{}) (stmt.StmtList, error) {
	return stmt.NewEmptyStmtList(sourcePosInfo), nil
}

func AppendStmt(stmtList, stmtElt interface{}) (stmt.StmtList, error) {
	return stmtList.(stmt.StmtList).AddToCorrectSlice(stmtElt), nil
}

func NewIf(cond interface{}, body interface{}, sourcePosInfo interface{}) (stmt.If, error) {
	return stmt.NewIf(cond.(interfaces.Expr), body.(stmt.StmtList), sourcePosInfo), nil
}

func NewIfElse(cond interface{}, ifBody interface{}, elseBody interface{}, sourcePosInfo interface{}) (stmt.IfElse, error) {
	return stmt.NewIfElse(cond.(interfaces.Expr), ifBody.(stmt.StmtList), elseBody.(stmt.StmtList), sourcePosInfo), nil
}

func NewVarDecl(ident interface{}, typeIdent interface{}, sourcePosInfo interface{}) (interfaces.VarDecl, error) {
	return vari.NewVarDecl(ident.(interfaces.VarId), typeIdent.(interfaces.VarType), sourcePosInfo), nil
}

func NewVarId(identToken interface{}) (vari.VarId, error) {
	sourcePosInfo := identToken.(*token.Token).Pos
	identifierString := string(identToken.(*token.Token).Lit)
	return vari.NewVarId(identifierString, sourcePosInfo), nil
}

func stringLiteralTokensToString(token *token.Token) (str string) {
	astr, err := strconv.Unquote(string(token.Lit))
	if err != nil {
		return ""
	}

	return astr
}
