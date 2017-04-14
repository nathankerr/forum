package users

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dhenkes/forum"
	"github.com/dhenkes/forum/postgres"
	"github.com/julienschmidt/httprouter"
)

func GetAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	users := []forum.User{}
	user := forum.User{}

	err := postgres.Query(users, &user, "SELECT uuid, username FROM users WHERE removed = $1", "0")
	if err != nil {
		fmt.Println(err)
	}

	js, err := json.Marshal(users)
	fmt.Fprint(w, string(js), "\n")
	fmt.Fprint(w, "\n")
}
