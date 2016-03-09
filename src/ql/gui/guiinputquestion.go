package gui

import (
	"errors"
	log "github.com/Sirupsen/logrus"
	"github.com/mattn/go-gtk/gtk"
	"ql/ast/expr/litexpr"
	"ql/ast/vari"
	"strconv"
)

type GUIInputQuestion struct {
	GUIQuestion
}

func CreateGUIInputQuestion(label string, questionType vari.VarType, callback func(interface{}, error)) GUIInputQuestion {
	questionLabel := createLabel(label)
	questionElement := createQuestionElement(questionType, callback)
	errorLabel := createLabel("")

	return GUIInputQuestion{GUIQuestion: GUIQuestion{questionLabel, questionElement, errorLabel}}
}

func createQuestionElement(questionType vari.VarType, callback func(interface{}, error)) gtk.IWidget {
	var GTKEntity gtk.IWidget

	switch questionType.(type) {
	case vari.BoolType:
		checkbox := CreateCheckboxConditional()
		checkbox.Connect("clicked", func() {
			log.WithFields(log.Fields{"value": checkbox.GetActive()}).Debug("Checkbox value changed")
			callback(litexpr.BoolLit{checkbox.GetActive()}, nil)
		})
		GTKEntity = checkbox
	case vari.StringType:
		inputField := CreateInputTextField(questionType.GetDefaultValue().(litexpr.Lit).String())
		inputField.Connect("changed", func() {
			inputText := inputField.GetText()

			log.WithFields(log.Fields{"value": inputText}).Debug("Input text value changed")

			callback(litexpr.StrLit{inputText}, nil) // TODO check if really is string
		})
		GTKEntity = inputField
	case vari.IntType:
		inputField := CreateInputTextField(questionType.GetDefaultValue().(litexpr.Lit).String())
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

			callback(litexpr.IntLit{inputTextAsInt}, nil)
		})
		GTKEntity = inputField
	default:
		errors.New("Unknown question type, can not create correct GTK object")
	}

	return GTKEntity
}
