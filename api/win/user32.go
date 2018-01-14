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

const (
	WS_BORDER           = 0x00800000
	WS_CAPTION          = 0x00C00000
	WS_CHILD            = 0x40000000
	WS_CHILDWINDOW      = 0x40000000
	WS_CLIPCHILDREN     = 0x02000000
	WS_CLIPSIBLINGS     = 0x04000000
	WS_DISABLED         = 0x08000000
	WS_DLGRAME          = 0x00400000
	WS_GROUP            = 0x00200000
	WS_HSCROLL          = 0x00100000
	WS_ICONIC           = 0x20000000
	WS_MAXIMIZE         = 0x01000000
	WS_MAXIMIZEBOX      = 0x00010000
	WS_MINIMIZE         = 0x20000000
	WS_MINIMIZEBOX      = 0x00020000
	WS_OVERLAPPED       = 0x00000000
	WS_OVERLAPPEDWINDOW = WS_OVERLAPPED | WS_CAPTION | WS_SYSMENU | WS_THICKFRAME | WS_MINIMIZEBOX | WS_MAXIMIZEBOX
	WS_POPUP            = 0x80000000
	WS_POPUPWINDOW      = WS_POPUP | WS_BORDER | WS_SYSMENU
	WS_SIZEBOX          = 0x00040000
	WS_SYSMENU          = 0x00080000
	WS_TABSTOP          = 0x00010000
	WS_THICKFRAME       = 0x00040000
	WS_TILED            = 0x00000000
	WS_TILEDWINDOW      = WS_OVERLAPPED | WS_CAPTION | WS_SYSMENU | WS_THICKFRAME | WS_MINIMIZEBOX | WS_MAXIMIZEBOX
	WS_VISIBLE          = 0x10000000
	WS_VSCROLL          = 0x00200000

	WS_EX_ACCEPTFILES         = 0x00000010
	WS_EX_APPWINDOW           = 0x00040000
	WS_EX_CLIENTEDGE          = 0x00000200
	WS_EX_COMPOSITED          = 0x02000000
	WS_EX_CONTEXTHELP         = 0x00000400
	WS_EX_CONTROLPARENT       = 0x00010000
	WS_EX_DLGMODALFRAME       = 0x00000001
	WS_EX_LAYERED             = 0x00080000
	WS_EX_LAYOUTRTL           = 0x00400000
	WS_EX_LEFT                = 0x00000000
	WS_EX_LEFTSCROLLBAR       = 0x00004000
	WS_EX_LTRREADING          = 0x00000000
	WS_EX_MDICHILD            = 0x00000040
	WS_EX_NOACTIVATE          = 0x08000000
	WS_EX_NOINHERITLAYOUT     = 0x00100000
	WS_EX_NOPARENTNOTIFY      = 0x00000004
	WS_EX_NOREDIRECTIONBITMAP = 0x00200000
	WS_EX_OVERLAPPEDWINDOW    = WS_EX_WINDOWEDGE | WS_EX_CLIENTEDGE
	WS_EX_PALETTEWINDOW       = WS_EX_WINDOWEDGE | WS_EX_TOOLWINDOW | WS_EX_TOPMOST
	WS_EX_RIGHT               = 0x00001000
	WS_EX_RIGHTSCROLLBAR      = 0x00000000
	WS_EX_RTLREADING          = 0x00002000
	WS_EX_STATICEDGE          = 0x00020000
	WS_EX_TOOLWINDOW          = 0x00000080
	WS_EX_TOPMOST             = 0x00000008
	WS_EX_TRANSPARENT         = 0x00000020
	WS_EX_WINDOWEDGE          = 0x00000100
)

type (
	HWND        = HANDLE
	LPINPUT     = unsafe.Pointer
	LPRECT      = *RECT
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

type RECT struct {
	Left   LONG
	Top    LONG
	Right  LONG
	Bottom LONG
}

var (
	// DLL
	user32 *windows.DLL

	// Functions
	createWindowEx           *windows.Proc
	enumWindows              *windows.Proc
	getDesktopWindow         *windows.Proc
	getWindowRect            *windows.Proc
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
	createWindowEx = user32.MustFindProc("CreateWindowExW")
	enumWindows = user32.MustFindProc("EnumWindows")
	getDesktopWindow = user32.MustFindProc("GetDesktopWindow")
	getWindowRect = user32.MustFindProc("GetWindowRect")
	getWindowThreadProcessId = user32.MustFindProc("GetWindowThreadProcessId")
	sendInput = user32.MustFindProc("SendInput")
	sendMessage = user32.MustFindProc("SendMessageW")
	setCursorPos = user32.MustFindProc("SetCursorPos")
	setForegroundWindow = user32.MustFindProc("SetForegroundWindow")
}

func CreateWindowEx(dwExStyle DWORD, lpClassName, lpWindowName LPCTSTR, dwStyle DWORD, x, y, nWidth, nHeight int, hWndParent HWND, hMenu HMENU, hInstance HINSTANCE, lpParam LPVOID) HWND {
	ret, _, _ := createWindowEx.Call(dwExStyle, lpClassName, lpWindowName, dwStyle, x, y, nWidth, nHeight, hWndParent, hMenu, hInstance, lpParam)
	return ret
}

func EnumWindows(lpEnumFunc WNDENUMPROC, lParam LPARAM) BOOL {
	ret, _, _ := enumWindows.Call(lpEnumFunc, uintptr(lParam))
	return BOOL(ret)
}

func GetDesktopWindow() HWND {
	ret, _, _ := getDesktopWindow.Call()
	return ret
}

func GetWindowRect(hWnd HWND, lpRect LPRECT) BOOL {
	ret, _, _ := getWindowRect.Call(hWnd, uintptr(unsafe.Pointer(lpRect)))
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
