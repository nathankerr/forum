package threads

import (
	"database/sql"

	"github.com/dhenkes/forum"
	"github.com/dhenkes/forum/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func Create(thread forum.Thread) ([]byte, error) {
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

	// Save thread in database.
	_, err = db.Exec("INSERT INTO threads(user, title, board) VALUES (?, ?, ?)", thread.User, thread.Title, thread.Board)
	if err != nil {
		return nil, err
	}

	// Return thread.
	return []byte("Thread created."), nil
}
