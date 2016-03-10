package gui

import (
	log "github.com/Sirupsen/logrus"
	"github.com/mattn/go-gtk/gtk"
	"ql/ast/vari"
)

type GUIQuestion struct {
	Label      *gtk.Label
	Element    gtk.IWidget
	ErrorLabel *gtk.Label
}

func CreateGUIQuestion(label string, questionType vari.VarType, callback func(interface{}, error)) GUIQuestion {
	questionLabel := CreateLabel(label)
	questionElement := createQuestionElement(questionType, callback)
	errorLabel := CreateLabel("")

	return GUIQuestion{questionLabel, questionElement, errorLabel}
}

func (g GUIQuestion) ChangeElementText(newText string) {
	log.WithFields(log.Fields{"oldLabelText": g.Label.GetText(), "newLabelText": newText}).Debug("Changing text of element")
	g.Element.(*gtk.Entry).SetText(newText)
}

func (g GUIQuestion) ChangeErrorLabelText(newText string) {
	g.ErrorLabel.SetText(newText)
}
