package stmt

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"ql/interfaces"
)

func (this Form) TypeCheck(typeChecker interfaces.TypeChecker, symbols interfaces.Symbols) {
	this.Content.TypeCheck(typeChecker, symbols)
}

func (this If) TypeCheck(typeChecker interfaces.TypeChecker, symbols interfaces.Symbols) {
	this.typeCheckIfForNonBoolConditions(typeChecker, symbols)

	typeChecker.AddConditionDependentOn(this.Cond)

	this.Body.TypeCheck(typeChecker, symbols)
	this.Cond.TypeCheck(typeChecker, symbols)

	typeChecker.PopLastConditionDependentOn()
}

func (this IfElse) TypeCheck(typeChecker interfaces.TypeChecker, symbols interfaces.Symbols) {
	this.typeCheckIfElseForNonBoolConditions(typeChecker, symbols)

	typeChecker.AddConditionDependentOn(this.Cond)

	this.Cond.TypeCheck(typeChecker, symbols)
	this.IfBody.TypeCheck(typeChecker, symbols)
	this.ElseBody.TypeCheck(typeChecker, symbols)

	typeChecker.PopLastConditionDependentOn()
}

func (this ComputedQuestion) TypeCheck(typeChecker interfaces.TypeChecker, symbols interfaces.Symbols) {
	typeCheckQuestionForDuplicateLabels(this, typeChecker)
	typeCheckQuestionForRedeclaration(this, typeChecker)

	typeChecker.SetCurrentVarIdVisited(this.VarDecl)

	this.Computation.TypeCheck(typeChecker, symbols)
	this.VarDecl.TypeCheck(typeChecker, symbols)

	for _, conditionDependentOn := range typeChecker.GetConditionsDependentOn() {
		conditionDependentOn.TypeCheck(typeChecker, symbols)
	}

	typeCheckForCyclicalDependencies(this, typeChecker, symbols)
	typeChecker.UnsetCurrentVarIdVisited()

}

func (this InputQuestion) TypeCheck(typeChecker interfaces.TypeChecker, symbols interfaces.Symbols) {
	typeCheckQuestionForDuplicateLabels(this, typeChecker)
	typeCheckQuestionForRedeclaration(this, typeChecker)

	typeChecker.SetCurrentVarIdVisited(this.VarDecl)

	this.VarDecl.TypeCheck(typeChecker, symbols)

	// for these condition expressions, running TypeCheck will collect VarIds in them
	for _, conditionDependentOn := range typeChecker.GetConditionsDependentOn() {
		conditionDependentOn.TypeCheck(typeChecker, symbols)
	}

	typeCheckForCyclicalDependencies(this, typeChecker, symbols)
	typeChecker.UnsetCurrentVarIdVisited()
}

func (this StmtList) TypeCheck(typeChecker interfaces.TypeChecker, symbols interfaces.Symbols) {
	for _, conditional := range this.Conditionals {
		switch conditional.(type) {
		case If:
			conditional.(If).TypeCheck(typeChecker, symbols)
		case IfElse:
			conditional.(IfElse).TypeCheck(typeChecker, symbols)
		}
	}

	for _, question := range this.Questions {
		switch question.(type) {
		case InputQuestion:
			question.(InputQuestion).TypeCheck(typeChecker, symbols)
		case ComputedQuestion:
			question.(ComputedQuestion).TypeCheck(typeChecker, symbols)
		}
	}
}

func typeCheckForCyclicalDependencies(question interfaces.Question, typeChecker interfaces.TypeChecker, symbols interfaces.Symbols) {
	varIdOfCurrentlyVisitingQuestion := question.GetVarDecl().GetIdent()

	numOfTimesQuestionVarIdFound := 0
	depencyChainForQuestionVarId := typeChecker.GetDependencyChainForVarId(varIdOfCurrentlyVisitingQuestion)
	for _, dependingVarId := range depencyChainForQuestionVarId {
		if dependingVarId == varIdOfCurrentlyVisitingQuestion {
			numOfTimesQuestionVarIdFound++
		}
	}

	// if we find our own var id more than once, the dependencyChain is cyclical
	if numOfTimesQuestionVarIdFound >= 2 {
		log.WithFields(log.Fields{"Cyclical": depencyChainForQuestionVarId}).Error("Cyclic dependency found")
		typeChecker.AddEncounteredErrorForCheckType("CyclicalDependencies", fmt.Errorf("Found cyclical dependency: %s", depencyChainForQuestionVarId))
	}
}

func (this If) typeCheckIfForNonBoolConditions(typeChecker interfaces.TypeChecker, symbols interfaces.Symbols) {
	evalCond := this.Cond.Eval(symbols)

	if _, CondIsBoolType := evalCond.(bool); !CondIsBoolType {
		typeChecker.AddEncounteredErrorForCheckType("NonBoolConditionals", fmt.Errorf("Non-boolean type used as condition: %T", evalCond))
	}
}

func (this IfElse) typeCheckIfElseForNonBoolConditions(typeChecker interfaces.TypeChecker, symbols interfaces.Symbols) {
	evalCond := this.Cond.Eval(symbols)

	if _, CondIsBoolType := evalCond.(bool); !CondIsBoolType {
		typeChecker.AddEncounteredErrorForCheckType("NonBoolConditionals", fmt.Errorf("Non-boolean type used as condition: %T", evalCond))
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
