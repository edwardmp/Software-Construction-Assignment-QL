package typechecker

import ("ql/interfaces"
    log "github.com/Sirupsen/logrus"
    "ql/ast/expr/litexpr"
    "ql/ast/vari")

type TypeChecker struct {
    interfaces.TypeChecker
	ErrorsEncounteredForCheckType map[string][]error
    UsedLabels map[litexpr.StrLit]vari.VarId
    KnownIdentifiers map[vari.VarId]vari.VarType
}

func NewTypeChecker() TypeChecker {
    return TypeChecker{ErrorsEncounteredForCheckType: make(map[string][]error), UsedLabels: make(map[litexpr.StrLit]vari.VarId), KnownIdentifiers: make(map[vari.VarId]vari.VarType)}
}

func (t *TypeChecker) AddEncounteredErrorForCheckType(checkType string, encounteredError error) {
    log.WithFields(log.Fields{"checkType": checkType, "errorEncountered": encounteredError}).Info("Added encountered error for check type")

    t.ErrorsEncounteredForCheckType[checkType] = append(t.ErrorsEncounteredForCheckType[checkType], encounteredError) 
}

func (t *TypeChecker) GetEncountedErrorsForCheckType(checkType string) []error {
    return t.ErrorsEncounteredForCheckType[checkType]
}

func (t *TypeChecker) IsLabelUsed(label litexpr.StrLit) bool {
    if _, exists := t.UsedLabels[label]; exists {
        return true
    }

    return false
}

func (t *TypeChecker) VarIdForLabel(label litexpr.StrLit) vari.VarId {
    return t.UsedLabels[label]
}

func (t *TypeChecker) MarkLabelAsUsed(label litexpr.StrLit, varDecl vari.VarDecl) {
    log.WithFields(log.Fields{"label": label}).Debug("Marking label as used")
    t.UsedLabels[label] = varDecl.Ident
}


func (t *TypeChecker) VarDeclIsKnown(varDecl vari.VarDecl) bool {
    if _, exists := t.KnownIdentifiers[varDecl.Ident]; exists {
        return true
    }

    return false
}

func (t *TypeChecker) MarkVarDeclAsKnown(varDecl vari.VarDecl) {
    log.WithFields(log.Fields{"VarDecl": varDecl}).Debug("Marking VarDecl as known")
    t.KnownIdentifiers[varDecl.Ident] = varDecl.Type
}

func (t *TypeChecker) TypeForVarDecl(varDecl vari.VarDecl) vari.VarType {
    return t.KnownIdentifiers[varDecl.Ident]
}


