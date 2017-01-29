package categories

import (
	"github.com/dhenkes/forum"
	"github.com/dhenkes/forum/mysql"
)

// Get selects the category with the given ID from the database and returns it.
func Get(mysql *mysql.MySQL, id int) (*forum.Category, error) {
	var u forum.Category

	row := mysql.DB.QueryRow("SELECT * FROM categories WHERE id = ?", id)
	mysql.Err = row.Scan(&u.ID, &u.Name, &u.Position)
	if mysql.Err != nil {
		return nil, mysql.Err
	}

	return &u, nil
}
