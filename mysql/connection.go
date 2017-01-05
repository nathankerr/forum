package mysql

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func Connect() {
	log.Println(" ")
	log.Println("Connecting to MySQL")
	db, err = sql.Open("mysql", username+":"+password+"@/"+database)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected successfully")

	// hashedPassword, err := bcrypt.GenerateFromPassword([]byte("123"), bcrypt.DefaultCost)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// _, err = db.Exec("INSERT INTO users(username, password, removed) VALUES (?, ?, ?)", "dhenkes", hashedPassword, 0)
}
