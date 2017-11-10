package main

import (
	"github.com/kookehs/automation/api/win"
)

const (
	BASE       uintptr = 0x01000000
	DISCOVERED uintptr = 0x010057A4
	FIELD      uintptr = 0x01005361
	FLAGS      uintptr = 0x01005194
	HEIGHT     uintptr = 0x01005338
	STATE      uintptr = 0x01005160
	TIME       uintptr = 0x0100579C
	WIDTH      uintptr = 0x01005334
	X          uintptr = 0x01005118
	Y          uintptr = 0x0100511C
)

type Game struct {
	Discovered uint
	Field      []byte
	Flags      uint
	Height     uint
	State      uint
	Time       uint
	Width      uint
	X          uint
	Y          uint
}

func NewGame(handle win.HANDLE) *Game {
	game := new(Game)
	var read win.SIZE_T
	win.ReadProcessMemory(handle, win.LPCVOID(DISCOVERED), win.LPVOID(&game.Discovered), 4, &read)
	win.ReadProcessMemory(handle, win.LPCVOID(FLAGS), win.LPVOID(&game.Flags), 4, &read)
	win.ReadProcessMemory(handle, win.LPCVOID(HEIGHT), win.LPVOID(&game.Height), 4, &read)
	win.ReadProcessMemory(handle, win.LPCVOID(STATE), win.LPVOID(&game.State), 4, &read)
	win.ReadProcessMemory(handle, win.LPCVOID(TIME), win.LPVOID(&game.Time), 2, &read)
	win.ReadProcessMemory(handle, win.LPCVOID(WIDTH), win.LPVOID(&game.Width), 4, &read)
	win.ReadProcessMemory(handle, win.LPCVOID(X), win.LPVOID(&game.X), 4, &read)
	win.ReadProcessMemory(handle, win.LPCVOID(Y), win.LPVOID(&game.Y), 4, &read)
	game.Field = []byte{}
	game.ReadFieldMemory(handle)
	return game
}

func (game *Game) ReadFieldMemory(handle win.HANDLE) {
	address := FIELD
	var read win.SIZE_T
	row := make([]byte, game.Width)

	for y := 0; y < int(game.Height); y++ {
		win.ReadProcessMemory(handle, win.LPCVOID(address), win.LPVOID(&row[0]), win.SIZE_T(game.Width), &read)
		game.Field = append(game.Field, row...)
		address += 0x20
	}
}
