package gui

import (
	log "github.com/Sirupsen/logrus"
	"github.com/andlabs/ui"
	"ql/interfaces"
)

type GUIForm struct {
	Form               interfaces.Form
	QuestionContainers []*ui.Box
	ComputedQuestions  []*GUIComputedQuestion
	FormContainer      *ui.Box
}

// NewGUIForm is a constructor method returning a new GUIForm with initialized embedded Form
func newGUIForm(form interfaces.Form) *GUIForm {
	if form == nil {
		panic("Passing nil form to GUI NewGUIForm")
	}

	return &GUIForm{Form: form, ComputedQuestions: make([]*GUIComputedQuestion, 0)}
}

// AddQuestionContainer appends the question container box to the form container
func (this *GUIForm) addQuestionContainer(questionContainer *ui.Box) {
	this.FormContainer.Append(questionContainer, false)
	log.Info("Adding question container to form")
	this.QuestionContainers = append(this.QuestionContainers, questionContainer)
}

func (this *GUIForm) addComputedQuestion(question *GUIComputedQuestion) {
	if question == nil {
		panic("Passing nil question to GUI AddComputedQuestion")
	}

	log.Info("Adding computed question to form")
	this.ComputedQuestions = append(this.ComputedQuestions, question)
}

// ShowForm displays the form box. It should only be called if no semantic errors are present.
func (this *GUIForm) show(window *ui.Window) {
	log.WithFields(log.Fields{"identifier": this.Form.GetIdentifier()}).Info("Showing form")

	this.FormContainer = ui.NewVerticalBox()
	window.SetChild(this.FormContainer)
}

func (this *GUIForm) contentCreationFinished() {
	this.addSubmitButton()
}

func extractEmbeddedGUIQuestions(inputQuestions []*GUIInputQuestion, computedQuestions []*GUIComputedQuestion) []*GUIQuestion {
	guiQuestions := make([]*GUIQuestion, 0)

	for _, question := range inputQuestions {
		guiQuestions = append(guiQuestions, question.GUIQuestion)
	}

	for _, question := range computedQuestions {
		guiQuestions = append(guiQuestions, question.GUIQuestion)
	}

	return guiQuestions
}

// CreateQuestionTableWithRows creates a table box containing the passed GUIQuestions.
func (this *GUIForm) createQuestionTable(questions []*GUIQuestion) *ui.Box {
	table := ui.NewVerticalBox()

	for _, question := range questions {
		attachQuestionToTable(table, question)
	}

	log.WithFields(log.Fields{"NumOfQuestions": len(questions)}).Info("Created question table")

	return table
}

// attachQuestionToTable is a helper method that attaches a GUIQuestion to the supplied box.
func attachQuestionToTable(table *ui.Box, question *GUIQuestion) {
	table.Append(question.Label, false)
	table.Append(question.Element, false)
	table.Append(question.ErrorLabel, false)
}

// addSubmitButton adds a submit button to the form.
func (this *GUIForm) addSubmitButton() {
	log.Info("Adding submit button to form")

	button := CreateButton("Submit", func(b *ui.Button) {
		log.Debug("Submit button clicked")
		// this.SaveDataCallback() FIXME place in Gui.go?
	})

	this.FormContainer.Append(button, false)

	/*
		display messagedialog that submit is OK
		log.Info("Submit dialog displayed")
	*/
}
