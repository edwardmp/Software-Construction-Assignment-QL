package typechecker

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"ql/ast/expr"
	"ql/ast/expr/binaryoperatorexpr"
	"ql/ast/expr/unaryoperatorexpr"
	"ql/ast/stmt"
	"ql/ast/visit"
	"ql/symboltable"
)

type UndefinedQuestionReferenceTypeChecker struct {
	visit.Visitor
	ErrorsEncountered []error
}

func CheckForReferencesToUndefinedQuestions(form stmt.Form, symbolTable symboltable.SymbolTable) []error {
	log.Info("Start check for references to undefined questions")
	undefinedQuestionReferenceTypeChecker := UndefinedQuestionReferenceTypeChecker{}

	undefinedQuestionReferenceTypeChecker.Visit(form, symbolTable)
	log.WithFields(log.Fields{"Errors Encountered": undefinedQuestionReferenceTypeChecker.ErrorsEncountered}).Info("Ended check for references to undefined questions")

	return undefinedQuestionReferenceTypeChecker.ErrorsEncountered
}

func (v *UndefinedQuestionReferenceTypeChecker) Visit(t interface{}, s interface{}) interface{} {
	symbolTable := s.(symboltable.SymbolTable)

	switch t.(type) {
	default:
		log.WithFields(log.Fields{"Node": fmt.Sprintf("%T", t)}).Debug("Ignoring unhandled node type")
	case stmt.Form:
		log.Debug("Visit Form")
		t.(stmt.Form).Content.Accept(v, symbolTable)
	case stmt.StmtList:
		log.Debug("Visit StmtList")

		for _, question := range t.(stmt.StmtList).Questions {
			question.Accept(v, symbolTable)
		}

		for _, conditional := range t.(stmt.StmtList).Conditionals {
			conditional.Accept(v, symbolTable)
		}
	case stmt.ComputedQuestion:
		log.Debug("Visit ComputedQuestion")
		t.(stmt.ComputedQuestion).Computation.Accept(v, symbolTable)
	case stmt.If:
		log.Debug("Visit If")
		t.(stmt.If).Cond.Accept(v, symbolTable)
		t.(stmt.If).Body.Accept(v, symbolTable)
	case stmt.IfElse:
		log.Debug("Visit IfElse")
		t.(stmt.IfElse).Cond.Accept(v, symbolTable)
		t.(stmt.IfElse).IfBody.Accept(v, symbolTable)
		t.(stmt.IfElse).ElseBody.Accept(v, symbolTable)
	case binaryoperatorexpr.BinaryOperatorExpr:
		log.Debug("Visit BinaryOperatorExpr")
		t.(binaryoperatorexpr.BinaryOperatorExpr).GetLhs().(expr.Expr).Accept(v, symbolTable)
		t.(binaryoperatorexpr.BinaryOperatorExpr).GetRhs().(expr.Expr).Accept(v, symbolTable)
	case unaryoperatorexpr.UnaryOperatorExpr:
		log.Debug("Visit UnaryOperatorExpr")
		t.(unaryoperatorexpr.UnaryOperatorExpr).GetValue().(expr.Expr).Accept(v, symbolTable)
	case expr.VarExpr:
		log.Debug("Visit VarExpr")

		identifier := t.(expr.VarExpr).GetIdentifier()
		if symbolTable.GetNodeForIdentifier(identifier) == nil {
			v.ErrorsEncountered = append(v.ErrorsEncountered, fmt.Errorf("Reference to unknown question identifier: %s", identifier))
		}
	}

	return nil
}
