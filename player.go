package main

import (
	"fmt"
	"math/rand"
	"strings"
)

// Personnage persistant utilisé hors combat (forge, affichage, etc.)
// Par défaut on démarre avec Steeve; pourra être changé par la création de personnage.
var currentPlayer = Steeve

// Inventaire du joueur
var playerInventory = map[string]int{
	"clés":           0,
	"potions":        0,
	"pièces":         0,
	"épées":          0,
	"clés_spéciales": 0,
	"puff_9k":        0, // Nouvel item boost d'attaque
	"hache":          0, // Permet de couper les arbres
	// Nouveaux objets utilisables en combat
	"potion_mineure":       0,
	"potion_majeure":       0,
	"potion_supreme":       0,
	"potion_degats":        0,
	"bombe_incendiaire":    0,
	"bombe_givrante":       0,
	"grenade_fumigene":     0,
	"parchemin_dispersion": 0,
	"elixir_force":         0,
	"elixir_vitesse":       0,
	"elixir_critique":      0,
	"antidote":             0,
	"sida":                 0, // Objet spécial drop boss final
	"vodka_vitaly":         0, // Potion qui régénère toute la vie
	// Loots spécifiques ennemis (communs / rares)
	"dent_rat":            0,
	"dent_rat_luisante":   0,
	"gelée_visqueuse":     0,
	"coeur_de_gelée":      0,
	"capuche_brigand":     0,
	"dague_ensorcelée":    0,
	"plume_fleche":        0,
	"carquois_gravé":      0,
	"cendre_infernale":    0,
	"braise_eternelle":    0,
	"insigne_chevalier":   0,
	"lame_ancient":        0,
	"sang_berserker":      0,
	"talisman_fureur":     0,
	"essence_sombre":      0,
	"noyau_occulte":       0,
	"corne_demon":         0,
	"fragment_demoniaque": 0,
	"parchemin_arcane":    0,
	"sceau_archimage":     0,
	"embleme_champion":    0,
	"aiguille_du_destin":  0,
}

// Stats du joueur pour les bonus
var playerStats = struct {
	attackBoost        int  // Bonus d'attaque en pourcentage
	hasLegendaryWeapon bool // Arme légendaire du gambling
	enemiesKilled      int  // Nombre total d'ennemis tués
}{
	attackBoost:        0,
	hasLegendaryWeapon: false,
	enemiesKilled:      0,
}

// Affiche l'inventaire du joueur
func showInventory() {
	fmt.Println("\n🎒 === INVENTAIRE ===")
	fmt.Printf("🔑 Clés: %d\n", playerInventory["clés"])
	fmt.Printf("🗝️  Clés spéciales: %d\n", playerInventory["clés_spéciales"])
	fmt.Printf("💰 Pièces: %d\n", playerInventory["pièces"])
	fmt.Printf("🪨 Roches d'évolution: %d\n", currentPlayer.Roches)
	// Consommables
	fmt.Printf("   Potion mineure:      %d\n", playerInventory["potion_mineure"])
	fmt.Printf("   Potion majeure:      %d\n", playerInventory["potion_majeure"])
	fmt.Printf("   Potion suprême:      %d\n", playerInventory["potion_supreme"])
	fmt.Printf("   Antidote:            %d\n", playerInventory["antidote"])
	fmt.Printf("   Potion de dégâts:    %d\n", playerInventory["potion_degats"])
	fmt.Printf("   Bombe incendiaire:   %d\n", playerInventory["bombe_incendiaire"])
	fmt.Printf("   Bombe givrante:      %d\n", playerInventory["bombe_givrante"])
	fmt.Printf("   Grenade fumigène:    %d\n", playerInventory["grenade_fumigene"])
	fmt.Printf("   Parchemin dispersion:%d\n", playerInventory["parchemin_dispersion"])
	fmt.Printf("   Élixir de force:     %d\n", playerInventory["elixir_force"])
	fmt.Printf("   Élixir de vitesse:   %d\n", playerInventory["elixir_vitesse"])
	fmt.Printf("   Élixir de précision: %d\n", playerInventory["elixir_critique"])
	fmt.Printf("💊 Puff 9K: %d\n", playerInventory["puff_9k"])
	fmt.Printf("🪓 Hache: %d\n", playerInventory["hache"])
	fmt.Printf("🍶 Vodka de Vitaly: %d\n", playerInventory["vodka_vitaly"])

	// Section loots spécifiques
	fmt.Println("-- Loots ennemis --")
	printIf := func(key, label string) {
		if playerInventory[key] > 0 {
			fmt.Printf("   %-22s %d\n", label+":", playerInventory[key])
		}
	}
	printIf("dent_rat", "Dent de Rat")
	printIf("dent_rat_luisante", "Dent de Rat Luisante")
	printIf("gelée_visqueuse", "Gelée Visqueuse")
	printIf("coeur_de_gelée", "Cœur de Gelée")
	printIf("capuche_brigand", "Capuche de Brigand")
	printIf("dague_ensorcelée", "Dague Ensorcelée")
	printIf("plume_fleche", "Plume de Flèche")
	printIf("carquois_gravé", "Carquois Gravé")
	printIf("cendre_infernale", "Cendre Infernale")
	printIf("braise_eternelle", "Braise Éternelle")
	printIf("insigne_chevalier", "Insigne de Chevalier")
	printIf("lame_ancient", "Lame Ancienne")
	printIf("sang_berserker", "Sang de Berserker")
	printIf("talisman_fureur", "Talisman de Fureur")
	printIf("essence_sombre", "Essence Sombre")
	printIf("noyau_occulte", "Noyau Occulte")
	printIf("corne_demon", "Corne de Démon")
	printIf("fragment_demoniaque", "Fragment Démoniaque")
	printIf("parchemin_arcane", "Parchemin Arcane")
	printIf("sceau_archimage", "Sceau d'Archimage")
	printIf("embleme_champion", "Emblème de Champion")
	printIf("aiguille_du_destin", "Aiguille du Destin")
	printIf("fragment_mentor", "Fragment de Mentor")
	printIf("souffle_mentor", "Souffle du Mentor")
	// Affiche les artefacts équipés
	artefacts := []string{}
	for _, a := range currentPlayer.ArtefactsEquipes {
		if a != nil {
			artefacts = append(artefacts, a.Nom)
		}
	}
	if len(artefacts) > 0 {
		fmt.Printf("🧿 Artefacts équipés: %s\n", strings.Join(artefacts, ", "))
	} else {
		fmt.Println("🧿 Artefacts équipés: Aucun")
	}
	if playerStats.hasLegendaryWeapon {
		fmt.Println("🌟 Excalibur Légendaire équipée!")
	}
	fmt.Println("===================")
}

// Ajoute un objet à l'inventaire
func addToInventory(item string, amount int) {
	playerInventory[item] += amount
	fmt.Printf("✨ Vous avez reçu %d %s !\n", amount, item)
}

// Calcule les dégâts d'attaque avec les bonus
func calculateAttackDamage() int {
	baseDamage := 20 + rand.Intn(15) // 20-34 dégâts de base

	// Bonus du Puff 9K : +15% par Puff 9K utilisé
	puffBonus := float64(playerStats.attackBoost) / 100.0

	// Bonus arme légendaire : +50%
	legendaryBonus := 0.0
	if playerStats.hasLegendaryWeapon {
		legendaryBonus = 0.5
	}

	totalDamage := float64(baseDamage) * (1.0 + puffBonus + legendaryBonus)

	return int(totalDamage)
}
