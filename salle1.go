package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var salle1 = [7][17]int{
	{9, 9, 9, 9, 9, 9, 9, 9, 7, 9, 9, 9, 9, 9, 9, 9, 9},
	{9, 0, 0, 0, 0, 0, 0, 0, 8, 0, 0, 0, 0, 0, 0, 0, 9},
	{9, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 9},
	{9, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 9},
	{9, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 9},
	{9, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 9},
	{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
}

// StartSalle1 lance la boucle de jeu pour cette salle
func StartSalle1() {
	reader := bufio.NewReader(os.Stdin)

	for {
		printMap()
		fmt.Print("Déplace-toi avec ZQSD (X pour quitter) : ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))

		px, py := findPlayer()
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

		// Vérifie que la case n'est pas un mur avant de bouger
		if salle1[newY][newX] != 9 {
			// Si c'est un ennemi
			if salle1[newY][newX] == 2 {
				fmt.Println("MESSIRE, UN SARAZIN !")
				// Ici, tu pourrais lancer un combat
			}

			// Déplace le joueur
			salle1[py][px] = 0
			salle1[newY][newX] = 1
		}
	}
}

// printMap affiche la salle dans le terminal
func printMap() {
	fmt.Print("\033[H\033[2J") // Efface l'écran avant chaque affichage
	for _, row := range salle1 {
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
func findPlayer() (int, int) {
	for y := 0; y < len(salle1); y++ {
		for x := 0; x < len(salle1[y]); x++ {
			if salle1[y][x] == 1 {
				return x, y
			}
		}
	}
	return -1, -1
}
