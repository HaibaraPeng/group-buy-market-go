package server

import (
	"context"
	"log"
	"net/http"
)

type Server struct {
	httpServer *http.Server
	router     *http.ServeMux
}

func New() *Server {
	mux := http.NewServeMux()

	return &Server{
		router: mux,
		httpServer: &http.Server{
			Addr:    ":8080",
			Handler: mux,
		},
	}
}

func (s *Server) Route(pattern string, handler http.HandlerFunc) {
	s.router.HandleFunc(pattern, handler)
}

func (s *Server) Start() error {
	log.Println("Server starting on :8080")
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
