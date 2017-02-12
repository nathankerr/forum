package users

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dhenkes/forum"
	"github.com/dhenkes/forum/couchbase"
	"github.com/julienschmidt/httprouter"
)

func Get(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	user := forum.User{}
	err := couchbase.Get("u:"+id, &user)
	if err != nil {
		fmt.Println(err)
	}
	jsonBytes, _ := json.Marshal(user)
	fmt.Fprint(w, string(jsonBytes), "\n")
}
