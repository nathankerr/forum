package users

import (
	"database/sql"

	"github.com/dhenkes/forum"
	"github.com/dhenkes/forum/mysql"
)

// Selects all users from the database and returns them.
func Getall(mysql *mysql.MySQL) ([]forum.User, error) {
	var u forum.User
	var result []forum.User

	var rows *sql.Rows
	rows, mysql.Err = mysql.DB.Query("SELECT * FROM users")
	if mysql.Err != nil {
		return nil, mysql.Err
	}
	defer rows.Close()

	for rows.Next() {
		if mysql.Err = rows.Scan(&u.ID, &u.Username, &u.Password, &u.Role, &u.Removed); mysql.Err != nil {
			return nil, mysql.Err
		}
		result = append(result, u)
	}

	return result, nil
}
