package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func New(dbPath string) (*sql.DB, error) {
	return sql.Open("sqlite3", dbPath)
}
