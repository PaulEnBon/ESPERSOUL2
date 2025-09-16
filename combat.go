package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// ---- Récompenses configurables ----
// Peut être déplacé dans un fichier de config plus tard.
const (
	baseMinCoins         = 5    // ancien 3
	baseMaxCoins         = 9    // génère range baseMin..baseMax inclus
	swordPerPieceBonus   = 0.10 // +10% des pièces de base par épée possédée
	legendaryWeaponBonus = 0.50 // +50% si arme légendaire
	puffAttackBonusShare = 0.20 // 20% de l'attackBoost converti en bonus or (ex: 30% atk -> +6% or)
	jackpotChancePercent = 6    // % de chance de jackpot
	jackpotMultiplier    = 4.0  // x4 sur le total final
)

// calcule le loot de pièces enrichi
func computeCoinLoot() (coins int, jackpot bool, breakdown string) {
	// Base aléatoire
	base := baseMinCoins + rand.Intn(baseMaxCoins-baseMinCoins+1)

	// Bonus épées
	swordBonus := int(float64(base) * swordPerPieceBonus * float64(playerInventory["épées"]))

	// Bonus arme légendaire
	legendaryBonus := 0
	if playerStats.hasLegendaryWeapon {
		legendaryBonus = int(float64(base) * legendaryWeaponBonus)
	}

	// Bonus Puff converti (attackBoost est un pourcentage cumulatif)
	puffBonus := int(float64(base) * (float64(playerStats.attackBoost) / 100.0) * puffAttackBonusShare)

	total := base + swordBonus + legendaryBonus + puffBonus

	// Jackpot ?
	if rand.Intn(100) < jackpotChancePercent {
		total = int(float64(total) * jackpotMultiplier)
		return total, true, fmt.Sprintf("base=%d +épées=%d +legend=%d +puff=%d xJackpot(%.1fx)", base, swordBonus, legendaryBonus, puffBonus, jackpotMultiplier)
	}

	return total, false, fmt.Sprintf("base=%d +épées=%d +legend=%d +puff=%d", base, swordBonus, legendaryBonus, puffBonus)
}

// Système de combat amélioré avec les nouveaux items
func combat(currentMap string, isSuper bool) interface{} {
	reader := bufio.NewReader(os.Stdin)
	rand.Seed(time.Now().UnixNano())

	playerHP := 100
	enemyHP := 80
	enemyAttackBase := 15
	if isSuper {
		enemyHP *= 2         // 2x vie
		enemyAttackBase *= 2 // 2x attaque
	}

	fmt.Println("\n🗡️  COMBAT ENGAGÉ ! 🗡️")
	if isSuper {
		fmt.Println("Vous affrontez un ENNEMI SURPUISSANT !")
	} else {
		fmt.Println("Vous affrontez une créature maudite !")
	}

	for playerHP > 0 && enemyHP > 0 {
		fmt.Printf("\n💚 Vos PV: %d | 💀 PV Ennemi: %d\n", playerHP, enemyHP)
		fmt.Printf("⚔️  Épées équipées: %d (+%d dégâts)\n", playerInventory["épées"], playerInventory["épées"]*3)
		if playerStats.hasLegendaryWeapon {
			fmt.Println("🌟 Excalibur équipée (+50% dégâts)")
		}
		fmt.Println("Actions: [A]ttaquer, [D]éfendre, [P]otion, [U]ser Puff 9K, [F]uir")
		fmt.Print("Choisissez une action: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))

		switch input {
		case "a":
			damage := calculateAttackDamage()
			enemyHP -= damage
			fmt.Printf("⚔️  Vous infligez %d dégâts !\n", damage)

		case "d":
			fmt.Println("🛡️  Vous vous défendez !")
			damage := 5 + rand.Intn(10)
			enemyDamage := enemyAttackBase + rand.Intn(10) - damage
			if enemyDamage < 0 {
				enemyDamage = 0
			}
			playerHP -= enemyDamage
			fmt.Printf("💥 L'ennemi vous inflige %d dégâts (réduits) !\n", enemyDamage)
			continue

		case "p":
			if playerInventory["potions"] > 0 {
				heal := 70 // fixed heal amount per potion
				playerHP += heal
				if playerHP > 100 {
					playerHP = 100
				}
				playerInventory["potions"]--
				fmt.Printf("🧪 Vous vous soignez de %d PV ! (PV actuels: %d)\n", heal, playerHP)
			} else {
				fmt.Println("❌ Vous n'avez pas de potions !")
				continue
			}

		case "u":
			if playerInventory["puff_9k"] > 0 {
				playerInventory["puff_9k"]--
				playerStats.attackBoost += 15 // +15% d'attaque
				playerHP -= 5                 // -5 HP
				fmt.Println("💊 Vous utilisez un Puff 9K !")
				fmt.Println("⚡ +15% de dégâts d'attaque pour ce combat !")
				fmt.Printf("💔 Vous perdez 5 HP. PV actuels: %d\n", playerHP)

				if playerHP <= 0 {
					fmt.Println("💀 Le Puff 9K vous a tué ! Attention à la surdose...")
					return false
				}
				continue
			} else {
				fmt.Println("❌ Vous n'avez pas de Puff 9K !")
				continue
			}

		case "f":
			fmt.Println("💨 Vous fuyez le combat !")
			// Reset les bonus temporaires
			playerStats.attackBoost = 0
			return false

		default:
			fmt.Println("Action invalide !")
			continue
		}

		if enemyHP <= 0 {
			fmt.Println("\n🎉 VICTOIRE ! Vous avez vaincu la créature !")

			// Incrémente le compteur d'ennemis tués
			playerStats.enemiesKilled++

			coins, jackpot, details := computeCoinLoot()
			addToInventory("pièces", coins)
			if jackpot {
				fmt.Printf("💎 JACKPOT ! Vous obtenez %d pièces (%s) !\n", coins, details)
			} else {
				fmt.Printf("✨ Vous avez reçu %d pièces (%s).\n", coins, details)
			}

			// Reset les bonus temporaires (après calcul qui pouvait s'en servir)
			playerStats.attackBoost = 0

			// Par défaut, les ennemis ne se transforment jamais en PNJ ici.
			// La seule exception (le seul mob de salle1) est gérée côté loop
			// après le combat, à partir de ses coordonnées exactes.
			fmt.Println("💨 La créature disparaît complètement dans un nuage de fumée...")
			return "disappear"
		}

		// Attaque de l'ennemi
		enemyDamage := enemyAttackBase + rand.Intn(10)
		playerHP -= enemyDamage
		fmt.Printf("💥 L'ennemi vous inflige %d dégâts !\n", enemyDamage)

		if playerHP <= 0 {
			fmt.Println("\n💀 DÉFAITE ! Vous avez été vaincu...")
			fmt.Println("🔄 Vous retournez au début de la salle.")
			// Reset les bonus temporaires
			playerStats.attackBoost = 0
			return false
		}
	}

	return false
}
