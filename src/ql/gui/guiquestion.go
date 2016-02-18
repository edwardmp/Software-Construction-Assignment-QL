package gui

import (
	"errors"
	"fmt"
	"github.com/mattn/go-gtk/gtk"
	"ql/ast/vari"
)

type GUIQuestion struct {
	Label   *gtk.Label
	Element gtk.IWidget
}

func CreateGUIQuestion(label string, questionType vari.VarType) GUIQuestion {
	questionLabel := createQuestionLabel(label)
	questionElement := createQuestionElement(questionType)

	return GUIQuestion{questionLabel, questionElement}
}

func createQuestionLabel(questionText string) *gtk.Label {
	label := gtk.NewLabel(questionText)
	label.ModifyFontEasy("DejaVu Serif 12")

	return label
}

func createQuestionElement(questionType vari.VarType) gtk.IWidget {
	var GTKEntity gtk.IWidget

	switch questionType {
	case vari.BOOLEAN:
		GTKEntity = CreateRadioButtons()
	case vari.STRING, vari.INT:
		GTKEntity = CreateInputTextField(fmt.Sprintf("Type: %v", questionType))
	default:
		errors.New("Unknown question type, can not create correct GTK object")
	}

	return GTKEntity
}
