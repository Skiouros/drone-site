package server

import "net/http"
import "github.com/gorilla/mux"
import "github.com/codegangsta/negroni"

var server Server

type Server struct {
	Router *mux.Router
}

func (server *Server) Route(path string, h http.HandlerFunc, middleware ...func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc)) *mux.Route {
	n := negroni.New()
	for _, m := range middleware {
		n.Use(negroni.HandlerFunc(m))
	}
	n.Use(negroni.Wrap(h))
	return server.Router.Handle(path, n)
}

func (server *Server) Start(port string) {
	http.Handle("/", server.Router)
	http.ListenAndServe(":" + port, nil)
}

func GetServer() *Server {
	return &server
}

func init() {
	r := mux.NewRouter()
	r = r.StrictSlash(true)
	server = Server{ r }
}
