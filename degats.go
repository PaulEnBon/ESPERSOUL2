package main

import (
	"math"
	"math/rand"
)

// ==========================
// definit si une attaque touche ou non sa cible
// ==========================

// Calcule la précision finale avec modificateurs
func CalculerPrecision(personnage *Personnage) float64 {
	_, _, modifPrecision, _, _, _ := CalculerModificateurs(personnage)
	precision := personnage.Precision * modifPrecision

	// S'assure que la précision reste dans [0.0, 1.0]
	if precision > 1.0 {
		precision = 1.0
	}
	if precision < 0.0 {
		precision = 0.0
	}

	return precision
}

// Vérifie si une attaque touche ou est esquivée
func AttaqueReussie(attaquant *Personnage) bool {
	// Si l'artefact 'Lunette d'Erwann' est équipé => touche toujours
	if PossedeArtefact(attaquant, "Lunette d'Erwann") {
		return true
	}
	tauxPrecision := CalculerPrecision(attaquant)
	return rand.Float64() < tauxPrecision
}

// ==========================
// definit si un coup est crtique
// ==========================

// Calcule le taux critique final avec modificateurs
func CalculerTauxCritique(personnage *Personnage) float64 {
	_, _, _, modifCritique, _, _ := CalculerModificateurs(personnage)
	critique := personnage.TauxCritique * modifCritique

	// Clamp entre 0 et 1
	if critique > 1.0 {
		critique = 1.0
	} else if critique < 0.0 {
		critique = 0.0
	}

	return critique
}

// Vérifie si une attaque est un coup critique
func EstCoupCritique(attaquant *Personnage) bool {
	// ⚡ éviter de reseed à chaque appel (coûteux et pas nécessaire)
	// tu peux initialiser rand.Seed() UNE FOIS dans init() ou main()
	tauxCritique := CalculerTauxCritique(attaquant)
	return rand.Float64() < tauxCritique
}

// ==========================
// calcul des degats physique ou magique en fonction du type de la compétence
// // ==========================

func CalculerDegatsPhysiques(attaquant, defenseur *Personnage, degatsBase int) int {
	modifDegatsPhys, _, _, _, _, _ := CalculerModificateurs(attaquant)

	for _, effetActif := range attaquant.EffetsActifs {
		if effetActif.Effet.Peur {
			modifDegatsPhys += effetActif.Effet.ModifDegats
		}
	}

	armure := defenseur.Armure
	reduction := 1.0 - (float64(armure) / (100.0 + float64(armure)))

	degats := float64(degatsBase) * modifDegatsPhys * reduction
	if degats < 1 {
		degats = 1
	}
	return int(math.Round(degats))
}

func CalculerDegatsMagiques(attaquant, defenseur *Personnage, degatsBase int) int {
	_, modifDegatsMag, _, _, _, _ := CalculerModificateurs(attaquant)

	resMag := defenseur.ResistMag
	reduction := 1.0 - (float64(resMag) / (100.0 + float64(resMag)))

	degats := float64(degatsBase) * modifDegatsMag * reduction
	if degats < 1 {
		degats = 1
	}
	return int(math.Round(degats))
}

// ==========================
// modification des stats en fonction des alterations
// ==========================

// Calcule les modificateurs actifs sur un personnage (dégâts, précision, critique, armure, résistance)
func CalculerModificateurs(personnage *Personnage) (
	modifDegatsPhys, modifDegatsMag, modifPrecision, modifCritique, modifArmure, modifResistMag float64,
) {
	modifDegatsPhys = 1.0
	modifDegatsMag = 1.0
	modifPrecision = 1.0
	modifCritique = 1.0
	modifArmure = 1.0
	modifResistMag = 1.0

	for _, effetActif := range personnage.EffetsActifs {
		effet := effetActif.Effet

		modifDegatsPhys += effet.ModifDegats
		modifDegatsMag += effet.ModifDegatsMag
		modifPrecision += effet.ModifPrecision
		modifCritique += effet.ModifCritique
		modifArmure += effet.ModifArmure
		modifResistMag += effet.ModifResistMag
	}

	// S'assure que les modificateurs ne deviennent pas négatifs
	if modifDegatsPhys < 0.1 {
		modifDegatsPhys = 0.1
	}
	if modifDegatsMag < 0.1 {
		modifDegatsMag = 0.1
	}
	if modifPrecision < 0.1 {
		modifPrecision = 0.1
	}
	if modifCritique < 0.1 {
		modifCritique = 0.1
	}
	if modifArmure < 0.1 {
		modifArmure = 0.1
	}
	if modifResistMag < 0.1 {
		modifResistMag = 0.1
	}

	return
}

// ==========================
// attaque finale avec les degats, le crit et si elle touche ou non
// ==========================

func CalculerDegatsAvecCritique(attaquant, defenseur *Personnage, degatsBase int, typeAttaque string) (degatsFinaux int, estCritique bool, aTouche bool) {
	// Vérifie si l’attaque touche
	aTouche = AttaqueReussie(attaquant)
	if !aTouche {
		return 0, false, false // raté : 0 dégâts
	}

	var degats int
	if typeAttaque == "magique" {
		degats = CalculerDegatsMagiques(attaquant, defenseur, degatsBase)
	} else {
		degats = CalculerDegatsPhysiques(attaquant, defenseur, degatsBase)
	}

	// Vérifie critique
	estCritique = EstCoupCritique(attaquant)
	if estCritique {
		degats = int(float64(degats) * attaquant.MultiplicateurCrit)
	}

	return degats, estCritique, true
}
