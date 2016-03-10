package binaryoperatorexpr

import (
	"fmt"
	//	"ql/ast/expr"
	"ql/interfaces"
)

func (b BinaryOperator) TypeCheck(typeChecker interfaces.TypeChecker, symbolTable interfaces.SymbolTable) {
	typeCheckForUnequalTypes(b, typeChecker, symbolTable)
	typeCheckInvalidOperatorForOperands(b, typeChecker, symbolTable)

	//b.Lhs.(expr.Expr).TypeCheck(typeChecker, symbolTable)
	//b.Rhs.(expr.Expr).TypeCheck(typeChecker, symbolTable)
}

func typeCheckForUnequalTypes(expr BinaryOperator, typeChecker interfaces.TypeChecker, symbolTable interfaces.SymbolTable) {
	lhsTypeOfValue := fmt.Sprintf("%T", expr.Lhs.Eval(symbolTable))
	rhsTypeOfValue := fmt.Sprintf("%T", expr.Rhs.Eval(symbolTable))

	if lhsTypeOfValue != rhsTypeOfValue {
		typeChecker.AddEncounteredErrorForCheckType("InvalidOperands", fmt.Errorf("Encountered BinaryOperator with operands of different types: %s and %s", lhsTypeOfValue, rhsTypeOfValue))
	}
}

func typeCheckInvalidOperatorForOperands(binaryOperatorExpr BinaryOperator, typeChecker interfaces.TypeChecker, symbolTable interfaces.SymbolTable) {
	exprType := fmt.Sprintf("%T", binaryOperatorExpr)
	lhsTypeOfValue := fmt.Sprintf("%T", binaryOperatorExpr.Lhs.Eval(symbolTable))

	if lhsTypeOfValue == "bool" {
		if exprType != "And" && exprType != "Or" && exprType != "Eq" && exprType != "NEq" {
			typeChecker.AddEncounteredErrorForCheckType("InvalidOperands", fmt.Errorf("Encountered invalid operation for bool operands: %s", exprType))
		}
	} else if lhsTypeOfValue == "int" {
		if exprType != "Add" && exprType != "Div" && exprType != "Eq" && exprType != "NEq" && exprType != "GEq" && exprType != "GT" && exprType != "GEq" && exprType != "LEq" && exprType != "LT" && exprType != "Mul" && exprType != "Sub" {
			typeChecker.AddEncounteredErrorForCheckType("InvalidOperands", fmt.Errorf("Encountered invalid operation for int operands: %s", exprType))
		}
	} else if lhsTypeOfValue == "string" {
		if exprType != "Eq" && exprType != "NEq" && exprType != "GEq" && exprType != "GT" && exprType != "GEq" && exprType != "LEq" && exprType != "LT" {
			typeChecker.AddEncounteredErrorForCheckType("InvalidOperands", fmt.Errorf("Encountered invalid operation for string operands: %s", exprType))
		}
	}
}
