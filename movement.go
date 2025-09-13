package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Map globale des salles
var salles = map[string][][]int{
	"salle1": salle1,
	"salle2": salle2,
	"salle3": salle3,
}

// Map des transitions : porte haut/porte bas
var transitions = map[string]map[int]struct {
	nextMap string
	spawnX  int
	spawnY  int
}{
	"salle1": {
		7: {nextMap: "salle2", spawnX: 2, spawnY: 1}, // ↑ vers salle2
	},
	"salle2": {
		7:  {nextMap: "salle3", spawnX: 8, spawnY: 1}, // ↑ vers salle3
		10: {nextMap: "salle1", spawnX: 2, spawnY: 5}, // ↓ vers salle1
	},
	"salle3": {
		10: {nextMap: "salle2", spawnX: 2, spawnY: 1}, // ↓ vers salle2
	},
}

// Boucle principale du jeu
func RunGameLoop(currentMap string) {
	reader := bufio.NewReader(os.Stdin)
	mapData := copyMap(salles[currentMap])

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
		}

		// Limites
		if newY < 0 || newY >= len(mapData) || newX < 0 || newX >= len(mapData[0]) {
			continue
		}

		cell := mapData[newY][newX]

		switch cell {
		case 9:
			continue
		case 2:
			fmt.Println("💥 Vous avez rencontré un ennemi !")
		case 7, 10: // portes
			if tr, ok := transitions[currentMap][cell]; ok {
				nextMapData := copyMap(salles[tr.nextMap])
				placePlayerAt(nextMapData, tr.spawnX, tr.spawnY)
				RunGameLoop(tr.nextMap)
				return
			}
		}

		mapData[py][px] = 0
		mapData[newY][newX] = 1
	}
}

func copyMap(src [][]int) [][]int {
	dst := make([][]int, len(src))
	for i := range src {
		dst[i] = make([]int, len(src[i]))
		copy(dst[i], src[i])
	}
	return dst
}

func printMap(mapData [][]int) {
	fmt.Print("\033[H\033[2J") // Clear terminal
	for _, row := range mapData {
		for _, val := range row {
			switch val {
			case 8:
				fmt.Print(" ๑ ")
			case 10:
				fmt.Print(" ↓ ")
			case 9:
				fmt.Print(" ▨ ")
			case 7:
				fmt.Print(" ↑ ")
			case 1:
				fmt.Print("💩 ")
			case 2:
				fmt.Print("😈 ")
			case 11:
				fmt.Print(" • ")
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
	for i := range mapData {
		for j := range mapData[i] {
			if mapData[i][j] == 1 {
				mapData[i][j] = 0
			}
		}
	}
	mapData[y][x] = 1
}
