package main

// Salle secrète 8x8 accessible uniquement avec une clé spéciale
var salle8 = [][]int{
	{9, 9, 9, 9, 9, 9, 9, 9},
	{9, 0, 0, 6, 6, 0, 0, 9}, // 6 = coffres au trésor
	{9, 0, 0, 0, 0, 0, 0, 9},
	{9, 6, 0, 3, 3, 0, 6, 9}, // 3 = PNJ gardiens, 6 = coffres
	{9, 0, 0, 3, 3, 0, 0, 9},
	{9, 0, 0, 0, 0, 0, 0, 9},
	{9, 0, 0, 0, 0, 0, 0, 9},   // Plus de coffres
	{9, 9, 9, 31, 32, 9, 9, 9}, // 31 = spawn point, 32 = porte de sortie
}
