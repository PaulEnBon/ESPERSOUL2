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
}

// Indique si un ennemi (par position) est "super" (2x stats). Clé: "x_y"
var superEnemyFlags = map[string]map[string]bool{
	"salle10": make(map[string]bool),
}
