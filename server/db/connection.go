package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func ConnDB(config Config) *sql.DB {
	var db *sql.DB
	dsn := config.DbUrl
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		errConnection(config.Environment, err)
	}

	return db
}

func errConnection(environment string, err error) {
	panic("failed to connect " + environment + " postgres database_infra:" + err.Error())
}
