package win

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

var (
	// DLL
	psapi *windows.DLL

	// Functions
	getModuleFileNameEx *windows.Proc
)

func init() {
	// DLL
	psapi = windows.MustLoadDLL("psapi.dll")

	// Functions
	getModuleFileNameEx = psapi.MustFindProc("GetModuleFileNameExW")
}

func GetModuleFileNameEx(hProcess HANDLE, hModule HMODULE, lpFilename LPTSTR, nSize DWORD) DWORD {
	ret, _, _ := getModuleFileNameEx.Call(hProcess, hModule, uintptr(unsafe.Pointer(&(*lpFilename)[0])), uintptr(nSize))
	return DWORD(ret)
}
