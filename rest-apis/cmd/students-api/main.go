package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gunjanghate/learning-go/internal/config"
)

func main() {
	slog.Info("Starting Students API...")
	fmt.Println("Starting Students API...")
	// load config
	cfg := config.MustLoad()
	// fmt.Printf("Loaded config: %+v\n", cfg)
	// databse setup
	// setup router
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("welcome to students api"))
	} )
	// setup server

	server := http.Server{
		Addr: cfg.Addr,
		Handler: router,
	}
	slog.Info("Server Started", slog.String("address", cfg.HTTPServer.Addr))
	fmt.Printf("Server Started at %s", cfg.HTTPServer.Addr)
    // garcefully shutdown
    done := make(chan os.Signal, 1)


	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	// Interrupt will occurr if user presses Ctrl+C or the process receives a termination signal
	// SIGINT and SIGTERM are specific types of termination signals
	go func(){

		err := server.ListenAndServe()
	
		if err != nil{
			log.Fatal("Error starting server:", err)
		}
	}()


	<- done

	// log an informational message indicating shutdown has startedful shutdown by signal.Notify
	slog.Info("shutting down the server")
	// create a context with a 5-second timeout to bound the graceful shutdown period
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// ensure the cancel function is called to release resources when the function returns
	defer cancel()
	
	// attempt to gracefully shut down the server using the context timeout
	err := server.Shutdown(ctx)
	// if shutdown returned an error, log it with a structured "error" field
	if err != nil {
		slog.Error("error during server shutdown", slog.String("error", err.Error()))
	}
	// log that the server has been stopped
	slog.Info("server stopped")
	slog.Info("server stopped")





}