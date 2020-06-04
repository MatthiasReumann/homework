package main

import (
	"log"
	"net/http"
	"time"
)

type server struct {
	srv *http.Server
}

func NewServer(port string) server {
	router := NewRouter()

	srv := &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return server{srv}
}

func (s *server) Serve() {
	log.Printf("Listen at %s\n", s.srv.Addr)
	err := s.srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func MethodNotAllowed(w http.ResponseWriter) {
	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
}
