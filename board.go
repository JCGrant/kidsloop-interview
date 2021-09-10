package main

// Board represents a 2D grid of booleans. Useful for dynamic programming purposes
// Creating a board with a width will center the grid at 0, 0
// This allows for negative indexes to be used in .Get and .Set
type Board struct {
	width  int
	height int
	center int
	cells  [][]bool
}

func NewBoard(width int) Board {
	height := width
	cells := make([][]bool, height)
	for i := 0; i < height; i++ {
		row := make([]bool, width)
		cells[i] = row
	}
	return Board{
		width:  width,
		height: height,
		center: width / 2,
		cells:  cells,
	}
}

func (b Board) Get(x, y int) bool {
	return b.cells[b.center-y][b.center+x]
}

func (b Board) Set(x, y int, value bool) {
	b.cells[b.center-y][b.center+x] = value
}
