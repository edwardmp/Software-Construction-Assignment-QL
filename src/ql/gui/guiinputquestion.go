package gui

import (
	"errors"
	log "github.com/Sirupsen/logrus"
	"github.com/mattn/go-gtk/gtk"
	"ql/ast/expr/lit"
	"ql/ast/vari"
	"strconv"
)

type GUIInputQuestion struct {
	GUIQuestion
}

func CreateGUIInputQuestion(label string, questionType vari.VarType, callback func(interface{}, error)) GUIInputQuestion {
	questionLabel := createQuestionLabel(label)
	questionElement := createQuestionElement(questionType, callback)
	errorLabel := createQuestionLabel("")

	return GUIInputQuestion{GUIQuestion: GUIQuestion{questionLabel, questionElement, errorLabel}}
}

func createQuestionElement(questionType vari.VarType, callback func(interface{}, error)) gtk.IWidget {
	var GTKEntity gtk.IWidget

	switch questionType {
	case vari.BOOLEAN:
		checkbox := CreateCheckboxConditional()
		checkbox.Connect("clicked", func() {
			log.WithFields(log.Fields{"value": checkbox.GetActive()}).Debug("First input radio button value changed")
			callback(lit.BoolLit{checkbox.GetActive()}, nil)
		})
		GTKEntity = checkbox
	case vari.STRING:
		inputField := CreateInputTextField("")
		inputField.Connect("changed", func() {
			inputText := inputField.GetText()
			log.WithFields(log.Fields{"value": inputText}).Debug("Input text value changed")

			callback(lit.StrLit{inputText}, nil) // TODO check if really is string
		})
		GTKEntity = inputField
	case vari.INT:
		inputField := CreateInputTextField("0")
		inputField.Connect("changed", func() {
			inputText := inputField.GetText()
			log.WithFields(log.Fields{"value": inputText}).Debug("Input text value changed")

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
