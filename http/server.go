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

func CreateServer(port string, cb *couchbase.Couchbase) *server {
	return &server{
		port:      port,
		Router:    httprouter.New(),
		Couchbase: cb,
	}
}

func (s *server) Run() error {
	log.Fatal(http.ListenAndServe(s.port, s.Router))
	return nil
}
