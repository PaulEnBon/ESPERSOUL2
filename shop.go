package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Interface de commerce du marchand
func showMerchantInterface() {
	// Helper: read a single key (last of any burst), return as string "1".."9" or letter lowercased
	readKey := func() rune {
		if globalKeyEvents == nil {
			return 0
		}
		e := <-globalKeyEvents
		draining := true
		for draining {
			select {
			case next := <-globalKeyEvents:
				e = next
			default:
				draining = false
			}
		}
		r := e.Rune
		if r >= 'A' && r <= 'Z' {
			r = r + 32
		}
		return r
	}

	for {
		fmt.Println("\n🛒 === MARCHAND ===")
		fmt.Printf("💰 Vos pièces: %d\n", playerInventory["pièces"])
		fmt.Printf("🧪 Vos potions: %d\n", playerInventory["potions"])
		fmt.Printf("🔑 Vos clés: %d\n", playerInventory["clés"])
		fmt.Printf("🗝️  Vos clés rouillées: %d\n", playerInventory["clés_spéciales"])
		fmt.Printf("💊 Vos Puff 9K: %d\n", playerInventory["puff_9k"])
		fmt.Println("\n📜 Articles disponibles:")
		fmt.Println("1. Clé magique - 10 pièces")
		fmt.Println("2. Vielle clé rouillée - 50 pièces")
		fmt.Println("3. 💊 Puff 9K - 20 pièces (+15% attaque, -5HP)")
		// Nouveaux objets
		fmt.Println("4. Antidote - 6 pièces")
		fmt.Println("5. Potion mineure - 3 pièces")
		fmt.Println("6. Potion majeure - 10 pièces")
		fmt.Println("7. Potion suprême - 25 pièces")
		fmt.Println("8. Potion de dégâts - 6 pièces")
		fmt.Println("a. Hache - 30 pièces (permet de couper les arbres)")
		fmt.Println("b. Vodka de Vitaly - 50 pièces (régénère toute la vie)")
		fmt.Println("c. Bombe incendiaire - 12 pièces")
		fmt.Println("d. Bombe givrante - 14 pièces")
		fmt.Println("e. Grenade fumigène - 8 pièces")
		fmt.Println("f. Parchemin de dispersion - 10 pièces")
		fmt.Println("g. Élixir de force - 12 pièces")
		fmt.Println("h. Élixir de vitesse - 10 pièces")
		fmt.Println("i. Élixir de précision - 15 pièces")
		fmt.Println("q. Quitter le magasin")
		fmt.Print("Choisissez un article (1-9, a-g, i, q): ")
		key := readKey()

		switch key {
		case 'b': // Vodka de Vitaly - 50 pièces
			if playerInventory["pièces"] >= 50 {
				playerInventory["pièces"] -= 50
				playerInventory["vodka_vitaly"]++
				fmt.Println("🍶 Vous avez acheté une Vodka de Vitaly ! Toute votre vie sera régénérée lors de son utilisation.")
			} else {
				fmt.Println("❌ Vous n'avez pas assez de pièces! (50 pièces nécessaires)")
			}

		case '1':
			if playerInventory["pièces"] >= 10 {
				playerInventory["pièces"] -= 10
				playerInventory["clés"]++
				fmt.Println("✨ Vous avez acheté une clé magique!")
			} else {
				fmt.Println("❌ Vous n'avez pas assez de pièces!")
			}

		case '2':
			if playerInventory["pièces"] >= 50 {
				playerInventory["pièces"] -= 50
				playerInventory["clés_spéciales"]++
				fmt.Println("🌟 Vous avez acheté une VIELLE CLÉ ROUILLÉE!")
			} else {
				fmt.Println("❌ Vous n'avez pas assez de pièces! (50 pièces nécessaires)")
			}

		case '3':
			if playerInventory["pièces"] >= 20 {
				playerInventory["pièces"] -= 20
				playerInventory["puff_9k"]++
				fmt.Println("💊 Vous avez acheté un Puff 9K!")
				fmt.Println("⚡ +15% de dégâts d'attaque mais attention : -5HP!")
			} else {
				fmt.Println("❌ Vous n'avez pas assez de pièces! (20 pièces nécessaires)")
			}

		case '4': // Antidote - 6 pièces
			if playerInventory["pièces"] >= 6 {
				playerInventory["pièces"] -= 6
				playerInventory["antidote"]++
				fmt.Println("✨ Vous avez acheté un antidote !")
			} else {
				fmt.Println("❌ Vous n'avez pas assez de pièces! (6 pièces nécessaires)")
			}

		case '5': // Potion mineure - 3 pièces
			if playerInventory["pièces"] >= 3 {
				playerInventory["pièces"] -= 3
				playerInventory["potion_mineure"]++
				fmt.Println("✨ Vous avez acheté une potion mineure !")
			} else {
				fmt.Println("❌ Vous n'avez pas assez de pièces! (3 pièces nécessaires)")
			}

		case '6': // Potion majeure - 10 pièces
			if playerInventory["pièces"] >= 10 {
				playerInventory["pièces"] -= 10
				playerInventory["potion_majeure"]++
				fmt.Println("✨ Vous avez acheté une potion majeure !")
			} else {
				fmt.Println("❌ Vous n'avez pas assez de pièces! (10 pièces nécessaires)")
			}

		case '7': // Potion suprême - 25 pièces
			if playerInventory["pièces"] >= 25 {
				playerInventory["pièces"] -= 25
				playerInventory["potion_supreme"]++
				fmt.Println("✨ Vous avez acheté une potion suprême !")
			} else {
				fmt.Println("❌ Vous n'avez pas assez de pièces! (25 pièces nécessaires)")
			}

		case '8': // Potion de dégâts - 6 pièces
			if playerInventory["pièces"] >= 6 {
				playerInventory["pièces"] -= 6
				playerInventory["potion_degats"]++
				fmt.Println("✨ Vous avez acheté une potion de dégâts !")
			} else {
				fmt.Println("❌ Vous n'avez pas assez de pièces! (6 pièces nécessaires)")
			}

		case 'a': // Hache - 30 pièces
			if playerInventory["pièces"] >= 30 {
				playerInventory["pièces"] -= 30
				playerInventory["hache"]++
				fmt.Println("⛏️ Vous avez acheté une HACHE ! Vous pouvez maintenant couper les arbres (🌳).")
				fmt.Println("Approchez-vous d'un arbre et marchez dessus pour le couper.")
			} else {
				fmt.Println("❌ Vous n'avez pas assez de pièces! (30 pièces nécessaires)")
			}

		case 'c': // Bombe incendiaire - 12 pièces
			if playerInventory["pièces"] >= 12 {
				playerInventory["pièces"] -= 12
				playerInventory["bombe_incendiaire"]++
				fmt.Println("✨ Vous avez acheté une bombe incendiaire !")
			} else {
				fmt.Println("❌ Vous n'avez pas assez de pièces! (12 pièces nécessaires)")
			}

		case 'd': // Bombe givrante - 14 pièces
			if playerInventory["pièces"] >= 14 {
				playerInventory["pièces"] -= 14
				playerInventory["bombe_givrante"]++
				fmt.Println("✨ Vous avez acheté une bombe givrante !")
			} else {
				fmt.Println("❌ Vous n'avez pas assez de pièces! (14 pièces nécessaires)")
			}

		case 'e': // Grenade fumigène - 8 pièces
			if playerInventory["pièces"] >= 8 {
				playerInventory["pièces"] -= 8
				playerInventory["grenade_fumigene"]++
				fmt.Println("✨ Vous avez acheté une grenade fumigène !")
			} else {
				fmt.Println("❌ Vous n'avez pas assez de pièces! (8 pièces nécessaires)")
			}

		case 'f': // Parchemin de dispersion - 10 pièces
			if playerInventory["pièces"] >= 10 {
				playerInventory["pièces"] -= 10
				playerInventory["parchemin_dispersion"]++
				fmt.Println("✨ Vous avez acheté un parchemin de dispersion !")
			} else {
				fmt.Println("❌ Vous n'avez pas assez de pièces! (10 pièces nécessaires)")
			}

		case 'g': // Élixir de force - 12 pièces
			if playerInventory["pièces"] >= 12 {
				playerInventory["pièces"] -= 12
				playerInventory["elixir_force"]++
				fmt.Println("✨ Vous avez acheté un élixir de force !")
			} else {
				fmt.Println("❌ Vous n'avez pas assez de pièces! (12 pièces nécessaires)")
			}

		case 'h': // Élixir de vitesse - 10 pièces
			if playerInventory["pièces"] >= 10 {
				playerInventory["pièces"] -= 10
				playerInventory["elixir_vitesse"]++
				fmt.Println("✨ Vous avez acheté un élixir de vitesse !")
			} else {
				fmt.Println("❌ Vous n'avez pas assez de pièces! (10 pièces nécessaires)")
			}

		case 'i': // Élixir de précision - 15 pièces
			if playerInventory["pièces"] >= 15 {
				playerInventory["pièces"] -= 15
				playerInventory["elixir_critique"]++
				fmt.Println("✨ Vous avez acheté un élixir de précision !")
			} else {
				fmt.Println("❌ Vous n'avez pas assez de pièces! (15 pièces nécessaires)")
			}

		case 'q':
			fmt.Println("👋 Merci de votre visite!")
			return

		default:
			fmt.Println("❌ Choix invalide!")
		}
		fmt.Print("Appuyez sur une touche pour continuer...")
		_ = readKey()
	}
}

// Interface du forgeron
func showForgeInterface() {
	readKey := func() rune {
		if globalKeyEvents == nil {
			return 0
		}
		e := <-globalKeyEvents
		draining := true
		for draining {
			select {
			case next := <-globalKeyEvents:
				e = next
			default:
				draining = false
			}
		}
		r := e.Rune
		if r >= 'A' && r <= 'Z' {
			r = r + 32
		}
		return r
	}

	for {
		fmt.Println("\n🔨 === FORGERON ===")
		fmt.Printf("💰 Pièces: %d  🪨 Roches: %d\n", playerInventory["pièces"], currentPlayer.Roches)

		// Calcul des coûts d'amélioration actuels
		coutArme := CoutAmelioration(currentPlayer.NiveauArme)
		coutArmure := CoutAmelioration(currentPlayer.NiveauArmure)

		fmt.Println("\n🛠️  Services disponibles:")
		fmt.Printf("1. Améliorer l'arme (Niv %d → %d) - %d roches\n", currentPlayer.NiveauArme, currentPlayer.NiveauArme+1, coutArme)
		fmt.Printf("2. Améliorer l'armure (Niv %d → %d) - %d roches\n", currentPlayer.NiveauArmure, currentPlayer.NiveauArmure+1, coutArmure)
		fmt.Println("3. Afficher les stats du joueur")
		fmt.Println("4. Quitter la forge")
		fmt.Print("Choisissez une option (1-4): ")
		key := readKey()

		switch key {
		case '1': // Améliorer l'arme
			if err := AmeliorerArme(&currentPlayer, len(currentPlayer.ArmesDisponibles)); err != nil {
				fmt.Printf("❌ %v\n", err)
			} else {
				// Ne pas écraser une arme spéciale déjà équipée (ex: Dragon Lore)
				if currentPlayer.ArmeEquipee.Nom == "" && currentPlayer.NiveauArme < len(currentPlayer.ArmesDisponibles) {
					currentPlayer.ArmeEquipee = currentPlayer.ArmesDisponibles[currentPlayer.NiveauArme]
				}
				fmt.Printf("✅ Arme améliorée → %s (niv %d)\n", currentPlayer.ArmeEquipee.Nom, currentPlayer.NiveauArme)
			}

		case '2': // Améliorer l'armure
			if err := AmeliorerArmure(&currentPlayer, len(currentPlayer.ArmuresDisponibles)); err != nil {
				fmt.Printf("❌ %v\n", err)
			} else {
				if currentPlayer.NiveauArmure < len(currentPlayer.ArmuresDisponibles) {
					currentPlayer.ArmureEquipee = currentPlayer.ArmuresDisponibles[currentPlayer.NiveauArmure]
				}
				fmt.Printf("✅ Armure améliorée → %s (niv %d)\n", currentPlayer.ArmureEquipee.Nom, currentPlayer.NiveauArmure)
			}

		case '3': // Afficher stats
			AfficherStats(&currentPlayer)

		case '4':
			fmt.Println("👋 Revenez quand vous voulez!")
			return

		default:
			fmt.Println("❌ Choix invalide!")
		}
		fmt.Print("Appuyez sur une touche pour continuer...")
		_ = readKey()
	}
}

// Interface de gambling
func showGamblingInterface() {
	// Note: Go 1.20+ auto-seeds the RNG

	readKey := func() rune {
		if globalKeyEvents == nil {
			return 0
		}
		e := <-globalKeyEvents
		draining := true
		for draining {
			select {
			case next := <-globalKeyEvents:
				e = next
			default:
				draining = false
			}
		}
		r := e.Rune
		if r >= 'A' && r <= 'Z' {
			r = r + 32
		}
		return r
	}

	for {
		fmt.Println("\n🎰 === CASINO SOUTERRAIN ===")
		fmt.Printf("💰 Vos pièces: %d\n", playerInventory["pièces"])
		// Retiré: affichage des épées
		fmt.Printf("💊 Vos Puff 9K: %d\n", playerInventory["puff_9k"])
		if currentPlayer.ArmeEquipee.Nom == dragonLore.Nom {
			fmt.Println("🌟 AWP Dragon Lore équipée !")
		}

		fmt.Println("\n📦 Caisses disponibles:")
		fmt.Println("1. Caisse Bronze - 5 pièces (Chances mystérieuses...)")
		fmt.Println("2. Caisse Argent - 25 pièces (Bonnes chances)")
		fmt.Println("3. Caisse Or - 75 pièces (Très bonnes chances)")
		fmt.Println("4. Caisse Legendary - 150 pièces (Garanti Dragon Lore !)")
		fmt.Println("5. Quitter le casino")
		fmt.Print("Choisissez une caisse (1-5): ")
		key := readKey()

		switch key {
		case '1': // Caisse Bronze - 5 pièces
			if playerInventory["pièces"] >= 5 {
				playerInventory["pièces"] -= 5
				fmt.Println("📦 *Ouverture de la caisse Bronze...*")
				time.Sleep(1 * time.Second)

				roll := rand.Intn(100)
				if roll < 2 { // 2% chance Dragon Lore
					_ = EquiperArme(&currentPlayer, dragonLore)
					fmt.Println("🐉 JACKPOT ! Vous obtenez l'AWP DRAGON LORE !")
					fmt.Println("🔥 Arme ultime équipée.")
				} else if roll < 10 { // 8% chance d'épées
					amount := 1 + rand.Intn(2) // 1-2 épées
					addToInventory("épées", amount)
				} else if roll < 30 { // 20% chance de puff 9k
					addToInventory("puff_9k", 1)
				} else if roll < 60 { // 30% chance de potions
					amount := 1 + rand.Intn(3) // 1-3 potions
					addToInventory("potions", amount)
				} else { // 40% chance de pièces
					amount := 2 + rand.Intn(8) // 2-9 pièces
					addToInventory("pièces", amount)
				}
			} else {
				fmt.Println("❌ Vous n'avez pas assez de pièces !")
			}

		case '2': // Caisse Argent - 25 pièces
			if playerInventory["pièces"] >= 25 {
				playerInventory["pièces"] -= 25
				fmt.Println("📦 *Ouverture de la caisse Argent...*")
				time.Sleep(1 * time.Second)

				roll := rand.Intn(100)
				if roll < 5 { // 5% chance Dragon Lore
					_ = EquiperArme(&currentPlayer, dragonLore)
					fmt.Println("🐉 INCROYABLE ! Vous obtenez l'AWP DRAGON LORE !")
					fmt.Println("🔥 Arme ultime équipée.")
				} else if roll < 25 { // 20% chance d'épées multiples
					amount := 2 + rand.Intn(3) // 2-4 épées
					addToInventory("épées", amount)
				} else if roll < 50 { // 25% chance de puff 9k multiple
					amount := 1 + rand.Intn(2) // 1-2 puff 9k
					addToInventory("puff_9k", amount)
				} else if roll < 80 { // 30% chance de potions
					amount := 3 + rand.Intn(3) // 3-5 potions
					addToInventory("potions", amount)
				} else { // 20% chance de pièces
					amount := 15 + rand.Intn(20) // 15-34 pièces
					addToInventory("pièces", amount)
				}
			} else {
				fmt.Println("❌ Vous n'avez pas assez de pièces !")
			}

		case '3': // Caisse Or - 75 pièces
			if playerInventory["pièces"] >= 75 {
				playerInventory["pièces"] -= 75
				fmt.Println("📦 *Ouverture de la caisse Or...*")
				time.Sleep(1 * time.Second)

				roll := rand.Intn(100)
				if roll < 15 { // 15% chance Dragon Lore
					_ = EquiperArme(&currentPlayer, dragonLore)
					fmt.Println("🐉 FANTASTIQUE ! Vous obtenez l'AWP DRAGON LORE !")
					fmt.Println("🔥 Arme ultime équipée.")
				} else if roll < 40 { // 25% chance d'épées premium
					amount := 3 + rand.Intn(3) // 3-5 épées
					addToInventory("épées", amount)
				} else if roll < 65 { // 25% chance de puff 9k premium
					amount := 2 + rand.Intn(2) // 2-3 puff 9k
					addToInventory("puff_9k", amount)
				} else if roll < 90 { // 25% chance de potions premium
					amount := 5 + rand.Intn(5) // 5-9 potions
					addToInventory("potions", amount)
				} else { // 10% chance de pièces premium
					amount := 50 + rand.Intn(50) // 50-99 pièces
					addToInventory("pièces", amount)
				}
			} else {
				fmt.Println("❌ Vous n'avez pas assez de pièces !")
			}

		case '4': // Caisse Legendary - 150 pièces (100% légendaire)
			if playerInventory["pièces"] >= 150 {
				playerInventory["pièces"] -= 150
				fmt.Println("📦 *Ouverture de la caisse LEGENDARY...*")
				time.Sleep(2 * time.Second)

				_ = EquiperArme(&currentPlayer, dragonLore)
				fmt.Println("🐉 DRAGON LORE GARANTI ! Vous obtenez l'AWP DRAGON LORE !")
				fmt.Println("🔥 Arme ultime équipée.")

				// Bonus supplémentaire
				bonusRoll := rand.Intn(3)
				switch bonusRoll {
				case 0:
					addToInventory("épées", 5)
					fmt.Println("🎁 Bonus : 5 épées supplémentaires !")
				case 1:
					addToInventory("puff_9k", 3)
					fmt.Println("🎁 Bonus : 3 Puff 9K supplémentaires !")
				case 2:
					addToInventory("potions", 10)
					fmt.Println("🎁 Bonus : 10 potions supplémentaires !")
				}
			} else {
				fmt.Println("❌ Vous n'avez pas assez de pièces !")
			}

		case '5':
			fmt.Println("🎰 À bientôt au casino !")
			return

		default:
			fmt.Println("❌ Choix invalide !")
		}
		fmt.Print("Appuyez sur une touche pour continuer...")
		_ = readKey()
	}
}
