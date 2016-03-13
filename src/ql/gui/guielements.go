package gui

import (
	"github.com/andlabs/ui"
)

// CreateInputTextField returns a new generic input text field.
func CreateInputTextField(defaultText string, disabled bool) *ui.Entry {
	textField := ui.NewEntry()
	textField.SetText(defaultText)

	if disabled {
		textField.Disable()
	}

	return textField
}

// CreateCheckboxConditional returns a new checkbox.
func CreateCheckboxConditional() *ui.Checkbox {
	return ui.NewCheckbox("")
}

// CreateButton returns a new button with passed onClick callback.
func CreateButton(buttonText string, onClick func(*ui.Button)) *ui.Button {
	button := ui.NewButton(buttonText)
	button.OnClicked(onClick)

	return button
}

// CreateCheckboxConditional returns a new text label.
func CreateLabel(text string) *ui.Label {
	return ui.NewLabel(text)
}
