package main

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
)


func (s *server) NewRouter() http.Handler{
	r := mux.NewRouter()
	r.HandleFunc("/links", s.Links)
	r.HandleFunc("/links/{uuid}", s.LinksUUID)
	r.HandleFunc("/homeworks", s.Homeworks)
	r.HandleFunc("/homeworks/{uuid}", s.HomeworksUUID)

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

