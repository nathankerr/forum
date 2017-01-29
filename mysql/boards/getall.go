package boards

import (
	"database/sql"

	"github.com/dhenkes/forum"
	"github.com/dhenkes/forum/mysql"
)

// Getall selects all boards from the database and returns them.
func Getall(mysql *mysql.MySQL) ([]forum.Board, error) {
	var u forum.Board
	var result []forum.Board

	var rows *sql.Rows
	rows, mysql.Err = mysql.DB.Query("SELECT * FROM boards")
	if mysql.Err != nil {
		return nil, mysql.Err
	}
	defer rows.Close()

	for rows.Next() {
		if mysql.Err = rows.Scan(&u.ID, &u.Name, &u.Category, &u.Position); mysql.Err != nil {
			return nil, mysql.Err
		}
		result = append(result, u)
	}

	return result, nil
}
