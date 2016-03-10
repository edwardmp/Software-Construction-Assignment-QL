package gui

import (
	"github.com/mattn/go-gtk/gtk"
)

// convenience method
func CreateDisabledInputTextField(defaultText string) *gtk.Entry {
	entry := CreateInputTextField(defaultText)
	entry.SetEditable(false)
	return entry
}

func CreateInputTextField(defaultText string) *gtk.Entry {
	entry := gtk.NewEntry()
	entry.SetText(defaultText)
	return entry
}

func CreateCheckboxConditional() *gtk.CheckButton {
	checkbutton := gtk.NewCheckButtonWithLabel("")

	return checkbutton
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
