package main

import (
	"fmt"
	"log"
	"net/http"
//	"os"

	"github.com/gin-gonic/gin"
	"github.com/Nico2018/openmarket-backend/internal/config"
	"github.com/Nico2018/openmarket-backend/internal/routes"
)

func main() {
	// Load the configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Set the Gin mode
	if cfg.Server.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// Create the Gin router
	router := gin.New()

	// Add middleware to the router
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Add the routes to the router
	routes.AddRoutes(router)

	// Start the HTTP server
	addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	srv := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	log.Printf("Starting server on %s\n", addr)

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to start server: %v", err)
	}
}

