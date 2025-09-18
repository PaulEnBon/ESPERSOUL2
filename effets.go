package main

import (
	"fmt"
	"math/rand"
)

// --- Structures ---

// Structure principale d'un effet
// --- Structures ---
type Effet struct {
	Nom             string
	ToursRestants   int     // Nombre de tours restant pour l'effet
	DegatsParTour   int     // Dégâts infligés par tour (pour Poison, Brûlure, Saignement)
	SoinsParTour    int     // Soins par tour (pour régénération)
	ChanceAppliquer float64 // Probabilité d'application de l'effet (0.0 à 1.0)

	// Statuts de contrôle
	Estourdissement bool
	Peur            bool

	// Modificateurs de stats (en pourcentage)
	ModifPrecision float64 // Modification de précision (-0.15 = -15%)
	ModifCritique  float64 // Modification du taux critique
	ModifDegats    float64 // Modification des dégâts physiques
	ModifDegatsMag float64 // Modification des dégâts magiques
	ModifArmure    float64 // Modification de l'armure
	ModifResistMag float64 // Modification de la résistance magique
}

// Effet actif sur un personnage
type EffetActif struct {
	Effet         Effet
	ToursRestants int
}

// Détermine si un effet est un buff/debuff (inclut DOT/soins, contrôles, modifs de stats)
func isBuffOrDebuff(e Effet) bool {
	if e.DegatsParTour > 0 || e.SoinsParTour > 0 {
		return true
	}
	if e.Estourdissement || e.Peur {
		return true
	}
	if e.ModifPrecision != 0 || e.ModifCritique != 0 || e.ModifDegats != 0 || e.ModifDegatsMag != 0 || e.ModifArmure != 0 || e.ModifResistMag != 0 {
		return true
	}
	if e.Nom == "Renforcement" {
		return true
	}
	return false
}

// Applique un effet sur un personnage
func AppliquerEffet(cible *Personnage, effet Effet) bool {
	// RNG auto-seeded in Go 1.20+

	if rand.Float64() > effet.ChanceAppliquer {
		fmt.Printf("L'effet %s a échoué sur %s\n", effet.Nom, cible.Nom)
		return false
	}

	// Cap la durée des buffs/debuffs à 3 tours
	if isBuffOrDebuff(effet) && effet.ToursRestants > 3 {
		effet.ToursRestants = 3
	}

	// Vérifie si l'effet existe déjà et le met à jour ou l'ajoute
	for i := range cible.EffetsActifs {
		if cible.EffetsActifs[i].Effet.Nom == effet.Nom {
			// Renouvelle la durée
			nouv := effet.ToursRestants
			if isBuffOrDebuff(effet) && nouv > 3 {
				nouv = 3
			}
			cible.EffetsActifs[i].ToursRestants = nouv
			fmt.Printf("%s : effet %s renouvelé (%d tours)\n", cible.Nom, effet.Nom, effet.ToursRestants)
			return true
		}
	}

	// Ajoute le nouvel effet
	effetActif := EffetActif{
		Effet:         effet,
		ToursRestants: effet.ToursRestants,
	}
	cible.EffetsActifs = append(cible.EffetsActifs, effetActif)

	fmt.Printf("%s subit l'effet %s pendant %d tours\n", cible.Nom, effet.Nom, effet.ToursRestants)
	return true
}

// Traite les effets à la fin du tour d'un personnage
func TraiterEffetsFinTour(personnage *Personnage) (degatsRecus, soinsRecus int) {
	var effetsRestants []EffetActif

	for _, effetActif := range personnage.EffetsActifs {
		effet := effetActif.Effet

		// Applique les dégâts ou soins par tour
		if effet.DegatsParTour > 0 {
			degatsRecus += effet.DegatsParTour
			fmt.Printf("%s subit %d dégâts de %s\n", personnage.Nom, effet.DegatsParTour, effet.Nom)
		}

		if effet.SoinsParTour > 0 {
			soinsRecus += effet.SoinsParTour
			fmt.Printf("%s récupère %d PV grâce à %s\n", personnage.Nom, effet.SoinsParTour, effet.Nom)
		}

		// Gestion spéciale pour la brûlure (50% chance de s'éteindre)
		if effet.Nom == "Brûlure" {
			if rand.Float64() < 0.5 {
				fmt.Printf("La brûlure de %s s'éteint\n", personnage.Nom)
				continue // N'ajoute pas à effetsRestants
			}
		}

		// Décremente la durée
		effetActif.ToursRestants--

		// Garde l'effet s'il reste des tours (ou durée infinie pour certains effets)
		if effetActif.ToursRestants > 0 {
			effetsRestants = append(effetsRestants, effetActif)
		} else {
			fmt.Printf("L'effet %s de %s se dissipe\n", effet.Nom, personnage.Nom)
		}
	}

	personnage.EffetsActifs = effetsRestants

	// Applique les dégâts/soins
	if degatsRecus > 0 {
		personnage.PV -= degatsRecus
		if personnage.PV < 0 {
			personnage.PV = 0
		}
	}

	if soinsRecus > 0 {
		personnage.PV += soinsRecus
		if personnage.PV > personnage.PVMax {
			personnage.PV = personnage.PVMax
		}
	}

	return degatsRecus, soinsRecus
}

// Vérifie si un personnage est étourdi
func EstEtourdi(personnage *Personnage) bool {
	for _, effet := range personnage.EffetsActifs {
		if effet.Effet.Estourdissement {
			return true
		}
	}
	return false
}

// Supprime tous les effets d'un type donné
func SupprimerEffet(personnage *Personnage, nomEffet string) {
	var effetsRestants []EffetActif

	for _, effet := range personnage.EffetsActifs {
		if effet.Effet.Nom != nomEffet {
			effetsRestants = append(effetsRestants, effet)
		}
	}

	if len(effetsRestants) != len(personnage.EffetsActifs) {
		fmt.Printf("%s : effet %s supprimé\n", personnage.Nom, nomEffet)
		personnage.EffetsActifs = effetsRestants
	}
}

// Guérit tous les poisons (ex: potion anti-poison)
func GuerirPoison(personnage *Personnage) {
	SupprimerEffet(personnage, "Poison")
}

// Supprime tous les débuffs
func SupprimerDebuffs(personnage *Personnage) {
	var effetsRestants []EffetActif
	debuffs := []string{"Nébulation", "Défavorisation", "Brise-Armure", "Brise-Armure Magique", "Peur"}

	for _, effet := range personnage.EffetsActifs {
		estDebuff := false
		for _, debuff := range debuffs {
			if effet.Effet.Nom == debuff {
				estDebuff = true
				break
			}
		}

		if !estDebuff {
			effetsRestants = append(effetsRestants, effet)
		}
	}

	if len(effetsRestants) != len(personnage.EffetsActifs) {
		fmt.Printf("%s : tous les débuffs supprimés\n", personnage.Nom)
		personnage.EffetsActifs = effetsRestants
	}
}

// Affiche tous les effets actifs sur un personnage
func AfficherEffets(personnage *Personnage) {
	if len(personnage.EffetsActifs) == 0 {
		fmt.Printf("%s n'a aucun effet actif\n", personnage.Nom)
		return
	}

	fmt.Printf("Effets actifs sur %s:\n", personnage.Nom)
	for _, effet := range personnage.EffetsActifs {
		duree := "permanent"
		if effet.ToursRestants < 999 {
			duree = fmt.Sprintf("%d tours", effet.ToursRestants)
		}
		fmt.Printf("  - %s (%s)\n", effet.Effet.Nom, duree)
	}
}

// Créer un effet personnalisé basé sur un type d'effet
func CreerEffet(typeEffet string, puissance int) *Effet {
	switch typeEffet {
	case "Renforcement":
		return &Effet{
			Nom:             "Renforcement",
			ToursRestants:   2,
			ChanceAppliquer: 1.0,
			// La réduction de dégâts est appliquée au moment du calcul via un check dédié
		}
	case "Focalisation":
		return &Effet{
			Nom:             "Focalisation",
			ToursRestants:   2,
			ChanceAppliquer: 1.0,
			ModifPrecision:  0.15 + float64(puissance)*0.05, // +20% à +40%
		}
	case "Fortification":
		return &Effet{
			Nom:             "Fortification",
			ToursRestants:   2,
			ChanceAppliquer: 1.0,
			ModifArmure:     0.25 + float64(puissance)*0.1, // +35% à +65% armure
			ModifResistMag:  0.20 + float64(puissance)*0.1, // +30% à +60% résistance magique
		}
	case "Imprégnation":
		return &Effet{
			Nom:             "Imprégnation",
			ToursRestants:   2,
			ChanceAppliquer: 1.0,
			ModifDegats:     0.15 + float64(puissance)*0.08, // +23% à +55%
			ModifDegatsMag:  0.10 + float64(puissance)*0.08, // +18% à +50%
		}
	case "Ivresse":
		return &Effet{
			Nom:             "Ivresse",
			ToursRestants:   3,     // dure 3 tours
			ChanceAppliquer: 1.0,   // toujours appliqué
			ModifPrecision:  -0.30, // -30% précision
		}
	case "Guérison Poison":
		return &Effet{
			Nom:             "Guérison Poison",
			ToursRestants:   0,   // Effet instantané
			ChanceAppliquer: 1.0, // Toujours appliqué
			// Pas de dégâts ni modif de stats
			// -> L’effet sera interprété comme une action de nettoyage
		}
	case "Brûlure":
		return &Effet{
			Nom:             "Brûlure",
			DegatsParTour:   3 + puissance*2,              // 3-13 dégâts selon puissance
			ToursRestants:   2 + puissance/2,              // 2-4 tours
			ChanceAppliquer: 0.5 + float64(puissance)*0.1, // 50%-100%
		}

	case "Poison":
		return &Effet{
			Nom:             "Poison",
			DegatsParTour:   2 + puissance,                 // 2-7 dégâts
			ToursRestants:   999,                           // permanent jusqu'à guérison
			ChanceAppliquer: 0.4 + float64(puissance)*0.12, // 40%-100%
		}

	case "Saignement":
		return &Effet{
			Nom:             "Saignement",
			DegatsParTour:   4 + puissance*2,               // 4-14 dégâts
			ToursRestants:   2 + puissance/3,               // 2-3 tours
			ChanceAppliquer: 0.5 + float64(puissance)*0.08, // 50%-90%
		}

	case "Étourdissement":
		tours := 1
		if puissance >= 4 {
			tours = 2
		}
		return &Effet{
			Nom:             "Étourdissement",
			ToursRestants:   tours,
			ChanceAppliquer: 0.3 + float64(puissance)*0.14, // 30%-100%
			Estourdissement: true,
		}

	case "Peur":
		return &Effet{
			Nom:             "Peur",
			ToursRestants:   2 + puissance/3,
			ChanceAppliquer: 0.4 + float64(puissance)*0.1, // 40%-90%
			Peur:            true,
			ModifDegats:     -0.2 - float64(puissance)*0.08, // -20% à -60%
		}

	case "Nébulation":
		return &Effet{
			Nom:             "Nébulation",
			ToursRestants:   2 + puissance/2,                // 2-4 tours
			ChanceAppliquer: 0.6 + float64(puissance)*0.08,  // 60%-100%
			ModifPrecision:  -0.1 - float64(puissance)*0.08, // -10% à -50%
		}

	case "Défavorisation":
		return &Effet{
			Nom:             "Défavorisation",
			ToursRestants:   1 + puissance/2,                 // 1-3 tours
			ChanceAppliquer: 0.5 + float64(puissance)*0.1,    // 50%-100%
			ModifCritique:   -0.15 - float64(puissance)*0.07, // -15% à -50%
		}

	case "Brise-Armure":
		return &Effet{
			Nom:             "Brise-Armure",
			ToursRestants:   2 + puissance/2,                // 2-4 tours
			ChanceAppliquer: 0.6 + float64(puissance)*0.08,  // 60%-100%
			ModifArmure:     -0.15 - float64(puissance)*0.1, // -15% à -65%
		}

	case "Brise-Armure Magique":
		return &Effet{
			Nom:             "Brise-Armure Magique",
			ToursRestants:   2 + puissance/2,                // 2-4 tours
			ChanceAppliquer: 0.6 + float64(puissance)*0.08,  // 60%-100%
			ModifResistMag:  -0.15 - float64(puissance)*0.1, // -15% à -65%
		}

	case "Augmentation de Dégâts":
		return &Effet{
			Nom:             "Augmentation de Dégâts",
			ToursRestants:   2 + puissance/2, // 2-4 tours
			ChanceAppliquer: 1.0,
			ModifDegats:     0.15 + float64(puissance)*0.07, // +15% à +50%
		}

	case "Augmentation de Dégâts Magiques":
		return &Effet{
			Nom:             "Augmentation de Dégâts Magiques",
			ToursRestants:   2 + puissance/2, // 2-4 tours
			ChanceAppliquer: 1.0,
			ModifDegatsMag:  0.15 + float64(puissance)*0.07, // +15% à +50%
		}

	case "Affaiblissement":
		return &Effet{
			Nom:             "Affaiblissement",
			ToursRestants:   2 + puissance/2,                // 2-4 tours
			ChanceAppliquer: 0.6 + float64(puissance)*0.08,  // 60%-100%
			ModifDegats:     -0.15 - float64(puissance)*0.07, // -15% à -50% dégâts physiques
			ModifDegatsMag:  -0.10 - float64(puissance)*0.07, // -10% à -45% dégâts magiques
		}

	case "Régénération":
		return &Effet{
			Nom:             "Régénération",
			ToursRestants:   3 + puissance/2, // 3-5 tours
			SoinsParTour:    5 + puissance*3, // 5-20 PV/tour
			ChanceAppliquer: 1.0,
		}

	default:
		return nil
	}
}
