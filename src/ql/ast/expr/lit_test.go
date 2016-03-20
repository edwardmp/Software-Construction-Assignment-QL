package expr

import "testing"

func TestIntLit(t *testing.T) {
	exampleIntLit := NewIntLit(10)

	if NewIntLit(10) != exampleIntLit {
		t.Errorf("IntLit creation test did not succeed")
	}
}

func TestIntLitGetValue(t *testing.T) {
	if NewIntLit(10).GetValue() != 10 {
		t.Errorf("IntLit value retrieval test did not succeed")
	}
}

func TestBoolLit(t *testing.T) {
	exampleBoolLit := NewBoolLit(true)

	if NewBoolLit(true) != exampleBoolLit {
		t.Errorf("BoolLit creation test did not succeed")
	}
}

func TestBoolLitGetValue(t *testing.T) {
	if NewBoolLit(true).GetValue() != true {
		t.Errorf("BoolLit value retrieval test did not succeed")
	}
}

func TestStrLit(t *testing.T) {
	exampleStrLit := NewStrLit("Test")

	if NewStrLit("Test") != exampleStrLit {
		t.Errorf("StrLit creation test did not succeed")
	}
}

func TestStrLitGetValue(t *testing.T) {
	if NewStrLit("Test").GetValue() != "Test" {
		t.Errorf("StrLit value retrieval test did not succeed")
	}
}
