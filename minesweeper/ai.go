package main

import (
	"math/rand"
)

func RandomClick(game *Game) {
	hidden := []byte{}

	for i := 0; i < game.Width*game.Height; i++ {
		if game.Field[i] == HIDDEN {
			hidden = append(hidden, i)
		}
	}

	index := rand.Int31n(len(hidden))
	cell := hidden[index]
	// TODO: Create a function for cell to coordinates
}
