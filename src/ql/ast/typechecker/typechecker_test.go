package typechecker

import (
	"fmt"
	"ql/ast/expr"
	"testing"
)

func TestInvalidOperandsCheckerForDifferentOperandEvalTypes(t *testing.T) {
	exampleExpr := expr.NewSub(expr.BoolLit{true}, expr.IntLit{10})

	typeChecker := NewTypeChecker()
	exampleExpr.TypeCheck(&typeChecker, nil)
	errorsReported := typeChecker.GetEncountedErrorsForCheckType("InvalidOperandsDifferentTypes")

	if len(errorsReported) != 1 || fmt.Sprintf("%v", errorsReported[0]) != fmt.Sprintf("%v", fmt.Errorf("Encountered BinaryOperator with operands of different types: bool and int")) {
		t.Errorf("Invalid operands checker did not correctly report operands of different types %v", errorsReported)
	}
}

func TestInvalidOperandsCheckerForInvalidBinaryOperationWithBools(t *testing.T) {
	exampleExpr := expr.NewSub(expr.BoolLit{true}, expr.BoolLit{false})

	typeChecker := NewTypeChecker()
	exampleExpr.TypeCheck(&typeChecker, nil)
	errorsReported := typeChecker.GetEncountedErrorsForCheckType("InvalidOperationOnOperands")

	if len(errorsReported) != 1 || fmt.Sprintf("%v", errorsReported[0]) != fmt.Sprintf("%v", fmt.Errorf("Encountered invalid operation for bool operands")) {
		t.Errorf("Invalid operand operation checker did not correctly report invalid binary operation on bool types %v", errorsReported)
	}
}

func TestInvalidOperandsCheckerForInvalidBinaryOperationWithIntegers(t *testing.T) {
	exampleExpr := expr.NewAnd(expr.IntLit{10}, expr.IntLit{8})

	typeChecker := NewTypeChecker()
	exampleExpr.TypeCheck(&typeChecker, nil)
	errorsReported := typeChecker.GetEncountedErrorsForCheckType("InvalidOperationOnOperands")

	if len(errorsReported) != 1 || fmt.Sprintf("%v", errorsReported[0]) != fmt.Sprintf("%v", fmt.Errorf("Encountered invalid operation for int operands")) {
		t.Errorf("Invalid operand operation checker did not correctly report invalid binary operation on int types")
	}
}

func TestInvalidOperandsCheckerForInvalidBinaryOperationWithStrings(t *testing.T) {
	exampleExpr := expr.NewAnd(expr.StrLit{"Test A"}, expr.StrLit{"Test B"})

	typeChecker := NewTypeChecker()
	exampleExpr.TypeCheck(&typeChecker, nil)
	errorsReported := typeChecker.GetEncountedErrorsForCheckType("InvalidOperationOnOperands")

	if len(errorsReported) != 1 || fmt.Sprintf("%v", errorsReported[0]) != fmt.Sprintf("%v", fmt.Errorf("Encountered invalid operation for string operands")) {
		t.Errorf("Invalid operand operation checker did not correctly report invalid binary operation on string types")
	}
}

func TestInvalidOperandsCheckerForInvalidUnaryOperationWithBool(t *testing.T) {
	exampleExpr := expr.NewNeg(expr.BoolLit{true})

	typeChecker := NewTypeChecker()
	exampleExpr.TypeCheck(&typeChecker, nil)
	errorsReported := typeChecker.GetEncountedErrorsForCheckType("InvalidOperationOnOperands")

	if len(errorsReported) != 1 || fmt.Sprintf("%v", errorsReported[0]) != fmt.Sprintf("%v", fmt.Errorf("Encountered invalid operation for bool operand")) {
		t.Errorf("Invalid operand operation checker did not correctly report invalid unary operation on bool type %v", errorsReported)
	}
}

func TestInvalidOperandsCheckerForInvalidUnaryOperationWithInt(t *testing.T) {
	exampleExpr := expr.NewNot(expr.IntLit{3})

	typeChecker := NewTypeChecker()
	exampleExpr.TypeCheck(&typeChecker, nil)
	errorsReported := typeChecker.GetEncountedErrorsForCheckType("InvalidOperationOnOperands")

	if len(errorsReported) != 1 || fmt.Sprintf("%v", errorsReported[0]) != fmt.Sprintf("%v", fmt.Errorf("Encountered invalid operation for int operand")) {
		t.Errorf("Invalid operand operation checker did not correctly report invalid unary operation on int type")
	}
}

func TestInvalidOperandsCheckerForInvalidUnaryOperationWithString(t *testing.T) {
	exampleExpr := expr.NewNot(expr.StrLit{"Test"})

	typeChecker := NewTypeChecker()
	exampleExpr.TypeCheck(&typeChecker, nil)
	errorsReported := typeChecker.GetEncountedErrorsForCheckType("InvalidOperationOnOperands")

	if len(errorsReported) != 1 || fmt.Sprintf("%v", errorsReported[0]) != fmt.Sprintf("%v", fmt.Errorf("Encountered invalid operation for string operand")) {
		t.Errorf("Invalid operand operation checker did not correctly report invalid unary operation on string type")
	}
}
