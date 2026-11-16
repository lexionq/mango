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
	fmt.Print(`Usage:
  mango <command> [arguments]

Commands:
  help               Show help message
  add                Add a new register
  ls                 List all registers
  search <keyword>   Search registers by keyword
  edit               Edit registers with nano
  cp                 Copy password to clipboard
  clp                Clear clipboard
  cls                Clear terminal screen
  import             Import registers from any source
  export             Export all registers
  change             Change master password
  gen      <length>  Generate a password with length

Flags:
  -h, --help         Show help message

Examples:
  mango add
  mango ls
  mango search github
  mango gen 16
`)
}
