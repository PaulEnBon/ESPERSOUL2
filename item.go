package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Interface de commerce du marchand
func showItemMerchantInterface() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n🛒 === MARCHAND ===")
		fmt.Printf("💰 Vos pièces: %d\n", playerInventory["pièces"])
		fmt.Printf("🧪 Vos potions: %d\n", playerInventory["potions"])
		fmt.Printf("🔑 Vos clés: %d\n", playerInventory["clés"])
		fmt.Printf("🗝️  Vos clés rouillées: %d\n", playerInventory["clés_spéciales"])
		fmt.Printf("💊 Vos Puff 9K: %d\n", playerInventory["puff_9k"])
		fmt.Println("\n📜 Articles disponibles:")
		fmt.Println("1. Potion de soin - 5 pièces")
		fmt.Println("2. Clé magique - 10 pièces")
		fmt.Println("3. Vielle clé rouillée - 50 pièces")
		fmt.Println("4. 💊 Puff 9K - 20 pièces (+15% attaque, -5HP)")
		fmt.Println("5. Quitter le magasin")
		fmt.Print("Choisissez un article (1-5): ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			if playerInventory["pièces"] >= 5 {
				playerInventory["pièces"] -= 5
				playerInventory["potions"]++
				fmt.Println("✨ Vous avez acheté une potion de soin!")
			} else {
				fmt.Println("❌ Vous n'avez pas assez de pièces!")
			}

		case "2":
			if playerInventory["pièces"] >= 10 {
				playerInventory["pièces"] -= 10
				playerInventory["clés"]++
				fmt.Println("✨ Vous avez acheté une clé magique!")
			} else {
				fmt.Println("❌ Vous n'avez pas assez de pièces!")
			}

		case "3":
			if playerInventory["pièces"] >= 50 {
				playerInventory["pièces"] -= 50
				playerInventory["clés_spéciales"]++
				fmt.Println("🌟 Vous avez acheté une VIELLE CLÉ ROUILLÉE!")
			} else {
				fmt.Println("❌ Vous n'avez pas assez de pièces! (50 pièces nécessaires)")
			}

		case "4":
			if playerInventory["pièces"] >= 20 {
				playerInventory["pièces"] -= 20
				playerInventory["puff_9k"]++
				fmt.Println("💊 Vous avez acheté un Puff 9K!")
				fmt.Println("⚡ +15% de dégâts d'attaque mais attention : -5HP!")
			} else {
				fmt.Println("❌ Vous n'avez pas assez de pièces! (20 pièces nécessaires)")
			}

		case "5":
			fmt.Println("👋 Merci de votre visite!")
			return

		default:
			fmt.Println("❌ Choix invalide!")
		}

		fmt.Print("Appuyez sur Entrée pour continuer...")
		reader.ReadString('\n')
	}
}

// Interface du forgeron
func showItemForgeInterface() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n🔨 === FORGERON ===")
		fmt.Printf("💰 Vos pièces: %d\n", playerInventory["pièces"])
		fmt.Printf("⚔️  Vos épées: %d\n", playerInventory["épées"])
		fmt.Println("\n🛠️  Services disponibles:")
		fmt.Println("1. Forger une épée - 15 pièces")
		fmt.Println("2. Quitter la forge")
		fmt.Print("Choisissez une option (1-2): ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			if playerInventory["pièces"] >= 15 {
				playerInventory["pièces"] -= 15
				playerInventory["épées"]++
				fmt.Println("🔨 *Clang clang clang*")
				fmt.Println("✨ Le forgeron vous forge une magnifique épée !")
				fmt.Println("⚔️  Vous avez reçu une épée forgée !")
			} else {
				fmt.Println("❌ Vous n'avez pas assez de pièces pour une épée!")
			}

		case "2":
			fmt.Println("👋 Revenez quand vous voulez!")
			return

		default:
			fmt.Println("❌ Choix invalide!")
		}

		fmt.Print("Appuyez sur Entrée pour continuer...")
		reader.ReadString('\n')
	}
}

// Interface de gambling (pour item.go)
func showItemGamblingInterface() {
	reader := bufio.NewReader(os.Stdin)
	rand.Seed(time.Now().UnixNano())

	for {
		fmt.Println("\n🎰 === CASINO SOUTERRAIN ===")
		fmt.Printf("💰 Vos pièces: %d\n", playerInventory["pièces"])
		fmt.Printf("⚔️  Vos épées: %d\n", playerInventory["épées"])
		fmt.Printf("💊 Vos Puff 9K: %d\n", playerInventory["puff_9k"])
		if playerStats.hasLegendaryWeapon {
			fmt.Println("🌟 Arme légendaire équipée !")
		}

		fmt.Println("\n📦 Caisses disponibles:")
		fmt.Println("1. Caisse Bronze - 5 pièces (Chances mystérieuses...)")
		fmt.Println("2. Caisse Argent - 25 pièces (Bonnes chances)")
		fmt.Println("3. Caisse Or - 75 pièces (Très bonnes chances)")
		fmt.Println("4. Caisse Legendary - 1000 pièces (Garanti légendaire !)")
		fmt.Println("5. Quitter le casino")
		fmt.Print("Choisissez une caisse (1-5): ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1": // Caisse Bronze - 5 pièces
			if playerInventory["pièces"] >= 5 {
				playerInventory["pièces"] -= 5
				fmt.Println("📦 *Ouverture de la caisse Bronze...*")
				time.Sleep(1 * time.Second)

				roll := rand.Intn(100)
				if roll < 2 { // 2% chance d'arme légendaire
					playerStats.hasLegendaryWeapon = true
					fmt.Println("🌟 JACKPOT ! Vous obtenez l'EXCALIBUR LÉGENDAIRE !")
					fmt.Println("⚡ +50% de dégâts d'attaque permanents !")
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

		case "2": // Caisse Argent - 25 pièces
			if playerInventory["pièces"] >= 25 {
				playerInventory["pièces"] -= 25
				fmt.Println("📦 *Ouverture de la caisse Argent...*")
				time.Sleep(1 * time.Second)

				roll := rand.Intn(100)
				if roll < 5 { // 5% chance d'arme légendaire
					playerStats.hasLegendaryWeapon = true
					fmt.Println("🌟 INCROYABLE ! Vous obtenez l'EXCALIBUR LÉGENDAIRE !")
					fmt.Println("⚡ +50% de dégâts d'attaque permanents !")
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

		case "3": // Caisse Or - 75 pièces
			if playerInventory["pièces"] >= 75 {
				playerInventory["pièces"] -= 75
				fmt.Println("📦 *Ouverture de la caisse Or...*")
				time.Sleep(1 * time.Second)

				roll := rand.Intn(100)
				if roll < 15 { // 15% chance d'arme légendaire
					playerStats.hasLegendaryWeapon = true
					fmt.Println("🌟 FANTASTIQUE ! Vous obtenez l'EXCALIBUR LÉGENDAIRE !")
					fmt.Println("⚡ +50% de dégâts d'attaque permanents !")
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

		case "4": // Caisse Legendary - 150 pièces (100% légendaire)
			if playerInventory["pièces"] >= 1000 {
				playerInventory["pièces"] -= 1000
				fmt.Println("📦 *Ouverture de la caisse LEGENDARY...*")
				time.Sleep(2 * time.Second)

				playerStats.hasLegendaryWeapon = true
				fmt.Println("🌟 LÉGENDAIRE GARANTI ! Vous obtenez l'EXCALIBUR LÉGENDAIRE !")
				fmt.Println("⚡ +50% de dégâts d'attaque permanents !")

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

		case "5":
			fmt.Println("🎰 À bientôt au casino !")
			return

		default:
			fmt.Println("❌ Choix invalide !")
		}

		fmt.Print("Appuyez sur Entrée pour continuer...")
		reader.ReadString('\n')
	}
}
