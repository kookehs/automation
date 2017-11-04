package main

import (
	"fmt"
	"github.com/kookehs/automation/api/win"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: bot.exe <pid>")
		return
	}

	var pid win.DWORD

	if id, err := strconv.ParseUint(os.Args[1], 10, 32); err != nil {
		fmt.Println(err)
		return
	} else {
		pid = win.DWORD(id)
	}

	var access win.DWORD = win.PROCESS_VM_READ | win.PROCESS_QUERY_INFORMATION
	var inherit win.BOOL = 0
	handle := win.OpenProcess(access, inherit, pid)

	if handle == win.NULL {
		return
	}

	defer win.CloseHandle(handle)
	fmt.Println("handle: ", handle)

	fileName := make([]byte, win.MAX_PATH)
	win.GetModuleFileNameEx(handle, 0, &fileName, win.MAX_PATH)
	fmt.Println("file name: ", string(fileName))
}
