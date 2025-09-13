package main

var salle2 = [][]int{
	{9, 9, 9, 7, 9},  // ↑ vers salle3
	{9, 0, 0, 11, 9}, // spawn retour depuis salle3
	{9, 0, 0, 0, 9},
	{9, 0, 0, 0, 9},
	{9, 0, 0, 0, 9},
	{9, 0, 0, 0, 9},
	{9, 0, 0, 0, 9},
	{9, 0, 1, 0, 9},  // spawn initial
	{9, 9, 10, 9, 9}, // ↓ vers salle1
}
