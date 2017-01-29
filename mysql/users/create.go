package users

import (
	"github.com/dhenkes/forum"
	"github.com/dhenkes/forum/mysql"
	"github.com/dhenkes/forum/utils"
)

// Inserts a user with an encrypted password into the database.
func Create(mysql *mysql.MySQL, user forum.User) ([]byte, error) {
	var hash []byte
	hash, mysql.Err = utils.HashPassword(user.Password)
	if mysql.Err != nil {
		return nil, mysql.Err
	}

	hashString := string(hash[:])

	_, mysql.Err = mysql.DB.Exec("INSERT INTO users(username, password, role) VALUES (?, ?, ?)", user.Username, hashString, user.Role)
	if mysql.Err != nil {
		return nil, mysql.Err
	}

	return []byte("User created."), nil
}
