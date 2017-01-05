package http

import (
	"log"
	"net"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type server struct {
	ln     net.Listener
	port   string
	Router *httprouter.Router
}

func CreateServer(port string) *server {
	return &server{
		port:   port,
		Router: httprouter.New(),
	}
}

func (s *server) Run() error {
	log.Fatal(http.ListenAndServe(s.port, s.Router))
	return nil
}
