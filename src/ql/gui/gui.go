package gui

import (
	"fmt"
	//"io/ioutil"
	log "github.com/Sirupsen/logrus"
	"github.com/andlabs/ui"
	"ql/ast/stmt"
	"ql/ast/visitor"
	"ql/interfaces"
	"strconv"
	"strings"
)

type GUI struct {
	visitor.BaseVisitor
	GUIForm                   *GUIForm
	Symbols                   interfaces.Symbols
	typeCheckerErrors         []error
	RegisteredOnShowCallbacks []func()
	SaveDataCallback          func() (interface{}, error)
	Window                    *ui.Window
}

// CreateGUI is a constructor method returning a new GUI
func CreateGUI(form stmt.Form, symbols interfaces.Symbols, typeCheckerErrors []error) GUI {
	gui := GUI{GUIForm: NewGUIForm(form), Symbols: symbols, typeCheckerErrors: typeCheckerErrors}

	gui.SaveDataCallback = symbols.SaveToDisk

	form.Accept(&gui, symbols)

	gui.Show()

	return gui
}

// RegisterOnShowCallback registers callback functions that are called once the UI is presented to the user.
func (this *GUI) RegisterOnShowCallback(callback func()) {
	this.RegisteredOnShowCallbacks = append(this.RegisteredOnShowCallbacks, callback)
}

// Show shows the UI after calling
func (this *GUI) Show() {
	log.Info("Showing GUI")

	err := ui.Main(func() {
		box := ui.NewVerticalBox()
		box.SetPadded(true)
		this.Window = ui.NewWindow("QL", 800, 600, false)

		this.Window.OnClosing(func(w *ui.Window) bool {
			log.Info("Destroy of window initiated")
			this.SaveDataCallback()
			ui.Quit()
			return true
		})

		this.Window.SetChild(box)

		this.Window.Show()

		this.GUIForm.Window = this.Window

		// FIXME reenable
		//if len(this.typeCheckerErrors) != 0 {
		// this.showErrorDialog()
		//} else {
		this.GUIForm.ShowForm()
		for _, registeredCallback := range this.RegisteredOnShowCallbacks {
			registeredCallback()
		}
		//}

	})

	if err != nil {
		panic(err)
	}
}

// VisitForm creates the top level questions
func (this *GUI) VisitForm(f interfaces.Form, symbols interfaces.Symbols) {
	guiQuestions := handleQuestions(this, f.GetQuestions(), symbols)

	this.RegisterOnShowCallback(func() {
		this.GUIForm.AddQuestionContainer(this.GUIForm.CreateQuestionTableWithRows(guiQuestions))
	})
}

func (this *GUI) VisitIf(i interfaces.If, symbols interfaces.Symbols) {
	guiQuestions := handleQuestions(this, i.GetBody().GetQuestions(), symbols)
	questionsEncompassingContainer := this.GUIForm.CreateQuestionTableWithRows(guiQuestions)

	this.RegisterOnShowCallback(func() {
		this.GUIForm.AddQuestionContainer(questionsEncompassingContainer)
		hideContainerWhenIfEvalToFalse(i, questionsEncompassingContainer, symbols)
	})

	this.Symbols.RegisterCallback(func(s interfaces.Symbols) {
		log.Debug("Received symbols update callback")

		hideContainerWhenIfEvalToFalse(i, questionsEncompassingContainer, symbols)
	})
}

func (this *GUI) VisitIfElse(i interfaces.IfElse, symbols interfaces.Symbols) {
	guiQuestionsIfBody := handleQuestions(this, i.GetIfBody().GetQuestions(), symbols)
	guiQuestionsElseBody := handleQuestions(this, i.GetElseBody().GetQuestions(), symbols)

	ifQuestionsEncompassingContainer := this.GUIForm.CreateQuestionTableWithRows(guiQuestionsIfBody)
	elseQuestionsEncompassingContainer := this.GUIForm.CreateQuestionTableWithRows(guiQuestionsElseBody)

	this.RegisterOnShowCallback(func() {
		this.GUIForm.AddQuestionContainer(ifQuestionsEncompassingContainer)
		this.GUIForm.AddQuestionContainer(elseQuestionsEncompassingContainer)
		hideContainerWhenIfElseEvalsToFalse(i, ifQuestionsEncompassingContainer, elseQuestionsEncompassingContainer, symbols)
	})

	this.Symbols.RegisterCallback(func(s interfaces.Symbols) {
		log.Debug("Received symbols update callback")

		hideContainerWhenIfElseEvalsToFalse(i, ifQuestionsEncompassingContainer, elseQuestionsEncompassingContainer, symbols)
	})
}

func hideContainerWhenIfEvalToFalse(conditionalStmt interfaces.Conditional, conditionalContainer *ui.Box, symbols interfaces.Symbols) {
	conditionValue := conditionalStmt.EvalCondition(symbols)

	if conditionValue {
		conditionalContainer.Show()
	} else {
		conditionalContainer.Hide()
	}
}

func hideContainerWhenIfElseEvalsToFalse(conditionalStmt interfaces.Conditional, conditionalContainerIfBody *ui.Box, conditionalContainerElseBody *ui.Box, symbols interfaces.Symbols) {
	conditionValue := conditionalStmt.EvalCondition(symbols)

	if conditionValue {
		conditionalContainerIfBody.Show()
		conditionalContainerElseBody.Hide()
	} else {
		conditionalContainerIfBody.Hide()
		conditionalContainerElseBody.Show()
	}
}

func handleQuestions(this *GUI, q []interfaces.Question, symbols interfaces.Symbols) []*GUIQuestion {
	guiQuestions := make([]*GUIQuestion, 0)

	for _, question := range q {
		switch question.(type) {
		case interfaces.ComputedQuestion:
			guiQuestion := this.handleComputedQuestion(question.(interfaces.ComputedQuestion), symbols).GUIQuestion
			guiQuestions = append(guiQuestions, guiQuestion)
		case interfaces.InputQuestion:
			guiQuestion := this.handleInputQuestion(question.(interfaces.InputQuestion), symbols).GUIQuestion
			guiQuestions = append(guiQuestions, guiQuestion)
		}
	}

	return guiQuestions
}

func (v *GUI) handleInputQuestion(question interfaces.InputQuestion, symbols interfaces.Symbols) *GUIInputQuestion {
	var guiQuestion *GUIInputQuestion
	questionCallback := func(input interface{}, err error) {
		guiQuestion.ChangeErrorLabelText("")
		if err != nil {
			if numError, ok := err.(*strconv.NumError); err != nil && ok {
				if numError.Err.Error() == "invalid syntax" {
					guiQuestion.ChangeErrorLabelText("Not a valid number!")
					log.Debug("Presenting invalid number error to user")
				}
			}

			return
		}

		questionIdentifier := question.GetVarDecl().GetIdent()
		log.WithFields(log.Fields{"input": input, "identifier": questionIdentifier}).Debug("Question input received")
		symbols.SetNodeForIdentifier(input, questionIdentifier)

		v.updateComputedQuestions(symbols)
	}

	guiQuestion = CreateGUIInputQuestion(question.GetLabelAsString(), question.GetVarDecl().GetType(), questionCallback)

	return guiQuestion
}

func (v *GUI) handleComputedQuestion(question interfaces.ComputedQuestion, symbols interfaces.Symbols) *GUIComputedQuestion {
	computation := question.GetComputation()
	guiQuestion := CreateGUIComputedQuestion(question.GetLabelAsString(), question.GetVarDecl().GetType(), computation, question.GetVarDecl().GetIdent())

	v.GUIForm.AddComputedQuestion(guiQuestion)

	return guiQuestion
}

func (this *GUI) updateComputedQuestions(symbols interfaces.Symbols) {
	for _, computedQuestion := range this.GUIForm.ComputedQuestions {
		computedQuestionEval := computedQuestion.Expr.Eval(symbols)
		computedQuestion.ChangeFieldValueText(fmt.Sprintf("%v", computedQuestionEval))

		// save the computed value to the symbol table
		symbols.SetNodeForIdentifier(computedQuestionEval, computedQuestion.VarId)
		log.WithFields(log.Fields{"eval": computedQuestionEval}).Info("Computed question value changed")
	}
}

func (this *GUI) showErrorDialog() {
	ui.MsgBoxError(this.Window, "Errors encountered", convertErrorStringListToString(this.typeCheckerErrors))
}

// convertErrorStringListToString converts a list of errors to a concatenated error string and returns it.
func convertErrorStringListToString(errorStringList []error) string {
	errorStrings := []string{}

	for _, singleError := range errorStringList {
		errorStrings = append(errorStrings, fmt.Sprintf("%s", singleError))
	}

	return strings.Join(errorStrings, "\n")
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
