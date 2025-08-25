package actions

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/lexionq/mango/internal/crypt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func ExportCSV(sourcePath string, key []byte) {
	// Read and decrypt data
	ciphertext, err := os.ReadFile(sourcePath)
	if err != nil {
		fmt.Println("[!] Error reading file:", err)
		return
	}

	plaintext := crypt.Decrypt(ciphertext, key)
	if plaintext == nil {
		fmt.Println("[!] No data to export.")
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
		fmt.Println("[!] No registers to export.")
		return
	}

	// Get export path from user
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter destination directory for CSV export: ")
	destDir, _ := reader.ReadString('\n')
	destDir = strings.TrimSpace(destDir)

	if destDir == "" {
		homeDir, _ := os.UserHomeDir()
		destDir = homeDir
	}

	// Create filename with timestamp
	timestamp := time.Now().Format("2006-01-02_15-04-05")
	filename := fmt.Sprintf("mango_export_%s.csv", timestamp)
	destPath := filepath.Join(destDir, filename)

	// Create CSV file
	file, err := os.Create(destPath)
	if err != nil {
		fmt.Println("[!] Error creating CSV file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV header
	header := []string{"Name", "Website", "Username", "Password", "Notes"}
	if err := writer.Write(header); err != nil {
		fmt.Println("[!] Error writing CSV header:", err)
		return
	}

	// Write data rows
	for _, line := range validLines {
		parts := strings.Split(line, ",")

		// Ensure we have at least 4 parts (name, website, username, password)
		for len(parts) < 5 {
			parts = append(parts, "")
		}

		// Clean up the parts
		for i := range parts {
			parts[i] = strings.TrimSpace(parts[i])
		}

		if err := writer.Write(parts); err != nil {
			fmt.Println("[!] Error writing CSV row:", err)
			return
		}
	}

	fmt.Printf("[✔] CSV export completed successfully!\n")
	fmt.Printf("[📁] File saved as: %s\n", destPath)
	fmt.Printf("[📊] Exported %d registers\n", len(validLines))

	// Security warning
	fmt.Println("\n⚠️  Security Warning:")
	fmt.Println("   The exported CSV file contains unencrypted passwords.")
	fmt.Println("   Please store it securely and delete it when no longer needed.")
}

func ImportCSV(destPath string, key []byte) {
	// Get import path from user
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter path to CSV file for import: ")
	csvPath, _ := reader.ReadString('\n')
	csvPath = strings.TrimSpace(csvPath)

	// Open CSV file
	file, err := os.Open(csvPath)
	if err != nil {
		fmt.Println("[!] Error opening CSV file:", err)
		return
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println("[!] Error reading CSV file:", err)
		return
	}

	if len(records) < 2 {
		fmt.Println("[!] CSV file appears to be empty or invalid.")
		return
	}

	// Skip header row and process data
	var newContent []string
	importCount := 0

	for i, record := range records[1:] {
		if len(record) < 4 {
			fmt.Printf("[!] Skipping row %d: insufficient data\n", i+2)
			continue
		}

		// Ensure we have 5 fields (name, website, username, password, notes)
		for len(record) < 5 {
			record = append(record, "")
		}

		// Clean and validate data
		name := strings.TrimSpace(record[0])
		website := strings.TrimSpace(record[1])
		username := strings.TrimSpace(record[2])
		password := strings.TrimSpace(record[3])
		notes := strings.TrimSpace(record[4])

		if name == "" || password == "" {
			fmt.Printf("[!] Skipping row %d: missing name or password\n", i+2)
			continue
		}

		// Format as mango format: name,website,username,password,notes
		line := fmt.Sprintf("%s,%s,%s,%s,%s", name, website, username, password, notes)
		newContent = append(newContent, line)
		importCount++
	}

	if importCount == 0 {
		fmt.Println("[!] No valid records found to import.")
		return
	}

	// Read existing data
	var existingContent []byte
	if ciphertext, err := os.ReadFile(destPath); err == nil && len(ciphertext) > 0 {
		existingContent = crypt.Decrypt(ciphertext, key)
	}

	// Combine existing and new content
	var finalContent string
	if len(existingContent) > 0 {
		finalContent = string(existingContent)
		if !strings.HasSuffix(finalContent, "\n") {
			finalContent += "\n"
		}
	}

	finalContent += strings.Join(newContent, "\n") + "\n"

	// Encrypt and save
	encryptedContent := crypt.Encrypt([]byte(finalContent), key)

	f, err := os.OpenFile(destPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("[!] Error opening destination file:", err)
		return
	}
	defer f.Close()

	_, err = f.Write(encryptedContent)
	if err != nil {
		fmt.Println("[!] Error writing file:", err)
		return
	}

	fmt.Printf("[✔] CSV import completed successfully!\n")
	fmt.Printf("[📊] Imported %d registers\n", importCount)
}
