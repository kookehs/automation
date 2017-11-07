package win

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

const (
	INPUT_MOUSE    = 0
	INPUT_KEYBOARD = 1
	INPUT_HARDWARE = 2
)

const (
	MOUSEEVENTF_ABSOLUTE        = 0x8000
	MOUSEEVENTF_HWHEEL          = 0x1000
	MOUSEEVENTF_MOVE            = 0x0001
	MOUSEEVENTF_MOVE_NOCOALESCE = 0x2000
	MOUSEEVENTF_LEFTDOWN        = 0x0002
	MOUSEEVENTF_LEFTUP          = 0x0004
	MOUSEEVENTF_RIGHTDOWN       = 0x0008
	MOUSEEVENTF_RIGHTUP         = 0x0010
	MOUSEEVENTF_MIDDLEDOWN      = 0x0020
	MOUSEEVENTF_MIDDLEUP        = 0x0040
	MOUSEEVENTF_VIRTUALDESK     = 0x4000
	MOUSEEVENTF_WHEEL           = 0x0800
	MOUSEEVENTF_XDOWN           = 0x0080
	MOUSEEVENTF_XUP             = 0x0100
	XBUTTON1                    = 0x0001
	XBUTTON2                    = 0x0002
)

type (
	HWND        = HANDLE
	LPINPUT     = unsafe.Pointer
	WNDENUMPROC = uintptr
)

type MOUSE_INPUT struct {
	Type DWORD
	Mi   MOUSEINPUT
}

type MOUSEINPUT struct {
	Dx          LONG
	Dy          LONG
	MouseData   DWORD
	DwFlags     DWORD
	Time        DWORD
	DwExtraInfo ULONG_PTR
}

var (
	// DLL
	user32 *windows.DLL

	// Functions
	enumWindows              *windows.Proc
	getWindowThreadProcessId *windows.Proc
	sendInput                *windows.Proc
	sendMessage              *windows.Proc
	setCursorPos             *windows.Proc
	setForegroundWindow      *windows.Proc
)

func init() {
	// DLL
	user32 = windows.MustLoadDLL("user32.dll")

	// Functions
	enumWindows = user32.MustFindProc("EnumWindows")
	getWindowThreadProcessId = user32.MustFindProc("GetWindowThreadProcessId")
	sendInput = user32.MustFindProc("SendInput")
	sendMessage = user32.MustFindProc("SendMessageW")
	setCursorPos = user32.MustFindProc("SetCursorPos")
	setForegroundWindow = user32.MustFindProc("SetForegroundWindow")
}

func EnumWindows(lpEnumFunc WNDENUMPROC, lParam LPARAM) BOOL {
	ret, _, _ := enumWindows.Call(lpEnumFunc, uintptr(lParam))
	return BOOL(ret)
}

func GetWindowThreadProcessId(hWnd HWND, lpdwProcessId LPDWORD) DWORD {
	ret, _, _ := getWindowThreadProcessId.Call(hWnd, uintptr(unsafe.Pointer(lpdwProcessId)))
	return DWORD(ret)
}

func SendInput(nInputs UINT, pInputs LPINPUT, cbSize int) UINT {
	ret, _, _ := sendInput.Call(uintptr(nInputs), uintptr(pInputs), uintptr(cbSize))
	return UINT(ret)
}

func SendMessage(hWnd HWND, msg UINT, wParam WPARAM, lParam LPARAM) LRESULT {
	ret, _, _ := sendMessage.Call(hWnd, uintptr(msg), uintptr(wParam), uintptr(lParam))
	return LRESULT(ret)
}

func SetCursorPos(x, y int32) BOOL {
	ret, _, _ := setCursorPos.Call(uintptr(x), uintptr(y))
	return BOOL(ret)
}

func SetForegroundWindow(hWnd HWND) BOOL {
	ret, _, _ := setForegroundWindow.Call(hWnd)
	return BOOL(ret)
}
