package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/eiannone/keyboard"
)

// État persistant du filtre artefacts (true = masquer non possédés)
var artefactHideNotOwned bool
var hideZeroConsumables bool // filtre consommables (true = masquer ceux à 0)

// showInventoryMenu affiche un écran d'inventaire détaillé. On reste dans cette vue
// tant que l'utilisateur ne presse pas I, ESC ou Entrée.
func showInventoryMenu(events <-chan keyboard.KeyEvent) {
	for {
		printInventoryScreen()
		// Attendre une touche utilisateur (drain répétitions)
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
		if input == "i" || e.Key == keyboard.KeyEsc || e.Key == keyboard.KeyEnter {
			return
		}
		if input == "a" { // ouvrir sous-menu artefacts
			showArtefactMenu(events)
		}
		if input == "f" { // toggle filtre consommables
			hideZeroConsumables = !hideZeroConsumables
			continue
		}
	}
}

func printInventoryScreen() {
	fmt.Print("\033[H\033[2J")
	fmt.Println("════════════════════════════════════════════════════════════")
	fmt.Println("🎒                    INVENTAIRE DU JOUEUR                   🎒")
	fmt.Println("════════════════════════════════════════════════════════════")

	// Section ressources & clés
	fmt.Println("📦 RESSOURCES")
	fmt.Printf("   💰 Pièces:          %d\n", playerInventory["pièces"])
	fmt.Printf("   🔑 Clés:            %d\n", playerInventory["clés"])
	fmt.Printf("   🗝️ Clés spéciales:  %d\n", playerInventory["clés_spéciales"])
	fmt.Printf("   ⛏️ Pioche:          %d\n", playerInventory["pioche"])
	fmt.Printf("   🪨 Roches évol.:    %d\n", currentPlayer.Roches)
	fmt.Println()

	// Section consommables
	fmt.Print("🧪 CONSOMMABLES")
	if hideZeroConsumables {
		fmt.Print(" (filtre actifs F pour tout)")
	} else {
		fmt.Print(" (F pour masquer ceux à 0)")
	}
	fmt.Println()
	printConsumable := func(label string, key string) {
		q := playerInventory[key]
		if hideZeroConsumables && q == 0 {
			return
		}
		fmt.Printf("   %-20s %d\n", label+":", q)
	}
	printConsumable("Potion (legacy +70)", "potions")
	printConsumable("Potion mineure", "potion_mineure")
	printConsumable("Potion majeure", "potion_majeure")
	printConsumable("Potion suprême", "potion_supreme")
	printConsumable("Antidote", "antidote")
	printConsumable("Potion de dégâts", "potion_degats")
	printConsumable("Bombe incendiaire", "bombe_incendiaire")
	printConsumable("Bombe givrante", "bombe_givrante")
	printConsumable("Grenade fumigène", "grenade_fumigene")
	printConsumable("Parchemin dispersion", "parchemin_dispersion")
	printConsumable("Élixir de force", "elixir_force")
	printConsumable("Élixir de vitesse", "elixir_vitesse")
	printConsumable("Élixir de précision", "elixir_critique")
	printConsumable("Puff 9K", "puff_9k")
	printConsumable("Vodka de Vitaly", "vodka_vitaly")
	fmt.Println()

	// Arme & Armure équipées (vue rapide)
	weaponName := "Aucune"
	if currentPlayer.ArmeEquipee.Nom != "" {
		weaponName = currentPlayer.ArmeEquipee.Nom
	} else if currentPlayer.NiveauArme >= 0 && currentPlayer.NiveauArme < len(currentPlayer.ArmesDisponibles) {
		weaponName = currentPlayer.ArmesDisponibles[currentPlayer.NiveauArme].Nom
	}
	armorName := "Aucune"
	if currentPlayer.ArmureEquipee.Nom != "" {
		armorName = currentPlayer.ArmureEquipee.Nom
	} else if currentPlayer.NiveauArmure >= 0 && currentPlayer.NiveauArmure < len(currentPlayer.ArmuresDisponibles) {
		armorName = currentPlayer.ArmuresDisponibles[currentPlayer.NiveauArmure].Nom
	}
	fmt.Println("🛡️ ÉQUIPEMENT")
	fmt.Printf("   Arme actuelle:      %s\n", weaponName)
	fmt.Printf("   Armure actuelle:    %s\n", armorName)
	if playerStats.hasLegendaryWeapon {
		fmt.Println("   🌟 Arme légendaire active: Excalibur")
	}
	fmt.Println()

	// Ligne des multiplicateurs actifs (valeurs absolues, 100% = base)
	modDeg, modMag, modPrec, modCrit, modArm, modResM := CalculerModificateurs(&currentPlayer)
	colorAbs := func(v float64) string {
		pct := v * 100
		if v > 1.0001 {
			return fmt.Sprintf("\033[32m%.0f%%\033[0m", pct)
		} else if v < 0.9999 {
			return fmt.Sprintf("\033[31m%.0f%%\033[0m", pct)
		}
		return fmt.Sprintf("%.0f%%", pct)
	}
	fmt.Printf("Multiplicateurs actifs: ATK %s | MATK %s | PREC %s | CRIT %s | ARM %s | RESM %s\n",
		colorAbs(modDeg), colorAbs(modMag), colorAbs(modPrec), colorAbs(modCrit), colorAbs(modArm), colorAbs(modResM))
	fmt.Println()

	// Artefacts équipés
	fmt.Println("🧿 ARTEFACTS")
	artNames := []string{}
	for _, a := range currentPlayer.ArtefactsEquipes {
		if a != nil {
			artNames = append(artNames, a.Nom)
		}
	}
	if len(artNames) == 0 {
		fmt.Println("   Aucun artefact équipé")
	} else {
		// Tri pour affichage stable
		sort.Strings(artNames)
		for _, n := range artNames {
			fmt.Printf("   • %s\n", n)
		}
	}
	fmt.Println()

	fmt.Printf("☠️  Ennemis tués au total: %d\n", playerStats.enemiesKilled)

	// Section Loots spécifiques (affiche seulement ceux possédés)
	fmt.Println()
	fmt.Println("🧬 LOOTS SPÉCIFIQUES")
	keys := []struct{ key, label string }{
		{"dent_rat", "Dent de Rat"},
		{"dent_rat_luisante", "Dent de Rat Luisante"},
		{"gelée_visqueuse", "Gelée Visqueuse"},
		{"coeur_de_gelée", "Cœur de Gelée"},
		{"capuche_brigand", "Capuche de Brigand"},
		{"dague_ensorcelée", "Dague Ensorcelée"},
		{"plume_fleche", "Plume de Flèche"},
		{"carquois_gravé", "Carquois Gravé"},
		{"cendre_infernale", "Cendre Infernale"},
		{"braise_eternelle", "Braise Éternelle"},
		{"insigne_chevalier", "Insigne de Chevalier"},
		{"lame_ancient", "Lame Ancienne"},
		{"sang_berserker", "Sang de Berserker"},
		{"talisman_fureur", "Talisman de Fureur"},
		{"essence_sombre", "Essence Sombre"},
		{"noyau_occulte", "Noyau Occulte"},
		{"corne_demon", "Corne de Démon"},
		{"fragment_demoniaque", "Fragment Démoniaque"},
		{"parchemin_arcane", "Parchemin Arcane"},
		{"sceau_archimage", "Sceau d'Archimage"},
		{"embleme_champion", "Emblème de Champion"},
		{"aiguille_du_destin", "Aiguille du Destin"},
	}
	shown := 0
	for _, it := range keys {
		if q := playerInventory[it.key]; q > 0 {
			fmt.Printf("   %-22s %d\n", it.label+":", q)
			shown++
		}
	}
	if shown == 0 {
		fmt.Println("   (Aucun pour l'instant)")
	}
	fmt.Println("════════════════════════════════════════════════════════════")
	fmt.Println("(I/Esc/Entrée = retour | A = artefacts | F = filtre consommables)")
}

// ------------------------- Sous-menu Artefacts -------------------------

func showArtefactMenu(events <-chan keyboard.KeyEvent) {
	for {
		letters := printArtefactScreen(&artefactHideNotOwned)
		e := <-events
		input := strings.ToLower(string(e.Rune))
		if e.Key == keyboard.KeyEsc || input == "r" || input == "i" || e.Key == keyboard.KeyEnter {
			return
		}
		if input == "f" { // toggle filtre
			artefactHideNotOwned = !artefactHideNotOwned
			continue
		}
		if handler, ok := letters[input]; ok {
			// toggle equip / unequip
			artefact := handler
			// déjà équipé ?
			slotIdx := -1
			for i, slot := range currentPlayer.ArtefactsEquipes {
				if slot != nil && slot.Nom == artefact.Nom {
					slotIdx = i
					break
				}
			}
			if slotIdx >= 0 { // déséquiper
				if err := DesequiperArtefactDuSlot(&currentPlayer, slotIdx); err != nil {
					fmt.Println("❌ ", err)
				} else {
					fmt.Printf("♻️  %s retiré (slot %d).\n", artefact.Nom, slotIdx)
				}
			} else { // équiper
				// vérifier possession
				possede := false
				for _, a := range currentPlayer.ArtefactsPossedes {
					if a.Nom == artefact.Nom {
						possede = true
						break
					}
				}
				if !possede {
					fmt.Println("❌ Non possédé.")
					continue
				}
				free := -1
				for i, slot := range currentPlayer.ArtefactsEquipes {
					if slot == nil {
						free = i
						break
					}
				}
				if free == -1 {
					fmt.Println("❌ Maximum (2) atteint. Déséquipez-en un.")
					continue
				}
				if err := EquiperArtefactDansSlot(&currentPlayer, artefact, free); err != nil {
					fmt.Println("❌ ", err)
				} else {
					fmt.Printf("✅ %s équipé (slot %d).\n", artefact.Nom, free)
				}
			}
		}
	}
}

func printArtefactScreen(hideNotOwned *bool) map[string]Artefact {
	fmt.Print("\033[H\033[2J")
	fmt.Println("════════════════════════════════════════════════════════════")
	fmt.Println("🧿               GESTION DES ARTEFACTS (max 2)               🧿")
	fmt.Println("════════════════════════════════════════════════════════════")
	fmt.Print("★ = Artefact équipé (en vert)")
	if hideNotOwned != nil && *hideNotOwned {
		fmt.Println(" | Filtre: seuls possédés (F pour afficher tous)")
	} else {
		fmt.Println(" | F pour masquer les non possédés")
	}

	// Construire set des possédés
	owned := map[string]bool{}
	for _, a := range currentPlayer.ArtefactsPossedes {
		owned[a.Nom] = true
	}
	// Construire map des équipés
	equipped := map[string]int{}
	for idx, slot := range currentPlayer.ArtefactsEquipes {
		if slot != nil {
			equipped[slot.Nom] = idx
		}
	}

	// Liste triée: équipés d'abord (ordre alpha), puis le reste (ordre alpha)
	all := make([]string, 0, len(ArtefactsDisponibles))
	for _, a := range ArtefactsDisponibles {
		all = append(all, a.Nom)
	}
	equippedNames := []string{}
	nonEquippedNames := []string{}
	for _, n := range all {
		if _, ok := equipped[n]; ok {
			equippedNames = append(equippedNames, n)
		} else {
			nonEquippedNames = append(nonEquippedNames, n)
		}
	}
	sort.Strings(equippedNames)
	sort.Strings(nonEquippedNames)
	names := append(equippedNames, nonEquippedNames...)

	fmt.Println("Let Nom                                | Statut        | Équipé | Effet")
	fmt.Println("------------------------------------------------------------------------")
	letterMap := map[string]Artefact{}
	// Optionnel: filtrer noms si demandé
	filtered := names
	if hideNotOwned != nil && *hideNotOwned {
		temp := []string{}
		for _, n := range names {
			if owned[n] {
				temp = append(temp, n)
			}
		}
		filtered = temp
	}
	letters := letterSequence(len(filtered))
	for i, nom := range filtered {
		art, _ := GetArtefactParNom(nom)
		letter := letters[i]
		statut := "Non possédé"
		if owned[nom] {
			statut = "Possédé"
		}
		slotTxt := "-"
		if idx, ok := equipped[nom]; ok {
			slotTxt = fmt.Sprintf("Slot %d", idx)
		}
		eff := resumeEffet(art.Effet)
		if eff == "—" { // aucun mod stat : tenter d'afficher le(s) débuff(s) contré(s)
			if alt := counterSummaryForArtefact(art.Nom); alt != "" {
				eff = alt
			}
		}
		nameCol := fmt.Sprintf("%-32s", nom)
		if _, ok := equipped[nom]; ok {
			nameCol = fmt.Sprintf("\033[32m%-32s\033[0m", "★ "+nom) // star + green
		}
		fmt.Printf("%s   %s | %-12s | %-6s | %s\n", letter, nameCol, statut, slotTxt, eff)
		letterMap[letter] = art
	}
	fmt.Println("------------------------------------------------------------------------")
	fmt.Println("Appuyez sur la lettre pour (dé)équiper | Esc = Retour")
	return letterMap
}

// Génère une séquence de lettres (a..z, puis aa, ab...) suffisante
func letterSequence(n int) []string {
	// Séquence: a..q, s..z (on saute r), puis 1..8,0, puis 9 tout à la fin, puis fallback si besoin.
	res := []string{}
	first := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q"}
	after := []string{"s", "t", "u", "v", "w", "x", "y", "z"}
	for _, l := range first {
		if len(res) >= n {
			return res
		}
		res = append(res, l)
	}
	for _, l := range after {
		if len(res) >= n {
			return res
		}
		res = append(res, l)
	}
	// chiffres (sans 9 pour le garder tout à la fin)
	digits := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for _, d := range digits {
		if len(res) >= n {
			return res
		}
		res = append(res, d)
	}
	// enfin 9
	if len(res) < n {
		res = append(res, "9")
	}
	// fallback si encore besoin
	extraAlphabet := []rune("abcdefghijklmnopqrstuvwxyz")
	idx := 0
	for len(res) < n {
		tag := fmt.Sprintf("1%c", extraAlphabet[idx%26])
		res = append(res, tag)
		idx++
	}
	return res
}

func resumeEffet(e Effet) string {
	parts := []string{}
	if e.ModifDegats != 0 {
		parts = append(parts, fmt.Sprintf("ATK %+d%%", int(e.ModifDegats*100)))
	}
	if e.ModifDegatsMag != 0 {
		parts = append(parts, fmt.Sprintf("MATK %+d%%", int(e.ModifDegatsMag*100)))
	}
	if e.ModifPrecision != 0 {
		parts = append(parts, fmt.Sprintf("PREC %+d%%", int(e.ModifPrecision*100)))
	}
	if e.ModifCritique != 0 {
		parts = append(parts, fmt.Sprintf("CRIT %+d%%", int(e.ModifCritique*100)))
	}
	if e.ModifArmure != 0 {
		parts = append(parts, fmt.Sprintf("ARM %+d%%", int(e.ModifArmure*100)))
	}
	if e.ModifResistMag != 0 {
		parts = append(parts, fmt.Sprintf("RESM %+d%%", int(e.ModifResistMag*100)))
	}
	if len(parts) == 0 {
		return "—"
	}
	return strings.Join(parts, "/")
}

// (Anciennes fonctions d'entrée texte supprimées après passage aux lettres)

// counterSummaryForArtefact retourne un court texte listant les débuffs nettoyés par un artefact utilitaire.
// On mappe par nom d'artefact -> liste courte des statuts contrés.
func counterSummaryForArtefact(nom string) string {
	m := map[string]string{
		"Amulette Anti-Poison":     "Nettoie Poison",
		"Antidote Éternel":         "Nettoie Poison",
		"Pendentif Purificateur":   "Nettoie TOUS débuffs",
		"Talisman Stoïque":         "Retire Peur/Étourdis.",
		"Totem de Refroidissement": "Retire Brûlure/Saign.",
		"Talisman Éteigneflamme":   "Retire Brûlure",
		"Sceau Hémostatique":       "Retire Saignement",
		"Pendentif de Courage":     "Retire Peur",
		"Talisman de Vigilance":    "Retire Étourdis.",
		"Sceau de Focalisation":    "Retire Nébul./Défavor.",
		"Glyphe de Bastion":        "Retire Brise-Armure",
		"Cachet de Détermination":  "Retire débuffs ATK",
	}
	return m[nom]
}
