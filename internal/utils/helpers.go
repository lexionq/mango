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
		search      Searches for the specified word among the registers
		edit        Edit registers (cross-platform editor)
		delete      Delete a register
		copy        Copy password to clipboard
		stats       Show password database statistics
		
	Password Management:
		generate    Generate a secure password
		strength    Check password strength
		change      Change your master password
		
	Backup & Export:
		backup      Create a backup of your database
		restore     Restore from a backup
		export      Export database (encrypted)
		exportcsv   Export to CSV format (unencrypted)
		import      Import database (encrypted)
		importcsv   Import from CSV format

	Flags:
		-h, --help  help for mango
		
	Examples:
		mango add                    # Add new password
		mango search google          # Search for 'google'
		mango generate 16            # Generate 16-char password
		mango strength mypass123     # Check password strength
		mango copy                   # Copy password to clipboard
		mango backup                 # Create backup
	`)
}
