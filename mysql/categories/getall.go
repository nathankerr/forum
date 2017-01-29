package categories

import (
	"database/sql"

	"github.com/dhenkes/forum"
	"github.com/dhenkes/forum/mysql"
)

// Selects all categories from the database and returns them.
func Getall(mysql *mysql.MySQL) ([]forum.Category, error) {
	var u forum.Category
	var result []forum.Category

	var rows *sql.Rows
	rows, mysql.Err = mysql.DB.Query("SELECT * FROM categories")
	if mysql.Err != nil {
		return nil, mysql.Err
	}
	defer rows.Close()

	for rows.Next() {
		if mysql.Err = rows.Scan(&u.ID, &u.Name, &u.Position); mysql.Err != nil {
			return nil, mysql.Err
		}
		result = append(result, u)
	}

	return result, nil
}
