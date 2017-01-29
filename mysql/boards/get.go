package boards

import (
	"github.com/dhenkes/forum"
	"github.com/dhenkes/forum/mysql"
)

// Selects the board with the given ID from the database and returns it.
func Get(mysql *mysql.MySQL, id int) (*forum.Board, error) {
	var u forum.Board

	row := mysql.DB.QueryRow("SELECT * FROM boards WHERE id = ?", id)
	mysql.Err = row.Scan(&u.ID, &u.Name, &u.Category, &u.Position)
	if mysql.Err != nil {
		return nil, mysql.Err
	}

	return &u, nil
}
