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

var Server server

func CreateServer(port *string) {
	Server = server{
		port:   *port,
		Router: httprouter.New(),
	}
}

func Run() error {
	log.Fatal(http.ListenAndServe(Server.port, Server.Router))
	return nil
}
