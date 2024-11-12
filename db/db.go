package db

import (
	"database/sql"
	"fmt"

	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	fmt.Println("Init DB executed")
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		log.Fatal(err)
		panic("Could not connect to database.")
	}

	// Test the connection
	err = DB.Ping()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	log.Println("Successfully connected to the SQLite database")
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUsertable := `CREATE TABLE IF NOT EXISTS users(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	username TEXT NOT NULL UNIQUE,
	email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`
	_, err := DB.Exec(createUsertable)
	if err != nil {
		panic("Could not Create user The Table.")
	}
	createEventTable := `CREATE TABLE IF NOT EXISTS events(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	description TEXT NOT NULL,
	location TEXT NOT NULL,
	datetime DATETIME NOT NULL,
	user_id INTEGER,
	FOREIGN KEY(user_id) REFERENCES users(id)
	)`

	_, err = DB.Exec(createEventTable)
	if err != nil {
		panic("Could not Create Event The Table.")
	}

	log.Println("Successfully Created The Tables")

}
