package main

import (
	"fmt"
	"os"
)

// Dialogues des PNJ par salle et position avec récompenses
var npcDialogues = map[string]map[string]struct {
	dialogues []string
	reward    string
	amount    int
}{
	"salle1": {
		"8_3": {
			dialogues: []string{
				"Merci de m'avoir libéré de cette malédiction !",
				"Grâce à toi, je peux enfin retrouver ma forme humaine.",
				"Maintenant, je peux t'aider dans ta quête.",
				"Tu vas devoir affronter de nombreux dangers dans ce donjon.",
				"Utilise cette clé spéciale pour ouvrir des portes verrouillées.",
				"Elles te dirigeront vers des trésors cachés.",
				"Prends cette clé, elle t'aidera dans ta quête.",
				"Les autres créatures dans ce donjon ont subi le même sort que moi.",
				"Ne les sous-estime pas, elles sont redoutables.",
			},
			reward: "clés",
			amount: 1,
		},
	},
	"salle3": {
		"8_3": {
			dialogues: []string{
				"Tu m'as sauvé ! Cette malédiction était terrible...",
				"Je gardais ce trésor depuis des années.",
				"Voici une potion de soin, tu en auras besoin !",
				"Attention, le boss final se trouve plus loin dans le donjon.",
			},
			reward: "potions",
			amount: 1,
		},
	},
	"salle4": {
		"1_3": {
			dialogues: []string{
				"Bienvenue, brave aventurier !",
				"Es-ce ce bon vieux Vitaly qui t'envoie me voir ?",
				"Je suis le marchand de ce donjon maudit.",
				"J'ai survécu ici en échangeant des objets avec les voyageurs.",
				"Que puis-je faire pour toi ?",
			},
			reward: "pièces",
			amount: 1,
		},
	},
	"salle5": {
		"2_2": {
			dialogues: []string{
				"Salut aventurier ! Je suis le forgeron de ce donjon.",
				"J'ai passé des années à perfectionner mon art ici.",
				"Si tu me donnes 15 pièces, je peux te forger une épée !",
				"Que dis-tu de cette offre ?",
			},
			reward: "",
			amount: 0,
		},
	},
	"salle7": {
		"2_2": {
			dialogues: []string{
				"Salut mec ! Bienvenue dans mon casino souterrain !",
				"Tu veux tenter ta chance ? J'ai des caisses mystères !",
				"Certaines sont cheap mais avec des trucs de ouf dedans...",
				"D'autres sont chères mais garantissent des armes légendaires !",
				"Alors, tu veux jouer ?",
			},
			reward: "",
			amount: 0,
		},
	},
	"salle8": {
		"3_3": {
			dialogues: []string{
				"Bienvenue dans la salle des trésors secrets !",
				"Seuls les plus braves aventuriers arrivent ici...",
				"Ces coffres contiennent des récompenses exceptionnelles !",
				"Que la chance soit avec toi, noble héros !",
			},
			reward: "pièces",
			amount: 10,
		},
		"4_3": {
			dialogues: []string{
				"Tu as découvert notre sanctuaire secret !",
				"Ces trésors étaient cachés depuis des siècles...",
				"Prends cette épée légendaire !",
				"Elle t'aidera dans tes futures aventures !",
			},
			reward: "épées",
			amount: 2,
		},
	},
	"salle11": {
		"3_2": { // PNJ soigneur au centre
			dialogues: []string{
				"Bienvenue au sanctuaire de repos.",
				"Je peux te soigner complètement pour 10 pièces.",
				"Appuie sur O pour accepter, N pour refuser.",
			},
			reward: "",
			amount: 0,
		},
	},
}

// Système de dialogue avec les PNJ
func showDialogue(currentMap string, x, y int) {
	key := fmt.Sprintf("%d_%d", x, y)

	// Helper: read a single key (last of any burst), returns lowercase rune
	readKey := func() rune {
		if globalKeyEvents == nil {
			// Fallback safety: no keyboard channel available
			return 0
		}
		e := <-globalKeyEvents
		draining := true
		for draining {
			select {
			case next := <-globalKeyEvents:
				e = next
			default:
				draining = false
			}
		}
		r := e.Rune
		// normalize to lowercase ASCII if applicable
		if r >= 'A' && r <= 'Z' {
			r = r + 32
		}
		return r
	}

	// Cas spécial pour le marchand de salle4
	if currentMap == "salle4" && key == "1_3" {
		npcData := npcDialogues[currentMap][key]
		fmt.Println("\n💬 === DIALOGUE ===")
		for i, line := range npcData.dialogues {
			fmt.Printf("👨 PNJ: %s\n", line)

			// Question oui/non après la phrase sur Vitaly
			if line == "Es-ce ce bon vieux Vitaly qui t'envoie me voir ?" {
				// S'assurer que la map interne existe
				if rewardsGiven[currentMap] == nil {
					rewardsGiven[currentMap] = make(map[string]bool)
				}
				// Ne donner la récompense qu'une seule fois
				if !rewardsGiven[currentMap][key] {
					fmt.Print("👨 PNJ: Est-ce bien cela ? (o/n): ")
					ans := readKey()
					if ans == 'o' {
						addToInventory("potions", 1)
						rewardsGiven[currentMap][key] = true
						fmt.Println("🎁 Vous recevez 1 potion.")
					} else {
						fmt.Println("👨 PNJ: Très bien, poursuivons…")
					}
				}
			}
			if i < len(npcData.dialogues)-1 {
				fmt.Print("Appuyez sur une touche pour continuer...")
				_ = readKey()
			}
		}
		fmt.Print("Appuyez sur une touche pour ouvrir le magasin...")
		_ = readKey()
		fmt.Println("===================")

		showMerchantInterface()
		return
	}

	// Cas spécial pour le forgeron de salle5
	if currentMap == "salle5" && key == "2_2" {
		npcData := npcDialogues[currentMap][key]
		fmt.Println("\n💬 === DIALOGUE ===")
		for i, line := range npcData.dialogues {
			fmt.Printf("🔨 Forgeron: %s\n", line)
			if i < len(npcData.dialogues)-1 {
				fmt.Print("Appuyez sur une touche pour continuer...")
				_ = readKey()
			}
		}
		fmt.Print("Appuyez sur une touche pour ouvrir la forge...")
		_ = readKey()
		fmt.Println("===================")

		showForgeInterface()
		return
	}

	// Cas spécial pour le gambling de salle7
	if currentMap == "salle7" && key == "2_2" {
		npcData := npcDialogues[currentMap][key]
		fmt.Println("\n💬 === DIALOGUE ===")
		for i, line := range npcData.dialogues {
			fmt.Printf("🎰 Croupier: %s\n", line)
			if i < len(npcData.dialogues)-1 {
				fmt.Print("Appuyez sur une touche pour continuer...")
				_ = readKey()
			}
		}
		fmt.Print("Appuyez sur une touche pour ouvrir le casino...")
		_ = readKey()
		fmt.Println("===================")

		showGamblingInterface()
		return
	}

	// Fin de jeu spéciale: salle1 PNJ (8,3) si item 'sida'
	if currentMap == "salle1" && key == "8_3" && playerInventory["sida"] > 0 {
		fmt.Println("\n💬 === DIALOGUE ===")
		fmt.Println("👨 PNJ: ...Cet objet... D'où le tiens-tu ?")
		fmt.Println("👨 PNJ: Si tu me le donnes, je pourrai briser définitivement la malédiction.")
		fmt.Println("👨 PNJ: Veux-tu ME DONNER 'sida' et mettre fin au jeu ? (o/n)")
		fmt.Print("Votre choix: ")
		ans := readKey()
		if ans == 'o' {
			playerInventory["sida"]--
			fmt.Println("✨ Le PNJ absorbe l'énergie sombre...")
			fmt.Println("🏁 La malédiction est levée. Vous avez terminé le jeu !")
			fmt.Println("Merci d'avoir joué.")
			os.Exit(0)
		} else {
			fmt.Println("👨 PNJ: Très bien... Peut-être plus tard.")
			fmt.Print("Appuyez sur une touche pour continuer...")
			_ = readKey()
			return
		}
	}

	// Dialogue normal pour les autres PNJ
	npcData, exists := npcDialogues[currentMap][key]
	if !exists {
		// Cas fallback: si c'est le mentor transformé mais dialogues non trouvés
		if currentMap == "salle1" && key == "8_3" {
			fmt.Println("🧙 Mentor Suprême: Merci de m'avoir libéré !")
		} else {
			fmt.Println("👨 PNJ: Merci de m'avoir libéré !")
		}
		fmt.Print("Appuyez sur une touche pour continuer...")
		_ = readKey()
		return
	}

	// Cas spécial: soigneur de salle11 (3,2)
	if currentMap == "salle11" && key == "3_2" {
		fmt.Println("\n💬 === DIALOGUE ===")
		for i, line := range npcData.dialogues {
			fmt.Printf("🧙 Soigneur: %s\n", line)
			if i < len(npcData.dialogues)-1 {
				fmt.Print("Appuyez sur une touche pour continuer...")
				_ = readKey()
			}
		}
		fmt.Print("🧙 Soigneur: Souhaitez-vous être soigné pour 10 pièces ? (o/n): ")
		ans := readKey()
		if ans == 'o' {
			if playerInventory["pièces"] >= 10 {
				playerInventory["pièces"] -= 10
				// Heal complet en fonction de l'armure équipée actuelle
				tmp := currentPlayer
				_ = EquiperArmure(&tmp, tmp.ArmuresDisponibles)
				currentPlayer.PV = tmp.PVMax
				fmt.Println("✨ Vous êtes complètement soigné !")
			} else {
				fmt.Println("🚫 Vous n'avez pas assez de pièces.")
			}
		} else {
			fmt.Println("Très bien, revenez si besoin.")
		}
		fmt.Print("Appuyez sur une touche pour fermer...")
		_ = readKey()
		fmt.Println("===================")
		return
	}

	fmt.Println("\n💬 === DIALOGUE ===")
	for i, line := range npcData.dialogues {
		if currentMap == "salle1" && key == "8_3" {
			fmt.Printf("🧙 Mentor Suprême: %s\n", line)
		} else {
			fmt.Printf("👨 PNJ: %s\n", line)
		}
		if i < len(npcData.dialogues)-1 {
			fmt.Print("Appuyez sur une touche pour continuer...")
			_ = readKey()
		}
	}

	// Vérifier si la récompense a déjà été donnée
	if rewardsGiven[currentMap][key] {
		if currentMap == "salle1" && key == "8_3" {
			fmt.Printf("🧙 Mentor Suprême: Je t'ai déjà remis ma récompense, mais merci encore !\n")
		} else {
			fmt.Printf("👨 PNJ: Je t'ai déjà donné ma récompense, mais merci encore!\n")
		}
	} else {
		// Donner la récompense une seule fois
		if npcData.reward != "" && npcData.amount > 0 {
			addToInventory(npcData.reward, npcData.amount)
			rewardsGiven[currentMap][key] = true
		}
	}

	fmt.Print("Appuyez sur une touche pour fermer...")
	_ = readKey()
	fmt.Println("===================")
}
