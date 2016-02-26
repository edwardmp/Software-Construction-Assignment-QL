package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
	"ql/ast/expr"
	"ql/ast/expr/binaryoperatorexpr"
	"ql/ast/expr/lit"
	"ql/ast/expr/unaryoperatorexpr"
	"ql/ast/stmt"
	"ql/ast/vari"
	"ql/ast/vari/vartype"
	"ql/ast/visit"
	"ql/symboltable"
	"ql/gui"
	"ql/lexer"
	"ql/parser"
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
		symbolTableStack := symboltable.NewSymbolTable()

		symbolTable := parsedForm.Accept(visitor, symbolTableStack).(symboltable.SymbolTable)

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
	stack := s.(symboltable.SymbolTable)

	switch t.(type) {
	default:
		log.WithFields(log.Fields{"Node": fmt.Sprintf("%T", t)}).Panic("Unexpected node type")
	case stmt.Form:
		log.Debug("Visit Form")
		t.(stmt.Form).Identifier.Accept(v, stack)
		return t.(stmt.Form).Content.Accept(v, stack)
	case vari.VarId:
		log.Debug("Visit VarId")
	case vartype.VarType:
		log.Debug("Visit VarType")
	case vari.VarDecl:
		log.Debug("Visit VarDecl")
		varDecl := t.(vari.VarDecl)
		stack.SetNodeForIdentifier(varDecl.GetType().GetDefaultValue(), varDecl.Ident)
		varDecl.Ident.Accept(v, stack)
	case stmt.StmtList:
		log.Debug("Visit StmtList")

		for _, question := range t.(stmt.StmtList).Questions {
			question.Accept(v, stack)
		}

		for _, conditional := range t.(stmt.StmtList).Conditionals {
			conditional.Accept(v, stack)
		}
	case stmt.InputQuestion:
		log.Debug("Visit InputQuestion")
		t.(stmt.InputQuestion).Label.Accept(v, stack)
		t.(stmt.InputQuestion).VarDecl.Accept(v, stack)
	case stmt.ComputedQuestion:
		log.Debug("Visit ComputedQuestion")
		t.(stmt.ComputedQuestion).Label.Accept(v, stack)
		t.(stmt.ComputedQuestion).VarDecl.Accept(v, stack)
		t.(stmt.ComputedQuestion).Computation.Accept(v, stack)
	case stmt.If:
		log.Debug("Visit If")
		t.(stmt.If).Cond.Accept(v, stack)
		t.(stmt.If).Body.Accept(v, stack)
	case stmt.IfElse:
		log.Debug("Visit IfElse")
		t.(stmt.IfElse).Cond.Accept(v, stack)
		t.(stmt.IfElse).IfBody.Accept(v, stack)
		t.(stmt.IfElse).ElseBody.Accept(v, stack)
	case lit.StrLit:
		log.Debug("Visit StrLit")
	case lit.BoolLit:
		log.Debug("Visit BoolLit")
	case lit.IntLit:
		log.Debug("Visit IntLit")
	case binaryoperatorexpr.BinaryOperatorExpr:
		log.Debug("Visit BinaryOperatorExpr")
		t.(binaryoperatorexpr.BinaryOperatorExpr).GetLhs().(expr.Expr).Accept(v, stack)
		t.(binaryoperatorexpr.BinaryOperatorExpr).GetRhs().(expr.Expr).Accept(v, stack)
	case unaryoperatorexpr.UnaryOperatorExpr:
		log.Debug("Visit UnaryOperatorExpr")
		t.(unaryoperatorexpr.UnaryOperatorExpr).GetValue().(expr.Expr).Accept(v, stack)
	case expr.VarExpr:
		log.Debug("Visit VarExpr")
		t.(expr.VarExpr).GetIdentifier().Accept(v, stack)
	}

	return stack
}
