package threads

import (
	"github.com/dhenkes/forum"
	"github.com/dhenkes/forum/mysql"
)

// Selects the thread with the given ID from the database and returns it.
func Get(mysql *mysql.MySQL, id int) (*forum.Thread, error) {
	var u forum.Thread

	row := mysql.DB.QueryRow("SELECT * FROM threads WHERE id = ?", id)
	mysql.Err = row.Scan(&u.ID, &u.User, &u.Title, &u.Board, &u.Removed)
	if mysql.Err != nil {
		return nil, mysql.Err
	}

	return &u, nil
}
