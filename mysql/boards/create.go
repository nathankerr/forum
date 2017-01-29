package boards

import (
	"github.com/dhenkes/forum"
	"github.com/dhenkes/forum/mysql"
)

// Create inserts a board into the database.
func Create(mysql *mysql.MySQL, board forum.Board) ([]byte, error) {
	_, mysql.Err = mysql.DB.Exec("INSERT INTO boards(name, cateogry, position) VALUES (?, ?, ?)", board.Name, board.Category, board.Position)
	if mysql.Err != nil {
		return nil, mysql.Err
	}

	return []byte("Board created."), nil
}
