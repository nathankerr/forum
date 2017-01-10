package boards

import (
	"database/sql"

	"github.com/dhenkes/forum"
	"github.com/dhenkes/forum/mysql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func Get(id int) (*forum.Board, error) {
	var u forum.Board

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

	// Select board.
	row := db.QueryRow("SELECT * FROM boards WHERE id = ?", id)
	err = row.Scan(&u.ID, &u.Name, &u.Category, &u.Position)
	if err != nil {
		return nil, err
	}

	// Return board.
	return &u, nil
}
