package utils

import (
	"fmt"
	"os"
)

var mango = `
🥭🥭🥭🥭🥭🥭🥭🥭🥭🥭🥭🥭🥭🥭🥭🥭🥭🥭🥭🥭
🥭 _ __ ___   __ _ _ __   __ _  ___   🥭
🥭| '_ ` + "`" + ` _ \ / _` + "`" + ` | '_ \ / _` + "`" + ` |/ _ \  🥭
🥭| | | | | | (_| | | | | (_| | (_) | 🥭
🥭|_| |_| |_|\__,_|_| |_|\__, |\___/  🥭
🥭                       |___/        🥭
🥭🥭🥭🥭🥭🥭🥭🥭🥭🥭🥭🥭🥭🥭🥭🥭🥭🥭🥭🥭
`

func PathExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func PassesFileControl(path string) {
	if !PathExists(path) {
		os.Create(path)
	}
}

func DisplayHelpMessage() {
	fmt.Println(mango)
	fmt.Println(`Usage:
		mango [command]

	Available Commands:
		help        Show this help message
		add         Add new register
		list        List all of registers
		search	    Searches for the specified word among the registers
		edit        Edit registers(with nano)
		copy        Copy password to clipboard
		change 	    Change your master password
		generate    Password generator for you

	Flags:
		-h, --help  help for mango
	`)
}
