package main

import (
	"github.com/kookehs/automation/api/win"
	"unsafe"
)

const (
	OFFSETX  uint8 = 16
	OFFSETY  uint8 = 102
	CELLSIZE uint8 = 16
)

const (
	CLICKED             byte = 0
	CLICKEDBOMB         byte = 128
	CLICKEDEXPLOSION    byte = 204
	CLICKEDQUESTIONBOMB byte = 137

	HIDDEN               byte = 15
	HIDDENBOMB           byte = 143
	HIDDENFLAGBOMB       byte = 142
	HIDDENFLAGNOBOMB     byte = 14
	HIDDENQUESTIONBOMB   byte = 141
	HIDDENQUESTIONNOBOMB byte = 13

	REVEALED             byte = 64
	REVEALEDBOMB         byte = 138
	REVEALEDFLAGNOBOMB   byte = 11
	REVEALEDQUESTIONBOMB byte = 138

	EIGHT byte = 72
	FIVE  byte = 69
	FOUR  byte = 68
	ONE   byte = 65
	SEVEN byte = 71
	SIX   byte = 70
	THREE byte = 67
	TWO   byte = 66
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
	Discovered uint16
	Field      []byte
	Flags      uint8
	Height     uint8
	State      uint8
	Time       uint16
	Width      uint8
	X          uint8
	Y          uint8
}

func NewGame(handle win.HANDLE) *Game {
	game := new(Game)
	var read win.SIZE_T

	win.ReadProcessMemory(handle, win.LPCVOID(DISCOVERED), win.LPVOID(&game.Discovered),
		win.SIZE_T(unsafe.Sizeof(game.Discovered)), &read)
	win.ReadProcessMemory(handle, win.LPCVOID(FLAGS), win.LPVOID(&game.Flags),
		win.SIZE_T(unsafe.Sizeof(game.Flags)), &read)
	win.ReadProcessMemory(handle, win.LPCVOID(HEIGHT), win.LPVOID(&game.Height),
		win.SIZE_T(unsafe.Sizeof(game.Height)), &read)
	win.ReadProcessMemory(handle, win.LPCVOID(STATE), win.LPVOID(&game.State),
		win.SIZE_T(unsafe.Sizeof(game.State)), &read)
	win.ReadProcessMemory(handle, win.LPCVOID(TIME), win.LPVOID(&game.Time),
		win.SIZE_T(unsafe.Sizeof(game.Time)), &read)
	win.ReadProcessMemory(handle, win.LPCVOID(WIDTH), win.LPVOID(&game.Width),
		win.SIZE_T(unsafe.Sizeof(game.Width)), &read)
	win.ReadProcessMemory(handle, win.LPCVOID(X), win.LPVOID(&game.X),
		win.SIZE_T(unsafe.Sizeof(game.X)), &read)
	win.ReadProcessMemory(handle, win.LPCVOID(Y), win.LPVOID(&game.Y),
		win.SIZE_T(unsafe.Sizeof(game.Y)), &read)

	game.Field = []byte{}
	game.ReadFieldMemory(handle)
	return game
}

func (game *Game) CellToCoordinates(cell uint8) (uint8, uint8) {
	return cell % game.Width, cell / game.Height
}

func (game *Game) GetAdjacentCells(cell uint8) map[uint8][]uint8 {
	cells := make(map[uint8][]uint8)
	x, y := game.CellToCoordinates(cell)

	for dx := -1; dx < 2; dx++ {
		for dy := -1; dy < 2; dy++ {
			if dx|dy == 0 {
				continue
			}

			rx := x + uint8(dx)
			ry := y + uint8(dy)

			if rx >= 0 && rx < game.Width && ry >= 0 && ry < game.Height {
				index := game.Width*ry + rx
				cells[game.Field[index]] = append(cells[game.Field[index]], index)
			}
		}
	}

	return cells
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

func CellByteToNumeric(cell uint8) uint8 {
	number := cell - ONE + 1

	if number < 1 && number > 8 {
		number = 0
	}

	return number
}
