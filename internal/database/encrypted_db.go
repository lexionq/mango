package database

import (
	"fmt"
	"github.com/lexionq/mango/internal/crypt"
)

type EncryptedDatabase struct {
	db  *Database
	key []byte
}

func NewEncryptedDatabase(dbPath string, key []byte) (*EncryptedDatabase, error) {
	db, err := NewDatabase(dbPath)
	if err != nil {
		return nil, err
	}

	return &EncryptedDatabase{
		db:  db,
		key: key,
	}, nil
}

func (edb *EncryptedDatabase) Close() error {
	return edb.db.Close()
}

func (edb *EncryptedDatabase) AddPassword(entry PasswordEntry) (int64, error) {
	// Encrypt sensitive fields
	encryptedEntry := entry
	encryptedEntry.Password = string(crypt.Encrypt([]byte(entry.Password), edb.key))
	if entry.Notes != "" {
		encryptedEntry.Notes = string(crypt.Encrypt([]byte(entry.Notes), edb.key))
	}

	return edb.db.AddPassword(encryptedEntry)
}

func (edb *EncryptedDatabase) GetAllPasswords() ([]PasswordEntry, error) {
	passwords, err := edb.db.GetAllPasswords()
	if err != nil {
		return nil, err
	}

	// Decrypt sensitive fields
	for i := range passwords {
		decryptedPassword := crypt.Decrypt([]byte(passwords[i].Password), edb.key)
		if decryptedPassword == nil {
			return nil, fmt.Errorf("failed to decrypt password for entry %d", passwords[i].ID)
		}
		passwords[i].Password = string(decryptedPassword)

		if passwords[i].Notes != "" {
			decryptedNotes := crypt.Decrypt([]byte(passwords[i].Notes), edb.key)
			if decryptedNotes != nil {
				passwords[i].Notes = string(decryptedNotes)
			}
		}
	}

	return passwords, nil
}

func (edb *EncryptedDatabase) GetPasswordByID(id int) (*PasswordEntry, error) {
	password, err := edb.db.GetPasswordByID(id)
	if err != nil {
		return nil, err
	}

	// Decrypt sensitive fields
	decryptedPassword := crypt.Decrypt([]byte(password.Password), edb.key)
	if decryptedPassword == nil {
		return nil, fmt.Errorf("failed to decrypt password for entry %d", password.ID)
	}
	password.Password = string(decryptedPassword)

	if password.Notes != "" {
		decryptedNotes := crypt.Decrypt([]byte(password.Notes), edb.key)
		if decryptedNotes != nil {
			password.Notes = string(decryptedNotes)
		}
	}

	return password, nil
}

func (edb *EncryptedDatabase) UpdatePassword(id int, entry PasswordEntry) error {
	// Encrypt sensitive fields
	encryptedEntry := entry
	encryptedEntry.Password = string(crypt.Encrypt([]byte(entry.Password), edb.key))
	if entry.Notes != "" {
		encryptedEntry.Notes = string(crypt.Encrypt([]byte(entry.Notes), edb.key))
	}

	return edb.db.UpdatePassword(id, encryptedEntry)
}

func (edb *EncryptedDatabase) DeletePassword(id int) error {
	return edb.db.DeletePassword(id)
}

func (edb *EncryptedDatabase) SearchPasswords(keyword string) ([]PasswordEntry, error) {
	// For search, we need to decrypt all passwords and search in memory
	// This is less efficient but necessary for encrypted data
	allPasswords, err := edb.GetAllPasswords()
	if err != nil {
		return nil, err
	}

	var results []PasswordEntry
	for _, password := range allPasswords {
		if containsKeyword(password, keyword) {
			results = append(results, password)
		}
	}

	return results, nil
}

func containsKeyword(entry PasswordEntry, keyword string) bool {
	fields := []string{
		entry.Name,
		entry.Website,
		entry.Username,
		entry.Notes,
	}

	for _, field := range fields {
		if len(field) > 0 && contains(field, keyword) {
			return true
		}
	}
	return false
}

func contains(s, substr string) bool {
	// Simple case-insensitive contains
	s = toLower(s)
	substr = toLower(substr)
	return len(s) >= len(substr) && (len(substr) == 0 || indexOf(s, substr) >= 0)
}

func toLower(s string) string {
	result := make([]rune, len(s))
	for i, r := range s {
		if r >= 'A' && r <= 'Z' {
			result[i] = r + 32
		} else {
			result[i] = r
		}
	}
	return string(result)
}

func indexOf(s, substr string) int {
	if len(substr) == 0 {
		return 0
	}
	if len(substr) > len(s) {
		return -1
	}
	
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}

func (edb *EncryptedDatabase) GetStats() (map[string]interface{}, error) {
	return edb.db.GetStats()
}

func (edb *EncryptedDatabase) GetDatabaseSize() (int64, error) {
	return edb.db.GetDatabaseSize()
}