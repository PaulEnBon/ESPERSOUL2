package main

import "fmt"

// Système mini-boss + boss salle12
// Codes tuiles utilisés:
// 67 = Mini-boss
// 66 = Spawner (apparaît quand 4 fragments obtenus)
// 68 = Boss

// Etat persistant pour la salle12
var salle12BossState = struct {
	defeatedMini map[string]bool // cle "x_y" pour chaque mini boss vaincu
	fragments    int             // nb de fragments collectés (max 4)
	spawnerSpawn bool            // spawner déjà apparu ?
	bossDefeated bool            // boss final vaincu ?
}{
	defeatedMini: make(map[string]bool),
}

// Coordonnées des mini boss (coins intérieurs) (x,y)
var salle12MiniBossCoords = [][2]int{{1, 1}, {5, 1}, {1, 5}, {5, 5}}

// Coordonnée centre pour spawner/boss
var salle12Center = [2]int{3, 3}

// Replace ou place une tuile si la case actuelle est vide ou différente
func setTileIfEmptyOr(mapData [][]int, x, y, val int) {
	if y < 0 || y >= len(mapData) || x < 0 || x >= len(mapData[0]) {
		return
	}
	if mapData[y][x] == 0 || mapData[y][x] == 66 || mapData[y][x] == 68 { // autorisé à réécrire phases
		mapData[y][x] = val
	}
}

// Appliqué dans applyEnemyStates après traitement générique
func applySalle12MiniBossSystem(mapData [][]int) {
	// Ne rien faire si boss final déjà vaincu (nettoyer spawner/boss)
	if salle12BossState.bossDefeated {
		// S'assurer qu'aucune entité spéciale ne reste
		if salle12BossState.spawnerSpawn { // on peut laisser le centre vide après victoire
			cx, cy := salle12Center[0], salle12Center[1]
			if mapData[cy][cx] == 66 || mapData[cy][cx] == 68 {
				mapData[cy][cx] = 0
			}
		}
		return
	}
	// Placer mini boss non vaincus
	for _, c := range salle12MiniBossCoords {
		x, y := c[0], c[1]
		key := fmt.Sprintf("%d_%d", x, y)
		if !salle12BossState.defeatedMini[key] {
			mapData[y][x] = 67 // mini boss
		}
	}
	// Si les 4 fragments collectés et spawner non placé -> placer spawner
	if salle12BossState.fragments >= 4 && !salle12BossState.spawnerSpawn {
		cx, cy := salle12Center[0], salle12Center[1]
		// Ne pas écraser le joueur (1) mais sinon placer
		if mapData[cy][cx] != 1 {
			mapData[cy][cx] = 66
		}
		salle12BossState.spawnerSpawn = true
	}
	// Si boss déjà invoqué (spawner transformé) mais pas encore vaincu conserver tuile 68
	if salle12BossState.spawnerSpawn && !salle12BossState.bossDefeated {
		// rien de plus ici; l'apparition du boss est gérée dans handleCellInteraction
	}
}
