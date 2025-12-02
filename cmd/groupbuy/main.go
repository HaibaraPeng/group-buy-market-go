package main

import (
	"fmt"
	"group-buy-market-go/internal/infrastructure"
	httpInterface "group-buy-market-go/internal/interfaces/http"
	"log"
	"net/http"
)

func main() {
	// Load configuration
	config, err := infrastructure.LoadConfig("configs/config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize database
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Initialize server with Wire
	server := httpInterface.NewServer(db)

	// Register routes
	server.RegisterRoutes()

	log.Printf("Server starting on %s:%d", config.Server.Host, config.Server.Port)
	// Start server
	addr := fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)
	err = http.ListenAndServe(addr, server)
	if err != nil {
		log.Fatal(err)
	}
}
