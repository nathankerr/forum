package users

import (
	"fmt"
	"net/http"

	"github.com/dhenkes/forum/postgres"
	"github.com/julienschmidt/httprouter"
)

func GetAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	q := postgres.NewQuery()
	q.Select("*").From("users").Where("removed").Equals("0")
	q.Execute()

	fmt.Fprint(w, "getall", "\n")
}
