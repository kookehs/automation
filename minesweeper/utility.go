package main

import (
	"github.com/kookehs/automation/api/win"
	"unsafe"
)

func mouseClick(button win.DWORD, x, y int32, clicks win.UINT) {
	inputs := make([]win.MOUSE_INPUT, clicks)
	x, y = normalizeCoordinates(x, y)

	for i := 0; i < int(clicks); i++ {
		inputs[i] = win.MOUSE_INPUT{
			Type: win.INPUT_MOUSE,
			Mi: win.MOUSEINPUT{
				Dx:          x,
				Dy:          y,
				MouseData:   0,
				DwFlags:     win.MOUSEEVENTF_ABSOLUTE | win.MOUSEEVENTF_MOVE | button,
				Time:        0,
				DwExtraInfo: 0,
			},
		}
	}

	win.SendInput(clicks, unsafe.Pointer(&inputs[0]), int(unsafe.Sizeof(win.MOUSE_INPUT{})))
}

func normalizeCoordinates(x, y int32) (int32, int32) {
	var desktopWindow win.RECT

	if win.GetWindowRect(win.GetDesktopWindow(), &desktopWindow) != 0 {
		x = x * 65536 / desktopWindow.Right
		y = y * 65536 / desktopWindow.Bottom
	}

	return x, y
}
