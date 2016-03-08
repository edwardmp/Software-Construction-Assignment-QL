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
	"ql/ast/vari/vartype"
	"ql/ast/visit"
	"ql/symboltable"
	"strconv"
)

type GUI struct {
	visit.Visitor
	Form *GUIForm
}

func CreateGUI(form stmt.Form, symbolTable symboltable.SymbolTable) {
	gui := GUI{Form: &GUIForm{Title: form.Identifier.Ident}}

	gui.Form.SaveDataCallback = symbolTable.SaveToDisk

	gui.Visit(form, symbolTable)
}

func (v GUI) Visit(t interface{}, s interface{}) interface{} {
	symbolTable := s.(symboltable.SymbolTable)

	switch t.(type) {
	default:
		panic(fmt.Sprintf("Unexpected node type %T", t))
	case stmt.Form:
		log.Debug("Visit Form")

		t.(stmt.Form).Identifier.Accept(v, symbolTable)
		t.(stmt.Form).Content.Accept(v, symbolTable)

		v.Form.Show()
	case stmt.StmtList:
		log.Debug("Visit StmtList")

		for _, question := range t.(stmt.StmtList).Questions {
			question.Accept(v, symbolTable)
		}

		for _, conditional := range t.(stmt.StmtList).Conditionals {
			conditional.Accept(v, symbolTable)
		}

	case stmt.InputQuestion:
		log.Debug("Visit InputQuestion")
		question := t.(stmt.InputQuestion)
		v.handleInputQuestion(question, symbolTable)

		question.Label.Accept(v, symbolTable)
		question.VarDecl.Accept(v, symbolTable)
	case stmt.ComputedQuestion:
		log.Debug("Visit ComputedQuestion")

		question := t.(stmt.ComputedQuestion)

		v.handleComputedQuestion(question, symbolTable)

		question.Label.Accept(v, symbolTable)
		question.VarDecl.Accept(v, symbolTable)
		question.Computation.Accept(v, symbolTable)

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
	case vartype.VarType:
		log.Debug("Visit VarType")
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

	return nil
}

func (v GUI) handleInputQuestion(question stmt.InputQuestion, symbolTable symboltable.SymbolTable) {
	var guiQuestion GUIInputQuestion
	questionCallback := func(input interface{}, err error) {
		if err != nil {
			if numError, ok := err.(*strconv.NumError); err != nil && ok {
				if numError.Err.Error() == "invalid syntax" {
					guiQuestion.ChangeErrorLabelText("not a valid number")
					log.Debug("Presenting invalid number error to user")
				}
			}

			return
		}

		questionIdentifier := question.GetVarDecl().Ident
		log.WithFields(log.Fields{"input": input, "identifier": questionIdentifier}).Debug("Question input received")
		symbolTable.SetNodeForIdentifier(input, questionIdentifier)

		v.updateComputedQuestions(symbolTable)
	}

	guiQuestion = CreateGUIInputQuestion(question.GetLabelAsString(), question.GetVarDecl().GetType(), questionCallback)
	v.Form.AddInputQuestion(guiQuestion)
}

func (v GUI) handleComputedQuestion(question stmt.ComputedQuestion, symbolTable symboltable.SymbolTable) {
	computation := question.Computation.(expr.Expr)
	guiQuestion := CreateGUIComputedQuestion(question.GetLabelAsString(), question.VarDecl.GetType(), computation, question.VarDecl.GetIdentifier())

	v.Form.AddComputedQuestion(guiQuestion)
}

func (g GUI) updateComputedQuestions(symbolTable symboltable.SymbolTable) {
	for _, computedQuestion := range g.Form.ComputedQuestions {
		computedQuestionEval := computedQuestion.Expr.Eval(symbolTable)
		computedQuestion.GUIQuestion.ChangeElementText(fmt.Sprintf("%v", computedQuestionEval))

		// save the computed value to the symbol table
		symbolTable.SetNodeForIdentifier(computedQuestionEval, computedQuestion.VarId)

		log.WithFields(log.Fields{"eval": computedQuestionEval}).Info("Computed question value changed")
	}
}

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
