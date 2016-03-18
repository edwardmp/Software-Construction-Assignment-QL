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

	filePath := "example.ql"
	fileContent, fileError := ioutil.ReadFile(filePath)

	if fileError != nil {
		log.WithFields(log.Fields{"filePath": filePath, "fileError": fileError}).Panic("Could not open input file")
	}

	log.WithFields(log.Fields{"filePath": filePath}).Info("Loaded file")

	parsedForm := lexAndParse(fileContent)
	typeCheckErrors, typeCheckWarnings := conductTypeChecking(parsedForm)

	varIdDefaultValueVisitor := NewDefaultVarIdSymbolValueVisitor()
	varIdValueSymbols := symbols.NewVarIdValueSymbols()
	parsedForm.Accept(varIdDefaultValueVisitor, varIdValueSymbols)

	gui.CreateGUI(parsedForm, varIdValueSymbols, typeCheckErrors, typeCheckWarnings)
}

func lexAndParse(fileContent []byte) interfaces.Form {
	lexer := lexer.NewLexer(fileContent)

	log.Info("Initiating parsing of file")
	parser := parser.NewParser()
	parseResult, parseErr := parser.Parse(lexer)

	if parseErr != nil {
		log.WithFields(log.Fields{"parseErr": parseErr}).Panic("Could not parse")
	}

	log.WithFields(log.Fields{"Result": parseResult}).Info("Form parsed")

	parsedForm, castAsFormOK := parseResult.(interfaces.Form)
	if !castAsFormOK {
		log.Panic("Parse result could not be casted to Form")
	}

	return parsedForm
}

func conductTypeChecking(form interfaces.Form) ([]error, []error) {
	typeChecker := typechecker.NewTypeChecker()
	form.TypeCheck(typeChecker, symbols.NewTypeCheckSymbols())
	warnings := typeChecker.GetEncounteredWarnings()
	errors := typeChecker.GetEncounteredErrors()

	log.WithFields(log.Fields{"errors": errors, "warnings": warnings}).Error("Type checking finished")

	return errors, warnings
}

func initLog() {
	log.SetLevel(log.DebugLevel)
}

// FIXME place somewhere else?
type DefaultVarIdSymbolValueVisitor struct {
	visitor.BaseVisitor
}

func NewDefaultVarIdSymbolValueVisitor() *DefaultVarIdSymbolValueVisitor {
	return &DefaultVarIdSymbolValueVisitor{}
}

func (this *DefaultVarIdSymbolValueVisitor) VisitVarDecl(varDecl interfaces.VarDecl, context interface{}) {
	symbols := context.(interfaces.VarIdValueSymbols)
	symbols.SetExprForVarId(varDecl.GetType().GetDefaultValue(), varDecl.GetIdent())
}
