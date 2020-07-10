package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/JustSteveKing/go-ship/database"
	"github.com/JustSteveKing/go-ship/handlers"
	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	// Create a new Logger instance to pass into our Handlers
	logger := log.New(os.Stdout, "go-ship-service ", log.LstdFlags)

	// Initialise our database
	database.InitDB()

	// register handlers
	pingHandler := handlers.NewPingHandler(logger)

	// Create a new serve mux and register handlers
	serveMux := mux.NewRouter()

	// Create specific sub routers
	getRouter := serveMux.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", pingHandler.Handle)

	// Handle CORS
	corsHandler := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))

	// Create a new server
	server := &http.Server{
		Addr:         ":8080",
		Handler:      corsHandler(serveMux),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// Listen and serve from the server
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	// close connection when done
	defer database.DBCon.Close()

	// Graceful Shutdown
	waitForShutdown(server, logger)
}

func waitForShutdown(server *http.Server, logger *log.Logger) {
	// Create a Signal chanel to listen to OS signals
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, os.Kill, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal
	shutdownSignal := <-interruptChan
	logger.Println("Received terminate signal, graceful shutdown", shutdownSignal)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	server.Shutdown(ctx)

	os.Exit(0)
}
