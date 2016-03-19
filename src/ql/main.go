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

	gui := gui.NewGUI()

	errors := []error{}

	filePath := "example.ql"
	fileContent, fileError := ioutil.ReadFile(filePath)

	if fileError != nil {
		log.WithFields(log.Fields{"filePath": filePath, "fileError": fileError}).Error("Could not open input file")
		errors = append(errors, fileError)

		// only show the error dialog
		gui.ShowWindow(errors, nil)
	}

	log.WithFields(log.Fields{"filePath": filePath}).Info("Loaded file")

	parsedForm, parseError := lexAndParse(fileContent)

	if parseError != nil {
		errors = append(errors, parseError)
		log.WithFields(log.Fields{"parseError": parseError}).Info("Adding parse error to errors list displayed in GUI")

		// only show the error dialog
		gui.ShowWindow(errors, nil)
		return
	}

	typeCheckErrors, typeCheckWarnings := conductTypeChecking(parsedForm)
	errors = append(errors, typeCheckErrors...)

	varIdDefaultValueVisitor := NewDefaultVarIdSymbolValueVisitor()
	varIdValueSymbols := symbols.NewVarIdValueSymbols()
	parsedForm.Accept(varIdDefaultValueVisitor, varIdValueSymbols)

	gui.InitializeGUIForm(parsedForm, varIdValueSymbols)
	gui.ShowWindow(errors, typeCheckWarnings)
}

func initLog() {
	log.SetLevel(log.DebugLevel)
}

func lexAndParse(fileContent []byte) (interfaces.Form, error) {
	lexer := lexer.NewLexer(fileContent)

	parser := parser.NewParser()
	log.Info("Initiating parsing of file")
	parseResult, parseErr := parser.Parse(lexer)

	if parseErr != nil {
		log.WithFields(log.Fields{"parseErr": parseErr}).Error("Could not parse")
	}

	log.WithFields(log.Fields{"Result": parseResult}).Info("Form parsed")

	parsedForm, parseResultIsForm := parseResult.(interfaces.Form)
	if parseErr == nil && !parseResultIsForm {
		log.Panic("Parse result could not be casted to Form")
	}

	return parsedForm, parseErr
}

func conductTypeChecking(form interfaces.Form) ([]error, []error) {
	typeChecker := typechecker.NewTypeChecker()
	form.TypeCheck(typeChecker, symbols.NewTypeCheckSymbols())
	warnings := typeChecker.GetEncounteredWarnings()
	errors := typeChecker.GetEncounteredErrors()

	log.WithFields(log.Fields{"errors": errors, "warnings": warnings}).Error("Type checking finished")

	return errors, warnings
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
