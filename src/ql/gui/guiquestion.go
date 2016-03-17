package gui

import (
	"errors"
	log "github.com/Sirupsen/logrus"
	"github.com/andlabs/ui"
	"ql/ast/expr"
	"ql/interfaces"
	"strconv"
)

type GUIQuestion struct {
	Label      *ui.Label
	Element    ui.Control
	ErrorLabel *ui.Label
}

// createGUIQuestion creates a GUIQuestion. The last argument indicates if the question should be disabled (no entry allowed).
func createGUIQuestion(label string, questionType interfaces.ValueType, callback func(interfaces.Expr, error), disabled bool) *GUIQuestion {
	questionLabel := CreateLabel(label)
	questionElement := createQuestionElement(questionType, callback, disabled)
	errorLabel := CreateLabel("")

	return &GUIQuestion{questionLabel, questionElement, errorLabel}
}

// CreateEnabledGUIQuestion is a convenience method for creating a GUIQuestion that is enabled
func CreateEnabledGUIQuestion(label string, questionType interfaces.ValueType, callback func(interfaces.Expr, error)) *GUIQuestion {
	return createGUIQuestion(label, questionType, callback, false)
}

// CreateDisabledGUIQuestion is a convenience method for creating a GUIQuestion that is disabled
func CreateDisabledGUIQuestion(label string, questionType interfaces.ValueType, callback func(interfaces.Expr, error)) *GUIQuestion {
	return createGUIQuestion(label, questionType, callback, true)
}

//
func (this *GUIQuestion) ChangeFieldValueText(newText string) {
	log.WithFields(log.Fields{"newLabelText": newText}).Debug("Changing text of element")
	this.Element.(*ui.Entry).SetText(newText)
}

// ChangeErrorLabelText changes the error text feedback presented when error occurs
func (this *GUIQuestion) ChangeErrorLabelText(newText string) {
	this.ErrorLabel.SetText(newText)
}

// ResetErrorLabelText removes the error text presented to the user
func (this *GUIQuestion) ResetErrorLabelText(newText string) {
	this.ChangeErrorLabelText("")
}

func createQuestionElement(questionType interfaces.ValueType, callback func(interfaces.Expr, error), disabled bool) ui.Control {
	var UIEntity ui.Control

	switch questionType.(type) {
	case expr.BoolType:
		checkbox := CreateCheckboxConditional()
		checkbox.OnToggled(func(*ui.Checkbox) {
			log.WithFields(log.Fields{"value": checkbox.Checked()}).Debug("Checkbox value changed")

			callback(expr.NewBoolLitNoSourceInfo(checkbox.Checked()), nil)
		})
		UIEntity = checkbox
	case expr.StringType:
		inputField := CreateInputTextField("", disabled)
		inputField.OnChanged(func(*ui.Entry) {
			inputText := inputField.Text()

			log.WithFields(log.Fields{"value": inputText}).Debug("Input text value changed (string field)")

			callback(expr.NewStrLitNoSourceInfo(inputText), nil)
		})
		UIEntity = inputField
	case expr.IntType:
		inputField := CreateInputTextField("", disabled)
		inputField.OnChanged(func(*ui.Entry) {
			inputText := inputField.Text()

			log.WithFields(log.Fields{"value": inputText}).Debug("Input text value changed (int field)")

			inputTextAsInt, err := strconv.Atoi(inputText)
			if inputText == "" {
				if callback != nil {
					callback(expr.NewIntLitNoSourceInfo(inputTextAsInt), nil)
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
