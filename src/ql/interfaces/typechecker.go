package interfaces

import (
    "ql/ast/vari"
    "ql/ast/expr/litexpr")

type TypeChecker interface {
    AddEncounteredErrorForCheckType(string, error)
    GetEncountedErrorsForCheckType(string) []error
    IsLabelUsed(litexpr.StrLit) bool
    MarkLabelAsUsed(litexpr.StrLit, vari.VarDecl)
    VarIdForLabel(litexpr.StrLit) vari.VarId
    VarDeclIsKnown(vari.VarDecl) bool
    MarkVarDeclAsKnown(varDecl vari.VarDecl)
    TypeForVarDecl(varDecl vari.VarDecl) vari.VarType 
}
