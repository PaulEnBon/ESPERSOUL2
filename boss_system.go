package main

import "fmt"

// Abstraction générique multi-salles pour mini-boss & boss progressifs
// Codes tuiles par salle:
//  Salle12: mini 67, spawner 66, boss 68 (déjà implémenté)
//  Salle13: mini 70, spawner 71, boss 72 (Niveau 2)
//  Salle14: mini 73, spawner 74, boss 75 (Niveau 3)
//  Salle15: mini 76, spawner 77, boss 78 (Niveau 4)

// État générique par salle
type BossRoomState struct {
	defeatedMini map[string]bool
	fragments    int
	spawnerSpawn bool
	bossDefeated bool
}

// util clé
func keyXY(x, y int) string { return fmt.Sprintf("%d_%d", x, y) }

// Registre des salles gérées
var bossRooms = map[string]*struct {
	state     BossRoomState
	minis     [][2]int
	center    [2]int
	codeMini  int
	codeSpawn int
	codeBoss  int
	level     int
}{
	"salle13": {state: BossRoomState{defeatedMini: make(map[string]bool)}, minis: [][2]int{{1, 1}, {5, 1}, {1, 5}, {5, 5}}, center: [2]int{3, 3}, codeMini: 70, codeSpawn: 71, codeBoss: 72, level: 2},
	"salle14": {state: BossRoomState{defeatedMini: make(map[string]bool)}, minis: [][2]int{{1, 1}, {5, 1}, {1, 5}, {5, 5}}, center: [2]int{3, 3}, codeMini: 73, codeSpawn: 74, codeBoss: 75, level: 3},
	"salle15": {state: BossRoomState{defeatedMini: make(map[string]bool)}, minis: [][2]int{{1, 1}, {5, 1}, {1, 5}, {5, 5}}, center: [2]int{3, 3}, codeMini: 76, codeSpawn: 77, codeBoss: 78, level: 4},
}

// Appliquer l'état (appelé dans applyEnemyStates pour chaque salle concernée)
func applyGenericBossRoom(mapName string, mapData [][]int) {
	cfg, ok := bossRooms[mapName]
	if !ok {
		return
	}
	st := &cfg.state
	if st.bossDefeated {
		// Nettoyer éventuel boss/spawner restant
		cx, cy := cfg.center[0], cfg.center[1]
		if mapData[cy][cx] == cfg.codeSpawn || mapData[cy][cx] == cfg.codeBoss {
			mapData[cy][cx] = 0
		}
		return
	}
	// Mini boss placement
	for _, c := range cfg.minis {
		x, y := c[0], c[1]
		key := keyXY(x, y)
		if !st.defeatedMini[key] {
			mapData[y][x] = cfg.codeMini
		}
	}
	// Spawner apparition
	if st.fragments >= 4 && !st.spawnerSpawn {
		cx, cy := cfg.center[0], cfg.center[1]
		if mapData[cy][cx] != 1 {
			mapData[cy][cx] = cfg.codeSpawn
		}
		st.spawnerSpawn = true
	}
}
