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

func question(table *gtk.Table, question stmt.Question, rowStart, rowEnd int) {
	label := gtk.NewLabel(question.GetLabelAsString())
	label.ModifyFontEasy("DejaVu Serif 12")

	table.AttachDefaults(label, 0, 1, uint(rowStart), uint(rowEnd))
	table.AttachDefaults(createRightGTKObjectForQuestionType(question).(gtk.IWidget), 1, 2, uint(rowStart), uint(rowEnd))
}

func createRightGTKObjectForQuestionType(question stmt.Question) interface{} {
	var GTKEntity gtk.IWidget
	switch question.GetVarDecl().GetType() {
	case vari.BOOLEAN:
		buttons := gtk.NewHBox(false, 1)
		radiofirst := gtk.NewRadioButtonWithLabel(nil, "Yes")
		buttons.Add(radiofirst)
		buttons.Add(gtk.NewRadioButtonWithLabel(radiofirst.GetGroup(), "No"))
		GTKEntity = buttons
	case vari.STRING, vari.INT:
		entry := gtk.NewEntry()
		entry.SetText(fmt.Sprintf("Type: %v", question.GetVarDecl().GetType()))
		entry.Connect("changed", func() {
			fmt.Printf("Input value changed: %s\n", entry.GetText())
		})
		GTKEntity = entry
	default:
		errors.New("Unknown question type")
	}

	return GTKEntity
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

func questions(questions []stmt.Question, vbox *gtk.VBox) {
	table := gtk.NewTable(uint(len(questions)), 1, false)
	for i := range questions {
		question(table, questions[i], i, i+1)
	}
	vbox.Add(table)
}

func createSubmitButton() *gtk.Button {
	button := gtk.NewButtonWithLabel("Submit")
	button.Connect("clicked", func() {
		fmt.Println("Submit button clicked")

		messagedialog := gtk.NewMessageDialog(
			button.GetTopLevelAsWindow(),
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

func presentOpenFileDialog() {
	messagedialog := gtk.NewMessageDialog(
		button.GetTopLevelAsWindow(),
		gtk.DIALOG_MODAL,
		gtk.MESSAGE_INFO,
		gtk.BUTTONS_OK,
		"Choose input QL file")
	messagedialog.Response(func() {
		fmt.Println("Dialog OK!")
		filechooserdialog := gtk.NewFileChooserDialog(
			"Choose QL File",
			button.GetTopLevelAsWindow(),
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

func openQLFile(filePath string) {

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

	frame1 := gtk.NewFrame("Form identifier")
	framebox1 := gtk.NewVBox(false, 1)
	frame1.Add(framebox1)
	vpaned.Pack1(frame1, false, false)

	firstQuestionOutput := stmt.InputQuestion{lit.StrLit{"Did you sell a house in 2010?"}, vari.VarDecl{vari.VarId{"hasSoldHouse"}, vari.BOOLEAN}}
	secondQuestionOutput := stmt.InputQuestion{lit.StrLit{"Did you not sell a house in 2010?"}, vari.VarDecl{vari.VarId{"hasSoldHouse"}, vari.BOOLEAN}}
	thirdQuestionOutput := stmt.InputQuestion{lit.StrLit{"Are you nuts?"}, vari.VarDecl{vari.VarId{"hasSoldHouse"}, vari.STRING}}
	questionsInput := []stmt.Question{firstQuestionOutput, secondQuestionOutput, thirdQuestionOutput, thirdQuestionOutput, thirdQuestionOutput, thirdQuestionOutput, thirdQuestionOutput, thirdQuestionOutput}

	vbox.PackStart(createOpenFileButton(), false, true, 0)
	questions(questionsInput, framebox1)

	vsep := gtk.NewVSeparator()
	vbox.PackStart(vsep, false, false, 1)

	vbox.PackStart(createSubmitButton(), false, true, 0)

	window.Add(vbox)
	window.SetSizeRequest(600, 600)
	window.ShowAll()
	gtk.Main()
}
