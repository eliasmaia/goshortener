package main

import (
	"database/sql"
	- "github.com/mattn/go-sqlite3"
)

func initDB() *sql.DB {
	db, _ := sql.Open("sqlite3", "./urls.db")
	statement, _ := db.Prepare("CREATE TABLE IF NOT EXISTS link (id INTEGER PRIMARY KEY, code TEXT, url TEXT)")
	statement.Exec()
	return db
}