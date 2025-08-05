package actions

import (
	"fmt"
	"github.com/lexionq/mango/internal/crypt"
	"os"
)

func ListRegisters(path string, key []byte) string {
	ciphertext, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	plaintext := crypt.Decrypt(ciphertext, key)

	fmt.Println("[✔] Registers listed successfully!")

	return string(plaintext)

}
