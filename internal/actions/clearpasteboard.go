package actions

import "github.com/atotto/clipboard"

func ClearPasteboard() error {
	return clipboard.WriteAll("")
}
