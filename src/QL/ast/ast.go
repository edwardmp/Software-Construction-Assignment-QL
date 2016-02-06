package ast

import (
	"QL/ast/expr"
)

/* Convert interface to correct type, either int or bool */
func convertToCorrectType(lhs interface{}, rhs interface{}) (lhsTypeConverted interface{}, rhsTypeConverted interface{}) {
	var lhsTyped interface{}
	var rhsTyped interface{}

	switch lhs.(type) {
	case int64:
		lhsTyped = int(lhs.(int64))
		rhsTyped = int(rhs.(int64))
	case bool:
		lhsTyped = lhs
		rhsTyped = rhs
	}

	return lhsTyped, rhsTyped
}

var (
	TRUE  = bool(true)
	FALSE = bool(false)
)

func NewPos(val interface{}) (expr.Expr, error) {
	return expr.Pos{int(val.(int64))}, nil
}

func NewNeg(value interface{}) (expr.Expr, error) {
	return expr.Neg{int(value.(int64))}, nil
}

func NewNot(value interface{}) (expr.Expr, error) {
	return expr.Not{bool(value.(bool))}, nil
}

func NewMul(lhs interface{}, rhs interface{}) (expr.Expr, error) {
	return expr.Mul{int(lhs.(int64)), int(rhs.(int64))}, nil
}

func NewDiv(lhs interface{}, rhs interface{}) (expr.Expr, error) {
	return expr.Div{int(lhs.(int64)), int(rhs.(int64))}, nil
}

func NewAdd(lhs interface{}, rhs interface{}) (expr.Expr, error) {
	return expr.Add{int(lhs.(int64)), int(rhs.(int64))}, nil
}

func NewSub(lhs interface{}, rhs interface{}) (expr.Expr, error) {
	return expr.Sub{int(lhs.(int64)), int(rhs.(int64))}, nil
}

func NewEq(lhs interface{}, rhs interface{}) (expr.Expr, error) {
	lhs, rhs = convertToCorrectType(lhs, rhs)
	return expr.Eq{lhs, rhs}, nil
}

func NewNEq(lhs interface{}, rhs interface{}) (expr.Expr, error) {
	lhs, rhs = convertToCorrectType(lhs, rhs)
	return expr.NEq{lhs, rhs}, nil
}

func NewGT(lhs interface{}, rhs interface{}) (expr.Expr, error) {
	return expr.GT{int(lhs.(int64)), int(rhs.(int64))}, nil
}

func NewLT(lhs interface{}, rhs interface{}) (expr.Expr, error) {
	return expr.LT{int(lhs.(int64)), int(rhs.(int64))}, nil
}

func NewGEq(lhs interface{}, rhs interface{}) (expr.Expr, error) {
	return expr.GEq{int(lhs.(int64)), int(rhs.(int64))}, nil
}

func NewLEq(lhs interface{}, rhs interface{}) (expr.Expr, error) {
	return expr.LEq{int(lhs.(int64)), int(rhs.(int64))}, nil
}

func NewAnd(lhs interface{}, rhs interface{}) (expr.Expr, error) {
	return expr.And{lhs.(bool), rhs.(bool)}, nil
}

func NewOr(lhs interface{}, rhs interface{}) (expr.Expr, error) {
	return expr.Or{lhs.(bool), rhs.(bool)}, nil
}
