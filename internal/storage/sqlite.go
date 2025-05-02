package storage

import (
	"database/sql"
	"errors"
	"time"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

var ErrLinkExpired = errors.New("link has expired")

func InitDB() error {
	db, err := sql.Open("sqlite", "./shortly.db")
	if err != nil {
		return err
	}

	createTable := `
		CREATE TABLE IF NOT EXISTS urls (
			id TEXT PRIMARY KEY,
			original_url TEXT NOT NULL,
			expires_at DATETIME DEFAULT (DATETIME('now', '+30 days'))
		);
    `
	_, err = db.Exec(createTable)
	if err != nil {
		return err
	}

	DB = db
	return nil
}

func FindShortIDByOriginalURL(original string) (string, error) {
	var id string
	row := DB.QueryRow("SELECT id FROM urls WHERE original_url = ?", original)
	err := row.Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}

func SaveURL(id string, url string, expiresAt *string) error {
	if expiresAt == nil {
		_, err := DB.Exec("INSERT INTO urls (id, original_url) VALUES (?, ?)", id, url)
		return err
	}
	_, err := DB.Exec("INSERT INTO urls (id, original_url, expires_at) VALUES (?, ?, ?)", id, url, expiresAt)
	return err
}

func GetOriginalURL(id string) (string, error) {
	var originalURL string
	var expiresAt sql.NullString

	row := DB.QueryRow(`
		SELECT original_url, expires_at FROM urls 
		WHERE id = ?
	`, id)

	err := row.Scan(&originalURL, &expiresAt)
	if err != nil {
		return "", err // id not found
	}

	// Check if expired
	if expiresAt.Valid {
		t, _ := time.Parse("2006-01-02 15:04:05", expiresAt.String)
		if t.Before(time.Now()) {
			return "", ErrLinkExpired
		}
	}

	return originalURL, nil
}

// Testing
func CleanupExpiredURLs() (int64, error) {
	result, err := DB.Exec(`
		DELETE FROM urls 
		WHERE expires_at IS NOT NULL 
		AND datetime(expires_at) <= datetime('now', 'localtime')
	`)
	if err != nil {
		return 0, err
	}
	rowsDeleted, _ := result.RowsAffected()
	return rowsDeleted, nil
}
