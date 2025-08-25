package actions

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func Restore(targetPath string) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("[!] Error getting home directory:", err)
		return
	}

	backupDir := filepath.Join(homeDir, ".mango", "backups")

	// List available backups
	files, err := os.ReadDir(backupDir)
	if err != nil {
		fmt.Println("[!] No backups found or error reading backup directory.")
		return
	}

	var backups []string
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".mango" {
			backups = append(backups, file.Name())
		}
	}

	if len(backups) == 0 {
		fmt.Println("[!] No backups found.")
		return
	}

	// Display available backups
	fmt.Println("\n[📦] Available Backups:")
	for i, backup := range backups {
		fmt.Printf("%d. %s\n", i+1, backup)
	}

	// Get user selection
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\nEnter the number of backup to restore (0 to cancel): ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if input == "0" {
		fmt.Println("[!] Operation cancelled.")
		return
	}

	index, err := strconv.Atoi(input)
	if err != nil || index < 1 || index > len(backups) {
		fmt.Println("[!] Invalid selection.")
		return
	}

	selectedBackup := backups[index-1]
	backupPath := filepath.Join(backupDir, selectedBackup)

	// Confirm restoration
	fmt.Printf("⚠️  This will overwrite your current password database.\n")
	fmt.Printf("Are you sure you want to restore from '%s'? (y/N): ", selectedBackup)
	confirm, _ := reader.ReadString('\n')
	confirm = strings.ToLower(strings.TrimSpace(confirm))

	if confirm != "y" && confirm != "yes" {
		fmt.Println("[!] Operation cancelled.")
		return
	}

	// Create backup of current file before restoring
	currentBackupName := fmt.Sprintf("pre_restore_backup_%d.mango", len(backups)+1)
	currentBackupPath := filepath.Join(backupDir, currentBackupName)

	if currentFile, err := os.Open(targetPath); err == nil {
		defer currentFile.Close()
		if currentBackup, err := os.Create(currentBackupPath); err == nil {
			defer currentBackup.Close()
			io.Copy(currentBackup, currentFile)
			fmt.Printf("[✔] Current database backed up as: %s\n", currentBackupName)
		}
	}

	// Restore from backup
	backupFile, err := os.Open(backupPath)
	if err != nil {
		fmt.Println("[!] Error opening backup file:", err)
		return
	}
	defer backupFile.Close()

	targetFile, err := os.Create(targetPath)
	if err != nil {
		fmt.Println("[!] Error creating target file:", err)
		return
	}
	defer targetFile.Close()

	// Copy backup contents to target
	_, err = io.Copy(targetFile, backupFile)
	if err != nil {
		fmt.Println("[!] Error restoring file:", err)
		return
	}

	fmt.Printf("[✔] Database restored successfully from: %s\n", selectedBackup)
	fmt.Println("[⚠️] Please verify your master password works with the restored database.")
}
