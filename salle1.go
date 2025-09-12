package main

import "fmt"

func Salle1() {
	salle1 := [7][17]int{
		{9, 9, 9, 9, 9, 9, 9, 9, 7, 9, 9, 9, 9, 9, 9, 9, 9},
		{9, 0, 0, 0, 0, 0, 0, 0, 8, 0, 0, 0, 0, 0, 0, 0, 9},
		{9, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 9},
		{9, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 9},
		{9, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 9},
		{9, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 9},
		{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
	}

	for _, row := range salle1 {
		for _, val := range row {
			switch val {

			case 8:
				fmt.Print("๑ ")
			case 9:
				fmt.Print("▨ ") // par ex. mur
			case 7:
				fmt.Print("↑ ")
			case 1:
				fmt.Print("☺ ") // joueur par ex.
			case 2:
				fmt.Print("😈 ")
			case 0:
				fmt.Print(". ") // sol vide
			default:
				fmt.Printf("%d ", val)
			}
		}
		fmt.Println()
	}
}
