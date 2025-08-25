package database

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	conn *sql.DB
	path string
}

type PasswordEntry struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Website   string    `json:"website"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Notes     string    `json:"notes"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewDatabase(dbPath string) (*Database, error) {
	// Ensure directory exists
	dir := filepath.Dir(dbPath)
	if err := os.MkdirAll(dir, 0700); err != nil {
		return nil, fmt.Errorf("failed to create directory: %w", err)
	}

	conn, err := sql.Open("sqlite3", dbPath+"?_foreign_keys=on")
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	db := &Database{
		conn: conn,
		path: dbPath,
	}

	if err := db.createTables(); err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to create tables: %w", err)
	}

	return db, nil
}

func (db *Database) createTables() error {
	query := `
	CREATE TABLE IF NOT EXISTS passwords (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		website TEXT,
		username TEXT,
		password TEXT NOT NULL,
		notes TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);

	CREATE INDEX IF NOT EXISTS idx_passwords_name ON passwords(name);
	CREATE INDEX IF NOT EXISTS idx_passwords_website ON passwords(website);
	CREATE INDEX IF NOT EXISTS idx_passwords_username ON passwords(username);

	-- Trigger to update updated_at timestamp
	CREATE TRIGGER IF NOT EXISTS update_passwords_timestamp 
	AFTER UPDATE ON passwords
	BEGIN
		UPDATE passwords SET updated_at = CURRENT_TIMESTAMP WHERE id = NEW.id;
	END;
	`

	_, err := db.conn.Exec(query)
	return err
}

func (db *Database) Close() error {
	if db.conn != nil {
		return db.conn.Close()
	}
	return nil
}

func (db *Database) AddPassword(entry PasswordEntry) (int64, error) {
	query := `
		INSERT INTO passwords (name, website, username, password, notes)
		VALUES (?, ?, ?, ?, ?)
	`
	
	result, err := db.conn.Exec(query, entry.Name, entry.Website, entry.Username, entry.Password, entry.Notes)
	if err != nil {
		return 0, fmt.Errorf("failed to insert password: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get last insert id: %w", err)
	}

	return id, nil
}

func (db *Database) GetAllPasswords() ([]PasswordEntry, error) {
	query := `
		SELECT id, name, website, username, password, notes, created_at, updated_at
		FROM passwords
		ORDER BY name ASC
	`

	rows, err := db.conn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query passwords: %w", err)
	}
	defer rows.Close()

	var passwords []PasswordEntry
	for rows.Next() {
		var p PasswordEntry
		err := rows.Scan(&p.ID, &p.Name, &p.Website, &p.Username, &p.Password, &p.Notes, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan password: %w", err)
		}
		passwords = append(passwords, p)
	}

	return passwords, rows.Err()
}

func (db *Database) GetPasswordByID(id int) (*PasswordEntry, error) {
	query := `
		SELECT id, name, website, username, password, notes, created_at, updated_at
		FROM passwords
		WHERE id = ?
	`

	var p PasswordEntry
	err := db.conn.QueryRow(query, id).Scan(&p.ID, &p.Name, &p.Website, &p.Username, &p.Password, &p.Notes, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("password with id %d not found", id)
		}
		return nil, fmt.Errorf("failed to get password: %w", err)
	}

	return &p, nil
}

func (db *Database) UpdatePassword(id int, entry PasswordEntry) error {
	query := `
		UPDATE passwords 
		SET name = ?, website = ?, username = ?, password = ?, notes = ?
		WHERE id = ?
	`

	result, err := db.conn.Exec(query, entry.Name, entry.Website, entry.Username, entry.Password, entry.Notes, id)
	if err != nil {
		return fmt.Errorf("failed to update password: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("password with id %d not found", id)
	}

	return nil
}

func (db *Database) DeletePassword(id int) error {
	query := `DELETE FROM passwords WHERE id = ?`

	result, err := db.conn.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete password: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("password with id %d not found", id)
	}

	return nil
}

func (db *Database) SearchPasswords(keyword string) ([]PasswordEntry, error) {
	query := `
		SELECT id, name, website, username, password, notes, created_at, updated_at
		FROM passwords
		WHERE name LIKE ? OR website LIKE ? OR username LIKE ? OR notes LIKE ?
		ORDER BY name ASC
	`

	searchTerm := "%" + keyword + "%"
	rows, err := db.conn.Query(query, searchTerm, searchTerm, searchTerm, searchTerm)
	if err != nil {
		return nil, fmt.Errorf("failed to search passwords: %w", err)
	}
	defer rows.Close()

	var passwords []PasswordEntry
	for rows.Next() {
		var p PasswordEntry
		err := rows.Scan(&p.ID, &p.Name, &p.Website, &p.Username, &p.Password, &p.Notes, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan password: %w", err)
		}
		passwords = append(passwords, p)
	}

	return passwords, rows.Err()
}

func (db *Database) GetStats() (map[string]interface{}, error) {
	stats := make(map[string]interface{})

	// Total passwords
	var totalCount int
	err := db.conn.QueryRow("SELECT COUNT(*) FROM passwords").Scan(&totalCount)
	if err != nil {
		return nil, fmt.Errorf("failed to get total count: %w", err)
	}
	stats["total_passwords"] = totalCount

	// Unique websites
	var uniqueWebsites int
	err = db.conn.QueryRow("SELECT COUNT(DISTINCT website) FROM passwords WHERE website != ''").Scan(&uniqueWebsites)
	if err != nil {
		return nil, fmt.Errorf("failed to get unique websites: %w", err)
	}
	stats["unique_websites"] = uniqueWebsites

	// Duplicate passwords
	var duplicatePasswords int
	err = db.conn.QueryRow(`
		SELECT COUNT(*) FROM (
			SELECT password FROM passwords 
			GROUP BY password 
			HAVING COUNT(*) > 1
		)
	`).Scan(&duplicatePasswords)
	if err != nil {
		return nil, fmt.Errorf("failed to get duplicate passwords: %w", err)
	}
	stats["duplicate_passwords"] = duplicatePasswords

	// Most recent entry
	var lastCreated time.Time
	err = db.conn.QueryRow("SELECT MAX(created_at) FROM passwords").Scan(&lastCreated)
	if err != nil && err != sql.ErrNoRows {
		return nil, fmt.Errorf("failed to get last created: %w", err)
	}
	if err != sql.ErrNoRows {
		stats["last_created"] = lastCreated
	}

	// Most common websites
	rows, err := db.conn.Query(`
		SELECT website, COUNT(*) as count 
		FROM passwords 
		WHERE website != '' 
		GROUP BY website 
		ORDER BY count DESC 
		LIMIT 5
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to get common websites: %w", err)
	}
	defer rows.Close()

	var commonWebsites []map[string]interface{}
	for rows.Next() {
		var website string
		var count int
		if err := rows.Scan(&website, &count); err != nil {
			return nil, fmt.Errorf("failed to scan website stats: %w", err)
		}
		commonWebsites = append(commonWebsites, map[string]interface{}{
			"website": website,
			"count":   count,
		})
	}
	stats["common_websites"] = commonWebsites

	return stats, nil
}

func (db *Database) GetDatabaseSize() (int64, error) {
	fileInfo, err := os.Stat(db.path)
	if err != nil {
		return 0, fmt.Errorf("failed to get database size: %w", err)
	}
	return fileInfo.Size(), nil
}