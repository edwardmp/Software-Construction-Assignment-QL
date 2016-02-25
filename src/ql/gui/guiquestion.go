package gui

import (
	"errors"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/mattn/go-gtk/gtk"
	"ql/ast/expr/lit"
	"ql/ast/vari"
	"strconv"
)

type GUIQuestion struct {
	Label      *gtk.Label
	Element    gtk.IWidget
	ErrorLabel *gtk.Label
}

func (g GUIQuestion) ChangeElementText(newText string) {
	g.Element.(*gtk.Entry).SetText(newText)
}

func (g GUIQuestion) ChangeErrorLabelText(newText string) {
	g.ErrorLabel.SetText(newText)
}

func CreateGUIInputQuestion(label string, questionType vari.VarType, callback func(interface{}, error)) GUIQuestion {
	questionLabel := createQuestionLabel(label)
	questionElement := createQuestionElement(questionType, callback)
	errorLabel := createQuestionLabel("")

	return GUIQuestion{questionLabel, questionElement, errorLabel}
}

func CreateGUIComputedQuestion(label string, elementValueString string, callback func()) GUIQuestion {
	questionLabel := createQuestionLabel(label)
	questionElement := CreateDisabledInputTextField(elementValueString)
	errorLabel := createQuestionLabel("")
	questionElement.Connect("changed", callback)

	return GUIQuestion{questionLabel, questionElement, errorLabel}
}

func createQuestionLabel(questionText string) *gtk.Label {
	label := gtk.NewLabel(questionText)
	label.ModifyFontEasy("DejaVu Serif 12")

	return label
}

func createQuestionElement(questionType vari.VarType, callback func(interface{}, error)) gtk.IWidget {
	var GTKEntity gtk.IWidget

	switch questionType {
	case vari.BOOLEAN:
		checkbox := CreateCheckboxConditional()
		checkbox.Connect("clicked", func() {
			log.WithFields(log.Fields{"value": checkbox.GetActive()}).Debug("First input radio button value changed")
			callback(lit.BoolLit{true}, nil)
		})
		GTKEntity = checkbox
	case vari.STRING:
		inputField := CreateInputTextField(fmt.Sprintf("Type: %v", questionType))
		inputField.Connect("changed", func() {
			inputText := inputField.GetText()
			log.WithFields(log.Fields{"value": inputText}).Debug("Input text value changed:")

			callback(lit.StrLit{inputText}, nil) // TODO check if really is string
		})
		GTKEntity = inputField
	case vari.INT:
		inputField := CreateInputTextField("")
		inputField.Connect("changed", func() {
			inputText := inputField.GetText()
			log.WithFields(log.Fields{"value": inputText}).Debug("Input text value changed:")

			inputTextAsInt, err := strconv.Atoi(inputText)
			if err != nil {
				log.Warn("Could not convert input text string to int")
				callback(nil, err)
			}

			callback(lit.IntLit{inputTextAsInt}, nil)
		})
		GTKEntity = inputField
	default:
		errors.New("Unknown question type, can not create correct GTK object")
	}

	return GTKEntity
}
