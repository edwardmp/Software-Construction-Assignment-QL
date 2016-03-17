package vari

import "ql/interfaces"

func (this VarDecl) TypeCheck(typeChecker interfaces.TypeChecker, symbols interfaces.TypeCheckSymbols) interface{} {
	// store type for identifier so when we find VarExpr with this id we know type (used during typechecking)
	symbols.SetTypeForVarId(this.Type, this.Ident)

	// we make it as known so earlier references to this VarId are known to be valid
	typeChecker.MarkVarIdAsKnown(this.Ident)

	return nil
}
