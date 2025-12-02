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

// RegisterRoutes registers all HTTP routes
func (s *Server) RegisterRoutes() {
	s.Route("/", s.handleHome)
	s.Route("/test", s.handleTest)
}

// handleHome handles requests to the home page
func (s *Server) handleHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, Group Buy Market!"))
}

// handleTest handles requests to the test endpoint
func (s *Server) handleTest(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Test endpoint is working!"))
}
