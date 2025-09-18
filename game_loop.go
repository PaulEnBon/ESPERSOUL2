package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/eiannone/keyboard"
)

// ===================== CHAUDRON : Craft d'artefacts via loots ennemis =====================
// Chaque recette consomme des objets de l'inventaire (drops spécifiques) pour créer un artefact permanent.
// Les artefacts ne sont ajoutés qu'une seule fois (si déjà possédé, la recette peut être ignorée ou redonnée comme confirmation simple).
type cauldronRecipe struct {
	inputs       map[string]int // item -> quantité nécessaire
	artefactName string         // Nom EXACT de l'artefact dans ArtefactsDisponibles
	description  string         // Texte affiché dans le menu
}

var cauldronRecipes = []cauldronRecipe{
	// Défense / bas-moyen (peu de composants, pas ou peu de rares)
	{inputs: map[string]int{"insigne_chevalier": 3, "gelée_visqueuse": 2}, artefactName: "Gant Anti-Émeute", description: "+Armure +Précision faible"},
	{inputs: map[string]int{"insigne_chevalier": 4, "gelée_visqueuse": 3}, artefactName: "Rune de Trempe", description: "+Armure +RésistMag"},
	{inputs: map[string]int{"insigne_chevalier": 5, "essence_sombre": 2}, artefactName: "Glyphe de Parade", description: "+Armure%"},
	{inputs: map[string]int{"gelée_visqueuse": 5, "coeur_de_gelée": 1}, artefactName: "Coquille Abyssale", description: "+RésistMag"},
	{inputs: map[string]int{"gelée_visqueuse": 4, "dent_rat": 4}, artefactName: "Relique de la Sylve", description: "+Armure +Précision"},

	// Physique / early (utilise dents, sang, insignes)
	{inputs: map[string]int{"dent_rat": 6, "capuche_brigand": 2}, artefactName: "Insigne du Sergent", description: "+Dégâts +Précision"},
	{inputs: map[string]int{"dent_rat": 8, "sang_berserker": 3}, artefactName: "Dent de Mammouth", description: "+Dégâts physiques"},
	{inputs: map[string]int{"sang_berserker": 6, "insigne_chevalier": 4, "dent_rat_luisante": 1}, artefactName: "Coutelas Runique", description: "+Dégâts +Crit"},
	{inputs: map[string]int{"sang_berserker": 5, "plume_fleche": 5, "dague_ensorcelée": 1}, artefactName: "Boussole de Chasseur", description: "+Précision +Dégâts"},
	{inputs: map[string]int{"sang_berserker": 6, "capuche_brigand": 4, "talisman_fureur": 1}, artefactName: "Médaillon du Chasseur de Mages", description: "+Physique +RésistMag"},

	// Crit / précision (archer, brigand, champion)
	{inputs: map[string]int{"plume_fleche": 6, "capuche_brigand": 4}, artefactName: "Carquois des Mille Flèches", description: "+Précision +Crit"},
	{inputs: map[string]int{"capuche_brigand": 6, "plume_fleche": 4, "carquois_gravé": 1}, artefactName: "Puce de Visée", description: "+Précision +Crit"},
	{inputs: map[string]int{"plume_fleche": 6, "embleme_champion": 4, "aiguille_du_destin": 1}, artefactName: "Bandeau du Ronin", description: "+Précision +Crit"},
	{inputs: map[string]int{"embleme_champion": 6, "plume_fleche": 4, "aiguille_du_destin": 2}, artefactName: "Œil de Lynx", description: "+Précision +Crit"},
	{inputs: map[string]int{"embleme_champion": 5, "capuche_brigand": 5, "aiguille_du_destin": 2}, artefactName: "Peau de Banane Sacrée", description: "+Précision"},

	// Magie (pyro, mage sombre, archimage)
	{inputs: map[string]int{"cendre_infernale": 6, "essence_sombre": 3}, artefactName: "Pierre d'Ignition", description: "+Magie +Crit"},
	{inputs: map[string]int{"cendre_infernale": 6, "braise_eternelle": 1, "parchemin_arcane": 4}, artefactName: "Talisman du Brasier", description: "+Magie +Physique"},
	{inputs: map[string]int{"essence_sombre": 6, "noyau_occulte": 1, "parchemin_arcane": 4}, artefactName: "Perle d'Æther", description: "+Magie +Précision"},
	{inputs: map[string]int{"parchemin_arcane": 7, "sceau_archimage": 2, "essence_sombre": 5}, artefactName: "Médaillon de Foudre Pure", description: "+Magie"},
	{inputs: map[string]int{"essence_sombre": 8, "noyau_occulte": 2, "fragment_demoniaque": 1}, artefactName: "Éclat de Foudre Gelée", description: "+Magie +Précision"},

	// Crit & Magie haute (démon / mix)
	{inputs: map[string]int{"corne_demon": 8, "fragment_demoniaque": 2, "noyau_occulte": 1}, artefactName: "Anneau des Tempêtes", description: "+Magie +Crit fort"},

	// Hybrides utilitaires supplémentaires
	{inputs: map[string]int{"insigne_chevalier": 4, "corne_demon": 3, "fragment_demoniaque": 1}, artefactName: "Anneau des Tempêtes", description: "(Variante)"},
	// (Note: L'Anneau des Tempêtes a 2 recettes possibles – le joueur peut en utiliser une)

	// Défense + résistance (mix chevalier + gelée + sombre)
	{inputs: map[string]int{"insigne_chevalier": 6, "gelée_visqueuse": 4, "essence_sombre": 2}, artefactName: "Conque des Profondeurs", description: "+RésistMag +Précision"},

	// Artefacts de dissipation (faible coût + 1 rare ciblée pour signifier spécialisation)
	{inputs: map[string]int{"gelée_visqueuse": 4, "dent_rat": 2}, artefactName: "Antidote Éternel", description: "Anti-Poison"},
	{inputs: map[string]int{"cendre_infernale": 4, "gelée_visqueuse": 2}, artefactName: "Talisman Éteigneflamme", description: "Anti-Brûlure"},
	{inputs: map[string]int{"sang_berserker": 4, "dent_rat": 2}, artefactName: "Sceau Hémostatique", description: "Anti-Saignement"},
	{inputs: map[string]int{"essence_sombre": 4, "capuche_brigand": 2}, artefactName: "Pendentif de Courage", description: "Anti-Peur"},
	{inputs: map[string]int{"plume_fleche": 4, "parchemin_arcane": 2}, artefactName: "Talisman de Vigilance", description: "Anti-Étourdissement"},
	{inputs: map[string]int{"parchemin_arcane": 4, "capuche_brigand": 2}, artefactName: "Sceau de Focalisation", description: "Anti Nébulation/Défavorisation"},
	{inputs: map[string]int{"insigne_chevalier": 4, "essence_sombre": 2}, artefactName: "Glyphe de Bastion", description: "Anti Brise-Armure"},
	{inputs: map[string]int{"sang_berserker": 3, "cendre_infernale": 3}, artefactName: "Cachet de Détermination", description: "Anti débuffs attaque"},

	// Restants / précision pure
	{inputs: map[string]int{"plume_fleche": 5, "capuche_brigand": 3}, artefactName: "Puce de Visée", description: "+Précision +Crit"},
}

// Vérifie si les ressources nécessaires sont présentes
func cauldronHasInputs(req map[string]int) bool {
	for k, v := range req {
		if playerInventory[k] < v {
			return false
		}
	}
	return true
}

// Consomme les ressources
func cauldronConsume(req map[string]int) {
	for k, v := range req {
		playerInventory[k] -= v
	}
}

// Récupère un artefact par nom exact
func cauldronGetArtefact(name string) (Artefact, bool) { return GetArtefactParNom(name) }

func formatRecipeInputs(req map[string]int) string {
	parts := make([]string, 0, len(req))
	for k, v := range req {
		parts = append(parts, fmt.Sprintf("%d×%s", v, k))
	}
	return strings.Join(parts, ", ")
}

// Menu principal du chaudron; bloquant jusqu'à sortie
func openCauldron() {
	// Collecte des ressources distinctes utilisées dans les recettes pour un affichage condensé
	neededKeys := map[string]struct{}{}
	for _, r := range cauldronRecipes {
		for k := range r.inputs {
			neededKeys[k] = struct{}{}
		}
	}

	// Helper pour afficher l'inventaire pertinent
	printInventory := func() {
		fmt.Println("📦 Ingrédients possédés (pertinents):")
		for k := range neededKeys {
			fmt.Printf(" - %s: %d\n", k, playerInventory[k])
		}
	}

	filterCraftable := false

	// Helper: construit et affiche la liste des recettes de façon épurée
	printRecipes := func() {
		// Calcul largeur max des noms d'artefacts pour alignement
		maxName := 0
		for _, r := range cauldronRecipes {
			if l := len(r.artefactName); l > maxName {
				maxName = l
			}
		}
		fmt.Println("\n🧪 === CHAUDRON ALCHEMIQUE ===")
		fmt.Println("(Entrer numéro pour forger | f = filtre craftables | m <num> = détails manquants | q = quitter)")
		printInventory()
		fmt.Println("\nRecettes:")
		for i, r := range cauldronRecipes {
			owned := PossedeArtefact(&currentPlayer, r.artefactName)
			has := cauldronHasInputs(r.inputs)
			if filterCraftable && (!has || owned) { // filtrer ceux qu'on ne peut pas (ou déjà possédés)
				continue
			}
			status := "❌"
			if has {
				status = "✅"
			}
			if owned {
				status = "✔️"
			}
			// Ligne épurée: index) Artefact  [status]  (si non craftable: liste courte manquants)
			shortMissing := ""
			if !has && !owned {
				// construire liste des manquants sous forme k: besoin-restant
				parts := []string{}
				for k, v := range r.inputs {
					have := playerInventory[k]
					if have < v {
						parts = append(parts, fmt.Sprintf("%s:%d", k, v-have))
					}
				}
				shortMissing = strings.Join(parts, ",")
				if shortMissing != "" {
					shortMissing = " - manquants: " + shortMissing
				}
			}
			// Format manquants: insère espaces après virgules pour lisibilité
			if strings.Contains(shortMissing, ",") {
				shortMissing = strings.ReplaceAll(shortMissing, ",", ", ")
			}
			fmt.Printf("%2d) %-*s  [%s]%s\n", i+1, maxName, r.artefactName, status, shortMissing)
		}
	}

	// Détails manquants complets
	showDetailedMissing := func(idx int) {
		if idx < 0 || idx >= len(cauldronRecipes) {
			return
		}
		r := cauldronRecipes[idx]
		if PossedeArtefact(&currentPlayer, r.artefactName) {
			fmt.Println("✔️ Déjà possédé.")
			return
		}
		fmt.Printf("🔍 Détails ingrédients pour %s:\n", r.artefactName)
		for k, v := range r.inputs {
			have := playerInventory[k]
			need := v
			status := "OK"
			if have < need {
				status = fmt.Sprintf("manque %d", need-have)
			}
			fmt.Printf(" - %-18s %d/%d (%s)\n", k, have, need, status)
		}
		fmt.Println("(Entrer numéro pour forger, f pour filtrer, q pour quitter)")
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		printRecipes()
		fmt.Print("> ")
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(strings.ToLower(line))
		if line == "q" {
			break
		}
		if line == "f" {
			filterCraftable = !filterCraftable
			continue
		}
		if strings.HasPrefix(line, "m ") { // m <num>
			var which int
			fmt.Sscanf(line[2:], "%d", &which)
			if which >= 1 && which <= len(cauldronRecipes) {
				showDetailedMissing(which - 1)
			} else {
				fmt.Println("Index invalide pour m <num>.")
			}
			continue
		}
		// Essayer parse numéro direct
		idx := -1
		fmt.Sscanf(line, "%d", &idx)
		if idx < 1 || idx > len(cauldronRecipes) {
			fmt.Println("Entrée inconnue.")
			continue
		}
		rec := cauldronRecipes[idx-1]
		if PossedeArtefact(&currentPlayer, rec.artefactName) {
			fmt.Println("Vous possédez déjà cet artefact.")
			continue
		}
		if !cauldronHasInputs(rec.inputs) {
			fmt.Println("Ingrédients insuffisants (utilisez m", idx, "pour voir les détails).")
			continue
		}
		art, ok := cauldronGetArtefact(rec.artefactName)
		if !ok {
			fmt.Println("Artefact introuvable (config)")
			continue
		}
		cauldronConsume(rec.inputs)
		AjouterArtefactPossede(&currentPlayer, art)
		equipped := false
		for slot := 0; slot < MaxArtefactsEquipes; slot++ {
			if slot >= len(currentPlayer.ArtefactsEquipes) || currentPlayer.ArtefactsEquipes[slot] == nil {
				EquiperArtefactDansSlot(&currentPlayer, art, slot)
				equipped = true
				break
			}
		}
		if equipped {
			fmt.Printf("✨ Artefact forgé et équipé: %s !\n", art.Nom)
		} else {
			fmt.Printf("✨ Artefact forgé: %s (ajouté à la collection).\n", art.Nom)
		}
	}
	fmt.Println("Fermeture du chaudron.")
}

// Séquence cheat: UP UP DOWN DOWN A B A B
var cheatSequence = []string{"up", "up", "down", "down", "a", "b", "a", "b"}
var cheatProgress int

// Traite un événement clavier pour la séquence; retourne true si séquence complétée
func processCheatSequence(e keyboard.KeyEvent) bool {
	// Normaliser entrée
	key := ""
	if e.Key == keyboard.KeyArrowUp {
		key = "up"
	} else if e.Key == keyboard.KeyArrowDown {
		key = "down"
	} else if e.Rune == 'a' || e.Rune == 'A' {
		key = "a"
	} else if e.Rune == 'b' || e.Rune == 'B' {
		key = "b"
	} else {
		// Toute autre touche reset partiel si pas vide
		if cheatProgress != 0 {
			cheatProgress = 0
		}
		return false
	}

	// Vérifier progression
	if key == cheatSequence[cheatProgress] {
		cheatProgress++
		if cheatProgress == len(cheatSequence) {
			cheatProgress = 0
			return true
		}
	} else {
		// Reset si la touche ne correspond pas
		cheatProgress = 0
		// Re-vérifier si cette touche est le début potentiel
		if key == cheatSequence[0] {
			cheatProgress = 1
		}
	}
	return false
}

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
	"salle11": salle11, // Salle PNJ soins
	"salle12": salle12, // Nouvelle salle 12
	"salle13": salle13, // Nouvelle salle 13
	"salle14": salle14, // Nouvelle salle 14
	"salle15": salle15, // Nouvelle salle 15
}

// Map pour suivre l'état des coffres ouverts
var chestOpened = make(map[string]bool)

// Map pour suivre l'état des coffres secrets ouverts
var secretChestsOpened = make(map[string]bool)

// Buffer de messages HUD à afficher sous la ligne de déplacement
var hudMessages []string

// État persistant : pierre bloquante de salle1 déjà détruite ?
var stoneBroken bool

// Ajoute un message au HUD (limite optionnelle pour éviter l'accumulation)
func addHUDMessage(msg string) {
	if len(strings.TrimSpace(msg)) == 0 {
		return
	}
	hudMessages = append(hudMessages, msg)
	if len(hudMessages) > 5 { // garder seulement les 5 derniers
		hudMessages = hudMessages[len(hudMessages)-5:]
	}
}

// Vide les messages HUD
func clearHUDMessages() {
	hudMessages = hudMessages[:0]
}

// Canal global pour le clavier, réutilisé par le combat pour éviter les conflits d'entrée
var globalKeyEvents <-chan keyboard.KeyEvent

// Flag pour demander un rechargement de la carte après utilisation du cheat menu
var cheatReloadRequested bool

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

	// Si on est en salle1 et que la pierre a été cassée, la retirer si présente
	if currentMap == "salle1" && stoneBroken {
		for y := 0; y < len(mapData); y++ {
			for x := 0; x < len(mapData[y]); x++ {
				if mapData[y][x] == 35 {
					mapData[y][x] = 0
				}
			}
		}
	}

	// Salle3: à chaque entrée, on régénère un set d'ennemis aléatoires
	if currentMap == "salle3" {
		// Nettoyer les ennemis existants (laisser PNJ et autres éléments intacts)
		for y := 0; y < len(mapData); y++ {
			for x := 0; x < len(mapData[y]); x++ {
				if mapData[y][x] == 2 {
					mapData[y][x] = 0
				}
			}
		}
		generateRandomMobs(mapData)
	} else if currentMap == "salle2" {
		// Comme salle3/salle9: à chaque entrée, nettoyer et régénérer aléatoirement 4 ennemis
		for y := 0; y < len(mapData); y++ {
			for x := 0; x < len(mapData[y]); x++ {
				if mapData[y][x] == 2 {
					mapData[y][x] = 0
				}
			}
		}
		generateRandomMobsSalle2(mapData)
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

	// Systèmes mini boss / boss personnalisés
	if currentMap == "salle12" {
		applySalle12MiniBossSystem(mapData)
	}
	if currentMap == "salle13" || currentMap == "salle14" || currentMap == "salle15" {
		applyGenericBossRoom(currentMap, mapData)
	}
}

// Vérifie si le joueur a vaincu le monstre obligatoire de la salle1 (coordonnées 8,3)
// On considère la salle "libérée" si l'ennemi est marqué vaincu OU transformé en PNJ.
func canLeaveSalle1() bool {
	key := "8_3"
	if defeatedMap, ok := enemiesDefeated["salle1"]; ok {
		if defeatedMap[key] { // Ennemi vaincu (disparu ou PNJ)
			return true
		}
	}
	if trMap, ok := pnjTransformed["salle1"]; ok {
		if trMap[key] { // Transformé en PNJ (cas spécial)
			return true
		}
	}
	return false
}

// Gère les interactions avec les différents types de cases
func handleCellInteraction(cell int, currentMap string, newX, newY int, mapData [][]int, px, py int) (bool, string) {
	switch cell {
	case 80: // Arbre
		// Si le joueur a une hache, couper l'arbre et avancer
		if playerInventory["hache"] > 0 {
			// Efface l'arbre côté persistance décor (si suivi) et sur la carte en mémoire
			RemoveTree(currentMap, newX, newY)
			MarkTreeCut(currentMap, newX, newY)
			mapData[newY][newX] = 0
			// Déplacer le joueur
			mapData[py][px] = 0
			mapData[newY][newX] = 1
			addHUDMessage("🪓 Vous coupez l'arbre. Le passage est libre.")
			return false, currentMap
		}
		// Sinon, bloqué comme un mur
		return false, currentMap
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
		addHUDMessage("⚔️ Vous rencontrez une créature maudite !")
		isSuper := (cell == 12)
		result := combat(currentMap, isSuper)

		enemyKey := fmt.Sprintf("%d_%d", newX, newY)

		// Si le joueur est mort (PV <= 0), régénérer, appliquer la perte de pièces et demander une transition vers salle1
		if currentPlayer.PV <= 0 {
			loss := playerInventory["pièces"] * 35 / 100
			if loss > 0 {
				playerInventory["pièces"] -= loss
				fmt.Printf("☠️ Vous êtes mort. Vous perdez %d pièces (35%%).\n", loss)
			} else {
				fmt.Println("☠️ Vous êtes mort.")
			}

			// Régénérer le personnage: PV = PVMax effectif avec armure équipée (sans modifier la base)
			tmp := currentPlayer
			_ = EquiperArmure(&tmp, tmp.ArmuresDisponibles)
			// tmp.PVMax inclut le bonus d'armure; heal complet
			currentPlayer.PV = tmp.PVMax
			fmt.Println("↩️  Retour à la salle 1 (spawn). Vous êtes régénéré.")
			// Demander une transition vers salle1, l'emplacement précis sera géré dans RunGameLoop
			return true, "salle1"
		}

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
				// Récompense unique : clé ET pioche (si pas déjà donnée)
				if playerInventory["pioche"] == 0 {
					addToInventory("pioche", 1)
					addHUDMessage("🪓 Vous obtenez une PIÔCHE ! Elle peut briser la pierre sacrée (๑).")
				}
				addHUDMessage("🤝 Le Mentor Maudit est libéré : il devient le Mentor Suprême !")
				showDialogue(currentMap, newX, newY)
			} else {
				enemiesDefeated[currentMap][enemyKey] = true
				mapData[py][px] = 0
				mapData[newY][newX] = 1
				addHUDMessage("✅ Ennemi vaincu. Passage dégagé.")
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
				if playerInventory["pioche"] == 0 {
					addToInventory("pioche", 1)
					addHUDMessage("🪓 Vous obtenez une PIÔCHE ! Elle peut briser la pierre sacrée (๑).")
				}
				addHUDMessage("🤝 Le Mentor Maudit est libéré : il devient le Mentor Suprême !")
				showDialogue(currentMap, newX, newY)
			} else {
				// Tous les autres mobs ne se transforment plus jamais
				enemiesDefeated[currentMap][enemyKey] = true
				mapData[py][px] = 0
				mapData[newY][newX] = 1
				addHUDMessage("✅ Ennemi vaincu. Passage dégagé.")
			}
		}
		return false, currentMap

	case 35: // Pierre sacrée bloquante devant la porte de salle1
		if currentMap != "salle1" {
			return false, currentMap
		}
		// Si déjà détruite (sécurité) : traiter comme sol
		if stoneBroken {
			mapData[newY][newX] = 0
			return false, currentMap
		}
		if !canLeaveSalle1() {
			addHUDMessage("🪨 La pierre est incassable tant que le monstre n'est pas vaincu.")
			return false, currentMap
		}
		if playerInventory["pioche"] == 0 {
			addHUDMessage("🪨 Il vous faut une pioche pour briser cette pierre (๑).")
			return false, currentMap
		}
		playExplosion(mapData, newX, newY)
		addHUDMessage("💥 La pierre se désintègre dans une explosion ! Passage libre.")
		stoneBroken = true
		return false, currentMap
	case 67: // Mini boss salle12
		if currentMap == "salle12" {
			addHUDMessage("🛡️ Mini-Boss !")
			// Combat super -> après combat on applique déjà scaling; on pourrait injecter un buff avant
			// (Simplification: on laisse CreateRandomEnemyForMap + isSuper pour PV/crit de base)
			result := combat(currentMap, true)
			if currentPlayer.PV <= 0 {
				loss := playerInventory["pièces"] * 35 / 100
				if loss > 0 {
					playerInventory["pièces"] -= loss
				}
				// respawn salle1
				return true, "salle1"
			}
			if result == true || result == "disappear" { // vaincu
				key := fmt.Sprintf("%d_%d", newX, newY)
				salle12BossState.defeatedMini[key] = true
				mapData[py][px] = 0
				mapData[newY][newX] = 1
				// Ajouter fragment inventaire
				addToInventory("fragment_spawn", 1)
				salle12BossState.fragments++
				addHUDMessage(fmt.Sprintf("🧩 Fragment obtenu (%d/4)", salle12BossState.fragments))
				if salle12BossState.fragments >= 4 && !salle12BossState.spawnerSpawn {
					addHUDMessage("⚙️ Un spawner apparaît au centre !")
					// sera placé à la prochaine application d'état (ou immédiatement)
					cx, cy := salle12Center[0], salle12Center[1]
					if mapData[cy][cx] != 1 {
						mapData[cy][cx] = 66
					}
					salle12BossState.spawnerSpawn = true
				}
			}
			return false, currentMap
		}
		return false, currentMap

	case 66: // Spawner (interagir pour invoquer boss si fragments OK)
		if currentMap == "salle12" && salle12BossState.spawnerSpawn && !salle12BossState.bossDefeated {
			if salle12BossState.fragments < 4 {
				addHUDMessage("❌ Il manque des fragments.")
				return false, currentMap
			}
			addHUDMessage("🔥 Le boss est invoqué !")
			mapData[newY][newX] = 68 // Boss
			return false, currentMap
		}
		return false, currentMap

	case 68: // Boss final salle12
		if currentMap == "salle12" && !salle12BossState.bossDefeated {
			addHUDMessage("👹 Boss Niveau 1/4 !")
			result := combat(currentMap, true)
			if currentPlayer.PV <= 0 {
				loss := playerInventory["pièces"] * 35 / 100
				if loss > 0 {
					playerInventory["pièces"] -= loss
				}
				return true, "salle1"
			}
			if result == true || result == "disappear" {
				addHUDMessage("🏆 Boss vaincu !")
				salle12BossState.bossDefeated = true
				mapData[newY][newX] = 1
				mapData[py][px] = 0
				addToInventory("pièces", 50)
			}
			return false, currentMap
		}
		return false, currentMap

	// --- Générique salles 13-15 ---
	case 70, 73, 76: // Mini boss niveaux 2,3,4
		if cfg, ok := bossRooms[currentMap]; ok {
			st := &cfg.state
			// Déterminer si case correspond à mini code de cette salle
			if cell == cfg.codeMini {
				addHUDMessage(fmt.Sprintf("🛡️ Mini-Boss Niveau %d !", cfg.level))
				result := combat(currentMap, true)
				if currentPlayer.PV <= 0 {
					loss := playerInventory["pièces"] * 35 / 100
					if loss > 0 {
						playerInventory["pièces"] -= loss
					}
					return true, "salle1"
				}
				if result == true || result == "disappear" {
					key := fmt.Sprintf("%d_%d", newX, newY)
					st.defeatedMini[key] = true
					mapData[py][px] = 0
					mapData[newY][newX] = 1
					addToInventory("fragment_spawn", 1)
					st.fragments++
					addHUDMessage(fmt.Sprintf("🧩 Fragment obtenu (%d/4)", st.fragments))
					if st.fragments >= 4 && !st.spawnerSpawn {
						addHUDMessage("⚙️ Un spawner apparaît au centre !")
						cx, cy := cfg.center[0], cfg.center[1]
						if mapData[cy][cx] != 1 {
							mapData[cy][cx] = cfg.codeSpawn
						}
						st.spawnerSpawn = true
					}
				}
			}
		}
		return false, currentMap

	case 71, 74, 77: // Spawner niveaux 2,3,4
		if cfg, ok := bossRooms[currentMap]; ok {
			st := &cfg.state
			if cell == cfg.codeSpawn && st.spawnerSpawn && !st.bossDefeated {
				if st.fragments < 4 {
					addHUDMessage("❌ Il manque des fragments.")
					return false, currentMap
				}
				addHUDMessage("🔥 Le boss est invoqué !")
				mapData[newY][newX] = cfg.codeBoss
			}
		}
		return false, currentMap

	case 72, 75, 78: // Boss niveaux 2,3,4
		if cfg, ok := bossRooms[currentMap]; ok {
			st := &cfg.state
			if cell == cfg.codeBoss && !st.bossDefeated {
				addHUDMessage(fmt.Sprintf("👹 Boss Niveau %d/4 !", cfg.level))
				result := combat(currentMap, true)
				if currentPlayer.PV <= 0 {
					loss := playerInventory["pièces"] * 35 / 100
					if loss > 0 {
						playerInventory["pièces"] -= loss
					}
					return true, "salle1"
				}
				if result == true || result == "disappear" {
					addHUDMessage("🏆 Boss vaincu !")
					st.bossDefeated = true
					mapData[newY][newX] = 1
					mapData[py][px] = 0
					// Récompense progressive
					reward := 50 + (cfg.level-1)*25
					addToInventory("pièces", reward)
					addHUDMessage(fmt.Sprintf("💰 Vous gagnez %d pièces !", reward))
					// Drop spécial boss final salle15
					if currentMap == "salle15" && cfg.level == 4 {
						addToInventory("sida", 1)
						addHUDMessage("🧬 Vous obtenez l'objet mystérieux 'sida'. Apportez-le au PNJ de la salle1...")
					}
				}
			}
		}
		return false, currentMap

	case 79: // Chaudron alchimique (sorcier)
		// Petit dialogue avant d'ouvrir l'interface du chaudron
		readKey := func() rune {
			if globalKeyEvents == nil { // fallback minimal
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
			if r >= 'A' && r <= 'Z' {
				r += 32
			}
			return r
		}

		fmt.Println("\n💬 === SORCIER ALCHEMISTE ===")
		fmt.Println("🧪 Sorcier: Approche, aventurier. Mon chaudron renferme des secrets anciens.")
		fmt.Println("🧪 Sorcier: Avec les bonnes composantes, je peux forger des artefacts puissants.")
		fmt.Println("🧪 Sorcier: Veux-tu accéder au chaudron ? (o/n)")
		fmt.Print("> ")
		ans := readKey()
		if ans == 'o' {
			fmt.Println("Sorcier: Très bien, contemple la transmutation...")
			openCauldron()
		} else {
			fmt.Println("Sorcier: Reviens quand tu auras plus de composants.")
		}
		fmt.Println("(Appuyez sur une touche pour continuer)")
		_ = readKey()
		return false, currentMap

	case 3: // PNJ
		fmt.Println("Vous parlez au PNJ...")
		showDialogue(currentMap, newX, newY)
		return false, currentMap

	case 30: // Porte spéciale vers la salle secrète dans salle1
		if currentMap == "salle1" {
			// Empêcher de quitter la salle1 avant d'avoir vaincu le monstre clé
			if !canLeaveSalle1() {
				addHUDMessage("⚠️ Une force mystérieuse vous empêche de partir... Vainquez d'abord le monstre de la salle !")
				return false, currentMap
			}
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

	// Toutes les autres portes - gestion unifiée (inclut nouvelles portes 50/51/52/53)
	case 7, 10, 13, 14, 15, 20, 21, 27, 28, 31, 33, 34, 38, 40, 42, 44, 50, 51, 52, 53, 54, 55, 56, 57:
		if tr, ok := transitions[currentMap][cell]; ok {
			if currentMap == "salle1" && !canLeaveSalle1() {
				addHUDMessage("⚠️ Vous ne pouvez pas encore quitter la salle1. Le monstre n'a pas été vaincu !")
				return false, currentMap
			}
			// Gating progression boss successifs
			if cell == 52 { // vers salle13
				if !salle12BossState.bossDefeated {
					addHUDMessage("⛔ La porte est scellée. Vainquez d'abord le Boss Niveau 1/4.")
					return false, currentMap
				}
			}
			if cell == 54 { // vers salle14
				if st, ok2 := bossRooms["salle13"]; ok2 {
					if !st.state.bossDefeated {
						addHUDMessage("⛔ La porte est scellée. Vainquez d'abord le Boss Niveau 2/4.")
						return false, currentMap
					}
				}
			}
			if cell == 56 { // vers salle15
				if st, ok2 := bossRooms["salle14"]; ok2 {
					if !st.state.bossDefeated {
						addHUDMessage("⛔ La porte est scellée. Vainquez d'abord le Boss Niveau 3/4.")
						return false, currentMap
					}
				}
			}
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
func getPlayerMovement(events <-chan keyboard.KeyEvent, px, py int) (int, int, bool) {
	fmt.Print("Déplacez-vous (ZQSD pour bouger, I=Inventaire, X=Quitter): ")

	// Afficher les messages HUD accumulés juste sous la ligne de déplacement
	if len(hudMessages) > 0 {
		fmt.Println()
		for _, m := range hudMessages {
			fmt.Println(m)
		}
		clearHUDMessages()
		fmt.Print("→ ")
	}

	// Lire le prochain événement dispo puis drainer rapidement les répétitions bufferisées,
	// pour que le dernier input (ex: gauche) prenne le dessus sur les répétitions (ex: droite).
	e := <-events
	draining := true
	for draining {
		select {
		case next := <-events:
			e = next
		default:
			draining = false
		}
	}

	input := strings.ToLower(string(e.Rune))

	// Ouverture via séquence (UP UP DOWN DOWN A B A B)
	if debugMode {
		if processCheatSequence(e) {
			showCheatMenu(&currentMapGlobalRef, &mapDataGlobalRef)
			// Appliquer les états ennemis sur la carte potentiellement modifiée par les cheats
			applyEnemyStates(mapDataGlobalRef, currentMapGlobalRef)
			// Demander au run loop d'adopter les nouvelles refs (TP, placements, etc.)
			cheatReloadRequested = true
			return px, py, true
		}
	}
	key := e.Key

	newX, newY := px, py
	switch {
	case input == "z" || key == keyboard.KeyArrowUp:
		newY = py - 1
	case input == "s" || key == keyboard.KeyArrowDown:
		newY = py + 1
	case input == "q" || key == keyboard.KeyArrowLeft:
		newX = px - 1
	case input == "d" || key == keyboard.KeyArrowRight:
		newX = px + 1
	case input == "i":
		showInventoryMenu(events)
		return px, py, true
	case input == "x":
		fmt.Println("Vous quittez la partie. Merci d'avoir joué !")
		return px, py, false
	default:
		// touche ignorée
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
	// Les arbres sont bloquants sauf si le joueur possède une hache
	if mapData[y][x] == 80 { // arbre
		if playerInventory["hache"] <= 0 {
			return false
		}
		// autoriser le mouvement: l'abattage sera géré dans handleCellInteraction
	}
	return true
}

// Boucle principale du jeu refactorisée
// Références globales temporaires nécessaires à l'appel cheat depuis getPlayerMovement
var currentMapGlobalRef string
var mapDataGlobalRef [][]int

func RunGameLoop(currentMap string) {
	// reader removed: using keyboard events for movement
	mapData := copyMap(salles[currentMap])
	// Appliquer les décorations (arbres, etc.) sur la carte chargée
	applyDecorations(currentMap, mapData)
	// Puis retirer les arbres déjà coupés (persistance)
	applyCutTrees(currentMap, mapData)

	// Auto-équipe la Lunette d'Erwann si la classe est Erwann
	if currentPlayer.Nom == "Erwann" {
		if !PossedeArtefact(&currentPlayer, "Lunette d'Erwann") {
			if a, ok := GetArtefactParNom("Lunette d'Erwann"); ok {
				EquiperArtefactDansSlot(&currentPlayer, a, 0)
				addHUDMessage("🎯 Lunette d'Erwann équipée automatiquement.")
			}
		}
	}

	// Auto-équipe le Halo de Gabriel si la classe est Gabriel
	if currentPlayer.Nom == "Gabriel" {
		if !PossedeArtefact(&currentPlayer, "Halo de Gabriel") {
			if a, ok := GetArtefactParNom("Halo de Gabriel"); ok {
				EquiperArtefactDansSlot(&currentPlayer, a, 0)
				addHUDMessage("✨ Halo de Gabriel équipé automatiquement.")
			}
		}
	}

	// Auto-équipe la Vodka de Vitaly si la classe est Vitaly
	if currentPlayer.Nom == "Vitaly" {
		if !PossedeArtefact(&currentPlayer, "Vodka de Vitaly") {
			if a, ok := GetArtefactParNom("Vodka de Vitaly"); ok {
				EquiperArtefactDansSlot(&currentPlayer, a, 0)
				addHUDMessage("🍾 Vodka de Vitaly équipée automatiquement.")
			}
		}
	}

	// Ouvrir une fois le clavier et récupérer un canal d'événements pour tout le loop
	if err := keyboard.Open(); err != nil {
		fmt.Println("Erreur d'accès clavier:", err)
		return
	}
	defer keyboard.Close()
	events, err := keyboard.GetKeys(10)
	if err != nil {
		fmt.Println("Erreur d'initialisation du clavier:", err)
		return
	}
	// Expose le canal aux autres modules (combat) tant que le jeu est ouvert
	globalKeyEvents = events

	// Applique l'état des ennemis vaincus
	applyEnemyStates(mapData, currentMap)

	// Assure que le joueur est présent dans la carte initiale
	if px, py := findPlayer(mapData); px == -1 || py == -1 {
		placePlayerAt(mapData, len(mapData[0])/2, len(mapData)/2)
	}

	// Met à jour les références globales utilisées par getPlayerMovement pour le cheat menu
	currentMapGlobalRef = currentMap
	mapDataGlobalRef = mapData

	for {
		assignEnemyEmojis(currentMap, mapData)
		printMap(mapData) // Le HUD est maintenant intégré dans printMap
		fmt.Printf("📍 Salle actuelle: %s\n", currentMap)

		px, py := findPlayer(mapData)
		if px == -1 || py == -1 {
			fmt.Println("Erreur: Joueur non trouvé dans la carte!")
			return
		}

		// (Hook cheat déplacé dans getPlayerMovement)

		// Utiliser la nouvelle fonction de mouvement (lecture instantanée via canal) si pas déjà traité
		newX, newY, shouldContinue := getPlayerMovement(events, px, py)

		if !shouldContinue {
			return // Le joueur a quitté le jeu
		}

		// Si le cheat menu a modifié la carte/salle, adopter les refs globales
		if cheatReloadRequested {
			currentMap = currentMapGlobalRef
			mapData = mapDataGlobalRef
			cheatReloadRequested = false
			// Recommencer l'itération pour afficher la nouvelle carte
			continue
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
			// Ré-appliquer les décorations pour la nouvelle carte puis retirer les arbres coupés
			applyDecorations(currentMap, mapData)
			applyCutTrees(currentMap, mapData)
			currentMapDisplayName = currentMap
			assignEnemyEmojis(currentMap, mapData)
			applyEnemyStates(mapData, currentMap)
			currentMapGlobalRef = currentMap
			mapDataGlobalRef = mapData

			// Placer le joueur selon la transition
			if currentMap == "salle8" { // Salle secrète => position spéciale inchangée
				placePlayerAt(mapData, 3, 6)
			} else if currentMap == "salle1" {
				// Deux cas: retour suite à mort (prevMap != salle1 via handleCellInteraction) ou simple transition (porte)
				if haveSpawn { // Transition via porte -> utiliser spawn défini dans transitions
					placePlayerAt(mapData, spawnX, spawnY)
				} else {
					// Respawn (mort ou cas sans spawn explicite) -> position fixe (8,5)
					placePlayerAt(mapData, 8, 5)
				}
			} else if haveSpawn {
				placePlayerAt(mapData, spawnX, spawnY)
			} else {
				placePlayerAt(mapData, len(mapData[0])/2, len(mapData)/2)
			}
			currentMapGlobalRef = currentMap
			mapDataGlobalRef = mapData
		}
	}
}

// Affiche une courte animation d'explosion pour la pierre (utilise les codes 46 et 47)
// Bloque environ 250ms au total – suffisamment court pour ne pas gêner le gameplay.
func playExplosion(mapData [][]int, x, y int) {
	frames := []int{46, 47}
	for i, f := range frames {
		mapData[y][x] = f
		printMap(mapData)
		fmt.Printf("📍 Salle actuelle: %s\n", "salle1")
		if i == 0 {
			time.Sleep(120 * time.Millisecond)
		} else {
			time.Sleep(150 * time.Millisecond)
		}
	}
	// Nettoie la case (sol vide)
	mapData[y][x] = 0
}
