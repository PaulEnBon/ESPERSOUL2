package main

import (
	"github.com/rivo/tview"
)

func Ecrant_de_départ() {
	asciiArt := `
/$$$$$$$$  /$$$$$$  /$$$$$$$  /$$$$$$$$ /$$$$$$$         /$$$$$$   /$$$$$$  /$$   /$$ /$$              /$$$$$$ 
| $$_____/ /$$__  $$| $$__  $$| $$_____/| $$__  $$       /$$__  $$ /$$__  $$| $$  | $$| $$             /$$__  $$
| $$      | $$  \__/| $$  \ $$| $$      | $$  \ $$      | $$  \__/| $$  \ $$| $$  | $$| $$            |__/  \ $$
| $$$$$   |  $$$$$$ | $$$$$$$/| $$$$$   | $$$$$$$/      |  $$$$$$ | $$  | $$| $$  | $$| $$              /$$$$$$/
| $$__/    \____  $$| $$____/ | $$__/   | $$__  $$       \____  $$| $$  | $$| $$  | $$| $$             /$$____/ 
| $$       /$$  \ $$| $$      | $$      | $$  \ $$       /$$  \ $$| $$  | $$| $$  | $$| $$            | $$      
| $$$$$$$$|  $$$$$$/| $$      | $$$$$$$$| $$  | $$      |  $$$$$$/|  $$$$$$/|  $$$$$$/| $$$$$$$$      | $$$$$$$$
|________/ \______/ |__/      |________/|__/  |__/       \______/  \______/  \______/ |________/      |________/`

	textView := tview.NewTextView().
		SetText(asciiArt).
		SetTextAlign(tview.AlignCenter) // Center the ASCII art horizontally

	// Optional: set dynamic word wrap off to preserve formatting
	textView.SetWrap(false)

	// Optional: Set border and title
	textView.SetBorder(true).SetTitle("")

	// Run the application with the TextView as root
	if err := tview.NewApplication().SetRoot(textView, true).Run(); err != nil {
		panic(err)
	}
}
