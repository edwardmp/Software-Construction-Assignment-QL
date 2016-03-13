package symboltable

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
	"ql/interfaces"
)

type Symbols struct {
	Table               SymbolTable
	RegisteredCallbacks []func(interfaces.SymbolTable)
}

func NewSymbols() *Symbols {
	return &Symbols{Table: NewSymbolTable(), RegisteredCallbacks: make([]func(interfaces.SymbolTable), 0)}
}

type SymbolTable map[interfaces.VarId]interface{}

func NewSymbolTable() SymbolTable {
	log.Debug("Creating new symbol table")
	return make(SymbolTable)
}

func (s *Symbols) RegisterCallback(callback func(interfaces.SymbolTable)) {
	s.RegisteredCallbacks = append(s.RegisteredCallbacks, callback)
}

func (s *Symbols) GetNodeForIdentifier(v interfaces.VarId) interface{} {
	value := s.Table[v]
	log.WithFields(log.Fields{"Identifier": v, "Value": value}).Debug("Looking up identifier in SymbolTable")

	return value
}

func (s *Symbols) SetNodeForIdentifier(e interface{}, v interfaces.VarId) {
	s.Table[v] = e
	log.WithFields(log.Fields{"Identifier": v, "Current": s.Table[v]}).Debug("Set node for identifier")

	for _, registeredCallback := range s.RegisteredCallbacks {
		registeredCallback(s)
	}
}

func (s *Symbols) SaveToDisk() (interface{}, error) {
	formDataAsJSON, _ := convertSymbolTableToJSON(convertSymbolTableKeysToStrings(s.Table))

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
		symbolTableWithStringKeys[k.GetIdent()] = v
	}

	return symbolTableWithStringKeys
}
