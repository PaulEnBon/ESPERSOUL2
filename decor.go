package main

import "fmt"

// Système simple de décorations (arbres, etc.)

// Code tuile pour un arbre
const TileTree = 80

// Placement de tuile décorative
type TilePlacement struct {
	X int
	Y int
}

// Dictionnaire: nomCarte -> liste de positions d'arbres
var mapTrees = map[string][]TilePlacement{}

// Arbres coupés persistants par carte: map[mapName]map["x_y"]bool
var cutTrees = map[string]map[string]bool{}

// Ajoute un arbre à une carte donnée (ne duplique pas)
func AddTree(mapName string, x, y int) {
	list := mapTrees[mapName]
	for _, p := range list {
		if p.X == x && p.Y == y { // déjà présent
			return
		}
	}
	mapTrees[mapName] = append(list, TilePlacement{X: x, Y: y})
}

// Retire un arbre d'une carte si présent
func RemoveTree(mapName string, x, y int) {
	list := mapTrees[mapName]
	out := make([]TilePlacement, 0, len(list))
	for _, p := range list {
		if !(p.X == x && p.Y == y) {
			out = append(out, p)
		}
	}
	mapTrees[mapName] = out
}

// Supprime tous les arbres d'une carte
func ClearTrees(mapName string) {
	delete(mapTrees, mapName)
}

// Applique les décorations à une carte (écrit TileTree uniquement sur sol vide)
func applyDecorations(mapName string, mapData [][]int) {
	placements := mapTrees[mapName]
	if len(placements) == 0 {
		return
	}
	h := len(mapData)
	if h == 0 {
		return
	}
	w := len(mapData[0])
	for _, p := range placements {
		if p.Y < 0 || p.Y >= h || p.X < 0 || p.X >= w {
			continue
		}
		// Ne pas écraser des entités importantes: on ne place que sur sol vide (0)
		cell := mapData[p.Y][p.X]
		if cell == 0 {
			mapData[p.Y][p.X] = TileTree
		}
	}
}

// Place immédiatement un arbre dans la carte en mémoire (utile pour debug/cheat)
func PlaceTreeImmediate(currentMap string, mapData [][]int, x, y int) {
	AddTree(currentMap, x, y)
	applyDecorations(currentMap, mapData)
}

// Marque un arbre comme coupé (persistance)
func MarkTreeCut(mapName string, x, y int) {
	if cutTrees[mapName] == nil {
		cutTrees[mapName] = map[string]bool{}
	}
	key := fmt.Sprintf("%d_%d", x, y)
	cutTrees[mapName][key] = true
}

// Applique la suppression des arbres coupés sur une carte chargée
func applyCutTrees(mapName string, mapData [][]int) {
	m := cutTrees[mapName]
	if m == nil {
		return
	}
	h := len(mapData)
	if h == 0 {
		return
	}
	w := len(mapData[0])
	for key := range m {
		var x, y int
		// parse key "x_y"
		if _, err := fmt.Sscanf(key, "%d_%d", &x, &y); err == nil {
			if y >= 0 && y < h && x >= 0 && x < w {
				if mapData[y][x] == TileTree {
					mapData[y][x] = 0
				}
			}
		}
	}
}
