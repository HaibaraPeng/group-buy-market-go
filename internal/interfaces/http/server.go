package http

import (
	"net/http"
)

type Server struct {
	router *http.ServeMux
}

func NewServer() *Server {
	return &Server{
		router: http.NewServeMux(),
	}
}

func (s *Server) Route(pattern string, handler http.HandlerFunc) {
	s.router.HandleFunc(pattern, handler)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
