package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dhenkes/forum/http"
	"github.com/dhenkes/forum/postgres"
)

func main() {
	enVars := [5]string{"http_port", "db_host", "db_user", "db_pass", "db_name"}
	notSet := 0

	for _, enVar := range enVars {
		if os.Getenv(enVar) == "" {
			notSet++
		}
	}

	if notSet > 0 {
		log.Fatal("Error: Not all environment variables are set.")
	}

	http_port := os.Getenv("http_port")
	db_host := os.Getenv("db_host")
	db_user := os.Getenv("db_user")
	db_pass := os.Getenv("db_pass")
	db_name := os.Getenv("db_name")

	err := postgres.Connect(db_host, db_user, db_pass, db_name)
	if err != nil {
		fmt.Println("Error: Could not open connection to PostgreSQL.")
		fmt.Printf("Error: ")
		fmt.Println(err)
		return
	}

	err = postgres.Ping()
	if err != nil {
		fmt.Println("Error: Could not ping PostgreSQL.")
		fmt.Printf("Error: ")
		fmt.Println(err)
		return
	}

	http.CreateServer(http_port)
	// http.Server.Router.GET("/users/:id", users.Get)
	// http.Server.Router.GET("/users", users.GetAll)
	// http.Server.Router.POST("/users", users.Create)
	http.Run()
}
