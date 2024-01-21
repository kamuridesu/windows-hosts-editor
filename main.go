package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var iPeeAdress string
var DomainName string

func init() {
	if !isAdmin() {
		runAsAdmin()
	}
	flag.Parse()
	args := flag.Args()
	if len(args) < 2 {
		fmt.Println("Erro: Expected IP and DNS in the following format: IP DNS")
		os.Exit(1)
	}
	iPeeAdress = args[0]
	DomainName = args[1]
}

func main() {
	if !isAdmin() {
		fmt.Println("Need admin permission to modify hosts file")
		os.Exit(1)
	}
	content := OpenHostsFile()

	content += fmt.Sprintf("\n    %s    %s", iPeeAdress, DomainName)

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
	SaveHostsFile(content)
}
