package expr

import (
	"fmt"
	"ql/interfaces"
)

func (this VarExpr) TypeCheck(typeChecker interfaces.TypeChecker, symbols interfaces.TypeCheckSymbols) interfaces.ValueType {
	typeChecker.AddDependencyForCurrentlyVisitedVarDecl(this.Identifier)

	// return the true type of the VarExpr; the type of the expr referred to
	if symbols.IsTypeSetForVarId(this.Identifier) {
		return symbols.GetTypeForVarId(this.Identifier).(interfaces.ValueType)
	}

	// We don't already mark it as an error; because there is only one scope, the VarDecl may be simply declared later on
	// After the whole Form is typechecked, it is checked which VarIds remain unknown (those that were not declared at a later point)
	typeChecker.MarkVarIdAsUnknown(this.GetIdentifier())

	// no type info in symboltable (reference to undefined question)
	return nil
}

func (this Add) TypeCheck(typeChecker interfaces.TypeChecker, s interfaces.TypeCheckSymbols) interfaces.ValueType {
	this.checkOperands(NewIntTypeNoSourceInfo(), typeChecker, s)

	return NewIntTypeNoSourceInfo()
}

func (this And) TypeCheck(typeChecker interfaces.TypeChecker, s interfaces.TypeCheckSymbols) interfaces.ValueType {
	this.checkOperands(NewBoolTypeNoSourceInfo(), typeChecker, s)

	return NewBoolTypeNoSourceInfo()
}

func (this Div) TypeCheck(typeChecker interfaces.TypeChecker, s interfaces.TypeCheckSymbols) interfaces.ValueType {
	this.checkOperands(NewIntTypeNoSourceInfo(), typeChecker, s)

	return NewIntTypeNoSourceInfo()
}

func (this Eq) TypeCheck(typeChecker interfaces.TypeChecker, s interfaces.TypeCheckSymbols) interfaces.ValueType {
	checkForUnequalTypes(this, NewBoolTypeNoSourceInfo(), typeChecker, s)

	return NewBoolTypeNoSourceInfo()
}

func (this GEq) TypeCheck(typeChecker interfaces.TypeChecker, s interfaces.TypeCheckSymbols) interfaces.ValueType {
	this.checkOperands(NewBoolTypeNoSourceInfo(), typeChecker, s)

	return NewBoolTypeNoSourceInfo()
}

func (this GT) TypeCheck(typeChecker interfaces.TypeChecker, s interfaces.TypeCheckSymbols) interfaces.ValueType {
	this.checkOperands(NewIntTypeNoSourceInfo(), typeChecker, s)

	return NewBoolTypeNoSourceInfo()
}

func (this LEq) TypeCheck(typeChecker interfaces.TypeChecker, s interfaces.TypeCheckSymbols) interfaces.ValueType {
	this.checkOperands(NewIntTypeNoSourceInfo(), typeChecker, s)

	return NewBoolTypeNoSourceInfo()
}

func (this LT) TypeCheck(typeChecker interfaces.TypeChecker, s interfaces.TypeCheckSymbols) interfaces.ValueType {
	this.checkOperands(NewIntTypeNoSourceInfo(), typeChecker, s)

	return NewBoolTypeNoSourceInfo()
}

func (this Mul) TypeCheck(typeChecker interfaces.TypeChecker, s interfaces.TypeCheckSymbols) interfaces.ValueType {
	this.checkOperands(NewIntTypeNoSourceInfo(), typeChecker, s)

	return NewIntTypeNoSourceInfo()
}

func (this Neg) TypeCheck(typeChecker interfaces.TypeChecker, s interfaces.TypeCheckSymbols) interfaces.ValueType {
	this.checkOperand(NewIntTypeNoSourceInfo(), typeChecker, s)

	return NewIntTypeNoSourceInfo()
}

func (this NEq) TypeCheck(typeChecker interfaces.TypeChecker, s interfaces.TypeCheckSymbols) interfaces.ValueType {
	checkForUnequalTypes(this, NewBoolTypeNoSourceInfo(), typeChecker, s)

	return NewBoolTypeNoSourceInfo()
}

func (this Not) TypeCheck(typeChecker interfaces.TypeChecker, s interfaces.TypeCheckSymbols) interfaces.ValueType {
	this.checkOperand(NewBoolTypeNoSourceInfo(), typeChecker, s)

	return NewBoolTypeNoSourceInfo()
}

func (this Or) TypeCheck(typeChecker interfaces.TypeChecker, s interfaces.TypeCheckSymbols) interfaces.ValueType {
	this.checkOperands(NewBoolTypeNoSourceInfo(), typeChecker, s)

	return NewBoolTypeNoSourceInfo()
}

func (this Pos) TypeCheck(typeChecker interfaces.TypeChecker, s interfaces.TypeCheckSymbols) interfaces.ValueType {
	this.checkOperand(NewIntTypeNoSourceInfo(), typeChecker, s)

	return NewIntTypeNoSourceInfo()
}

func (this Sub) TypeCheck(typeChecker interfaces.TypeChecker, s interfaces.TypeCheckSymbols) interfaces.ValueType {
	this.checkOperands(NewIntTypeNoSourceInfo(), typeChecker, s)

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

// checkOperand checks if the value of a unaryExpr is of the expected type
func (unaryExpr UnaryOperator) checkOperand(expectedType interfaces.ValueType, typeChecker interfaces.TypeChecker, s interfaces.TypeCheckSymbols) {
	checkForInvalidOperandForOperator(unaryExpr.GetValue(), expectedType, typeChecker, s)
}

// checkOperands checks if the left-hand and right-hand sides are of the expected type
func (binaryExpression BinaryOperator) checkOperands(expectedType interfaces.ValueType, typeChecker interfaces.TypeChecker, s interfaces.TypeCheckSymbols) {
	checkForInvalidOperandForOperator(binaryExpression.GetLhs(), expectedType, typeChecker, s)
	checkForInvalidOperandForOperator(binaryExpression.GetRhs(), expectedType, typeChecker, s)
}

/* TODO, is this a better approach
func (this Eq) checkForUnequalTypes() expectedType interfaces.ValueType, typeChecker interfaces.TypeChecker, s interfaces.TypeCheckSymbols) {

}

Because
func (this BinaryOperator) checkForUnequalTypes() will be callable from other operators than Eq & NEq
*/

func checkForUnequalTypes(binaryExpression interfaces.BinaryOperatorExpr, expectedType interfaces.ValueType, typeChecker interfaces.TypeChecker, s interfaces.TypeCheckSymbols) {
	lhsType := binaryExpression.GetLhs().TypeCheck(typeChecker, s)
	rhsType := binaryExpression.GetRhs().TypeCheck(typeChecker, s)

	// this occurs when we have no type info (e.g. VarExpr with reference to undefined question)
	// this case is already handled by the undefined question reference checker, so don't continue here
	if lhsType == nil || rhsType == nil {
		return
	}

	if lhsType != rhsType {
		typeChecker.AddEncounteredError(fmt.Errorf("Encountered BinaryOperator with operands of different types: %s and %s", lhsType, rhsType))
	}
}

func checkForInvalidOperandForOperator(expr interfaces.Expr, expectedType interfaces.ValueType, typeChecker interfaces.TypeChecker, symbols interfaces.TypeCheckSymbols) {
	actualType := expr.TypeCheck(typeChecker, symbols)

	// this occurs when we have no type info (e.g. VarExpr with reference to undefined question)
	// this case is already handled by the undefined question reference checker, so don't continue here
	if actualType == nil {
		return
	}

	if actualType != expectedType {
		typeChecker.AddEncounteredError(fmt.Errorf("Encountered invalid operand type for operator, expected type: %s, actual type: %s", expectedType, actualType))
	}
}
