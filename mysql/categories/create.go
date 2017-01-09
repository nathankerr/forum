package categories

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

	// Save category in database.
	_, err = db.Exec("INSERT INTO categories(name, position) VALUES (?, ?)", "First category", 1)
	if err != nil {
		return nil, err
	}

	// Return category.
	return []byte("Category created."), nil
}