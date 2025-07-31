package utils

import (
	"fmt"
	"os"
)

func PathExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func PassesFileControl(path string){
	if !PathExists(path){
		os.Create(path)
	}
}

func DisplayHelpMessage(){
	fmt.Println(`Usage:
		mango [command]

	Available Commands:
		help        Show this help message
		add         Add new register
		list        List all of registers
		search	    Searches for the specified word among the registers
		edit        Edit registers(with nano)
		change 		Change your master password

	Flags:
		-h, --help  help for mango
	`)
}