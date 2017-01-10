package categories

import (
	"database/sql"

	"github.com/dhenkes/forum"
	"github.com/dhenkes/forum/mysql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func Get(id int) (*forum.Category, error) {
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

	// Select category.
	row := db.QueryRow("SELECT * FROM categories WHERE id = ?", id)
	err = row.Scan(&u.ID, &u.Name, &u.Position)
	if err != nil {
		return nil, err
	}

	// Return category.
	return &u, nil
}
