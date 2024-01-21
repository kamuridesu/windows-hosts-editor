package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var (
	iPeeAdress       string
	domainName       string
	editMode         bool
	TempFileLocation string
)

const HostsFileLocation = `C:\Windows\System32\drivers\etc\hosts`

func init() {
	TempFileLocation, _ = os.UserHomeDir()
	TempFileLocation += `\AppData\Local\Temp\hosts.edit.tmp`
	editMode = *flag.Bool("edit", true, "Edit mode, opens Notepad to edit hosts file")
	if !isAdmin() {
		runAsAdmin()
	}
	flag.Parse()
	args := flag.Args()
	if editMode {
		CheckAndCreateTempFile()
		return
	}
	if len(args) < 2 {
		fmt.Println("Erro: Expected IP and DNS in the following format: IP DNS")
		os.Exit(1)
	}
	iPeeAdress = args[0]
	domainName = args[1]
}

func main() {
	if !isAdmin() {
		fmt.Println("Need admin permission to modify hosts file")
		os.Exit(1)
	}
	fmt.Println("Running as admin")
	if !editMode {
		CLI()
		os.Exit(0)
	}
	notepadEdit()
}

func CLI() {
	content := OpenFile(HostsFileLocation)
	content += fmt.Sprintf("\n    %s    %s", iPeeAdress, domainName)
	_continue := "N"
	fmt.Print("Please, review the changes:\n\n")
	fmt.Print(content + "\n\n")
	fmt.Print("Do you wish to continue [y/N]: ")
	fmt.Scanln(&_continue)
	_continue = strings.ToLower(_continue)
	if _continue != "y" {
		fmt.Println("Aborted!")
		os.Exit(3)
	}
	SaveFile(HostsFileLocation, content)
}

func notepadEdit() {
	content := OpenFile(HostsFileLocation)
	SaveFile(TempFileLocation, content)
	cmd := exec.Command("notepad", TempFileLocation)
	if err := cmd.Start(); err != nil {
		panic(err)
	}
	fmt.Println("Waiting for notepad to be closed...")
	cmd.Wait()
	content = OpenFile(TempFileLocation)
	SaveFile(HostsFileLocation, content)
	os.Remove(TempFileLocation)
}
