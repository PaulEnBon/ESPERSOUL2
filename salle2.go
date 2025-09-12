package main

var salle2 = [][]int{
	{9, 9, 9, 7, 9}, // ↑ porte haut vers salle3
	{9, 0, 0, 11, 9},
	{9, 0, 0, 0, 9}, // 11 = spawn si retour depuis salle3 (2,4)
	{9, 0, 0, 0, 9},
	{9, 0, 0, 0, 9},
	{9, 0, 0, 0, 9},
	{9, 0, 0, 0, 9},
	{9, 0, 1, 0, 9},  // spawn initial joueur
	{9, 9, 10, 9, 9}, // ↓ porte bas vers salle1
}
