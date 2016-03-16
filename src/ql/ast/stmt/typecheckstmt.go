package stmt

import (
	"fmt"
	"ql/ast/expr"
	"ql/interfaces"
)

func (this Form) TypeCheck(typeChecker interfaces.TypeChecker, symbols interfaces.TypeCheckSymbols) {
	this.Content.TypeCheck(typeChecker, symbols)
}

func (this If) TypeCheck(typeChecker interfaces.TypeChecker, symbols interfaces.TypeCheckSymbols) {
	this.typeCheckIfForNonBoolConditions(typeChecker, symbols)

	typeChecker.AddConditionDependentOn(this.Cond)

	this.Body.TypeCheck(typeChecker, symbols)
	this.Cond.TypeCheck(typeChecker, symbols)

	typeChecker.PopLastConditionDependentOn()
}

func (this IfElse) TypeCheck(typeChecker interfaces.TypeChecker, symbols interfaces.TypeCheckSymbols) {
	this.typeCheckIfElseForNonBoolConditions(typeChecker, symbols)

	typeChecker.AddConditionDependentOn(this.Cond)

	this.Cond.TypeCheck(typeChecker, symbols)
	this.IfBody.TypeCheck(typeChecker, symbols)
	this.ElseBody.TypeCheck(typeChecker, symbols)

	typeChecker.PopLastConditionDependentOn()
}

func (this ComputedQuestion) TypeCheck(typeChecker interfaces.TypeChecker, symbols interfaces.TypeCheckSymbols) {
	typeCheckQuestionForDuplicateLabels(this, typeChecker)
	typeCheckQuestionForRedeclaration(this, typeChecker)

	typeChecker.SetCurrentVarIdVisited(this.VarDecl)

	this.Computation.TypeCheck(typeChecker, symbols)
	this.VarDecl.TypeCheck(typeChecker, symbols)

	for _, conditionDependentOn := range typeChecker.GetConditionsDependentOn() {
		conditionDependentOn.TypeCheck(typeChecker, symbols)
	}

	typeCheckForCyclicDependencies(this, typeChecker, symbols)
	typeChecker.UnsetCurrentVarIdVisited()

}

func (this InputQuestion) TypeCheck(typeChecker interfaces.TypeChecker, symbols interfaces.TypeCheckSymbols) {
	typeCheckQuestionForDuplicateLabels(this, typeChecker)
	typeCheckQuestionForRedeclaration(this, typeChecker)

	typeChecker.SetCurrentVarIdVisited(this.VarDecl)

	this.VarDecl.TypeCheck(typeChecker, symbols)

	// for these condition expressions, running TypeCheck will collect VarIds in them
	for _, conditionDependentOn := range typeChecker.GetConditionsDependentOn() {
		conditionDependentOn.TypeCheck(typeChecker, symbols)
	}

	typeCheckForCyclicDependencies(this, typeChecker, symbols)
	typeChecker.UnsetCurrentVarIdVisited()
}

func (this StmtList) TypeCheck(typeChecker interfaces.TypeChecker, symbols interfaces.TypeCheckSymbols) {
	for _, question := range this.Questions {
		switch question.(type) {
		case InputQuestion:
			question.(InputQuestion).TypeCheck(typeChecker, symbols)
		case ComputedQuestion:
			question.(ComputedQuestion).TypeCheck(typeChecker, symbols)
		}
	}

	for _, conditional := range this.Conditionals {
		switch conditional.(type) {
		case If:
			conditional.(If).TypeCheck(typeChecker, symbols)
		case IfElse:
			conditional.(IfElse).TypeCheck(typeChecker, symbols)
		}
	}
}

func (this If) typeCheckIfForNonBoolConditions(typeChecker interfaces.TypeChecker, symbols interfaces.TypeCheckSymbols) {
	checkForNonBoolCondition(this.GetCondition(), typeChecker, symbols)
}

func (this IfElse) typeCheckIfElseForNonBoolConditions(typeChecker interfaces.TypeChecker, symbols interfaces.TypeCheckSymbols) {
	checkForNonBoolCondition(this.GetCondition(), typeChecker, symbols)
}

func checkForNonBoolCondition(condition interfaces.Expr, typeChecker interfaces.TypeChecker, symbols interfaces.TypeCheckSymbols) {
	typeOfCondition := condition.TypeCheck(typeChecker, symbols)

	if typeOfCondition != expr.NewBoolTypeNoSourceInfo() {
		typeChecker.AddEncounteredErrorForCheckType("NonBoolConditionals", fmt.Errorf("Non-boolean type used as condition: %s", typeOfCondition))
	}
}

func typeCheckForCyclicDependencies(question interfaces.Question, typeChecker interfaces.TypeChecker, symbols interfaces.TypeCheckSymbols) {
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
		typeChecker.AddEncounteredErrorForCheckType("CyclicDependencies", fmt.Errorf("Found cyclic dependency: %s", depencyChainForQuestionId))
	}
}

func typeCheckQuestionForDuplicateLabels(question interfaces.Question, typeChecker interfaces.TypeChecker) {
	labelKnown := typeChecker.IsLabelUsed(question.GetLabel())

	if labelKnown {
		typeChecker.AddEncounteredErrorForCheckType("DuplicateLabels", fmt.Errorf("Label \"%s\" already used for question with identifier %s, using again for question with identifier %s", question.GetLabel(), typeChecker.VarIdForLabel(question.GetLabel()), question.GetVarDecl().GetIdent()))
	} else {
		typeChecker.MarkLabelAsUsed(question.GetLabel(), question.GetVarDecl())
	}
}

func typeCheckQuestionForRedeclaration(question interfaces.Question, typeChecker interfaces.TypeChecker) {
	varDecl := question.GetVarDecl()
	labelKnown := typeChecker.VarDeclIsKnown(varDecl)

	if labelKnown && typeChecker.TypeForVarDecl(varDecl) != varDecl.GetType() {
		typeChecker.AddEncounteredErrorForCheckType("DuplicateVarDeclarations", fmt.Errorf("Question redeclared with different types: %T and %T", varDecl.GetType(), typeChecker.TypeForVarDecl(varDecl)))
	} else {
		typeChecker.MarkVarDeclAsKnown(varDecl)
	}
}
