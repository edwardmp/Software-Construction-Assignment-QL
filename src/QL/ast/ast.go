package ast

import (
	"QL/ast/expr"
	"QL/ast/stmt"
	"QL/token"
	"strconv"
)

func tokenToString(a interface{}) (str string) {
	astr, err := strconv.Unquote(string(a.(*token.Token).Lit))
	if err != nil {
		return ""
	}
	return astr
}

var (
	TRUE  = bool(true)
	FALSE = bool(false)
)

/* expressions */
func NewPos(value interface{}) (expr.Expr, error) {
	return expr.Pos{value.(expr.Expr)}, nil
}

func NewNeg(value interface{}) (expr.Expr, error) {
	return expr.Neg{value.(expr.Expr)}, nil
}

func NewNot(value interface{}) (expr.Expr, error) {
	return expr.Not{value.(expr.Expr)}, nil
}

func NewMul(lhs interface{}, rhs interface{}) (expr.Expr, error) {
	return expr.Mul{lhs.(expr.Expr), rhs.(expr.Expr)}, nil
}

func NewDiv(lhs interface{}, rhs interface{}) (expr.Expr, error) {
	return expr.Div{lhs.(expr.Expr), rhs.(expr.Expr)}, nil
}

func NewAdd(lhs interface{}, rhs interface{}) (expr.Expr, error) {
	return expr.Add{lhs.(expr.Expr), rhs.(expr.Expr)}, nil
}

func NewSub(lhs interface{}, rhs interface{}) (expr.Expr, error) {
	return expr.Sub{lhs.(expr.Expr), rhs.(expr.Expr)}, nil
}

func NewEq(lhs interface{}, rhs interface{}) (expr.Expr, error) {
	return expr.Eq{lhs.(expr.Expr), rhs.(expr.Expr)}, nil
}

func NewNEq(lhs interface{}, rhs interface{}) (expr.Expr, error) {
	return expr.NEq{lhs.(expr.Expr), rhs.(expr.Expr)}, nil
}

func NewGT(lhs interface{}, rhs interface{}) (expr.Expr, error) {
	return expr.GT{lhs.(expr.Expr).(expr.Expr), rhs.(expr.Expr)}, nil
}

func NewLT(lhs interface{}, rhs interface{}) (expr.Expr, error) {
	return expr.LT{lhs.(expr.Expr), rhs.(expr.Expr)}, nil
}

func NewGEq(lhs interface{}, rhs interface{}) (expr.Expr, error) {
	return expr.GEq{lhs.(expr.Expr), rhs.(expr.Expr)}, nil
}

func NewLEq(lhs interface{}, rhs interface{}) (expr.Expr, error) {
	return expr.LEq{lhs.(expr.Expr), rhs.(expr.Expr)}, nil
}

func NewAnd(lhs interface{}, rhs interface{}) (expr.Expr, error) {
	return expr.And{lhs.(expr.Expr), rhs.(expr.Expr)}, nil
}

func NewOr(lhs interface{}, rhs interface{}) (expr.Expr, error) {
	return expr.Or{lhs.(expr.Expr), rhs.(expr.Expr)}, nil
}

/* Literals */
func NewIntLit(value int64, e error) (expr.Expr, error) {
	return expr.IntLit{int(value)}, nil
}

func NewBoolLit(value bool) (expr.Expr, error) {
	return expr.BoolLit{value}, nil
}

/* statements */
func NewQuestion(label interface{}, identifier interface{}, typeIdentifier interface{}) (stmt.Stmt, error) {
	labelString := tokenToString(label)
	identifierString := tokenToString(identifier)
	return stmt.Question{labelString, identifierString, typeIdentifier.(stmt.TypeIdentifier)}, nil
}

func NewForm(identifier interface{}, body interface{}) (stmt.Stmt, error) {
	identifierString := tokenToString(identifier)
	return stmt.Form{identifierString, body.(stmt.StmtList)}, nil
}

func NewStmtList(stmtElt interface{}) (stmt.StmtList, error) {
	return stmt.StmtList{stmtElt.(stmt.Stmt)}, nil
}

func AppendStmt(stmtList, stmtElt interface{}) (stmt.StmtList, error) {
	return append(stmtList.(stmt.StmtList), stmtElt.(stmt.Stmt)), nil
}

func NewIf(cond interface{}, stmtList interface{}) (stmt.Stmt, error) {
	return stmt.If{cond.(expr.Expr), stmtList.(*token.Token).Lit}, nil
}

func NewIfElse(cond interface{}, thenBody interface{}, elseBody interface{}) (stmt.Stmt, error) {
	return stmt.IfElse{cond.(expr.Expr), thenBody.(*token.Token).Lit, elseBody.(*token.Token).Lit}, nil
}
