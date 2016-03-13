package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
	"ql/ast"
	"ql/ast/stmt"
	"ql/ast/typechecker"
	"ql/ast/visitor"
	"ql/gui"
	"ql/interfaces"
	"ql/lexer"
	"ql/parser"
	"ql/symbols"
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
		fmt.Println(ast.SourcePosInformation)
		visitor := SymbolTableFillerVisitor{}
		symbols := symboltable.NewSymbols()
		parsedForm.Accept(&visitor, symbols)

		typeChecker := typechecker.NewTypeChecker()
		parsedForm.TypeCheck(&typeChecker, symbols)

		errors := make([]error, 0)
		warnings := make([]error, 0)

		warnings = append(warnings, typeChecker.GetEncountedErrorsForCheckType("DuplicateLabels")...)
		errors = append(errors, typeChecker.GetEncountedErrorsForCheckType("InvalidOperandsDifferentTypes")...)
		errors = append(errors, typeChecker.GetEncountedErrorsForCheckType("ReferenceToUndefinedQuestion")...)
		errors = append(errors, typeChecker.GetEncountedErrorsForCheckType("InvalidOperationOnOperands")...)
		errors = append(errors, typeChecker.GetEncountedErrorsForCheckType("NonBoolConditionals")...)

		log.WithFields(log.Fields{"errors": errors, "warnings": warnings}).Error("Type checking finished")
		gui.CreateGUI(parsedForm, symbols, errors)
	}
}

func initLog() {
	log.SetLevel(log.DebugLevel)
}

type SymbolTableFillerVisitor struct {
	visitor.BaseVisitor
}

func (s *SymbolTableFillerVisitor) VisitVarDecl(va interfaces.VarDecl, sy interface{}) {
	sy.(interfaces.SymbolTable).SetNodeForIdentifier(va.GetType().GetDefaultValue(), va.GetIdent())
}
