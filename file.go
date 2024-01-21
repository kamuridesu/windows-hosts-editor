package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"syscall"

	"golang.org/x/sys/windows"
)

func isAdmin() bool {
	_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
	if err != nil {
		return false
	}
	return true
}

func runAsAdmin() {
	verb := "runas"
	exe, _ := os.Executable()
	cwd, _ := os.Getwd()
	args := strings.Join(os.Args[1:], " ")

	verbPtr, _ := syscall.UTF16PtrFromString(verb)
	exePtr, _ := syscall.UTF16PtrFromString(exe)
	cwdPtr, _ := syscall.UTF16PtrFromString(cwd)
	argPtr, _ := syscall.UTF16PtrFromString(args)

	var showCmd int32 = 1 //SW_NORMAL

	err := windows.ShellExecute(0, verbPtr, exePtr, argPtr, cwdPtr, showCmd)
	if err != nil {
		fmt.Println(err)
	}
}

func OpenFile(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func SaveFile(filename, content string) {
	if err := os.WriteFile(filename, []byte(content), os.ModeAppend); err != nil {
		panic(err)
	}
}

func CheckAndCreateTempFile() {
	if _, err := os.Stat(TempFileLocation); err == nil {
		SaveFile(TempFileLocation, "")
	} else if errors.Is(err, os.ErrNotExist) {
		os.WriteFile(TempFileLocation, []byte(""), 0644)
	} else {
		panic(err)
	}
}
