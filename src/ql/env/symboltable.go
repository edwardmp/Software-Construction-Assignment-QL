package env

import (
	log "github.com/Sirupsen/logrus"
	"ql/ast/expr"
	"ql/ast/vari"
)

/* Symbol Table */

type SymbolTable map[vari.VarId]expr.Expr

func (s SymbolTable) getNodeForIdentifier(v vari.VarId) expr.Expr {
	return s[v]
}

func (s SymbolTable) setNodeForIdentifier(e expr.Expr, v vari.VarId) SymbolTable {
	if previousValue, keyExists := s[v]; keyExists {
		s[v] = e
		log.WithFields(log.Fields{"Identifier": v, "Current": s[v], "Previous": previousValue}).Debug("Set node for identifier")
	} else {
		s[v] = e
		log.WithFields(log.Fields{"Identifier": v, "Current": s[v]}).Debug("Set node for identifier")
	}

	return s
}

/* Symbol Table Stack */

type SymbolTableStack []SymbolTable

func NewSymbolTableStack() SymbolTableStack {
	log.Info("Creating new symbol table stack")
	return make(SymbolTableStack, 1)
}

func NewSymbolTable() SymbolTable {
	log.Debug("Creating new symbol table")
	return make(SymbolTable)
}

func (s SymbolTableStack) Push(v SymbolTable) SymbolTableStack {
	s = append(s, v)
	log.WithFields(log.Fields{"SymbolTableStack": s}).Info("Pushed new symbol table")
	return s
}

func (s SymbolTableStack) Pop() (SymbolTableStack, SymbolTable) {
	if len(s) == 0 {
		panic("Trying to pop from empty stack")
	}

	l := len(s)
	return s[:l-1], s[l-1]
}

func (s SymbolTableStack) Peek() SymbolTable {
	if len(s) == 0 {
		panic("Trying to peek from empty stack")
	}

	log.WithFields(log.Fields{"SymbolTableStack": s}).Debug("Peeking from symbol table stack")
	return s[len(s)-1]
}

func (s SymbolTableStack) NewSymbolTableWithParentScope() SymbolTableStack {
	NewSymbolTable := NewSymbolTable()

	// copy all symbols in parent scope table
	for k, v := range s.Peek() {
		NewSymbolTable[k] = v
	}

	s = s.Push(NewSymbolTable)

	log.WithFields(log.Fields{"SymbolTable": s}).Info("Pushed new symbol table with parent scope on stack")

	return s
}

func (s SymbolTableStack) SetValueForIdentifierInTopSymbolTable(e expr.Expr, v vari.VarId) {
	s.Peek().setNodeForIdentifier(e, v)
}

func (s SymbolTableStack) GetValueForIdentifierInTopSymbolTable(v vari.VarId) expr.Expr {
	return s.Peek().getNodeForIdentifier(v)
}
