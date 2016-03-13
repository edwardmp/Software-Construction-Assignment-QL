package gui

import (
	"errors"
	log "github.com/Sirupsen/logrus"
	"github.com/andlabs/ui"
	"ql/ast/expr"
	"ql/ast/vari"
	"ql/interfaces"
	"strconv"
)

type GUIQuestion struct {
	Label      *ui.Label
	Element    ui.Control
	ErrorLabel *ui.Label
}

func CreateEnabledGUIQuestion(label string, questionType interfaces.VarType, callback func(interface{}, error)) GUIQuestion {
	return createGUIQuestion(label, questionType, callback, false)
}

func CreateDisabledGUIQuestion(label string, questionType interfaces.VarType, callback func(interface{}, error)) GUIQuestion {
	return createGUIQuestion(label, questionType, callback, true)
}

func createGUIQuestion(label string, questionType interfaces.VarType, callback func(interface{}, error), disabled bool) GUIQuestion {
	questionLabel := CreateLabel(label)
	questionElement := createQuestionElement(questionType, callback, disabled)
	errorLabel := CreateLabel("")

	return GUIQuestion{questionLabel, questionElement, errorLabel}
}

func (g GUIQuestion) ChangeElementText(newText string) {
	text := g.Label.Text()
	log.WithFields(log.Fields{"oldLabelText": text, "newLabelText": newText}).Debug("Changing text of element")
	g.Element.(*ui.Entry).SetText(newText)
}

func (g GUIQuestion) ChangeErrorLabelText(newText string) {
	g.ErrorLabel.SetText(newText)
}

func createQuestionElement(questionType interfaces.VarType, callback func(interface{}, error), disabled bool) ui.Control {
	var UIEntity ui.Control

	switch questionType.(type) {
	case vari.BoolType:
		checkbox := CreateCheckboxConditional()
		checkbox.OnToggled(func(*ui.Checkbox) {
			log.WithFields(log.Fields{"value": checkbox.Checked()}).Debug("Checkbox value changed")
			callback(expr.NewBoolLitNoSourceInfo(checkbox.Checked()), nil)
		})
		UIEntity = checkbox
	case vari.StringType:
		inputField := CreateInputTextField("", disabled)
		inputField.OnChanged(func(*ui.Entry) {
			inputText := inputField.Text()

			log.WithFields(log.Fields{"value": inputText}).Debug("Input text value changed (string field)")

			callback(expr.NewStrLitNoSourceInfo(inputText), nil) // TODO check if really is string
		})
		UIEntity = inputField
	case vari.IntType:
		inputField := CreateInputTextField("", disabled)
		inputField.OnChanged(func(*ui.Entry) {
			inputText := inputField.Text()
			log.WithFields(log.Fields{"value": inputText}).Debug("Input text value changed (int field)")

			inputTextAsInt, err := strconv.Atoi(inputText)
			if inputText == "" {
				if callback != nil {
					callback(questionType.GetDefaultValue(), nil)
				}
				return
			} else if err != nil {
				log.Warn("Could not convert input text string to int")

				if callback != nil {
					callback(nil, err)
				}

				return
			}

			if callback != nil {
				callback(expr.NewIntLitNoSourceInfo(inputTextAsInt), nil)
			}
		})
		UIEntity = inputField
	default:
		errors.New("Unknown question type, can not create correct GTK object")
	}

	return UIEntity
}
