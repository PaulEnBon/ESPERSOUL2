package main

// État des ennemis vaincus par salle
var enemiesDefeated = map[string]map[string]bool{
	"salle1":  make(map[string]bool),
	"salle2":  make(map[string]bool),
	"salle3":  make(map[string]bool),
	"salle4":  make(map[string]bool),
	"salle5":  make(map[string]bool),
	"salle6":  make(map[string]bool),
	"salle7":  make(map[string]bool),
	"salle8":  make(map[string]bool),
	"salle9":  make(map[string]bool),
	"salle10": make(map[string]bool),
	"salle11": make(map[string]bool),
	"salle12": make(map[string]bool),
	"salle13": make(map[string]bool),
	"salle14": make(map[string]bool),
	"salle15": make(map[string]bool),
}

// État des ennemis transformés en PNJ (pour différencier d'une disparition)
var pnjTransformed = map[string]map[string]bool{
	"salle1":  make(map[string]bool),
	"salle2":  make(map[string]bool),
	"salle3":  make(map[string]bool),
	"salle4":  make(map[string]bool),
	"salle5":  make(map[string]bool),
	"salle6":  make(map[string]bool),
	"salle7":  make(map[string]bool),
	"salle8":  make(map[string]bool),
	"salle9":  make(map[string]bool),
	"salle10": make(map[string]bool),
	"salle11": make(map[string]bool),
	"salle12": make(map[string]bool),
	"salle13": make(map[string]bool),
	"salle14": make(map[string]bool),
	"salle15": make(map[string]bool),
}

// État des récompenses déjà récupérées
var rewardsGiven = map[string]map[string]bool{
	"salle1":  make(map[string]bool),
	"salle2":  make(map[string]bool),
	"salle3":  make(map[string]bool),
	"salle4":  make(map[string]bool),
	"salle5":  make(map[string]bool),
	"salle6":  make(map[string]bool),
	"salle7":  make(map[string]bool),
	"salle8":  make(map[string]bool),
	"salle9":  make(map[string]bool),
	"salle10": make(map[string]bool),
	"salle11": make(map[string]bool),
	"salle12": make(map[string]bool),
	"salle13": make(map[string]bool),
	"salle14": make(map[string]bool),
	"salle15": make(map[string]bool),
}

// Transitions entre les salles
var transitions = map[string]map[int]struct {
	nextMap string
	spawnX  int
	spawnY  int
}{
	"salle1": {
		7: {nextMap: "salle2", spawnX: 2, spawnY: 7},
	},
	"salle2": {
		20: {nextMap: "salle3", spawnX: 8, spawnY: 13},
		10: {nextMap: "salle1", spawnX: 8, spawnY: 1},
	},
	"salle3": {
		21: {nextMap: "salle2", spawnX: 2, spawnY: 1},
		13: {nextMap: "salle4", spawnX: 8, spawnY: 1},
		33: {nextMap: "salle5", spawnX: 2, spawnY: 3},
		34: {nextMap: "salle6", spawnX: 2, spawnY: 3},
		38: {nextMap: "salle9", spawnX: 1, spawnY: 1},
		50: {nextMap: "salle12", spawnX: 3, spawnY: 6}, // Porte vers salle12
	},
	"salle4": {
		14: {nextMap: "salle3", spawnX: 1, spawnY: 11},
	},
	"salle5": {
		31: {nextMap: "salle3", spawnX: 1, spawnY: 1},
	},
	"salle6": {
		15: {nextMap: "salle3", spawnX: 12, spawnY: 1},
		28: {nextMap: "salle3", spawnX: 12, spawnY: 1},
		27: {nextMap: "salle7", spawnX: 2, spawnY: 3},
	},
	"salle7": {
		27: {nextMap: "salle3", spawnX: 7, spawnY: 1},
		31: {nextMap: "salle6", spawnX: 2, spawnY: 3},
	},
	"salle8": {
		32: {nextMap: "salle1", spawnX: 15, spawnY: 6},
	},
	"salle9": {
		38: {nextMap: "salle3", spawnX: 13, spawnY: 12},
		44: {nextMap: "salle11", spawnX: 3, spawnY: 3},
	},
	"salle10": {
		42: {nextMap: "salle9", spawnX: 7, spawnY: 1},
	},
	"salle11": {
		42: {nextMap: "salle9", spawnX: 7, spawnY: 1},
	},
	"salle12": {
		51: {nextMap: "salle3", spawnX: 7, spawnY: 1},  // Retour salle3
		52: {nextMap: "salle13", spawnX: 1, spawnY: 3}, // Vers salle13 (spawn EXACT sur la tuile 61 située en (1,3))
		54: {nextMap: "salle14", spawnX: 5, spawnY: 3}, // Vers salle14 (spawn sur la tuile 63 en (5,3))
		56: {nextMap: "salle15", spawnX: 3, spawnY: 6}, // Vers salle15 (spawn sur la tuile 65 en (3,6))
	},
	"salle13": {
		53: {nextMap: "salle12", spawnX: 5, spawnY: 3}, // Retour salle12 (spawn sur 60)
	},
	"salle14": {
		55: {nextMap: "salle12", spawnX: 1, spawnY: 3}, // Retour salle12 (spawn sur la tuile 62)
	},
	"salle15": {
		57: {nextMap: "salle12", spawnX: 3, spawnY: 1}, // Retour salle12 (spawn sur la tuile 64 en (3,1))
	},
}

// Indique si un ennemi (par position) est "super" (2x stats). Clé: "x_y"
var superEnemyFlags = map[string]map[string]bool{
	"salle10": make(map[string]bool),
}
