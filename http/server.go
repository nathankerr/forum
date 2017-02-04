package http

import (
	"log"
	"net"
	"net/http"

	"github.com/dhenkes/forum/couchbase"
	"github.com/julienschmidt/httprouter"
)

type server struct {
	ln        net.Listener
	port      string
	Router    *httprouter.Router
	Couchbase *couchbase.Couchbase
}

var Server server

func CreateServer(port string, cb *couchbase.Couchbase) {
	Server = server{
		port:      port,
		Router:    httprouter.New(),
		Couchbase: cb,
	}
}

func Run() error {
	log.Fatal(http.ListenAndServe(Server.port, Server.Router))
	return nil
}
