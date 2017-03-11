package postgres

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

func Connect(host string, user string, pass string, name string) error {
	db, err = sql.Open("postgres", "postgres://"+user+":"+pass+"@"+host+"/"+name+"?sslmode=disable")
	return err
}

func Ping() error {
	err = db.Ping()
	return err
}
