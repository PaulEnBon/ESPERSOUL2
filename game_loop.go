package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// Note: Since Go 1.20, the global RNG is automatically seeded; no need to seed manually.

// Map globale des salles (à importer depuis votre fichier maps.go existant)
var salles = map[string][][]int{
	"salle1":  salle1,
	"salle2":  salle2,
	"salle3":  salle3,
	"salle4":  salle4,
	"salle5":  salle5,  // Salle forgeron
	"salle6":  salle6,  // Salle coffre
	"salle7":  salle7,  // Salle gambling
	"salle8":  salle8,  // Salle secrète
	"salle9":  salle9,  // Nouvelle salle
	"salle10": salle10, // Nouvelle salle 10
}

// Map pour suivre l'état des coffres ouverts
var chestOpened = make(map[string]bool)

// Map pour suivre l'état des coffres secrets ouverts
var secretChestsOpened = make(map[string]bool)

// Applique l'état des ennemis vaincus sur la map
func applyEnemyStates(mapData [][]int, currentMap string) {
	for y := 0; y < len(mapData); y++ {
		for x := 0; x < len(mapData[y]); x++ {
			if mapData[y][x] == 2 { // Si c'est un ennemi
				enemyKey := fmt.Sprintf("%d_%d", x, y)
				if enemiesDefeated[currentMap][enemyKey] {
					// Ne transforme plus automatiquement en PNJ.
					// L'ennemi vaincu disparaît.
					mapData[y][x] = 0
				}
			}
		}
	}

	// Cas spécial: le seul ennemi de salle1 (8,3) reste PNJ si transformé
	if currentMap == "salle1" {
		key := fmt.Sprintf("%d_%d", 8, 3)
		if pnjTransformed[currentMap][key] {
			mapData[3][8] = 3
		}
	}

	// Générer des mobs aléatoires dans salle3 si c'est la première visite
	if currentMap == "salle3" && len(randomMobsSalle3) == 0 {
		generateRandomMobs(mapData)
	} else if currentMap == "salle3" {
		// Replacer les mobs aléatoires s'ils n'ont pas été vaincus
		for _, mob := range randomMobsSalle3 {
			enemyKey := fmt.Sprintf("%d_%d", mob.x, mob.y)
			if !enemiesDefeated[currentMap][enemyKey] {
				mapData[mob.y][mob.x] = 2 // Remettre l'ennemi
			} else {
				mapData[mob.y][mob.x] = 0 // Disparu si vaincu
			}
		}
	} else if currentMap == "salle2" {
		// Générer 4 ennemis aléatoires lors de la première visite de salle2
		if len(randomMobsSalle2) == 0 {
			generateRandomMobsSalle2(mapData)
		} else {
			// Replacer/convertir selon l'état de défaite
			for _, mob := range randomMobsSalle2 {
				enemyKey := fmt.Sprintf("%d_%d", mob.x, mob.y)
				if !enemiesDefeated[currentMap][enemyKey] {
					mapData[mob.y][mob.x] = 2
				} else {
					mapData[mob.y][mob.x] = 0
				}
			}
		}
	} else if currentMap == "salle9" {
		// Salle9: à chaque entrée, on nettoie ennemis et on régénère 10-15 ennemis
		for y := 0; y < len(mapData); y++ {
			for x := 0; x < len(mapData[y]); x++ {
				if mapData[y][x] == 2 || mapData[y][x] == 3 { // efface ennemis temporaires ou anciens PNJ
					mapData[y][x] = 0
				}
			}
		}
		generateRandomMobsSalle9(mapData)
	} else if currentMap == "salle10" {
		// Générer positions si première visite
		if len(randomMobsSalle10) == 0 {
			generateRandomMobsSalle10(mapData)
			// 50/50 super flag pour chaque ennemi placé
			for _, mob := range randomMobsSalle10 {
				key := fmt.Sprintf("%d_%d", mob.x, mob.y)
				isSuper := rand.Intn(2) == 0
				superEnemyFlags[currentMap][key] = isSuper
				mapData[mob.y][mob.x] = 2
				if isSuper {
					mapData[mob.y][mob.x] = 12
				}
			}
		} else {
			// Replacer selon l'état
			for _, mob := range randomMobsSalle10 {
				key := fmt.Sprintf("%d_%d", mob.x, mob.y)
				if enemiesDefeated[currentMap][key] {
					// Ennemi disparu: 50/50 respawn (plus de PNJ)
					if rand.Intn(2) == 0 {
						// respawn; 50% super
						isSuper := rand.Intn(2) == 0
						superEnemyFlags[currentMap][key] = isSuper
						mapData[mob.y][mob.x] = 2
						if isSuper {
							mapData[mob.y][mob.x] = 12
						}
						enemiesDefeated[currentMap][key] = false
					} else {
						mapData[mob.y][mob.x] = 0 // reste vide
					}
				} else {
					// non défait encore
					if superEnemyFlags[currentMap][key] {
						mapData[mob.y][mob.x] = 12
					} else {
						mapData[mob.y][mob.x] = 2
					}
				}
			}
		}
	}
}

// Gère les interactions avec les différents types de cases
func handleCellInteraction(cell int, currentMap string, newX, newY int, mapData [][]int, px, py int) (bool, string) {
	switch cell {
	case 9: // mur
		return false, currentMap
	case 4: // Marchand permanent ou PNJ gambling
		if currentMap == "salle7" {
			fmt.Println("Vous approchez du croupier...")
			showDialogue(currentMap, newX, newY)
		} else {
			fmt.Println("Vous approchez du marchand...")
			showDialogue(currentMap, newX, newY)
		}
		return false, currentMap
	case 5: // Forgeron
		fmt.Println("Vous approchez du forgeron...")
		showDialogue(currentMap, newX, newY)
		return false, currentMap
	case 6: // Coffre normal
		fmt.Println("Vous trouvez un coffre mystérieux !")
		openChest(currentMap, newX, newY)
		return false, currentMap
	case 8: // Coffre secret (salle secrète)
		if currentMap == "salle8" {
			openSecretChest(newX, newY) // Utilise la fonction spéciale pour les coffres secrets
			return false, currentMap
		}
	case 2, 12: // ennemi (2=normal, 12=super)
		fmt.Println("Vous rencontrez une créature maudite !")
		isSuper := (cell == 12)
		result := combat(currentMap, isSuper)

		enemyKey := fmt.Sprintf("%d_%d", newX, newY)

		if result == "disappear" {
			// Cas spécial: dans salle1 à (8,3) toujours transformer en PNJ
			if currentMap == "salle1" && newX == 8 && newY == 3 {
				enemiesDefeated[currentMap][enemyKey] = true
				pnjTransformed[currentMap][enemyKey] = true
				mapData[py][px] = 0
				mapData[newY][newX] = 3
				if py != newY || px != newX {
					mapData[py][px] = 1
				} else {
					placePlayerNearby(mapData, newX, newY)
				}
				fmt.Println("La créature retrouve sa forme humaine et devient un PNJ amical !")
				showDialogue(currentMap, newX, newY)
			} else {
				enemiesDefeated[currentMap][enemyKey] = true
				mapData[py][px] = 0
				mapData[newY][newX] = 1
				fmt.Println("Vous pouvez maintenant passer par cette case.")
			}
		} else if result == true {
			// Cas spécial: autoriser la transformation en PNJ UNIQUEMENT
			// pour l'unique mob de salle1 (coordonnées 8,3 dans salle1).
			if currentMap == "salle1" && newX == 8 && newY == 3 {
				enemiesDefeated[currentMap][enemyKey] = true
				mapData[py][px] = 0
				mapData[newY][newX] = 3
				pnjTransformed[currentMap][enemyKey] = true
				if py != newY || px != newX {
					mapData[py][px] = 1
				} else {
					placePlayerNearby(mapData, newX, newY)
				}
				fmt.Println("La créature retrouve sa forme humaine et devient un PNJ amical !")
				showDialogue(currentMap, newX, newY)
			} else {
				// Tous les autres mobs ne se transforment plus jamais
				enemiesDefeated[currentMap][enemyKey] = true
				mapData[py][px] = 0
				mapData[newY][newX] = 1
				fmt.Println("Vous pouvez maintenant passer par cette case.")
			}
		} else {
			fmt.Println("Vous restez à votre position.")
		}
		return false, currentMap

	case 3: // PNJ
		fmt.Println("Vous parlez au PNJ...")
		showDialogue(currentMap, newX, newY)
		return false, currentMap

	case 30: // Porte spéciale vers la salle secrète dans salle1
		if currentMap == "salle1" {
			if playerInventory["clés_spéciales"] > 0 {
				fmt.Println("Vous utilisez votre clé spéciale !")
				fmt.Println("Un passage secret s'ouvre...")
				playerInventory["clés_spéciales"]--
				return true, "salle8"
			} else {
				fmt.Println("Cette porte nécessite une clé spéciale...")
				fmt.Println("Peut-être que le marchand en a une ?")
				return false, currentMap
			}
		}

	case 32: // Retour depuis la salle secrète
		if currentMap == "salle8" {
			if tr, ok := transitions[currentMap][cell]; ok {
				fmt.Printf("Retour vers %s\n", tr.nextMap)
				return true, tr.nextMap
			}
		}

	// Toutes les autres portes - gestion unifiée
	case 7, 10, 13, 14, 15, 20, 21, 27, 28, 31, 33, 34, 38, 40, 42, 44:
		if tr, ok := transitions[currentMap][cell]; ok {
			fmt.Printf("Transition vers %s aux coordonnées (%d,%d)\n", tr.nextMap, tr.spawnX, tr.spawnY)
			return true, tr.nextMap
		} else {
			fmt.Printf("Aucune transition définie pour la case %d dans %s\n", cell, currentMap)
			return false, currentMap
		}
	}

	// Déplacement normal
	if cell == 0 || cell >= 16 {
		mapData[py][px] = 0     // Vide l'ancienne position
		mapData[newY][newX] = 1 // Place le joueur à la nouvelle position
		return false, currentMap
	}

	return false, currentMap
}

func openChest(currentMap string, chestX, chestY int) {
	// Clé unique basée sur la carte et la position du coffre
	chestKey := fmt.Sprintf("%s_%d_%d", currentMap, chestX, chestY)
	if chestOpened[chestKey] {
		fmt.Println("📦 Le coffre est vide... Vous l'avez déjà ouvert.")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Voulez-vous l'ouvrir ? (o/n): ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(strings.ToLower(input))

	if input == "o" || input == "oui" {
		fmt.Println("🔓 *clic* Le coffre s'ouvre...")
		items := []struct {
			name        string
			amount      int
			probability int
		}{
			{"épées", 1, 30},
			{"potions", 2, 25},
			{"pièces", 20, 25},
			{"clés", 2, 15},
			{"pièces", 50, 5}, // Jackpot rare
		}

		roll := rand.Intn(100)
		cumulative := 0
		for _, item := range items {
			cumulative += item.probability
			if roll < cumulative {
				addToInventory(item.name, item.amount)
				chestOpened[chestKey] = true
				fmt.Println("✨ Le coffre se referme magiquement après avoir donné son trésor.")
				break
			}
		}
	} else {
		fmt.Println("🤔 Vous décidez de laisser le coffre fermé pour l'instant.")
	}

	fmt.Print("Appuyez sur Entrée pour continuer...")
	reader.ReadString('\n')
}

// Fonction pour ouvrir un coffre secret avec des objets rares
func openSecretChest(x, y int) {
	// Vérifier si ce coffre spécifique a déjà été ouvert
	chestKey := fmt.Sprintf("secret_%d_%d", x, y)
	if secretChestsOpened[chestKey] {
		fmt.Println("📦 Ce coffre secret est vide... Vous l'avez déjà ouvert.")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("🔮 Vous trouvez un coffre secret orné de symboles anciens!")
	fmt.Print("Voulez-vous l'ouvrir ? (o/n): ")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(strings.ToLower(input))

	if input == "o" || input == "oui" {
		fmt.Println("✨ *Une aura magique enveloppe le coffre qui s'ouvre lentement*")

		// Objets plus rares et plus précieux que les coffres normaux
		items := []struct {
			name        string
			amount      int
			probability int
		}{
			{"épées", 2, 20},
			{"potions", 3, 20},
			{"pièces", 100, 25},
			{"clés", 3, 15},
			{"clés_spéciales", 1, 15},
			{"artefacts", 1, 5},
		}

		roll := rand.Intn(100)
		cumulative := 0

		for _, item := range items {
			cumulative += item.probability
			if roll < cumulative {
				addToInventory(item.name, item.amount)
				secretChestsOpened[chestKey] = true
				fmt.Println("🌟 Le coffre secret se referme dans un halo de lumière après avoir révélé son trésor.")
				break
			}
		}
	} else {
		fmt.Println("🤔 Vous décidez de laisser le coffre secret fermé pour l'instant.")
	}

	fmt.Print("Appuyez sur Entrée pour continuer...")
	reader.ReadString('\n')
}

// Gère l'entrée utilisateur pour le déplacement et actions.
// Retourne les nouvelles coordonnées (ou les mêmes si pas de déplacement) et un bool indiquant si la boucle doit continuer.
func getPlayerMovement(reader *bufio.Reader, px, py int) (int, int, bool) {
	fmt.Print("Déplacez-vous (ZQSD Pour se Déplacer, I=Inventaire, X=Quitter): ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(strings.ToLower(input))

	newX, newY := px, py
	switch input {
	case "z", "haut":
		newY = py - 1
	case "s", "bas":
		newY = py + 1
	case "q", "gauche":
		newX = px - 1
	case "d", "droite":
		newX = px + 1
	case "i", "inv", "inventaire":
		showInventory()
		return px, py, true
	case "x", "quit", "exit":
		fmt.Println("Vous quittez la partie. Merci d'avoir joué !")
		return px, py, false
	default:
		fmt.Println("Entrée invalide.")
		return px, py, true
	}
	return newX, newY, true
}

// Valide un mouvement sur la carte (bords et murs)
func isValidMovement(x, y int, mapData [][]int) bool {
	if y < 0 || y >= len(mapData) || x < 0 || x >= len(mapData[0]) {
		return false
	}
	if mapData[y][x] == 9 { // mur
		return false
	}
	return true
}

// Boucle principale du jeu refactorisée
func RunGameLoop(currentMap string) {
	reader := bufio.NewReader(os.Stdin)
	mapData := copyMap(salles[currentMap])

	// Applique l'état des ennemis vaincus
	applyEnemyStates(mapData, currentMap)

	// Assure que le joueur est présent dans la carte initiale
	if px, py := findPlayer(mapData); px == -1 || py == -1 {
		placePlayerAt(mapData, len(mapData[0])/2, len(mapData)/2)
	}

	for {
		printMap(mapData) // Le HUD est maintenant intégré dans printMap
		fmt.Printf("📍 Salle actuelle: %s\n", currentMap)

		px, py := findPlayer(mapData)
		if px == -1 || py == -1 {
			fmt.Println("Erreur: Joueur non trouvé dans la carte!")
			return
		}

		// Utiliser la nouvelle fonction de mouvement
		newX, newY, shouldContinue := getPlayerMovement(reader, px, py)

		if !shouldContinue {
			return // Le joueur a quitté le jeu
		}

		// Si les coordonnées n'ont pas changé (entrée invalide), continuer la boucle
		if newX == px && newY == py {
			continue
		}

		// Vérifier si le mouvement est valide
		if !isValidMovement(newX, newY, mapData) {
			continue
		}

		cell := mapData[newY][newX]

		// Gérer l'interaction avec la case
		transitionNeeded, newMap := handleCellInteraction(cell, currentMap, newX, newY, mapData, px, py)

		// Préparer les coordonnées d'apparition de la prochaine carte si transition
		spawnX, spawnY := -1, -1
		haveSpawn := false
		if transitionNeeded {
			if tr, ok := transitions[currentMap][cell]; ok {
				spawnX, spawnY = tr.spawnX, tr.spawnY
				haveSpawn = true
			}
		}

		if transitionNeeded {
			currentMap = newMap
			mapData = copyMap(salles[currentMap])
			applyEnemyStates(mapData, currentMap)

			// Placer le joueur selon la transition
			if currentMap == "salle8" {
				placePlayerAt(mapData, 3, 6) // Position spéciale pour la salle secrète
			} else if haveSpawn {
				placePlayerAt(mapData, spawnX, spawnY)
			} else {
				// Position par défaut
				placePlayerAt(mapData, len(mapData[0])/2, len(mapData)/2)
			}
		}
	}
}
