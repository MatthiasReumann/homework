package main

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
)


func NewRouter() http.Handler{
	r := mux.NewRouter()
	r.HandleFunc("/links", Links)
	r.HandleFunc("/homeworks", Homeworks)
	r.HandleFunc("/homeworks/{uuid}", HomeworksUUID)

	r.Use(log_request)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:8080"},
		AllowedMethods: []string{"GET", "POST", "PUT", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	})

	return c.Handler(r)
}

func log_request(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s -> %s %s\n", r.RemoteAddr, r.URL.Path, r.Method)
		next.ServeHTTP(w, r)
	})
}

