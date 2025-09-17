package main

import (
	"fmt"
)

// Activer/désactiver les cheats (mettre true pour tester). L'ouverture se fait avec la touche 'k'.
// Mettre à false avant de distribuer le jeu.
var debugMode = true

// God mode: rend le joueur invincible (à implémenter dans calcul des dégâts combat si besoin)
var godMode = false

// Menu de cheats déclenché par la touche 'c' dans la boucle de jeu
func showCheatMenu(currentMap *string, mapData *[][]int) {
	if globalKeyEvents == nil {
		return
	}
	for {
		fmt.Println("\n===== CHEAT MENU =====")
		fmt.Println("1. Donner l'objet spécial (sida)")
		fmt.Println("2. Aller directement en salle15")
		fmt.Println("3. Faire apparaître le boss salle15 maintenant")
		fmt.Println("4. Soigner complètement le joueur")
		fmt.Println("5. Ajouter 500 pièces")
		fmt.Println("6. Tuer instantanément le boss (flag vaincu)")
		fmt.Println("7. Toggle God Mode (invincibilité en combat)")
		fmt.Println("8. +1 Niveau joueur (augmente stats de base)")
		fmt.Println("9. Donner un lot de potions (mineure/majeure/suprême x5)")
		fmt.Println("a. Téléportation salle1")
		fmt.Println("b. Reset progression séquence cheat (cheatProgress=0)")
		fmt.Println("c. Donner Puff 9K (+1 boost attaque)")
		fmt.Println("d. Booster attack x2 (bonus temporaire)")
		fmt.Println("e. Apparaitre PNJ soigneur (pose un code tile)")
		fmt.Println("q. Quitter le menu")
		fmt.Print("Choix: ")
		ev := <-globalKeyEvents
		c := ev.Rune
		if c >= 'A' && c <= 'Z' {
			c += 32
		}
		switch c {
		case '1':
			addToInventory("sida", 1)
			addHUDMessage("🧪 Cheat: objet 'sida' ajouté.")
		case '2':
			*currentMap = "salle15"
			*mapData = copyMap(salles[*currentMap])
			applyEnemyStates(*mapData, *currentMap)
			placePlayerAt(*mapData, 3, 6)
			addHUDMessage("🗺️ Cheat: téléporté en salle15.")
		case '3':
			if *currentMap == "salle15" {
				(*mapData)[3][3] = 78
				addHUDMessage("👹 Cheat: Boss final invoqué.")
			} else {
				addHUDMessage("⚠️ Pas en salle15.")
			}
		case '4':
			currentPlayer.PV = currentPlayer.PVMax
			addHUDMessage("❤️ Cheat: PV restaurés.")
		case '5':
			addToInventory("pièces", 500)
			addHUDMessage("💰 Cheat: +500 pièces.")
		case '6':
			if *currentMap == "salle15" {
				if cfg, ok := bossRooms["salle15"]; ok {
					cfg.state.bossDefeated = true
					addHUDMessage("💀 Cheat: Boss marqué vaincu.")
				}
			} else {
				addHUDMessage("⚠️ Pas en salle15.")
			}
		case '7':
			// Toggle d'un flag global de god mode (ajouter variable si inexistante)
			godMode = !godMode
			if godMode {
				addHUDMessage("🛡️ God Mode ACTIVÉ")
			} else {
				addHUDMessage("🛡️ God Mode désactivé")
			}
		case '8':
			// Simule un gain de niveau : on augmente quelques stats basiques
			currentPlayer.PVMax += 5
			currentPlayer.PV = currentPlayer.PVMax
			currentPlayer.Armure += 1
			currentPlayer.ResistMag += 1
			currentPlayer.Precision += 0.01
			currentPlayer.TauxCritique += 0.01
			addHUDMessage("📈 Gain simulé: +5 PVMax, +1 Armure, +1 RM, +1% Prec & Crit")
		case '9':
			addToInventory("potion_mineure", 5)
			addToInventory("potion_majeure", 5)
			addToInventory("potion_supreme", 5)
			addHUDMessage("🧪 Pack potions ajouté x5 chacun")
		case 'a':
			*currentMap = "salle1"
			*mapData = copyMap(salles[*currentMap])
			applyEnemyStates(*mapData, *currentMap)
			placePlayerAt(*mapData, 2, 2)
			addHUDMessage("🚪 Téléporté salle1")
		case 'b':
			cheatProgress = 0
			addHUDMessage("♻️ Séquence cheat reset")
		case 'c':
			addToInventory("puff_9k", 1)
			playerStats.attackBoost += 15
			addHUDMessage("💊 Puff 9K ajouté (+15% attaque)")
		case 'd':
			playerStats.attackBoost += 100
			addHUDMessage("⚔️ Attack boost massif (+100%)")
		case 'e':
			// On place un PNJ soigneur codé 90 (valeur arbitraire libre si non utilisée)
			// On vérifie la taille avant de poser
			if len(*mapData) > 4 && len((*mapData)[4]) > 4 {
				(*mapData)[4][4] = 90
				addHUDMessage("🚑 PNJ soigneur apparu (4,4)")
			} else {
				addHUDMessage("⚠️ Carte trop petite pour PNJ soigneur")
			}
		case 'q':
			addHUDMessage("Fermeture cheat menu.")
			return
		default:
			addHUDMessage("Entrée invalide cheat.")
		}
		// Mettre à jour refs globales
		currentMapGlobalRef = *currentMap
		mapDataGlobalRef = *mapData
	}
}
