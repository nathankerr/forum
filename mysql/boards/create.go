package boards

import (
	"database/sql"

	"github.com/dhenkes/forum/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func Create() ([]byte, error) {
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

	// Save board in database.
	_, err = db.Exec("INSERT INTO boards(name, cateogry, position) VALUES (?, ?, ?)", "First board", 1, 1)
	if err != nil {
		return nil, err
	}

	// Return board.
	return []byte("Board created."), nil
}
