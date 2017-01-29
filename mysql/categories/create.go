package categories

import (
	"github.com/dhenkes/forum"
	"github.com/dhenkes/forum/mysql"
)

// Inserts a category into the database.
func Create(mysql *mysql.MySQL, category forum.Category) ([]byte, error) {
	_, mysql.Err = mysql.DB.Exec("INSERT INTO categories(name, position) VALUES (?, ?)", category.Name, category.Position)
	if mysql.Err != nil {
		return nil, mysql.Err
	}

	return []byte("Category created."), nil
}
