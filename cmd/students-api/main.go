package main

import (
	"fmt"
	"log"
	"net/http"

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

	fmt.Println(cfg)
	fmt.Println("started  server on ", cfg.Addr)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("failed to start server")
	}

}
