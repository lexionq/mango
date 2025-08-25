package actions

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

func Backup(sourcePath string) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("[!] Error getting home directory:", err)
		return
	}

	// Create backup directory if it doesn't exist
	backupDir := filepath.Join(homeDir, ".mango", "backups")
	if err := os.MkdirAll(backupDir, 0700); err != nil {
		fmt.Println("[!] Error creating backup directory:", err)
		return
	}

	// Generate backup filename with timestamp
	timestamp := time.Now().Format("2006-01-02_15-04-05")
	backupFileName := fmt.Sprintf("mango_backup_%s.mango", timestamp)
	backupPath := filepath.Join(backupDir, backupFileName)

	// Copy the source file to backup location
	sourceFile, err := os.Open(sourcePath)
	if err != nil {
		fmt.Println("[!] Error opening source file:", err)
		return
	}
	defer sourceFile.Close()

	backupFile, err := os.Create(backupPath)
	if err != nil {
		fmt.Println("[!] Error creating backup file:", err)
		return
	}
	defer backupFile.Close()

	// Copy file contents
	_, err = io.Copy(backupFile, sourceFile)
	if err != nil {
		fmt.Println("[!] Error copying file:", err)
		return
	}

	fmt.Printf("[✔] Backup created successfully: %s\n", backupFileName)
	fmt.Printf("[📁] Backup location: %s\n", backupPath)

	// Also backup the password hash file
	hashSourcePath := filepath.Join(homeDir, ".mango", "pass.hash")
	hashBackupPath := filepath.Join(backupDir, fmt.Sprintf("pass_hash_%s.hash", timestamp))

	if hashSource, err := os.Open(hashSourcePath); err == nil {
		defer hashSource.Close()
		if hashBackup, err := os.Create(hashBackupPath); err == nil {
			defer hashBackup.Close()
			io.Copy(hashBackup, hashSource)
			fmt.Printf("[✔] Password hash backed up: pass_hash_%s.hash\n", timestamp)
		}
	}
}

func ListBackups() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("[!] Error getting home directory:", err)
		return
	}

	backupDir := filepath.Join(homeDir, ".mango", "backups")

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

	fmt.Println("\n[📦] Available Backups:")
	for i, backup := range backups {
		fmt.Printf("%d. %s\n", i+1, backup)
	}
}
