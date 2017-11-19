package main

import (
	"math/rand"
)

func SolveStraightforward(game *Game) map[uint8]byte {
	commands := make(map[uint8]byte)

	for i := 0; i < int(game.Width*game.Height); i++ {
		adjacent := game.GetAdjacentCells(uint8(i))
		hidden := len(adjacent[HIDDEN])
		hiddenBomb := len(adjacent[HIDDENBOMB])
		hiddenFlag := len(adjacent[HIDDENFLAGBOMB]) + len(adjacent[HIDDENFLAGNOBOMB])
		hiddenQuestion := len(adjacent[HIDDENQUESTIONBOMB]) + len(adjacent[HIDDENQUESTIONNOBOMB])

		switch CellByteToNumeric(game.Field[i]) {
		case 0:
			break
		case uint8(hiddenFlag + hiddenQuestion):
			for _, cell := range adjacent[HIDDEN] {
				commands[cell] = 'L'
			}

			break
		case uint8(hidden + hiddenBomb + hiddenFlag + hiddenQuestion):
			for _, cell := range adjacent[HIDDEN] {
				commands[cell] = 'R'
			}

			for _, cell := range adjacent[HIDDENBOMB] {
				commands[cell] = 'R'
			}

			break
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
