package gui

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
)

type GUIForm struct {
	Title             string
	InputQuestions    []GUIInputQuestion
	ComputedQuestions []GUIComputedQuestion
	SaveDataCallback  func() (interface{}, error)
}

func (g *GUIForm) AddInputQuestion(question GUIInputQuestion) {
	g.InputQuestions = append(g.InputQuestions, question)
}

func (g *GUIForm) AddComputedQuestion(question GUIComputedQuestion) {
	g.ComputedQuestions = append(g.ComputedQuestions, question)
}

func (g *GUIForm) Show() {
	log.Info("Showing form")

	gtk.Init(nil)

	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetPosition(gtk.WIN_POS_CENTER)
	window.SetTitle("QL")
	window.SetIconName("gtk-dialog-info")

	window.Connect("destroy", func(ctx *glib.CallbackContext) {
		fmt.Println("Destroy of window initiated", ctx.Data().(string))
		g.SaveDataCallback()
		gtk.MainQuit()
	}, "foo")

	vbox := gtk.NewVBox(false, 1)

	vpaned := gtk.NewVPaned()
	vbox.Add(vpaned)

	frame := gtk.NewFrame(fmt.Sprintf("Form %s", g.Title))
	framebox := gtk.NewVBox(false, 5)
	frame.Add(framebox)
	vpaned.Pack1(frame, false, false)

	createQuestions(extractEmbeddedGUIQuestions(g.InputQuestions, g.ComputedQuestions), framebox)

	vsep := gtk.NewVSeparator()
	vbox.PackStart(vsep, false, false, 1)

	vbox.PackStart(createSubmitButton(g, window), false, true, 1)

	window.Add(vbox)
	//window.SetSizeRequest(400, 400)
	window.ShowAll()
	gtk.Main()
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

func createQuestions(questions []GUIQuestion, vbox *gtk.VBox) {
	table := gtk.NewTable(uint(len(questions)), 1, false)
	for index, question := range questions {
		attachToTable(table, question, index)
	}

	vbox.Add(table)

	log.WithFields(log.Fields{"NumOfQuestions": len(questions)}).Info("Created question table")
}

func attachToTable(table *gtk.Table, question GUIQuestion, rowStart int) {
	table.AttachDefaults(question.Label, 0, 1, uint(rowStart), uint(rowStart+1))
	table.AttachDefaults(question.Element.(gtk.IWidget), 1, 2, uint(rowStart), uint(rowStart+1))
	table.AttachDefaults(question.ErrorLabel, 2, 3, uint(rowStart), uint(rowStart+1))
}

func createSubmitButton(form *GUIForm, window *gtk.Window) *gtk.Button {
	button := CreateButton("Submit", func() {
		log.Debug("Submit button clicked")
		form.SaveDataCallback()
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
	})

	return button
}
