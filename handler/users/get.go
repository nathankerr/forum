package users

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dhenkes/forum"
	fhttp "github.com/dhenkes/forum/http"
	"github.com/julienschmidt/httprouter"
)

func Get(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	user := forum.User{}
	fhttp.Server.Couchbase.Bucket.Get("u:"+id, &user)
	jsonBytes, _ := json.Marshal(user)
	fmt.Fprint(w, string(jsonBytes), "\n")
}
