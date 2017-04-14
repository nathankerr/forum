package users

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dhenkes/forum"
	"github.com/dhenkes/forum/postgres"
	"github.com/julienschmidt/httprouter"
)

func Get(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	user := forum.User{}

	err := postgres.QueryRow(&user, "SELECT uuid, username FROM users WHERE uuid = $1 AND removed = $2", id, "0")
	if err != nil {
		fmt.Println(err)
	}

	js, err := json.Marshal(user)
	fmt.Fprint(w, string(js), "\n")

	fmt.Fprint(w, "\n")
}
