package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Bienvenue dans ESPER SOUL 2!")
	fmt.Println("Appuyez sur Entrée pour commencer...")

	reader := bufio.NewReader(os.Stdin)
	_, _ = reader.ReadString('\n')

	// Initialiser la salle1 avec le joueur
	currentMap := "salle1"
	fmt.Println("Initialisation de la partie dans la salle1...")

	// Lancement du jeu
	RunGameLoop(currentMap)
}

// RunGameLoop gère la boucle principale du jeu
// La fonction RunGameLoop complète est définie dans game_loop.go
