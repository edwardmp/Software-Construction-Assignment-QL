package gui

import (
	"fmt"
	//"io/ioutil"
	log "github.com/Sirupsen/logrus"
	"github.com/andlabs/ui"
	"ql/ast/visitor"
	"ql/interfaces"
	"strconv"
)

type GUI struct {
	visitor.BaseVisitor
	GUIForm                   *GUIForm
	Symbols                   interfaces.VarIdValueSymbols
	RegisteredOnShowCallbacks []func()
	SaveDataCallback          func() (interface{}, error)
	Window                    *ui.Window
}

// NewGUI is a constructor method returning a new GUI
func NewGUI() *GUI {
	return &GUI{}
}

func (this *GUI) InitializeGUIForm(form interfaces.Form, symbols interfaces.VarIdValueSymbols) {
	this.setSymbols(symbols)
	this.createAndSetGUIFormFromForm(form)
}

func (this *GUI) setSymbols(symbols interfaces.VarIdValueSymbols) {
	if symbols == nil {
		panic("Passing nil symbols to GUI setSymbols")
	}

	this.Symbols = symbols
	this.SaveDataCallback = symbols.SaveToDisk
}

func (this *GUI) createAndSetGUIFormFromForm(form interfaces.Form) {
	if form == nil {
		panic("Passing nil form to GUI createAndSetGUIFormFromForm")
	}

	this.GUIForm = newGUIForm(form)
	form.Accept(this, this.Symbols)
}

// RegisterOnShowCallback registers callback functions that are called once the UI is presented to the user.
func (this *GUI) RegisterOnShowCallback(callback func()) {
	this.RegisteredOnShowCallbacks = append(this.RegisteredOnShowCallbacks, callback)
}

// Show shows the UI after calling
func (this *GUI) ShowWindow(errorsToDisplay, warningsToDisplay []error) {
	log.Info("Showing GUI")

	guiError := ui.Main(func() {
		contentBox := ui.NewVerticalBox()
		contentBox.SetPadded(true)
		this.Window = ui.NewWindow("QL", 800, 600, false)

		this.Window.OnClosing(func(w *ui.Window) bool {
			log.Info("Destroy of window initiated")
			this.SaveDataCallback()
			ui.Quit()
			return true
		})

		this.Window.SetChild(contentBox)
		this.Window.SetMargined(true)
		this.Window.Show()

		// if there are any errors/warnings, show a dialog
		if errorsPresent := this.showErrorDialogIfNecessary(errorsToDisplay); !errorsPresent {
			this.showWarningDialogIfNecessary(warningsToDisplay)
			this.ShowForm()
		}
	})

	if guiError != nil {
		log.WithFields(log.Fields{"guiError": guiError}).Panic("Encountered GUI error")
	}
}

func (this *GUI) ShowForm() {
	this.GUIForm.show(this.Window)

	for _, registeredOnShowCallback := range this.RegisteredOnShowCallbacks {
		registeredOnShowCallback()
	}

	this.GUIForm.contentCreationFinished()
}

// VisitForm creates the top level questions
func (this *GUI) VisitForm(f interfaces.Form, context interface{}) {
	guiQuestions := handleQuestions(this, f.GetQuestions())

	this.RegisterOnShowCallback(func() {
		this.GUIForm.addQuestionContainer(this.GUIForm.createQuestionTable(guiQuestions))
	})
}

func (this *GUI) VisitIf(ifStmt interfaces.If, context interface{}) {
	guiQuestions := handleQuestions(this, ifStmt.GetBody().GetQuestions())
	questionsEncompassingContainer := this.GUIForm.createQuestionTable(guiQuestions)

	this.RegisterOnShowCallback(func() {
		this.GUIForm.addQuestionContainer(questionsEncompassingContainer)
		this.showOrHideContainerDependingOnIfEval(ifStmt, questionsEncompassingContainer)
	})

	this.Symbols.RegisterCallback(func() {
		log.Debug("Received symbols update callback")
		this.showOrHideContainerDependingOnIfEval(ifStmt, questionsEncompassingContainer)
	})
}

func (this *GUI) VisitIfElse(ifElse interfaces.IfElse, context interface{}) {
	guiQuestionsIfBody := handleQuestions(this, ifElse.GetIfBody().GetQuestions())
	guiQuestionsElseBody := handleQuestions(this, ifElse.GetElseBody().GetQuestions())

	ifQuestionsEncompassingContainer := this.GUIForm.createQuestionTable(guiQuestionsIfBody)
	elseQuestionsEncompassingContainer := this.GUIForm.createQuestionTable(guiQuestionsElseBody)

	this.RegisterOnShowCallback(func() {
		this.GUIForm.addQuestionContainer(ifQuestionsEncompassingContainer)
		this.GUIForm.addQuestionContainer(elseQuestionsEncompassingContainer)

		this.showOrHideContainerDependingOnIfElseEval(ifElse, ifQuestionsEncompassingContainer, elseQuestionsEncompassingContainer)
	})

	this.Symbols.RegisterCallback(func() {
		log.Debug("Received symbols update callback")

		this.showOrHideContainerDependingOnIfElseEval(ifElse, ifQuestionsEncompassingContainer, elseQuestionsEncompassingContainer)
	})
}

func (this *GUI) showOrHideContainerDependingOnIfEval(ifStmt interfaces.If, conditionalContainer *ui.Box) {
	conditionValue := ifStmt.EvalCondition(this.Symbols)

	// if condition evals to true
	if conditionValue {
		conditionalContainer.Show()
		return
	}

	// condition evals to false
	conditionalContainer.Hide()
}

func (this *GUI) showOrHideContainerDependingOnIfElseEval(ifElseStmt interfaces.IfElse, conditionalContainerIfBody, conditionalContainerElseBody *ui.Box) {
	conditionValue := ifElseStmt.EvalCondition(this.Symbols)

	// ifElse condition evals to true
	if conditionValue {
		conditionalContainerIfBody.Show()
		conditionalContainerElseBody.Hide()
		return
	}

	// condition evals to false, show else body
	conditionalContainerIfBody.Hide()
	conditionalContainerElseBody.Show()
}

func handleQuestions(this *GUI, q []interfaces.Question) []*GUIQuestion {
	guiQuestions := make([]*GUIQuestion, 0)

	for _, question := range q {
		var guiQuestion *GUIQuestion

		switch question := question.(type) {
		case interfaces.ComputedQuestion:
			guiQuestion = this.handleComputedQuestion(question).GUIQuestion
		case interfaces.InputQuestion:
			guiQuestion = this.handleInputQuestion(question).GUIQuestion
		}

		guiQuestions = append(guiQuestions, guiQuestion)
	}

	return guiQuestions
}

func (this *GUI) handleInputQuestion(question interfaces.InputQuestion) *GUIInputQuestion {
	var guiQuestion *GUIInputQuestion
	questionCallback := func(inputExpr interfaces.Expr, err error) {
		guiQuestion.ChangeErrorLabelText("")
		if err != nil {
			if numError, isNumError := err.(*strconv.NumError); isNumError && numError.Err.Error() == "invalid syntax" {
				guiQuestion.ChangeErrorLabelText("Not a valid number!")
				log.Debug("Presenting invalid number error to user")
			}

			return
		}

		questionIdentifier := question.GetVarDecl().GetIdent()
		log.WithFields(log.Fields{"input": inputExpr, "identifier": questionIdentifier}).Debug("Question input received")
		this.Symbols.SetExprForVarId(inputExpr, questionIdentifier)

		this.updateComputedQuestions()
	}

	guiQuestion = CreateGUIInputQuestion(fmt.Sprintf("%s", question.GetLabel()), question.GetVarDecl().GetType(), questionCallback)

	return guiQuestion
}

func (this *GUI) handleComputedQuestion(question interfaces.ComputedQuestion) *GUIComputedQuestion {
	computation := question.GetComputation()
	guiQuestion := CreateGUIComputedQuestion(fmt.Sprintf("%s", question.GetLabel()), question.GetVarDecl().GetType(), computation, question.GetVarDecl().GetIdent())

	this.GUIForm.addComputedQuestion(guiQuestion)

	return guiQuestion
}

func (this *GUI) updateComputedQuestions() {
	for _, computedQuestion := range this.GUIForm.ComputedQuestions {
		computedQuestionEval := computedQuestion.Expr.Eval(this.Symbols)
		computedQuestion.ChangeFieldValueText(fmt.Sprintf("%v", computedQuestionEval))

		// save the computed value to the symbol table
		this.Symbols.SetExprForVarId(computedQuestion.Expr, computedQuestion.VarId)
		log.WithFields(log.Fields{"evaluatesTo": computedQuestionEval}).Info("Computed question value changed")
	}
}

func (this *GUI) showErrorDialogIfNecessary(errorsToDisplay []error) bool {
	if len(errorsToDisplay) == 0 {
		return false
	}

	ShowMessageBoxForErrors("Errors encountered", errorsToDisplay, this.Window)

	return true
}

func (this *GUI) showWarningDialogIfNecessary(warningsToDisplay []error) bool {
	if len(warningsToDisplay) == 0 {
		return false
	}

	ShowMessageBoxForErrors("Warnings encountered", warningsToDisplay, this.Window)

	return true
}
