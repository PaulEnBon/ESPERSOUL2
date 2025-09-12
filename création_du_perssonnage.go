package main

import (
	"github.com/rivo/tview"
)

func Création_du_perssonnage() {

	PhotoPlaceholder := tview.NewTextView().
		SetText("kitten icat image.jpeg").
		SetTextAlign(tview.AlignCenter).
		SetBorder(true).
		SetTitle("Photo")

	form := tview.NewForm().
		AddInputField("First name:", "", 20, nil, nil).
		AddInputField("Last name:", "", 20, nil, nil).
		AddDropDown("Role:", []string{
			"Engineer",
			"Manager",
			"Administration",
		}, 0, nil).
		AddCheckbox("On vacation:", false, nil).
		AddPasswordField("Password:", "", 10, '*', nil).
		AddTextArea("Notes:", "", 0, 5, 0, nil).
		AddButton("Save", func() { /* Save data */ }).
		AddButton("Cancel", func() { /* Cancel */ })

	flex := tview.NewFlex().
		AddItem(PhotoPlaceholder, 20, 1, false).
		AddItem(form, 0, 2, true)

	tview.NewApplication().
		SetRoot(flex, true).
		Run()
}
