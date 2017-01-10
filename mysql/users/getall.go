package users

import (
	"database/sql"

	"github.com/dhenkes/forum"
	"github.com/dhenkes/forum/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func Getall() ([]forum.User, error) {
	var u forum.User
	var result []forum.User

	// Connect to database.
	db, err = sql.Open("mysql", mysql.Username+":"+mysql.Password+"@/"+mysql.Database)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Test connection.
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// Select users.
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Testing
	for rows.Next() {
		if err := rows.Scan(&u.ID, &u.Username, &u.Password, &u.Role, &u.Removed); err != nil {
			return nil, err
		}

		// Add row scan to result array
		result = append(result, u)
	}

	// Return user.
	return result, nil
}
