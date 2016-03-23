package stmt

import (
	"fmt"
	"ql/ast/expr"
	"ql/interfaces"
)

func (this Form) TypeCheck(typeCheckArgs interfaces.TypeCheckArgs) {
	this.Content.TypeCheck(typeCheckArgs)

	// Only at the end of the form we truly know the question is not declared anywhere (since QL is not a sequential program)
	this.checkForUndefinedReferences(typeCheckArgs.TypeChecker())
}

func (this If) TypeCheck(typeCheckArgs interfaces.TypeCheckArgs) {
	checkForNonBoolCondition(this.GetCondition(), typeCheckArgs)

    this.Cond.TypeCheck(typeCheckArgs)
    typeCheckArgs = typeCheckArgs.AddConditionDependentOn(this.Cond)

	this.Body.TypeCheck(typeCheckArgs)
}

func (this IfElse) TypeCheck(typeCheckArgs interfaces.TypeCheckArgs) {
	checkForNonBoolCondition(this.GetCondition(), typeCheckArgs)

    this.Cond.TypeCheck(typeCheckArgs)
    typeCheckArgs = typeCheckArgs.AddConditionDependentOn(this.Cond)

	this.IfBody.TypeCheck(typeCheckArgs)
	this.ElseBody.TypeCheck(typeCheckArgs)
}

func (this ComputedQuestion) TypeCheck(typeCheckArgs interfaces.TypeCheckArgs) {
	checkQuestionForDuplicateLabels(this, typeCheckArgs.TypeChecker())
	checkQuestionForRedeclaration(this, typeCheckArgs)

	typeCheckArgs = typeCheckArgs.SetCurrentVarDeclVisited(this.VarDecl)

	this.checkIfQuestionTypeMatchesComputationType(typeCheckArgs)

	this.VarDecl.TypeCheck(typeCheckArgs)

    collectVarIdsInExpressions(typeCheckArgs)

	checkForCyclicDependencies(this, typeCheckArgs.TypeChecker())
}

func (this InputQuestion) TypeCheck(typeCheckArgs interfaces.TypeCheckArgs) {
	checkQuestionForDuplicateLabels(this, typeCheckArgs.TypeChecker())
	checkQuestionForRedeclaration(this, typeCheckArgs)

    typeCheckArgs = typeCheckArgs.SetCurrentVarDeclVisited(this.VarDecl)

	this.VarDecl.TypeCheck(typeCheckArgs)

    collectVarIdsInExpressions(typeCheckArgs)

	checkForCyclicDependencies(this, typeCheckArgs.TypeChecker())
}

func (this StmtList) TypeCheck(typeCheckArgs interfaces.TypeCheckArgs) {
	for _, question := range this.Questions {
		question.TypeCheck(typeCheckArgs)
	}

	for _, conditional := range this.Conditionals {
		conditional.TypeCheck(typeCheckArgs)
	}
}

func (this Stmt) TypeCheck(typeCheckArgs interfaces.TypeCheckArgs) {
	panic("Stmt TypeCheck method not overridden")
}

func (this ComputedQuestion) checkIfQuestionTypeMatchesComputationType(typeCheckArgs interfaces.TypeCheckArgs) {
    actualType := this.Computation.TypeCheck(typeCheckArgs)
    expectedType := this.GetVarDecl().GetType()

    // check if question declaration type matches the type of the computation
    if actualType != expr.NewUnknownType() && actualType != expectedType {
        typeCheckArgs.TypeChecker().AddEncounteredError(fmt.Errorf("Encountered computed question with mismatch between declared type (%s) and actual computation type (%s)", expectedType, actualType))
    }
}

// checkForUndefinedReferences looks at all identifiers encountered to see if they have been marked as known
// It should only run at the end of a form (when all identifiers) are collected, hence it's an instance method on Form
func (this Form) checkForUndefinedReferences(typeChecker interfaces.TypeChecker) {
	for identifier, identifierKnown := range typeChecker.GetIdentifiersEncountered() {
		if !identifierKnown {
			typeChecker.AddEncounteredError(fmt.Errorf("Reference to unknown question identifier: %s", identifier))
		}
	}
}

func collectVarIdsInExpressions(typeCheckArgs interfaces.TypeCheckArgs) {
    // for these condition expressions, running TypeCheck will collect VarIds in them and add them as dependencies
    for _, conditionDependentOn := range typeCheckArgs.ConditionsDependentOn() {
        conditionDependentOn.TypeCheck(typeCheckArgs)
    }
}

func checkForNonBoolCondition(condition interfaces.Expr, typeCheckArgs interfaces.TypeCheckArgs) {
	typeOfCondition := condition.TypeCheck(typeCheckArgs)

	if typeOfCondition != expr.NewBoolType() && typeOfCondition != expr.NewUnknownType() {
		typeCheckArgs.TypeChecker().AddEncounteredError(fmt.Errorf("Non-boolean type used as condition: %s", typeOfCondition))
	}
}

func checkForCyclicDependencies(question interfaces.Question, typeChecker interfaces.TypeChecker) {
	// if we find our own VarId as a dependency at least once, the dependencyChain is cyclic
	if typeChecker.DependencyListForVarDeclContainsReferenceToSelf(question.GetVarDecl()) {
		typeChecker.AddEncounteredError(fmt.Errorf("Found cyclic dependency"))
	}
}

func checkQuestionForDuplicateLabels(question interfaces.Question, typeChecker interfaces.TypeChecker) {
	labelKnown := typeChecker.IsLabelUsed(question.GetLabel())

	if labelKnown {
		typeChecker.AddEncounteredWarning(fmt.Errorf("Label \"%s\" already used for question with identifier %s, using again for question with identifier %s", question.GetLabel(), typeChecker.VarIdForLabel(question.GetLabel()), question.GetVarDecl().GetIdent()))
	} else {
		typeChecker.MarkLabelAsUsed(question.GetLabel(), question.GetVarDecl())
	}
}

func checkQuestionForRedeclaration(question interfaces.Question, typeCheckArgs interfaces.TypeCheckArgs) {
	varDecl := question.GetVarDecl()

	if typeCheckArgs.Symbols().IsTypeSetForVarId(varDecl.GetIdent()) && typeCheckArgs.Symbols().GetTypeForVarId(varDecl.GetIdent()) != varDecl.GetType() {
		typeCheckArgs.TypeChecker().AddEncounteredError(fmt.Errorf("Question redeclared with different types: %s and %s", varDecl.GetType(), typeCheckArgs.Symbols().GetTypeForVarId(varDecl.GetIdent())))
	}
}
