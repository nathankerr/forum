package users

// func GetAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	start := time.Now()

// 	overview := forum.Overview{}
// 	fhttp.Server.Couchbase.Bucket.Get("f:overview", &overview)

// 	log.Printf("Time for GET request %s", time.Since(start))
// 	start = time.Now()

// 	var items []gocb.BulkOp
// 	for i := 0; i < len(overview.Users); i++ {
// 		items = append(items, &gocb.GetOp{Key: overview.Users[i], Value: &forum.User{}})
// 	}

// 	log.Printf("Time for FOR loop %s", time.Since(start))
// 	start = time.Now()

// 	err := fhttp.Server.Couchbase.Bucket.Do(items)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	log.Printf("Time for DO request %s", time.Since(start))
// 	start = time.Now()

// 	var users []*forum.User
// 	for i := 0; i < len(items); i++ {
// 		users = append(users, items[i].(*gocb.GetOp).Value.(*forum.User))
// 	}

// 	log.Printf("Time for second FOR loop %s", time.Since(start))

// 	jsonBytes, _ := json.Marshal(users)
// 	fmt.Fprint(w, string(jsonBytes), "\n")
// }
