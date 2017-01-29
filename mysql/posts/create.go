package posts

import (
	"time"

	"github.com/dhenkes/forum"
	"github.com/dhenkes/forum/mysql"
)

// Inserts a post into the database.
func Create(mysql *mysql.MySQL, post forum.Post) ([]byte, error) {
	_, mysql.Err = mysql.DB.Exec("INSERT INTO posts(is_sticky, user, title, thread, content, created) VALUES (?, ?, ?, ?, ?, ?)", post.IsSticky, post.User, post.Title, post.Thread, post.Content, int32(time.Now().Unix()))
	if mysql.Err != nil {
		return nil, mysql.Err
	}

	return []byte("Post created."), nil
}
