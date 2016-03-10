package vari

import (
	"ql/ast/expr"
	"testing"
)

func TestVarTypeDefaultValueForIntType(t *testing.T) {
	varTypeExample := IntType{}
	correctDefaultValue := expr.NewIntLit(0)
	if varTypeExample.GetDefaultValue() != correctDefaultValue {
		t.Errorf("Int Type default value invalid")
	}
}

func TestVarTypeDefaultValueForBoolType(t *testing.T) {
	varTypeExample := BoolType{}
	correctDefaultValue := expr.NewBoolLit(false)
	if varTypeExample.GetDefaultValue() != correctDefaultValue {
		t.Errorf("Bool Type default value invalid")
	}
}

func TestVarTypeDefaultValueForStringType(t *testing.T) {
	varTypeExample := StringType{}
	correctDefaultValue := expr.NewStrLit("")
	if varTypeExample.GetDefaultValue() != correctDefaultValue {
		t.Errorf("String Type default value invalid")
	}
}
