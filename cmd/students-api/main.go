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

	"github.com/NEVIL77/students-api/internal/config" // 1. loading config
)

func main() {
	fmt.Println("Hello World")

	// 1 load config
	// 2 database setup
	// 3 setup route
	// 4 start server

	cfg := config.MustLoad() // 1. loading config

	router := http.NewServeMux() // 3. setup route

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {

		// func(w http.ResponseWriter, r *http.Request) -> This is the function that runs when the route is matched.

		// w is used to send a response back to the browser/client.

		// r contains information about the incoming request
		// Instead of giving the function a complete copy of the request, Go gives it the address/reference of the request.
		// Why? : Because http.Request is a relatively large struct. Passing a pointer avoids making a full copy and lets the function work with the original request.

		// * → pointer / value at an address
		// & → address of a variable

		w.Write([]byte("Hello World"))
	})

	// 4 start server
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	slog.Info("Server started", slog.String("adress", cfg.Addr))

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal("failed to start server")
		}

	}()

	<-done
	slog.Info("shutting down the server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("server shutdown failed", "error", err)
	}
	slog.Info("server gracefully stopped")

}
