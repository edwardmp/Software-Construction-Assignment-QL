package symbols

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
	"ql/interfaces"
)

type VarIdToExprSymbolTable map[interfaces.VarId]interfaces.Expr

type SymbolCallBackFunction func(interfaces.VarIdValueSymbols)

type VarIdValueSymbols struct {
	Table               VarIdToExprSymbolTable
	RegisteredCallbacks []func()
}

func NewVarIdValueSymbols() *VarIdValueSymbols {
	return &VarIdValueSymbols{Table: make(VarIdToExprSymbolTable), RegisteredCallbacks: make([]func(), 0)}
}

func (this *VarIdValueSymbols) SetExprForVarId(expr interfaces.Expr, varId interfaces.VarId) {
	if expr == nil || varId == nil {
		panic("Trying to set Expr for VarId to nil or varId is nil")
	}

	this.Table[varId] = expr
	log.WithFields(log.Fields{"Identifier": varId, "Expr": expr}).Debug("Set Expr for VarId")

	for _, registeredCallback := range this.RegisteredCallbacks {
		registeredCallback()
	}
}

func (this *VarIdValueSymbols) ExprForVarId(varId interfaces.VarId) interfaces.Expr {
	if varId == nil {
		panic("Trying to get Expr for nil VarId")
	}

	expr := this.Table[varId]
	log.WithFields(log.Fields{"Identifier": varId, "Expr": expr}).Debug("Looking up Expr for VarId in SymbolTable")

	return expr
}

func (this *VarIdValueSymbols) RegisterCallback(callback func()) {
	this.RegisteredCallbacks = append(this.RegisteredCallbacks, callback)
}

func (this *VarIdValueSymbols) SaveToDisk() (interface{}, error) {
	formDataAsJSON, _ := convertSymbolTableToJSON(this.convertSymbolTableKeysToStrings())

	writeErr := ioutil.WriteFile("savedForm.json", formDataAsJSON, 0644)

	if writeErr != nil {
		log.WithFields(log.Fields{"error": writeErr}).Error("Encountered error during writing to disk of symbolTable data")
		return nil, writeErr
	}

	log.Info("SymbolTable written to disk as JSON")

	return formDataAsJSON, nil
}

func (this *VarIdValueSymbols) convertSymbolTableKeysToStrings() map[string]interface{} {
	var symbolTableWithStringKeys map[string]interface{} = make(map[string]interface{})
	for varId, expr := range this.Table {
		symbolTableWithStringKeys[varId.Identifier()] = expr.Eval(this)
	}

	return symbolTableWithStringKeys
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
