package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

const (
	maxHeaderBytes    = 1024
	readHeaderTimeout = 0o3
	readTimeout       = 0o3
	writeTimeout      = 15
	idleTimeout       = 60
	graceTime         = 15
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Print("Received request")
}

func router() {
	router := mux.NewRouter()
	router.HandleFunc("/hi", handler).Methods("GET")

	listen(":8081", router)
}

// we have configured the server
func listen(address string, handler http.Handler) {
	server := &http.Server{
		Addr:              address,
		Handler:           handler,
		MaxHeaderBytes:    maxHeaderBytes,
		ReadHeaderTimeout: readHeaderTimeout * time.Second,
		ReadTimeout:       readTimeout * time.Second,
		WriteTimeout:      writeTimeout * time.Second, // response stream
		IdleTimeout:       idleTimeout * time.Second,
	}

	log.Printf("Started listening on: %s\n", address)

	// Graceful shutdown
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	log.Println("shutting down")
	// Create a deadline to wait for.

	// context - package in golan, all the functions that you are going to write in a webservice
	// first param will be a context.

	// your benchmarking and your microservice
	ctx, cancel := context.WithTimeout(context.Background(), graceTime*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Print("Error resulted in the shutdown")
	}
}

// Create more handlers, try to follow REST API standards
// 4 handlers to handle
// get all employees
// save employee
// get one employee by id
// update an employee

// Read about project layout in golang
// https://github.com/golang-standards/project-layout
// controller - service- repository
// packages should look like
