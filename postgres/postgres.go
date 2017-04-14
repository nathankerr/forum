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

func QueryRow(dest interface{}, query string, args ...interface{}) error {
	err := db.QueryRowx(query, args...).StructScan(dest)
	return err
}

func Query(dest []interface{}, class interface{}, query string, args ...interface{}) error {
	rows, err := db.Queryx(query, args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.StructScan(&class)
		if err != nil {
			return err
		}
		dest = append(dest, class)
	}
	return nil
}
