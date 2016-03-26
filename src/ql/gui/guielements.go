package gui

import (
	"github.com/andlabs/ui"
	"strings"
)

// createInputTextField returns a new generic input text field.
func createInputTextField(defaultText string, disabled bool) *ui.Entry {
	textField := ui.NewEntry()
	textField.SetText(defaultText)

	if disabled {
		textField.Disable()
	}

	return textField
}

// createCheckboxConditional returns a new checkbox.
func createCheckboxConditional() *ui.Checkbox {
	return ui.NewCheckbox("")
}

// createButton returns a new button with passed onClick callback.
func createButton(buttonText string, onClick func(*ui.Button)) *ui.Button {
	button := ui.NewButton(buttonText)
	button.OnClicked(onClick)

	return button
}

// createCheckboxConditional returns a new text label.
func createLabel(text string) *ui.Label {
	return ui.NewLabel(text)
}

func showMessageBoxForErrors(title string, errors []error, window *ui.Window) {
	ui.MsgBoxError(window, title, convertErrorStringListToString(errors))
}

// convertErrorStringListToString converts a list of errors to a concatenated error string and returns it.
func convertErrorStringListToString(errorStringList []error) string {
	errorStrings := []string{}

	for _, singleError := range errorStringList {
		errorStrings = append(errorStrings, singleError.Error())
	}

	return strings.Join(errorStrings, "\n\n")
}
