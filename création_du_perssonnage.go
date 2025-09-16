package main

import (
	"github.com/rivo/tview"
)

// This function now returns the layout instead of running a new app.
func Création_du_perssonnage(pages *tview.Pages) tview.Primitive {
	photoPlaceholder := tview.NewTextView().
		SetText("kitten icat image.jpeg").
		SetTextAlign(tview.AlignCenter).
		SetBorder(true).
		SetTitle("Photo")

	form := tview.NewForm().
		AddDropDown("Perssonnage:", []string{
			"Engineer",
			"Manager",
			"Administration",
		}, 0, nil).
		AddDropDown("Artéfact:", []string{
			"ytftyfyty",
		}, 0, nil).
		AddCheckbox("Hard mode:", false, nil).
		AddPasswordField("Secret:", "", 10, '*', nil).
		AddButton("Play", func() {
			pages.SwitchToPage("main")
		}).
		AddButton("Back", func() {
			// Go back to main menu
			pages.SwitchToPage("main")
		})

	flex := tview.NewFlex().
		AddItem(photoPlaceholder, 20, 1, false).
		AddItem(form, 0, 2, true)

	return flex
}
