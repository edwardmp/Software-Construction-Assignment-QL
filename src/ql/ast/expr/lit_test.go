package expr

import "testing"

func testIntLit(t *testing.T) {
	exampleIntLit := IntLit{10}

	if NewIntLit(10) != exampleIntLit {
		t.Errorf("IntLit creation test did not succeed")
	}
}

func testIntLitGetValue(t *testing.T) {
	if NewIntLit(10).GetValue() != 10 {
		t.Errorf("IntLit value retrieval test did not succeed")
	}
}

func testBoolLit(t *testing.T) {
	exampleBoolLit := BoolLit{true}

	if NewBoolLit(true) != exampleBoolLit {
		t.Errorf("BoolLit creation test did not succeed")
	}
}

func testBoolLitGetValue(t *testing.T) {
	if NewBoolLit(true).GetValue() != true {
		t.Errorf("BoolLit value retrieval test did not succeed")
	}
}

func testStrLit(t *testing.T) {
	exampleStrLit := StrLit{"Test"}

	if NewStrLit("Test") != exampleStrLit {
		t.Errorf("StrLit creation test did not succeed")
	}
}

func testStrLitGetValue(t *testing.T) {
	if NewStrLit("Test").GetValue() != "Test" {
		t.Errorf("StrLit value retrieval test did not succeed")
	}
}
