package main

// Salle 5x5 avec coffre
var salle6 = [][]int{
	{9, 9, 9, 28, 9}, // (3,0) flèche vers salle7
	{9, 0, 0, 0, 9},
	{9, 0, 6, 0, 9}, // 6 = coffre
	{9, 0, 25, 0, 9},
	{9, 9, 15, 9, 9}, // 15 = porte sud vers salle3, spawn point 25
}
