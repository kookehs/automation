package main

import (
	"math/rand"
)

func SolveStraightforward(game *Game) map[byte][]uint8 {
	// TODO: Fix checking of adjacent cells to account for flags and question marks
	commands := make(map[byte][]uint8)

	for i := 0; i < int(game.Width*game.Height); i++ {
		adjacent := game.GetAdjacentCells(uint8(i))

		if CellByteToNumeric(game.Field[i]) == uint8(len(adjacent[HIDDEN])+len(adjacent[HIDDENBOMB])) {
			commands['F'] = append(commands['F'], adjacent[HIDDEN]...)
			commands['F'] = append(commands['F'], adjacent[HIDDENBOMB]...)
		}
	}

	return commands
}

func RandomCell(game *Game) uint8 {
	hidden := []uint8{}

	for i := 0; i < int(game.Width*game.Height); i++ {
		cell := game.Field[i]

		if cell == HIDDEN || cell == HIDDENBOMB {
			hidden = append(hidden, uint8(i))
		}
	}

	return hidden[rand.Int31n(int32(len(hidden)))]
}
