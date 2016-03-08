package typechecker

import (
	"fmt"
	"ql/ast/expr/lit"
	"ql/ast/stmt"
	"ql/ast/vari"
	"ql/ast/vari/vartype"
	"ql/typechecker"
	"testing"
)

func TestDuplicateLabelChecker(t *testing.T) {
	firstQuestionOutput := stmt.InputQuestion{lit.StrLit{"Did you sell a house in 2010?"}, vari.VarDecl{vari.VarId{"hasSoldHouse"}, vartype.BoolType{}}}
	secondQuestionOutput := stmt.InputQuestion{lit.StrLit{"Did you sell a house in 2010?"}, vari.VarDecl{vari.VarId{"hasMaintLoan"}, vartype.BoolType{}}}
	exampleBodyOutput := stmt.StmtList{[]stmt.Question{firstQuestionOutput, secondQuestionOutput}, []stmt.Conditional{}}
	exampleOutputForm := stmt.Form{vari.VarId{"TestForm"}, exampleBodyOutput}

	errorsReported := typechecker.CheckForDuplicateLabels(exampleOutputForm)

	if len(errorsReported) != 1 || fmt.Sprintf("%v", errorsReported[0]) != fmt.Sprintf("%v", fmt.Errorf("Label \"Did you sell a house in 2010?\" already used for question with identifier hasSoldHouse, using again for question with identifier hasMaintLoan")) {
		t.Errorf("Duplicate label not reported correctly by type checker")
	}
}

func TestDuplicateVarDeclChecker(t *testing.T) {
	firstQuestionOutput := stmt.InputQuestion{lit.StrLit{"Did you sell a house in 2010?"}, vari.VarDecl{vari.VarId{"hasSoldHouse"}, vartype.BoolType{}}}
	secondQuestionOutput := stmt.InputQuestion{lit.StrLit{"Did you sell a house in 2010?"}, vari.VarDecl{vari.VarId{"hasSoldHouse"}, vartype.IntType{}}}
	exampleBodyOutput := stmt.StmtList{[]stmt.Question{firstQuestionOutput, secondQuestionOutput}, []stmt.Conditional{}}
	exampleOutputForm := stmt.Form{vari.VarId{"TestForm"}, exampleBodyOutput}

	errorsReported := typechecker.CheckForDuplicateVarDeclWithDiffTypes(exampleOutputForm)

	if len(errorsReported) != 1 || fmt.Sprintf("%v", errorsReported[0]) != fmt.Sprintf("%v", fmt.Errorf("Question redeclared with different types: vartype.IntType and vartype.BoolType")) {
		t.Errorf("Duplicate var decl not reported correctly by type checker")
	}
}
