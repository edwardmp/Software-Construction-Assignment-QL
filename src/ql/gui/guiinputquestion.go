package gui

import (
	"errors"
	log "github.com/Sirupsen/logrus"
	"github.com/mattn/go-gtk/gtk"
	"ql/ast/expr/lit"
	"ql/ast/vari/vartype"
	"strconv"
)

type GUIInputQuestion struct {
	GUIQuestion
}

func CreateGUIInputQuestion(label string, questionType vartype.VarType, callback func(interface{}, error)) GUIInputQuestion {
	questionLabel := createLabel(label)
	questionElement := createQuestionElement(questionType, callback)
	errorLabel := createLabel("")

	return GUIInputQuestion{GUIQuestion: GUIQuestion{questionLabel, questionElement, errorLabel}}
}

func createQuestionElement(questionType vartype.VarType, callback func(interface{}, error)) gtk.IWidget {
	var GTKEntity gtk.IWidget

	switch questionType.(type) {
	case vartype.BoolType:
		checkbox := CreateCheckboxConditional()
		checkbox.Connect("clicked", func() {
			log.WithFields(log.Fields{"value": checkbox.GetActive()}).Debug("First input radio button value changed")
			callback(lit.BoolLit{checkbox.GetActive()}, nil)
		})
		GTKEntity = checkbox
	case vartype.StringType:
		inputField := CreateInputTextField(questionType.GetDefaultValue().(lit.Lit).String())
		inputField.Connect("changed", func() {
			inputText := inputField.GetText()

			log.WithFields(log.Fields{"value": inputText}).Debug("Input text value changed")

			callback(lit.StrLit{inputText}, nil) // TODO check if really is string
		})
		GTKEntity = inputField
	case vartype.IntType:
		inputField := CreateInputTextField(questionType.GetDefaultValue().(lit.Lit).String())
		inputField.Connect("changed", func() {
			inputText := inputField.GetText()
			log.WithFields(log.Fields{"value": inputText}).Debug("Input text value changed")

			inputTextAsInt, err := strconv.Atoi(inputText)
			if inputText == "" {
				callback(questionType.GetDefaultValue(), nil)
				return
			} else if err != nil {
				log.Warn("Could not convert input text string to int")
				callback(nil, err)
				return
			}

			callback(lit.IntLit{inputTextAsInt}, nil)
		})
		GTKEntity = inputField
	default:
		errors.New("Unknown question type, can not create correct GTK object")
	}

	return GTKEntity
}
