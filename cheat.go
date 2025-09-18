package main

import (
	"fmt"
	"strings"

	"github.com/eiannone/keyboard"
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
	// Helper: read a line of text from keyboard events until Enter
	readLine := func() string {
		var b strings.Builder
		for {
			ev := <-globalKeyEvents
			// Drain burst
			draining := true
			for draining {
				select {
				case next := <-globalKeyEvents:
					ev = next
				default:
					draining = false
				}
			}
			if ev.Key == keyboard.KeyEnter { // finish
				break
			}
			if ev.Key == keyboard.KeyBackspace || ev.Rune == 127 {
				s := b.String()
				if len(s) > 0 {
					// remove last rune
					rs := []rune(s)
					rs = rs[:len(rs)-1]
					b.Reset()
					b.WriteString(string(rs))
					// simple backspace visual
					fmt.Print("\b \b")
				}
				continue
			}
			r := ev.Rune
			if r >= 'A' && r <= 'Z' {
				r = r + 32
			}
			// accept letters, digits, underscore
			if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '_' {
				b.WriteRune(r)
				fmt.Printf("%c", r)
			}
		}
		fmt.Println()
		return strings.TrimSpace(b.String())
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
		fmt.Println("f. Toggle rendu ASCII/Emoji")
		fmt.Println("g. Donner une Hache (coupe les arbres)")
		fmt.Println("h. Placer un arbre à droite du joueur")
		fmt.Println("i. Couper/enlever l'arbre à droite du joueur")
		fmt.Println("j. TP salle4 (Marchand)")
		fmt.Println("k. TP salle5 (Forgeron)")
		fmt.Println("l. TP salle7 (Casino)")
		fmt.Println("m. TP salle12 (Mini-boss)")
		fmt.Println("n. TP salle13")
		fmt.Println("o. TP salle14")
		fmt.Println("p. +20 Roches d'évolution")
		fmt.Println("t. TP vers salle par nom (ex: salle1..salle15)")
		fmt.Println("w. Lister les salles disponibles")
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
			applyDecorations(*currentMap, *mapData)
			applyCutTrees(*currentMap, *mapData)
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
			applyDecorations(*currentMap, *mapData)
			applyCutTrees(*currentMap, *mapData)
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
		case 'f':
			useASCII = !useASCII
			if useASCII {
				addHUDMessage("🔤 Rendu ASCII activé")
			} else {
				addHUDMessage("🖼️ Rendu Emoji activé")
			}
		case 'g':
			addToInventory("hache", 1)
			addHUDMessage("🪓 Cheat: Hache ajoutée")
		case 'h':
			// Place un arbre à droite du joueur si libre
			px, py := findPlayer(*mapData)
			x, y := px+1, py
			if y >= 0 && y < len(*mapData) && x >= 0 && x < len((*mapData)[0]) {
				if (*mapData)[y][x] == 0 { // sol vide
					PlaceTreeImmediate(*currentMap, *mapData, x, y)
					addHUDMessage("🌳 Arbre placé à droite")
				} else {
					addHUDMessage("⚠️ Case non vide à droite")
				}
			}
		case 'i':
			// Coupe/enlève un arbre à droite du joueur
			px, py := findPlayer(*mapData)
			x, y := px+1, py
			if y >= 0 && y < len(*mapData) && x >= 0 && x < len((*mapData)[0]) {
				if (*mapData)[y][x] == TileTree {
					RemoveTree(*currentMap, x, y)
					MarkTreeCut(*currentMap, x, y)
					(*mapData)[y][x] = 0
					addHUDMessage("🪓 Arbre enlevé à droite")
				} else {
					addHUDMessage("⚠️ Pas d'arbre à droite")
				}
			}
		case 'j':
			*currentMap = "salle4"
			*mapData = copyMap(salles[*currentMap])
			applyDecorations(*currentMap, *mapData)
			applyCutTrees(*currentMap, *mapData)
			applyEnemyStates(*mapData, *currentMap)
			placePlayerAt(*mapData, 2, 2)
			addHUDMessage("🛒 TP salle4 (Marchand)")
		case 'k':
			*currentMap = "salle5"
			*mapData = copyMap(salles[*currentMap])
			applyDecorations(*currentMap, *mapData)
			applyCutTrees(*currentMap, *mapData)
			applyEnemyStates(*mapData, *currentMap)
			placePlayerAt(*mapData, 2, 2)
			addHUDMessage("⚒️ TP salle5 (Forgeron)")
		case 'l':
			*currentMap = "salle7"
			*mapData = copyMap(salles[*currentMap])
			applyDecorations(*currentMap, *mapData)
			applyCutTrees(*currentMap, *mapData)
			applyEnemyStates(*mapData, *currentMap)
			placePlayerAt(*mapData, 2, 2)
			addHUDMessage("🎰 TP salle7 (Casino)")
		case 'm':
			*currentMap = "salle12"
			*mapData = copyMap(salles[*currentMap])
			applyDecorations(*currentMap, *mapData)
			applyCutTrees(*currentMap, *mapData)
			applyEnemyStates(*mapData, *currentMap)
			placePlayerAt(*mapData, 2, 2)
			addHUDMessage("🗡️ TP salle12 (Mini-boss)")
		case 'n':
			*currentMap = "salle13"
			*mapData = copyMap(salles[*currentMap])
			applyDecorations(*currentMap, *mapData)
			applyCutTrees(*currentMap, *mapData)
			applyEnemyStates(*mapData, *currentMap)
			placePlayerAt(*mapData, 2, 2)
			addHUDMessage("👑 TP salle13")
		case 'o':
			*currentMap = "salle14"
			*mapData = copyMap(salles[*currentMap])
			applyDecorations(*currentMap, *mapData)
			applyCutTrees(*currentMap, *mapData)
			applyEnemyStates(*mapData, *currentMap)
			placePlayerAt(*mapData, 2, 2)
			addHUDMessage("🔥 TP salle14")
		case 'p':
			currentPlayer.Roches += 20
			addHUDMessage("🪨 +20 Roches d'évolution")
		case 't':
			fmt.Print("Nom de la salle à rejoindre: ")
			name := readLine()
			if name == "" {
				addHUDMessage("⚠️ Nom vide.")
				break
			}
			if _, ok := salles[name]; !ok {
				addHUDMessage("❌ Salle inconnue: " + name)
				break
			}
			*currentMap = name
			*mapData = copyMap(salles[*currentMap])
			applyDecorations(*currentMap, *mapData)
			applyCutTrees(*currentMap, *mapData)
			applyEnemyStates(*mapData, *currentMap)
			// spawn par défaut près du coin (2,2)
			placePlayerAt(*mapData, 2, 2)
			addHUDMessage("🧭 TP vers " + name)
		case 'w':
			// Liste simple des clés de 'salles'
			fmt.Println("Salles disponibles:")
			for k := range salles {
				fmt.Println(" -", k)
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
