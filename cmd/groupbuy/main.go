package main

import (
	"group-buy-market-go/pkg/server"
	"log"
	"net/http"
)

func main() {
	s := server.New()
	s.Route("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, Group Buy Market!"))
	})

	s.Route("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Test endpoint is working!"))
	})

	log.Println("Server starting on :8080")
	err := s.Start()
	if err != nil {
		log.Fatal(err)
	}
}
