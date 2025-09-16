package main

// Salle 5x5 avec forgeron
var salle5 = [][]int{
	{9, 9, 9, 9, 9},
	{9, 0, 0, 0, 9},
	{9, 0, 5, 0, 9}, // 5 = forgeron
	{9, 0, 24, 0, 9},
	{9, 9, 31, 9, 9}, // 31 = porte sud vers salle3, spawn point 24
}
