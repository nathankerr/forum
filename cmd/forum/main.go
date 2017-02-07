package main

import (
	"os"

	"github.com/dhenkes/forum/couchbase"
	"github.com/dhenkes/forum/handler/users"
	"github.com/dhenkes/forum/http"
)

func main() {
	enVars := [4]string{"http_port", "cb_url", "cb_bucket", "cb_pass"}
	notSet := 0

	for _, enVar := range enVars {
		if os.Getenv(enVar) == "" {
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

	couchbase.Connect(&cb_url)
	if couchbase.DB.Err != nil {
		os.Exit(2)
	}

	couchbase.OpenBucket(&cb_bucket, &cb_pass)
	if couchbase.DB.Err != nil {
		os.Exit(3)
	}

	http.CreateServer(&http_port)
	http.Server.Router.GET("/users/:id", users.Get)
	http.Server.Router.GET("/users", users.GetAll)
	http.Server.Router.POST("/users", users.Create)
	http.Run()
}
