package gui

import (
	"github.com/andlabs/ui"
)

func CreateInputTextField(defaultText string, disabled bool) *ui.Entry {
	textField := ui.NewEntry()
	textField.SetText(defaultText)

	if disabled {
		textField.Disable()
	}

	return textField
}

func CreateCheckboxConditional() *ui.Checkbox {
	return ui.NewCheckbox("")
}

func CreateButton(buttonText string, onClick func(*ui.Button)) *ui.Button {
	button := ui.NewButton("Submit")
	button.OnClicked(onClick)

	return button
}

func CreateLabel(text string) *ui.Label {
	return ui.NewLabel(text)
}
