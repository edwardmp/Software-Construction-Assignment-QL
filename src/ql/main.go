package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
	"ql/ast/expr"
	"ql/ast/expr/binaryoperatorexpr"
	"ql/ast/expr/litexpr"
	"ql/ast/expr/unaryoperatorexpr"
	"ql/ast/stmt"
	"ql/ast/vari"
	"ql/ast/visit"
	"ql/gui"
	"ql/lexer"
	"ql/parser"
	"ql/symboltable"
	"ql/typechecker"
)

func main() {
	initLog()

	log.Info("Initiating parsing of file")

	fileName := "example.ql"
	qlFile, fileError := ioutil.ReadFile(fileName) // TODO handle error

	if fileError != nil {
		log.WithFields(log.Fields{"fileName": fileName}).Panic("Could not open input file")
	}

	lex := lexer.NewLexer([]byte(string(qlFile)))

	p := parser.NewParser()
	parseResult, parseErr := p.Parse(lex)

	if parseErr != nil {
		log.WithFields(log.Fields{"err": parseErr}).Panic("Could not parse")
	}

	if parsedForm, ok := parseResult.(stmt.Form); !ok {
		log.Panic("Parse result is not form")
	} else {
		log.WithFields(log.Fields{"Result": parsedForm}).Info("Form parsed")

		visitor := VisitorAdapter{}
		symbolTable := symboltable.NewSymbolTable()
		symbolTable = parsedForm.Accept(visitor, symbolTable).(symboltable.SymbolTable)

		errors := make([]error, 0)
		warnings := make([]error, 0)

		warnings = append(warnings, typechecker.CheckForDuplicateLabels(parsedForm)...)
		errors = append(errors, typechecker.CheckForDuplicateVarDeclWithDiffTypes(parsedForm)...)
		errors = append(errors, typechecker.CheckForReferencesToUndefinedQuestions(parsedForm, symbolTable)...)
		errors = append(errors, typechecker.CheckForNonBoolConditions(parsedForm, symbolTable)...)
		errors = append(errors, typechecker.CheckForOperatorsWithInvalidOperands(parsedForm, symbolTable)...)

		log.WithFields(log.Fields{"errors": errors, "warnings": warnings}).Error("Type checking finished")
		gui.CreateGUI(parsedForm, symbolTable)
	}
}

func initLog() {
	log.SetLevel(log.DebugLevel)
}

type VisitorAdapter struct {
	visit.Visitor
}

func (v VisitorAdapter) Visit(t interface{}, s interface{}) interface{} {
	symbolTable := s.(symboltable.SymbolTable)

	switch t.(type) {
	default:
		log.WithFields(log.Fields{"Node": fmt.Sprintf("%T", t)}).Panic("Unexpected node type")
	case stmt.Form:
		log.Debug("Visit Form")
		t.(stmt.Form).Identifier.Accept(v, symbolTable)
		return t.(stmt.Form).Content.Accept(v, symbolTable)
	case vari.VarId:
		log.Debug("Visit VarId")
	case vari.VarType:
		log.Debug("Visit VarType")
	case vari.VarDecl:
		log.Debug("Visit VarDecl")
		varDecl := t.(vari.VarDecl)
		symbolTable.SetNodeForIdentifier(varDecl.GetType().GetDefaultValue(), varDecl.Ident)
		varDecl.Ident.Accept(v, symbolTable)
	case stmt.StmtList:
		log.Debug("Visit StmtList")

		for _, question := range t.(stmt.StmtList).Questions {
			question.Accept(v, symbolTable)
		}

		for _, conditional := range t.(stmt.StmtList).Conditionals {
			conditional.Accept(v, symbolTable)
		}
	case stmt.InputQuestion:
		log.Debug("Visit InputQuestion")
		t.(stmt.InputQuestion).Label.Accept(v, symbolTable)
		t.(stmt.InputQuestion).VarDecl.Accept(v, symbolTable)
	case stmt.ComputedQuestion:
		log.Debug("Visit ComputedQuestion")
		t.(stmt.ComputedQuestion).Label.Accept(v, symbolTable)
		t.(stmt.ComputedQuestion).VarDecl.Accept(v, symbolTable)
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
	case litexpr.StrLit:
		log.Debug("Visit StrLit")
	case litexpr.BoolLit:
		log.Debug("Visit BoolLit")
	case litexpr.IntLit:
		log.Debug("Visit IntLit")
	case binaryoperatorexpr.BinaryOperatorExpr:
		log.Debug("Visit BinaryOperatorExpr")
		t.(binaryoperatorexpr.BinaryOperatorExpr).GetLhs().(expr.Expr).Accept(v, symbolTable)
		t.(binaryoperatorexpr.BinaryOperatorExpr).GetRhs().(expr.Expr).Accept(v, symbolTable)
	case unaryoperatorexpr.UnaryOperatorExpr:
		log.Debug("Visit UnaryOperatorExpr")
		t.(unaryoperatorexpr.UnaryOperatorExpr).GetValue().(expr.Expr).Accept(v, symbolTable)
	case unaryoperatorexpr.VarExpr:
		log.Debug("Visit VarExpr")
		t.(unaryoperatorexpr.VarExpr).GetIdentifier().Accept(v, symbolTable)
	}

	return symbolTable
}
