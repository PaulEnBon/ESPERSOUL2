package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/eiannone/keyboard"
)

// ---- Récompenses configurables ----
// Peut être déplacé dans un fichier de config plus tard.
const (
	baseMinCoins         = 5    // ancien 3
	baseMaxCoins         = 9    // génère range baseMin..baseMax inclus
	legendaryWeaponBonus = 0.50 // +50% si arme légendaire
	puffAttackBonusShare = 0.20 // 20% de l'attackBoost converti en bonus or (ex: 30% atk -> +6% or)
	jackpotChancePercent = 6    // % de chance de jackpot
	jackpotMultiplier    = 4.0  // x4 sur le total final
)

// calcule le loot de pièces enrichi
func computeCoinLoot() (coins int, jackpot bool, breakdown string) {
	// Base aléatoire
	base := baseMinCoins + rand.Intn(baseMaxCoins-baseMinCoins+1)

	// Bonus arme légendaire
	legendaryBonus := 0
	if playerStats.hasLegendaryWeapon {
		legendaryBonus = int(float64(base) * legendaryWeaponBonus)
	}

	// Bonus Puff converti (attackBoost est un pourcentage cumulatif)
	puffBonus := int(float64(base) * (float64(playerStats.attackBoost) / 100.0) * puffAttackBonusShare)

	total := base + legendaryBonus + puffBonus

	// Jackpot ?
	if rand.Intn(100) < jackpotChancePercent {
		total = int(float64(total) * jackpotMultiplier)
		return total, true, fmt.Sprintf("base=%d +legend=%d +puff=%d xJackpot(%.1fx)", base, legendaryBonus, puffBonus, jackpotMultiplier)
	}

	return total, false, fmt.Sprintf("base=%d +legend=%d +puff=%d", base, legendaryBonus, puffBonus)
}

// ——————————————————————————————————————————————————————————
// Système de combat intégrant classes/armes/dégâts/effets/artefacts
// ——————————————————————————————————————————————————————————

// fabrique un personnage joueur de base (en attendant la vraie sélection)
func buildPlayerCharacter() Personnage {
	// Part d'une copie du joueur persistant
	p := currentPlayer
	// Applique l'armure et l'arme sur la copie (pas sur l'état persistant)
	_ = EquiperArmure(&p, p.ArmuresDisponibles)
	if p.NiveauArme >= 0 && p.NiveauArme < len(p.ArmesDisponibles) {
		_ = EquiperArme(&p, p.ArmesDisponibles[p.NiveauArme])
	}
	// Préserve les PV persistants et les borne au nouveau PVMax
	if currentPlayer.PV > 0 {
		if currentPlayer.PV > p.PVMax {
			p.PV = p.PVMax
		} else {
			p.PV = currentPlayer.PV
		}
	} else {
		// Si PV persistants à 0, démarre à 0 (pas de heal auto)
		if p.PV > p.PVMax {
			p.PV = p.PVMax
		}
	}
	return p
}

// fabrique un ennemi générique en fonction du niveau de menace
func buildEnemy(isSuper bool) Personnage {
	base := Personnage{
		Nom:                "Créature",
		PV:                 8,
		PVMax:              8,
		Armure:             10,
		ResistMag:          8,
		Precision:          0.80,
		TauxCritique:       0.10,
		MultiplicateurCrit: 1.5,
	}
	// Arme simple pour l'ennemi
	arme := epeePierre
	_ = EquiperArme(&base, arme)
	return base
}

// réalise une attaque avec calcul précision/crit/type en utilisant les helpers de degats.go
func resolveAttack(attaquant, defenseur *Personnage, degatsBase int, typeAttaque string) (degats int, touche bool, crit bool) {
	d, estCrit, aTouche := CalculerDegatsAvecCritique(attaquant, defenseur, degatsBase, typeAttaque)
	return d, aTouche, estCrit
}

// choisit une compétence « simple » (priorité aux dégâts > 0) sur l'arme équipée
func pickCompetence(p *Personnage) (Competence, bool) {
	if len(p.ArmeEquipee.Competences) == 0 {
		return Competence{}, false
	}
	// Choisit la première compétence avec dégâts positifs, sinon la première dispo
	for _, c := range p.ArmeEquipee.Competences {
		if c.Degats > 0 {
			return c, true
		}
	}
	return p.ArmeEquipee.Competences[0], true
}

// applique un effet éventuel sur la cible en fonction de la compétence
func maybeApplyEffect(defenseur *Personnage, comp Competence) {
	if comp.TypeEffet == "" {
		return
	}
	if eff := CreerEffet(comp.TypeEffet, comp.Puissance); eff != nil {
		AppliquerEffet(defenseur, *eff)
	}
}

// détermine si l'effet est plutôt un buff sur soi que sur la cible
func isSelfBuff(effectName string) bool {
	switch effectName {
	case "Augmentation de Dégâts", "Augmentation de Dégâts Magiques", "Régénération", "Guérison Poison":
		return true
	default:
		return false
	}
}

// propose la liste des compétences de l'arme et retourne le choix de l'utilisateur
func chooseCompetence(p *Personnage) (Competence, bool) {
	comps := p.ArmeEquipee.Competences
	if len(comps) == 0 {
		return Competence{}, false
	}
	fmt.Println("\nChoisissez une compétence:")
	for i, c := range comps {
		eff := c.TypeEffet
		extra := ""
		if eff != "" {
			extra = fmt.Sprintf(" | Effet: %s (puiss.%d)", eff, c.Puissance)
		}
		fmt.Printf("  %d) %s [%s] Dégâts:%d%s\n", i+1, c.Nom, c.Type, c.Degats, extra)
	}
	fmt.Print("Votre choix (1-", len(comps), "): ")
	// Lire une seule touche depuis le canal global
	if globalKeyEvents == nil {
		// Fallback extrême si le canal n'est pas prêt
		return comps[0], true
	}
	e := <-globalKeyEvents
	r := e.Rune
	if r >= '1' && r <= '9' {
		idx := int(r - '0')
		if idx >= 1 && idx <= len(comps) {
			return comps[idx-1], true
		}
	}
	fmt.Println("Saisie invalide, compétence par défaut utilisée.")
	return comps[0], true
}

// Système de combat amélioré avec les modules existants
func combat(currentMap string, isSuper bool) interface{} {
	rand.Seed(time.Now().UnixNano())

	// Crée les entités combat
	player := buildPlayerCharacter()
	enemy := CreateRandomEnemyForMap(currentMap, isSuper)
	enemyAttackBase := enemy.ArmeEquipee.DegatsPhysiques
	if enemyAttackBase <= 0 {
		enemyAttackBase = 12
	}

	fmt.Println("\n🗡️  COMBAT ENGAGÉ ! 🗡️")
	if isSuper {
		fmt.Println("Vous affrontez un ENNEMI SURPUISSANT !")
	} else {
		fmt.Println("Vous affrontez une créature maudite !")
	}

	for player.PV > 0 && enemy.PV > 0 {
		fmt.Printf("\n💚 Vos PV: %d/%d | 💀 PV Ennemi: %d/%d\n", player.PV, player.PVMax, enemy.PV, enemy.PVMax)
		if playerStats.hasLegendaryWeapon {
			fmt.Println("🌟 Excalibur équipée (+50% dégâts de loot)")
		}

		// Affiche compétence de base
		if comp, ok := pickCompetence(&player); ok {
			fmt.Printf("Compétence de base: %s (%s, %d dégâts)\n", comp.Nom, comp.Type, comp.Degats)
		}

		fmt.Println("Actions: [A]ttaquer, [D]éfendre, [P]otion, [U]ser Puff 9K, [F]uir")
		fmt.Print("Choisissez une action: ")
		// Utilise le même canal que la boucle de jeu
		if globalKeyEvents == nil {
			fmt.Println("(clavier non initialisé)")
			return false
		}
		e := <-globalKeyEvents
		input := strings.ToLower(string(e.Rune))
		if e.Key == keyboard.KeyEsc {
			input = "f"
		}

		// Début de tour: protections d'artefacts éventuelles
		AppliquerProtectionsArtefactsDebutTour(&player)

		switch input {
		case "a":
			// Sélection de compétence
			comp, ok := chooseCompetence(&player)
			if !ok {
				// Fallback absolument minimal
				comp = Competence{Nom: "Attaque", Degats: 15, Type: "physique"}
			}

			// Buffs/soins auto-ciblés
			if comp.Degats <= 0 && comp.TypeEffet != "" && isSelfBuff(comp.TypeEffet) {
				if eff := CreerEffet(comp.TypeEffet, comp.Puissance); eff != nil {
					AppliquerEffet(&player, *eff)
					fmt.Printf("✨ Vous utilisez %s sur vous-même.\n", comp.Nom)
				}
			} else if comp.Degats <= 0 && comp.TypeEffet != "" {
				// Utilitaires offensifs sans dégâts (débuffs)
				if eff := CreerEffet(comp.TypeEffet, comp.Puissance); eff != nil {
					AppliquerEffet(&enemy, *eff)
					fmt.Printf("✨ Vous appliquez %s à l'ennemi.\n", comp.Nom)
				}
			} else {
				// Attaque avec dégâts
				degatsBase := comp.Degats
				typeAtk := comp.Type
				if degatsBase <= 0 {
					degatsBase = 15
				}
				if typeAtk == "" {
					typeAtk = "physique"
				}

				dmg, touche, crit := resolveAttack(&player, &enemy, degatsBase, typeAtk)
				if !touche {
					fmt.Println("🙃 Votre attaque rate sa cible !")
				} else {
					enemy.PV -= dmg
					if enemy.PV < 0 {
						enemy.PV = 0
					}
					if crit {
						fmt.Printf("⚔️  Coup critique ! Vous infligez %d dégâts.\n", dmg)
					} else {
						fmt.Printf("⚔️  Vous infligez %d dégâts.\n", dmg)
					}
					if comp.TypeEffet != "" {
						maybeApplyEffect(&enemy, comp)
					}
				}
			}

		case "d":
			fmt.Println("🛡️  Vous vous défendez !")
			// La défense réduit les dégâts du prochain coup ennemi
			// On simule un bouclier temporaire stocké dans un effet léger
			shield := Effet{Nom: "Bouclier", ToursRestants: 1, ModifArmure: 0.30, ChanceAppliquer: 1.0}
			AppliquerEffet(&player, shield)
			// Pas d'attaque joueur ce tour, on enchaîne vers l'ennemi

		case "p":
			if playerInventory["potions"] > 0 {
				heal := 70
				player.PV += heal
				if player.PV > player.PVMax {
					player.PV = player.PVMax
				}
				playerInventory["potions"]--
				fmt.Printf("🧪 Vous vous soignez de %d PV ! (PV actuels: %d/%d)\n", heal, player.PV, player.PVMax)
			} else {
				fmt.Println("❌ Vous n'avez pas de potions !")
				// saute le tour ennemi si pas d'action ? non, on continue le tour normal
			}

		case "u":
			if playerInventory["puff_9k"] > 0 {
				playerInventory["puff_9k"]--
				playerStats.attackBoost += 15 // conserve l'ancien bonus pour le loot
				// Ajoute un effet d'Augmentation de Dégâts sur le joueur
				if eff := CreerEffet("Augmentation de Dégâts", 2); eff != nil {
					AppliquerEffet(&player, *eff)
				}
				player.PV -= 5
				if player.PV < 0 {
					player.PV = 0
				}
				fmt.Println("💊 Vous utilisez un Puff 9K !")
				fmt.Println("⚡ +15% de dégâts (loot) et buff de dégâts temporaire !")
				fmt.Printf("💔 Vous perdez 5 PV. PV actuels: %d/%d\n", player.PV, player.PVMax)

				if player.PV <= 0 {
					fmt.Println("💀 Le Puff 9K vous a tué ! Attention à la surdose...")
					return false
				}
			} else {
				fmt.Println("❌ Vous n'avez pas de Puff 9K !")
			}

		case "f":
			fmt.Println("💨 Vous fuyez le combat !")
			// Persiste les PV du joueur
			currentPlayer.PV = player.PV
			// Reset les bonus temporaires côté anciens stats loot
			playerStats.attackBoost = 0
			return false

		default:
			fmt.Println("Action invalide !")
			// on passe quand même au tour adverse, comme avant
		}

		// Fin d'action joueur: traitements d'effets sur les deux
		TraiterEffetsFinTour(&player)
		TraiterEffetsFinTour(&enemy)

		if enemy.PV <= 0 {
			fmt.Println("\n🎉 VICTOIRE ! Vous avez vaincu la créature !")

			// Incrémente le compteur d'ennemis tués (stat héritée)
			playerStats.enemiesKilled++

			coins, jackpot, details := computeCoinLoot()
			addToInventory("pièces", coins)
			if jackpot {
				fmt.Printf("💎 JACKPOT ! Vous obtenez %d pièces (%s) !\n", coins, details)
			} else {
				fmt.Printf("✨ Vous avez reçu %d pièces (%s).\n", coins, details)
			}

			// Drop de roches d'évolution selon la difficulté
			tier := tierForMap(currentMap)
			rocks := 0
			switch tier {
			case TierTutorial, TierEarly:
				rocks = 1 // easy
			case TierMid:
				rocks = 2 // moyen
			case TierLate:
				rocks = 3 // hard
			}
			if rocks > 0 {
				currentPlayer.Roches += rocks
				fmt.Printf("🪨 Vous obtenez %d roche(s) d'évolution. Total roches: %d\n", rocks, currentPlayer.Roches)
			}

			// Persiste les PV du joueur
			currentPlayer.PV = player.PV
			// Reset bonus temporaires hérités
			playerStats.attackBoost = 0

			// Tous les ennemis disparaissent (PNJ géré côté loop pour le cas spécial)
			fmt.Println("💨 La créature disparaît complètement dans un nuage de fumée...")
			return "disappear"
		}

		// Tour de l'ennemi — saute si étourdi
		if EstEtourdi(&enemy) {
			fmt.Println("😵‍💫 L'ennemi est étourdi et rate son tour !")
		} else {
			// L'ennemi tente une compétence simple si disponible
			ecomp, ok := pickCompetence(&enemy)
			edeg := enemyAttackBase
			etype := "physique"
			if ok {
				if ecomp.Degats > 0 {
					edeg = ecomp.Degats
				}
				if ecomp.Type != "" {
					etype = ecomp.Type
				}
			}
			edmg, touche, crit := resolveAttack(&enemy, &player, edeg, etype)
			if !touche {
				fmt.Println("🌀 L'ennemi rate son attaque !")
			} else {
				player.PV -= edmg
				if player.PV < 0 {
					player.PV = 0
				}
				if crit {
					fmt.Printf("💥 Coup critique ennemi ! Vous subissez %d dégâts.\n", edmg)
				} else {
					fmt.Printf("💥 L'ennemi vous inflige %d dégâts.\n", edmg)
				}
				if ok {
					maybeApplyEffect(&player, ecomp)
				}
			}
		}

		// Fin de tour: ticks d'effets
		TraiterEffetsFinTour(&player)
		TraiterEffetsFinTour(&enemy)

		if player.PV <= 0 {
			fmt.Println("\n💀 DÉFAITE ! Vous avez été vaincu...")
			fmt.Println("🔄 Vous retournez au début de la salle.")
			// Persiste les PV du joueur (reste à 0)
			currentPlayer.PV = player.PV
			playerStats.attackBoost = 0
			return false
		}
	}

	return false
}
