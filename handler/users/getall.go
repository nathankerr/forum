package users

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/couchbase/gocb"
	"github.com/dhenkes/forum"
	"github.com/dhenkes/forum/couchbase"
	"github.com/julienschmidt/httprouter"
)

func GetAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	overview := forum.Overview{}
	couchbase.DB.Bucket.Get("f:overview", &overview)

	var items []gocb.BulkOp
	for i := 0; i < len(overview.Users); i++ {
		items = append(items, &gocb.GetOp{Key: overview.Users[i], Value: &forum.User{}})
	}

	err := couchbase.DB.Bucket.Do(items)
	if err != nil {
		fmt.Println(err)
	}

	var users []*forum.User
	for i := 0; i < len(items); i++ {
		users = append(users, items[i].(*gocb.GetOp).Value.(*forum.User))
	}

	jsonBytes, _ := json.Marshal(users)
	fmt.Fprint(w, string(jsonBytes), "\n")
}
