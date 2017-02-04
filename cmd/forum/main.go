package main

import (
	"fmt"
	"os"

	"github.com/dhenkes/forum/couchbase"
	"github.com/dhenkes/forum/http"
)

func main() {
	enVars := [4]string{"http_port", "cb_url", "cb_bucket", "cb_pass"}
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

	http_port := os.Getenv("http_port")
	cb_url := os.Getenv("cb_url")
	cb_bucket := os.Getenv("cb_bucket")
	cb_pass := os.Getenv("cb_pass")

	cb := couchbase.Connect(cb_url)
	if cb.Err != nil {
		fmt.Println(cb.Err)
		os.Exit(2)
	}

	cb.OpenBucket(cb_bucket, cb_pass)
	if cb.Err != nil {
		fmt.Println(cb.Err)
		os.Exit(3)
	}

	s := http.CreateServer(http_port, cb)
	// s.Router.GET("/", http.Index)
	// s.Router.GET("/users", s.UsersGetAll)
	s.Router.GET("/users/:id", s.UsersGetByID)
	// s.Router.GET("/boards", s.BoardsGetAll)
	// s.Router.GET("/boards/:id", s.BoardsGetByID)
	// s.Router.GET("/categories", s.CategoriesGetAll)
	// s.Router.GET("/categories/:id", s.CategoriesGetByID)
	// s.Router.GET("/threads", s.ThreadsGetAll)
	// s.Router.GET("/threads/:id", s.ThreadsGetByID)
	// s.Router.GET("/posts", s.PostsGetAll)
	// s.Router.GET("/posts/:id", s.PostsGetByID)
	s.Run()
}
