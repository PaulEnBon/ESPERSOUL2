package main

import (
	"fmt"
	"math/rand"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/eiannone/keyboard"
)

// Controls whether enemies use the synthesized 4-skill kit (true for all normal monsters)
var combatForceSynthKit = true

// Tracks the last enemy skill name used, for HUD display in header
var combatLastEnemySkill string

// --------- Simple UI helpers (ASCII/emoji) for a Pokémon-like battle window ---------

func hpBar(current, max, width int) string {
	if max <= 0 {
		max = 1
	}
	if current < 0 {
		current = 0
	}
	if current > max {
		current = max
	}
	filled := int(float64(current) / float64(max) * float64(width))
	if filled < 0 {
		filled = 0
	}
	if filled > width {
		filled = width
	}
	bar := strings.Repeat("█", filled) + strings.Repeat("░", width-filled)
	// Color hint via emojis and counts
	return bar
}

// Computes effective armor after modifiers for display purposes
func effectiveArmor(p *Personnage) int {
	if p == nil {
		return 0
	}
	_, _, _, _, modifArmure, _ := CalculerModificateurs(p)
	if modifArmure < 0.1 {
		modifArmure = 0.1
	}
	eff := int(float64(p.Armure) * modifArmure)
	if eff < 0 {
		eff = 0
	}
	return eff
}

// Armor-aware HP bar: increases visual fill based on damage mitigation from armor
func hpBarWithArmor(current, max, armor, width int) string {
	if max <= 0 {
		max = 1
	}
	if current < 0 {
		current = 0
	}
	if current > max {
		current = max
	}
	// Reduction model from degats.go: damageTaken = base * 100/(100+armor)
	// So effective HP factor = (100+armor)/100
	factor := 1.0 + (float64(armor) / 100.0)
	scaledRatio := (float64(current) / float64(max)) * factor
	if scaledRatio > 1.0 {
		scaledRatio = 1.0
	}
	if scaledRatio < 0.0 {
		scaledRatio = 0.0
	}
	filled := int(scaledRatio * float64(width))
	if filled < 0 {
		filled = 0
	}
	if filled > width {
		filled = width
	}
	return strings.Repeat("█", filled) + strings.Repeat("░", width-filled)
}

// ANSI helpers for dynamic coloring
func hpRatio(cur, max int) float64 {
	if max <= 0 {
		return 0
	}
	if cur < 0 {
		cur = 0
	}
	if cur > max {
		cur = max
	}
	return float64(cur) / float64(max)
}

func ansiColorForHP(cur, max int) string {
	r := hpRatio(cur, max)
	// Red <20%, Yellow <50%, Green otherwise
	if r <= 0.20 {
		return "\033[31m" // red
	}
	if r <= 0.50 {
		return "\033[33m" // yellow
	}
	return "\033[32m" // green
}

func ansiReset() string { return "\033[0m" }

func colorize(s, color string) string {
	if color == "" {
		return s
	}
	return color + s + ansiReset()
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func printBattleHeader(player, enemy Personnage, isSuper bool) {
	clearScreen()
	eIcon := emojiForEnemyName(enemy.Nom)
	title := "⚔️ COMBAT"
	if isSuper {
		title = "⚔️ COMBAT — ENNEMI SURPUISSANT"
	}
	// Determine frame color based on HP status (priority: player critical -> red; enemy critical -> green; player mid -> yellow)
	pR := hpRatio(player.PV, player.PVMax)
	eR := hpRatio(enemy.PV, enemy.PVMax)
	frameColor := "\033[36m" // cyan default
	if pR <= 0.20 {
		frameColor = "\033[31m" // red: danger
	} else if eR <= 0.20 {
		frameColor = "\033[32m" // green: avantage
	} else if pR <= 0.50 {
		frameColor = "\033[33m" // yellow: attention
	}

	// Prepare names (visible trimming)
	ename := fmt.Sprintf("%s %s", eIcon, enemy.Nom)
	if len(ename) > 28 {
		ename = ename[:28]
	}
	pname := currentPlayer.Nom
	if pname == "" {
		pname = "Joueur"
	}
	if len(pname) > 28 {
		pname = pname[:28]
	}

	// Compute visible lengths using raw (non-colored) bar for sizing
	rawBar := strings.Repeat("█", 24)
	rawEnemyLine := fmt.Sprintf(" Ennemi: %-28s  PV: [%-24s] %d/%d ", ename, rawBar, enemy.PV, enemy.PVMax)
	rawPlayerLine := fmt.Sprintf(" Joueur: %-28s  PV: [%-24s] %d/%d ", pname, rawBar, player.PV, player.PVMax)
	actionsLine := "  ⚔️  Attaque (A)   🎒 Objets (O)   🏃 Fuir (F)               "
	// Base inner width and dynamic expansion to fit longer PV values
	innerWidth := 64
	if l := len(title); l > innerWidth {
		innerWidth = l
	}
	if l := len(rawEnemyLine); l > innerWidth {
		innerWidth = l
	}
	if l := len(rawPlayerLine); l > innerWidth {
		innerWidth = l
	}
	if l := len(actionsLine); l > innerWidth {
		innerWidth = l
	}

	// Top frame (colored)
	fmt.Println(colorize("╔"+strings.Repeat("═", innerWidth)+"╗", frameColor))
	// Title row padded
	fmt.Println(colorize("║ "+fmt.Sprintf("%-*s", innerWidth-2, title)+" ║", frameColor))
	fmt.Println(colorize("╠"+strings.Repeat("═", innerWidth)+"╣", frameColor))
	// Enemy panel (top right style like Pokémon)
	ebar := hpBar(enemy.PV, enemy.PVMax, 24)
	ebarColored := colorize(ebar, ansiColorForHP(enemy.PV, enemy.PVMax))
	enemyContent := fmt.Sprintf(" Ennemi: %-28s  PV: [%-24s] %d/%d ", ename, ebarColored, enemy.PV, enemy.PVMax)
	fmt.Printf("║ %-*s ║\n", innerWidth, enemyContent)
	// Spacer row (or enemy action HUD if available)
	if combatLastEnemySkill != "" {
		// Keep it short and unobtrusive under the enemy panel
		line := "🧠 Action ennemie: " + combatLastEnemySkill
		if len(line) > innerWidth {
			line = line[:innerWidth]
		}
		fmt.Printf("║ %-*s ║\n", innerWidth, line)
	} else {
		fmt.Printf("║ %-*s ║\n", innerWidth, "")
	}
	// Player panel (bottom left style)
	pbar := hpBar(player.PV, player.PVMax, 24)
	pbarColored := colorize(pbar, ansiColorForHP(player.PV, player.PVMax))
	playerContent := fmt.Sprintf(" Joueur: %-28s  PV: [%-24s] %d/%d ", pname, pbarColored, player.PV, player.PVMax)
	fmt.Printf("║ %-*s ║\n", innerWidth, playerContent)
	fmt.Println(colorize("╠"+strings.Repeat("═", innerWidth)+"╣", frameColor))
	// Actions row with icons (padded)
	fmt.Printf("║ %-*s ║\n", innerWidth, actionsLine)
	fmt.Println(colorize("╚"+strings.Repeat("═", innerWidth)+"╝", frameColor))
}

// Emoji par classe d'ennemi (affiché dans l'intro du combat)
func emojiForEnemyName(name string) string {
	// Normaliser: enlever préfixes d'emojis / symboles et espaces
	clean := strings.TrimSpace(name)
	for len(clean) > 0 {
		r, size := utf8.DecodeRuneInString(clean)
		// On considère qu'un nom commence au premier caractère lettre ou chiffre
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			break
		}
		clean = strings.TrimSpace(clean[size:])
	}
	// Certains noms peuvent être composés avec indicateurs supplémentaires; on fait une comparaison directe
	switch clean {
	case "Rat":
		return "🐀"
	case "Gelée":
		return "🟢"
	case "Brigand":
		return "🗡️"
	case "Archer":
		return "🏹"
	case "Apprenti Pyro":
		return "🔥"
	case "Chevalier":
		return "🛡️"
	case "Berserker":
		return "⚔️"
	case "Mage Sombre":
		return "🪄"
	case "Seigneur Démon":
		return "👿"
	case "Archimage":
		return "📜"
	case "Champion déchu":
		return "🥷"
	case "Mentor Maudit":
		return "🧙"
	case "Mentor Suprême":
		return "🎓"
	}
	// Fallback heuristique: si contient "Mentor" on force l'emoji mage
	if strings.Contains(clean, "Mentor") {
		return "🧙"
	}
	return "👾"
}

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

// centralise l'affichage de la récompense en pièces pour éviter les doublons
func printCoinReward(coins int, jackpot bool) {
	if jackpot {
		fmt.Printf("💎 JACKPOT ! Vous obtenez %d pièces !\n", coins)
	} else {
		fmt.Printf("✨ Vous avez reçu %d pièces.\n", coins)
	}
}

// (emojiForEnemyName restauré ici)

// ——————————————————————————————————————————————————————————
// Système de combat intégrant classes/armes/dégâts/effets/artefacts
// ——————————————————————————————————————————————————————————

// fabrique un personnage joueur de base (en attendant la vraie sélection)
func buildPlayerCharacter() Personnage {
	// Part d'une copie du joueur persistant
	p := currentPlayer
	// Recalcule proprement les stats depuis la base et ré-applique l'équipement sauvegardé
	RecomputeFromBaseAndEquip(&p)
	// Préserve les PV persistants et les borne au nouveau PVMax
	if currentPlayer.PV > 0 {
		if currentPlayer.PV > p.PVMax {
			p.PV = p.PVMax
		} else {
			p.PV = currentPlayer.PV
		}
	} else {
		if p.PV > p.PVMax {
			p.PV = p.PVMax
		}
	}
	return p
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

// Sélection aléatoire d'une compétence pour l'IA ennemie
//   - 70%: privilégie une compétence avec dégâts (>0) si disponible
//   - 30%: choix totalement aléatoire (utilitaire/buff compris)
func pickRandomCompetence(p *Personnage) (Competence, bool) {
	// When forcing the synthesized kit, ignore weapon competences for enemies
	if !combatForceSynthKit {
		// Use weapon competences if available
		comps := p.ArmeEquipee.Competences
		if len(comps) > 0 {
			offensives := make([]Competence, 0, len(comps))
			for _, c := range comps {
				if c.Degats > 0 {
					offensives = append(offensives, c)
				}
			}
			if len(offensives) > 0 && rand.Intn(100) < 70 {
				return offensives[rand.Intn(len(offensives))], true
			}
			return comps[rand.Intn(len(comps))], true
		}
		// If no weapon competences, fall back to synthesized kit below
	}

	// Synthesized 4-skill kit for generic monsters (forced or fallback)
	r := rand.Intn(100)
	switch {
	case r < 15:
		// 15% Reinforcement: 50% damage reduction for 2 turns
		return Competence{Nom: "Renforcement", Description: "Réduit les dégâts subis de 50% pendant 2 tours", Degats: 0, Type: "buff", TypeEffet: "Renforcement", Puissance: 1}, true
	case r < 65:
		// Next 50%: Main moderate attack scaled vs player later; use placeholder damage here
		// Damage will be computed from enemyAttackBase in combat loop; just mark type
		return Competence{Nom: "Frappe Mesurée", Description: "Attaque principale modérée", Degats: 0, Type: "physique"}, true
	case r < 85:
		// Next 20%: Strong attack
		return Competence{Nom: "Assaut Puissant", Description: "Attaque puissante", Degats: 0, Type: "physique", Puissance: 2}, true
	default:
		// Last 15%: buff self or debuff player; randomly pick one
		if rand.Intn(2) == 0 {
			// Self buff
			opts := []Competence{
				{Nom: "Focalisation", Description: "+Précision pour l'ennemi", Degats: 0, Type: "buff", TypeEffet: "Focalisation", Puissance: 2},
				{Nom: "Fortification", Description: "+Défense/ResMag pour l'ennemi", Degats: 0, Type: "buff", TypeEffet: "Fortification", Puissance: 2},
				{Nom: "Imprégnation", Description: "+Dégâts pour l'ennemi", Degats: 0, Type: "buff", TypeEffet: "Imprégnation", Puissance: 2},
			}
			return opts[rand.Intn(len(opts))], true
		}
		// Debuff player: inverse of buffs (precision, defense, attack via damage reduction)
		opts := []Competence{
			{Nom: "Nébulation", Description: "-Précision du joueur", Degats: 0, Type: "magique", TypeEffet: "Nébulation", Puissance: 2},
			{Nom: "Brise-Armure", Description: "-Défense du joueur", Degats: 0, Type: "physique", TypeEffet: "Brise-Armure", Puissance: 2},
			{Nom: "Brise-Armure Magique", Description: "-Résistance magique du joueur", Degats: 0, Type: "magique", TypeEffet: "Brise-Armure Magique", Puissance: 2},
			{Nom: "Affaiblissement", Description: "-Dégâts infligés par le joueur", Degats: 0, Type: "magique", TypeEffet: "Affaiblissement", Puissance: 2},
		}
		return opts[rand.Intn(len(opts))], true
	}
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
func chooseCompetence(p *Personnage) (Competence, bool, bool) {
	comps := p.ArmeEquipee.Competences
	if len(comps) == 0 {
		return Competence{}, false, false
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
	fmt.Println("  R) Retour")
	fmt.Print("Votre choix (1-", len(comps), " ou R): ")
	if globalKeyEvents == nil {
		// Fallback extrême si le canal n'est pas prêt
		return comps[0], true, false
	}
	for {
		e := <-globalKeyEvents
		if e.Key == keyboard.KeyEsc {
			return Competence{}, false, true
		}
		r := e.Rune
		if r == 'r' || r == 'R' {
			return Competence{}, false, true
		}
		if r >= '1' && r <= '9' {
			idx := int(r - '0')
			if idx >= 1 && idx <= len(comps) {
				return comps[idx-1], true, false
			}
		}
		fmt.Print("(touche non assignée) Choisissez 1-", len(comps), " ou R: ")
	}
}

// Sous-menu Objets (potion, Puff 9K, etc.) — n'utilise pas le tour
// Retourne true si le joueur meurt pendant l'utilisation (ex: Puff 9K)
func objectMenu(player, enemy *Personnage) bool {
	for {
		fmt.Println("\n🎒 Objets:")
		printed := 0
		have := func(key string) int { return playerInventory[key] }
		// Soins
		if have("potions") > 0 {
			fmt.Printf("  1) Potion (x%d) — +70 PV\n", have("potions"))
			printed++
		}
		if have("potion_mineure") > 0 {
			fmt.Printf("  2) Potion Mineure (x%d) — soin léger\n", have("potion_mineure"))
			printed++
		}
		if have("potion_majeure") > 0 {
			fmt.Printf("  3) Potion Majeure (x%d) — soin puissant\n", have("potion_majeure"))
			printed++
		}
		if have("potion_supreme") > 0 {
			fmt.Printf("  4) Potion Suprême (x%d) — soin massif\n", have("potion_supreme"))
			printed++
		}
		if have("antidote") > 0 {
			fmt.Printf("  5) Antidote (x%d) — retire poison\n", have("antidote"))
			printed++
		}
		if have("vodka_vitaly") > 0 {
			fmt.Printf("  V) Vodka de Vitaly (x%d) — régénère toute la vie !\n", have("vodka_vitaly"))
			printed++
		}
		// Buffs
		if have("puff_9k") > 0 {
			fmt.Printf("  6) Puff 9K (x%d) — +15%%%% dégâts (loot) + buff, -5 PV\n", have("puff_9k"))
			printed++
		}
		if have("elixir_force") > 0 {
			fmt.Printf("  7) Élixir de Force (x%d) — buff dégâts\n", have("elixir_force"))
			printed++
		}
		if have("elixir_vitesse") > 0 {
			fmt.Printf("  8) Élixir de Vitesse (x%d) — buff dégâts magiques\n", have("elixir_vitesse"))
			printed++
		}
		if have("elixir_critique") > 0 {
			fmt.Printf("  9) Élixir de Précision (x%d) — buff dégâts/crit\n", have("elixir_critique"))
			printed++
		}
		// Offensifs/Utilitaires
		if have("potion_degats") > 0 {
			fmt.Printf("  A) Potion de Dégâts (x%d) — dégâts magiques\n", have("potion_degats"))
			printed++
		}
		if have("bombe_incendiaire") > 0 {
			fmt.Printf("  B) Bombe Incendiaire (x%d) — dégâts + brûlure\n", have("bombe_incendiaire"))
			printed++
		}
		if have("bombe_givrante") > 0 {
			fmt.Printf("  C) Bombe Givrante (x%d) — dégâts + étourdissement\n", have("bombe_givrante"))
			printed++
		}
		if have("grenade_fumigene") > 0 {
			fmt.Printf("  D) Grenade Fumigène (x%d) — nébulation (aveugle)\n", have("grenade_fumigene"))
			printed++
		}
		if have("parchemin_dispersion") > 0 {
			fmt.Printf("  E) Parchemin de Dispersion (x%d) — affaiblissement\n", have("parchemin_dispersion"))
			printed++
		}
		if printed == 0 {
			fmt.Println("  (Aucun objet disponible)")
		}
		fmt.Println("  [R]etour")
		fmt.Print("Votre choix: ")
		if globalKeyEvents == nil {
			return false
		}
		e := <-globalKeyEvents
		input := strings.ToLower(string(e.Rune))
		if e.Key == keyboard.KeyEsc {
			input = "r"
		}
		switch input {
		case "v": // Vodka de Vitaly — régénère toute la vie
			if playerInventory["vodka_vitaly"] > 0 {
				heal := player.PVMax - player.PV
				player.PV = player.PVMax
				playerInventory["vodka_vitaly"]--
				// Applique le malus d'ivresse (-30% précision pendant 3 tours)
				if eff := CreerEffet("Ivresse", 0); eff != nil {
					AppliquerEffet(player, *eff)
				}
				fmt.Printf("🍶 Vodka de Vitaly: +%d PV (PV: %d/%d) — Toute votre vie est régénérée, mais votre précision chute temporairement !\n", heal, player.PV, player.PVMax)
			} else {
				fmt.Println("❌ Vous n'avez pas de Vodka de Vitaly !")
			}
		case "1": // Potion simple +70 PV (compat historique)
			if playerInventory["potions"] > 0 {
				heal := 70
				player.PV += heal
				if player.PV > player.PVMax {
					player.PV = player.PVMax
				}
				playerInventory["potions"]--
				fmt.Printf("🧪 Potion: +%d PV (PV: %d/%d)\n", heal, player.PV, player.PVMax)
			} else {
				fmt.Println("❌ Vous n'avez pas de potion !")
			}
		case "2": // Potion Mineure
			if playerInventory["potion_mineure"] > 0 {
				comp := potionMineure.Competences[0]
				heal := -comp.Degats
				if heal < 0 {
					heal = 30
				}
				player.PV += heal
				if player.PV > player.PVMax {
					player.PV = player.PVMax
				}
				playerInventory["potion_mineure"]--
				fmt.Printf("🧪 Potion mineure: +%d PV (PV: %d/%d)\n", heal, player.PV, player.PVMax)
			} else {
				fmt.Println("❌ Vous n'avez pas de potion mineure !")
			}
		case "3": // Potion Majeure
			if playerInventory["potion_majeure"] > 0 {
				comp := potionMajeure.Competences[0]
				heal := -comp.Degats
				if heal < 0 {
					heal = 80
				}
				player.PV += heal
				if player.PV > player.PVMax {
					player.PV = player.PVMax
				}
				playerInventory["potion_majeure"]--
				fmt.Printf("🧪 Potion majeure: +%d PV (PV: %d/%d)\n", heal, player.PV, player.PVMax)
			} else {
				fmt.Println("❌ Vous n'avez pas de potion majeure !")
			}
		case "4": // Potion Suprême
			if playerInventory["potion_supreme"] > 0 {
				comp := potionSupreme.Competences[0]
				heal := -comp.Degats
				if heal < 0 {
					heal = 200
				}
				player.PV += heal
				if player.PV > player.PVMax {
					player.PV = player.PVMax
				}
				playerInventory["potion_supreme"]--
				fmt.Printf("🧪 Potion suprême: +%d PV (PV: %d/%d)\n", heal, player.PV, player.PVMax)
			} else {
				fmt.Println("❌ Vous n'avez pas de potion suprême !")
			}
		case "5": // Antidote
			if playerInventory["antidote"] > 0 {
				if eff := CreerEffet("Guérison Poison", 1); eff != nil {
					AppliquerEffet(player, *eff)
				}
				playerInventory["antidote"]--
				fmt.Println("🧯 Antidote utilisé: le poison est dissipé.")
			} else {
				fmt.Println("❌ Vous n'avez pas d'antidote !")
			}
		case "6": // Puff 9K
			if playerInventory["puff_9k"] > 0 {
				playerInventory["puff_9k"]--
				playerStats.attackBoost += 15 // bonus de loot cumulatif
				if eff := CreerEffet("Augmentation de Dégâts", 2); eff != nil {
					AppliquerEffet(player, *eff)
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
					return true
				}
			} else {
				fmt.Println("❌ Vous n'avez pas de Puff 9K !")
			}
		case "7": // Élixir de Force
			if playerInventory["elixir_force"] > 0 {
				if eff := CreerEffet("Augmentation de Dégâts", 4); eff != nil {
					AppliquerEffet(player, *eff)
				}
				playerInventory["elixir_force"]--
				fmt.Println("🧃 Élixir de Force: vos dégâts sont augmentés !")
			} else {
				fmt.Println("❌ Vous n'avez pas d'élixir de force !")
			}
		case "8": // Élixir de Vitesse
			if playerInventory["elixir_vitesse"] > 0 {
				if eff := CreerEffet("Augmentation de Dégâts Magiques", 3); eff != nil {
					AppliquerEffet(player, *eff)
				}
				playerInventory["elixir_vitesse"]--
				fmt.Println("🧃 Élixir de Vitesse: vos dégâts magiques sont augmentés !")
			} else {
				fmt.Println("❌ Vous n'avez pas d'élixir de vitesse !")
			}
		case "9": // Élixir de Précision
			if playerInventory["elixir_critique"] > 0 {
				if eff := CreerEffet("Augmentation de Dégâts", 5); eff != nil {
					AppliquerEffet(player, *eff)
				}
				playerInventory["elixir_critique"]--
				fmt.Println("🧃 Élixir de Précision: vos coups deviennent plus meurtriers !")
			} else {
				fmt.Println("❌ Vous n'avez pas d'élixir de précision !")
			}
		case "a": // Potion de Dégâts (attaque magique directe)
			if playerInventory["potion_degats"] > 0 {
				comp := potionDegats.Competences[0]
				dmg, touche, crit := resolveAttack(player, enemy, comp.Degats, comp.Type)
				if !touche {
					fmt.Println("🙃 Votre lancer de potion rate !")
				} else {
					enemy.PV -= dmg
					if enemy.PV < 0 {
						enemy.PV = 0
					}
					if crit {
						fmt.Printf("💥 Potion de dégâts critique ! %d dégâts.\n", dmg)
					} else {
						fmt.Printf("💥 Potion de dégâts inflige %d dégâts.\n", dmg)
					}
				}
				playerInventory["potion_degats"]--
			} else {
				fmt.Println("❌ Vous n'avez pas de potion de dégâts !")
			}
		case "b": // Bombe Incendiaire
			if playerInventory["bombe_incendiaire"] > 0 {
				comp := bombeIncendiaire.Competences[0]
				dmg, touche, crit := resolveAttack(player, enemy, comp.Degats, comp.Type)
				if !touche {
					fmt.Println("🧨 La bombe incendiaire n'atteint pas sa cible !")
				} else {
					enemy.PV -= dmg
					if enemy.PV < 0 {
						enemy.PV = 0
					}
					if crit {
						fmt.Printf("🔥 Explosion critique ! %d dégâts.\n", dmg)
					} else {
						fmt.Printf("🔥 Explosion de feu: %d dégâts.\n", dmg)
					}
					if eff := CreerEffet("Brûlure", comp.Puissance); eff != nil {
						AppliquerEffet(enemy, *eff)
					}
				}
				playerInventory["bombe_incendiaire"]--
			} else {
				fmt.Println("❌ Vous n'avez pas de bombe incendiaire !")
			}
		case "c": // Bombe Givrante
			if playerInventory["bombe_givrante"] > 0 {
				comp := bombeGivrante.Competences[0]
				dmg, touche, crit := resolveAttack(player, enemy, comp.Degats, comp.Type)
				if !touche {
					fmt.Println("❄️ La bombe givrante rate sa cible !")
				} else {
					enemy.PV -= dmg
					if enemy.PV < 0 {
						enemy.PV = 0
					}
					if crit {
						fmt.Printf("❄️ Explosion glaciale critique ! %d dégâts.\n", dmg)
					} else {
						fmt.Printf("❄️ Explosion de glace: %d dégâts.\n", dmg)
					}
					if eff := CreerEffet("Étourdissement", comp.Puissance); eff != nil {
						AppliquerEffet(enemy, *eff)
					}
				}
				playerInventory["bombe_givrante"]--
			} else {
				fmt.Println("❌ Vous n'avez pas de bombe givrante !")
			}
		case "d": // Grenade Fumigène
			if playerInventory["grenade_fumigene"] > 0 {
				if eff := CreerEffet("Nébulation", 3); eff != nil {
					AppliquerEffet(enemy, *eff)
				}
				playerInventory["grenade_fumigene"]--
				fmt.Println("🌫️ Grenade fumigène: l'ennemi voit mal !")
			} else {
				fmt.Println("❌ Vous n'avez pas de grenade fumigène !")
			}
		case "e": // Parchemin de Dispersion
			if playerInventory["parchemin_dispersion"] > 0 {
				if eff := CreerEffet("Affaiblissement", 2); eff != nil {
					AppliquerEffet(enemy, *eff)
				}
				playerInventory["parchemin_dispersion"]--
				fmt.Println("📜 Parchemin de Dispersion: l'ennemi est affaibli !")
			} else {
				fmt.Println("❌ Vous n'avez pas de parchemin de dispersion !")
			}
		case "r":
			return false
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

// Système de combat amélioré avec les modules existants
// Optional override for enemy naming in scripted encounters
var combatCustomEnemyName string

func combat(currentMap string, isSuper bool) interface{} {
	// Crée les entités combat
	player := buildPlayerCharacter()
	enemy := CreateRandomEnemyForMap(currentMap, isSuper)

	// Boss final personnalisé pour salle15
	if currentMap == "salle15" && combatCustomEnemyName == "" {
		// Définition explicite du boss final (ignorer le scaling générique ensuite)
		custom := Personnage{
			Nom:                "Mia Khalifa",
			PV:                 69,
			PVMax:              69,
			Armure:             69,
			ResistMag:          69,
			Precision:          0.90,
			TauxCritique:       0.69,
			MultiplicateurCrit: 1.8,
		}
		// Équipe l'arme foutre de Zeus
		_ = EquiperArme(&custom, foutreDeZeus)
		// Ajuste les dégâts pour refléter "69 attaque"
		custom.ArmeEquipee.DegatsPhysiques = 69
		custom.ArmeEquipee.DegatsMagiques = 69
		enemy = custom
	}

	// Apply a custom name if provided (used for salle15 mini-bosses and boss)
	if combatCustomEnemyName != "" {
		enemy.Nom = combatCustomEnemyName
		// reset after use to avoid leaking to future fights
		combatCustomEnemyName = ""
	}

	// Scaling supplémentaire pour salles boss progressives
	levelMultiplier := 1.0
	switch currentMap {
	case "salle12":
		levelMultiplier = 1.2 // Niveau 1/4
	case "salle13":
		levelMultiplier = 1.5 // Niveau 2/4
	case "salle14":
		levelMultiplier = 1.9 // Niveau 3/4
	case "salle15":
		// Pas de scaling : boss déjà défini avec ses stats personnalisées
		levelMultiplier = 1.0
	}
	if levelMultiplier > 1.0 {
		enemy.PV = int(float64(enemy.PV) * levelMultiplier)
		enemy.PVMax = int(float64(enemy.PVMax) * levelMultiplier)
		// Buff dégâts via augmentation base dégâts arme
		if enemy.ArmeEquipee.Nom != "" {
			enemy.ArmeEquipee.DegatsPhysiques = int(float64(enemy.ArmeEquipee.DegatsPhysiques) * (0.85 + levelMultiplier/1.5))
		}
		// Légère hausse critique
		enemy.TauxCritique += 0.03 * (levelMultiplier - 1)
		if enemy.TauxCritique > 0.60 {
			enemy.TauxCritique = 0.60
		}
	}
	enemyAttackBase := enemy.ArmeEquipee.DegatsPhysiques
	if enemyAttackBase <= 0 {
		enemyAttackBase = 12
	}

	// Reset enemy HUD action at fight start
	combatLastEnemySkill = ""
	printBattleHeader(player, enemy, isSuper)

	for player.PV > 0 && enemy.PV > 0 {
		printBattleHeader(player, enemy, isSuper)
		if playerStats.hasLegendaryWeapon {
			fmt.Println("🌟 Excalibur équipée (+50% dégâts de loot)")
		}
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
			comp, ok, back := chooseCompetence(&player)
			if back {
				// Retour au menu principal sans consommer le tour
				continue
			}
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

		case "o":
			// Sous-menu objets: n'utilise pas le tour
			if died := objectMenu(&player, &enemy); died {
				// Persiste la mort immédiate
				currentPlayer.PV = player.PV
				playerStats.attackBoost = 0
				return false
			}
			// Si l'objet a tué l'ennemi, accorder la victoire immédiatement
			if enemy.PV <= 0 {
				fmt.Println("\n🎉 VICTOIRE ! Vous avez vaincu la créature !")
				playerStats.enemiesKilled++
				coins, jackpot, _ := computeCoinLoot()
				addToInventory("pièces", coins)
				printCoinReward(coins, jackpot)
				tier := tierForMap(currentMap)
				rocks := 0
				switch tier {
				case TierTutorial, TierEarly:
					rocks = 1
				case TierMid:
					rocks = 2
				case TierLate:
					rocks = 3
				}
				if rocks > 0 {
					currentPlayer.Roches += rocks
					fmt.Printf("🪨 Vous obtenez %d roche(s) d'évolution. Total roches: %d\n", rocks, currentPlayer.Roches)
				}
				currentPlayer.PV = player.PV
				playerStats.attackBoost = 0
				fmt.Println("💨 La créature disparaît complètement dans un nuage de fumée...")
				return "disappear"
			}
			// Sinon, ne consomme pas le tour ennemi
			continue

		case "f":
			fmt.Println("💨 Vous fuyez le combat !")
			// Persiste les PV du joueur
			currentPlayer.PV = player.PV
			// Reset les bonus temporaires côté anciens stats loot
			playerStats.attackBoost = 0
			return false

		default:
			// Touche non assignée: ne consomme pas le tour
			fmt.Println("(touche non assignée) Actions valides: A, O, F")
			continue
		}

		// Fin d'action joueur: traitements d'effets sur les deux
		TraiterEffetsFinTour(&player)
		TraiterEffetsFinTour(&enemy)

		if enemy.PV <= 0 {
			fmt.Println("\n🎉 VICTOIRE ! Vous avez vaincu la créature !")

			// Incrémente le compteur d'ennemis tués (stat héritée)
			playerStats.enemiesKilled++

			// Loot spécifique ennemi
			awardEnemyLoot(enemy.Nom, isSuper)

			coins, jackpot, _ := computeCoinLoot()
			addToInventory("pièces", coins)
			printCoinReward(coins, jackpot)

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
			// No action this turn
			combatLastEnemySkill = "(étourdi)"
		} else {
			// L'ennemi choisit une compétence au hasard (biais offensif)
			ecomp, ok := pickRandomCompetence(&enemy)
			edeg := enemyAttackBase
			etype := "physique"
			if ok {
				// Scale damage based on competence category
				switch ecomp.Nom {
				case "Frappe Mesurée":
					// Moderate attack: scale around player's gear; use baseline
					edeg = int(float64(enemyAttackBase) * 1.0)
				case "Assaut Puissant":
					// Stronger attack: 1.5x baseline
					edeg = int(float64(enemyAttackBase) * 1.5)
				default:
					if ecomp.Degats > 0 {
						edeg = ecomp.Degats
					}
				}
				if ecomp.Type != "" {
					etype = ecomp.Type
				}
			}
			// Record the chosen skill name for HUD
			if ok {
				name := ecomp.Nom
				if name == "" {
					name = "Attaque"
				}
				combatLastEnemySkill = name
			} else {
				combatLastEnemySkill = "Attaque"
			}
			edmg, touche, crit := resolveAttack(&enemy, &player, edeg, etype)
			if !touche {
				fmt.Println("🌀 L'ennemi rate son attaque !")
				combatLastEnemySkill += " (raté)"
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
					// Buff/debuff handling independent of damage
					if ecomp.TypeEffet != "" {
						if isSelfBuff(ecomp.TypeEffet) || ecomp.Nom == "Renforcement" {
							if eff := CreerEffet(ecomp.TypeEffet, ecomp.Puissance); eff != nil {
								// Special chance for Renforcement: 15%
								AppliquerEffet(&enemy, *eff)
								if ecomp.Nom == "Renforcement" {
									fmt.Println("🛡️ L'ennemi se renforce et subira 50% de dégâts en moins pendant un moment !")
								} else {
									fmt.Printf("✨ L'ennemi s'applique %s.\n", ecomp.Nom)
								}
							}
						} else {
							maybeApplyEffect(&player, ecomp)
						}
					}
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

// Variante qui force un type d'ennemi si name est non vide
func combatWithAssignedType(currentMap string, isSuper bool, name string) interface{} {
	// Joueur
	player := buildPlayerCharacter()

	// Choix de l'ennemi
	var enemy Personnage
	if name == "" {
		enemy = CreateRandomEnemyForMap(currentMap, isSuper)
	} else {
		// Cherche le template par nom à partir du tier de la salle
		tier := tierForMap(currentMap)
		var pool []EnemyTemplate
		switch tier {
		case TierTutorial:
			pool = tutorialPool
		case TierEarly:
			pool = earlyPool
		case TierMid:
			pool = midPool
		case TierLate:
			pool = latePool
		default:
			pool = earlyPool
		}
		found := false
		for _, t := range pool {
			if t.Name == name {
				enemy = NewEnemyFromTemplate(t, isSuper)
				// Ancien comportement: pas de préfixe de tier, éventuellement marqueur super
				if isSuper {
					enemy.Nom = "💀 " + enemy.Nom
				}
				found = true
				break
			}
		}
		if !found {
			enemy = CreateRandomEnemyForMap(currentMap, isSuper)
		}
	}

	enemyAttackBase := enemy.ArmeEquipee.DegatsPhysiques
	if enemyAttackBase <= 0 {
		enemyAttackBase = 12
	}

	fmt.Println("\n🗡️  COMBAT ENGAGÉ ! 🗡️")
	enemyEmoji := emojiForEnemyName(enemy.Nom)
	if isSuper {
		fmt.Printf("Vous affrontez %s %s (SURPUISSANT)\n", enemyEmoji, enemy.Nom)
	} else {
		fmt.Printf("Vous affrontez %s %s\n", enemyEmoji, enemy.Nom)
	}

	for player.PV > 0 && enemy.PV > 0 {
		fmt.Printf("\n💚 Vos PV: %d/%d | 💀 PV Ennemi: %d/%d\n", player.PV, player.PVMax, enemy.PV, enemy.PVMax)
		if playerStats.hasLegendaryWeapon {
			fmt.Println("🌟 Excalibur équipée (+50% dégâts de loot)")
		}

		fmt.Println("Actions: [A]ttaquer, [O]bjet, [F]uir")
		fmt.Print("Choisissez une action: ")
		if globalKeyEvents == nil {
			fmt.Println("(clavier non initialisé)")
			return false
		}
		e := <-globalKeyEvents
		input := strings.ToLower(string(e.Rune))
		if e.Key == keyboard.KeyEsc {
			input = "f"
		}

		AppliquerProtectionsArtefactsDebutTour(&player)

		switch input {
		case "a":
			comp, ok, back := chooseCompetence(&player)
			if back {
				continue
			}
			if !ok {
				comp = Competence{Nom: "Attaque", Degats: 15, Type: "physique"}
			}
			if comp.Degats <= 0 && comp.TypeEffet != "" && isSelfBuff(comp.TypeEffet) {
				if eff := CreerEffet(comp.TypeEffet, comp.Puissance); eff != nil {
					AppliquerEffet(&player, *eff)
				}
			} else if comp.Degats <= 0 && comp.TypeEffet != "" {
				if eff := CreerEffet(comp.TypeEffet, comp.Puissance); eff != nil {
					AppliquerEffet(&enemy, *eff)
				}
			} else {
				degatsBase := comp.Degats
				typeAtk := comp.Type
				if degatsBase <= 0 {
					degatsBase = 15
				}
				if typeAtk == "" {
					typeAtk = "physique"
				}
				dmg, touche, crit := resolveAttack(&player, &enemy, degatsBase, typeAtk)
				if touche {
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
				} else {
					fmt.Println("🙃 Votre attaque rate sa cible !")
				}
			}
		case "o":
			if died := objectMenu(&player, &enemy); died {
				currentPlayer.PV = player.PV
				playerStats.attackBoost = 0
				return false
			}
			if enemy.PV <= 0 {
				fmt.Println("\n🎉 VICTOIRE ! Vous avez vaincu la créature !")
				playerStats.enemiesKilled++
				coins, jackpot, _ := computeCoinLoot()
				addToInventory("pièces", coins)
				printCoinReward(coins, jackpot)
				tier := tierForMap(currentMap)
				rocks := 0
				switch tier {
				case TierTutorial, TierEarly:
					rocks = 1
				case TierMid:
					rocks = 2
				case TierLate:
					rocks = 3
				}
				if rocks > 0 {
					currentPlayer.Roches += rocks
					fmt.Printf("🪨 Vous obtenez %d roche(s) d'évolution. Total roches: %d\n", rocks, currentPlayer.Roches)
				}
				currentPlayer.PV = player.PV
				playerStats.attackBoost = 0
				fmt.Println("💨 La créature disparaît complètement dans un nuage de fumée...")
				return "disappear"
			}
			continue
		case "f":
			fmt.Println("💨 Vous fuyez le combat !")
			currentPlayer.PV = player.PV
			playerStats.attackBoost = 0
			return false
		default:
			// Touche non assignée: ne consomme pas le tour
			fmt.Println("(touche non assignée) Actions valides: A, O, F")
			continue
		}

		TraiterEffetsFinTour(&player)
		TraiterEffetsFinTour(&enemy)

		if enemy.PV <= 0 {
			fmt.Println("\n🎉 VICTOIRE ! Vous avez vaincu la créature !")
			playerStats.enemiesKilled++
			awardEnemyLoot(enemy.Nom, isSuper)
			awardEnemyLoot(enemy.Nom, isSuper)
			awardEnemyLoot(enemy.Nom, isSuper)
			coins, jackpot, _ := computeCoinLoot()
			addToInventory("pièces", coins)
			printCoinReward(coins, jackpot)
			tier := tierForMap(currentMap)
			rocks := 0
			switch tier {
			case TierTutorial, TierEarly:
				rocks = 1
			case TierMid:
				rocks = 2
			case TierLate:
				rocks = 3
			}
			if rocks > 0 {
				currentPlayer.Roches += rocks
				fmt.Printf("🪨 Vous obtenez %d roche(s) d'évolution. Total roches: %d\n", rocks, currentPlayer.Roches)
			}
			currentPlayer.PV = player.PV
			playerStats.attackBoost = 0
			fmt.Println("💨 La créature disparaît complètement dans un nuage de fumée...")
			return "disappear"
		}

		if EstEtourdi(&enemy) {
			fmt.Println("😵‍💫 L'ennemi est étourdi et rate son tour !")
			combatLastEnemySkill = "(étourdi)"
		} else {
			ecomp, ok := pickRandomCompetence(&enemy)
			edeg := enemyAttackBase
			etype := "physique"
			if ok {
				// Scale damage based on competence like the main combat loop
				switch ecomp.Nom {
				case "Frappe Mesurée":
					edeg = int(float64(enemyAttackBase) * 1.0)
				case "Assaut Puissant":
					edeg = int(float64(enemyAttackBase) * 1.5)
				default:
					if ecomp.Degats > 0 {
						edeg = ecomp.Degats
					}
				}
				if ecomp.Type != "" {
					etype = ecomp.Type
				}
			}
			if ok {
				if ecomp.Nom != "" {
					combatLastEnemySkill = ecomp.Nom
				} else {
					combatLastEnemySkill = "Attaque"
				}
			} else {
				combatLastEnemySkill = "Attaque"
			}
			edmg, touche, crit := resolveAttack(&enemy, &player, edeg, etype)
			if !touche {
				fmt.Println("🌀 L'ennemi rate son attaque !")
				combatLastEnemySkill += " (raté)"
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
					if ecomp.Degats <= 0 && ecomp.TypeEffet != "" && isSelfBuff(ecomp.TypeEffet) {
						if eff := CreerEffet(ecomp.TypeEffet, ecomp.Puissance); eff != nil {
							AppliquerEffet(&enemy, *eff)
							fmt.Printf("✨ L'ennemi s'applique %s.\n", ecomp.Nom)
						}
					} else if ecomp.TypeEffet != "" {
						maybeApplyEffect(&player, ecomp)
					}
				}
			}
		}

		TraiterEffetsFinTour(&player)
		TraiterEffetsFinTour(&enemy)

		if player.PV <= 0 {
			fmt.Println("\n💀 DÉFAITE ! Vous avez été vaincu...")
			fmt.Println("🔄 Vous retournez au début de la salle.")
			currentPlayer.PV = player.PV
			playerStats.attackBoost = 0
			return false
		}
	}

	return false
}
