package main

// Salle 11 : petite salle avec un PNJ soigneur
var salle11 = [][]int{
	{9, 9, 9, 9, 9, 9, 9},
	{9, 0, 0, 0, 0, 0, 9},  // porte gauche (←) en (1,1) pour retour salle9
	{9, 0, 3, 0, 79, 0, 9}, // PNJ au centre en (3,2)
	{9, 0, 0, 0, 0, 0, 9},
	{9, 9, 9, 42, 9, 9, 9},
}
