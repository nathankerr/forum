package users

import (
	"database/sql"

	"github.com/dhenkes/forum"
	_ "github.com/go-sql-driver/mysql"
)

type UserService struct {
	DB *sql.DB
}

func (s *UserService) User(id int) (*forum.User, error) {
	var u forum.User
	row := db.QueryRow(`SELECT id, name FROM users WHERE id = $1`, id)
	if row.Scan(&u.ID, &u.Name); err != nil {
		return nil, err
	}
	return &u, nil
}
