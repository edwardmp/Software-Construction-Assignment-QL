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

func TestInvalidOperandsCheckerForDifferentOperandEvalTypes(t *testing.T) {
	exampleExpr := expr.NewSubNoSourceInfo(expr.NewBoolLitNoSourceInfo(true), expr.NewIntLitNoSourceInfo(10))

	typeChecker := NewTypeChecker()
	exampleExpr.TypeCheck(&typeChecker, nil)
	errorsReported := typeChecker.GetEncountedErrorsForCheckType("InvalidOperandsDifferentTypes")

	if len(errorsReported) != 1 || fmt.Sprintf("%v", errorsReported[0]) != fmt.Sprintf("%v", fmt.Errorf("Encountered BinaryOperator with operands of different types: bool and int")) {
		t.Errorf("Invalid operands checker did not correctly report operands of different types %v", errorsReported)
	}
}

func TestInvalidOperandsCheckerForInvalidBinaryOperationWithBools(t *testing.T) {
	exampleExpr := expr.NewSubNoSourceInfo(expr.NewBoolLitNoSourceInfo(true), expr.NewBoolLitNoSourceInfo(false))

	typeChecker := NewTypeChecker()
	exampleExpr.TypeCheck(&typeChecker, nil)
	errorsReported := typeChecker.GetEncountedErrorsForCheckType("InvalidOperationOnOperands")

	if len(errorsReported) != 1 || fmt.Sprintf("%v", errorsReported[0]) != fmt.Sprintf("%v", fmt.Errorf("Encountered invalid operation for bool operands")) {
		t.Errorf("Invalid operand operation checker did not correctly report invalid binary operation on bool types %v", errorsReported)
	}
}

func TestInvalidOperandsCheckerForInvalidBinaryOperationWithIntegers(t *testing.T) {
	exampleExpr := expr.NewAndNoSourceInfo(expr.NewIntLitNoSourceInfo(10), expr.NewIntLitNoSourceInfo(8))

	typeChecker := NewTypeChecker()
	exampleExpr.TypeCheck(&typeChecker, nil)
	errorsReported := typeChecker.GetEncountedErrorsForCheckType("InvalidOperationOnOperands")

	if len(errorsReported) != 1 || fmt.Sprintf("%v", errorsReported[0]) != fmt.Sprintf("%v", fmt.Errorf("Encountered invalid operation for int operands")) {
		t.Errorf("Invalid operand operation checker did not correctly report invalid binary operation on int types")
	}
}

func TestInvalidOperandsCheckerForInvalidBinaryOperationWithStrings(t *testing.T) {
	exampleExpr := expr.NewAndNoSourceInfo(expr.NewStrLitNoSourceInfo("Test A"), expr.NewStrLitNoSourceInfo("Test B"))

	typeChecker := NewTypeChecker()
	exampleExpr.TypeCheck(&typeChecker, nil)
	errorsReported := typeChecker.GetEncountedErrorsForCheckType("InvalidOperationOnOperands")

	if len(errorsReported) != 1 || fmt.Sprintf("%v", errorsReported[0]) != fmt.Sprintf("%v", fmt.Errorf("Encountered invalid operation for string operands")) {
		t.Errorf("Invalid operand operation checker did not correctly report invalid binary operation on string types")
	}
}

func TestInvalidOperandsCheckerForInvalidUnaryOperationWithBool(t *testing.T) {
	exampleExpr := expr.NewNegNoSourceInfo(expr.NewBoolLitNoSourceInfo(true))

	typeChecker := NewTypeChecker()
	exampleExpr.TypeCheck(&typeChecker, nil)
	errorsReported := typeChecker.GetEncountedErrorsForCheckType("InvalidOperationOnOperands")

	if len(errorsReported) != 1 || fmt.Sprintf("%v", errorsReported[0]) != fmt.Sprintf("%v", fmt.Errorf("Encountered invalid operation for bool operand")) {
		t.Errorf("Invalid operand operation checker did not correctly report invalid unary operation on bool type %v", errorsReported)
	}
}

func TestInvalidOperandsCheckerForInvalidUnaryOperationWithInt(t *testing.T) {
	exampleExpr := expr.NewNotNoSourceInfo(expr.NewIntLitNoSourceInfo(3))

	typeChecker := NewTypeChecker()
	exampleExpr.TypeCheck(&typeChecker, nil)
	errorsReported := typeChecker.GetEncountedErrorsForCheckType("InvalidOperationOnOperands")

	if len(errorsReported) != 1 || fmt.Sprintf("%v", errorsReported[0]) != fmt.Sprintf("%v", fmt.Errorf("Encountered invalid operation for int operand")) {
		t.Errorf("Invalid operand operation checker did not correctly report invalid unary operation on int type")
	}
}

func TestInvalidOperandsCheckerForInvalidUnaryOperationWithString(t *testing.T) {
	exampleExpr := expr.NewNotNoSourceInfo(expr.NewStrLitNoSourceInfo("Test"))

	typeChecker := NewTypeChecker()
	exampleExpr.TypeCheck(&typeChecker, nil)
	errorsReported := typeChecker.GetEncountedErrorsForCheckType("InvalidOperationOnOperands")

	if len(errorsReported) != 1 || fmt.Sprintf("%v", errorsReported[0]) != fmt.Sprintf("%v", fmt.Errorf("Encountered invalid operation for string operand")) {
		t.Errorf("Invalid operand operation checker did not correctly report invalid unary operation on string type")
	}
}

func TestUndefinedQuestionReferenceChecker(t *testing.T) {
	computedQuestion := stmt.NewComputedQuestionNoSourceInfo(expr.NewStrLitNoSourceInfo("Value residue:"), vari.NewVarDeclNoSourceInfo(vari.NewVarIdNoSourceInfo("valueResidue"), vari.NewIntTypeNoSourceInfo()), expr.NewSubNoSourceInfo(expr.NewVarExprNoSourceInfo(vari.NewVarIdNoSourceInfo("hasSoldHouse")), expr.NewVarExprNoSourceInfo(vari.NewVarIdNoSourceInfo("hasMaintLoan"))))
	exampleBody := stmt.NewStmtListNoSourceInfo([]interfaces.Question{computedQuestion}, []interfaces.Conditional{})
	exampleForm := stmt.NewFormNoSourceInfo(vari.NewVarIdNoSourceInfo("TestForm"), exampleBody)

	typeChecker := NewTypeChecker()
	exampleForm.TypeCheck(&typeChecker, symboltable.NewSymbols())
	errorsReported := typeChecker.GetEncountedErrorsForCheckType("ReferenceToUndefinedQuestion")

	if len(errorsReported) != 2 || fmt.Sprintf("%v", errorsReported[0]) != fmt.Sprintf("%v", fmt.Errorf("Reference to unknown question identifier: hasSoldHouse")) || fmt.Sprintf("%v", errorsReported[1]) != fmt.Sprintf("%v", fmt.Errorf("Reference to unknown question identifier: hasMaintLoan")) {
		t.Errorf("Undefined questions references not reported correctly by type checker")
	}
}

func TestNonBoolConditionalChecker(t *testing.T) {
	exampleQuestion := stmt.NewInputQuestionNoSourceInfo(expr.NewStrLitNoSourceInfo("Did you sell a house in 2010?"), vari.NewVarDeclNoSourceInfo(vari.NewVarIdNoSourceInfo("hasSoldHouse"), vari.NewBoolTypeNoSourceInfo()))
	exampleIf := stmt.NewIfNoSourceInfo(expr.NewIntLitNoSourceInfo(10), stmt.NewStmtListNoSourceInfo([]interfaces.Question{exampleQuestion}, []interfaces.Conditional{}))
	exampleBody := stmt.NewStmtListNoSourceInfo([]interfaces.Question{}, []interfaces.Conditional{exampleIf})
	exampleForm := stmt.NewFormNoSourceInfo(vari.NewVarIdNoSourceInfo("TestForm"), exampleBody)

	typeChecker := NewTypeChecker()
	exampleForm.TypeCheck(&typeChecker, symboltable.NewSymbols())
	errorsReported := typeChecker.GetEncountedErrorsForCheckType("NonBoolConditionals")

	if len(errorsReported) != 1 || fmt.Sprintf("%v", errorsReported[0]) != fmt.Sprintf("%v", fmt.Errorf("Non-boolean type used as condition: int")) {
		t.Errorf("Non bool condition type checker did not correctly report condition of invalid type %v", errorsReported)
	}
}

func TestDuplicateLabelChecker(t *testing.T) {
	firstQuestion := stmt.NewInputQuestionNoSourceInfo(expr.NewStrLitNoSourceInfo("Did you sell a house in 2010?"), vari.NewVarDeclNoSourceInfo(vari.NewVarIdNoSourceInfo("hasSoldHouse"), vari.NewBoolTypeNoSourceInfo()))
	secondQuestion := stmt.NewInputQuestionNoSourceInfo(expr.NewStrLitNoSourceInfo("Did you sell a house in 2010?"), vari.NewVarDeclNoSourceInfo(vari.NewVarIdNoSourceInfo("hasMaintLoan"), vari.NewBoolTypeNoSourceInfo()))
	exampleBody := stmt.NewStmtListNoSourceInfo([]interfaces.Question{firstQuestion, secondQuestion}, []interfaces.Conditional{})
	exampleForm := stmt.NewFormNoSourceInfo(vari.NewVarIdNoSourceInfo("TestForm"), exampleBody)

	typeChecker := NewTypeChecker()
	exampleForm.TypeCheck(&typeChecker, symboltable.NewSymbols())
	warningsReported := typeChecker.GetEncountedErrorsForCheckType("DuplicateLabels")

	if len(warningsReported) != 1 || fmt.Sprintf("%v", warningsReported[0]) != fmt.Sprintf("%v", fmt.Errorf("Label \"Did you sell a house in 2010?\" already used for question with identifier hasSoldHouse, using again for question with identifier hasMaintLoan")) {
		t.Errorf("Duplicate label not reported correctly by type checker")
	}
}

func TestDuplicateVarDeclChecker(t *testing.T) {
	firstQuestion := stmt.NewInputQuestionNoSourceInfo(expr.NewStrLitNoSourceInfo("Did you sell a house in 2010?"), vari.NewVarDeclNoSourceInfo(vari.NewVarIdNoSourceInfo("hasSoldHouse"), vari.NewBoolTypeNoSourceInfo()))
	secondQuestion := stmt.NewInputQuestionNoSourceInfo(expr.NewStrLitNoSourceInfo("Did you sell a house in 2010?"), vari.NewVarDeclNoSourceInfo(vari.NewVarIdNoSourceInfo("hasSoldHouse"), vari.NewIntTypeNoSourceInfo()))
	exampleBody := stmt.NewStmtListNoSourceInfo([]interfaces.Question{firstQuestion, secondQuestion}, []interfaces.Conditional{})
	exampleForm := stmt.NewFormNoSourceInfo(vari.NewVarIdNoSourceInfo("TestForm"), exampleBody)

	typeChecker := NewTypeChecker()
	exampleForm.TypeCheck(&typeChecker, symboltable.NewSymbols())
	errorsReported := typeChecker.GetEncountedErrorsForCheckType("DuplicateVarDeclarations")

	if len(errorsReported) != 1 || fmt.Sprintf("%v", errorsReported[0]) != fmt.Sprintf("%v", fmt.Errorf("Question redeclared with different types: vari.IntType and vari.BoolType")) {
		t.Errorf("Duplicate var decl not reported correctly by type checker")
	}
}
