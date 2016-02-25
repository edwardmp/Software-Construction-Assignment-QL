package gui

import (
	"fmt"
	//"io/ioutil"
	log "github.com/Sirupsen/logrus"
	"ql/ast/expr"
	"ql/ast/expr/binaryoperatorexpr"
	"ql/ast/expr/lit"
	"ql/ast/expr/unaryoperatorexpr"
	"ql/ast/stmt"
	"ql/ast/vari"
	"ql/ast/visit"
	"ql/env"
	"strconv"
)

/*
func presentOpenFileDialog(window *gtk.Window) {
	messagedialog := gtk.NewMessageDialog(
		window,
		gtk.DIALOG_MODAL,
		gtk.MESSAGE_INFO,
		gtk.BUTTONS_OK,
		"Choose input QL file")
	messagedialog.Response(func() {
		fmt.Println("Dialog OK!")
		filechooserdialog := gtk.NewFileChooserDialog(
			"Choose QL File",
			window,
			gtk.FILE_CHOOSER_ACTION_OPEN,
			gtk.STOCK_OK,
			gtk.RESPONSE_ACCEPT)
		filter := gtk.NewFileFilter()
		filter.AddPattern("*.ql")
		filechooserdialog.AddFilter(filter)
		filechooserdialog.Response(func() {
			fmt.Println(filechooserdialog.GetFilename())
			openQLFile(filechooserdialog.GetFilename())
			filechooserdialog.Destroy()
		})
		filechooserdialog.Run()
		messagedialog.Destroy()
	})
	messagedialog.Run()
}
*/

/*
func openQLFile(filePath string) string {
	qlFile, _ := ioutil.ReadFile(filePath)
	return string(qlFile)
}
*/

func CreateGUI(form stmt.Form, symbolTable env.SymbolTable) {
	gui := GUI{}

	gui.Visit(form, symbolTable)
}

type GUI struct {
	visit.Visitor
	Form *GUIForm
}

func (v GUI) Visit(t interface{}, s interface{}) interface{} {
	symbolTable := s.(env.SymbolTable)

	switch t.(type) {
	default:
		panic(fmt.Sprintf("Unexpected node type %T", t))
	case stmt.Form:
		log.Debug("Visit Form")

		v.Form = &GUIForm{Title: t.(stmt.Form).Identifier.Ident}

		t.(stmt.Form).Identifier.Accept(v, symbolTable)
		t.(stmt.Form).Content.Accept(v, symbolTable)

		v.Form.Show()
	case stmt.StmtList:
		log.Debug("Visit StmtList")

		for i := range t.(stmt.StmtList).Questions {
			t.(stmt.StmtList).Questions[i].(stmt.Question).Accept(v, symbolTable)
		}

		for i := range t.(stmt.StmtList).Conditionals {
			t.(stmt.StmtList).Conditionals[i].(stmt.Conditional).Accept(v, symbolTable)
		}

	case stmt.InputQuestion:
		log.Debug("Visit InputQuestion")

		question := t.(stmt.InputQuestion)

		var guiQuestion GUIQuestion
		QuestionCallback := func(input interface{}, err error) {
			if numError, ok := err.(*strconv.NumError); err != nil && ok {
				if numError.Err.Error() == "invalid syntax" {
					guiQuestion.ChangeErrorLabelText("not a valid number")
					log.Debug("Presenting invalid number error to user")
				}

				return
			}

			questionIdentifier := question.GetVarDecl().Ident
			log.WithFields(log.Fields{"input": input, "identifier": questionIdentifier}).Debug("Question input received")
			symbolTable.SetNodeForIdentifier(input.(expr.Expr), questionIdentifier)
		}

		guiQuestion = CreateGUIInputQuestion(question.GetLabelAsString(), question.GetVarDecl().GetType(), QuestionCallback)
		v.Form.AddQuestion(guiQuestion)

		question.Label.Accept(v, symbolTable)
		question.VarDecl.Accept(v, symbolTable)
	case stmt.ComputedQuestion:
		log.Debug("Visit ComputedQuestion")

		question := t.(stmt.ComputedQuestion)

		question.Label.Accept(v, symbolTable)
		question.VarDecl.Accept(v, symbolTable)
		question.Computation.Accept(v, symbolTable)

		var guiQuestion GUIQuestion
		QuestionCallback := func() {
			fmt.Printf("Symbol table %s\n", symbolTable)
			questionIdentifier := question.GetVarDecl().Ident
			computedQuestionEval := question.Computation.Eval(symbolTable)
			log.WithFields(log.Fields{"identifier": questionIdentifier, "eval": computedQuestionEval}).Info("Computed question value changed")
			guiQuestion.ChangeElementText(strconv.Itoa(computedQuestionEval.(int)))
		}

		guiQuestion = CreateGUIComputedQuestion(question.GetLabelAsString(), "Empty", QuestionCallback)

		v.Form.AddQuestion(guiQuestion)
	case stmt.If:
		log.Debug("Visit If")
		t.(stmt.If).Cond.Accept(v, symbolTable)
		t.(stmt.If).Body.Accept(v, symbolTable)
	case stmt.IfElse:
		log.Debug("Visit IfElse")
		t.(stmt.IfElse).Cond.Accept(v, symbolTable)
		t.(stmt.IfElse).IfBody.Accept(v, symbolTable)
		t.(stmt.IfElse).ElseBody.Accept(v, symbolTable)
	case vari.VarId:
		log.Debug("Visit VarId")
	case vari.VarDecl:
		log.Debug("Visit VarDecl")
		t.(vari.VarDecl).Ident.Accept(v, symbolTable)
	case lit.StrLit:
		log.Debug("Visit StrLit")
	case lit.BoolLit:
		log.Debug("Visit BoolLit")
	case lit.IntLit:
		log.Debug("Visit IntLit")
	case binaryoperatorexpr.BinaryOperatorExpr:
		log.Debug("Visit BinaryOperatorExpr")
		t.(binaryoperatorexpr.BinaryOperatorExpr).GetLhs().(expr.Expr).Accept(v, symbolTable)
		t.(binaryoperatorexpr.BinaryOperatorExpr).GetRhs().(expr.Expr).Accept(v, symbolTable)
	case unaryoperatorexpr.UnaryOperatorExpr:
		log.Debug("Visit UnaryOperatorExpr")
		t.(unaryoperatorexpr.UnaryOperatorExpr).GetValue().(expr.Expr).Accept(v, symbolTable)
	case expr.VarExpr:
		log.Debug("Visit VarExpr")
	}

	return false
}
