package env

import (
	log "github.com/Sirupsen/logrus"
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
