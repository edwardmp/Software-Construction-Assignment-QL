package stmt

import (
    "fmt"
    log "github.com/Sirupsen/logrus"
    "ql/symboltable"
)

type ConditiontypeCheckNonBoolConditionsNonBoolConditionser struct {
    ErrorsEncountered []error
}

func CheckForNonBoolConditions(form Form, symbolTable symboltable.SymbolTable) []error {
    log.Info("Start check for non-boolean conditions")
    conditionTypeChecker := ConditiontypeCheckNonBoolConditionsNonBoolConditionser{}

    form.typeCheckNonBoolConditionsNonBoolConditions(&conditionTypeChecker, symbolTable)
    log.WithFields(log.Fields{"Errors Encountered": conditionTypeChecker.ErrorsEncountered}).Info("Ended check for non-boolean conditions")

    return conditionTypeChecker.ErrorsEncountered
}

func (f Form) typeCheckNonBoolConditionsNonBoolConditions(conditionTypeChecker *ConditiontypeCheckNonBoolConditionsNonBoolConditionser, symbolTable symboltable.SymbolTable) {
    f.Content.typeCheckNonBoolConditionsNonBoolConditions(conditionTypeChecker, symbolTable)
}

func (i If) typeCheckNonBoolConditionsNonBoolConditions(conditionTypeChecker *ConditiontypeCheckNonBoolConditionsNonBoolConditionser, symbolTable symboltable.SymbolTable) {
    evalCond := i.Cond.Eval(symbolTable)

    if _, CondIsBoolType := evalCond.(bool); !CondIsBoolType {
        conditionTypeChecker.ErrorsEncountered = append(conditionTypeChecker.ErrorsEncountered, fmt.Errorf("Non-boolean type used as condition: %T", evalCond))
    }

    i.Body.typeCheckNonBoolConditionsNonBoolConditions(conditionTypeChecker, symbolTable)
}

func (i IfElse) typeCheckNonBoolConditionsNonBoolConditions(conditionTypeChecker *ConditiontypeCheckNonBoolConditionsNonBoolConditionser, symbolTable symboltable.SymbolTable) {
    evalCond := i.Cond.Eval(symbolTable)

    if _, CondIsBoolType := evalCond.(bool); !CondIsBoolType {
        conditionTypeChecker.ErrorsEncountered = append(conditionTypeChecker.ErrorsEncountered, fmt.Errorf("Non-boolean type used as condition: %T", evalCond))
    }

    i.IfBody.typeCheckNonBoolConditionsNonBoolConditions(conditionTypeChecker, symbolTable)
    i.ElseBody.typeCheckNonBoolConditionsNonBoolConditions(conditionTypeChecker, symbolTable)
}

func (s StmtList) typeCheckNonBoolConditionsNonBoolConditions(conditionTypeChecker *ConditiontypeCheckNonBoolConditionsNonBoolConditionser, symbolTable symboltable.SymbolTable) {
    for _, conditional := range s.Conditionals {
        switch conditional.(type) {
        case If:
            conditional.(If).typeCheckNonBoolConditionsNonBoolConditions(conditionTypeChecker, symbolTable)
        case IfElse:
            conditional.(IfElse).typeCheckNonBoolConditionsNonBoolConditions(conditionTypeChecker, symbolTable)
        }
    }
}
