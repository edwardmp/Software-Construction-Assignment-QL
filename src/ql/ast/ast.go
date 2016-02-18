package ast

import (
	//"fmt"
	"ql/ast/expr"
	"ql/ast/expr/binaryoperatorexpr"
	"ql/ast/expr/lit"
	"ql/ast/expr/unaryoperatorexpr"
	"ql/ast/stmt"
	"ql/ast/vari"
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
func NewPos(value interface{}) (expr.Expr, error) {
	return unaryoperatorexpr.Pos{value.(expr.Expr)}, nil
}

func NewNeg(value interface{}) (expr.Expr, error) {
	return unaryoperatorexpr.Neg{value.(expr.Expr)}, nil
}

func NewNot(value interface{}) (expr.Expr, error) {
	return unaryoperatorexpr.Not{value.(expr.Expr)}, nil
}

/* binary operator expressins */
func NewMul(lhs interface{}, rhs interface{}) (expr.Expr, error) {
	return binaryoperatorexpr.Mul{lhs.(expr.Expr), rhs.(expr.Expr)}, nil
}

func NewDiv(lhs interface{}, rhs interface{}) (expr.Expr, error) {
	return binaryoperatorexpr.Div{lhs.(expr.Expr), rhs.(expr.Expr)}, nil
}

func NewAdd(lhs interface{}, rhs interface{}) (expr.Expr, error) {
	return binaryoperatorexpr.Add{lhs.(expr.Expr), rhs.(expr.Expr)}, nil
}

func NewSub(lhs interface{}, rhs interface{}) (expr.Expr, error) {
	return binaryoperatorexpr.Sub{lhs.(expr.Expr), rhs.(expr.Expr)}, nil
}

func NewEq(lhs interface{}, rhs interface{}) (expr.Expr, error) {
	return binaryoperatorexpr.Eq{lhs.(expr.Expr), rhs.(expr.Expr)}, nil
}

func NewNEq(lhs interface{}, rhs interface{}) (expr.Expr, error) {
	return binaryoperatorexpr.NEq{lhs.(expr.Expr), rhs.(expr.Expr)}, nil
}

func NewGT(lhs interface{}, rhs interface{}) (expr.Expr, error) {
	return binaryoperatorexpr.GT{lhs.(expr.Expr).(expr.Expr), rhs.(expr.Expr)}, nil
}

func NewLT(lhs interface{}, rhs interface{}) (expr.Expr, error) {
	return binaryoperatorexpr.LT{lhs.(expr.Expr), rhs.(expr.Expr)}, nil
}

func NewGEq(lhs interface{}, rhs interface{}) (expr.Expr, error) {
	return binaryoperatorexpr.GEq{lhs.(expr.Expr), rhs.(expr.Expr)}, nil
}

func NewLEq(lhs interface{}, rhs interface{}) (expr.Expr, error) {
	return binaryoperatorexpr.LEq{lhs.(expr.Expr), rhs.(expr.Expr)}, nil
}

func NewAnd(lhs interface{}, rhs interface{}) (expr.Expr, error) {
	return binaryoperatorexpr.And{lhs.(expr.Expr), rhs.(expr.Expr)}, nil
}

func NewOr(lhs interface{}, rhs interface{}) (expr.Expr, error) {
	return binaryoperatorexpr.Or{lhs.(expr.Expr), rhs.(expr.Expr)}, nil
}

/* Literals */
func NewIntLit(value int64, e error) (expr.Expr, error) {
	return lit.IntLit{int(value)}, nil
}

func NewBoolLit(value bool) (expr.Expr, error) {
	return lit.BoolLit{value}, nil
}

func NewStrLit(value interface{}) (expr.Expr, error) {
	literalString := stringLiteralTokensToString(value)
	return lit.StrLit{literalString}, nil
}

func NewVarExpr(identifier interface{}) (expr.Expr, error) {
	return expr.VarExpr{identifier.(vari.VarId), nil}, nil
}

/* statements */

func NewForm(identifier interface{}, body interface{}) (stmt.Form, error) {
	return stmt.Form{identifier.(vari.VarId), body.(stmt.StmtList)}, nil
}

func NewInputQuestion(label interface{}, varDecl interface{}) (stmt.InputQuestion, error) {
	return stmt.InputQuestion{label.(lit.StrLit), varDecl.(vari.VarDecl)}, nil
}

func NewComputedQuestion(label interface{}, varDecl interface{}, computation interface{}) (stmt.ComputedQuestion, error) {
	return stmt.ComputedQuestion{label.(lit.StrLit), varDecl.(vari.VarDecl), computation.(expr.Expr)}, nil
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
	return stmt.If{cond.(expr.Expr), body.(stmt.StmtList)}, nil
}

func NewIfElse(cond interface{}, ifBody interface{}, elseBody interface{}) (stmt.IfElse, error) {
	return stmt.IfElse{cond.(expr.Expr), ifBody.(stmt.StmtList), elseBody.(stmt.StmtList)}, nil
}

func NewVarDecl(ident interface{}, typeIdent interface{}) (vari.VarDecl, error) {
	return vari.VarDecl{ident.(vari.VarId), typeIdent.(vari.VarTypeId)}, nil
}

func NewVarId(ident interface{}) (vari.VarId, error) {
	identifierString := string(ident.(*token.Token).Lit)
	return vari.VarId{identifierString}, nil
}
