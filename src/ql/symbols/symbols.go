package symbols

import (
	log "github.com/Sirupsen/logrus"
	"ql/interfaces"
)

type SymbolTable map[interfaces.VarId]interface{}

type Symbols struct {
	Table SymbolTable
}

func newSymbols() *Symbols {
	return &Symbols{Table: newSymbolTable()}
}

func newSymbolTable() SymbolTable {
	log.Debug("Creating new symbol table")
	return make(SymbolTable)
}

func (this *Symbols) getNodeForIdentifier(v interfaces.VarId) interface{} {
	value := this.Table[v]
	log.WithFields(log.Fields{"Identifier": v, "Value": value}).Debug("Looking up node for identifier in SymbolTable")

	return value
}

func (this *Symbols) setNodeForIdentifier(e interface{}, v interfaces.VarId) {
	this.Table[v] = e
	log.WithFields(log.Fields{"Identifier": v, "Current": this.Table[v]}).Debug("Set node for identifier")
}
