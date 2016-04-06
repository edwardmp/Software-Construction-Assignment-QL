package symbols

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
	"ql/interfaces"
)

type varIdToExprSymbolTable map[interfaces.VarId]interfaces.Expr

type symbolCallBackFunction func(interfaces.VarIdValueSymbols)

type VarIdValueSymbols struct {
	Table               varIdToExprSymbolTable
	RegisteredCallbacks []func()
}

func NewVarIdValueSymbols() *VarIdValueSymbols {
	log.Info("Creating new VarIdValueSymbols")
	return &VarIdValueSymbols{Table: make(varIdToExprSymbolTable), RegisteredCallbacks: make([]func(), 0)}
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

/* methods related to export of symboltable to file */

type VarIdStringExprEvalSymbolTable map[string]interface{}

func (this *VarIdValueSymbols) SaveToDisk() (interface{}, error) {
	formDataAsJSON, _ := convertSymbolTableToJSON(this.convertSymbolTableKeysToStringsAndEvalValues())

	writeErr := ioutil.WriteFile("savedForm.json", formDataAsJSON, 0644)

	if writeErr != nil {
		log.WithFields(log.Fields{"error": writeErr}).Error("Encountered error during writing to disk of symbolTable data")
		return nil, writeErr
	}

	log.Info("SymbolTable written to disk as JSON")

	return formDataAsJSON, nil
}

func (this *VarIdValueSymbols) convertSymbolTableKeysToStringsAndEvalValues() VarIdStringExprEvalSymbolTable {
	symbolTableWithStringKeys := make(VarIdStringExprEvalSymbolTable)
	for varId, expr := range this.Table {
		symbolTableWithStringKeys[varId.Identifier()] = expr.Eval(this).PrimitiveValue()
	}

	return symbolTableWithStringKeys
}

func convertSymbolTableToJSON(symbolTableWithStringKeys VarIdStringExprEvalSymbolTable) ([]byte, error) {
	formDataAsJSON, jsonErr := json.MarshalIndent(symbolTableWithStringKeys, "", "  ")

	if jsonErr != nil {
		log.WithFields(log.Fields{"error": jsonErr}).Error("Encountered error during symbolTable to JSON conversion")
		return nil, jsonErr
	}

	log.WithFields(log.Fields{"formDataAsJSON": string(formDataAsJSON[:])}).Debug("Successful conversion of symbolTable to JSON")

	return formDataAsJSON, nil
}
