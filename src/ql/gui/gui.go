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
	switch t.(type) {
	default:
		panic(fmt.Sprintf("Unexpected node type %T", t))
	case stmt.Form:
		log.Debug("Visit Form")

		v.Form = &GUIForm{Title: t.(stmt.Form).Identifier.Ident} // TODO naming

		t.(stmt.Form).Identifier.Accept(v, s)
		t.(stmt.Form).Content.Accept(v, s)

		v.Form.Show()
	case stmt.StmtList:
		log.Debug("Visit StmtList")

		for i := range t.(stmt.StmtList).Questions {
			t.(stmt.StmtList).Questions[i].(stmt.Question).Accept(v, s)
		}

		for i := range t.(stmt.StmtList).Conditionals {
			t.(stmt.StmtList).Conditionals[i].(stmt.Conditional).Accept(v, s)
		}

	case stmt.InputQuestion:
		log.Debug("Visit InputQuestion")

		question := t.(stmt.InputQuestion)
		guiQuestion := CreateGUIQuestion(question.GetLabelAsString(), question.GetVarDecl().GetType())
		v.Form.AddQuestion(guiQuestion)

		t.(stmt.InputQuestion).Label.Accept(v, s)
		t.(stmt.InputQuestion).VarDecl.Accept(v, s)
	case stmt.ComputedQuestion:
		log.Debug("Visit ComputedQuestion")
		t.(stmt.ComputedQuestion).Label.Accept(v, s)
		t.(stmt.ComputedQuestion).VarDecl.Accept(v, s)
		t.(stmt.ComputedQuestion).Computation.Accept(v, s)
	case stmt.If:
		log.Debug("Visit If")
		t.(stmt.If).Cond.Accept(v, s)
		t.(stmt.If).Body.Accept(v, s)
	case stmt.IfElse:
		log.Debug("Visit IfElse")
		t.(stmt.IfElse).Cond.Accept(v, s)
		t.(stmt.IfElse).IfBody.Accept(v, s)
		t.(stmt.IfElse).ElseBody.Accept(v, s)
	case vari.VarId:
		log.Debug("Visit VarId")
	case vari.VarDecl:
		log.Debug("Visit VarDecl")
		t.(vari.VarDecl).Ident.Accept(v, s)
	case lit.StrLit:
		log.Debug("Visit StrLit")
	case lit.BoolLit:
		log.Debug("Visit BoolLit")
	case lit.IntLit:
		log.Debug("Visit IntLit")
	case binaryoperatorexpr.BinaryOperatorExpr:
		log.Debug("Visit BinaryOperatorExpr")
		t.(binaryoperatorexpr.BinaryOperatorExpr).GetLhs().(expr.Expr).Accept(v, s)
		t.(binaryoperatorexpr.BinaryOperatorExpr).GetRhs().(expr.Expr).Accept(v, s)
	case unaryoperatorexpr.UnaryOperatorExpr:
		log.Debug("Visit UnaryOperatorExpr")
		t.(unaryoperatorexpr.UnaryOperatorExpr).GetValue().(expr.Expr).Accept(v, s)
	case expr.VarExpr:
		log.Debug("Visit VarExpr")
	}

	return false
}
