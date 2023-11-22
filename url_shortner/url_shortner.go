package url_shortener

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"github.com/mattn/go-sqlite3"
	"github.com/google/uuid"
)
var database *sql.DB

func initializeDB(){
	var err error
	database, err = sql.Open("sqlite3", "url_shortener.db")
	if err != nil {
		log.Fatal(err)
	}
}

func initializeDatabase() {
	_, err := database.Exec(`
        CREATE TABLE IF NOT EXISTS urls (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            original_url TEXT NOT NULL,
            short_url TEXT NOT NULL
        );
    `)
	if err != nil {
		log.Fatal(err)
	}
}

func generateShortURL() string {
	uuid := uuid.New()
	return strings.Replace(uuid.String(), "-", "", -1)[:8]
}

func shortenURL(originalURL string) {
	shortURL := generateShortURL()

	// Check if the short URL already exists
	row := database.QueryRow("SELECT * FROM urls WHERE short_url = ?", shortURL)
	var id int
	err := row.Scan(&id, nil, nil)
	if err == nil {
		fmt.Printf("Short URL already exists: %s\n", shortURL)
		return
	}

	// Insert the new URL into the database
	_, err = database.Exec("INSERT INTO urls (original_url, short_url) VALUES (?, ?)", originalURL, shortURL)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Short URL created: %s\n", shortURL)
}

func retrieveOriginalURL(shortURL string) {
	row := database.QueryRow("SELECT original_url FROM urls WHERE short_url = ?", shortURL)
	var originalURL string
	err := row.Scan(&originalURL)
	if err == sql.ErrNoRows {
		fmt.Println("Short URL not found.")
	} else if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Original URL: %s\n", originalURL)
	}
}