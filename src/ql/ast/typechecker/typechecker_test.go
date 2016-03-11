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
	exampleExpr := expr.NewSub(expr.BoolLit{true}, expr.IntLit{10})

	typeChecker := NewTypeChecker()
	exampleExpr.TypeCheck(&typeChecker, nil)
	errorsReported := typeChecker.GetEncountedErrorsForCheckType("InvalidOperandsDifferentTypes")

	if len(errorsReported) != 1 || fmt.Sprintf("%v", errorsReported[0]) != fmt.Sprintf("%v", fmt.Errorf("Encountered BinaryOperator with operands of different types: bool and int")) {
		t.Errorf("Invalid operands checker did not correctly report operands of different types %v", errorsReported)
	}
}

func testInvalidOperandsCheckerForInvalidBinaryOperationWithBools(t *testing.T) {
	exampleExpr := expr.NewSub(expr.BoolLit{true}, expr.BoolLit{false})

	typeChecker := NewTypeChecker()
	exampleExpr.TypeCheck(&typeChecker, nil)
	errorsReported := typeChecker.GetEncountedErrorsForCheckType("InvalidOperationOnOperands")

	if len(errorsReported) != 1 || fmt.Sprintf("%v", errorsReported[0]) != fmt.Sprintf("%v", fmt.Errorf("Encountered invalid operation for bool operands")) {
		t.Errorf("Invalid operand operation checker did not correctly report invalid binary operation on bool types %v", errorsReported)
	}
}

func testInvalidOperandsCheckerForInvalidBinaryOperationWithIntegers(t *testing.T) {
	exampleExpr := expr.NewAnd(expr.IntLit{10}, expr.IntLit{8})

	typeChecker := NewTypeChecker()
	exampleExpr.TypeCheck(&typeChecker, nil)
	errorsReported := typeChecker.GetEncountedErrorsForCheckType("InvalidOperationOnOperands")

	if len(errorsReported) != 1 || fmt.Sprintf("%v", errorsReported[0]) != fmt.Sprintf("%v", fmt.Errorf("Encountered invalid operation for int operands")) {
		t.Errorf("Invalid operand operation checker did not correctly report invalid binary operation on int types")
	}
}

func testInvalidOperandsCheckerForInvalidBinaryOperationWithStrings(t *testing.T) {
	exampleExpr := expr.NewAnd(expr.StrLit{"Test A"}, expr.StrLit{"Test B"})

	typeChecker := NewTypeChecker()
	exampleExpr.TypeCheck(&typeChecker, nil)
	errorsReported := typeChecker.GetEncountedErrorsForCheckType("InvalidOperationOnOperands")

	if len(errorsReported) != 1 || fmt.Sprintf("%v", errorsReported[0]) != fmt.Sprintf("%v", fmt.Errorf("Encountered invalid operation for string operands")) {
		t.Errorf("Invalid operand operation checker did not correctly report invalid binary operation on string types")
	}
}

func testInvalidOperandsCheckerForInvalidUnaryOperationWithBool(t *testing.T) {
	exampleExpr := expr.NewNeg(expr.BoolLit{true})

	typeChecker := NewTypeChecker()
	exampleExpr.TypeCheck(&typeChecker, nil)
	errorsReported := typeChecker.GetEncountedErrorsForCheckType("InvalidOperationOnOperands")

	if len(errorsReported) != 1 || fmt.Sprintf("%v", errorsReported[0]) != fmt.Sprintf("%v", fmt.Errorf("Encountered invalid operation for bool operand")) {
		t.Errorf("Invalid operand operation checker did not correctly report invalid unary operation on bool type %v", errorsReported)
	}
}

func testInvalidOperandsCheckerForInvalidUnaryOperationWithInt(t *testing.T) {
	exampleExpr := expr.NewNot(expr.IntLit{3})

	typeChecker := NewTypeChecker()
	exampleExpr.TypeCheck(&typeChecker, nil)
	errorsReported := typeChecker.GetEncountedErrorsForCheckType("InvalidOperationOnOperands")

	if len(errorsReported) != 1 || fmt.Sprintf("%v", errorsReported[0]) != fmt.Sprintf("%v", fmt.Errorf("Encountered invalid operation for int operand")) {
		t.Errorf("Invalid operand operation checker did not correctly report invalid unary operation on int type")
	}
}

func testInvalidOperandsCheckerForInvalidUnaryOperationWithString(t *testing.T) {
	exampleExpr := expr.NewNot(expr.StrLit{"Test"})

	typeChecker := NewTypeChecker()
	exampleExpr.TypeCheck(&typeChecker, nil)
	errorsReported := typeChecker.GetEncountedErrorsForCheckType("InvalidOperationOnOperands")

	if len(errorsReported) != 1 || fmt.Sprintf("%v", errorsReported[0]) != fmt.Sprintf("%v", fmt.Errorf("Encountered invalid operation for string operand")) {
		t.Errorf("Invalid operand operation checker did not correctly report invalid unary operation on string type")
	}
}

func testUndefinedQuestionReferenceChecker(t *testing.T) {
	computedQuestion := stmt.ComputedQuestion{expr.NewStrLit("Value residue:"), vari.VarDecl{vari.VarId{"valueResidue"}, vari.IntType{}}, expr.NewSub(expr.VarExpr{vari.VarId{"hasSoldHouse"}}, expr.VarExpr{vari.VarId{"hasMaintLoan"}})}
	exampleBody := stmt.StmtList{[]interfaces.Question{computedQuestion}, []interfaces.Conditional{}}
	exampleForm := stmt.Form{vari.VarId{"TestForm"}, exampleBody}

	typeChecker := NewTypeChecker()
	exampleForm.TypeCheck(&typeChecker, symboltable.NewSymbolTable())
	errorsReported := typeChecker.GetEncountedErrorsForCheckType("ReferenceToUndefinedQuestion")

	if len(errorsReported) != 2 || fmt.Sprintf("%v", errorsReported[0]) != fmt.Sprintf("%v", fmt.Errorf("Reference to unknown question identifier: hasSoldHouse")) || fmt.Sprintf("%v", errorsReported[1]) != fmt.Sprintf("%v", fmt.Errorf("Reference to unknown question identifier: hasMaintLoan")) {
		t.Errorf("Undefined questions references not reported correctly by type checker")
	}
}
