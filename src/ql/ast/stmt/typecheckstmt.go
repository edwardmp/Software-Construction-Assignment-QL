package stmt

import (
	"fmt"
	"ql/ast/expr"
	"ql/interfaces"
)

func (this Form) TypeCheck(typeChecker interfaces.TypeChecker, symbols interfaces.TypeCheckSymbols) {
	this.Content.TypeCheck(typeChecker, symbols)

	// If at the end of the form we still have undefined references, they will be added as errors
	// Only at the end of the form we truly know the question is not declared anywhere
	this.checkForUndefinedReferences(typeChecker)
}

func (this If) TypeCheck(typeChecker interfaces.TypeChecker, symbols interfaces.TypeCheckSymbols) {
	checkForNonBoolCondition(this.GetCondition(), typeChecker, symbols)

	typeChecker.AddConditionDependentOn(this.Cond)

	this.Body.TypeCheck(typeChecker, symbols)
	this.Cond.TypeCheck(typeChecker, symbols)

	typeChecker.PopLastConditionDependentOn()
}

func (this IfElse) TypeCheck(typeChecker interfaces.TypeChecker, symbols interfaces.TypeCheckSymbols) {
	checkForNonBoolCondition(this.GetCondition(), typeChecker, symbols)

	typeChecker.AddConditionDependentOn(this.Cond)

	this.Cond.TypeCheck(typeChecker, symbols)
	this.IfBody.TypeCheck(typeChecker, symbols)
	this.ElseBody.TypeCheck(typeChecker, symbols)

	typeChecker.PopLastConditionDependentOn()
}

func (this ComputedQuestion) TypeCheck(typeChecker interfaces.TypeChecker, symbols interfaces.TypeCheckSymbols) {
	checkQuestionForDuplicateLabels(this, typeChecker)
	checkQuestionForRedeclaration(this, typeChecker, symbols)
	typeChecker.SetCurrentVarIdVisited(this.VarDecl)

	this.checkIfQuestionTypeMatchesComputationType(typeChecker, symbols)

	this.VarDecl.TypeCheck(typeChecker, symbols)

	for _, conditionDependentOn := range typeChecker.GetConditionsDependentOn() {
		conditionDependentOn.TypeCheck(typeChecker, symbols)
	}

	checkForCyclicDependencies(this, typeChecker, symbols)

	typeChecker.UnsetCurrentVarIdVisited()
}

func (this InputQuestion) TypeCheck(typeChecker interfaces.TypeChecker, symbols interfaces.TypeCheckSymbols) {
	checkQuestionForDuplicateLabels(this, typeChecker)
	checkQuestionForRedeclaration(this, typeChecker, symbols)

	typeChecker.SetCurrentVarIdVisited(this.VarDecl)

	this.VarDecl.TypeCheck(typeChecker, symbols)

	// for these condition expressions, running TypeCheck will collect VarIds in them
	for _, conditionDependentOn := range typeChecker.GetConditionsDependentOn() {
		conditionDependentOn.TypeCheck(typeChecker, symbols)
	}

	checkForCyclicDependencies(this, typeChecker, symbols)

	typeChecker.UnsetCurrentVarIdVisited()
}

func (this StmtList) TypeCheck(typeChecker interfaces.TypeChecker, symbols interfaces.TypeCheckSymbols) {
	for _, question := range this.Questions {
		question.TypeCheck(typeChecker, symbols)
	}

	for _, conditional := range this.Conditionals {
		conditional.TypeCheck(typeChecker, symbols)
	}
}

func (this Stmt) TypeCheck(typeChecker interfaces.TypeChecker, symbols interfaces.TypeCheckSymbols) {
	panic("Stmt TypeCheck method not overridden")
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

func checkForNonBoolCondition(condition interfaces.Expr, typeChecker interfaces.TypeChecker, symbols interfaces.TypeCheckSymbols) {
	typeOfCondition := condition.TypeCheck(typeChecker, symbols)

	if typeOfCondition != expr.NewBoolTypeNoSourceInfo() && typeOfCondition != expr.NewUnknownType() {
		typeChecker.AddEncounteredError(fmt.Errorf("Non-boolean type used as condition: %s", typeOfCondition))
	}
}

func checkForCyclicDependencies(question interfaces.Question, typeChecker interfaces.TypeChecker, symbols interfaces.TypeCheckSymbols) {
	varIdOfCurrentlyVisitingQuestion := question.GetVarDecl().GetIdent()

	numOfTimesQuestionIdFound := 0
	depencyChainForQuestionId := typeChecker.GetDependencyChainForVarId(varIdOfCurrentlyVisitingQuestion)
	for _, dependingVarId := range depencyChainForQuestionId {
		if dependingVarId == varIdOfCurrentlyVisitingQuestion {
			numOfTimesQuestionIdFound++
		}
	}

	// if we find our own var id more than once, the dependencyChain is cyclic
	if numOfTimesQuestionIdFound >= 2 {
		typeChecker.AddEncounteredError(fmt.Errorf("Found cyclic dependency: %s", depencyChainForQuestionId))
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

func checkQuestionForRedeclaration(question interfaces.Question, typeChecker interfaces.TypeChecker, symbols interfaces.TypeCheckSymbols) {
	varDecl := question.GetVarDecl()

	if symbols.IsTypeSetForVarId(varDecl.GetIdent()) && symbols.GetTypeForVarId(varDecl.GetIdent()) != varDecl.GetType() {
		typeChecker.AddEncounteredError(fmt.Errorf("Question redeclared with different types: %s and %s", varDecl.GetType(), symbols.GetTypeForVarId(varDecl.GetIdent())))
	}
}

func (this ComputedQuestion) checkIfQuestionTypeMatchesComputationType(typeChecker interfaces.TypeChecker, symbols interfaces.TypeCheckSymbols) {
	actualType := this.Computation.TypeCheck(typeChecker, symbols)
	expectedType := this.GetVarDecl().GetType()

	// check if question declaration type matches the type of the computation
	if actualType != expr.NewUnknownType() && actualType != expectedType {
		typeChecker.AddEncounteredError(fmt.Errorf("Encountered computed question with mismatch between declared type (%s) and actual computation type (%s)", expectedType, actualType))
	}
}
