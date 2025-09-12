package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// --- Types ---
type Spawn struct{ X, Y int }
type Door struct {
	NextMap string
	SpawnX  int
	SpawnY  int
}

// --- Portes par salle ---
var doors = map[string]map[[2]int]Door{
	"salle1": {
		{8, 0}: {"salle2", 2, 7}, // porte haut salle1 → spawn salle2 X=2,Y=7
	},
	"salle2": {
		{2, 0}: {"salle3", 8, 13}, // porte haut salle2 → spawn salle3
		{2, 8}: {"salle3", 8, 1},  // porte bas salle2 → spawn salle3
	},
	"salle3": {
		{8, 0}: {"salle4", 8, 13}, // future salle4
	},
}

// --- Map des salles ---
var salles = map[string][][]int{
	"salle1": salle1,
	"salle2": salle2,
	"salle3": salle3,
}

// --- Boucle principale ---
func RunGameLoopSafe() {
	currentMap := "salle1"
	mapData := copyMap(salles[currentMap])
	reader := bufio.NewReader(os.Stdin)

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

		if newY >= 0 && newY < len(mapData) && newX >= 0 && newX < len(mapData[0]) {
			switch mapData[newY][newX] {
			case 9: // mur
				continue
			case 2: // ennemi
				fmt.Println("💥 Vous avez rencontré un ennemi !")
			case 7, 10: // porte haut/bas
				if door, ok := doors[currentMap][[2]int{newX, newY}]; ok {
					fmt.Println("🚪 Vous passez dans la salle suivante...")
					currentMap = door.NextMap
					mapData = copyMap(salles[currentMap])
					// Spawn sécurisé
					if door.SpawnX >= 0 && door.SpawnX < len(mapData[0]) &&
						door.SpawnY >= 0 && door.SpawnY < len(mapData) {
						placePlayerAt(mapData, door.SpawnX, door.SpawnY)
					} else {
						c := centerSpawn(mapData)
						placePlayerAt(mapData, c.X, c.Y)
					}
					continue
				} else {
					fmt.Println("Porte inconnue.")
				}
			}

			// Déplacement normal
			mapData[py][px] = 0
			mapData[newY][newX] = 1
		}
	}
}

// --- Spawn au centre ---
func centerSpawn(mapData [][]int) (s Spawn) {
	s.X = len(mapData[0]) / 2
	s.Y = len(mapData) / 2
	return
}

// --- Affichage ---
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

// --- Utilitaires ---
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

func copyMap(original [][]int) [][]int {
	cpy := make([][]int, len(original))
	for i := range original {
		cpy[i] = make([]int, len(original[i]))
		copy(cpy[i], original[i])
	}
	return cpy
}
