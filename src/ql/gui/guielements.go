package gui

import (
	"fmt"
	"github.com/mattn/go-gtk/gtk"
)

func CreateRadioButtons() *gtk.HBox {
	buttons := gtk.NewHBox(false, 1)
	firstRadioButton := gtk.NewRadioButtonWithLabel(nil, "Yes")
	buttons.Add(firstRadioButton)
	buttons.Add(gtk.NewRadioButtonWithLabel(firstRadioButton.GetGroup(), "No"))

	return buttons
}

func CreateInputTextField(defaultText string) *gtk.Entry {
	entry := gtk.NewEntry()
	entry.SetText(defaultText)
	entry.Connect("changed", func() {
		fmt.Printf("Input value changed: %s\n", entry.GetText())
	})

	return entry
}

func CreateCheckboxConditional() *gtk.CheckButton {
	checkbutton := gtk.NewCheckButtonWithLabel("CheckButton with label")
	checkbutton.Connect("toggled", func() {
		if checkbutton.GetActive() {
			checkbutton.SetLabel("CheckButton CHECKED!")
		} else {
			checkbutton.SetLabel("CheckButton UNCHECKED!")
		}
	})

	return checkbutton
}

func CreateButton(buttonText string, onClick func()) *gtk.Button {
	button := gtk.NewButtonWithLabel(buttonText)
	button.Connect("clicked", onClick)
	return button
}
