package win

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type (
	HWND        = HANDLE
	WNDENUMPROC = uintptr
)

var (
	// DLL
	user32 *windows.DLL

	// Functions
	enumWindows              *windows.Proc
	getWindowThreadProcessId *windows.Proc
	setForegroundWindow      *windows.Proc
)

func init() {
	// DLL
	user32 = windows.MustLoadDLL("user32.dll")

	// Functions
	enumWindows = user32.MustFindProc("EnumWindows")
	getWindowThreadProcessId = user32.MustFindProc("GetWindowThreadProcessId")
	setForegroundWindow = user32.MustFindProc("SetForegroundWindow")
}

func EnumWindows(lpEnumFunc WNDENUMPROC, lParam LPARAM) BOOL {
	ret, _, _ := enumWindows.Call(lpEnumFunc, lParam)
	return BOOL(ret)
}

func GetWindowThreadProcessId(hWnd HWND, lpdwProcessId LPDWORD) DWORD {
	ret, _, _ := getWindowThreadProcessId.Call(hWnd, uintptr(unsafe.Pointer(&lpdwProcessId)))
	return DWORD(ret)
}

func SetForegroundWindow(hWnd HWND) BOOL {
	ret, _, _ := setForegroundWindow.Call(hWnd)
	return BOOL(ret)
}
