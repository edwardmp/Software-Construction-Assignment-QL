package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"io/ioutil"
	"os"
	"ql/ast/visitor"
	"ql/gui"
	"ql/interfaces"
	"ql/lexer"
	"ql/parser"
	"ql/symbols"
	"ql/typechecker"
)

func main() {
	app := cli.NewApp()
	app.Name = "QL"
	app.Usage = "generate a questionnaire form from a DSL"
	app.Version = "1.0.0"

	var filePathPassed string
	var debugModeEnabled bool

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "filepath",
			Value:       "form.ql",
			Usage:       "path to a .ql file",
			Destination: &filePathPassed,
		},
		cli.BoolFlag{
			Name:        "debugmode",
			Usage:       "if passed, the minimum log level displayed is set to debug",
			Destination: &debugModeEnabled,
		},
	}

	app.Action = func(c *cli.Context) {
		filePath := filePathPassed
		if c.NArg() > 0 {
			filePath = c.Args()[0]
		}

		if debugModeEnabled {
			enableDebugLevelLogs()
		} else {
			enableFatalLevelLogs()
		}

		log.WithFields(log.Fields{"filePath": filePath}).Info("Passed file path as CLI argument/flag")

		// no file at path exists
		if src, err := os.Stat(filePath); os.IsNotExist(err) {
			fmt.Printf("Could not continue: no file exists at path: %s \n", filePath)
			return
		} else if src.IsDir() {
			fmt.Printf("Could not continue: path to directory instead of file passed: %s\n", filePath)
			return
		}

		initQL(filePath)
	}

	app.Run(os.Args)
}

// enableDebugLevelLogs sets the minimum level of logs displayed to debug
func enableDebugLevelLogs() {
	log.SetLevel(log.DebugLevel)
	log.Info("Enabled debug level logging")
}

// enableDebugLevelLogs sets the minimum level of logs displayed to fatal
func enableFatalLevelLogs() {
	log.SetLevel(log.FatalLevel)
}

// initQL loads the contents of the file at the passed file path, starts lexer/parser and type checking and presents the GUI
func initQL(filePath string) {
	gui := gui.NewGUI()

	fileContent, _ := ioutil.ReadFile(filePath)

	log.WithFields(log.Fields{"filePath": filePath}).Info("Loaded file")

	parsedForm, parseError := lexAndParse(fileContent)

	if parseError != nil {
		errors := []error{parseError}
		log.WithFields(log.Fields{"parseError": parseError}).Info("Adding parse error to errors list displayed in GUI")

		// only show the error dialog
		gui.ShowWindow(errors, nil)
		return
	}

	typeCheckErrors, typeCheckWarnings := conductTypeChecking(parsedForm)

	varIdDefaultValueVisitor := NewDefaultVarIdSymbolValueVisitor()
	varIdValueSymbols := symbols.NewVarIdValueSymbols()
	parsedForm.Accept(varIdDefaultValueVisitor, varIdValueSymbols)

	gui.InitializeGUIForm(parsedForm, varIdValueSymbols)
	gui.ShowWindow(typeCheckErrors, typeCheckWarnings)
}

// lexAndParse conducts lexical analysis and parses the past file content, returns a parsed form and a potential error
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

// conductTypeChecking starts type checking the passed form for errors and warnings
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
