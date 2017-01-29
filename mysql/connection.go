package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type MySQL struct {
	DB  *sql.DB
	Err error
}

// Opens a MySQL database connection and returns a pointer to it.
func OpenConnection(username string, password string, database string) *MySQL {
	db, err := sql.Open("mysql", username+":"+password+"@/"+database)
	return &MySQL{
		DB:  db,
		Err: err,
	}
}

// Pings the databse.
func (mysql *MySQL) PingDatabase() {
	mysql.Err = mysql.DB.Ping()
}

// Closes the MySQL database connection.
func (mysql *MySQL) CloseConnection() {
	mysql.DB.Close()
}
