package main

import (
	"fmt"
	"math/rand"
)

// Personnage persistant utilisé hors combat (forge, affichage, etc.)
// Par défaut on démarre avec Steeve; pourra être changé par la création de personnage.
var currentPlayer = Pyromane

// Inventaire du joueur
var playerInventory = map[string]int{
	"clés":           0,
	"potions":        0,
	"pièces":         0,
	"épées":          0,
	"clés_spéciales": 0,
	"puff_9k":        0, // Nouvel item boost d'attaque
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
