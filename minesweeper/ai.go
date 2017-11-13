package main

import (
	"fmt"
	"math/rand"
)

func CellToCoordinates(cell int32, game *Game, process *Process) (int32, int32) {
	x := process.Rect.Left + int32(OFFSETX) + cell%int32(game.Width)*int32(CELLSIZE)
	y := process.Rect.Top + int32(OFFSETY) + cell%int32(game.Height)*int32(CELLSIZE)
	fmt.Println(cell, x, y)
	return x, y
}

func RandomClick(game *Game, process *Process) {
	// TODO: Fix seeding and hidden cells
	hidden := []int{}

	for i := 0; i < int(game.Width*game.Height); i++ {
		if game.Field[i] == HIDDEN {
			hidden = append(hidden, i)
		}
	}

	index := rand.Int31n(int32(len(hidden)))
	cell := int32(hidden[index])
	x, y := CellToCoordinates(cell, game, process)
	MouseClick(MOUSE_CLICKLEFT, x, y, 1)
}
