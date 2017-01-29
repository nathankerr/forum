package posts

import (
	"database/sql"

	"github.com/dhenkes/forum"
	"github.com/dhenkes/forum/mysql"
)

// Selects all posts from the database and returns them.
func Getall(mysql *mysql.MySQL) ([]forum.Post, error) {
	var u forum.Post
	var result []forum.Post

	var rows *sql.Rows
	rows, mysql.Err = mysql.DB.Query("SELECT * FROM posts")
	if mysql.Err != nil {
		return nil, mysql.Err
	}
	defer rows.Close()

	for rows.Next() {
		if mysql.Err = rows.Scan(&u.ID, &u.IsSticky, &u.User, &u.Title, &u.Thread, &u.Content, &u.Created, &u.Modified, &u.Removed); mysql.Err != nil {
			return nil, mysql.Err
		}
		result = append(result, u)
	}

	return result, nil
}
