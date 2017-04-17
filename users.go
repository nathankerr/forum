package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dhenkes/forum/logger"
	"github.com/dhenkes/forum/uuid"
	"github.com/julienschmidt/httprouter"
)

func getUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	user := User{}

	err := postgres.Get(&user, "SELECT uuid, username FROM users WHERE uuid = $1 AND removed = $2", ps.ByName("id"), "0")
	if err != nil {
		logger.Warning("%s", "Error during Request: SELECT uuid, username FROM users WHERE uuid = "+ps.ByName("id")+" AND removed = 0")
		logger.Warning("%s", err.Error())

		w.WriteHeader(http.StatusNotFound)

		response := Response{
			Error: &Error{
				Code:    http.StatusNotFound,
				Message: "NOT_FOUND",
			},
		}

		js, _ := json.Marshal(response)
		fmt.Fprint(w, string(js), "\n")
	} else {
		response := Response{
			Data: user,
		}

		js, _ := json.Marshal(response)
		fmt.Fprint(w, string(js), "\n")
	}
}

func getAllUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	users := []User{}

	err := postgres.Select(&users, "SELECT uuid, username FROM users WHERE removed = $1", "0")
	if err != nil {
		logger.Warning("%s", "Error during Request: SELECT uuid, username FROM users WHERE removed = 0")
		logger.Warning("%s", err.Error())

		w.WriteHeader(http.StatusNotFound)
	}

	js, _ := json.Marshal(users)
	fmt.Fprint(w, string(js), "\n")
}

func createUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("content-type", "application/json")

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		output := CouldNotReadBody(w)
		w.Write(output)
		return
	}

	var message User
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

	_, err = postgres.Exec("INSERT INTO users(uuid, username, password, created) VALUES($1,$2,$3, $4)", message.Uuid, message.Username, message.Password, message.Created)
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

	response := Response{
		Error: &Error{
			Code:    http.StatusNotFound,
			Message: "COULD_NOT_READ_BODY",
		},
	}

	output, _ := json.Marshal(response)
	return output
}

func CouldNotInsert(w http.ResponseWriter) []byte {
	w.WriteHeader(http.StatusInternalServerError)

	response := Response{
		Error: &Error{
			Code:    http.StatusNotFound,
			Message: "COULD_NOT_INSERT_INTO_DB",
		},
	}

	output, _ := json.Marshal(response)
	return output
}
