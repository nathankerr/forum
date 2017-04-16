package users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dhenkes/forum"
	"github.com/dhenkes/forum/postgres"
	"github.com/dhenkes/forum/uuid"
	"github.com/julienschmidt/httprouter"
)

func Post(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("content-type", "application/json")

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		output := CouldNotReadBody(w)
		w.Write(output)
		return
	}

	var message forum.User
	err = json.Unmarshal(body, &message)
	if err != nil {
		output := CouldNotReadBody(w)
		w.Write(output)
		return
	}

	message.Uuid = uuid.NewV5(message.Username)
	message.Created = int(time.Now().Unix())

	output, err := json.Marshal(message)
	if err != nil {
		output := CouldNotReadBody(w)
		w.Write(output)
		return
	}

	err = postgres.Insert("INSERT INTO users(uuid, username, password, created) VALUES($1,$2,$3, $4)", message.Uuid, message.Username, message.Password, message.Created)
	if err != nil {
		fmt.Println(err)
		output := CouldNotInsert(w)
		w.Write(output)
		return
	}

	w.Write(output)
}

func CouldNotReadBody(w http.ResponseWriter) []byte {
	w.WriteHeader(http.StatusInternalServerError)

	response := forum.Response{
		Error: &forum.Error{
			Code:    http.StatusNotFound,
			Message: "COULD_NOT_READ_BODY",
		},
	}

	output, _ := json.Marshal(response)
	return output
}

func CouldNotInsert(w http.ResponseWriter) []byte {
	w.WriteHeader(http.StatusInternalServerError)

	response := forum.Response{
		Error: &forum.Error{
			Code:    http.StatusNotFound,
			Message: "COULD_NOT_INSERT_INTO_DB",
		},
	}

	output, _ := json.Marshal(response)
	return output
}
