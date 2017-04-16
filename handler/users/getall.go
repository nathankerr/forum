package users

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dhenkes/forum"
	"github.com/dhenkes/forum/logger"
	"github.com/dhenkes/forum/postgres"
	"github.com/julienschmidt/httprouter"
)

func GetAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	users := []forum.User{}

	err := postgres.SelectUsers(&users, "SELECT uuid, username FROM users WHERE removed = $1", "0")
	if err != nil {
		logger.Warning("%s", "Error during Request: SELECT uuid, username FROM users WHERE removed = 0")
		logger.Warning("%s", err.Error())

		w.WriteHeader(http.StatusNotFound)
	}

	js, _ := json.Marshal(users)
	fmt.Fprint(w, string(js), "\n")
}
