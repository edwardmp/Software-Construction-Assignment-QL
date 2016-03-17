package symbols

import (
	log "github.com/Sirupsen/logrus"
	"ql/interfaces"
)

type VarIdTovalueTypeSymbolTable map[interfaces.VarId]interfaces.ValueType

type TypeCheckSymbols struct {
	Table VarIdTovalueTypeSymbolTable
}

func NewTypeCheckSymbols() *TypeCheckSymbols {
	log.Debug("Creating new TypeCheckSymbols")
	return &TypeCheckSymbols{Table: make(VarIdTovalueTypeSymbolTable)}
}

func (this *TypeCheckSymbols) SetTypeForVarId(valueType interfaces.ValueType, varId interfaces.VarId) {
	if valueType == nil || varId == nil {
		panic("Trying to set value type for VarId to nil or varId is nil")
	}

	this.Table[varId] = valueType
	log.WithFields(log.Fields{"Identifier": varId, "valueType": valueType}).Debug("Set ValueType for VarId")
}

func (this *TypeCheckSymbols) GetTypeForVarId(varId interfaces.VarId) interfaces.ValueType {
	if varId == nil {
		panic("Trying to get type for nil VarId")
	}

	valueType := this.Table[varId]
	log.WithFields(log.Fields{"Identifier": varId, "Value": valueType}).Debug("Looking up valueType for VarId in SymbolTable")

	return valueType
}

func (this *TypeCheckSymbols) IsTypeSetForVarId(varId interfaces.VarId) bool {
	if this.GetTypeForVarId(varId) == nil {
		return false
	}

	return true
}
