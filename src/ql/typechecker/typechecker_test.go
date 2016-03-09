package typechecker

import (
	"fmt"
	"ql/ast/expr/binaryoperatorexpr"
	"ql/ast/expr/litexpr"
	"ql/ast/expr/unaryoperatorexpr"
	"ql/ast/stmt"
	"ql/ast/vari"
	"ql/symboltable"
	"testing"
)

func TestDuplicateLabelChecker(t *testing.T) {
	firstQuestionOutput := stmt.InputQuestion{litexpr.StrLit{"Did you sell a house in 2010?"}, vari.VarDecl{vari.VarId{"hasSoldHouse"}, vari.BoolType{}}}
	secondQuestionOutput := stmt.InputQuestion{litexpr.StrLit{"Did you sell a house in 2010?"}, vari.VarDecl{vari.VarId{"hasMaintLoan"}, vari.BoolType{}}}
	exampleBodyOutput := stmt.StmtList{[]stmt.Question{firstQuestionOutput, secondQuestionOutput}, []stmt.Conditional{}}
	exampleOutputForm := stmt.Form{vari.VarId{"TestForm"}, exampleBodyOutput}

	warningsReported := CheckForDuplicateLabels(exampleOutputForm)

	if len(warningsReported) != 1 || fmt.Sprintf("%v", warningsReported[0]) != fmt.Sprintf("%v", fmt.Errorf("Label \"Did you sell a house in 2010?\" already used for question with identifier hasSoldHouse, using again for question with identifier hasMaintLoan")) {
		t.Errorf("Duplicate label not reported correctly by type checker")
	}
}

func TestDuplicateVarDeclChecker(t *testing.T) {
	firstQuestionOutput := stmt.InputQuestion{litexpr.StrLit{"Did you sell a house in 2010?"}, vari.VarDecl{vari.VarId{"hasSoldHouse"}, vari.BoolType{}}}
	secondQuestionOutput := stmt.InputQuestion{litexpr.StrLit{"Did you sell a house in 2010?"}, vari.VarDecl{vari.VarId{"hasSoldHouse"}, vari.IntType{}}}
	exampleBodyOutput := stmt.StmtList{[]stmt.Question{firstQuestionOutput, secondQuestionOutput}, []stmt.Conditional{}}
	exampleOutputForm := stmt.Form{vari.VarId{"TestForm"}, exampleBodyOutput}

	errorsReported := CheckForDuplicateVarDeclWithDiffTypes(exampleOutputForm)

	if len(errorsReported) != 1 || fmt.Sprintf("%v", errorsReported[0]) != fmt.Sprintf("%v", fmt.Errorf("Question redeclared with different types: vari.IntType and vari.BoolType")) {
		t.Errorf("Duplicate var decl not reported correctly by type checker")
	}
}

func TestUndefinedQuestionReferenceChecker(t *testing.T) {
	computedQuestion := stmt.ComputedQuestion{litexpr.StrLit{"Value residue:"}, vari.VarDecl{vari.VarId{"valueResidue"}, vari.IntType{}}, binaryoperatorexpr.Sub{unaryoperatorexpr.VarExpr{vari.VarId{"hasSoldHouse"}}, unaryoperatorexpr.VarExpr{vari.VarId{"hasMaintLoan"}}}}
	exampleBody := stmt.StmtList{[]stmt.Question{computedQuestion}, []stmt.Conditional{}}
	exampleForm := stmt.Form{vari.VarId{"TestForm"}, exampleBody}

	errorsReported := CheckForReferencesToUndefinedQuestions(exampleForm, symboltable.NewSymbolTable())

	if len(errorsReported) != 2 || fmt.Sprintf("%v", errorsReported[0]) != fmt.Sprintf("%v", fmt.Errorf("Reference to unknown question identifier: hasSoldHouse")) || fmt.Sprintf("%v", errorsReported[1]) != fmt.Sprintf("%v", fmt.Errorf("Reference to unknown question identifier: hasMaintLoan")) {
		t.Errorf("Undefined questions references not reported correctly by type checker")
	}
}

func TestNonBoolConditionalChecker(t *testing.T) {
	exampleQuestion := stmt.InputQuestion{litexpr.StrLit{"Did you sell a house in 2010?"}, vari.VarDecl{vari.VarId{"hasSoldHouse"}, vari.BoolType{}}}
	exampleIf := stmt.If{litexpr.IntLit{10}, stmt.StmtList{[]stmt.Question{exampleQuestion}, []stmt.Conditional{}}}
	exampleBody := stmt.StmtList{[]stmt.Question{}, []stmt.Conditional{exampleIf}}
	exampleForm := stmt.Form{vari.VarId{"TestForm"}, exampleBody}

	errorsReported := CheckForNonBoolConditions(exampleForm, symboltable.NewSymbolTable())

	if len(errorsReported) != 1 || fmt.Sprintf("%v", errorsReported[0]) != fmt.Sprintf("%v", fmt.Errorf("Non-boolean type used as condition: int")) {
		t.Errorf("Non bool condition type checker did not correctly report condition of invalid type")
	}
}
