package models

import (
	"github.com/Yadier01/notes-backend/db"
)

type Note struct {
	ID      int    `json:"id" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

var notes []Note = []Note{}

func (n Note) Save() error {
	query := `INSERT INTO notes (title, content) VALUES (?, ?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(n.Title, n.Content)
	if err != nil {
		return err
	}

	return nil
}
