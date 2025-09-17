package main

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/mattn/go-runewidth"
)

// Positions des mobs aléatoires dans salle3
var randomMobsSalle3 = []struct{ x, y int }{}

// Positions des mobs aléatoires dans salle2
var randomMobsSalle2 = []struct{ x, y int }{}

// Positions des mobs (salle10)
var randomMobsSalle10 = []struct{ x, y int }{}

// Configuration du nombre de mobs aléatoires dans salle3
const (
	minRandomMobsSalle3 = 5
	maxRandomMobsSalle3 = 8
)

// Nombre fixe de mobs aléatoires dans salle2
const numRandomMobsSalle2 = 4

// Nombre de mobs aléatoires dans salle10
const numRandomMobsSalle10 = 5

// Copie de la map
func copyMap(src [][]int) [][]int {
	dst := make([][]int, len(src))
	for i := range src {
		dst[i] = make([]int, len(src[i]))
		copy(dst[i], src[i])
	}
	return dst
}

// Convertit une cellule en symbole
var useASCII = false // Passez à true pour un mode ASCII aligné sans emojis

func cellToSymbol(val, x, y int, currentMap string) string {
	if useASCII {
		switch val {
		case 9:
			return "#"
		case 1:
			return "@"
		case 2:
			return "E"
		case 3:
			return "N"
		case 4:
			return "M"
		case 5:
			return "F"
		case 6:
			return "C"
		case 7, 10, 13, 14, 15, 20, 21, 27, 28, 31, 33, 34, 38, 40, 42, 44:
			return "P"
		case 30:
			return "S"
		case 32:
			return "X"
		default:
			return "."
		}
	}
	switch val {
	case 8:
		return "⬜"
	case 14, 38:
		return "→"
	case 13, 40:
		return "←"
	case 15, 27, 31, 10, 21, 42:
		return "↓"
	case 28, 33, 34, 44:
		return "↑"
	case 9:
		return "▨"
	case 30: // Porte secrète
		return "▧"
	case 32: // Sortie de la salle secrète
		return "🚪"
	case 7, 20:
		return "↑"
	case 1:
		return "🎮" // Joueur plus visible
	case 2, 12:
		// Affiche un emoji spécifique selon le type assigné
		key := fmt.Sprintf("%d_%d", x, y)
		if currentMap != "" {
			if name, ok := enemyAssignments[currentMap][key]; ok && name != "" {
				return emojiForEnemyName(name)
			}
		}
		// Fallbacks si pas d'assignation
		if val == 12 {
			return "💀"
		}
		return "👹"
	case 3:
		return "👨" // PNJ
	case 4:
		return "🛒" // Marchand
	case 5:
		return "🔨" // Forgeron
	case 6:
		return "💰" // Coffre
	case 11, 16, 17, 18, 19, 22, 23, 24, 25, 26, 29, 36, 37, 39, 41, 43, 45:
		return "•"
	case 0:
		return "•"
	default:
		return "?"
	}
}

// Affiche la map avec HUD optimisé
func printMap(mapData [][]int, currentMap string) {
	fmt.Print("\033[H\033[2J") // Nettoie l'écran

	// En-tête du jeu
	fmt.Println("═══════════════════════════════════════════════════════════════════════")
	fmt.Println("🎮                        DONJON MYSTÉRIEUX                        🎮")
	fmt.Println("═══════════════════════════════════════════════════════════════════════")

	// Prépare les lignes d'informations pour l'affichage côte à côte
	// Préparer la ligne des artefacts équipés
	artNames := []string{}
	for _, slot := range currentPlayer.ArtefactsEquipes {
		if slot != nil {
			artNames = append(artNames, slot.Nom)
		}
	}
	artefactsStr := "Aucun"
	if len(artNames) > 0 {
		artefactsStr = strings.Join(artNames, ", ")
	}

	infoLines := []string{
		"📊 === STATISTIQUES ===",
		fmt.Sprintf("💰 Pièces: %d", playerInventory["pièces"]),
		fmt.Sprintf("� Roches: %d", currentPlayer.Roches),
		fmt.Sprintf("🔑 Clés: %d", playerInventory["clés"]),
		fmt.Sprintf("🗝️  Clés spéciales: %d", playerInventory["clés_spéciales"]),
		fmt.Sprintf("💊 Puff 9K: %d", playerInventory["puff_9k"]),
		fmt.Sprintf("🧿 Artefacts: %s", artefactsStr),
		fmt.Sprintf("☠️ Ennemis tués: %d", playerStats.enemiesKilled),
		"",
		"🏆 === BONUS ===",
	}

	if playerStats.hasLegendaryWeapon {
		infoLines = append(infoLines, "🌟 Excalibur (+50% ATK)")
	} else {
		infoLines = append(infoLines, "🌟 Pas d'arme légendaire")
	}

	// Affiche la carte avec les infos côte à côte
	mapHeight := len(mapData)
	maxLines := max(mapHeight, len(infoLines))

	cellWidth := 2 // Largeur cible minimale visuelle
	if useASCII {
		cellWidth = 1
	}

	for i := 0; i < maxLines; i++ {
		if i < mapHeight {
			var b strings.Builder
			for j, val := range mapData[i] {
				sym := cellToSymbol(val, j, i, currentMap)
				w := runewidth.StringWidth(sym)
				pad := cellWidth - w
				if pad < 0 {
					pad = 0
				}
				// Un espace devant pour aérer
				b.WriteString(" ")
				b.WriteString(sym)
				b.WriteString(strings.Repeat(" ", pad))
			}
			fmt.Print(b.String())
		} else {
			// Ligne vide alignée
			lineWidth := len(mapData[0]) * (cellWidth + 1)
			fmt.Print(strings.Repeat(" ", lineWidth))
		}

		fmt.Print("   │   ")
		if i < len(infoLines) {
			fmt.Print(infoLines[i])
		}
		fmt.Println()
	}

	// Ligne de séparation
	fmt.Println("═══════════════════════════════════════════════════════════════════════")
}

// Affiche un HUD compact pour l'inventaire uniquement
func showCompactInventory() {
	// Compte des artefacts équipés
	artCount := 0
	for _, a := range currentPlayer.ArtefactsEquipes {
		if a != nil {
			artCount++
		}
	}

	fmt.Printf("💰:%d 🔑:%d 🗝️:%d 💊:%d 🪨:%d",
		playerInventory["pièces"],
		playerInventory["clés"],
		playerInventory["clés_spéciales"],
		playerInventory["puff_9k"],
		currentPlayer.Roches)
	if artCount > 0 {
		fmt.Printf(" 🧿:%d", artCount)
	}
	if playerStats.hasLegendaryWeapon {
		fmt.Print(" 🌟:Excalibur")
	}
	fmt.Println()
}

// Fonction helper pour max
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Trouve le joueur
func findPlayer(mapData [][]int) (int, int) {
	for y := 0; y < len(mapData); y++ {
		for x := 0; x < len(mapData[y]); x++ {
			if mapData[y][x] == 1 {
				return x, y
			}
		}
	}
	return -1, -1
}

// Place le joueur à un point spécifique
func placePlayerAt(mapData [][]int, x, y int) {
	if y < 0 || y >= len(mapData) || x < 0 || x >= len(mapData[0]) {
		for i := 0; i < len(mapData); i++ {
			for j := 0; j < len(mapData[i]); j++ {
				if mapData[i][j] == 0 {
					x, y = j, i
					goto placePlayer
				}
			}
		}
		return
	}

placePlayer:
	for i := range mapData {
		for j := range mapData[i] {
			if mapData[i][j] == 1 {
				mapData[i][j] = 0
			}
		}
	}

	if mapData[y][x] == 9 {
		for _, offset := range []struct{ dx, dy int }{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
			nx, ny := x+offset.dx, y+offset.dy
			if ny >= 0 && ny < len(mapData) && nx >= 0 && nx < len(mapData[0]) && mapData[ny][nx] != 9 {
				x, y = nx, ny
				break
			}
		}
	}

	mapData[y][x] = 1
}

// Place le joueur près d'une position donnée
func placePlayerNearby(mapData [][]int, targetX, targetY int) {
	for _, offset := range []struct{ dx, dy int }{{0, 1}, {1, 0}, {0, -1}, {-1, 0}, {1, 1}, {-1, -1}, {1, -1}, {-1, 1}} {
		nx, ny := targetX+offset.dx, targetY+offset.dy
		if ny >= 0 && ny < len(mapData) && nx >= 0 && nx < len(mapData[0]) && mapData[ny][nx] == 0 {
			mapData[ny][nx] = 1
			return
		}
	}

	for y := 0; y < len(mapData); y++ {
		for x := 0; x < len(mapData[y]); x++ {
			if mapData[y][x] == 0 {
				mapData[y][x] = 1
				return
			}
		}
	}
}

// Génère des mobs aléatoires dans la salle3
func generateRandomMobs(mapData [][]int) {
	randomMobsSalle3 = []struct{ x, y int }{}

	// génère entre minRandomMobsSalle3 et maxRandomMobsSalle3 (inclus)
	rangeSize := maxRandomMobsSalle3 - minRandomMobsSalle3 + 1
	if rangeSize < 1 {
		rangeSize = 1
	}
	numMobs := minRandomMobsSalle3 + rand.Intn(rangeSize)
	attempts := 0
	maxAttempts := 200

	for len(randomMobsSalle3) < numMobs && attempts < maxAttempts {
		x := 1 + rand.Intn(13)
		y := 1 + rand.Intn(12)

		if mapData[y][x] == 0 {
			isOccupied := false
			for _, mob := range randomMobsSalle3 {
				if mob.x == x && mob.y == y {
					isOccupied = true
					break
				}
			}

			if !isOccupied {
				randomMobsSalle3 = append(randomMobsSalle3, struct{ x, y int }{x, y})
				mapData[y][x] = 2
				// Assigner un type d'ennemi aléatoire selon le tier de la salle
				pool := tutorialPool // fallback
				if tierForMap("salle3") == TierEarly {
					pool = earlyPool
				}
				if len(pool) > 0 {
					chosen := pool[rand.Intn(len(pool))]
					key := fmt.Sprintf("%d_%d", x, y)
					enemyAssignments["salle3"][key] = chosen.Name
				}
			}
		}
		attempts++
	}

	fmt.Printf("🎲 %d mobs aléatoires générés dans la salle3!\n", len(randomMobsSalle3))
}

// Génère 4 mobs aléatoires dans la salle2
func generateRandomMobsSalle2(mapData [][]int) {
	randomMobsSalle2 = []struct{ x, y int }{}

	h := len(mapData)
	if h == 0 {
		return
	}
	w := len(mapData[0])

	attempts := 0
	maxAttempts := 200

	// Restreindre aux cases intérieures pour éviter murs/portes sur les bords
	for len(randomMobsSalle2) < numRandomMobsSalle2 && attempts < maxAttempts {
		if w <= 2 || h <= 2 {
			break
		}
		x := 1 + rand.Intn(w-2)
		y := 1 + rand.Intn(h-2)

		if mapData[y][x] == 0 { // seulement sur sol vide
			// éviter doublons
			isOccupied := false
			for _, mob := range randomMobsSalle2 {
				if mob.x == x && mob.y == y {
					isOccupied = true
					break
				}
			}

			if !isOccupied {
				randomMobsSalle2 = append(randomMobsSalle2, struct{ x, y int }{x, y})
				mapData[y][x] = 2
				// Assigner un type d'ennemi aléatoire selon le tier de la salle
				pool := earlyPool
				if len(pool) > 0 {
					chosen := pool[rand.Intn(len(pool))]
					key := fmt.Sprintf("%d_%d", x, y)
					enemyAssignments["salle2"][key] = chosen.Name
				}
			}
		}
		attempts++
	}

	fmt.Printf("🎲 %d mobs aléatoires générés dans la salle2!\n", len(randomMobsSalle2))
}

// Génère entre 10 et 15 ennemis dans la salle9 à chaque entrée
func generateRandomMobsSalle9(mapData [][]int) {
	h := len(mapData)
	if h == 0 {
		return
	}
	w := len(mapData[0])

	// Nombre aléatoire entre 10 et 15 inclus
	numMobs := 10 + rand.Intn(6)

	placed := 0
	attempts := 0
	maxAttempts := 500

	for placed < numMobs && attempts < maxAttempts {
		if w <= 2 || h <= 2 {
			break
		}
		// Choisir uniquement les cases intérieures
		x := 1 + rand.Intn(w-2)
		y := 1 + rand.Intn(h-2)

		// Placer uniquement sur sol vide (0), éviter portes/marqueurs
		if mapData[y][x] == 0 {
			mapData[y][x] = 2
			// Assigner un type d'ennemi aléatoire selon le tier Late
			pool := latePool
			if len(pool) > 0 {
				chosen := pool[rand.Intn(len(pool))]
				key := fmt.Sprintf("%d_%d", x, y)
				enemyAssignments["salle9"][key] = chosen.Name
			}
			placed++
		}
		attempts++
	}
}

// Génère des mobs (emplacements) dans la salle10
func generateRandomMobsSalle10(mapData [][]int) {
	randomMobsSalle10 = []struct{ x, y int }{}

	h := len(mapData)
	if h == 0 {
		return
	}
	w := len(mapData[0])

	attempts := 0
	maxAttempts := 300

	for len(randomMobsSalle10) < numRandomMobsSalle10 && attempts < maxAttempts {
		if w <= 2 || h <= 2 {
			break
		}
		x := 1 + rand.Intn(w-2)
		y := 1 + rand.Intn(h-2)

		if mapData[y][x] == 0 {
			// éviter doublons
			occupied := false
			for _, m := range randomMobsSalle10 {
				if m.x == x && m.y == y {
					occupied = true
					break
				}
			}
			if !occupied {
				randomMobsSalle10 = append(randomMobsSalle10, struct{ x, y int }{x, y})
				// On n'écrit pas ici: l'écriture (2 ou 12) est faite dans applyEnemyStates (salle10)
			}
		}
		attempts++
	}
}
