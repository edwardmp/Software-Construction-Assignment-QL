package vari

import (
	"ql/ast/expr"
	"testing"
)

func testVarTypeDefaultValueForIntType(t *testing.T) {
	varTypeExample := IntType{}
	correctDefaultValue := expr.NewIntLit(0)
	if varTypeExample.GetDefaultValue() != correctDefaultValue {
		t.Errorf("Int Type default value invalid")
	}
}

func testVarTypeDefaultValueForBoolType(t *testing.T) {
	varTypeExample := BoolType{}
	correctDefaultValue := expr.NewBoolLit(false)
	if varTypeExample.GetDefaultValue() != correctDefaultValue {
		t.Errorf("Bool Type default value invalid")
	}
}

func testVarTypeDefaultValueForStringType(t *testing.T) {
	varTypeExample := StringType{}
	correctDefaultValue := expr.NewStrLit("")
	if varTypeExample.GetDefaultValue() != correctDefaultValue {
		t.Errorf("String Type default value invalid")
	}
}
