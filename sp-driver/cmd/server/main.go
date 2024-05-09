package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/cycle-labs/example-step-driver/sp-driver/api"
	"github.com/cycle-labs/example-step-driver/sp-driver/driver"
	"github.com/gin-gonic/gin"
)

func main() {

	// Parse command line flags
	port := flag.Int("port", 8081, "the port to listen on")
	flag.Parse()

	// Create a new Gin router
	router := gin.Default()

	// Define your routes here
	// provide a health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// wire up the API handlers
	service := driver.NewDriverService()
	handlers := driver.NewGinHandlers(service)
	api.RegisterHandlersWithOptions(router, handlers, api.GinServerOptions{})

	// Start the server
	addr := fmt.Sprintf(":%d", *port)
	log.Printf("Server listening on port %d", *port)
	log.Fatal(http.ListenAndServe(addr, router))
}
