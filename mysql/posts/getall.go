package posts

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

	// Select posts.
	rows, err := db.Query("SELECT * FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Testing
	for rows.Next() {
		if err := rows.Scan(&u.ID, &u.IsSticky, &u.User, &u.Title, &u.Thread, &u.Created, &u.Modified, &u.Removed); err != nil {
			return nil, err
		}
	}

	// Return posts.
	return &u, nil
}
