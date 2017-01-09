package threads

import (
	"database/sql"

	"github.com/dhenkes/forum"
	"github.com/dhenkes/forum/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func Getall() (*forum.Thread, error) {

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

	// Select threads.
	rows, err := db.Query("SELECT * FROM threads")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Testing
	for rows.Next() {
		if err := rows.Scan(&u.ID, &u.User, &u.Title, &u.Board, &u.Removed); err != nil {
			return nil, err
		}
	}

	// Return threads.
	return &u, nil
}
