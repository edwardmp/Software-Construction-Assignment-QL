package main

import (
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
	"ql/ast/visitor"
	"ql/gui"
	"ql/interfaces"
	"ql/lexer"
	"ql/parser"
	"ql/symbols"
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

	if parsedForm, ok := parseResult.(interfaces.Form); !ok {
		log.Panic("Parse result is not form")
	} else {
		log.WithFields(log.Fields{"Result": parsedForm}).Info("Form parsed")

		visitor := SymbolTableFillerVisitor{}
		symbols := symbols.NewSymbols()
		parsedForm.Accept(&visitor, symbols)

		typeChecker := typechecker.NewTypeChecker()
		parsedForm.TypeCheck(&typeChecker, symbols)

		warnings := typeChecker.GetEncountedWarnings()
		errors := typeChecker.GetEncountedErrors()

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

func (this *SymbolTableFillerVisitor) VisitVarDecl(v interfaces.VarDecl, s interfaces.Symbols) {
	s.SetNodeForIdentifier(v.GetType().GetDefaultValue(), v.GetIdent())
}
