package actions

import (
	"bufio"
	"fmt"
	"github.com/lexionq/mango/internal/crypt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func copyToClipboard(text string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin": // macOS
		cmd = exec.Command("pbcopy")
	case "linux":
		// Try xclip first, then xsel
		if _, err := exec.LookPath("xclip"); err == nil {
			cmd = exec.Command("xclip", "-selection", "clipboard")
		} else if _, err := exec.LookPath("xsel"); err == nil {
			cmd = exec.Command("xsel", "--clipboard", "--input")
		} else {
			return fmt.Errorf("no clipboard utility found (install xclip or xsel)")
		}
	case "windows":
		cmd = exec.Command("clip")
	default:
		return fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}

	cmd.Stdin = strings.NewReader(text)
	return cmd.Run()
}

func Copy(path string, key []byte) {
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
	fmt.Print("\nEnter the number of register to copy password (0 to cancel): ")
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

	// Extract password from selected line
	selectedLine := validLines[index-1]
	parts := strings.Split(selectedLine, ",")

	if len(parts) < 4 {
		fmt.Println("[!] Invalid register format.")
		return
	}

	registerName := parts[0]
	password := parts[3]

	// Copy to clipboard
	err = copyToClipboard(password)
	if err != nil {
		fmt.Printf("[!] Error copying to clipboard: %v\n", err)
		fmt.Printf("[📋] Password for '%s': %s\n", registerName, password)
		return
	}

	fmt.Printf("[✔] Password for '%s' copied to clipboard!\n", registerName)
	fmt.Println("[⏰] Password will be cleared from clipboard in 30 seconds...")

	// Clear clipboard after 30 seconds
	go func() {
		time.Sleep(30 * time.Second)
		copyToClipboard("")
		fmt.Println("\n[🔒] Clipboard cleared for security.")
	}()
}
