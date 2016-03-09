package stmt

import (
    "fmt"
    log "github.com/Sirupsen/logrus"
    "ql/symboltable"
)

type ConditiontypeChecker struct {
    ErrorsEncountered []error
}

func CheckForNonBoolConditions(form Form, symbolTable symboltable.SymbolTable) []error {
    log.Info("Start check for non-boolean conditions")
    conditionTypeChecker := ConditiontypeChecker{}

    form.typeCheck(&conditionTypeChecker, symbolTable)
    log.WithFields(log.Fields{"Errors Encountered": conditionTypeChecker.ErrorsEncountered}).Info("Ended check for non-boolean conditions")

    return conditionTypeChecker.ErrorsEncountered
}

func (f Form) typeCheck(conditionTypeChecker *ConditiontypeChecker, symbolTable symboltable.SymbolTable) {
    f.Content.typeCheck(conditionTypeChecker, symbolTable)
}

func (i If) typeCheck(conditionTypeChecker *ConditiontypeChecker, symbolTable symboltable.SymbolTable) {
    evalCond := i.Cond.Eval(symbolTable)

    if _, CondIsBoolType := evalCond.(bool); !CondIsBoolType {
        conditionTypeChecker.ErrorsEncountered = append(conditionTypeChecker.ErrorsEncountered, fmt.Errorf("Non-boolean type used as condition: %T", evalCond))
    }

    i.Body.typeCheck(conditionTypeChecker, symbolTable)
}

func (i IfElse) typeCheck(conditionTypeChecker *ConditiontypeChecker, symbolTable symboltable.SymbolTable) {
    evalCond := i.Cond.Eval(symbolTable)

    if _, CondIsBoolType := evalCond.(bool); !CondIsBoolType {
        conditionTypeChecker.ErrorsEncountered = append(conditionTypeChecker.ErrorsEncountered, fmt.Errorf("Non-boolean type used as condition: %T", evalCond))
    }

    i.IfBody.typeCheck(conditionTypeChecker, symbolTable)
    i.ElseBody.typeCheck(conditionTypeChecker, symbolTable)
}

func (s StmtList) typeCheck(conditionTypeChecker *ConditiontypeChecker, symbolTable symboltable.SymbolTable) {
    for _, conditional := range s.Conditionals {
        switch conditional.(type) {
        case If:
            conditional.(If).typeCheck(conditionTypeChecker, symbolTable)
        case IfElse:
            conditional.(IfElse).typeCheck(conditionTypeChecker, symbolTable)
        }
    }
}
