package boards

import (
	"database/sql"

	"github.com/dhenkes/forum"
	"github.com/dhenkes/forum/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func Getall() (*forum.Board, error) {

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

	// Select boards.
	rows, err := db.Query("SELECT * FROM boards")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Testing
	for rows.Next() {
		if err := rows.Scan(&u.ID, &u.Name, &u.Category, &u.Position); err != nil {
			return nil, err
		}
	}

	// Return boards.
	return &u, nil
}
