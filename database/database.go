package database

import (
	"database/sql"

	"gotodo/database/todos"
)

// Database defines our database struct.
type Database struct {
	Todos *todos.Database
}

// New returns as new gotodo database.
// able to pass down the database in todos folder
func New(db *sql.DB) *Database {
	return &Database{
		Todos: todos.New(db),
	}
}
