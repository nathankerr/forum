package users

// type user struct {
// 	username string
// 	password string
// }

// func Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	var hash []byte
// 	hash, _ = utils.HashPassword("abc")

// 	hashString := string(hash[:])

// 	user := forum.User{hashString}
// 	_, err := fhttp.Server.Couchbase.Bucket.Insert("u:"+hashString, user, 0)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	overview := forum.Overview{}
// 	fhttp.Server.Couchbase.Bucket.Get("f:overview", &overview)
// 	overview.Users = append(overview.Users, hashString)
// 	_, err = fhttp.Server.Couchbase.Bucket.Upsert("f:overview", &overview, 0)

// 	fmt.Fprint(w, string("User created"), "\n")
// }
