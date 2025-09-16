package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func Ecrant_de_départ() {
	app := tview.NewApplication()

	// ASCII Art
	asciiArt := `
/$$$$$$$$  /$$$$$$  /$$$$$$$  /$$$$$$$$ /$$$$$$$         /$$$$$$   /$$$$$$  /$$   /$$ /$$              /$$$$$$ 
| $$_____/ /$$__  $$| $$__  $$| $$_____/| $$__  $$       /$$__  $$ /$$__  $$| $$  | $$| $$             /$$__  $$
| $$      | $$  \__/| $$  \ $$| $$      | $$  \ $$      | $$  \__/| $$  \ $$| $$  | $$| $$            |__/  \ $$
| $$$$$   |  $$$$$$ | $$$$$$$/| $$$$$   | $$$$$$$/      |  $$$$$$ | $$  | $$| $$  | $$| $$              /$$$$$$/
| $$__/    \____  $$| $$____/ | $$__/   | $$__  $$       \____  $$| $$  | $$| $$  | $$| $$             /$$____/ 
| $$       /$$  \ $$| $$      | $$      | $$  \ $$       /$$  \ $$| $$  | $$| $$  | $$| $$            | $$      
| $$$$$$$$|  $$$$$$/| $$      | $$$$$$$$| $$  | $$      |  $$$$$$/|  $$$$$$/|  $$$$$$/| $$$$$$$$      | $$$$$$$$
|________/ \______/ |__/      |________/|__/  |__/       \______/  \______/  \______/ |________/      |________/`

	pages := tview.NewPages()

	// TextView for ASCII Art
	textView := tview.NewTextView().
		SetText(asciiArt).
		SetTextAlign(tview.AlignCenter).
		SetWrap(false).
		SetDynamicColors(true)

	// Buttons
	playButton := tview.NewButton("Play").SetSelectedFunc(func() {
		pages.SwitchToPage("second")
	})

	settingsButton := tview.NewButton("Settings").SetSelectedFunc(func() {
		textView.SetText("You pressed [yellow]Settings[-]!")
	})

	quitButton := tview.NewButton("Quit").SetSelectedFunc(func() {
		app.Stop()
	})

	Back := tview.NewButton("Back").SetSelectedFunc(func() {
		textView.SetText("You pressed [red]Back[-]!")
	})

	// Button styling and focus indication
	styleButton := func(btn *tview.Button) {
		btn.SetLabelColor(tcell.ColorWhite).
			SetLabelColorActivated(tcell.ColorBlack).
			SetBackgroundColorActivated(tcell.ColorWhite)
	}
	styleButton(playButton)
	styleButton(settingsButton)
	styleButton(quitButton)
	styleButton(Back)

	// Button layout (centered horizontally)
	buttonRow := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(nil, 0, 1, false).             // spacer
		AddItem(playButton, 12, 0, true).      // Play button
		AddItem(nil, 2, 0, false).             // spacer
		AddItem(settingsButton, 12, 0, false). // Settings button
		AddItem(nil, 2, 0, false).             // spacer
		AddItem(Back, 12, 0, false).           // Back button
		AddItem(nil, 2, 0, false).             // spacer
		AddItem(quitButton, 12, 0, false).     // Quit button
		AddItem(nil, 0, 1, false)              // spacer

	// Vertical layout
	mainLayout := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(nil, 0, 1, false).       // top spacer
		AddItem(textView, 10, 0, false). // ASCII art
		AddItem(nil, 1, 0, false).       // spacing
		AddItem(buttonRow, 1, 0, true).  // buttons
		AddItem(nil, 0, 1, false)        // bottom spacer

	pages.AddPage("main", mainLayout, true, true)

	// Start app
	if err := app.EnableMouse(true).SetRoot(pages, true).Run(); err != nil {
		panic(err)
	}
}
