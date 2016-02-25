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
	"ql/ast/visit"
	"ql/env"
	"ql/gui"
	"ql/lexer"
	"ql/parser"
)

func main() {
	initLog()

	log.Info("Initiating parsing of file")

	qlFile, _ := ioutil.ReadFile("example.ql")
	qlFileAsString := string(qlFile)

	lex := lexer.NewLexer([]byte(qlFileAsString))

	p := parser.NewParser()
	result, err := p.Parse(lex)

	if err != nil {
		panic(err)
	}

	if parsedForm, ok := result.(stmt.Form); !ok {
		panic("Parse result is not form")
	} else {
		log.WithFields(log.Fields{"Result": result}).Info("Form parsed")

		visitor := VisitorAdapter{}
		symbolTableStack := env.NewSymbolTable()

		symbolTable := parsedForm.Accept(visitor, symbolTableStack).(env.SymbolTable)

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
	stack := s.(env.SymbolTable)

	switch t.(type) {
	default:
		panic(fmt.Sprintf("Unexpected node type %T", t))
	case stmt.Form:
		log.Debug("Visit Form")
		t.(stmt.Form).Identifier.Accept(v, stack)
		return t.(stmt.Form).Content.Accept(v, stack)
	case vari.VarId:
		log.Debug("Visit VarId")
	case vari.VarDecl:
		log.Debug("Visit VarDecl")
		stack.SetNodeForIdentifier(lit.IntLit{19}, t.(vari.VarDecl).Ident)
		t.(vari.VarDecl).Ident.Accept(v, stack)
	case stmt.StmtList:
		log.Debug("Visit StmtList")

		for i := range t.(stmt.StmtList).Questions {
			t.(stmt.StmtList).Questions[i].(stmt.Question).Accept(v, stack)
		}

		for i := range t.(stmt.StmtList).Conditionals {
			t.(stmt.StmtList).Conditionals[i].(stmt.Conditional).Accept(v, stack)
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
