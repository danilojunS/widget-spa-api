package database

import (
	"database/sql"
	"fmt"
)

const (
	user = ""
	pass = ""
	db   = "postgres"
)

// Connect to postgres database
func Connect() (*sql.DB, error) {
	dbinfo := fmt.Sprintf("dbname=%s sslmode=disable", db)
	return sql.Open("postgres", dbinfo)
}
