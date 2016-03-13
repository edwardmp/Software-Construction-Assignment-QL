package gui

import (
	"fmt"
	//"io/ioutil"
	log "github.com/Sirupsen/logrus"
	"github.com/andlabs/ui"
	"ql/ast/stmt"

	"ql/ast/visitor"
	"ql/interfaces"
	"ql/symboltable"
	"strconv"
	"strings"
)

type GUI struct {
	visitor.BaseVisitor
	Form              *GUIForm
	Symbols           *symboltable.Symbols
	typeCheckerErrors []error
}

func CreateGUI(form stmt.Form, symbols *symboltable.Symbols, typeCheckerErrors []error) GUI {
	gui := GUI{Form: &GUIForm{Title: form.Identifier.GetIdent()}, Symbols: symbols, typeCheckerErrors: typeCheckerErrors}

	gui.Form.SaveDataCallback = symbols.SaveToDisk

	form.Accept(&gui, symbols)

	gui.Show()

	return gui
}

func (g *GUI) Show() {
	log.Info("Showing GUI")

	err := ui.Main(func() {
		box := ui.NewVerticalBox()
		window := ui.NewWindow("QL", 200, 100, false)

		window.SetChild(box)

		window.Show()

		g.Form.Window = window

		//if len(g.typeCheckerErrors) != 0 {
		// g.ShowErrorDialog() FIXME reenable
		//} else {
		g.Form.ShowForm()
		//}

	})

	if err != nil {
		panic(err)
	}
}

// VisitForm creates the top level questions
func (g *GUI) VisitForm(f interfaces.Form, s interface{}) {
	handleQuestions(g, f.GetQuestions(), s.(interfaces.SymbolTable))
}

func (g *GUI) VisitIf(i interfaces.If, s interface{}) {
	body := i.GetBody()
	handleQuestions(g, body.GetQuestions(), s.(interfaces.SymbolTable))
	/*
		g.Symbols.RegisterCallback(func(varId interfaces.VarId) {
			fmt.Println("Callback working")
			question := body.GetQuestions()[0]
			q := CreateGUIInputQuestion("s5er", question.GetVarDecl().GetType(), nil)
			//g.Form.Container.Remove(CreateQuestions([]GUIQuestion{q.GUIQuestion}, g.Form.Container))

			//if varId == i.GetIdent()
		})
	*/
}

func handleQuestions(g *GUI, q []interfaces.Question, s interfaces.SymbolTable) {
	for _, question := range q {
		switch question.(type) {
		case interfaces.ComputedQuestion:
			g.handleComputedQuestion(question.(interfaces.ComputedQuestion), s.(interfaces.SymbolTable))
		case interfaces.InputQuestion:
			g.handleInputQuestion(question.(interfaces.InputQuestion), s.(interfaces.SymbolTable))
		}
	}
}

func (v *GUI) handleInputQuestion(question interfaces.InputQuestion, symbolTable interfaces.SymbolTable) {
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

func (v *GUI) handleComputedQuestion(question interfaces.ComputedQuestion, symbolTable interfaces.SymbolTable) {

	computation := question.GetComputation()
	guiQuestion := CreateGUIComputedQuestion(question.GetLabelAsString(), question.GetVarDecl().GetType(), computation, question.GetVarDecl().GetIdent())

	v.Form.AddComputedQuestion(guiQuestion)
}

func (g *GUI) updateComputedQuestions(symbolTable interfaces.SymbolTable) {
	for _, computedQuestion := range g.Form.ComputedQuestions {
		computedQuestionEval := computedQuestion.Expr.Eval(symbolTable)
		computedQuestion.ChangeElementText(fmt.Sprintf("%v", computedQuestionEval))

		// save the computed value to the symbol table
		symbolTable.SetNodeForIdentifier(computedQuestionEval, computedQuestion.VarId)
		log.WithFields(log.Fields{"eval": computedQuestionEval}).Info("Computed question value changed")
	}
}

func (g *GUI) ShowErrorDialog() {
	errorStrings := []string{}
	for _, singleError := range g.typeCheckerErrors {
		errorStrings = append(errorStrings, fmt.Sprintf("%s", singleError))
	}

	errorsAsString := strings.Join(errorStrings, "\n")
	ui.MsgBoxError(g.Form.Window, "Errors encountered", errorsAsString)
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
