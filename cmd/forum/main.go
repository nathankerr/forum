package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/dhenkes/forum/handler/users"
	"github.com/dhenkes/forum/http"
	"github.com/dhenkes/forum/logger"
	"github.com/dhenkes/forum/postgres"
)

var configVars = [5]string{"http_port", "db_host", "db_user", "db_pass", "db_name"}
var notSet []string

func checkEnv(config map[string]string) {
	for _, v := range configVars {
		config[v] = os.Getenv(v)
		if len(config[v]) == 0 {
			notSet = append(notSet, v)
		}
	}
}

func getInput(config map[string]string) {
	scanner := bufio.NewScanner(os.Stdin)
	for _, v := range notSet {
		logAPIMsg(v + ": ")
		scanner.Scan()
		config[v] = scanner.Text()
	}
}

func logAPIMsg(message string) {
	colour := "\x1b[35m"
	reset := "\x1b[0m"
	now := time.Now().Format("2006-01-02 15:04:05")

	fmt.Printf("%s%s - API%s: %s", colour, now, reset, message)
}

func main() {
	config := make(map[string]string)

	checkEnv(config)
	if len(notSet) > 0 {
		logAPIMsg("Not all environment variables set.\n")
		logAPIMsg("Please enter the value(s) for the following variable(s).\n")
		getInput(config)
	}

	logger.Info("%s", "Starting API")

	err := postgres.Connect(config["db_host"], config["db_user"], config["db_pass"], config["db_name"])
	if err != nil {
		logger.Error("%s", "Could not open connection to PostgreSQL.")
		return
	}
	logger.Info("%s", "Connected to PostgreSQL")

	err = postgres.Ping()
	if err != nil {
		logger.Error("%s", "Could not ping PostgreSQL.")
		logger.Error("%s", err.Error())
		return
	}
	logger.Info("%s", "Pinged PostgreSQL")

	http.CreateServer(config["http_port"])
	logger.Info("%s %s", "Starting server with port", config["http_port"])

	http.Server.Router.GET("/users/:id", users.Get)
	http.Server.Router.GET("/users", users.GetAll)
	http.Server.Router.POST("/users", users.Post)

	logger.Info("%s", "Registered route [GET] /users/:id")
	logger.Info("%s", "Registered route [GET] /users")
	logger.Info("%s", "Registered route [POST] /users")

	http.Run()
}
