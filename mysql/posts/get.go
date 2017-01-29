package posts

import (
	"github.com/dhenkes/forum"
	"github.com/dhenkes/forum/mysql"
)

// Selects the post with the given ID from the database and returns it.
func Get(mysql *mysql.MySQL, id int) (*forum.Post, error) {
	var u forum.Post

	row := mysql.DB.QueryRow("SELECT * FROM posts WHERE id = ?", id)
	mysql.Err = row.Scan(&u.ID, &u.IsSticky, &u.User, &u.Title, &u.Thread, &u.Content, &u.Created, &u.Modified, &u.Removed)
	if mysql.Err != nil {
		return nil, mysql.Err
	}

	return &u, nil
}
