package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	middleware "server.simplifycontrol.com/middlewares"
	paymentHelpers "server.simplifycontrol.com/payment"

	"github.com/rs/cors"
	"server.simplifycontrol.com/helpers/database"
	"server.simplifycontrol.com/routes"
	"server.simplifycontrol.com/secrets"
)

func main() {
	if err := secrets.FetchSecretFile(); err != nil {
		fmt.Printf("Failed to fetch secrets: %v\n", err)
		os.Exit(1)
	}

	// Initialize the database connection
	database.InitializeDatabase()
	defer database.CloseDatabase()

	// Initialize payment gateway
	paymentHelpers.InitPayment()

	// Setup routes
	mux := routes.InitializeRoutes()

	// Chain your internal middlewares (logging and JWT auth)
	chainedHandler := middleware.Chain(
		mux,
		middleware.LoggingMiddleware,
		middleware.JWTAuthMiddleware,
	)

	// Now wrap the entire chain with the CORS handler.
	// This will allow all origins; adjust settings as necessary.
	handler := cors.AllowAll().Handler(chainedHandler)

	// Create a new HTTP server with a graceful shutdown mechanism
	server := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	// Start the server in a goroutine
	go func() {
		fmt.Println("Server is running on http://localhost:8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Server stopped: %v\n", err)
		}
	}()

	// Setup channel to listen for shutdown signals (SIGINT, SIGTERM)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Block until we receive a shutdown signal
	<-stop
	fmt.Println("Shutting down the server...")

	// Graceful shutdown: allow up to 5 seconds for pending requests to finish
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt to gracefully stop the server
	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("Server shutdown failed: %v\n", err)
	} else {
		fmt.Println("Server gracefully stopped.")
	}
}
