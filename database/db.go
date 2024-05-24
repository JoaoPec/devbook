package database

import (
	"database/sql"
    "fmt"
)

func StartDB() (*sql.DB, error) {

	db, err := sql.Open("sqlite3", "devbook.db")
	if err != nil {
		return nil, err
	}

	defer db.Close()

	query, err := db.Prepare(`CREATE TABLE IF NOT EXISTS users (
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        name TEXT,
        email TEXT,
        password TEXT);`)
	if err != nil {
		return nil, err
	}

	_, err = query.Exec()
	if err != nil {
		return nil, err
	}

    fmt.Println("Database started successfully");

	return db, nil
}

