package actions

import (
	"fmt"
	"github.com/lexionq/mango/internal/crypt"
	"os"
	"strings"
	"time"
)

func Stats(path string, key []byte) {
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
	var websites []string
	var passwords []string
	duplicatePasswords := make(map[string]int)

	// Filter out empty lines and collect data
	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			validLines = append(validLines, line)
			parts := strings.Split(line, ",")
			if len(parts) >= 4 {
				websites = append(websites, parts[1])
				password := parts[3]
				passwords = append(passwords, password)
				duplicatePasswords[password]++
			}
		}
	}

	// Count unique websites
	uniqueWebsites := make(map[string]bool)
	for _, website := range websites {
		if website != "" {
			uniqueWebsites[website] = true
		}
	}

	// Count duplicate passwords
	duplicateCount := 0
	for _, count := range duplicatePasswords {
		if count > 1 {
			duplicateCount++
		}
	}

	// Get file info for last modified
	fileInfo, err := os.Stat(path)
	var lastModified string
	if err == nil {
		lastModified = fileInfo.ModTime().Format("2006-01-02 15:04:05")
	} else {
		lastModified = "Unknown"
	}

	// Display statistics
	fmt.Println("\n🥭 Mango Password Manager Statistics")
	fmt.Println("=====================================")
	fmt.Printf("📊 Total Registers: %d\n", len(validLines))
	fmt.Printf("🌐 Unique Websites: %d\n", len(uniqueWebsites))
	fmt.Printf("🔑 Total Passwords: %d\n", len(passwords))
	fmt.Printf("⚠️  Duplicate Passwords: %d\n", duplicateCount)
	fmt.Printf("📅 Last Modified: %s\n", lastModified)
	fmt.Printf("💾 Database Size: %.2f KB\n", float64(len(ciphertext))/1024)

	if duplicateCount > 0 {
		fmt.Println("\n⚠️  Security Warning: You have duplicate passwords!")
		fmt.Println("   Consider using unique passwords for better security.")
	}

	// Show most common websites (top 5)
	websiteCount := make(map[string]int)
	for _, website := range websites {
		if website != "" {
			websiteCount[website]++
		}
	}

	if len(websiteCount) > 0 {
		fmt.Println("\n🏆 Most Used Websites:")
		count := 0
		for website, freq := range websiteCount {
			if count >= 5 {
				break
			}
			if freq > 1 {
				fmt.Printf("   %s: %d accounts\n", website, freq)
			}
			count++
		}
	}

	fmt.Printf("\n⏰ Report generated at: %s\n", time.Now().Format("2006-01-02 15:04:05"))
}
