package typechecker

import (
	"fmt"
	"ql/ast/expr"
	"ql/ast/stmt"
	"ql/ast/vari"
	"ql/interfaces"
	"ql/symboltable"
	"testing"
)

func testInvalidOperandsCheckerForDifferentOperandEvalTypes(t *testing.T) {
	exampleExpr := expr.NewSubNoSourceInfo(expr.NewBoolLitNoSourceInfo(true), expr.NewIntLitNoSourceInfo(10))

	typeChecker := NewTypeChecker()
	exampleExpr.TypeCheck(&typeChecker, nil)
	errorsReported := typeChecker.GetEncountedErrorsForCheckType("InvalidOperandsDifferentTypes")

	if len(errorsReported) != 1 || fmt.Sprintf("%v", errorsReported[0]) != fmt.Sprintf("%v", fmt.Errorf("Encountered BinaryOperator with operands of different types: bool and int")) {
		t.Errorf("Invalid operands checker did not correctly report operands of different types %v", errorsReported)
	}
}

func testInvalidOperandsCheckerForInvalidBinaryOperationWithBools(t *testing.T) {
	exampleExpr := expr.NewSubNoSourceInfo(expr.NewBoolLitNoSourceInfo(true), expr.NewBoolLitNoSourceInfo(false))

	typeChecker := NewTypeChecker()
	exampleExpr.TypeCheck(&typeChecker, nil)
	errorsReported := typeChecker.GetEncountedErrorsForCheckType("InvalidOperationOnOperands")

	if len(errorsReported) != 1 || fmt.Sprintf("%v", errorsReported[0]) != fmt.Sprintf("%v", fmt.Errorf("Encountered invalid operation for bool operands")) {
		t.Errorf("Invalid operand operation checker did not correctly report invalid binary operation on bool types %v", errorsReported)
	}
}

func testInvalidOperandsCheckerForInvalidBinaryOperationWithIntegers(t *testing.T) {
	exampleExpr := expr.NewAndNoSourceInfo(expr.NewIntLitNoSourceInfo(10), expr.NewIntLitNoSourceInfo(8))

	typeChecker := NewTypeChecker()
	exampleExpr.TypeCheck(&typeChecker, nil)
	errorsReported := typeChecker.GetEncountedErrorsForCheckType("InvalidOperationOnOperands")

	if len(errorsReported) != 1 || fmt.Sprintf("%v", errorsReported[0]) != fmt.Sprintf("%v", fmt.Errorf("Encountered invalid operation for int operands")) {
		t.Errorf("Invalid operand operation checker did not correctly report invalid binary operation on int types")
	}
}

func testInvalidOperandsCheckerForInvalidBinaryOperationWithStrings(t *testing.T) {
	exampleExpr := expr.NewAndNoSourceInfo(expr.NewStrLitNoSourceInfo("Test A"), expr.NewStrLitNoSourceInfo("Test B"))

	typeChecker := NewTypeChecker()
	exampleExpr.TypeCheck(&typeChecker, nil)
	errorsReported := typeChecker.GetEncountedErrorsForCheckType("InvalidOperationOnOperands")

	if len(errorsReported) != 1 || fmt.Sprintf("%v", errorsReported[0]) != fmt.Sprintf("%v", fmt.Errorf("Encountered invalid operation for string operands")) {
		t.Errorf("Invalid operand operation checker did not correctly report invalid binary operation on string types")
	}
}

func testInvalidOperandsCheckerForInvalidUnaryOperationWithBool(t *testing.T) {
	exampleExpr := expr.NewNegNoSourceInfo(expr.NewBoolLitNoSourceInfo(true))

	typeChecker := NewTypeChecker()
	exampleExpr.TypeCheck(&typeChecker, nil)
	errorsReported := typeChecker.GetEncountedErrorsForCheckType("InvalidOperationOnOperands")

	if len(errorsReported) != 1 || fmt.Sprintf("%v", errorsReported[0]) != fmt.Sprintf("%v", fmt.Errorf("Encountered invalid operation for bool operand")) {
		t.Errorf("Invalid operand operation checker did not correctly report invalid unary operation on bool type %v", errorsReported)
	}
}

func testInvalidOperandsCheckerForInvalidUnaryOperationWithInt(t *testing.T) {
	exampleExpr := expr.NewNotNoSourceInfo(expr.NewIntLitNoSourceInfo(3))

	typeChecker := NewTypeChecker()
	exampleExpr.TypeCheck(&typeChecker, nil)
	errorsReported := typeChecker.GetEncountedErrorsForCheckType("InvalidOperationOnOperands")

	if len(errorsReported) != 1 || fmt.Sprintf("%v", errorsReported[0]) != fmt.Sprintf("%v", fmt.Errorf("Encountered invalid operation for int operand")) {
		t.Errorf("Invalid operand operation checker did not correctly report invalid unary operation on int type")
	}
}

func testInvalidOperandsCheckerForInvalidUnaryOperationWithString(t *testing.T) {
	exampleExpr := expr.NewNotNoSourceInfo(expr.NewStrLitNoSourceInfo("Test"))

	typeChecker := NewTypeChecker()
	exampleExpr.TypeCheck(&typeChecker, nil)
	errorsReported := typeChecker.GetEncountedErrorsForCheckType("InvalidOperationOnOperands")

	if len(errorsReported) != 1 || fmt.Sprintf("%v", errorsReported[0]) != fmt.Sprintf("%v", fmt.Errorf("Encountered invalid operation for string operand")) {
		t.Errorf("Invalid operand operation checker did not correctly report invalid unary operation on string type")
	}
}

func testUndefinedQuestionReferenceChecker(t *testing.T) {
	computedQuestion := stmt.NewComputedQuestionNoSourceInfo(expr.NewStrLitNoSourceInfo("Value residue:"), vari.NewVarDeclNoSourceInfo(vari.NewVarIdNoSourceInfo("valueResidue"), vari.NewIntTypeNoSourceInfo()), expr.NewSubNoSourceInfo(expr.NewVarExprNoSourceInfo(vari.NewVarIdNoSourceInfo("hasSoldHouse")), expr.NewVarExprNoSourceInfo(vari.NewVarIdNoSourceInfo("hasMaintLoan"))))
	exampleBody := stmt.NewStmtListNoSourceInfo([]interfaces.Question{computedQuestion}, []interfaces.Conditional{})
	exampleForm := stmt.NewFormNoSourceInfo(vari.NewVarIdNoSourceInfo("TestForm"), exampleBody)

	typeChecker := NewTypeChecker()
	exampleForm.TypeCheck(&typeChecker, symboltable.NewSymbolTable())
	errorsReported := typeChecker.GetEncountedErrorsForCheckType("ReferenceToUndefinedQuestion")

	if len(errorsReported) != 2 || fmt.Sprintf("%v", errorsReported[0]) != fmt.Sprintf("%v", fmt.Errorf("Reference to unknown question identifier: hasSoldHouse")) || fmt.Sprintf("%v", errorsReported[1]) != fmt.Sprintf("%v", fmt.Errorf("Reference to unknown question identifier: hasMaintLoan")) {
		t.Errorf("Undefined questions references not reported correctly by type checker")
	}
}
