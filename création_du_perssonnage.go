package main

import (
	"fmt"
	"strings"

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

	// Préparer la liste des classes depuis classes.go
	classNames := AllClassNames()
	selectedIdx := 0
	selectedClass := func() Personnage { return AllClasses()[selectedIdx] }

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
	form.AddCheckbox("Hard mode:", false, nil)
	form.AddButton("Play", func() {
		// Appliquer la classe choisie au joueur courant
		chosen := selectedClass()
		currentPlayer.Nom = chosen.Nom
		currentPlayer.ArmesDisponibles = chosen.ArmesDisponibles
		currentPlayer.ArmuresDisponibles = chosen.ArmuresDisponibles
		currentPlayer.NiveauArme = 0
		currentPlayer.NiveauArmure = 0
		currentPlayer.Roches = chosen.Roches
		currentPlayer.ArtefactsEquipes = make([]*Artefact, MaxArtefactsEquipes)
		// Recompute base stats and apply first equipment
		RecomputeFromBaseAndEquip(&currentPlayer)

		currentMap := "salle1"
		fmt.Println("Initialisation de la partie dans la salle1...")
		app.Stop()
		RunGameLoop(currentMap)
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
