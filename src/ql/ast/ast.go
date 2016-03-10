package ast

import (
	//"fmt"
	"ql/ast/expr"
	"ql/ast/stmt"
	"ql/ast/vari"
	"ql/interfaces"
	"ql/token"
	"strconv"
)

func stringLiteralTokensToString(a interface{}) (str string) {
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

/** expressions **/

/* unary operator expressions */
func NewPos(value interface{}) (interfaces.Expr, error) {
	return expr.NewPos(value.(interfaces.Expr)), nil
}

func NewNeg(value interface{}) (interfaces.Expr, error) {
	return expr.NewNeg(value.(interfaces.Expr)), nil
}

func NewNot(value interface{}) (interfaces.Expr, error) {
	return expr.NewNot(value.(interfaces.Expr)), nil
}

func NewVarExpr(identifier interface{}) (interfaces.Expr, error) {
	return expr.VarExpr{identifier.(vari.VarId)}, nil
}

/* binary operator expressins */
func NewMul(lhs interface{}, rhs interface{}) (interfaces.Expr, error) {
	return expr.NewMul(lhs.(interfaces.Expr), rhs.(interfaces.Expr)), nil
}

func NewDiv(lhs interface{}, rhs interface{}) (interfaces.Expr, error) {
	return expr.NewDiv(lhs.(interfaces.Expr), rhs.(interfaces.Expr)), nil
}

func NewAdd(lhs interface{}, rhs interface{}) (interfaces.Expr, error) {
	return expr.NewAdd(lhs.(interfaces.Expr), rhs.(interfaces.Expr)), nil
}

func NewSub(lhs interface{}, rhs interface{}) (interfaces.Expr, error) {
	return expr.NewSub(lhs.(interfaces.Expr), rhs.(interfaces.Expr)), nil
}

func NewEq(lhs interface{}, rhs interface{}) (interfaces.Expr, error) {
	return expr.NewEq(lhs.(interfaces.Expr), rhs.(interfaces.Expr)), nil
}

func NewNEq(lhs interface{}, rhs interface{}) (interfaces.Expr, error) {
	return expr.NewNEq(lhs.(interfaces.Expr), rhs.(interfaces.Expr)), nil
}

func NewGT(lhs interface{}, rhs interface{}) (interfaces.Expr, error) {
	return expr.NewGT(lhs.(interfaces.Expr).(interfaces.Expr), rhs.(interfaces.Expr)), nil
}

func NewLT(lhs interface{}, rhs interface{}) (interfaces.Expr, error) {
	return expr.NewLT(lhs.(interfaces.Expr), rhs.(interfaces.Expr)), nil
}

func NewGEq(lhs interface{}, rhs interface{}) (interfaces.Expr, error) {
	return expr.NewGEq(lhs.(interfaces.Expr), rhs.(interfaces.Expr)), nil
}

func NewLEq(lhs interface{}, rhs interface{}) (interfaces.Expr, error) {
	return expr.NewLEq(lhs.(interfaces.Expr), rhs.(interfaces.Expr)), nil
}

func NewAnd(lhs interface{}, rhs interface{}) (interfaces.Expr, error) {
	return expr.NewAnd(lhs.(interfaces.Expr), rhs.(interfaces.Expr)), nil
}

func NewOr(lhs interface{}, rhs interface{}) (interfaces.Expr, error) {
	return expr.NewOr(lhs.(interfaces.Expr), rhs.(interfaces.Expr)), nil
}

/* Literals */

func NewIntLit(value int64, e error) (interfaces.Expr, error) {
	return expr.IntLit{int(value)}, nil
}

func NewBoolLit(value bool) (interfaces.Expr, error) {
	return expr.BoolLit{value}, nil
}

func NewStrLit(value interface{}) (interfaces.Expr, error) {
	literalString := stringLiteralTokensToString(value)
	return expr.StrLit{literalString}, nil
}

/* vartypes */

func NewIntType() (vari.IntType, error) {
	return vari.IntType{}, nil
}

func NewBoolType() (vari.BoolType, error) {
	return vari.BoolType{}, nil
}

func NewStringType() (vari.StringType, error) {
	return vari.StringType{}, nil
}

/** statements **/

func NewForm(identifier interface{}, body interface{}) (stmt.Form, error) {
	return stmt.Form{identifier.(vari.VarId), body.(stmt.StmtList)}, nil
}

func NewInputQuestion(label interface{}, varDecl interface{}) (stmt.InputQuestion, error) {
	return stmt.InputQuestion{label.(expr.StrLit), varDecl.(vari.VarDecl)}, nil
}

func NewComputedQuestion(label interface{}, varDecl interface{}, computation interface{}) (stmt.ComputedQuestion, error) {
	return stmt.ComputedQuestion{label.(expr.StrLit), varDecl.(vari.VarDecl), computation.(interfaces.Expr)}, nil
}

func NewStmtList(stmtElt interface{}) (stmt.StmtList, error) {
	s := stmt.StmtList{}
	return s.AddToCorrectSlice(stmtElt), nil
}

func NewEmptyStmtList() (stmt.StmtList, error) {
	return stmt.StmtList{}, nil
}

func AppendStmt(stmtList, stmtElt interface{}) (stmt.StmtList, error) {
	return stmtList.(stmt.StmtList).AddToCorrectSlice(stmtElt), nil
}

func NewIf(cond interface{}, body interface{}) (stmt.If, error) {
	return stmt.If{cond.(interfaces.Expr), body.(stmt.StmtList)}, nil
}

func NewIfElse(cond interface{}, ifBody interface{}, elseBody interface{}) (stmt.IfElse, error) {
	return stmt.IfElse{cond.(interfaces.Expr), ifBody.(stmt.StmtList), elseBody.(stmt.StmtList)}, nil
}

func NewVarDecl(ident interface{}, typeIdent interface{}) (interfaces.VarDecl, error) {
	return vari.VarDecl{ident.(interfaces.VarId), typeIdent.(interfaces.VarType)}, nil
}

func NewVarId(ident interface{}) (vari.VarId, error) {
	identifierString := string(ident.(*token.Token).Lit)
	return vari.VarId{identifierString}, nil
}
