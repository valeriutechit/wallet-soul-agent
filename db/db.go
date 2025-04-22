package db

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./journal.db")
	if err != nil {
		log.Fatal("❌ Failed to connect to DB:", err)
	}

	createTable := `
	CREATE TABLE IF NOT EXISTS soul_reports (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		address TEXT,
		profile TEXT,
		reflection TEXT,
		timestamp DATETIME
	);`

	if _, err := DB.Exec(createTable); err != nil {
		log.Fatal("❌ Failed to create table:", err)
	}

	log.Println("🗂️ SQLite initialized")
}

func SaveReport(address, profile, reflection string) {
	stmt, err := DB.Prepare("INSERT INTO soul_reports(address, profile, reflection, timestamp) VALUES(?, ?, ?, ?)")
	if err != nil {
		log.Println("❌ Prepare failed:", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(address, profile, reflection, time.Now())
	if err != nil {
		log.Println("❌ Insert failed:", err)
	}
}
