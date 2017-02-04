package main

import (
	"fmt"
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

	http.CreateServer(http_port, cb)
	http.Server.Router.GET("/users/:id", users.Get)
	http.Run()
}
