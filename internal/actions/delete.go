package actions

import (
	"bufio"
	"fmt"
	"github.com/lexionq/mango/internal/crypt"
	"os"
	"strconv"
	"strings"
)

func Delete(path string, key []byte) {
	ciphertext, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("[!] Error: ", err)
		return
	}

	// Check if file is empty
	if len(ciphertext) == 0 {
		fmt.Println("[!] No registers found.")
		return
	}

	plaintext := crypt.Decrypt(ciphertext, key)
	if plaintext == nil || len(plaintext) == 0 {
		fmt.Println("[!] No registers found.")
		return
	}

	lines := strings.Split(string(plaintext), "\n")
	var validLines []string

	// Filter out empty lines
	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			validLines = append(validLines, line)
		}
	}

	if len(validLines) == 0 {
		fmt.Println("[!] No registers found.")
		return
	}

	// Display all registers with numbers
	fmt.Println("\n[📋] Your Registers:")
	for i, line := range validLines {
		parts := strings.Split(line, ",")
		if len(parts) >= 2 {
			fmt.Printf("%d. %s (%s)\n", i+1, parts[0], parts[1])
		}
	}

	// Get user selection
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\nEnter the number of register to delete (0 to cancel): ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if input == "0" {
		fmt.Println("[!] Operation cancelled.")
		return
	}

	index, err := strconv.Atoi(input)
	if err != nil || index < 1 || index > len(validLines) {
		fmt.Println("[!] Invalid selection.")
		return
	}

	// Confirm deletion
	selectedLine := validLines[index-1]
	parts := strings.Split(selectedLine, ",")
	registerName := parts[0]

	fmt.Printf("Are you sure you want to delete '%s'? (y/N): ", registerName)
	confirm, _ := reader.ReadString('\n')
	confirm = strings.ToLower(strings.TrimSpace(confirm))

	if confirm != "y" && confirm != "yes" {
		fmt.Println("[!] Operation cancelled.")
		return
	}

	// Remove the selected line
	validLines = append(validLines[:index-1], validLines[index:]...)

	// Reconstruct the content
	newContent := strings.Join(validLines, "\n")
	if len(validLines) > 0 {
		newContent += "\n"
	}

	// Encrypt and save
	encryptedContent := crypt.Encrypt([]byte(newContent), key)

	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("[!] Error opening file:", err)
		return
	}
	defer f.Close()

	_, err = f.Write(encryptedContent)
	if err != nil {
		fmt.Println("[!] Error writing file:", err)
		return
	}

	fmt.Printf("[✔] Register '%s' deleted successfully!\n", registerName)
}
