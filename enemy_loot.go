package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
)

// Configuration simple de loot par type d'ennemi.
// Chaque ennemi donne toujours un objet thématique (commonItem) en quantité 1 (modifiée si super),
// et peut donner un objet rare supplémentaire (rareItem) avec une probabilité (rareChancePct).
// isSuper double la quantité du common et augmente la chance du rare.
type enemyLootConfig struct {
	commonItem    string
	rareItem      string
	rareChancePct int // 0-100
}

var enemyLootTable = map[string]enemyLootConfig{
	"rat":            {commonItem: "dent_rat", rareItem: "dent_rat_luisante", rareChancePct: 8},
	"gelée":          {commonItem: "gelée_visqueuse", rareItem: "coeur_de_gelée", rareChancePct: 6},
	"brigand":        {commonItem: "capuche_brigand", rareItem: "dague_ensorcelée", rareChancePct: 5},
	"archer":         {commonItem: "plume_fleche", rareItem: "carquois_gravé", rareChancePct: 5},
	"apprenti pyro":  {commonItem: "cendre_infernale", rareItem: "braise_eternelle", rareChancePct: 7},
	"chevalier":      {commonItem: "insigne_chevalier", rareItem: "lame_ancient", rareChancePct: 5},
	"berserker":      {commonItem: "sang_berserker", rareItem: "talisman_fureur", rareChancePct: 6},
	"mage sombre":    {commonItem: "essence_sombre", rareItem: "noyau_occulte", rareChancePct: 9},
	"seigneur démon": {commonItem: "corne_demon", rareItem: "fragment_demoniaque", rareChancePct: 10},
	"archimage":      {commonItem: "parchemin_arcane", rareItem: "sceau_archimage", rareChancePct: 10},
	"champion déchu": {commonItem: "embleme_champion", rareItem: "aiguille_du_destin", rareChancePct: 7},
}

var nonLetter = regexp.MustCompile(`[^a-zA-ZÀ-ÖØ-öø-ÿ\s]`)

// normalise un nom d'ennemi pour la table (supprime emojis/préfixes et met en minuscules)
func normalizeEnemyName(name string) string {
	n := strings.TrimSpace(strings.ToLower(name))
	n = nonLetter.ReplaceAllString(n, "")
	n = strings.ReplaceAll(n, "  ", " ")
	return n
}

// Attribue le loot spécifique d'un ennemi.
func awardEnemyLoot(enemyName string, isSuper bool) {
	key := normalizeEnemyName(enemyName)
	cfg, ok := enemyLootTable[key]
	if !ok || cfg.commonItem == "" { // aucun loot configuré
		return
	}
	qty := 1
	if isSuper { // super ennemis donnent plus
		qty = 2
	}
	addToInventory(cfg.commonItem, qty)
	// Chance du rare
	chance := cfg.rareChancePct
	if isSuper {
		// bonus +50% relatif (ex: 8 -> 12) sans dépasser 100
		chance = int(float64(chance) * 1.5)
		if chance > 100 {
			chance = 100
		}
	}
	if cfg.rareItem != "" && rand.Intn(100) < chance {
		addToInventory(cfg.rareItem, 1)
		fmt.Printf("🌟 Objet rare obtenu: %s !\n", cfg.rareItem)
	}
}
