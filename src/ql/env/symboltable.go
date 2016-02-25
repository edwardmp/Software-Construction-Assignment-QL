package env

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
	"ql/ast/vari"
)

/* Symbol Table */

type SymbolTable map[vari.VarId]interface{}

func (s SymbolTable) GetNodeForIdentifier(v vari.VarId) interface{} {
	return s[v]
}

func (s SymbolTable) SetNodeForIdentifier(e interface{}, v vari.VarId) *SymbolTable {
	if previousValue, keyExists := s[v]; keyExists {
		s[v] = e
		log.WithFields(log.Fields{"Identifier": v, "Current": s[v], "Previous": previousValue}).Debug("Set node for identifier")
	} else {
		s[v] = e
		log.WithFields(log.Fields{"Identifier": v, "Current": s[v]}).Debug("Set node for identifier")
	}

	return &s
}

func NewSymbolTable() SymbolTable {
	log.Debug("Creating new symbol table")
	return make(SymbolTable)
}

func (s SymbolTable) SaveToDisk() (interface{}, error) {
	formDataAsJSON, _ := convertSymbolTableToJSON(convertSymbolTableKeysToStrings(s))

	writeErr := ioutil.WriteFile("savedForm.json", formDataAsJSON, 0644)

	if writeErr != nil {
		log.WithFields(log.Fields{"error": writeErr}).Error("Encountered error during writing to disk of symbolTable data")
		return nil, writeErr
	}

	log.Info("SymbolTable written to disk as JSON")

	return formDataAsJSON, nil
}

func convertSymbolTableToJSON(symbolTableWithStringKeys map[string]interface{}) ([]byte, error) {
	formDataAsJSON, jsonErr := json.MarshalIndent(symbolTableWithStringKeys, "", "  ")

	if jsonErr != nil {
		log.WithFields(log.Fields{"error": jsonErr}).Error("Encountered error during symbolTable to JSON conversion")
		return nil, jsonErr
	}

	log.WithFields(log.Fields{"formDataAsJSON": string(formDataAsJSON[:])}).Debug("Successful conversion of symbolTable to JSON")

	return formDataAsJSON, nil
}

func convertSymbolTableKeysToStrings(s SymbolTable) map[string]interface{} {
	var symbolTableWithStringKeys map[string]interface{} = make(map[string]interface{})
	for k, v := range s {
		symbolTableWithStringKeys[k.Ident] = v
	}

	return symbolTableWithStringKeys
}
