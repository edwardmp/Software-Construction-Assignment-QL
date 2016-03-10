package stmt

import ("ql/interfaces"
"fmt")

func (f Form) TypeCheck(typeChecker interfaces.TypeChecker, symbolTable interfaces.SymbolTable) {
    f.Content.TypeCheck(typeChecker, symbolTable)
}

func (i If) TypeCheck(typeChecker interfaces.TypeChecker, symbolTable interfaces.SymbolTable) {
    i.typeCheckIfForNonBoolConditions(typeChecker, symbolTable)
    i.Body.TypeCheck(typeChecker, symbolTable)
}

func (i IfElse) TypeCheck(typeChecker interfaces.TypeChecker, symbolTable interfaces.SymbolTable) {
    i.typeCheckIfElseForNonBoolConditions(typeChecker, symbolTable)
    i.IfBody.TypeCheck(typeChecker, symbolTable)
    i.ElseBody.TypeCheck(typeChecker, symbolTable)
}

func (c ComputedQuestion) TypeCheck(typeChecker interfaces.TypeChecker, symbolTable interfaces.SymbolTable) {
    typeCheckQuestionForDuplicateLabels(c, typeChecker)
    typeCheckQuestionForRedeclaration(c, typeChecker)
}

func (i InputQuestion) TypeCheck(typeChecker interfaces.TypeChecker, symbolTable interfaces.SymbolTable) {
    typeCheckQuestionForDuplicateLabels(i, typeChecker)
    typeCheckQuestionForRedeclaration(i, typeChecker)
}

func (s StmtList) TypeCheck(typeChecker interfaces.TypeChecker, symbolTable interfaces.SymbolTable) {
    for _, conditional := range s.Conditionals {
        switch conditional.(type) {
        case If:
            conditional.(If).TypeCheck(typeChecker, symbolTable)
        case IfElse:
            conditional.(IfElse).TypeCheck(typeChecker, symbolTable)
        }
    }

    for _, question := range s.Questions {
        switch question.(type) {
            case InputQuestion:
                question.(InputQuestion).TypeCheck(typeChecker, symbolTable)
            case ComputedQuestion:
                question.(ComputedQuestion).TypeCheck(typeChecker, symbolTable)
        }
    }
}

func (i If) typeCheckIfForNonBoolConditions(typeChecker interfaces.TypeChecker, symbolTable interfaces.SymbolTable) {
    evalCond := i.Cond.Eval(symbolTable)

    if _, CondIsBoolType := evalCond.(bool); !CondIsBoolType {
        typeChecker.AddEncounteredErrorForCheckType("NonBoolConditionals", fmt.Errorf("Non-boolean type used as condition: %T", evalCond))
    }
}

func (i IfElse) typeCheckIfElseForNonBoolConditions(typeChecker interfaces.TypeChecker, symbolTable interfaces.SymbolTable) {
    evalCond := i.Cond.Eval(symbolTable)

    if _, CondIsBoolType := evalCond.(bool); !CondIsBoolType {
        typeChecker.AddEncounteredErrorForCheckType("NonBoolConditionals", fmt.Errorf("Non-boolean type used as condition: %T", evalCond))
    }
}

func typeCheckQuestionForDuplicateLabels(question interfaces.Question, typeChecker interfaces.TypeChecker) {
    labelKnown := typeChecker.IsLabelUsed(question.GetLabel())

    if labelKnown {
        typeChecker.AddEncounteredErrorForCheckType("DuplicateLabels", fmt.Errorf("Label \"%s\" already used for question with identifier %s, using again for question with identifier %s", question.GetLabel(), typeChecker.VarIdForLabel(question.GetLabel()), question.GetVarDecl().Ident))
    } else {
        typeChecker.MarkLabelAsUsed(question.GetLabel(), question.GetVarDecl())
    }
}

func typeCheckQuestionForRedeclaration(question interfaces.Question, typeChecker interfaces.TypeChecker) {
    varDecl := question.GetVarDecl()
    labelKnown := typeChecker.VarDeclIsKnown(varDecl)

    if labelKnown && typeChecker.TypeForVarDecl(varDecl) != varDecl.Type {
        typeChecker.AddEncounteredErrorForCheckType("DuplicateVarDeclarations", fmt.Errorf("Question redeclared with different types: %T and %T", varDecl.Type, typeChecker.TypeForVarDecl(varDecl)))
    } else {
        typeChecker.MarkVarDeclAsKnown(varDecl)
    }
}

