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
	Window             *ui.Window
	Container          *ui.Box
}

// NewGUIForm is a constructor method returning a new GUIForm with initialized embedded Form
func NewGUIForm(form interfaces.Form) *GUIForm {
	return &GUIForm{Form: form}
}

// AddQuestionContainer appends the question container box to the form container
func (this *GUIForm) AddQuestionContainer(questionContainer *ui.Box) {
	this.Container.Append(questionContainer, false)
	log.Info("Adding question container to form")
	this.QuestionContainers = append(this.QuestionContainers, questionContainer)
}

func (this *GUIForm) AddComputedQuestion(question *GUIComputedQuestion) {
	log.Info("Adding computed question to form")
	this.ComputedQuestions = append(this.ComputedQuestions, question)
}

// ShowForm displays the form box. It should only be called if no semantic errors are present.
func (this *GUIForm) ShowForm() {
	log.WithFields(log.Fields{"identifier": this.Form.GetIdentifier()}).Info("Showing form")

	box := ui.NewVerticalBox()
	this.Container = box

	this.addSubmitButton()
	this.Container.Append(ui.NewHorizontalSeparator(), false)
	this.Window.SetChild(this.Container)
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
func (this *GUIForm) CreateQuestionTableWithRows(questions []*GUIQuestion) *ui.Box {
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

	this.Container.Append(button, false)

	/*
		display messagedialog that submit is OK
		log.Info("Submit dialog displayed")
	*/
}
