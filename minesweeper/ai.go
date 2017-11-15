package main

import (
	"math/rand"
)

func GetAdjacentCells() []int {
	cells := []int{}

	for i := 0; i < ADJACENTCELLS; i++ {

	}
}

func SolveStraightforward() []int {
	cells := []int{}

	for i := 0; i < int(game.Width*game.Height); i++ {

	}

	return cells
}

func RandomClick(game *Game, process *Process) {
	hidden := []int{}

	for i := 0; i < int(game.Width*game.Height); i++ {
		cell := game.Field[i]

		if cell == HIDDEN || cell == HIDDENBOMB {
			hidden = append(hidden, i)
		}
	}

	index := rand.Int31n(int32(len(hidden)))
	x, y := CellToCoordinates(int32(hidden[index]), game, process)
	MouseClick(MOUSE_CLICKLEFT, x, y, 1)
}
