package main

import (
	"fmt"
)

var (
	safeSum = 23
	steps   = []Vector2D{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}
	/*
	  With some initial analysis of the problem, we can figure out that the area will have a roughly
	  diamond shape, the longest points being near the axes. So we can find the first number where it's
	  digits sum to over 23 and this happens to be 699. With this knowledge, we know that the minimum width
	  for our board we will use for dynamic programming will be 699 * 2 + 1, which equals 1399.
	*/
	dfsBoardWidth = 1399
)

func haventVisited(visited Board, cell Vector2D) bool {
	return !visited.Get(cell.X, cell.Y)
}

func isSafe(coord Vector2D) bool {
	return sumOfDigits(intAbs(coord.X))+sumOfDigits(intAbs(coord.Y)) <= safeSum
}

func solveByDFS(visited Board, cell Vector2D, area *uint64) {
	visited.Set(cell.X, cell.Y, true)
	(*area)++
	for _, step := range steps {
		nextCell := cell.AddVector(step)
		if haventVisited(visited, nextCell) && isSafe(nextCell) {
			solveByDFS(visited, nextCell, area)
		}
	}
}

func main() {
	visited := NewBoard(dfsBoardWidth)
	initialCell := Vector2D{0, 0}
	var area uint64
	solveByDFS(visited, initialCell, &area)
	fmt.Println("area:", area)
}
