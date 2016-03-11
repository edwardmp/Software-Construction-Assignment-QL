package gui

import (
	"fmt"
	//"io/ioutil"
	log "github.com/Sirupsen/logrus"
	"ql/ast/stmt"
	"ql/ast/visitor"
	"ql/interfaces"
	"strconv"
)

type GUI struct {
	visitor.BaseVisitor
	Form *GUIForm
}

func CreateGUI(form stmt.Form, symbolTable interfaces.SymbolTable) {
	gui := GUI{Form: &GUIForm{Title: form.Identifier.GetIdent()}}

	gui.Form.SaveDataCallback = symbolTable.SaveToDisk

	form.Accept(gui, symbolTable)

	gui.Form.Show()
}

func (g GUI) VisitComputedQuestion(c interfaces.ComputedQuestion, s interface{}) {
	g.handleComputedQuestion(c, s.(interfaces.SymbolTable))
}

func (g GUI) VisitInputQuestion(i interfaces.InputQuestion, s interface{}) {
	g.handleInputQuestion(i, s.(interfaces.SymbolTable))
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
