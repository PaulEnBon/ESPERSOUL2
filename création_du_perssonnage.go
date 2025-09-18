package main

import (
	"fmt"

	"github.com/rivo/tview"
)

func CreationDuPersonnage(app *tview.Application, pages *tview.Pages) tview.Primitive {
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
		AddButton("Play", func() {
			currentMap := "salle1"
			fmt.Println("Initialisation de la partie dans la salle1...")
			app.Stop()              // stop tview
			RunGameLoop(currentMap) // start your console loop
		}).
		AddButton("Back", func() {
			pages.SwitchToPage("main")
		})

	flex := tview.NewFlex().
		AddItem(photoPlaceholder, 20, 1, false).
		AddItem(form, 0, 2, true)

	return flex
}
