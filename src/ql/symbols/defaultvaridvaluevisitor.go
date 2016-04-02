package symbols

import (
	"ql/ast/visitor"
	"ql/interfaces"
)

type DefaultVarIdValueVisitor struct {
	visitor.BaseVisitor
}

func NewDefaultVarIdValueVisitor() *DefaultVarIdValueVisitor {
	return &DefaultVarIdValueVisitor{visitor.BaseVisitor{}}
}

func (this *DefaultVarIdValueVisitor) StartSettingDefaultValuesForVarIds(form interfaces.Form) *VarIdValueSymbols {
	varIdValueSymbols := NewVarIdValueSymbols()
	form.Accept(this, varIdValueSymbols)

	return varIdValueSymbols
}

func (this *DefaultVarIdValueVisitor) VisitVarDecl(varDecl interfaces.VarDecl, context interface{}) {
	symbols := context.(interfaces.VarIdValueSymbols)
	symbols.SetExprForVarId(varDecl.Type().DefaultValue(), varDecl.VariableIdentifier())
}
