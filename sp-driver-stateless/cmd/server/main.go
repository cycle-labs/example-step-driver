package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/cycle-labs/example-step-driver/sp-driver-stateless/api"
	"github.com/cycle-labs/example-step-driver/sp-driver-stateless/driver"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {

	// Parse command line flags
	port := flag.Int("port", 8082, "the port to listen on")
	flag.Parse()

	done := make(chan os.Signal)

	// Create a new Gin router
	router := gin.Default()

	// Define your routes here
	// provide a health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	router.DELETE("/shutdown", func(c *gin.Context) {
		defer func() {
			done <- os.Kill
		}()
		c.JSON(200, gin.H{"status": "ok"})
	})

	dbUrl, exists := os.LookupEnv("DATABASE_URL")
	if !exists {
		dbUrl = "postgresql://cycle:labs1@localhost:5432/accountdb"
	}

	// set up the DB connection pool
	pool, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}

	// wire up the API handlers
	service := driver.NewDriverService(pool)
	handlers := driver.NewGinHandlers(service)
	api.RegisterHandlersWithOptions(router, handlers, api.GinServerOptions{})

	// Start the server
	addr := fmt.Sprintf(":%d", *port)
	log.Printf("Server listening on port %d", *port)

	server := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("error: %s\n", err)
		}
	}()

	// blocks until a signal is received
	<-done

	log.Println("Shutting down server...")
	pool.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Error while shutting down Server. Initiating force shutdown...", err.Error())
	} else {
		log.Default().Println("Server exiting")
	}
}
