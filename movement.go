package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// map de salles : relie un nom à une salle
var salles = map[string][][]int{
	"salle1": salle1,
	"salle2": salle2, // salle2 must be defined below
}

// Define salle2 as a 2D slice of integers
var salle2 = [][]int{
	{9, 9, 9, 9, 9},
	{9, 0, 0, 7, 9},
	{9, 0, 1, 0, 9},
	{9, 9, 9, 9, 9},
}

// map de transitions : quand on marche sur une porte
var transitions = map[string]string{
	"salle1": "salle2",
	"salle2": "", // fin du donjon
}

// RunGameLoop gère les déplacements et l'affichage
func RunGameLoop(mapData [][]int, currentMap string) {
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
		}

		if newY >= 0 && newY < len(mapData) && newX >= 0 && newX < len(mapData[0]) {
			switch mapData[newY][newX] {
			case 9: // mur
				continue
			case 2: // ennemi
				fmt.Println("💥 Vous avez rencontré un ennemi !")
			case 7: // porte
				fmt.Println("🚪 Vous passez dans la salle suivante...")
				nextMap := transitions[currentMap]
				if nextMap != "" {
					RunGameLoop(salles[nextMap], nextMap)
				} else {
					fmt.Println("✅ Vous avez fini le donjon !")
				}
				return
			}

			// Déplace le joueur
			mapData[py][px] = 0
			mapData[newY][newX] = 1
		}
	}
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
			case 1:
				fmt.Print("💩 ")
			case 2:
				fmt.Print("😈 ")
			case 0:
				fmt.Print(" • ")
			default:
				fmt.Printf("%d ", val)
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
