package vari

import "ql/interfaces"

func (this VarDecl) TypeCheck(typeCheckArgs interfaces.TypeCheckArgs) interface{} {
	// store type for identifier so when we find VarExpr with this VarId we know its real type (used during typechecking)
	typeCheckArgs.Symbols().SetTypeForVarId(this.Type(), this.Identifier())

	// we mark it as known to indicate that earlier references to this VarId are valid
	typeCheckArgs.TypeChecker().MarkVarIdAsKnown(this.Identifier())

	return nil
}
