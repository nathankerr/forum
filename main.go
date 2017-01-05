package main

import (
	"forum/controllers/users"
	"forum/http"
	// "forum/mysql"
	"log"
)

func main() {
	log.Println("Forum API v0.1")

	// mysql.Connect()

	server := http.CreateServer(":8080")
	server.Router.GET("/users", users.Getall)
	server.Router.GET("/users/:id", users.Get)
	server.Run()
}
