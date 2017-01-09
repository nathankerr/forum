package categories

import (
	"database/sql"

	"github.com/dhenkes/forum"
	"github.com/dhenkes/forum/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func Getall() (*forum.Category, error) {

	var u forum.Category

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

	// Select categories.
	rows, err := db.Query("SELECT * FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Testing
	for rows.Next() {
		if err := rows.Scan(&u.ID, &u.Name, &u.Position); err != nil {
			return nil, err
		}
	}

	// Return categories.
	return &u, nil
}
