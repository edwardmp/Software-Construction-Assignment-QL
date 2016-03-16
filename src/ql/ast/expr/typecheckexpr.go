package expr

import (
	"fmt"
	"ql/interfaces"
)

func (this VarExpr) TypeCheck(typeChecker interfaces.TypeChecker, symbols interfaces.TypeCheckSymbols) interfaces.ValueType {
	checkUndefinedQuestionReference(this, typeChecker, symbols)

	typeChecker.AddDependencyForCurrentlyVisitedVarDecl(this.Identifier)

	if symbols.IsTypeSetForVarId(this.Identifier) {
		return symbols.GetTypeForVarId(this.Identifier).(interfaces.ValueType)
	}

	return nil
}

func (this Add) TypeCheck(typeChecker interfaces.TypeChecker, s interfaces.TypeCheckSymbols) interfaces.ValueType {
	checkOperands(this, NewIntTypeNoSourceInfo(), typeChecker, s)

	return NewIntTypeNoSourceInfo()
}

func (this And) TypeCheck(typeChecker interfaces.TypeChecker, s interfaces.TypeCheckSymbols) interfaces.ValueType {
	checkOperands(this, NewBoolTypeNoSourceInfo(), typeChecker, s)

	return NewBoolTypeNoSourceInfo()
}

func (this Div) TypeCheck(typeChecker interfaces.TypeChecker, s interfaces.TypeCheckSymbols) interfaces.ValueType {
	checkOperands(this, NewIntTypeNoSourceInfo(), typeChecker, s)

	return NewIntTypeNoSourceInfo()
}

func (this Eq) TypeCheck(typeChecker interfaces.TypeChecker, s interfaces.TypeCheckSymbols) interfaces.ValueType {
	checkOperands(this, NewBoolTypeNoSourceInfo(), typeChecker, s)

	return NewBoolTypeNoSourceInfo()
}

func (this GEq) TypeCheck(typeChecker interfaces.TypeChecker, s interfaces.TypeCheckSymbols) interfaces.ValueType {
	checkOperands(this, NewBoolTypeNoSourceInfo(), typeChecker, s)

	return NewBoolTypeNoSourceInfo()
}

func (this GT) TypeCheck(typeChecker interfaces.TypeChecker, s interfaces.TypeCheckSymbols) interfaces.ValueType {
	checkOperands(this, NewBoolTypeNoSourceInfo(), typeChecker, s)

	return NewBoolTypeNoSourceInfo()
}

func (this LEq) TypeCheck(typeChecker interfaces.TypeChecker, s interfaces.TypeCheckSymbols) interfaces.ValueType {
	checkOperands(this, NewBoolTypeNoSourceInfo(), typeChecker, s)

	return NewBoolTypeNoSourceInfo()
}

func (this LT) TypeCheck(typeChecker interfaces.TypeChecker, s interfaces.TypeCheckSymbols) interfaces.ValueType {
	checkOperands(this, NewBoolTypeNoSourceInfo(), typeChecker, s)

	return NewBoolTypeNoSourceInfo()
}

func (this Mul) TypeCheck(typeChecker interfaces.TypeChecker, s interfaces.TypeCheckSymbols) interfaces.ValueType {
	checkOperands(this, NewIntTypeNoSourceInfo(), typeChecker, s)

	return NewIntTypeNoSourceInfo()
}

func (this Neg) TypeCheck(typeChecker interfaces.TypeChecker, s interfaces.TypeCheckSymbols) interfaces.ValueType {
	checkOperand(this, NewIntTypeNoSourceInfo(), typeChecker, s)

	return NewIntTypeNoSourceInfo()
}

func (this NEq) TypeCheck(typeChecker interfaces.TypeChecker, s interfaces.TypeCheckSymbols) interfaces.ValueType {
	checkOperands(this, NewBoolTypeNoSourceInfo(), typeChecker, s)

	return NewBoolTypeNoSourceInfo()
}

func (this Not) TypeCheck(typeChecker interfaces.TypeChecker, s interfaces.TypeCheckSymbols) interfaces.ValueType {
	checkOperand(this, NewBoolTypeNoSourceInfo(), typeChecker, s)

	return NewBoolTypeNoSourceInfo()
}

func (this Or) TypeCheck(typeChecker interfaces.TypeChecker, s interfaces.TypeCheckSymbols) interfaces.ValueType {
	checkOperands(this, NewBoolTypeNoSourceInfo(), typeChecker, s)

	return NewBoolTypeNoSourceInfo()
}

func (this Pos) TypeCheck(typeChecker interfaces.TypeChecker, s interfaces.TypeCheckSymbols) interfaces.ValueType {
	checkOperand(this, NewIntTypeNoSourceInfo(), typeChecker, s)

	return NewIntTypeNoSourceInfo()
}

func (this Sub) TypeCheck(typeChecker interfaces.TypeChecker, s interfaces.TypeCheckSymbols) interfaces.ValueType {
	checkOperands(this, NewIntTypeNoSourceInfo(), typeChecker, s)

	return NewIntTypeNoSourceInfo()
}

func (this IntLit) TypeCheck(typeChecker interfaces.TypeChecker, symbols interfaces.TypeCheckSymbols) interfaces.ValueType {
	return NewIntTypeNoSourceInfo()
}

func (this BoolLit) TypeCheck(typeChecker interfaces.TypeChecker, symbols interfaces.TypeCheckSymbols) interfaces.ValueType {
	return NewBoolTypeNoSourceInfo()
}

func (this StrLit) TypeCheck(typeChecker interfaces.TypeChecker, symbols interfaces.TypeCheckSymbols) interfaces.ValueType {
	return NewStringTypeNoSourceInfo()
}

func (this Expr) TypeCheck(typeChecker interfaces.TypeChecker, symbols interfaces.TypeCheckSymbols) interfaces.ValueType {
	panic("Expr TypeCheck method not overridden")
}

func checkOperand(unaryExpr interfaces.UnaryOperatorExpr, expectedType interfaces.ValueType, typeChecker interfaces.TypeChecker, s interfaces.TypeCheckSymbols) {
	checkForInvalidOperationOperand(unaryExpr.GetValue(), expectedType, typeChecker, s)
}

func checkOperands(binaryExpression interfaces.BinaryOperatorExpr, expectedType interfaces.ValueType, typeChecker interfaces.TypeChecker, s interfaces.TypeCheckSymbols) {
	lhsType := binaryExpression.GetLhs().TypeCheck(typeChecker, s)
	rhsType := binaryExpression.GetRhs().TypeCheck(typeChecker, s)

	typesEqual := checkForUnequalTypes(lhsType, rhsType, typeChecker)

	if typesEqual {
		checkForInvalidOperationOperand(binaryExpression.GetLhs(), expectedType, typeChecker, s)
		checkForInvalidOperationOperand(binaryExpression.GetRhs(), expectedType, typeChecker, s)
	}
}

func checkForUnequalTypes(lhsType, rhsType interfaces.ValueType, typeChecker interfaces.TypeChecker) bool {
	if lhsType != rhsType {
		typeChecker.AddEncounteredErrorForCheckType("InvalidOperandsDifferentTypes", fmt.Errorf("Encountered BinaryOperator with operands of different types: %s and %s", lhsType, rhsType))
		return false
	}

	return true
}

func checkForInvalidOperationOperand(expr interfaces.Expr, expectedType interfaces.ValueType, typeChecker interfaces.TypeChecker, symbols interfaces.TypeCheckSymbols) {
	actualType := expr.TypeCheck(typeChecker, symbols)

	if actualType != expectedType {
		typeChecker.AddEncounteredErrorForCheckType("InvalidOperandForOperation", fmt.Errorf("Encountered invalid operand type for operator, expected type: %s, actual type: %s", expectedType, actualType))
	}
}

func checkUndefinedQuestionReference(varExpr VarExpr, typeChecker interfaces.TypeChecker, symbols interfaces.TypeCheckSymbols) {
	if !symbols.IsTypeSetForVarId(varExpr.GetIdentifier()) {
		typeChecker.AddEncounteredErrorForCheckType("ReferenceToUndefinedQuestion", fmt.Errorf("Reference to unknown question identifier: %s", varExpr.GetIdentifier()))
	}
}
