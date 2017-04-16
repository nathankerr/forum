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

func Get(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	user := forum.User{}

	err := postgres.Get(&user, "SELECT uuid, username FROM users WHERE uuid = $1 AND removed = $2", ps.ByName("id"), "0")
	if err != nil {
		logger.Warning("%s", "Error during Request: SELECT uuid, username FROM users WHERE uuid = "+ps.ByName("id")+" AND removed = 0")
		logger.Warning("%s", err.Error())

		w.WriteHeader(http.StatusNotFound)

		response := forum.Response{
			Error: &forum.Error{
				Code:    http.StatusNotFound,
				Message: "NOT_FOUND",
			},
		}

		js, _ := json.Marshal(response)
		fmt.Fprint(w, string(js), "\n")
	} else {
		response := forum.Response{
			Data: user,
		}

		js, _ := json.Marshal(response)
		fmt.Fprint(w, string(js), "\n")
	}
}
