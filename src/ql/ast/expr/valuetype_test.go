package expr

import "testing"

func TestVarTypeDefaultValueForIntType(t *testing.T) {
	varTypeExample := NewIntTypeNoSourceInfo()
	correctDefaultValue := NewIntLitNoSourceInfo(0)
	if varTypeExample.GetDefaultValue() != correctDefaultValue {
		t.Errorf("Int Type default value invalid")
	}
}

func TestVarTypeDefaultValueForBoolType(t *testing.T) {
	varTypeExample := NewBoolTypeNoSourceInfo()
	correctDefaultValue := NewBoolLitNoSourceInfo(false)
	if varTypeExample.GetDefaultValue() != correctDefaultValue {
		t.Errorf("Bool Type default value invalid")
	}
}

func TestVarTypeDefaultValueForStringType(t *testing.T) {
	varTypeExample := NewStringTypeNoSourceInfo()
	correctDefaultValue := NewStrLitNoSourceInfo("")
	if varTypeExample.GetDefaultValue() != correctDefaultValue {
		t.Errorf("String Type default value invalid")
	}
}
