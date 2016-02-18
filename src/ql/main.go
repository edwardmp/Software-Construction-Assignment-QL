package main

import (
	"errors"
	"fmt"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
	"ql/ast/expr/lit"
	"ql/ast/stmt"
	"ql/ast/vari"
)

func createQuestionLabel(questionText string) *gtk.Label {
	label := gtk.NewLabel(questionText)
	label.ModifyFontEasy("DejaVu Serif 12")

	return label
}

func attachQuestionToTable(table *gtk.Table, questionLabel *gtk.Label, questionElement interface{}, rowStart int) {
	table.AttachDefaults(questionLabel, 0, 1, uint(rowStart), uint(rowStart+1))
	table.AttachDefaults(questionElement.(gtk.IWidget), 1, 2, uint(rowStart), uint(rowStart+1))
}

func createQuestionElement(question stmt.Question) interface{} {
	return createRightGTKObjectForQuestionType(question)
}

func createRightGTKObjectForQuestionType(question stmt.Question) interface{} {
	var GTKEntity gtk.IWidget

	switch question.GetVarDecl().GetType() {
	case vari.BOOLEAN:
		GTKEntity = createRadioButtons()
	case vari.STRING, vari.INT:
		GTKEntity = createInputTextField(fmt.Sprintf("Type: %v", question.GetVarDecl().GetType()))
	default:
		errors.New("Unknown question type, can not create correct GTK object")
	}

	return GTKEntity
}

func createRadioButtons() *gtk.HBox {
	buttons := gtk.NewHBox(false, 1)
	firstRadioButton := gtk.NewRadioButtonWithLabel(nil, "Yes")
	buttons.Add(firstRadioButton)
	buttons.Add(gtk.NewRadioButtonWithLabel(firstRadioButton.GetGroup(), "No"))

	return buttons
}

func createInputTextField(defaultText string) *gtk.Entry {
	entry := gtk.NewEntry()
	entry.SetText(defaultText)
	entry.Connect("changed", func() {
		fmt.Printf("Input value changed: %s\n", entry.GetText())
	})

	return entry
}

func createCheckboxConditional() *gtk.CheckButton {
	checkbutton := gtk.NewCheckButtonWithLabel("CheckButton with label")
	checkbutton.Connect("toggled", func() {
		if checkbutton.GetActive() {
			checkbutton.SetLabel("CheckButton CHECKED!")
		} else {
			checkbutton.SetLabel("CheckButton UNCHECKED!")
		}
	})

	return checkbutton
}

func createQuestionLabelAndElement(question stmt.Question) (label *gtk.Label, element interface{}) {
	label = createQuestionLabel(question.GetLabelAsString())
	element = createQuestionElement(question)

	return
}

func createQuestions(questions []stmt.Question, vbox *gtk.VBox) {
	table := gtk.NewTable(uint(len(questions)), 1, false)
	for i := range questions {
		questionLabel, questionElement := createQuestionLabelAndElement(questions[i])
		attachQuestionToTable(table, questionLabel, questionElement, i)
	}

	vbox.Add(table)
}

func createSubmitButton(window *gtk.Window) *gtk.Button {
	button := createButton("Submit", func() {
		fmt.Println("Submit button clicked")

		messagedialog := gtk.NewMessageDialog(
			window,
			gtk.DIALOG_MODAL,
			gtk.MESSAGE_INFO,
			gtk.BUTTONS_OK,
			"Form saved")
		messagedialog.Response(func() {
			fmt.Println("Dialog OK!")

			messagedialog.Destroy()
		})
		messagedialog.Run()
	})

	return button
}

func createButton(buttonText string, onClick func()) *gtk.Button {
	button := gtk.NewButtonWithLabel(buttonText)
	button.Connect("clicked", onClick)
	return button
}

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

func openQLFile(filePath string) string {
	qlFile, _ := ioutil.ReadFile(filePath)
	return string(qlFile)
}

func main() {
	gtk.Init(nil)

	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetPosition(gtk.WIN_POS_CENTER)
	window.SetTitle("QL")
	window.SetIconName("gtk-dialog-info")

	window.Connect("destroy", func(ctx *glib.CallbackContext) {
		fmt.Println("Destroy of window initiated", ctx.Data().(string))
		gtk.MainQuit()
	}, "foo")

	vbox := gtk.NewVBox(false, 1)

	menubar := gtk.NewMenuBar()
	vbox.PackStart(menubar, false, false, 0)

	vpaned := gtk.NewVPaned()
	vbox.Add(vpaned)

	frame := gtk.NewFrame("Form identifier")
	framebox := gtk.NewVBox(false, 5)
	frame.Add(framebox)
	vpaned.Pack1(frame, false, false)

	firstQuestionOutput := stmt.InputQuestion{lit.StrLit{"Did you sell a house in 2010?"}, vari.VarDecl{vari.VarId{"hasSoldHouse"}, vari.BOOLEAN}}
	secondQuestionOutput := stmt.InputQuestion{lit.StrLit{"Did you not sell a house in 2010?"}, vari.VarDecl{vari.VarId{"hasSoldHouse"}, vari.BOOLEAN}}
	thirdQuestionOutput := stmt.InputQuestion{lit.StrLit{"Are you nuts?"}, vari.VarDecl{vari.VarId{"hasSoldHouse"}, vari.STRING}}
	questionsInput := []stmt.Question{firstQuestionOutput, secondQuestionOutput, thirdQuestionOutput, thirdQuestionOutput, thirdQuestionOutput, thirdQuestionOutput, thirdQuestionOutput, thirdQuestionOutput}

	createQuestions(questionsInput, framebox)

	vsep := gtk.NewVSeparator()
	vbox.PackStart(vsep, false, false, 1)

	vbox.PackStart(createSubmitButton(window), false, true, 1)

	window.Add(vbox)
	window.SetSizeRequest(400, 400)
	window.ShowAll()
	gtk.Main()
}
