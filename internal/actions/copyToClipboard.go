package actions

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/atotto/clipboard"
	"github.com/lexionq/mango/internal/crypt"
)

func CopyToClipboard(path string, key []byte) error {
	ciphertext, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("[!] failed to read file: %w", err)
	}

	plaintext := crypt.Decrypt(ciphertext, key)
	if plaintext == nil {
		return fmt.Errorf("[!] Error: Decryption failed or returned nil")
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Register Name (or search term): ")
	searchTerm, _ := reader.ReadString('\n')
	searchTerm = Clean(strings.TrimSpace(searchTerm))

	if searchTerm == "" {
		return fmt.Errorf("[!] Error: Register name cannot be empty")
	}

	text := strings.Split(string(plaintext), "\n")
	var foundLine string
	var foundIndex int = -1

	for i, line := range text {
		if Clean(line) == "" {
			continue
		}
		fields := strings.Split(line, ",")
		if len(fields) < 4 {
			continue
		}
		if strings.Contains(Clean(fields[0]), searchTerm) || 
		   strings.Contains(Clean(fields[1]), searchTerm) ||
		   strings.Contains(Clean(fields[2]), searchTerm) {
			foundLine = line
			foundIndex = i
			break
		}
	}

	if foundIndex == -1 {
		return fmt.Errorf("[!] Error: Register not found")
	}

	fields := strings.Split(foundLine, ",")
	if len(fields) < 4 {
		return fmt.Errorf("[!] Error: Password field not found")
	}

	password := Clean(fields[3])
	clipboard.WriteAll(password)
	fmt.Println("[✔] Password copied succesfully!")
	fmt.Println("[*] Clipboard will be cleared after 10 seconds...")
	
	go func() {
		time.Sleep(10 * time.Second)
		clipboard.WriteAll("")
		fmt.Println("[✔] Clipboard cleared!")
	}()
	
	return nil
}