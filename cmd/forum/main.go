package main

import (
	"github.com/dhenkes/forum/http"
	"github.com/dhenkes/forum/mysql"
)

func main() {
	mysql.CheckTables()

	s := http.CreateServer(":8080")
	s.Router.GET("/", http.Index)
	s.Router.GET("/users", http.UsersGetAll)
	s.Router.GET("/users/:id", http.UsersGetByID)
	s.Router.GET("/boards", http.BoardsGetAll)
	s.Router.GET("/boards/:id", http.BoardsGetByID)
	s.Router.GET("/categories", http.CategoriesGetAll)
	s.Router.GET("/categories/:id", http.CategoriesGetByID)
	s.Router.GET("/threads", http.ThreadsGetAll)
	s.Router.GET("/threads/:id", http.ThreadsGetByID)
	s.Router.GET("/posts", http.PostsGetAll)
	s.Router.GET("/posts/:id", http.PostsGetByID)
	s.Run()
}
