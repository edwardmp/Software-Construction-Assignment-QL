package env

import (
	"io/ioutil"
	"os"
	"ql/ast/vari"
	"ql/env"
	"strings"
	"testing"
)

func TestSymbolTableAdd(t *testing.T) {
	newSymbolTable := env.NewSymbolTable()
	exampleVarId := vari.VarId{"testIdentifier"}
	newSymbolTable.SetNodeForIdentifier("testValue", exampleVarId)

	if lookupValue := newSymbolTable.GetNodeForIdentifier(exampleVarId); lookupValue != "testValue" {
		t.Errorf("New symbol table created not correct, expected value %s for key %s, is %s", "testValue", exampleVarId, lookupValue)
	}
}

func TestSaveToDisk(t *testing.T) {
	newSymbolTable := env.NewSymbolTable()
	exampleVarId := vari.VarId{"testIdentifier"}
	newSymbolTable.SetNodeForIdentifier("testValue", exampleVarId)

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
