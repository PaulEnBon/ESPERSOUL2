package main

import "fmt"

// EnemyInstance représente un ennemi placé sur la carte et persistant tant qu'il n'est pas vaincu.
type EnemyInstance struct {
	MapName string
	X, Y    int
	Nom     string
	Emoji   string
	Super   bool
	Alive   bool
	Persona Personnage // Copie complète des stats au moment de la génération
}

// enemiesByMap[map]["x,y"] = *EnemyInstance
var enemiesByMap = map[string]map[string]*EnemyInstance{}

// tempEnemyInstanceDuringCombat référence l'instance ennemie utilisée dans un combat en cours
var tempEnemyInstanceDuringCombat *EnemyInstance

func enemyCoordKey(x, y int) string { return fmt.Sprintf("%d,%d", x, y) }

// InitEnemiesForMap parcourt la map et crée des instances persistantes pour chaque cellule val==2 ou 12
// sans recréer celles déjà existantes (permet d'appeler la fonction plusieurs fois sans duplication).
func InitEnemiesForMap(mapName string, mapData [][]int) {
	if enemiesByMap[mapName] == nil {
		enemiesByMap[mapName] = map[string]*EnemyInstance{}
	}
	for y := 0; y < len(mapData); y++ {
		for x := 0; x < len(mapData[y]); x++ {
			val := mapData[y][x]
			if val == 2 || val == 12 { // Ennemi / Super Ennemi
				key := enemyCoordKey(x, y)
				if _, exists := enemiesByMap[mapName][key]; exists {
					continue // déjà instancié
				}
				// Génère le personnage via la logique existante
				isSuper := (val == 12)
				perso := CreateRandomEnemyForMap(mapName, isSuper)
				// Emoji déterministe basé sur le nom (réutilise emojiForEnemyName)
				emoji := emojiForEnemyName(perso.Nom)
				enemiesByMap[mapName][key] = &EnemyInstance{
					MapName: mapName,
					X:       x,
					Y:       y,
					Nom:     perso.Nom,
					Emoji:   emoji,
					Super:   isSuper,
					Alive:   true,
					Persona: perso,
				}
			}
		}
	}
}

// RerollEnemiesForMap force la régénération de tous les ennemis de la carte passée.
// Elle efface les instances précédentes puis recrée de nouveaux ennemis basés sur l'état actuel de mapData.
// Cas spéciaux :
//  - Si le Mentor est déjà transformé (salle1 coord 8,3 -> devenu PNJ 3), aucune instance n'est créée à cet emplacement.
//  - Les cases qui ne sont plus 2 ou 12 ne régénèrent pas d'ennemi.
func RerollEnemiesForMap(mapName string, mapData [][]int) {
	// Réinitialiser le conteneur pour cette map
	enemiesByMap[mapName] = map[string]*EnemyInstance{}

	// Si salle1 et Mentor déjà transformé, on note sa coord pour le skip
	skipMentor := false
	if mapName == "salle1" {
		key := "8_3"
		if trMap, ok := pnjTransformed[mapName]; ok && trMap[key] {
			skipMentor = true
		}
	}

	for y := 0; y < len(mapData); y++ {
		for x := 0; x < len(mapData[y]); x++ {
			val := mapData[y][x]
			if val == 2 || val == 12 { // cellule ennemie dans la carte
				if skipMentor && mapName == "salle1" && x == 8 && y == 3 {
					continue // ne pas recréer l'ancien Mentor
				}
				isSuper := (val == 12)
				perso := CreateRandomEnemyForMap(mapName, isSuper)
				emoji := emojiForEnemyName(perso.Nom)
				enemiesByMap[mapName][enemyCoordKey(x, y)] = &EnemyInstance{
					MapName: mapName,
					X:       x,
					Y:       y,
					Nom:     perso.Nom,
					Emoji:   emoji,
					Super:   isSuper,
					Alive:   true,
					Persona: perso,
				}
			}
		}
	}
}

// GetEnemyInstance returns the persistent enemy at (x,y) if any and alive
func GetEnemyInstance(mapName string, x, y int) (*EnemyInstance, bool) {
	if mm := enemiesByMap[mapName]; mm != nil {
		if inst, ok := mm[enemyCoordKey(x, y)]; ok && inst != nil && inst.Alive {
			return inst, true
		}
	}
	return nil, false
}
