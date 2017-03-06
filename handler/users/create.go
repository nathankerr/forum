package users

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/dhenkes/forum"
	"github.com/dhenkes/forum/couchbase"
	"github.com/dhenkes/forum/utils"
	"github.com/dhenkes/forum/uuid"
	"github.com/julienschmidt/httprouter"
)

func Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)

	var newUser forum.NewUser
	err := decoder.Decode(&newUser)
	if err != nil {
		fmt.Println(err)
	}

	newUser.Username = strings.ToLower(newUser.Username)
	newUser.Username = strings.TrimSpace(newUser.Username)
	uuid := uuid.NewV5(newUser.Username)

	var user forum.User
	err = couchbase.Get("u:"+uuid, &user)
	if err == nil {
		fmt.Fprint(w, string("User already exists"), "\n")
		return
	}

	var hash []byte
	hash, _ = utils.HashPassword(newUser.Password)
	hashString := string(hash[:])
	newUser.Password = hashString

	users := forum.Users{}
	couchbase.Get("f:users", &users)
	users.Users = append(users.Users, "u:"+uuid)
	err = couchbase.Upsert("f:users", &users, 0)
	if err != nil {
		fmt.Println(err)
	}

	err = couchbase.Insert("u:"+uuid, newUser, 0)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprint(w, string("User created"), "\n")
}
