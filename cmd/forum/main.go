package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/dhenkes/forum/handler/users"
	"github.com/dhenkes/forum/http"
	"github.com/dhenkes/forum/logging"
	"github.com/dhenkes/forum/postgres"
)

func main() {

	m := make(map[string]string)
	getInput(m)

	err := postgres.Connect(m["db_host"], m["db_user"], m["db_pass"], m["db_name"])
	if err != nil {
		logging.PrintError("Could not open connection to PostgreSQL.")
		return
	}

	err = postgres.Ping()
	if err != nil {
		logging.PrintError("Could not ping PostgreSQL.")
		logging.PrintError(err.Error())
		return
	}

	http.CreateServer(m["http_port"])
	http.Server.Router.GET("/users/:id", users.Get)
	http.Server.Router.GET("/users", users.GetAll)
	http.Run()
}

func getInput(m map[string]string) {
	scanner := bufio.NewScanner(os.Stdin)
	for _, v := range [5]string{"http_port", "db_host", "db_user", "db_pass", "db_name"} {
		fmt.Print(v, ": ")
		scanner.Scan()
		m[v] = scanner.Text()
	}
}
