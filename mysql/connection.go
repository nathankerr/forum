package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func OpenConnection() error {
	db, err = sql.Open("mysql", Username+":"+Password+"@/"+Database)
	if err != nil {
		return err
	}
	return nil
}

func PingDatabase() error {
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func CloseConnection() {
	db.Close()
}
