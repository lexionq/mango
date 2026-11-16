package actions

import (
	"os"
)

func ClearScreen() {
	// CSI H  : cursor home
	// CSI 2J : erase the visible screen
	// CSI 3J : erase scrollback (otherwise many terminals only scroll up)
	_, _ = os.Stdout.Write([]byte("\033[H\033[2J\033[3J"))
	_ = os.Stdout.Sync()
}