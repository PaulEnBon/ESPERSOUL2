package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/eiannone/keyboard"
)

// showInventoryMenu affiche un écran d'inventaire détaillé. On reste dans cette vue
// tant que l'utilisateur ne presse pas I, ESC ou Entrée.
func showInventoryMenu(events <-chan keyboard.KeyEvent) {
	for {
		printInventoryScreen()
		// Attendre une touche utilisateur (drain répétitions)
		e := <-events
		draining := true
		for draining {
			select {
			case next := <-events:
				e = next
			default:
				draining = false
			}
		}
		input := strings.ToLower(string(e.Rune))
		if input == "i" || e.Key == keyboard.KeyEsc || e.Key == keyboard.KeyEnter {
			return
		}
	}
}

func printInventoryScreen() {
	fmt.Print("\033[H\033[2J")
	fmt.Println("════════════════════════════════════════════════════════════")
	fmt.Println("🎒                    INVENTAIRE DU JOUEUR                   🎒")
	fmt.Println("════════════════════════════════════════════════════════════")

	// Section ressources & clés
	fmt.Println("📦 RESSOURCES")
	fmt.Printf("   💰 Pièces:          %d\n", playerInventory["pièces"])
	fmt.Printf("   🔑 Clés:            %d\n", playerInventory["clés"])
	fmt.Printf("   🗝️ Clés spéciales:  %d\n", playerInventory["clés_spéciales"])
	fmt.Printf("   ⛏️ Pioche:          %d\n", playerInventory["pioche"])
	fmt.Printf("   🪨 Roches évol.:    %d\n", currentPlayer.Roches)
	fmt.Printf("   ⚔️ Épées:           %d\n", playerInventory["épées"])
	fmt.Println()

	// Section consommables
	fmt.Println("🧪 CONSOMMABLES")
	fmt.Printf("   Potions:            %d\n", playerInventory["potions"])
	fmt.Printf("   Puff 9K:            %d\n", playerInventory["puff_9k"])
	fmt.Println()

	// Arme & Armure équipées (vue rapide)
	weaponName := "Aucune"
	if currentPlayer.ArmeEquipee.Nom != "" {
		weaponName = currentPlayer.ArmeEquipee.Nom
	} else if currentPlayer.NiveauArme >= 0 && currentPlayer.NiveauArme < len(currentPlayer.ArmesDisponibles) {
		weaponName = currentPlayer.ArmesDisponibles[currentPlayer.NiveauArme].Nom
	}
	armorName := "Aucune"
	if currentPlayer.ArmureEquipee.Nom != "" {
		armorName = currentPlayer.ArmureEquipee.Nom
	} else if currentPlayer.NiveauArmure >= 0 && currentPlayer.NiveauArmure < len(currentPlayer.ArmuresDisponibles) {
		armorName = currentPlayer.ArmuresDisponibles[currentPlayer.NiveauArmure].Nom
	}
	fmt.Println("🛡️ ÉQUIPEMENT")
	fmt.Printf("   Arme actuelle:      %s\n", weaponName)
	fmt.Printf("   Armure actuelle:    %s\n", armorName)
	if playerStats.hasLegendaryWeapon {
		fmt.Println("   🌟 Arme légendaire active: Excalibur")
	}
	fmt.Println()

	// Artefacts équipés
	fmt.Println("🧿 ARTEFACTS")
	artNames := []string{}
	for _, a := range currentPlayer.ArtefactsEquipes {
		if a != nil {
			artNames = append(artNames, a.Nom)
		}
	}
	if len(artNames) == 0 {
		fmt.Println("   Aucun artefact équipé")
	} else {
		// Tri pour affichage stable
		sort.Strings(artNames)
		for _, n := range artNames {
			fmt.Printf("   • %s\n", n)
		}
	}
	fmt.Println()

	fmt.Printf("☠️  Ennemis tués au total: %d\n", playerStats.enemiesKilled)
	fmt.Println("════════════════════════════════════════════════════════════")
	fmt.Println("(Appuyez sur I, ESC ou Entrée pour revenir au jeu)")
}
