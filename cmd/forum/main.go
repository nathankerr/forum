package main

import (
	"bufio"
	"os"

	"github.com/dhenkes/forum/handler/users"
	"github.com/dhenkes/forum/http"
	"github.com/dhenkes/forum/logger"
	"github.com/dhenkes/forum/postgres"
)

func loadConfig() map[string]string {
	config := map[string]string{}

	var notSet []string
	for _, v := range []string{"http_port", "db_host", "db_user", "db_pass", "db_name"} {
		config[v] = os.Getenv(v)
		if len(config[v]) == 0 {
			notSet = append(notSet, v)
		}
	}

	if len(notSet) > 0 {
		logger.API("Not all environment variables set.\n")
		logger.API("Please enter the value(s) for the following variable(s).\n")

		scanner := bufio.NewScanner(os.Stdin)
		for _, v := range notSet {
			logger.API(v + ": ")
			scanner.Scan()
			config[v] = scanner.Text()
		}
	}

	return config
}

func main() {
	config := loadConfig()

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
