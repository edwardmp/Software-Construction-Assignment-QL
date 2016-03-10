package gui

import (
	"fmt"
	//"io/ioutil"
	log "github.com/Sirupsen/logrus"
	"ql/ast/stmt"
	"ql/interfaces"
	"ql/symboltable"
	"strconv"
)

type GUI struct {
	interfaces.Visitor
	Form *GUIForm
}

func CreateGUI(form stmt.Form, symbolTable symboltable.SymbolTable) {
	gui := GUI{Form: &GUIForm{Title: form.Identifier.GetIdent()}}

	gui.Form.SaveDataCallback = symbolTable.SaveToDisk

	form.Accept(gui, symbolTable)

	gui.Form.Show()
}

func (g GUI) VisitAdd(a interfaces.Add, s interface{}) {

}

func (g GUI) VisitAnd(a interfaces.And, s interface{}) {

}

func (g GUI) VisitDiv(d interfaces.Div, s interface{}) {

}

func (g GUI) VisitEq(e interfaces.Eq, s interface{}) {

}

func (g GUI) VisitGEq(ge interfaces.GEq, s interface{}) {

}

func (g GUI) VisitGT(gt interfaces.GT, s interface{}) {

}

func (g GUI) VisitLEq(l interfaces.LEq, s interface{}) {

}

func (g GUI) VisitLT(l interfaces.LT, s interface{}) {

}

func (g GUI) VisitMul(m interfaces.Mul, s interface{}) {

}

func (g GUI) VisitNEq(n interfaces.NEq, s interface{}) {

}

func (g GUI) VisitOr(o interfaces.Or, s interface{}) {

}

func (g GUI) VisitSub(su interfaces.Sub, s interface{}) {

}

func (g GUI) VisitBoolLit(b interfaces.BoolLit, s interface{}) {

}

func (g GUI) VisitIntLit(i interfaces.IntLit, s interface{}) {

}

func (g GUI) VisitStrLit(st interfaces.StrLit, s interface{}) {

}

func (g GUI) VisitNeg(n interfaces.Neg, s interface{}) {

}

func (g GUI) VisitNot(n interfaces.Not, s interface{}) {

}

func (g GUI) VisitPos(p interfaces.Pos, s interface{}) {

}

func (g GUI) VisitVarExpr(va interfaces.VarExpr, s interface{}) {

}

func (g GUI) VisitVarDecl(va interfaces.VarDecl, s interface{}) {

}

func (g GUI) VisitVarId(va interfaces.VarId, s interface{}) {

}

func (g GUI) VisitIntType(i interfaces.IntType, s interface{}) {

}

func (g GUI) VisitBoolType(b interfaces.BoolType, s interface{}) {

}

func (g GUI) VisitStringType(st interfaces.StringType, s interface{}) {

}

func (g GUI) VisitForm(f interfaces.Form, s interface{}) {
}

func (g GUI) VisitComputedQuestion(c interfaces.ComputedQuestion, s interface{}) {
	g.handleComputedQuestion(c, s.(interfaces.SymbolTable))
}

func (g GUI) VisitInputQuestion(i interfaces.InputQuestion, s interface{}) {
	g.handleInputQuestion(i, s.(interfaces.SymbolTable))
}

func (g GUI) VisitIf(i interfaces.If, s interface{}) {

}

func (g GUI) VisitIfElse(i interfaces.IfElse, s interface{}) {

}

func (g GUI) VisitStmtList(st interfaces.StmtList, sy interface{}) {
}

func (v GUI) handleInputQuestion(question interfaces.InputQuestion, symbolTable interfaces.SymbolTable) {
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

		questionIdentifier := question.GetVarDecl().GetIdent()
		log.WithFields(log.Fields{"input": input, "identifier": questionIdentifier}).Debug("Question input received")
		symbolTable.SetNodeForIdentifier(input, questionIdentifier)

		v.updateComputedQuestions(symbolTable)
	}

	guiQuestion = CreateGUIInputQuestion(question.GetLabelAsString(), question.GetVarDecl().GetType(), questionCallback)
	v.Form.AddInputQuestion(guiQuestion)
}

func (v GUI) handleComputedQuestion(question interfaces.ComputedQuestion, symbolTable interfaces.SymbolTable) {
	computation := question.GetComputation()
	guiQuestion := CreateGUIComputedQuestion(question.GetLabelAsString(), question.GetVarDecl().GetType(), computation, question.GetVarDecl().GetIdent())

	v.Form.AddComputedQuestion(guiQuestion)
}

func (g GUI) updateComputedQuestions(symbolTable interfaces.SymbolTable) {
	for _, computedQuestion := range g.Form.ComputedQuestions {
		computedQuestionEval := computedQuestion.Expr.Eval(symbolTable)
		computedQuestion.ChangeElementText(fmt.Sprintf("%v", computedQuestionEval))

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
