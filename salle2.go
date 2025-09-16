package main

var salle2 = [][]int{
	{9, 9, 9, 20, 9}, // ↑ porte haut vers salle3
	{9, 0, 0, 17, 9}, // 17 = spawn depuis salle3
	{9, 0, 0, 0, 9},
	{9, 0, 0, 0, 9},
	{9, 0, 0, 0, 9},
	{9, 0, 0, 0, 9},
	{9, 0, 0, 0, 9},
	{9, 0, 19, 0, 9}, // spawn initial joueur
	{9, 9, 10, 9, 9}, // ↓ porte bas vers salle1
}
