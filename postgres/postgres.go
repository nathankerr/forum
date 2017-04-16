package postgres

import (
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

var db *sqlx.DB
var err error

func Connect(host string, user string, pass string, name string) error {
	db, err = sqlx.Connect("postgres", "postgres://"+user+":"+pass+"@"+host+"/"+name+"?sslmode=disable")
	return err
}

func Ping() error {
	err = db.Ping()
	return err
}
