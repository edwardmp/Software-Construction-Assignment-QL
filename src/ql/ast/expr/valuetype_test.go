package expr

import "testing"

func TestVarTypeDefaultValueForIntType(t *testing.T) {
	varTypeExample := NewIntType()
	correctDefaultValue := NewIntLit(0)
	if varTypeExample.GetDefaultValue() != correctDefaultValue {
		t.Errorf("Int Type default value invalid")
	}
}

func TestVarTypeDefaultValueForBoolType(t *testing.T) {
	varTypeExample := NewBoolType()
	correctDefaultValue := NewBoolLit(false)
	if varTypeExample.GetDefaultValue() != correctDefaultValue {
		t.Errorf("Bool Type default value invalid")
	}
}

func TestVarTypeDefaultValueForStringType(t *testing.T) {
	varTypeExample := NewStringType()
	correctDefaultValue := NewStrLit("")
	if varTypeExample.GetDefaultValue() != correctDefaultValue {
		t.Errorf("String Type default value invalid")
	}
}
