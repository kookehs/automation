package main

import (
	"fmt"
	"github.com/kookehs/automation/api/win"
	"os"
	"strconv"
	"syscall"
	"unsafe"
)

type Process struct {
	FileName []byte
	Handle   win.HANDLE
	HWnd     win.HWND
	Pid      win.DWORD
}

func NewProcess() *Process {
	return &Process{
		FileName: make([]byte, win.MAX_PATH),
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: bot.exe <pid>")
		return
	}

	var minesweeper *Process = NewProcess()

	if pid, err := strconv.ParseUint(os.Args[1], 10, 32); err != nil {
		fmt.Println(err)
		return
	} else {
		minesweeper.Pid = win.DWORD(pid)
	}

	var access win.DWORD = win.PROCESS_VM_READ | win.PROCESS_QUERY_INFORMATION
	var inherit win.BOOL = 0
	minesweeper.Handle = win.OpenProcess(access, inherit, minesweeper.Pid)

	if minesweeper.Handle == win.NULL {
		return
	}

	defer win.CloseHandle(minesweeper.Handle)
	fmt.Println("handle: ", minesweeper.Handle)

	if win.GetModuleFileNameEx(minesweeper.Handle, 0, &minesweeper.FileName, win.MAX_PATH) != 0 {
		fmt.Println("file name: ", string(minesweeper.FileName))
	}

	var address uintptr = 0x010057A4
	var discovered [4]byte
	var read win.SIZE_T

	if win.ReadProcessMemory(minesweeper.Handle, win.LPCVOID(address), win.LPVOID(&discovered), 4, &read) != 0 {
		fmt.Println("read: ", read)
		fmt.Println("discovered: ", discovered)
	}

	callback := syscall.NewCallback(func(hWnd win.HWND, lParam win.LPARAM) win.BOOL {
		var pid win.DWORD
		win.GetWindowThreadProcessId(hWnd, &pid)

		if uintptr(pid) == uintptr(lParam) {
			minesweeper.HWnd = hWnd
			return 0
		}

		return 1
	})

	if win.EnumWindows(callback, win.LPARAM(minesweeper.Pid)) != 0 {
		fmt.Println("hWnd: ", minesweeper.HWnd)
	}

	if win.SetForegroundWindow(minesweeper.HWnd) != 0 {
		fmt.Println("foreground: ", minesweeper.HWnd)
	}

	// TODO: Fix issue with absolute positioning
	var size win.UINT = 1
	inputs := make([]win.MOUSE_INPUT, size)

	input := win.MOUSE_INPUT{
		Type: win.INPUT_MOUSE,
		Mi: win.MOUSEINPUT{
			Dx:          480,
			Dy:          10,
			MouseData:   0,
			DwFlags:     win.MOUSEEVENTF_ABSOLUTE | win.MOUSEEVENTF_MOVE,
			Time:        0,
			DwExtraInfo: 0,
		},
	}

	inputs[0] = input
	win.SendInput(size, unsafe.Pointer(&inputs[0]), int(unsafe.Sizeof(win.MOUSE_INPUT{})))

	inputs[0].Mi.DwFlags = win.MOUSEEVENTF_ABSOLUTE | win.MOUSEEVENTF_LEFTDOWN
	win.SendInput(size, unsafe.Pointer(&inputs[0]), int(unsafe.Sizeof(win.MOUSE_INPUT{})))

	inputs[0].Mi.DwFlags = win.MOUSEEVENTF_ABSOLUTE | win.MOUSEEVENTF_LEFTUP
	win.SendInput(size, unsafe.Pointer(&inputs[0]), int(unsafe.Sizeof(win.MOUSE_INPUT{})))
}
