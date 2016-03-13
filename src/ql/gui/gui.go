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
	GUIForm                   *GUIForm
	Symbols                   *symboltable.Symbols
	typeCheckerErrors         []error
	RegisteredOnShowCallbacks []func()
	SaveDataCallback          func() (interface{}, error)
	Window                    *ui.Window
}

// CreateGUI is a constructor method returning a new GUI
func CreateGUI(form stmt.Form, symbols *symboltable.Symbols, typeCheckerErrors []error) GUI {
	gui := GUI{GUIForm: NewGUIForm(form), Symbols: symbols, typeCheckerErrors: typeCheckerErrors}

	gui.SaveDataCallback = symbols.SaveToDisk

	form.Accept(&gui, symbols)

	gui.Show()

	return gui
}

// RegisterOnShowCallback registers callback functions that are called once the UI is presented to the user.
func (g *GUI) RegisterOnShowCallback(callback func()) {
	g.RegisteredOnShowCallbacks = append(g.RegisteredOnShowCallbacks, callback)
}

// Show shows the UI after calling
func (g *GUI) Show() {
	log.Info("Showing GUI")

	err := ui.Main(func() {
		box := ui.NewVerticalBox()
		box.SetPadded(true)
		g.Window = ui.NewWindow("QL", 800, 600, false)

		g.Window.OnClosing(func(w *ui.Window) bool {
			log.Info("Destroy of window initiated")
			g.SaveDataCallback()
			ui.Quit()
			return true
		})

		g.Window.SetChild(box)

		g.Window.Show()

		g.GUIForm.Window = g.Window

		//if len(g.typeCheckerErrors) != 0 {
		// g.showErrorDialog() FIXME reenable
		//} else {
		g.GUIForm.ShowForm()
		for _, registeredCallback := range g.RegisteredOnShowCallbacks {
			registeredCallback()
		}
		//}

	})

	if err != nil {
		panic(err)
	}
}

// VisitForm creates the top level questions
func (g *GUI) VisitForm(f interfaces.Form, s interface{}) {
	guiQuestions := handleQuestions(g, f.GetQuestions(), s.(interfaces.SymbolTable))

	g.RegisterOnShowCallback(func() {
		g.GUIForm.AddQuestionContainer(g.GUIForm.CreateQuestionTableWithRows(guiQuestions))
	})
}

func (g *GUI) VisitIf(i interfaces.If, s interface{}) {
	symbolTable := s.(interfaces.SymbolTable)

	guiQuestions := handleQuestions(g, i.GetBody().GetQuestions(), symbolTable)
	questionsEncompassingContainer := g.GUIForm.CreateQuestionTableWithRows(guiQuestions)

	g.RegisterOnShowCallback(func() {
		g.GUIForm.AddQuestionContainer(questionsEncompassingContainer)
		hideContainerWhenIfEvalToFalse(i, questionsEncompassingContainer, symbolTable)
	})

	g.Symbols.RegisterCallback(func(s interfaces.SymbolTable) {
		log.Debug("Received symbolTable update callback")

		hideContainerWhenIfEvalToFalse(i, questionsEncompassingContainer, symbolTable)
	})
}

func (g *GUI) VisitIfElse(i interfaces.IfElse, s interface{}) {
	symbolTable := s.(interfaces.SymbolTable)

	guiQuestionsIfBody := handleQuestions(g, i.GetIfBody().GetQuestions(), symbolTable)
	guiQuestionsElseBody := handleQuestions(g, i.GetElseBody().GetQuestions(), symbolTable)

	ifQuestionsEncompassingContainer := g.GUIForm.CreateQuestionTableWithRows(guiQuestionsIfBody)
	elseQuestionsEncompassingContainer := g.GUIForm.CreateQuestionTableWithRows(guiQuestionsElseBody)

	g.RegisterOnShowCallback(func() {
		g.GUIForm.AddQuestionContainer(ifQuestionsEncompassingContainer)
		g.GUIForm.AddQuestionContainer(elseQuestionsEncompassingContainer)
		hideContainerWhenIfElseEvalsToFalse(i, ifQuestionsEncompassingContainer, elseQuestionsEncompassingContainer, symbolTable)
	})

	g.Symbols.RegisterCallback(func(s interfaces.SymbolTable) {
		log.Debug("Received symbolTable update callback")

		hideContainerWhenIfElseEvalsToFalse(i, ifQuestionsEncompassingContainer, elseQuestionsEncompassingContainer, symbolTable)
	})
}

func hideContainerWhenIfEvalToFalse(conditionalStmt interfaces.Conditional, conditionalContainer *ui.Box, symbolTable interfaces.SymbolTable) {
	conditionValue := conditionalStmt.EvalCondition(symbolTable)

	if conditionValue {
		conditionalContainer.Show()
	} else {
		conditionalContainer.Hide()
	}
}

func hideContainerWhenIfElseEvalsToFalse(conditionalStmt interfaces.Conditional, conditionalContainerIfBody *ui.Box, conditionalContainerElseBody *ui.Box, symbolTable interfaces.SymbolTable) {
	conditionValue := conditionalStmt.EvalCondition(symbolTable)

	if conditionValue {
		conditionalContainerIfBody.Show()
		conditionalContainerElseBody.Hide()
	} else {
		conditionalContainerIfBody.Hide()
		conditionalContainerElseBody.Show()
	}
}

func handleQuestions(g *GUI, q []interfaces.Question, s interfaces.SymbolTable) []*GUIQuestion {
	guiQuestions := make([]*GUIQuestion, 0)

	for _, question := range q {
		switch question.(type) {
		case interfaces.ComputedQuestion:
			guiQuestion := g.handleComputedQuestion(question.(interfaces.ComputedQuestion), s.(interfaces.SymbolTable)).GUIQuestion
			guiQuestions = append(guiQuestions, guiQuestion)
		case interfaces.InputQuestion:
			guiQuestion := g.handleInputQuestion(question.(interfaces.InputQuestion), s.(interfaces.SymbolTable)).GUIQuestion
			guiQuestions = append(guiQuestions, guiQuestion)
		}
	}

	return guiQuestions
}

func (v *GUI) handleInputQuestion(question interfaces.InputQuestion, symbolTable interfaces.SymbolTable) *GUIInputQuestion {
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
		symbolTable.SetNodeForIdentifier(input, questionIdentifier)

		v.updateComputedQuestions(symbolTable)
	}

	guiQuestion = CreateGUIInputQuestion(question.GetLabelAsString(), question.GetVarDecl().GetType(), questionCallback)

	return guiQuestion
}

func (v *GUI) handleComputedQuestion(question interfaces.ComputedQuestion, symbolTable interfaces.SymbolTable) *GUIComputedQuestion {
	computation := question.GetComputation()
	guiQuestion := CreateGUIComputedQuestion(question.GetLabelAsString(), question.GetVarDecl().GetType(), computation, question.GetVarDecl().GetIdent())

	v.GUIForm.AddComputedQuestion(guiQuestion)

	return guiQuestion
}

func (g *GUI) updateComputedQuestions(symbolTable interfaces.SymbolTable) {
	for _, computedQuestion := range g.GUIForm.ComputedQuestions {
		computedQuestionEval := computedQuestion.Expr.Eval(symbolTable)
		computedQuestion.ChangeFieldValueText(fmt.Sprintf("%v", computedQuestionEval))

		// save the computed value to the symbol table
		symbolTable.SetNodeForIdentifier(computedQuestionEval, computedQuestion.VarId)
		log.WithFields(log.Fields{"eval": computedQuestionEval}).Info("Computed question value changed")
	}
}

func (g *GUI) showErrorDialog() {
	ui.MsgBoxError(g.Window, "Errors encountered", convertErrorStringListToString(g.typeCheckerErrors))
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
