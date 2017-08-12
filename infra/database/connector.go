package database

import (
	"database/sql"
	"fmt"
	config "github.com/danilojunS/widgets-spa-api/config"
)

// Connect to postgres database
func Connect() (*sql.DB, error) {
	var (
		db   = config.Get().DB
		host = config.Get().DBHost
		user = config.Get().DBUser
		pass = config.Get().DBPass
	)

	hostInfo := ""
	if host != "" {
		hostInfo = fmt.Sprintf("host=%s", host)
	}

	userInfo := ""
	if user != "" {
		userInfo = fmt.Sprintf("user=%s", user)
	}

	passInfo := ""
	if pass != "" {
		passInfo = fmt.Sprintf("user=%s", pass)
	}

	dbinfo := fmt.Sprintf("dbname=%s sslmode=disable %s %s %s", db, hostInfo, userInfo, passInfo)
	return sql.Open("postgres", dbinfo)
}
