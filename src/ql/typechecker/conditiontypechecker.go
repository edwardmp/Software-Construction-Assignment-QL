package typechecker

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"ql/ast/expr"
	"ql/ast/stmt"
	"ql/ast/visit"
	"ql/symboltable"
)

type ConditionTypeChecker struct {
	visit.Visitor
	ErrorsEncountered []error
}

func CheckForNonBoolConditions(form stmt.Form, symbolTable symboltable.SymbolTable) []error {
	log.Info("Start check for non-boolean conditions")
	ConditionTypeChecker := ConditionTypeChecker{}

	ConditionTypeChecker.Visit(form, symbolTable)
	log.WithFields(log.Fields{"Errors Encountered": ConditionTypeChecker.ErrorsEncountered}).Info("Ended check for non-boolean conditions")

	return ConditionTypeChecker.ErrorsEncountered
}

func (v *ConditionTypeChecker) Visit(t interface{}, s interface{}) interface{} {
	switch t.(type) {
	default:
		log.WithFields(log.Fields{"Node": fmt.Sprintf("%T", t)}).Debug("Ignoring unhandled node type")
	case stmt.Form:
		log.Debug("Visit Form")
		t.(stmt.Form).Identifier.Accept(v, s)
		t.(stmt.Form).Content.Accept(v, s)
	case stmt.StmtList:
		log.Debug("Visit StmtList")

		for _, question := range t.(stmt.StmtList).Questions {
			question.Accept(v, s)
		}

		for _, conditional := range t.(stmt.StmtList).Conditionals {
			conditional.Accept(v, s)
		}
	case stmt.If, stmt.IfElse:
		log.Debug("Visit If")

		var cond expr.Expr
		if _, castAsIfOK := t.(stmt.If); castAsIfOK {
			cond = t.(stmt.If).Cond
			t.(stmt.If).Body.Accept(v, s)
		} else {
			cond = t.(stmt.IfElse).Cond

			t.(stmt.IfElse).IfBody.Accept(v, s)
			t.(stmt.IfElse).ElseBody.Accept(v, s)
		}

		evalCond := cond.Eval(s)

		if _, CondIsBoolType := evalCond.(bool); !CondIsBoolType {
			v.ErrorsEncountered = append(v.ErrorsEncountered, fmt.Errorf("Non-boolean type used as condition: %T", evalCond))
		}

		cond.Accept(v, s)
	}

	return nil
}
