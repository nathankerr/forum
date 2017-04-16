package postgres

import "github.com/dhenkes/forum"

func Get(dest interface{}, query string, args ...interface{}) error {
	err = db.Get(dest, query, args...)
	return err
}

func Insert(query string, args ...interface{}) error {
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(args...)
	return err
}

func SelectUsers(dest *[]forum.User, query string, args ...interface{}) error {
	err = db.Select(dest, query, args...)
	return err
}
