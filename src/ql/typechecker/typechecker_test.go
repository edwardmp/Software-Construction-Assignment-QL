package typechecker

import (
	"fmt"
	"ql/ast/expr"
	"ql/ast/stmt"
	"ql/ast/vari"
	"ql/interfaces"
	"ql/symbols"
	"testing"
)

func TestInvalidOperandsCheckerForDifferentOperandEvalTypes(t *testing.T) {
	exampleExpr := expr.NewSub(expr.NewBoolLit(true), expr.NewIntLit(10))

	typeChecker := NewTypeChecker()
	exampleExpr.TypeCheck(typeChecker, nil)
	errorsReported := typeChecker.ErrorsEncountered

	if len(errorsReported) != 1 || fmt.Sprintf("%v", errorsReported[0]) != fmt.Sprintf("%v", fmt.Errorf("Encountered unexpected operand type, expected type: Integer, actual type: Boolean")) {
		t.Errorf("Invalid operands checker did not correctly report operands of different types %v", errorsReported)
	}
}

func TestInvalidOperandsCheckerForInvalidBinaryOperationWithBools(t *testing.T) {
	exampleExpr := expr.NewSub(expr.NewBoolLit(true), expr.NewBoolLit(false))

	typeChecker := NewTypeChecker()
	exampleExpr.TypeCheck(typeChecker, nil)
	errorsReported := typeChecker.ErrorsEncountered

	expectedError := fmt.Sprintf("%v", fmt.Errorf("Encountered unexpected operand type, expected type: Integer, actual type: Boolean"))
	if len(errorsReported) != 2 || fmt.Sprintf("%v", errorsReported[0]) != expectedError || fmt.Sprintf("%v", errorsReported[1]) != expectedError {
		t.Errorf("Invalid operand operation checker did not correctly report invalid binary operation on bool types %v", errorsReported)
	}
}

func TestInvalidOperandsCheckerForInvalidBinaryOperationWithIntegers(t *testing.T) {
	exampleExpr := expr.NewAnd(expr.NewIntLit(10), expr.NewIntLit(8))

	typeChecker := NewTypeChecker()
	exampleExpr.TypeCheck(typeChecker, nil)
	errorsReported := typeChecker.ErrorsEncountered

	expectedError := fmt.Sprintf("%v", fmt.Errorf("Encountered unexpected operand type, expected type: Boolean, actual type: Integer"))

	if len(errorsReported) != 2 || fmt.Sprintf("%v", errorsReported[0]) != expectedError || fmt.Sprintf("%v", errorsReported[1]) != expectedError {
		t.Errorf("Invalid operand operation checker did not correctly report invalid binary operation on int types")
	}
}

func TestInvalidOperandsCheckerForInvalidBinaryOperationWithStrings(t *testing.T) {
	exampleExpr := expr.NewAnd(expr.NewStrLit("Test A"), expr.NewStrLit("Test B"))

	typeChecker := NewTypeChecker()
	exampleExpr.TypeCheck(typeChecker, nil)
	errorsReported := typeChecker.ErrorsEncountered

	expectedError := fmt.Sprintf("%v", fmt.Errorf("Encountered unexpected operand type, expected type: Boolean, actual type: String"))

	if len(errorsReported) != 2 || fmt.Sprintf("%v", errorsReported[0]) != expectedError || fmt.Sprintf("%v", errorsReported[1]) != expectedError {
		t.Errorf("Invalid operand operation checker did not correctly report invalid binary operation on string types")
	}
}

func TestInvalidOperandsCheckerForInvalidUnaryOperationWithBool(t *testing.T) {
	exampleExpr := expr.NewNeg(expr.NewBoolLit(true))

	typeChecker := NewTypeChecker()
	exampleExpr.TypeCheck(typeChecker, nil)
	errorsReported := typeChecker.ErrorsEncountered

	if len(errorsReported) != 1 || fmt.Sprintf("%v", errorsReported[0]) != fmt.Sprintf("%v", fmt.Errorf("Encountered unexpected operand type, expected type: Integer, actual type: Boolean")) {
		t.Errorf("Invalid operand operation checker did not correctly report invalid unary operation on bool type %v", errorsReported)
	}
}

func TestInvalidOperandsCheckerForInvalidUnaryOperationWithInt(t *testing.T) {
	exampleExpr := expr.NewNot(expr.NewIntLit(3))

	typeChecker := NewTypeChecker()
	exampleExpr.TypeCheck(typeChecker, nil)
	errorsReported := typeChecker.ErrorsEncountered

	if len(errorsReported) != 1 || fmt.Sprintf("%v", errorsReported[0]) != fmt.Sprintf("%v", fmt.Errorf("Encountered unexpected operand type, expected type: Boolean, actual type: Integer")) {
		t.Errorf("Invalid operand operation checker did not correctly report invalid unary operation on int type %v", errorsReported)
	}
}

func TestInvalidOperandsCheckerForInvalidUnaryOperationWithString(t *testing.T) {
	exampleExpr := expr.NewNot(expr.NewStrLit("Test"))

	typeChecker := NewTypeChecker()
	exampleExpr.TypeCheck(typeChecker, nil)
	errorsReported := typeChecker.ErrorsEncountered

	if len(errorsReported) != 1 || fmt.Sprintf("%v", errorsReported[0]) != fmt.Sprintf("%v", fmt.Errorf("Encountered unexpected operand type, expected type: Boolean, actual type: String")) {
		t.Errorf("Invalid operand operation checker did not correctly report invalid unary operation on string type")
	}
}

func TestUndefinedQuestionReferenceChecker(t *testing.T) {
	computedQuestion := stmt.NewComputedQuestion(expr.NewStrLit("Value residue:"), vari.NewVarDecl(vari.NewVarId("valueResidue"), expr.NewIntType()), expr.NewSub(expr.NewIntLit(100), expr.NewVarExpr(vari.NewVarId("hasMaintLoan"))))
	exampleBody := stmt.NewStmtList([]interfaces.Question{computedQuestion}, []interfaces.Conditional{})
	exampleForm := stmt.NewForm(vari.NewVarId("TestForm"), exampleBody)

	typeChecker := NewTypeChecker()
	exampleForm.TypeCheck(typeChecker, symbols.NewTypeCheckSymbols())
	errorsReported := typeChecker.ErrorsEncountered

	if len(errorsReported) != 1 || fmt.Sprintf("%v", errorsReported[0]) != fmt.Sprintf("%v", fmt.Errorf("Reference to unknown question identifier: hasMaintLoan")) {
		t.Errorf("Undefined questions references not reported correctly by type checker %v", errorsReported)
	}
}

func TestNonBoolConditionalChecker(t *testing.T) {
	exampleQuestion := stmt.NewInputQuestion(expr.NewStrLit("Did you sell a house in 2010?"), vari.NewVarDecl(vari.NewVarId("hasSoldHouse"), expr.NewBoolType()))
	exampleIf := stmt.NewIf(expr.NewIntLit(10), stmt.NewStmtList([]interfaces.Question{exampleQuestion}, []interfaces.Conditional{}))
	exampleBody := stmt.NewStmtList([]interfaces.Question{}, []interfaces.Conditional{exampleIf})
	exampleForm := stmt.NewForm(vari.NewVarId("TestForm"), exampleBody)

	typeChecker := NewTypeChecker()
	exampleForm.TypeCheck(typeChecker, symbols.NewTypeCheckSymbols())
	errorsReported := typeChecker.ErrorsEncountered

	if len(errorsReported) != 1 || fmt.Sprintf("%v", errorsReported[0]) != fmt.Sprintf("%v", fmt.Errorf("Non-boolean type used as condition: Integer")) {
		t.Errorf("Non bool condition type checker did not correctly report condition of invalid type %v", errorsReported)
	}
}

func TestDuplicateLabelChecker(t *testing.T) {
	firstQuestion := stmt.NewInputQuestion(expr.NewStrLit("Did you sell a house in 2010?"), vari.NewVarDecl(vari.NewVarId("hasSoldHouse"), expr.NewBoolType()))
	secondQuestion := stmt.NewInputQuestion(expr.NewStrLit("Did you sell a house in 2010?"), vari.NewVarDecl(vari.NewVarId("hasMaintLoan"), expr.NewBoolType()))
	exampleBody := stmt.NewStmtList([]interfaces.Question{firstQuestion, secondQuestion}, []interfaces.Conditional{})
	exampleForm := stmt.NewForm(vari.NewVarId("TestForm"), exampleBody)

	typeChecker := NewTypeChecker()
	exampleForm.TypeCheck(typeChecker, symbols.NewTypeCheckSymbols())
	warningsReported := typeChecker.WarningsEncountered

	if len(warningsReported) != 1 || fmt.Sprintf("%v", warningsReported[0]) != fmt.Sprintf("%v", fmt.Errorf("Label \"Did you sell a house in 2010?\" already used for question with identifier hasSoldHouse, using again for question with identifier hasMaintLoan")) {
		t.Errorf("Duplicate label not reported correctly by type checker %v", warningsReported)
	}
}

func TestDuplicateVarDeclChecker(t *testing.T) {
	firstQuestion := stmt.NewInputQuestion(expr.NewStrLit("Did you sell a house in 2010?"), vari.NewVarDecl(vari.NewVarId("hasSoldHouse"), expr.NewBoolType()))
	secondQuestion := stmt.NewInputQuestion(expr.NewStrLit("Did you sell a house in 2010?"), vari.NewVarDecl(vari.NewVarId("hasSoldHouse"), expr.NewIntType()))
	exampleBody := stmt.NewStmtList([]interfaces.Question{firstQuestion, secondQuestion}, []interfaces.Conditional{})
	exampleForm := stmt.NewForm(vari.NewVarId("TestForm"), exampleBody)

	typeChecker := NewTypeChecker()
	exampleForm.TypeCheck(typeChecker, symbols.NewTypeCheckSymbols())
	errorsReported := typeChecker.ErrorsEncountered

	if len(errorsReported) != 1 || fmt.Sprintf("%v", errorsReported[0]) != fmt.Sprintf("%v", fmt.Errorf("Question redeclared with different types: Integer and Boolean")) {
		t.Errorf("Duplicate var decl not reported correctly by type checker %v", errorsReported)
	}
}

func TestCyclicReferenceCheckerReferenceToEachOther(t *testing.T) {
	questionPointingToSecondQuestion := stmt.NewComputedQuestion(expr.NewStrLit("Did you sell a house in 2010?"), vari.NewVarDecl(vari.NewVarId("hasSoldHouse"), expr.NewBoolType()), expr.NewVarExpr(vari.NewVarId("hasBoughtHouse")))
	questionPointingToFirstQuestion := stmt.NewComputedQuestion(expr.NewStrLit("Did you buy a house in 2010?"), vari.NewVarDecl(vari.NewVarId("hasBoughtHouse"), expr.NewBoolType()), expr.NewVarExpr(vari.NewVarId("hasSoldHouse")))
	exampleBody := stmt.NewStmtList([]interfaces.Question{questionPointingToFirstQuestion, questionPointingToSecondQuestion}, []interfaces.Conditional{})
	exampleForm := stmt.NewForm(vari.NewVarId("TestForm"), exampleBody)

	typeChecker := NewTypeChecker()
	exampleForm.TypeCheck(typeChecker, symbols.NewTypeCheckSymbols())
	errorsReported := typeChecker.ErrorsEncountered

	if len(errorsReported) != 1 || fmt.Sprintf("%v", errorsReported[0]) != fmt.Sprintf("%v", fmt.Errorf("Found cyclic dependency: [hasSoldHouse hasBoughtHouse hasSoldHouse]")) {
		t.Errorf("Cyclic reference to self not reported correctly by type checker %v", errorsReported)
	}
}

func TestCyclicReferenceCheckerIfConditionRefersToBody(t *testing.T) {
	questionExample := stmt.NewInputQuestion(expr.NewStrLit("Did you sell a house in 2010?"), vari.NewVarDecl(vari.NewVarId("hasSoldHouse"), expr.NewBoolType()))
	ifBodyExample := stmt.NewStmtList([]interfaces.Question{questionExample}, []interfaces.Conditional{})
	ifExample := stmt.NewIf(expr.NewVarExpr(vari.NewVarId("hasSoldHouse")), ifBodyExample)
	exampleFormBody := stmt.NewStmtList([]interfaces.Question{}, []interfaces.Conditional{ifExample})
	exampleForm := stmt.NewForm(vari.NewVarId("TestForm"), exampleFormBody)

	typeChecker := NewTypeChecker()
	exampleForm.TypeCheck(typeChecker, symbols.NewTypeCheckSymbols())
	errorsReported := typeChecker.ErrorsEncountered

	if len(errorsReported) != 1 || fmt.Sprintf("%v", errorsReported[0]) != fmt.Sprintf("%v", fmt.Errorf("Found cyclic dependency: [hasSoldHouse hasSoldHouse]")) {
		t.Errorf("Cyclic reference not reported correctly by type checker %v", errorsReported)
	}
}

func TestQuestionTypeAndComputationTypeMismatch(t *testing.T) {
	exampleQuestion := stmt.NewComputedQuestion(expr.NewStrLit("Did you sell a house in 2010?"), vari.NewVarDecl(vari.NewVarId("hasSoldHouse"), expr.NewBoolType()), expr.NewIntLit(10))
	exampleBody := stmt.NewStmtList([]interfaces.Question{exampleQuestion}, []interfaces.Conditional{})
	exampleForm := stmt.NewForm(vari.NewVarId("TestForm"), exampleBody)

	typeChecker := NewTypeChecker()
	exampleForm.TypeCheck(typeChecker, symbols.NewTypeCheckSymbols())
	errorsReported := typeChecker.ErrorsEncountered

	if len(errorsReported) != 1 || fmt.Sprintf("%v", errorsReported[0]) != fmt.Sprintf("%v", fmt.Errorf("Encountered computed question with mismatch between declared type (Boolean) and actual computation type (Integer)")) {
		t.Errorf("Non bool condition type checker did not correctly report condition of invalid type %v", errorsReported)
	}
}
