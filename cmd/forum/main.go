package main

import (
	"fmt"
	"os"

	"github.com/dhenkes/forum/http"
	"github.com/dhenkes/forum/mysql"
)

// Checks environment variables and continues if they are set.
// If they are not set, the execution stops.
func main() {
	enVars := [4]string{"port", "user", "pass", "db"}
	notSet := 0

	for _, enVar := range enVars {
		if os.Getenv(enVar) == "" {
			fmt.Println(enVar, "is not set.")
			notSet++
		}
	}

	if notSet > 0 {
		os.Exit(1)
	}

	port := os.Getenv("port")
	username := os.Getenv("user")
	password := os.Getenv("pass")
	database := os.Getenv("db")

	m := mysql.OpenConnection(username, password, database)
	m.PingDatabase()
	if m.Err != nil {
		os.Exit(2)
	}
	m.CheckTables()
	if m.Err != nil {
		os.Exit(2)
	}

	s := http.CreateServer(port)
	s.UseMySQL(m)
	s.Router.GET("/", http.Index)
	s.Router.GET("/users", s.UsersGetAll)
	s.Router.GET("/users/:id", s.UsersGetByID)
	s.Router.GET("/boards", s.BoardsGetAll)
	s.Router.GET("/boards/:id", s.BoardsGetByID)
	s.Router.GET("/categories", s.CategoriesGetAll)
	s.Router.GET("/categories/:id", s.CategoriesGetByID)
	s.Router.GET("/threads", s.ThreadsGetAll)
	s.Router.GET("/threads/:id", s.ThreadsGetByID)
	s.Router.GET("/posts", s.PostsGetAll)
	s.Router.GET("/posts/:id", s.PostsGetByID)
	s.Run()
}
