package users

import (
	"database/sql"

	"github.com/dhenkes/forum"
	"github.com/dhenkes/forum/mysql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func Get(id int) (*forum.User, error) {

	var u forum.User

	// Connect to database.
	db, err = sql.Open("mysql", mysql.Username+":"+mysql.Password+"@/"+mysql.Database)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Test connection.
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// Select user.
	row := db.QueryRow("SELECT * FROM users WHERE id = ?", id)
	err = row.Scan(&u.ID, &u.Username, &u.Password, &u.Role, &u.Removed)
	if err != nil {
		return nil, err
	}

	// Return user.
	return &u, nil
}
