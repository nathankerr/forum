package users

import (
	"github.com/dhenkes/forum"
	"github.com/dhenkes/forum/mysql"
)

// Selects the user with the given ID from the database and returns it.
func Get(mysql *mysql.MySQL, id int) (*forum.User, error) {
	var u forum.User

	row := mysql.DB.QueryRow("SELECT * FROM users WHERE id = ?", id)
	mysql.Err = row.Scan(&u.ID, &u.Username, &u.Password, &u.Role, &u.Removed)
	if mysql.Err != nil {
		return nil, mysql.Err
	}

	return &u, nil
}
