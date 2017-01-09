package posts

import (
	"database/sql"
	"time"

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

	// Save post in database.
	_, err = db.Exec("INSERT INTO posts(is_sticky, user, title, thread, created) VALUES (?, ?, ?, ?, ?)", 0, 1, "First post", 1, int32(time.Now().Unix()))
	if err != nil {
		return nil, err
	}

	// Return post.
	return []byte("Post created."), nil
}
