package threads

import (
	"github.com/dhenkes/forum"
	"github.com/dhenkes/forum/mysql"
)

// Inserts a thread into the database.
func Create(mysql *mysql.MySQL, thread forum.Thread) ([]byte, error) {
	_, mysql.Err = mysql.DB.Exec("INSERT INTO threads(user, title, board) VALUES (?, ?, ?)", thread.User, thread.Title, thread.Board)
	if mysql.Err != nil {
		return nil, mysql.Err
	}

	return []byte("Thread created."), nil
}
