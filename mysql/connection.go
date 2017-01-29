package mysql

import (
	"database/sql"
	// Import the MySQL driver here so all MySQL related data is contained in
	// one package.
	_ "github.com/go-sql-driver/mysql"
)

// MySQL contains the DB connection and an error.
type MySQL struct {
	DB  *sql.DB
	Err error
}

// OpenConnection opens a MySQL database connection and returns a pointer to it.
func OpenConnection(username string, password string, database string) *MySQL {
	db, err := sql.Open("mysql", username+":"+password+"@/"+database)
	return &MySQL{
		DB:  db,
		Err: err,
	}
}

// PingDatabase pings the databse.
func (mysql *MySQL) PingDatabase() {
	mysql.Err = mysql.DB.Ping()
}

// CloseConnection closes the MySQL database connection.
func (mysql *MySQL) CloseConnection() {
	mysql.DB.Close()
}
