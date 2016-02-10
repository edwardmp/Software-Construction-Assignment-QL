package ast

import (
	"QL/ast/expr"
)

var (
	TRUE  = bool(true)
	FALSE = bool(false)
)

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
