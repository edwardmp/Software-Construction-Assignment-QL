package vari

import (
	"testing"
    "ql/ast/expr/litexpr"
)

func TestVarTypeDefaultValueForIntType(t *testing.T) {
	varTypeExample := IntType{}
    correctDefaultValue := litexpr.IntLit{0}
	if varTypeExample.GetDefaultValue() != correctDefaultValue {
		t.Errorf("Int Type default value invalid")
	}
}

func TestVarTypeDefaultValueForBoolType(t *testing.T) {
    varTypeExample := BoolType{}
    correctDefaultValue := litexpr.BoolLit{false}
    if varTypeExample.GetDefaultValue() != correctDefaultValue {
        t.Errorf("Bool Type default value invalid")
    }
}

func TestVarTypeDefaultValueForStringType(t *testing.T) {
    varTypeExample := StringType{}
    correctDefaultValue := litexpr.StrLit{""}
    if varTypeExample.GetDefaultValue() != correctDefaultValue {
        t.Errorf("String Type default value invalid")
    }
}
