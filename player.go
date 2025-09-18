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
	if currentPlayer.ArmeEquipee.Nom == dragonLore.Nom {
		fmt.Println("🌟 AWP Dragon Lore équipée!")
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

	// Bonus arme légendaire (ancien) retiré: Dragon Lore applique ses propres dégâts via ArmeEquipee
	legendaryBonus := 0.0

	totalDamage := float64(baseDamage) * (1.0 + puffBonus + legendaryBonus)

	return int(totalDamage)
}
