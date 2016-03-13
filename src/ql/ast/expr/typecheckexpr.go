package expr

import (
	"fmt"
	"ql/interfaces"
)

func (b BinaryOperator) TypeCheck(typeChecker interfaces.TypeChecker, symbolTable interfaces.SymbolTable) {
	typeCheckForUnequalTypes(b, typeChecker, symbolTable)
	typeCheckInvalidBinaryOperatorForOperands(b, typeChecker, symbolTable)

	b.Lhs.TypeCheck(typeChecker, symbolTable)
	b.Rhs.TypeCheck(typeChecker, symbolTable)
}

func typeCheckForUnequalTypes(expr BinaryOperator, typeChecker interfaces.TypeChecker, symbolTable interfaces.SymbolTable) {
	lhsTypeOfValue := fmt.Sprintf("%T", expr.Lhs.Eval(symbolTable))
	rhsTypeOfValue := fmt.Sprintf("%T", expr.Rhs.Eval(symbolTable))

	if lhsTypeOfValue != rhsTypeOfValue {
		typeChecker.AddEncounteredErrorForCheckType("InvalidOperandsDifferentTypes", fmt.Errorf("Encountered BinaryOperator with operands of different types: %s and %s", lhsTypeOfValue, rhsTypeOfValue))
	}
}

func typeCheckInvalidBinaryOperatorForOperands(binaryOperator BinaryOperator, typeChecker interfaces.TypeChecker, symbolTable interfaces.SymbolTable) {
	exprType := fmt.Sprintf("%T", binaryOperator)
	lhsTypeOfValue := binaryOperator.Lhs.Eval(symbolTable)

	switch lhsTypeOfValue.(type) {
	case bool:
		if exprType != "And" && exprType != "Or" && exprType != "Eq" && exprType != "NEq" {
			typeChecker.AddEncounteredErrorForCheckType("InvalidOperationOnOperands", fmt.Errorf("Encountered invalid operation for bool operands"))
		}
	case int:
		if exprType != "Add" && exprType != "Div" && exprType != "Eq" && exprType != "NEq" && exprType != "GEq" && exprType != "GT" && exprType != "GEq" && exprType != "LEq" && exprType != "LT" && exprType != "Mul" && exprType != "Sub" {
			typeChecker.AddEncounteredErrorForCheckType("InvalidOperationOnOperands", fmt.Errorf("Encountered invalid operation for int operands"))
		}
	case string:
		if exprType != "Eq" && exprType != "NEq" && exprType != "GEq" && exprType != "GT" && exprType != "GEq" && exprType != "LEq" && exprType != "LT" {
			typeChecker.AddEncounteredErrorForCheckType("InvalidOperationOnOperands", fmt.Errorf("Encountered invalid operation for string operands"))
		}
	}
}

func (u UnaryOperator) TypeCheck(typeChecker interfaces.TypeChecker, symbolTable interfaces.SymbolTable) {
	typeCheckInvalidUnaryOperatorForOperands(u, typeChecker, symbolTable)

	u.Value.(interfaces.Expr).TypeCheck(typeChecker, symbolTable)
}

func typeCheckInvalidUnaryOperatorForOperands(unaryOperator UnaryOperator, typeChecker interfaces.TypeChecker, symbolTable interfaces.SymbolTable) {
	exprType := fmt.Sprintf("%T", unaryOperator)
	evalValue := unaryOperator.Value.Eval(symbolTable)

	switch evalValue.(type) {
	case bool:
		if exprType != "unaryoperatorNot" {
			typeChecker.AddEncounteredErrorForCheckType("InvalidOperationOnOperands", fmt.Errorf("Encountered invalid operation for bool operand"))
		}
	case int:
		if exprType != "unaryoperatorNeg" && exprType != "unaryoperatorPos" {
			typeChecker.AddEncounteredErrorForCheckType("InvalidOperationOnOperands", fmt.Errorf("Encountered invalid operation for int operand"))
		}
	case string:
		// there are no unary operators on strings, so always return error
		typeChecker.AddEncounteredErrorForCheckType("InvalidOperationOnOperands", fmt.Errorf("Encountered invalid operation for string operand"))
	}
}

func (i IntLit) TypeCheck(typeChecker interfaces.TypeChecker, symbolTable interfaces.SymbolTable) {

}

func (b BoolLit) TypeCheck(typeChecker interfaces.TypeChecker, symbolTable interfaces.SymbolTable) {

}

func (s StrLit) TypeCheck(typeChecker interfaces.TypeChecker, symbolTable interfaces.SymbolTable) {

}

func (v VarExpr) TypeCheck(typeChecker interfaces.TypeChecker, symbolTable interfaces.SymbolTable) {
	typeCheckUndefinedQuestionReference(v, typeChecker, symbolTable)

	v.Identifier.TypeCheck(typeChecker, symbolTable)
}

func typeCheckUndefinedQuestionReference(varExpr VarExpr, typeChecker interfaces.TypeChecker, symbolTable interfaces.SymbolTable) {
	if symbolTable.GetNodeForIdentifier(varExpr.GetIdentifier()) == nil {
		typeChecker.AddEncounteredErrorForCheckType("ReferenceToUndefinedQuestion", fmt.Errorf("Reference to unknown question identifier: %s", varExpr.GetIdentifier()))
	}
}

func (this Expr) TypeCheck(typeChecker interfaces.TypeChecker, symbolTable interfaces.SymbolTable) {
	panic("Expr TypeCheck method not overridden")
}
