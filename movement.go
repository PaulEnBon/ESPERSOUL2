package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// map des salles
var salles = map[string][][]int{
	"salle1": salle1,
	"salle2": salle2,
	"salle3": salle3,
}

// transitions par porte : 7=haut, 10=bas
var transitions = map[string]map[int]string{
	"salle1": {7: "salle2"},
	"salle2": {7: "salle3", 10: "salle1"},
	"salle3": {10: "salle2"},
}

// RunGameLoop gère le jeu, les déplacements et les transitions
func RunGameLoop(mapData [][]int, currentMap string, cameFrom string) {
	reader := bufio.NewReader(os.Stdin)

	// spawn automatique si on revient d'une salle précédente
	if cameFrom != "" {
		spawnX, spawnY := findSpawn(mapData)
		if spawnX != -1 && spawnY != -1 {
			placePlayerAt(mapData, spawnX, spawnY)
		}
	}

	for {
		printMap(mapData)
		fmt.Print("Déplace-toi avec ZQSD (X pour quitter) : ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))

		px, py := findPlayer(mapData)
		newX, newY := px, py

		switch input {
		case "z":
			newY--
		case "s":
			newY++
		case "q":
			newX--
		case "d":
			newX++
		case "x":
			fmt.Println("Jeu terminé.")
			return
		default:
			fmt.Println("Touche inconnue.")
			continue
		}

		// vérifier bornes
		if newY < 0 || newY >= len(mapData) || newX < 0 || newX >= len(mapData[0]) {
			continue
		}

		cell := mapData[newY][newX]

		switch cell {
		case 9: // mur
			continue
		case 2: // ennemi
			fmt.Println("💥 Vous avez rencontré un ennemi !")
		case 7, 10: // porte haut ou bas
			nextMap := transitions[currentMap][cell]
			if nextMap != "" {
				nextMapData := salles[nextMap]

				// ajuster newX si la salle suivante est plus petite
				if newX >= len(nextMapData[0]) {
					newX = len(nextMapData[0]) - 2
				} else if newX <= 0 {
					newX = 1
				}

				// spawn dynamique : chercher case 11
				spawnX, spawnY := findSpawn(nextMapData)
				if spawnX != -1 && spawnY != -1 {
					placePlayerAt(nextMapData, spawnX, spawnY)
				} else {
					// fallback : spawn bas ou haut selon porte
					switch currentMap + "->" + nextMap {
					case "salle1->salle2", "salle2->salle3":
						// on monte → spawn en bas de la salle suivante
						placePlayerAt(nextMapData, newX, len(nextMapData)-2)
					case "salle3->salle2", "salle2->salle1":
						// on descend → spawn en haut de la salle suivante
						placePlayerAt(nextMapData, newX, 1)
					default:
						if cell == 7 {
							placePlayerAt(nextMapData, newX, newY+1)
						} else {
							placePlayerAt(nextMapData, newX, newY-1)
						}
					}
				}

				RunGameLoop(nextMapData, nextMap, currentMap)
				return
			} else {
				fmt.Println("✅ Vous avez fini le donjon !")
				return
			}
		}

		// déplacement normal
		mapData[py][px] = 0
		mapData[newY][newX] = 1
	}
}

// findSpawn retourne les coordonnées de la case 11
func findSpawn(mapData [][]int) (int, int) {
	for y := 0; y < len(mapData); y++ {
		for x := 0; x < len(mapData[y]); x++ {
			if mapData[y][x] == 11 {
				return x, y
			}
		}
	}
	return -1, -1
}

// printMap affiche la salle
func printMap(mapData [][]int) {
	fmt.Print("\033[H\033[2J")
	for _, row := range mapData {
		for _, val := range row {
			switch val {
			case 8:
				fmt.Print(" ๑ ")
			case 9:
				fmt.Print(" ▨ ")
			case 7:
				fmt.Print(" ↑ ")
			case 10:
				fmt.Print(" ↓ ")
			case 1:
				fmt.Print("💩 ")
			case 2:
				fmt.Print("😈 ")
			case 11:
				fmt.Print(" • ") // invisible, juste pour spawn dynamique
			case 0:
				fmt.Print(" • ")
			default:
				fmt.Print(" ? ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

// findPlayer retourne la position du joueur
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

// placePlayerAt place le joueur à une position spécifique
func placePlayerAt(mapData [][]int, x, y int) {
	for row := range mapData {
		for col := range mapData[row] {
			if mapData[row][col] == 1 {
				mapData[row][col] = 0
			}
		}
	}
	mapData[y][x] = 1
}
