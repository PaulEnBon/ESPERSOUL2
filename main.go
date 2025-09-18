package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
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

	// Buttons (no direct Play on home screen)

	settingsButton := tview.NewButton("Settings").SetSelectedFunc(func() {
		textView.SetText("You pressed [yellow]Settings[-]!")
	})

	quitButton := tview.NewButton("Quit").SetSelectedFunc(func() {
		app.Stop()
	})

	// Save manager page
	saveManager := func() tview.Primitive {
		list := tview.NewList().
			ShowSecondaryText(false)
		// Populate 4 slots
		refresh := func() {
			list.Clear()
			for i := 1; i <= 4; i++ {
				exists, line := ReadSlotSummary(i)
				idx := i
				if exists {
					list.AddItem(line, "Entrer pour charger, 'Suppr' pour effacer", '0'+rune(i), func() {
						// load
						go func(slot int) {
							_ = LoadFromSlot(app.Stop, slot)
						}(idx)
					})
				} else {
					list.AddItem(line, "Commencer une nouvelle partie ici", '0'+rune(i), func() {
						// Start new game flow -> character creation, then auto-save to this slot
						targetSlot := idx
						// Provide creation screen with a callback
						pages.AddAndSwitchToPage("character_creation", CreationDuPersonnageWithSave(app, pages, targetSlot), true)
					})
				}
			}
		}
		refresh()

		// Delete with Delete key
		list.SetDoneFunc(func() { pages.SwitchToPage("main") })
		list.SetSelectedFunc(func(index int, mainText, secondaryText string, shortcut rune) {
			// handled in item callbacks
		})
		list.SetInputCapture(func(ev *tcell.EventKey) *tcell.EventKey {
			if ev.Key() == tcell.KeyDelete {
				// Delete selected slot if exists
				idx := list.GetCurrentItem() + 1
				if err := DeleteSlot(idx); err == nil {
					refresh()
				}
				return nil
			}
			return ev
		})

		frame := tview.NewFrame(list).
			SetBorders(1, 1, 1, 1, 2, 2).
			AddText("Gestion des sauvegardes (Suppr pour effacer)", true, tview.AlignCenter, tcell.ColorYellow)
		return frame
	}

	savesButton := tview.NewButton("Play").SetSelectedFunc(func() {
		pages.AddAndSwitchToPage("saves", saveManager(), true)
	})

	// Button styling and focus indication
	styleButton := func(btn *tview.Button) {
		btn.SetLabelColor(tcell.ColorWhite).
			SetLabelColorActivated(tcell.ColorBlack).
			SetBackgroundColorActivated(tcell.ColorWhite)
	}
	styleButton(settingsButton)
	styleButton(quitButton)
	styleButton(savesButton)

	// Button layout (centered horizontally)
	buttonRow := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(nil, 0, 1, false).             // spacer
		AddItem(savesButton, 16, 0, true).     // Save manager button (focused)
		AddItem(nil, 2, 0, false).             // spacer
		AddItem(settingsButton, 12, 0, false). // Settings button
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

	if err := app.EnableMouse(true).SetRoot(pages, true).Run(); err != nil {
		panic(err)
	}
}

// RunGameLoop gère la boucle principale du jeu
// La fonction RunGameLoop complète est définie dans game_loop.go
