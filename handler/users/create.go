package users

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dhenkes/forum"
	"github.com/dhenkes/forum/couchbase"
	"github.com/dhenkes/forum/utils"
	"github.com/julienschmidt/httprouter"
)

type user struct {
	Username string
	Password string
}

func Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)

	var u user
	err := decoder.Decode(&u)

	if err != nil {
		fmt.Println(err)
	}

	var hash []byte
	hash, _ = utils.HashPassword(u.Password)
	hashString := string(hash[:])

	u.Password = hashString

	var userCheck user
	_, err = couchbase.DB.Bucket.Get("u:"+u.Username, &userCheck)
	if err == nil {
		fmt.Println("User does exist")
	}

	overview := forum.Overview{}
	couchbase.DB.Bucket.Get("f:overview", &overview)
	overview.Users = append(overview.Users, "u:"+u.Username)
	_, err = couchbase.DB.Bucket.Upsert("f:overview", &overview, 0)

	couchbase.DB.Bucket.Insert("u:"+u.Username, u, 0)

	fmt.Fprint(w, string("User created"), "\n")
}
