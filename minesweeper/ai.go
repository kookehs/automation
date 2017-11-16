package main

import (
	"fmt"
	"math/rand"
)

func SolveStraightforward(game *Game) []uint8 {
	cells := []uint8{}

	for i := 0; i < int(game.Width*game.Height); i++ {

	}

	return cells
}

func RandomClick(game *Game, process *Process) {
	fmt.Println(game.GetAdjacentCells(0))

	hidden := []uint8{}

	for i := 0; i < int(game.Width*game.Height); i++ {
		cell := game.Field[i]

		if cell == HIDDEN || cell == HIDDENBOMB {
			hidden = append(hidden, uint8(i))
		}
	}

	index := rand.Int31n(int32(len(hidden)))
	x, y := CellToScreenCoordinates(hidden[index], game, process)
	MouseClick(MOUSE_CLICKLEFT, x, y, 1)
}
