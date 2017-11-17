package main

import (
	"fmt"
	"github.com/kookehs/automation/api/win"
	"math/rand"
	"os"
	"strconv"
	"syscall"
	"time"
)

type Process struct {
	FileName []byte
	Handle   win.HANDLE
	HWnd     win.HWND
	Pid      win.DWORD
	Rect     win.RECT
}

func NewProcess(pid win.DWORD) *Process {
	process := new(Process)
	process.FileName = make([]byte, win.MAX_PATH)
	process.Pid = pid

	var access win.DWORD = win.PROCESS_VM_READ | win.PROCESS_QUERY_INFORMATION
	var inherit win.BOOL = 0
	process.Handle = win.OpenProcess(access, inherit, process.Pid)

	if process.Handle == win.NULL {
		return nil
	}

	win.GetModuleFileNameEx(process.Handle, 0, &process.FileName, win.MAX_PATH)

	callback := syscall.NewCallback(func(hWnd win.HWND, lParam win.LPARAM) win.BOOL {
		var pid win.DWORD
		win.GetWindowThreadProcessId(hWnd, &pid)

		if uintptr(pid) == uintptr(lParam) {
			process.HWnd = hWnd
			return 0
		}

		return 1
	})

	win.EnumWindows(callback, win.LPARAM(process.Pid))
	win.GetWindowRect(process.HWnd, &process.Rect)
	return process
}

func CellToScreenCoordinates(cell uint8, game *Game, process *Process) (int32, int32) {
	x := process.Rect.Left + int32(OFFSETX+cell%game.Width*CELLSIZE)
	y := process.Rect.Top + int32(OFFSETY+cell/game.Height*CELLSIZE)
	return x, y
}

func ExecuteCommands(commands map[byte][]uint8, game *Game, process *Process) {
	for command, cells := range commands {
		button := MOUSE_CLICKLEFT

		if command == 'F' {
			button = MOUSE_CLICKRIGHT
		} else if command == 'M' {
			button = MOUSE_CLICKMIDDLE
		}

		for _, cell := range cells {
			x, y := CellToScreenCoordinates(cell, game, process)
			MouseClick(button, x, y, 1)
		}
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: bot.exe <pid>")
		return
	}

	var process *Process

	if pid, err := strconv.ParseUint(os.Args[1], 10, 32); err != nil {
		fmt.Println(err)
		return
	} else {
		process = NewProcess(win.DWORD(pid))
	}

	defer win.CloseHandle(process.Handle)
	fmt.Println(process)

	game := NewGame(process.Handle)
	fmt.Println(game)

	rand.Seed(time.Now().UTC().UnixNano())

	win.SetForegroundWindow(process.HWnd)

	// ExecuteCommands([]uint8{RandomCell(game)}, game, process)
	ExecuteCommands(SolveStraightforward(game), game, process)
}
