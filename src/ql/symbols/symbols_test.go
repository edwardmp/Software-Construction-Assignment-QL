package symbols

import (
	"io/ioutil"
	"os"
	"ql/ast/expr"
	"ql/ast/vari"
	"strings"
	"testing"
)

func TestSymbolTableAdd(t *testing.T) {
	newSymbolTable := newSymbols()
	exampleVarId := vari.NewVarIdNoSourceInfo("testIdentifier")
	newSymbolTable.setNodeForIdentifier("testValue", exampleVarId)

	if lookupValue := newSymbolTable.getNodeForIdentifier(exampleVarId); lookupValue != "testValue" {
		t.Errorf("New symbol table created not correct, expected value %s for key %s, is %s", "testValue", exampleVarId, lookupValue)
	}
}

// TODO move to own file for that type of symboltable
func TestSaveToDisk(t *testing.T) {
	newSymbolTable := NewVarIdValueSymbols()
	exampleVarId := vari.NewVarIdNoSourceInfo("testIdentifier")
	newSymbolTable.SetExprForVarId(expr.NewStrLitNoSourceInfo("testValue"), exampleVarId)

	newSymbolTable.SaveToDisk()

	qlFile, err := ioutil.ReadFile("savedForm.json")
	if err != nil || !strings.Contains(string(qlFile), "testIdentifier\": \"testValue") {
		t.Errorf("Output file does not contain correct data %s", qlFile)
	}

	// clean up file
	removeOutputFileAfterTest()
}

func removeOutputFileAfterTest() {
	err := os.Remove("savedForm.json")

	if err != nil {
		panic(err)
	}
}
