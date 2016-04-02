package stmt

import (
	"fmt"
	"ql/ast/expr"
	"ql/interfaces"
)

func (this Form) TypeCheck(typeCheckArgs interfaces.TypeCheckArgs) {
	this.Content().TypeCheck(typeCheckArgs)

	// Only at the end of the form we truly know the question is not declared anywhere (since QL is not a sequential program)
	this.checkForUndefinedReferences(typeCheckArgs.TypeChecker())
}

func (this If) TypeCheck(typeCheckArgs interfaces.TypeCheckArgs) {
	checkForNonBoolCondition(this.Condition(), typeCheckArgs)

	this.Condition().TypeCheck(typeCheckArgs)
	typeCheckArgs = typeCheckArgs.AddConditionDependentOn(this.Condition())

	this.Body().TypeCheck(typeCheckArgs)
}

func (this IfElse) TypeCheck(typeCheckArgs interfaces.TypeCheckArgs) {
	checkForNonBoolCondition(this.Condition(), typeCheckArgs)

	this.Condition().TypeCheck(typeCheckArgs)
	typeCheckArgs = typeCheckArgs.AddConditionDependentOn(this.Condition())

	this.IfBody().TypeCheck(typeCheckArgs)
	this.ElseBody().TypeCheck(typeCheckArgs)
}

func (this ComputedQuestion) TypeCheck(typeCheckArgs interfaces.TypeCheckArgs) {
	checkQuestionForDuplicateLabels(this, typeCheckArgs.TypeChecker())
	checkQuestionForRedeclarationWithDifferentTypes(this, typeCheckArgs)

	typeCheckArgs = typeCheckArgs.SetCurrentVarDeclVisited(this.VarDecl())

	this.checkIfQuestionTypeMatchesComputationType(typeCheckArgs)

	this.VarDecl().TypeCheck(typeCheckArgs)

	collectVarIdsInExpressions(typeCheckArgs)

	checkForCyclicDependencies(this, typeCheckArgs.TypeChecker())
}

func (this InputQuestion) TypeCheck(typeCheckArgs interfaces.TypeCheckArgs) {
	checkQuestionForDuplicateLabels(this, typeCheckArgs.TypeChecker())
	checkQuestionForRedeclarationWithDifferentTypes(this, typeCheckArgs)

	typeCheckArgs = typeCheckArgs.SetCurrentVarDeclVisited(this.VarDecl())

	this.VarDecl().TypeCheck(typeCheckArgs)

	collectVarIdsInExpressions(typeCheckArgs)

	checkForCyclicDependencies(this, typeCheckArgs.TypeChecker())
}

func (this StmtList) TypeCheck(typeCheckArgs interfaces.TypeCheckArgs) {
	for _, question := range this.Questions() {
		question.TypeCheck(typeCheckArgs)
	}

	for _, conditional := range this.Conditionals() {
		conditional.TypeCheck(typeCheckArgs)
	}
}

func (this Stmt) TypeCheck(typeCheckArgs interfaces.TypeCheckArgs) {
	panic("Stmt TypeCheck method not overridden")
}

// checkIfQuestionTypeMatchesComputationType checks if the declared computed question type and its actual type match, and if not, adds an error to the typechecker
func (this ComputedQuestion) checkIfQuestionTypeMatchesComputationType(typeCheckArgs interfaces.TypeCheckArgs) {
	actualType := this.Computation().TypeCheck(typeCheckArgs)
	expectedType := this.VarDecl().Type()

	// check if question declaration type matches the type of the computation
	if actualType != expr.NewUnknownType() && actualType != expectedType {
		typeCheckArgs.TypeChecker().AddEncounteredError(fmt.Errorf("Encountered computed question with mismatch between declared type (%s) and actual computation type (%s)", expectedType, actualType))
	}
}

// checkForUndefinedReferences seeks to confirm that all identifiers encountered have been marked as known (being declared)
func (this Form) checkForUndefinedReferences(typeChecker interfaces.TypeChecker) {
	for identifier, identifierKnown := range typeChecker.IdentifiersEncountered() {
		if !identifierKnown {
			typeChecker.AddEncounteredError(fmt.Errorf("Reference to unknown question identifier: %s", identifier))
		}
	}
}

// FIXME side-effect
func collectVarIdsInExpressions(typeCheckArgs interfaces.TypeCheckArgs) {
	// for these condition expressions, running TypeCheck will collect VarIds in them and add them as dependencies
	for _, conditionDependentOn := range typeCheckArgs.ConditionsDependentOn() {
		conditionDependentOn.TypeCheck(typeCheckArgs)
	}
}

// checkForNonBoolCondition checks if the condition is of a boolean type, and if not, adds an error to the typechecker
func checkForNonBoolCondition(condition interfaces.Expr, typeCheckArgs interfaces.TypeCheckArgs) {
	typeOfCondition := condition.TypeCheck(typeCheckArgs)

	if typeOfCondition != expr.NewBoolType() && typeOfCondition != expr.NewUnknownType() {
		typeCheckArgs.TypeChecker().AddEncounteredError(fmt.Errorf("Non-boolean type used as condition: %s", typeOfCondition))
	}
}

// checkForCyclicDependencies checks if the this question depends on itself, and if so, adds an error to the typechecker
func checkForCyclicDependencies(question interfaces.Question, typeChecker interfaces.TypeChecker) {
	// if we find our own VarId as a dependency at least once, the dependencyChain is cyclic
	if typeChecker.DependencyListForVarDeclContainsReferenceToSelf(question.VarDecl()) {
		typeChecker.AddEncounteredError(fmt.Errorf("Found cyclic dependency"))
	}
}

// checkQuestionForDuplicateLabels checks if there are multiple questions with the same label, and if so, adds an warning to the typechecker
func checkQuestionForDuplicateLabels(question interfaces.Question, typeChecker interfaces.TypeChecker) {
	labelKnown := typeChecker.IsLabelUsed(question.Label())

	if labelKnown {
		typeChecker.AddEncounteredWarning(fmt.Errorf("Label \"%s\" already used for question with identifier %s, using again for question with identifier %s", question.Label(), typeChecker.VarIdForLabel(question.Label()), question.VarDecl().Identifier()))
	} else {
		typeChecker.MarkLabelAsUsed(question.Label(), question.VarDecl())
	}
}

// checkQuestionForRedeclaration checks if the passed question has been declared before with a different type, and if so, adds an error to the typechecker
func checkQuestionForRedeclarationWithDifferentTypes(question interfaces.Question, typeCheckArgs interfaces.TypeCheckArgs) {
	varDecl := question.VarDecl()

	if typeCheckArgs.Symbols().IsTypeSetForVarId(varDecl.Identifier()) && typeCheckArgs.Symbols().TypeForVarId(varDecl.Identifier()) != varDecl.Type() {
		typeCheckArgs.TypeChecker().AddEncounteredError(fmt.Errorf("Question redeclared with different types: %s and %s", varDecl.Type(), typeCheckArgs.Symbols().TypeForVarId(varDecl.Identifier())))
	}
}
