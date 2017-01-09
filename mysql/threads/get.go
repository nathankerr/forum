package threads

import (
	"database/sql"

	"github.com/dhenkes/forum"
	"github.com/dhenkes/forum/mysql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func Get(id int) (*forum.Thread, error) {

	var u forum.Thread

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

	// Select thread.
	row := db.QueryRow("SELECT * FROM threads WHERE id = ?", id)
	err = row.Scan(&u.ID, &u.User, &u.Title, &u.Board, &u.Removed)
	if err != nil {
		return nil, err
	}

	// Return thread.
	return &u, nil
}
