package http

import (
	"log"
	"net"
	"net/http"

	"github.com/dhenkes/forum/mysql"
	"github.com/julienschmidt/httprouter"
)

type server struct {
	ln     net.Listener
	port   string
	Router *httprouter.Router
	MySQL  *mysql.MySQL
}

// Creates the http server with the given port and initiates a router.
func CreateServer(port string) *server {
	return &server{
		port:   port,
		Router: httprouter.New(),
	}
}

// Sets MySQL as the database being used.
func (s *server) UseMySQL(db *mysql.MySQL) {
	s.MySQL = db
}

// Starts the http server.
func (s *server) Run() error {
	log.Fatal(http.ListenAndServe(s.port, s.Router))
	return nil
}
