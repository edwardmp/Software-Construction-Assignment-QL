package main

import (
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
	"ql/ast/stmt"
	"ql/ast/typechecker"
	"ql/gui"
	"ql/interfaces"
	"ql/lexer"
	"ql/parser"
	"ql/symboltable"
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

		visitor := SymbolTableFillerVisitor{}
		symbolTable := symboltable.NewSymbolTable()
		parsedForm.Accept(visitor, symbolTable)

		typeChecker := typechecker.NewTypeChecker()
		parsedForm.TypeCheck(&typeChecker, symbolTable)

		errors := make([]error, 0)
		warnings := make([]error, 0)

		warnings = append(warnings, typeChecker.GetEncountedErrorsForCheckType("DuplicateLabels")...)
		errors = append(errors, typeChecker.GetEncountedErrorsForCheckType("InvalidOperandsDifferentTypes")...)
		//errors = append(errors, typechecker.CheckForReferencesToUndefinedQuestions(parsedForm, symbolTable)...)
		errors = append(errors, typeChecker.GetEncountedErrorsForCheckType("InvalidOperationOnOperands")...)
		errors = append(errors, typeChecker.GetEncountedErrorsForCheckType("NonBoolConditionals")...)

		log.WithFields(log.Fields{"errors": errors, "warnings": warnings}).Error("Type checking finished")
		gui.CreateGUI(parsedForm, symbolTable)
	}
}

func initLog() {
	log.SetLevel(log.DebugLevel)
}

type SymbolTableFillerVisitor struct {
	interfaces.Visitor
}

func (s SymbolTableFillerVisitor) VisitAdd(a interfaces.Add, sy interface{}) {

}

func (s SymbolTableFillerVisitor) VisitAnd(a interfaces.And, sy interface{}) {

}

func (s SymbolTableFillerVisitor) VisitDiv(d interfaces.Div, sy interface{}) {

}

func (s SymbolTableFillerVisitor) VisitEq(e interfaces.Eq, sy interface{}) {

}

func (s SymbolTableFillerVisitor) VisitGEq(g interfaces.GEq, sy interface{}) {

}

func (s SymbolTableFillerVisitor) VisitGT(g interfaces.GT, sy interface{}) {

}

func (s SymbolTableFillerVisitor) VisitLEq(l interfaces.LEq, sy interface{}) {

}

func (s SymbolTableFillerVisitor) VisitLT(l interfaces.LT, sy interface{}) {

}

func (s SymbolTableFillerVisitor) VisitMul(m interfaces.Mul, sy interface{}) {

}

func (s SymbolTableFillerVisitor) VisitNEq(n interfaces.NEq, sy interface{}) {

}

func (s SymbolTableFillerVisitor) VisitOr(o interfaces.Or, sy interface{}) {

}

func (s SymbolTableFillerVisitor) VisitSub(su interfaces.Sub, sy interface{}) {

}

func (s SymbolTableFillerVisitor) VisitBoolLit(b interfaces.BoolLit, sy interface{}) {

}

func (s SymbolTableFillerVisitor) VisitIntLit(i interfaces.IntLit, sy interface{}) {

}

func (s SymbolTableFillerVisitor) VisitStrLit(st interfaces.StrLit, sy interface{}) {

}

func (s SymbolTableFillerVisitor) VisitNeg(n interfaces.Neg, sy interface{}) {

}

func (s SymbolTableFillerVisitor) VisitNot(n interfaces.Not, sy interface{}) {

}

func (s SymbolTableFillerVisitor) VisitPos(p interfaces.Pos, sy interface{}) {

}

func (s SymbolTableFillerVisitor) VisitVarExpr(va interfaces.VarExpr, sy interface{}) {

}

func (s SymbolTableFillerVisitor) VisitVarDecl(va interfaces.VarDecl, sy interface{}) {
	sy.(symboltable.SymbolTable).SetNodeForIdentifier(va.GetType().GetDefaultValue(), va.GetIdent())
}

func (s SymbolTableFillerVisitor) VisitVarId(va interfaces.VarId, sy interface{}) {

}

func (s SymbolTableFillerVisitor) VisitIntType(i interfaces.IntType, sy interface{}) {

}

func (s SymbolTableFillerVisitor) VisitBoolType(b interfaces.BoolType, sy interface{}) {

}

func (s SymbolTableFillerVisitor) VisitStringType(st interfaces.StringType, sy interface{}) {

}

func (s SymbolTableFillerVisitor) VisitForm(f interfaces.Form, sy interface{}) {

}

func (s SymbolTableFillerVisitor) VisitComputedQuestion(c interfaces.ComputedQuestion, sy interface{}) {

}

func (s SymbolTableFillerVisitor) VisitInputQuestion(i interfaces.InputQuestion, sy interface{}) {

}

func (s SymbolTableFillerVisitor) VisitIf(i interfaces.If, sy interface{}) {

}

func (s SymbolTableFillerVisitor) VisitIfElse(i interfaces.IfElse, sy interface{}) {

}

func (s SymbolTableFillerVisitor) VisitStmtList(st interfaces.StmtList, sy interface{}) {

}
