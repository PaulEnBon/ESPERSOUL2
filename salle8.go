package main

// Salle secrète 8x8 accessible uniquement avec une clé spéciale
var salle8 = [][]int{
	{9, 9, 9, 9, 9, 9, 9},
	{9, 0, 0, 0, 0, 0, 9},
	{9, 0, 0, 0, 0, 0, 9},
	{9, 0, 3, 3, 3, 0, 9}, // PNJ Matheo, Paul, Michael centrés (3,3), (4,3), (5,3)
	{9, 0, 0, 0, 0, 0, 9},
	{9, 0, 0, 0, 0, 0, 9},
	{9, 0, 0, 0, 0, 0, 9},
	{9, 9, 9, 32, 9, 9, 9}, // 32 = porte de sortie (centrée)
}
