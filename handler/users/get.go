package users

import (
	"fmt"
	"net/http"

	"github.com/dhenkes/forum/postgres"
	"github.com/julienschmidt/httprouter"
)

func Get(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	q := postgres.NewQuery()
	q.Select("*").From("users").Where("uuid", "removed").Equals(id, "0")
	q.Execute()

	fmt.Fprint(w, "get", "\n")
}
