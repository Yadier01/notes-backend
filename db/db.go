package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not open db")
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	//TODO Falta anadir al notesTable el tags tiene que estar linkeado con el note table
	//posiblemente con un foreign key y hacer un tagsTable
	createNotesTable := `
	CREATE TABLE IF NOT EXISTS notes (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		content TEXT NOT NULL,
		archived BOOLEAN NOT NULL

);`

	_, err := DB.Exec(createNotesTable)
	if err != nil {
		panic(err)
	}
}
