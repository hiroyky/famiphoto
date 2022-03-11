package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hiroyky/famiphoto/config"
	"github.com/hiroyky/famiphoto/infrastructures/repositories"
)

var db *sql.DB = nil

func NewDatabaseDriver() repositories.SQLExecutor {
	if db != nil {
		return db
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
		panic(err)
	}
	db = newDB
	return db
}
