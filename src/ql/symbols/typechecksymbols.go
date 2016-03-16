package symbols

import (
	"fmt"
	"ql/interfaces"
)

type TypeCheckSymbols struct {
	*Symbols
}

func NewTypeCheckSymbols() *TypeCheckSymbols {
	return &TypeCheckSymbols{Symbols: newSymbols()}
}

func (this *TypeCheckSymbols) SetTypeForVarId(varType interfaces.ValueType, varId interfaces.VarId) {
	if varType == nil || varId == nil {
		panic("Trying to set type for VarId to nil or varId is nil")
	}

	this.setNodeForIdentifier(varType, varId)
}

func (this *TypeCheckSymbols) GetTypeForVarId(varId interfaces.VarId) interfaces.ValueType {
	if varId == nil {
		panic("Trying to get type for nil VarId")
	}

	fmt.Println(this.Symbols)
	return this.getNodeForIdentifier(varId).(interfaces.ValueType)
}

func (this *TypeCheckSymbols) IsTypeSetForVarId(varId interfaces.VarId) bool {
	if this.getNodeForIdentifier(varId) == nil {
		return false
	}

	return true
}
