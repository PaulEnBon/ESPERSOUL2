package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// GameOverChoice enumerates possible choices from the Game Over screen
type GameOverChoice int

const (
	GameOverContinue GameOverChoice = iota // Respawn at spawn in salle1
	GameOverTitle                          // Return to title screen
)

// ShowGameOverScreen displays a modal Game Over screen and returns the user's choice.
func ShowGameOverScreen() GameOverChoice {
	app := tview.NewApplication()

	// Title text
	title := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true)
	title.SetText("[red::b]\n\n   GAME OVER\n\n[-]")

	// Buttons
	var choice GameOverChoice = GameOverContinue
	continueBtn := tview.NewButton("CONTINUER").SetSelectedFunc(func() {
		choice = GameOverContinue
		app.Stop()
	})
	titleBtn := tview.NewButton("RETOUR TITRE").SetSelectedFunc(func() {
		choice = GameOverTitle
		app.Stop()
	})

	// Style
	styleButton := func(b *tview.Button) {
		b.SetLabelColor(tcell.ColorWhite).
			SetLabelColorActivated(tcell.ColorBlack).
			SetBackgroundColorActivated(tcell.ColorWhite)
	}
	styleButton(continueBtn)
	styleButton(titleBtn)

	// Layout
	row := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(nil, 0, 1, false).
		AddItem(continueBtn, 18, 0, true).
		AddItem(nil, 2, 0, false).
		AddItem(titleBtn, 18, 0, false).
		AddItem(nil, 0, 1, false)

	layout := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(nil, 0, 2, false).
		AddItem(title, 5, 0, false).
		AddItem(nil, 1, 0, false).
		AddItem(row, 1, 0, true).
		AddItem(nil, 0, 2, false)

	if err := app.EnableMouse(true).SetRoot(layout, true).Run(); err != nil {
		// In case of UI failure, default to continue
		return GameOverContinue
	}
	return choice
}
