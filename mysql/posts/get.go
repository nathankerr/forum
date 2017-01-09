package posts

import (
	"database/sql"

	"github.com/dhenkes/forum"
	"github.com/dhenkes/forum/mysql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func Get(id int) (*forum.Post, error) {

	var u forum.Post

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

	// Select post.
	row := db.QueryRow("SELECT * FROM posts WHERE id = ?", id)
	err = row.Scan(&u.ID, &u.IsSticky, &u.User, &u.Title, &u.Thread, &u.Created, &u.Modified, &u.Removed)
	if err != nil {
		return nil, err
	}

	// Return post.
	return &u, nil
}