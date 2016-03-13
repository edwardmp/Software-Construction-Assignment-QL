package gui

import (
	log "github.com/Sirupsen/logrus"
	"github.com/andlabs/ui"
)

type GUIForm struct {
	Title             string
	InputQuestions    []GUIInputQuestion
	ComputedQuestions []GUIComputedQuestion
	SaveDataCallback  func() (interface{}, error)
	Window            *ui.Window
	Container         *ui.Box
}

func (g *GUIForm) AddInputQuestion(question GUIInputQuestion) {
	log.Info("Adding input question to form")
	g.InputQuestions = append(g.InputQuestions, question)
}

func (g *GUIForm) AddComputedQuestion(question GUIComputedQuestion) {
	log.Info("Adding computed question to form")
	g.ComputedQuestions = append(g.ComputedQuestions, question)
}

func (g *GUIForm) ShowForm() {
	log.Info("Showing form")

	g.Window.OnClosing(func(w *ui.Window) bool {
		log.Info("Destroy of window initiated")
		g.SaveDataCallback()
		ui.Quit()
		return true
	})

	box := ui.NewVerticalBox()
	g.Container = box
	g.Window.SetChild(box)

	box.Append(CreateQuestions(extractEmbeddedGUIQuestions(g.InputQuestions, g.ComputedQuestions), box), false)
	createSubmitButton(g, box)
}

func extractEmbeddedGUIQuestions(inputQuestions []GUIInputQuestion, computedQuestions []GUIComputedQuestion) []GUIQuestion {
	guiQuestions := make([]GUIQuestion, 0)

	for _, question := range inputQuestions {
		guiQuestions = append(guiQuestions, question.GUIQuestion)
	}

	for _, question := range computedQuestions {
		guiQuestions = append(guiQuestions, question.GUIQuestion)
	}

	return guiQuestions
}

func CreateQuestions(questions []GUIQuestion, box *ui.Box) *ui.Box {
	table := ui.NewVerticalBox()

	for _, question := range questions {
		attachQuestionToTable(table, question)
	}

	log.WithFields(log.Fields{"NumOfQuestions": len(questions)}).Info("Created question table")

	return table
}

func attachQuestionToTable(table *ui.Box, question GUIQuestion) {
	table.Append(question.Label, false)
	table.Append(question.Element, false)
	table.Append(question.ErrorLabel, false)
}

func createSubmitButton(form *GUIForm, box *ui.Box) {
	button := CreateButton("Submit", func(b *ui.Button) {
		log.Debug("Submit button clicked")
		form.SaveDataCallback()
	})

	box.Append(button, false)

	/*
		messagedialog := gtk.NewMessageDialog(
			window,
			gtk.DIALOG_MODAL,
			gtk.MESSAGE_INFO,
			gtk.BUTTONS_OK,
			"Form saved")
		messagedialog.Response(func() {
			log.Info("Submit dialog displayed")

			messagedialog.Destroy()
		})
		messagedialog.Run()
	*/
}
