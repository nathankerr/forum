package users

import (
	"database/sql"

	"github.com/dhenkes/forum/mysql"
	"github.com/dhenkes/forum/utils"
	_ "github.com/go-sql-driver/mysql"
)

func Create() ([]byte, error) {
	// Hash the password.
	hash, err := utils.HashPassword("test")
	if err != nil {
		return nil, err
	}

	// Convert []byte to string.
	hashString := string(hash[:])

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

	// Save user in database.
	_, err = db.Exec("INSERT INTO users(username, password, role) VALUES (?, ?, ?)", "dhenkes", hashString, 2)
	if err != nil {
		return nil, err
	}

	// Return user.
	return []byte("User created."), nil
}
