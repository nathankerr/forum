package threads

import (
	"database/sql"

	"github.com/dhenkes/forum"
	"github.com/dhenkes/forum/mysql"
)

// Getall selects all threads from the database and returns them.
func Getall(mysql *mysql.MySQL) ([]forum.Thread, error) {
	var u forum.Thread
	var result []forum.Thread

	var rows *sql.Rows
	rows, mysql.Err = mysql.DB.Query("SELECT * FROM threads")
	if mysql.Err != nil {
		return nil, mysql.Err
	}
	defer rows.Close()

	for rows.Next() {
		if mysql.Err = rows.Scan(&u.ID, &u.User, &u.Title, &u.Board, &u.Removed); mysql.Err != nil {
			return nil, mysql.Err
		}
		result = append(result, u)
	}

	return result, nil
}
