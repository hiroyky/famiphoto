package mysql

import (
	"database/sql"
	"fmt"
	"github.com/hiroyky/famiphoto/config"
)

var db *sql.DB = nil

func NewDatabaseDriver() (*sql.DB, error) {
	if db != nil {
		return db, nil
	}

	source := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.Env.MySQLUserName,
		config.Env.MySQLPassword,
		config.Env.MySQLHostName,
		config.Env.MySQLPort,
		config.Env.MySQLDatabase,
	)
	newDB, err := sql.Open("mysql", source)
	if err != nil {
		return nil, err
	}
	db = newDB
	return db, nil
}
