package http

import (
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

type server struct {
	ln     net.Listener
	port   string
	Router *httprouter.Router
}

var Server server

func CreateServer(port string) {
	if strings.HasPrefix(port, ":") == false {
		port = ":" + port
	}
	Server = server{
		port:   port,
		Router: httprouter.New(),
	}
}

func Run() error {
	log.Fatal(http.ListenAndServe(Server.port, Server.Router))
	return nil
}
