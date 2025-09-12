package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// RunGameLoop gère les déplacements du joueur et l'affichage
func RunGameLoop(mapData [][]int) {
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

		// Vérifie les limites de la map avant de bouger
		if newY >= 0 && newY < len(mapData) && newX >= 0 && newX < len(mapData[0]) {
			// Vérifie que la case n'est pas un mur
			if mapData[newY][newX] != 9 {
				if mapData[newY][newX] == 2 {
					fmt.Println("💥 MESSIR,UN SARAZIN ! 💥")
				}

				mapData[py][px] = 0
				mapData[newY][newX] = 1
			}
		}
	}
}

// printMap affiche la salle dans le terminal
func printMap(mapData [][]int) {
	fmt.Print("\033[H\033[2J") // Efface l'écran avant chaque affichage
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

// findPlayer retourne la position du joueur (x, y)
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
