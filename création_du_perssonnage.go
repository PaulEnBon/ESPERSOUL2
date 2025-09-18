package main

import (
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func CreationDuPersonnage(app *tview.Application, pages *tview.Pages) tview.Primitive {
	photoPlaceholder := tview.NewTextView().
		SetText("Choisissez votre classe à gauche.").
		SetTextAlign(tview.AlignCenter).
		SetBorder(true).
		SetTitle("Aperçu")

	// Zone d'aperçu des armes/armures
	details := tview.NewTextView()
	details.SetDynamicColors(true)
	details.SetBorder(true)
	details.SetTitle("Équipement")

	// Préparer la liste des classes depuis classes.go, en EXCLUANT les classes secrètes
	all := AllClasses()
	filtered := make([]Personnage, 0, len(all))
	for _, c := range all {
		if c.Nom == "Erwann" || c.Nom == "Gabriel" || c.Nom == "Vitaly" {
			continue
		}
		filtered = append(filtered, c)
	}
	classNames := make([]string, 0, len(filtered))
	for _, c := range filtered {
		classNames = append(classNames, c.Nom)
	}
	selectedIdx := 0
	selectedClass := func() Personnage { return filtered[selectedIdx] }

	updateDetails := func(p Personnage) {
		var b strings.Builder
		fmt.Fprintf(&b, "[yellow]Classe:[-] %s\n\n", p.Nom)
		fmt.Fprintf(&b, "[white]Armes disponibles:[-]\n")
		for i, a := range p.ArmesDisponibles {
			fmt.Fprintf(&b, "  %d. %s (ATK %d, MATK %d, PREC %.0f%%, CRIT %.0f%%)\n", i+1, a.Nom, a.DegatsPhysiques, a.DegatsMagiques, a.Precision*100, a.TauxCritique*100)
		}
		fmt.Fprintf(&b, "\n[white]Armures disponibles:[-]\n")
		for i, a := range p.ArmuresDisponibles {
			fmt.Fprintf(&b, "  %d. %s (DEF %d, RESM %d, HP +%d)\n", i+1, a.Nom, a.Defense, a.Resistance, a.HP)
		}
		details.SetText(b.String())
	}

	// Formulaire
	form := tview.NewForm()
	form.AddDropDown("Personnage:", classNames, 0, func(option string, index int) {
		selectedIdx = index
		updateDetails(selectedClass())
	})
	// Champ secret masqué
	secretField := tview.NewInputField().
		SetLabel("Secret:").
		SetMaskCharacter('*')
	form.AddFormItem(secretField)

	// Helper: démarre une partie avec une classe donnée
	startWith := func(chosen Personnage) {
		currentPlayer.Nom = chosen.Nom
		currentPlayer.ArmesDisponibles = chosen.ArmesDisponibles
		currentPlayer.ArmuresDisponibles = chosen.ArmuresDisponibles
		currentPlayer.NiveauArme = 0
		currentPlayer.NiveauArmure = 0
		currentPlayer.Roches = chosen.Roches
		currentPlayer.ArtefactsEquipes = make([]*Artefact, MaxArtefactsEquipes)
		RecomputeFromBaseAndEquip(&currentPlayer)
		currentMap := "salle1"
		// Create an initial save in the first empty slot
		_ = ensureSavesDir()
		_ = SaveInitialNewGame(currentMap, 8, 5)
		fmt.Println("Initialisation de la partie dans la salle1...")
		app.Stop()
		RunGameLoop(currentMap)
	}

	// Validation du champ secret par Entrée → sélection automatique de la classe correspondante
	secretField.SetDoneFunc(func(key tcell.Key) {
		if key != tcell.KeyEnter {
			return
		}
		code := strings.TrimSpace(strings.ToLower(secretField.GetText()))
		switch code {
		case "mac":
			startWith(Erwann)
		case "jean":
			startWith(Gabriel)
		case "vodka":
			startWith(Vitaly)
		default:
			// Rien: on ne fait pas démarrer si le code est invalide
		}
	})
	form.AddButton("Play", func() {
		// Si un code secret valide est entré, ignorer la sélection et charger la classe secrète
		code := strings.TrimSpace(strings.ToLower(secretField.GetText()))
		switch code {
		case "mac":
			startWith(Erwann)
			return
		case "jean":
			startWith(Gabriel)
			return
		case "vodka":
			startWith(Vitaly)
			return
		}
		// Sinon, démarrer avec la classe sélectionnée dans la liste filtrée
		chosen := selectedClass()
		startWith(chosen)
	})
	form.AddButton("Back", func() { pages.SwitchToPage("main") })

	// Mettre à jour l'aperçu initial
	updateDetails(selectedClass())

	// Layout: photo/aperçu à gauche, form au centre, détails à droite
	left := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(photoPlaceholder, 7, 1, false).
		AddItem(details, 0, 1, false)

	flex := tview.NewFlex().
		AddItem(left, 0, 2, false).
		AddItem(form, 0, 3, true)

	return flex
}

// Variant of the creation screen that reserves a specific save slot for the new game
func CreationDuPersonnageWithSave(app *tview.Application, pages *tview.Pages, targetSlot int) tview.Primitive {
	photoPlaceholder := tview.NewTextView().
		SetText("Choisissez votre classe à gauche.").
		SetTextAlign(tview.AlignCenter).
		SetBorder(true).
		SetTitle("Aperçu")

	details := tview.NewTextView()
	details.SetDynamicColors(true)
	details.SetBorder(true)
	details.SetTitle("Équipement")

	all := AllClasses()
	filtered := make([]Personnage, 0, len(all))
	for _, c := range all {
		if c.Nom == "Erwann" || c.Nom == "Gabriel" || c.Nom == "Vitaly" {
			continue
		}
		filtered = append(filtered, c)
	}
	classNames := make([]string, 0, len(filtered))
	for _, c := range filtered {
		classNames = append(classNames, c.Nom)
	}
	selectedIdx := 0
	selectedClass := func() Personnage { return filtered[selectedIdx] }

	updateDetails := func(p Personnage) {
		var b strings.Builder
		fmt.Fprintf(&b, "[yellow]Classe:[-] %s\n\n", p.Nom)
		fmt.Fprintf(&b, "[white]Armes disponibles:[-]\n")
		for i, a := range p.ArmesDisponibles {
			fmt.Fprintf(&b, "  %d. %s (ATK %d, MATK %d, PREC %.0f%%, CRIT %.0f%%)\n", i+1, a.Nom, a.DegatsPhysiques, a.DegatsMagiques, a.Precision*100, a.TauxCritique*100)
		}
		fmt.Fprintf(&b, "\n[white]Armures disponibles:[-]\n")
		for i, a := range p.ArmuresDisponibles {
			fmt.Fprintf(&b, "  %d. %s (DEF %d, RESM %d, HP +%d)\n", i+1, a.Nom, a.Defense, a.Resistance, a.HP)
		}
		details.SetText(b.String())
	}

	form := tview.NewForm()
	form.AddDropDown("Personnage:", classNames, 0, func(option string, index int) {
		selectedIdx = index
		updateDetails(selectedClass())
	})
	secretField := tview.NewInputField().
		SetLabel("Secret:").
		SetMaskCharacter('*')
	form.AddFormItem(secretField)

	startWith := func(chosen Personnage) {
		currentPlayer.Nom = chosen.Nom
		currentPlayer.ArmesDisponibles = chosen.ArmesDisponibles
		currentPlayer.ArmuresDisponibles = chosen.ArmuresDisponibles
		currentPlayer.NiveauArme = 0
		currentPlayer.NiveauArmure = 0
		currentPlayer.Roches = chosen.Roches
		currentPlayer.ArtefactsEquipes = make([]*Artefact, MaxArtefactsEquipes)
		RecomputeFromBaseAndEquip(&currentPlayer)
		currentMap := "salle1"
		// Reserve target slot immediately
		_ = ensureSavesDir()
		currentMapGlobalRef = currentMap
		if targetSlot >= 1 && targetSlot <= 4 {
			_ = SaveToSlot(targetSlot)
		} else {
			_ = SaveInitialNewGame(currentMap, 8, 5)
		}
		fmt.Println("Initialisation de la partie dans la salle1...")
		app.Stop()
		RunGameLoop(currentMap)
	}

	secretField.SetDoneFunc(func(key tcell.Key) {
		if key != tcell.KeyEnter {
			return
		}
		code := strings.TrimSpace(strings.ToLower(secretField.GetText()))
		switch code {
		case "mac":
			startWith(Erwann)
		case "jean":
			startWith(Gabriel)
		case "vodka":
			startWith(Vitaly)
		}
	})
	form.AddButton("Play",
		func() {
			code := strings.TrimSpace(strings.ToLower(secretField.GetText()))
			switch code {
			case "mac":
				startWith(Erwann)
				return
			case "jean":
				startWith(Gabriel)
				return
			case "vodka":
				startWith(Vitaly)
				return
			}
			chosen := selectedClass()
			startWith(chosen)
		})
	form.AddButton("Back", func() { pages.SwitchToPage("main") })

	updateDetails(selectedClass())

	left := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(photoPlaceholder, 7, 1, false).
		AddItem(details, 0, 1, false)

	flex := tview.NewFlex().
		AddItem(left, 0, 2, false).
		AddItem(form, 0, 3, true)

	return flex
}
