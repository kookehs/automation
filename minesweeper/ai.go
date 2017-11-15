package main

import (
	"math/rand"
)

func CellToCoordinates(cell int32, game *Game, process *Process) (int32, int32) {
	x := process.Rect.Left + int32(OFFSETX) + cell%int32(game.Width)*int32(CELLSIZE)
	y := process.Rect.Top + int32(OFFSETY) + cell/int32(game.Height)*int32(CELLSIZE)
	return x, y
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
