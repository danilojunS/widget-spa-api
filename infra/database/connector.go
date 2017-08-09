package database

import (
	"database/sql"
	"fmt"
	config "github.com/danilojunS/widgets-spa-api/config"
)

// Connect to postgres database
func Connect() (*sql.DB, error) {
	var (
		db = config.Get().DB
		// user = config.Get().DBUser
		// pass = config.Get().DBPass
	)
	dbinfo := fmt.Sprintf("dbname=%s sslmode=disable", db)
	return sql.Open("postgres", dbinfo)
}
