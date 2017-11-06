package win

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

const MAX_PATH = 260
const NULL = 0

// Standard access rights
const (
	DELETE       = 0x00010000
	READ_CONTROL = 0x00020000
	SYNCHRONIZE  = 0x00100000
	WRITE_DAC    = 0x00040000
	WRITE_OWNER  = 0x00080000
)

// Process-specific access rights
const (
	PROCESS_ALL_ACCESS                = 0x001F0FFF
	PROCESS_CREATE_PROCESS            = 0x00000080
	PROCESS_CREATE_THREAD             = 0x00000002
	PROCESS_DUP_HANDLE                = 0x00000040
	PROCESS_QUERY_INFORMATION         = 0x00000400
	PROCESS_QUERY_LIMITED_INFORMATION = 0x00001000
	PROCESS_SET_INFORMATION           = 0x00000200
	PROCESS_SET_QUOTA                 = 0x00000100
	PROCESS_SUSPEND_RESUME            = 0x00000800
	PROCESS_TERMINATE                 = 0x00000001
	PROCESS_VM_OPERATION              = 0x00000008
	PROCESS_VM_READ                   = 0x00000010
	PROCESS_VM_WRITE                  = 0x00000020
)

type (
	BOOL      = int
	DWORD     = uint32
	HANDLE    = uintptr
	HINSTANCE = HANDLE
	HMODULE   = HINSTANCE
	LONG      = int32
	LONG_PTR  = int64
	LPARAM    = LONG_PTR
	LPVOID    = unsafe.Pointer
	LPCVOID   = unsafe.Pointer
	LPDWORD   = *DWORD
	LPTSTR    = LPWSTR
	LPWSTR    = *WCHAR
	LRESULT   = LONG_PTR
	SHORT     = int16
	SIZE_T    = int64
	UINT      = uint32
	UINT_PTR  = uintptr
	ULONG_PTR = uintptr
	WCHAR     = []byte
	WORD      = uint16
	WPARAM    = UINT_PTR
)

var (
	// DLL
	kernel32 *windows.DLL

	// Functions
	closeHandle       *windows.Proc
	openProcess       *windows.Proc
	readProcessMemory *windows.Proc
)

func init() {
	// DLL
	kernel32 = windows.MustLoadDLL("kernel32.dll")

	// Functions
	closeHandle = kernel32.MustFindProc("CloseHandle")
	openProcess = kernel32.MustFindProc("OpenProcess")
	readProcessMemory = kernel32.MustFindProc("ReadProcessMemory")
}

func CloseHandle(hObject HANDLE) BOOL {
	ret, _, _ := closeHandle.Call(hObject)
	return BOOL(ret)
}

func OpenProcess(dwDesiredAccess DWORD, bInheritHandle BOOL, dwProcessId DWORD) HANDLE {
	ret, _, _ := openProcess.Call(uintptr(dwDesiredAccess), uintptr(bInheritHandle), uintptr(dwProcessId))
	return ret
}

func ReadProcessMemory(hProcess HANDLE, lpBaseAddress LPCVOID, lpBuffer LPVOID, nSize SIZE_T, lpNumberOfBytesRead *SIZE_T) BOOL {
	ret, _, _ := readProcessMemory.Call(hProcess, uintptr(lpBaseAddress), uintptr(lpBuffer), uintptr(nSize), uintptr(unsafe.Pointer(lpNumberOfBytesRead)))
	return BOOL(ret)
}
