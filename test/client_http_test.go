package test

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func setupDb() *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStatement := `
		CREATE TABLE IF NOT EXISTS clients 
		(
			id TEXT NOT NULL PRIMARY KEY,
			name TEXT,
			email TEXT,
			points INTEGER
		);
	`

	_, err = db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
	return db
}

func tearDownDb(db *sql.DB) {
	db.Exec("DROP TABLE clients")
	db.Close()
}
