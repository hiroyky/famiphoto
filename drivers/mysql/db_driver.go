package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hiroyky/famiphoto/config"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type SQLExecutor interface {
	boil.ContextExecutor
	boil.ContextBeginner
}

var db SQLExecutor = nil

func NewDatabaseDriver() SQLExecutor {
	if db != nil {
		return db
	}

	source := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.Env.MySQLUser,
		config.Env.MySQLPassword,
		config.Env.MySQLHostName,
		config.Env.MySQLPort,
		config.Env.MySQLDatabase,
	)
	newDB, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	if err := newDB.Ping(); err != nil {
		panic(fmt.Sprintf("%s %+v", config.Env.MySQLHostName, err))
	}

	db = newDB
	return db
}
