package vari

import "ql/interfaces"

func (this VarDecl) TypeCheck(typeChecker interfaces.TypeChecker, symbols interfaces.TypeCheckSymbols) interface{} {
	// store type for identifier so when we find VarExpr with this VarId we know its rea type (used during typechecking)
	symbols.SetTypeForVarId(this.GetType(), this.GetIdent())

	// we mark it as known to indicate that earlier references to this VarId are valid
	typeChecker.MarkVarIdAsKnown(this.GetIdent())

	return nil
}
