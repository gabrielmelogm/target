package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func ConnDB(config Config) *sql.DB {
	var db *sql.DB
	driver := config.Driver
	dsn := config.Driver + "://" + config.User + ":" + config.Password + "@" +
		config.Host + ":" + config.Port + "/" + config.Database + config.SSLMode
	db, err := sql.Open(driver, dsn)
	if err != nil {
		errConnection(config.Environment, err)
	}

	return db
}

func errConnection(environment string, err error) {
	panic("failed to connect " + environment + " postgres database_infra:" + err.Error())
}
