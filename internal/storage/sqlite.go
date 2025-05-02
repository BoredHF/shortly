package storage

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() error {
	db, err := sql.Open("sqlite", "./shortly.db")
	if err != nil {
		return err
	}

	createTable := `
        CREATE TABLE IF NOT EXISTS urls (
            id TEXT PRIMARY KEY,
            original_url TEXT NOT NULL
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

func SaveURL(id string, url string) error {
	_, err := DB.Exec("INSERT INTO urls (id, original_url) VALUES (?, ?)", id, url)
	return err
}

func GetOriginalURL(id string) (string, error) {
	var originalURL string
	row := DB.QueryRow("SELECT original_url FROM urls WHERE id = ?", id)
	err := row.Scan(&originalURL)
	if err != nil {
		return "", err // return empty + error if not found or failed
	}
	return originalURL, nil
}
