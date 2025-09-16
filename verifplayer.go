package main

import "fmt"

func ArmeEquipee(p *Personnage) Arme {
	if p.NiveauArme < 0 || p.NiveauArme >= len(p.ArmesDisponibles) {
		fmt.Printf("Niveau Arme hors limites : %d\n", p.NiveauArme)
		return Arme{}
	}
	fmt.Printf("Arme choisie : %s\n", p.ArmesDisponibles[p.NiveauArme].Nom)
	return p.ArmesDisponibles[p.NiveauArme]
}

func ArmureEquipee(p *Personnage) Armure {
	if p.NiveauArmure < 0 || p.NiveauArmure >= len(p.ArmuresDisponibles) {
		fmt.Printf("Niveau Armure hors limites : %d\n", p.NiveauArmure)
		return Armure{} // Armure vide si niveau invalide
	}
	fmt.Printf("Armure choisie : %s\n", p.ArmuresDisponibles[p.NiveauArmure].Nom)
	return p.ArmuresDisponibles[p.NiveauArmure]
}

func AfficherStats(p *Personnage) {
	arme := ArmeEquipee(p)
	armure := ArmureEquipee(p)

	fmt.Printf("=== Stats de %s ===\n", p.Nom)
	fmt.Printf("PV: %d / %d\n", p.PV, p.PVMax)
	fmt.Printf("Armure: %d\n", p.Armure)
	fmt.Printf("Résistance Magique: %d\n", p.ResistMag)
	fmt.Printf("Précision: %.2f\n", p.Precision)
	fmt.Printf("Taux Critique: %.2f\n", p.TauxCritique)
	fmt.Printf("Niveau Arme: %d (%s)\n", p.NiveauArme, arme.Nom)
	fmt.Printf("  - Dégâts Physiques: %d\n", arme.DegatsPhysiques)
	fmt.Printf("  - Dégâts Magiques: %d\n", arme.DegatsMagiques)
	fmt.Printf("  - Précision Arme: %.2f\n", arme.Precision)
	fmt.Printf("  - Taux Critique Arme: %.2f\n", arme.TauxCritique)
	fmt.Printf("  - Durabilité: %d\n", arme.Durabilite)
	fmt.Printf("Niveau Armure: %d (%s)\n", p.NiveauArmure, armure.Nom)
	fmt.Printf("  - Défense: %d\n", armure.Defense)
	fmt.Printf("  - Résistance: %d\n", armure.Resistance)
	fmt.Printf("  - HP Bonus: %d\n", armure.HP)
	fmt.Println("====================")
}
