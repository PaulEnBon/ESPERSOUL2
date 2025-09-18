package main

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/mattn/go-runewidth"
)

// Variables globales nécessaires (restaurées)
var useASCII bool = false
var randomMobsSalle3 []struct{ x, y int }
var randomMobsSalle2 []struct{ x, y int }
var randomMobsSalle10 []struct{ x, y int }

// Nom de la carte courante (pour affichage des emojis par coordonnée)
var currentMapDisplayName string

// Stocke pour chaque carte et coordonnée l'emoji d'ennemi choisi (clé "x,y")
var enemyDisplayedEmoji = map[string]map[string]string{}

// Paramètres pour génération aléatoire (valeurs précédentes supposées)
const (
	// Salle3: augmentation massive du nombre d'ennemis
	minRandomMobsSalle3  = 8  // ancien 3
	maxRandomMobsSalle3  = 14 // ancien 6
	numRandomMobsSalle2  = 5  // augmenté de 4 à 5 (inchangé ici)
	numRandomMobsSalle10 = 5
)

// copyMap : fait une copie profonde de la carte (restauration)
func copyMap(src [][]int) [][]int {
	dst := make([][]int, len(src))
	for i := range src {
		dst[i] = make([]int, len(src[i]))
		copy(dst[i], src[i])
	}
	return dst
}

// Restauration de la fonction cellToSymbol (version avant l'ajout des émojis dynamiques par tier)
func cellToSymbol(val int) string {
	if useASCII { // Mode ASCII simplifié
		switch val {
		case 1:
			return "P"
		case 2:
			return "E"
		case 12:
			return "S" // Super ennemi
		case 3:
			return "N" // PNJ
		case 4:
			return "M" // Marchand
		case 5:
			return "F" // Forgeron
		case 6:
			return "C" // Coffre
		case 35:
			return "B"
		case 46, 47:
			return "*"
		case 66:
			return "G" // spawner
		case 67, 70, 73, 76:
			return "m" // mini boss
		case 68, 72, 75, 78:
			return "B" // boss
		case 71, 74, 77:
			return "G" // spawner niv 2/3/4
		case 8, 9:
			return "#"
		default:
			return "."
		}
	}

	switch val { // Version emoji
	case 8:
		return "⬜"
	case 14, 38, 52, 55:
		return "→"
	case 13, 40, 53, 54:
		return "←"
	case 15, 27, 31, 10, 21, 42, 51, 57:
		return "↓"
	case 28, 33, 34, 44, 50, 56:
		return "↑"
	case 9:
		return "▨"
	case 30: // Porte secrète
		return "▧"
	case 32: // Sortie de la salle secrète
		return "🚪"
	case 7, 20:
		return "↑"
	case 1:
		return "🎮" // Joueur
	case 2:
		return "👹" // Ennemi (fallback si pas de mapping spécifique)
	case 12:
		return "💀" // Super Ennemi
	case 3:
		return "👨" // PNJ
	case 4:
		return "🛒" // Marchand
	case 5:
		return "🔨" // Forgeron
	case 6:
		return "💰" // Coffre
	case 35: // Pierre bloquant la sortie de salle1
		return "๑"
	case 46: // Frame explosion 1
		return "💥"
	case 47: // Frame explosion 2
		return "🔥"
	case 66:
		return "⚙️" // spawner
	case 67:
		return "🛡️" // mini boss
	case 68:
		return "👑" // boss
	case 70, 73, 76: // mini boss niveaux 2,3,4
		return "🛡️"
	case 71, 74, 77: // spawner niveaux 2,3,4
		return "⚙️"
	case 72, 75, 78: // boss niveaux 2,3,4
		return "👑"
	case 11, 16, 17, 18, 19, 22, 23, 24, 25, 26, 29, 36, 37, 39, 41, 43, 45, 58, 59, 60, 61, 62, 63, 64, 65, 0:
		return "•"
	default:
		return "?"
	}
}

// Retourne un symbole adapté tenant compte de la coordonnée (per-class emoji)
func cellToSymbolAt(x, y, val int) string {
	// Cas spécial: Mentor (salle1 coord (8,3)) -> toujours 🧙 (version simplifiée)
	if !useASCII {
		if currentMapDisplayName == "salle1" && x == 8 && y == 3 {
			// Que ce soit encore un ennemi (2) ou déjà PNJ (3) => afficher le même emoji mentor
			if val == 2 || val == 3 || val == 12 { // inclure safety super flag improbable
				return "🧙"
			}
		}
		if val == 2 || val == 12 {
			if m := enemyDisplayedEmoji[currentMapDisplayName]; m != nil {
				key := fmt.Sprintf("%d,%d", x, y)
				if e, ok := m[key]; ok {
					return e
				}
			}
		}
	}
	return cellToSymbol(val)
}

// Choisit aléatoirement un emoji d'ennemi compatible avec le tier de la carte
func randomEnemyEmojiForMap(mapName string) string {
	tier := tierForMap(mapName)
	var pool []string
	switch tier {
	case TierTutorial:
		pool = []string{"Rat", "Gelée"}
	case TierEarly:
		pool = []string{"Brigand", "Archer", "Apprenti Pyro", "Rat", "Gelée"}
	case TierMid:
		pool = []string{"Chevalier", "Berserker", "Mage Sombre"}
	case TierLate:
		pool = []string{"Seigneur Démon", "Archimage", "Champion déchu", "Mage Sombre"}
	default:
		pool = []string{"Rat"}
	}
	name := pool[rand.Intn(len(pool))]
	// Local mini switch d'emoji (évite dépendance sur combat.go)
	switch name {
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
	default:
		return "👹"
	}
}

// Assigne des emojis aux cases ennemies d'une carte si non déjà définis
func assignEnemyEmojis(mapName string, mapData [][]int) {
	if enemyDisplayedEmoji[mapName] == nil {
		enemyDisplayedEmoji[mapName] = map[string]string{}
	}
	for y := range mapData {
		for x := range mapData[y] {
			if mapData[y][x] == 2 || mapData[y][x] == 12 { // ennemi ou super ennemi
				key := fmt.Sprintf("%d,%d", x, y)
				if _, exists := enemyDisplayedEmoji[mapName][key]; !exists {
					enemyDisplayedEmoji[mapName][key] = randomEnemyEmojiForMap(mapName)
				}
			}
		}
	}
}

// Helpers d'abréviation pour la colonne compétences
func shortType(t string) string {
	switch t {
	case "physique":
		return "P"
	case "magique":
		return "M"
	default:
		if len(t) > 1 {
			return strings.ToUpper(string([]rune(t)[0]))
		}
		return "?"
	}
}

func abbrevEffet(e string) string {
	switch e {
	case "Saignement":
		return "Saig"
	case "Brise-Armure":
		return "BrArm"
	case "Brise-Armure Magique":
		return "BrArM"
	case "Étourdissement":
		return "Stun"
	case "Brûlure":
		return "Brul"
	case "Nébulation":
		return "Nébu"
	case "Affaiblissement":
		return "Affaib"
	case "Défavorisation":
		return "Défav"
	case "Augmentation de Dégâts":
		return "+ATK"
	case "Augmentation de Dégâts Magiques":
		return "+MATK"
	case "Régénération":
		return "Regen"
	case "Guérison Poison":
		return "Antid"
	default:
		if len(e) > 6 {
			return e[:6]
		}
		return e
	}
}

// Affiche la map avec HUD optimisé
func printMap(mapData [][]int) {
	fmt.Print("\033[H\033[2J") // Nettoie l'écran

	// En-tête du jeu
	fmt.Println("═══════════════════════════════════════════════════════════════════════")
	fmt.Println("🎮                        DONJON MYSTÉRIEUX                        🎮")
	fmt.Println("═══════════════════════════════════════════════════════════════════════")

	// Prépare les lignes d'informations pour l'affichage côte à côte
	// Préparer la ligne des artefacts équipés
	artNames := []string{}
	for _, slot := range currentPlayer.ArtefactsEquipes {
		if slot != nil {
			artNames = append(artNames, slot.Nom)
		}
	}
	artefactsStr := "Aucun"
	if len(artNames) > 0 {
		artefactsStr = strings.Join(artNames, ", ")
	}

	// Colonne 1 : Statistiques globales simplifiées (seulement ce que l'utilisateur souhaite conserver)
	infoLines := []string{
		"📊 === STATISTIQUES ===",
		fmt.Sprintf("💰 Pièces: %d", playerInventory["pièces"]),
		fmt.Sprintf("☠️ Ennemis tués: %d", playerStats.enemiesKilled),
		fmt.Sprintf("🧿 Artefacts: %s", artefactsStr),
	}

	// Colonne 2 : Statistiques détaillées du joueur
	// Arme équipée : si vide (ex: jamais explicitement équipée encore), tenter de récupérer via niveau
	weaponName := "Aucune"
	phys, mag := 0, 0
	critW := 0.0
	if currentPlayer.ArmeEquipee.Nom == "" && currentPlayer.NiveauArme >= 0 && currentPlayer.NiveauArme < len(currentPlayer.ArmesDisponibles) {
		// Auto-récupération silencieuse (n'affecte pas les stats cumulatives déjà appliquées ailleurs)
		weapon := currentPlayer.ArmesDisponibles[currentPlayer.NiveauArme]
		weaponName = weapon.Nom
		phys = weapon.DegatsPhysiques
		mag = weapon.DegatsMagiques
		critW = weapon.TauxCritique * 100
	} else if currentPlayer.ArmeEquipee.Nom != "" {
		weaponName = currentPlayer.ArmeEquipee.Nom
		phys = currentPlayer.ArmeEquipee.DegatsPhysiques
		mag = currentPlayer.ArmeEquipee.DegatsMagiques
		critW = currentPlayer.ArmeEquipee.TauxCritique * 100
	}

	// Armure équipée : même logique de fallback
	armorName := "Aucune"
	armorDef, armorRes, armorHP := 0, 0, 0
	if currentPlayer.ArmureEquipee.Nom == "" && currentPlayer.NiveauArmure >= 0 && currentPlayer.NiveauArmure < len(currentPlayer.ArmuresDisponibles) {
		arm := currentPlayer.ArmuresDisponibles[currentPlayer.NiveauArmure]
		armorName = arm.Nom
		armorDef = arm.Defense
		armorRes = arm.Resistance
		armorHP = arm.HP
	} else if currentPlayer.ArmureEquipee.Nom != "" {
		arm := currentPlayer.ArmureEquipee
		armorName = arm.Nom
		armorDef = arm.Defense
		armorRes = arm.Resistance
		armorHP = arm.HP
	}
	playerLines := []string{
		"🧝 === JOUEUR ===",
		fmt.Sprintf("👤 %s", currentPlayer.Nom),
		fmt.Sprintf("❤️ PV: %d/%d", currentPlayer.PV, currentPlayer.PVMax),
		fmt.Sprintf("🛡️ Armure Totale: %d", currentPlayer.Armure),
		fmt.Sprintf("🔮 Résist. Mag Totale: %d", currentPlayer.ResistMag),
		fmt.Sprintf("🎯 Précision: %.0f%%", currentPlayer.Precision*100),
		fmt.Sprintf("💥 Crit Base: %.0f%% x%.2f", currentPlayer.TauxCritique*100, currentPlayer.MultiplicateurCrit),
		"",
		"⚔️ === ARME ===",
		fmt.Sprintf("Nom: %s", weaponName),
		fmt.Sprintf("Phys/Mag: %d / %d", phys, mag),
		fmt.Sprintf("Crit Arme: %.0f%%", critW),
		"",
		"🛡️ === ARMURE ===",
		fmt.Sprintf("Nom: %s", armorName),
		fmt.Sprintf("Déf/Res: %d / %d", armorDef, armorRes),
		fmt.Sprintf("Bonus PV: %d", armorHP),
	}
	if playerStats.attackBoost > 0 {
		playerLines = append(playerLines, fmt.Sprintf("📈 Boost ATK: +%d%%", playerStats.attackBoost))
	}
	if playerInventory["potions"] > 0 {
		playerLines = append(playerLines, fmt.Sprintf("🧪 Potions: %d", playerInventory["potions"]))
	}
	// Ancienne colonne 3 : compétences -> intégrée sous les statistiques (infoLines)
	compLines := []string{"", "🗡️ === COMPÉTENCES ==="}
	var displayWeapon Arme
	if currentPlayer.ArmeEquipee.Nom != "" {
		displayWeapon = currentPlayer.ArmeEquipee
	} else if currentPlayer.NiveauArme >= 0 && currentPlayer.NiveauArme < len(currentPlayer.ArmesDisponibles) {
		displayWeapon = currentPlayer.ArmesDisponibles[currentPlayer.NiveauArme]
	}
	if displayWeapon.Nom == "" || len(displayWeapon.Competences) == 0 {
		compLines = append(compLines, "(Aucune compétence)")
	} else {
		for _, c := range displayWeapon.Competences {
			line := fmt.Sprintf("• %s [%s %d", c.Nom, shortType(c.Type), c.Degats)
			if c.TypeEffet != "" {
				line += fmt.Sprintf(" | %s", abbrevEffet(c.TypeEffet))
			}
			line += "]"
			compLines = append(compLines, line)
		}
	}
	// Fusion : ajouter compLines à la suite des infoLines
	infoLines = append(infoLines, compLines...)

	// Largeurs & itérations (désormais 2 colonnes : stats+compétences et joueur)
	mapHeight := len(mapData)
	maxLines := mapHeight
	for _, l := range []int{len(infoLines), len(playerLines)} {
		if l > maxLines {
			maxLines = l
		}
	}

	cellWidth := 2
	if useASCII {
		cellWidth = 1
	}

	// Calcul largeurs des colonnes texte (1=info+comp, 2=player)
	col1Width, col2Width := 0, 0
	for _, l := range infoLines {
		if w := runewidth.StringWidth(l); w > col1Width {
			col1Width = w
		}
	}
	for _, l := range playerLines {
		if w := runewidth.StringWidth(l); w > col2Width {
			col2Width = w
		}
	}

	// Largeur fixe de la partie carte pour aligner les séparateurs
	// (Empêche le décalage vertical des lignes blanches dû à la variation de largeur des emojis.)
	mapLineWidth := len(mapData[0]) * (cellWidth + 1) // chaque case = 1 espace + symbole (et pad)
	for i := 0; i < maxLines; i++ {
		var mapLine string
		if i < mapHeight {
			var b strings.Builder
			for xIdx, val := range mapData[i] {
				sym := cellToSymbolAt(xIdx, i, val)
				w := runewidth.StringWidth(sym)
				pad := cellWidth - w
				if pad < 0 {
					pad = 0
				}
				b.WriteString(" ")
				b.WriteString(sym)
				b.WriteString(strings.Repeat(" ", pad))
			}
			mapLine = b.String()
		} else {
			mapLine = strings.Repeat(" ", mapLineWidth)
		}
		// Normalise la longueur pour éviter les décalages dus à des largeurs de runes imprévues
		currentWidth := runewidth.StringWidth(mapLine)
		if currentWidth < mapLineWidth {
			mapLine += strings.Repeat(" ", mapLineWidth-currentWidth)
		}
		fmt.Print(mapLine)

		// Colonnes texte (2 colonnes désormais)
		fmt.Print("   │ ")
		c1 := ""
		if i < len(infoLines) {
			c1 = infoLines[i]
		}
		fmt.Print(c1)
		if pad := col1Width - runewidth.StringWidth(c1); pad > 0 {
			fmt.Print(strings.Repeat(" ", pad))
		}

		fmt.Print(" │ ")
		c2 := ""
		if i < len(playerLines) {
			c2 = playerLines[i]
		}
		fmt.Print(c2)
		if pad := col2Width - runewidth.StringWidth(c2); pad > 0 {
			fmt.Print(strings.Repeat(" ", pad))
		}
		fmt.Println()
	}

	fmt.Println("═══════════════════════════════════════════════════════════════════════")
}

// Affiche un HUD compact pour l'inventaire uniquement
// (Ancien HUD compact et helpers supprimés car non utilisés après refonte inventaire)

// Trouve le joueur
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

// Place le joueur à un point spécifique
func placePlayerAt(mapData [][]int, x, y int) {
	if y < 0 || y >= len(mapData) || x < 0 || x >= len(mapData[0]) {
		for i := 0; i < len(mapData); i++ {
			for j := 0; j < len(mapData[i]); j++ {
				if mapData[i][j] == 0 {
					x, y = j, i
					goto placePlayer
				}
			}
		}
		return
	}

placePlayer:
	for i := range mapData {
		for j := range mapData[i] {
			if mapData[i][j] == 1 {
				mapData[i][j] = 0
			}
		}
	}

	if mapData[y][x] == 9 {
		for _, offset := range []struct{ dx, dy int }{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
			nx, ny := x+offset.dx, y+offset.dy
			if ny >= 0 && ny < len(mapData) && nx >= 0 && nx < len(mapData[0]) && mapData[ny][nx] != 9 {
				x, y = nx, ny
				break
			}
		}
	}

	mapData[y][x] = 1
}

// Place le joueur près d'une position donnée
func placePlayerNearby(mapData [][]int, targetX, targetY int) {
	for _, offset := range []struct{ dx, dy int }{{0, 1}, {1, 0}, {0, -1}, {-1, 0}, {1, 1}, {-1, -1}, {1, -1}, {-1, 1}} {
		nx, ny := targetX+offset.dx, targetY+offset.dy
		if ny >= 0 && ny < len(mapData) && nx >= 0 && nx < len(mapData[0]) && mapData[ny][nx] == 0 {
			mapData[ny][nx] = 1
			return
		}
	}

	for y := 0; y < len(mapData); y++ {
		for x := 0; x < len(mapData[y]); x++ {
			if mapData[y][x] == 0 {
				mapData[y][x] = 1
				return
			}
		}
	}
}

// Génère des mobs aléatoires dans la salle3
func generateRandomMobs(mapData [][]int) {
	randomMobsSalle3 = []struct{ x, y int }{}

	// Sécurité: dimensions minimales
	h := len(mapData)
	if h == 0 {
		return
	}
	w := len(mapData[0])
	if w < 3 || h < 3 { // trop petit pour placement intérieur
		return
	}

	// génère entre minRandomMobsSalle3 et maxRandomMobsSalle3 (inclus)
	rangeSize := maxRandomMobsSalle3 - minRandomMobsSalle3 + 1
	if rangeSize < 1 {
		rangeSize = 1
	}
	numMobs := minRandomMobsSalle3 + rand.Intn(rangeSize)
	attempts := 0
	maxAttempts := 200

	innerW := w - 2
	innerH := h - 2
	if innerW <= 0 || innerH <= 0 {
		return
	}
	for len(randomMobsSalle3) < numMobs && attempts < maxAttempts {
		x := 1 + rand.Intn(innerW)
		y := 1 + rand.Intn(innerH)

		if mapData[y][x] == 0 {
			isOccupied := false
			for _, mob := range randomMobsSalle3 {
				if mob.x == x && mob.y == y {
					isOccupied = true
					break
				}
			}

			if !isOccupied {
				randomMobsSalle3 = append(randomMobsSalle3, struct{ x, y int }{x, y})
				mapData[y][x] = 2
			}
		}
		attempts++
	}

	fmt.Printf("🎲 %d mobs aléatoires générés dans la salle3!\n", len(randomMobsSalle3))
}

// Génère 4 mobs aléatoires dans la salle2
func generateRandomMobsSalle2(mapData [][]int) {
	randomMobsSalle2 = []struct{ x, y int }{}

	h := len(mapData)
	if h == 0 {
		return
	}
	w := len(mapData[0])

	attempts := 0
	maxAttempts := 200

	// Restreindre aux cases intérieures pour éviter murs/portes sur les bords
	for len(randomMobsSalle2) < numRandomMobsSalle2 && attempts < maxAttempts {
		if w <= 2 || h <= 2 {
			break
		}
		x := 1 + rand.Intn(w-2)
		y := 1 + rand.Intn(h-2)

		if mapData[y][x] == 0 { // seulement sur sol vide
			// éviter doublons
			isOccupied := false
			for _, mob := range randomMobsSalle2 {
				if mob.x == x && mob.y == y {
					isOccupied = true
					break
				}
			}

			if !isOccupied {
				randomMobsSalle2 = append(randomMobsSalle2, struct{ x, y int }{x, y})
				mapData[y][x] = 2
			}
		}
		attempts++
	}

	fmt.Printf("🎲 %d mobs aléatoires générés dans la salle2!\n", len(randomMobsSalle2))
}

// Génère entre 10 et 15 ennemis dans la salle9 à chaque entrée
func generateRandomMobsSalle9(mapData [][]int) {
	h := len(mapData)
	if h == 0 {
		return
	}
	w := len(mapData[0])

	// Nombre aléatoire entre 10 et 15 inclus
	numMobs := 10 + rand.Intn(6)

	placed := 0
	attempts := 0
	maxAttempts := 500

	for placed < numMobs && attempts < maxAttempts {
		if w <= 2 || h <= 2 {
			break
		}
		// Choisir uniquement les cases intérieures
		x := 1 + rand.Intn(w-2)
		y := 1 + rand.Intn(h-2)

		// Placer uniquement sur sol vide (0), éviter portes/marqueurs
		if mapData[y][x] == 0 {
			mapData[y][x] = 2
			placed++
		}
		attempts++
	}
}

// Génère des mobs (emplacements) dans la salle10
func generateRandomMobsSalle10(mapData [][]int) {
	randomMobsSalle10 = []struct{ x, y int }{}

	h := len(mapData)
	if h == 0 {
		return
	}
	w := len(mapData[0])

	attempts := 0
	maxAttempts := 300

	for len(randomMobsSalle10) < numRandomMobsSalle10 && attempts < maxAttempts {
		if w <= 2 || h <= 2 {
			break
		}
		x := 1 + rand.Intn(w-2)
		y := 1 + rand.Intn(h-2)

		if mapData[y][x] == 0 {
			// éviter doublons
			occupied := false
			for _, m := range randomMobsSalle10 {
				if m.x == x && m.y == y {
					occupied = true
					break
				}
			}
			if !occupied {
				randomMobsSalle10 = append(randomMobsSalle10, struct{ x, y int }{x, y})
				// On n'écrit pas ici: l'écriture (2 ou 12) est faite dans applyEnemyStates (salle10)
			}
		}
		attempts++
	}
}
