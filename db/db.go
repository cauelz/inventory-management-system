package db

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)


type Database interface {
	Select(dest interface{}, query string, args ...interface{}) error
	Get(dest interface{}, query string, args ...interface{}) error
	Exec(query string, args ...interface{}) (sql.Result, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

type DB struct {
	*sqlx.DB
}