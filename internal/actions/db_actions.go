package actions

import (
	"bufio"
	"fmt"
	"github.com/lexionq/mango/internal/database"
	"github.com/lexionq/mango/internal/password"
	"golang.org/x/term"
	"os"
	"strconv"
	"strings"
	"syscall"
	"time"
)

func AddPasswordDB(db *database.EncryptedDatabase) {
	var entry database.PasswordEntry
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Register Name: ")
	input, _ := reader.ReadString('\n')
	entry.Name = strings.TrimSpace(input)

	fmt.Print("Website: ")
	input, _ = reader.ReadString('\n')
	entry.Website = strings.TrimSpace(input)

	fmt.Print("Username: ")
	input, _ = reader.ReadString('\n')
	entry.Username = strings.TrimSpace(input)

	fmt.Print("Password: ")
	passBytes, err := term.ReadPassword(int(syscall.Stdin))
	fmt.Println()

	if err != nil {
		fmt.Println("Error reading password:", err)
		return
	}

	entry.Password = string(passBytes)

	// Check password strength
	fmt.Println("\n🔍 Checking password strength...")
	password.DisplayPasswordStrength(entry.Password)

	fmt.Print("Notes (optional): ")
	input, _ = reader.ReadString('\n')
	entry.Notes = strings.TrimSpace(input)

	id, err := db.AddPassword(entry)
	if err != nil {
		fmt.Printf("[!] Error adding password: %v\n", err)
		return
	}

	fmt.Printf("[✔] Password added successfully! (ID: %d)\n", id)
}

func ListPasswordsDB(db *database.EncryptedDatabase) {
	passwords, err := db.GetAllPasswords()
	if err != nil {
		fmt.Printf("[!] Error listing passwords: %v\n", err)
		return
	}

	if len(passwords) == 0 {
		fmt.Println("[!] No passwords found.")
		return
	}

	fmt.Printf("\n[📋] Your Passwords (%d total):\n", len(passwords))
	fmt.Println("=====================================")

	for _, p := range passwords {
		fmt.Printf("ID: %d\n", p.ID)
		fmt.Printf("Name: %s\n", p.Name)
		if p.Website != "" {
			fmt.Printf("Website: %s\n", p.Website)
		}
		if p.Username != "" {
			fmt.Printf("Username: %s\n", p.Username)
		}
		fmt.Printf("Password: %s\n", p.Password)
		if p.Notes != "" {
			fmt.Printf("Notes: %s\n", p.Notes)
		}
		fmt.Printf("Created: %s\n", p.CreatedAt.Format("2006-01-02 15:04:05"))
		if !p.UpdatedAt.Equal(p.CreatedAt) {
			fmt.Printf("Updated: %s\n", p.UpdatedAt.Format("2006-01-02 15:04:05"))
		}
		fmt.Println("-------------------------------------")
	}

	fmt.Println("[✔] Passwords listed successfully!")
}

func SearchPasswordsDB(db *database.EncryptedDatabase, keyword string) {
	if keyword == "" {
		fmt.Println("[!] Error: Please provide a search term")
		return
	}

	passwords, err := db.SearchPasswords(keyword)
	if err != nil {
		fmt.Printf("[!] Error searching passwords: %v\n", err)
		return
	}

	if len(passwords) == 0 {
		fmt.Printf("[!] No passwords found matching '%s'\n", keyword)
		return
	}

	fmt.Printf("\n[🔍] Search Results for '%s' (%d found):\n", keyword, len(passwords))
	fmt.Println("=====================================")

	for _, p := range passwords {
		fmt.Printf("ID: %d | %s", p.ID, p.Name)
		if p.Website != "" {
			fmt.Printf(" (%s)", p.Website)
		}
		if p.Username != "" {
			fmt.Printf(" - %s", p.Username)
		}
		fmt.Println()
	}

	fmt.Println("[✔] Search completed successfully!")
}

func DeletePasswordDB(db *database.EncryptedDatabase) {
	passwords, err := db.GetAllPasswords()
	if err != nil {
		fmt.Printf("[!] Error getting passwords: %v\n", err)
		return
	}

	if len(passwords) == 0 {
		fmt.Println("[!] No passwords found.")
		return
	}

	// Display all passwords with IDs
	fmt.Println("\n[📋] Your Passwords:")
	for _, p := range passwords {
		fmt.Printf("%d. %s", p.ID, p.Name)
		if p.Website != "" {
			fmt.Printf(" (%s)", p.Website)
		}
		fmt.Println()
	}

	// Get user selection
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\nEnter the ID of password to delete (0 to cancel): ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if input == "0" {
		fmt.Println("[!] Operation cancelled.")
		return
	}

	id, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("[!] Invalid ID.")
		return
	}

	// Get password details for confirmation
	password, err := db.GetPasswordByID(id)
	if err != nil {
		fmt.Printf("[!] Password with ID %d not found.\n", id)
		return
	}

	// Confirm deletion
	fmt.Printf("Are you sure you want to delete '%s'? (y/N): ", password.Name)
	confirm, _ := reader.ReadString('\n')
	confirm = strings.ToLower(strings.TrimSpace(confirm))

	if confirm != "y" && confirm != "yes" {
		fmt.Println("[!] Operation cancelled.")
		return
	}

	// Delete password
	err = db.DeletePassword(id)
	if err != nil {
		fmt.Printf("[!] Error deleting password: %v\n", err)
		return
	}

	fmt.Printf("[✔] Password '%s' deleted successfully!\n", password.Name)
}

func CopyPasswordDB(db *database.EncryptedDatabase) {
	passwords, err := db.GetAllPasswords()
	if err != nil {
		fmt.Printf("[!] Error getting passwords: %v\n", err)
		return
	}

	if len(passwords) == 0 {
		fmt.Println("[!] No passwords found.")
		return
	}

	// Display all passwords with IDs
	fmt.Println("\n[📋] Your Passwords:")
	for _, p := range passwords {
		fmt.Printf("%d. %s", p.ID, p.Name)
		if p.Website != "" {
			fmt.Printf(" (%s)", p.Website)
		}
		fmt.Println()
	}

	// Get user selection
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\nEnter the ID of password to copy (0 to cancel): ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if input == "0" {
		fmt.Println("[!] Operation cancelled.")
		return
	}

	id, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("[!] Invalid ID.")
		return
	}

	// Get password
	password, err := db.GetPasswordByID(id)
	if err != nil {
		fmt.Printf("[!] Password with ID %d not found.\n", id)
		return
	}

	// Copy to clipboard
	err = copyToClipboard(password.Password)
	if err != nil {
		fmt.Printf("[!] Error copying to clipboard: %v\n", err)
		fmt.Printf("[📋] Password for '%s': %s\n", password.Name, password.Password)
		return
	}

	fmt.Printf("[✔] Password for '%s' copied to clipboard!\n", password.Name)
	fmt.Println("[⏰] Password will be cleared from clipboard in 30 seconds...")

	// Clear clipboard after 30 seconds
	go func() {
		time.Sleep(30 * time.Second)
		copyToClipboard("")
		fmt.Println("\n[🔒] Clipboard cleared for security.")
	}()
}

func StatsDB(db *database.EncryptedDatabase) {
	stats, err := db.GetStats()
	if err != nil {
		fmt.Printf("[!] Error getting statistics: %v\n", err)
		return
	}

	dbSize, err := db.GetDatabaseSize()
	if err != nil {
		fmt.Printf("[!] Error getting database size: %v\n", err)
		return
	}

	// Display statistics
	fmt.Println("\n🥭 Mango Password Manager Statistics")
	fmt.Println("=====================================")
	fmt.Printf("📊 Total Passwords: %d\n", stats["total_passwords"])
	fmt.Printf("🌐 Unique Websites: %d\n", stats["unique_websites"])
	fmt.Printf("⚠️  Duplicate Passwords: %d\n", stats["duplicate_passwords"])
	fmt.Printf("💾 Database Size: %.2f KB\n", float64(dbSize)/1024)

	if lastCreated, ok := stats["last_created"].(time.Time); ok {
		fmt.Printf("📅 Last Entry: %s\n", lastCreated.Format("2006-01-02 15:04:05"))
	}

	if stats["duplicate_passwords"].(int) > 0 {
		fmt.Println("\n⚠️  Security Warning: You have duplicate passwords!")
		fmt.Println("   Consider using unique passwords for better security.")
	}

	// Show most common websites
	if commonWebsites, ok := stats["common_websites"].([]map[string]interface{}); ok && len(commonWebsites) > 0 {
		fmt.Println("\n🏆 Most Used Websites:")
		for _, website := range commonWebsites {
			if website["count"].(int) > 1 {
				fmt.Printf("   %s: %d accounts\n", website["website"], website["count"])
			}
		}
	}

	fmt.Printf("\n⏰ Report generated at: %s\n", time.Now().Format("2006-01-02 15:04:05"))
}

func EditPasswordDB(db *database.EncryptedDatabase) {
	passwords, err := db.GetAllPasswords()
	if err != nil {
		fmt.Printf("[!] Error getting passwords: %v\n", err)
		return
	}

	if len(passwords) == 0 {
		fmt.Println("[!] No passwords found.")
		return
	}

	// Display all passwords with IDs
	fmt.Println("\n[📋] Your Passwords:")
	for _, p := range passwords {
		fmt.Printf("%d. %s", p.ID, p.Name)
		if p.Website != "" {
			fmt.Printf(" (%s)", p.Website)
		}
		fmt.Println()
	}

	// Get user selection
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\nEnter the ID of password to edit (0 to cancel): ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if input == "0" {
		fmt.Println("[!] Operation cancelled.")
		return
	}

	id, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("[!] Invalid ID.")
		return
	}

	// Get existing password
	existing, err := db.GetPasswordByID(id)
	if err != nil {
		fmt.Printf("[!] Password with ID %d not found.\n", id)
		return
	}

	fmt.Printf("\n[📝] Editing '%s' (press Enter to keep current value):\n", existing.Name)

	var updated database.PasswordEntry = *existing

	// Edit fields
	fmt.Printf("Name [%s]: ", existing.Name)
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input != "" {
		updated.Name = input
	}

	fmt.Printf("Website [%s]: ", existing.Website)
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input != "" {
		updated.Website = input
	}

	fmt.Printf("Username [%s]: ", existing.Username)
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input != "" {
		updated.Username = input
	}

	fmt.Print("Password (press Enter to keep current): ")
	passBytes, err := term.ReadPassword(int(syscall.Stdin))
	fmt.Println()
	if err != nil {
		fmt.Println("Error reading password:", err)
		return
	}
	if len(passBytes) > 0 {
		updated.Password = string(passBytes)
		// Check password strength for new password
		fmt.Println("\n🔍 Checking new password strength...")
		password.DisplayPasswordStrength(updated.Password)
	}

	fmt.Printf("Notes [%s]: ", existing.Notes)
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input != "" {
		updated.Notes = input
	}

	// Update password
	err = db.UpdatePassword(id, updated)
	if err != nil {
		fmt.Printf("[!] Error updating password: %v\n", err)
		return
	}

	fmt.Printf("[✔] Password '%s' updated successfully!\n", updated.Name)
}