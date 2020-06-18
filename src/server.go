package main

import (
	"log"
	"net/http"
	"time"
)

type server struct {
	httpServer *http.Server
	db *databaseConnection
}

func NewServer(port string, db *databaseConnection) server {
	srv := &http.Server{
		Addr:           ":" + port,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return server{srv, db}
}

func (s *server) Serve() {
	log.Printf("Listen at %s\n", s.httpServer.Addr)
	s.httpServer.Handler = s.NewRouter()

	err := s.httpServer.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func (s *server) MethodNotAllowed(w http.ResponseWriter) {
	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
}
