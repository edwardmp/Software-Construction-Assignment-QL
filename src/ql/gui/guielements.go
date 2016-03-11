package gui

import (
	"github.com/mattn/go-gtk/gtk"
)

func CreateInputTextField(defaultText string, disabled bool) *gtk.Entry {
	textField := gtk.NewEntry()
	textField.SetText(defaultText)

	if disabled {
		textField.SetEditable(false)
	}

	return textField
}

func CreateCheckboxConditional() *gtk.CheckButton {
	return gtk.NewCheckButtonWithLabel("")
}

func CreateButton(buttonText string, onClick func()) *gtk.Button {
	button := gtk.NewButtonWithLabel(buttonText)
	button.Connect("clicked", onClick)
	return button
}

func CreateLabel(text string) *gtk.Label {
	label := gtk.NewLabel(text)
	label.ModifyFontEasy("DejaVu Serif 12")

	return label
}
