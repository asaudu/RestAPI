package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// var DB *sql.DB

type Database struct {
	DB *sql.DB
}

func NewDatabase(db *sql.DB) *Database {
	return &Database{DB: db}
}

func (db *Database) InitDB() {
	var err error
	db.DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to database!")
	}

	db.DB.SetMaxOpenConns(10)
	db.DB.SetMaxIdleConns(5)

	createTables(db.DB)
}

func createTables(db *sql.DB) {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER
	)
	`
	_, err := db.Exec(createEventsTable)
	if err != nil {
		panic("Could not create events table")
	}
}
