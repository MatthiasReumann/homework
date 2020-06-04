package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

const (
	HEStatusUnsubmitted = "Unsubmitted"
	HEStatusSubmitted = "Submitted"
	HEStatusCorrected= "Corrected"
)

type server struct{
	srv *http.Server
}

type HELink struct{
	HELinkUuid string
}

type Student struct{
	Firstname string
	Lastname string
}

type File struct{
	Text string
}

type HE struct{
	HELinkUuid string
	HeUuid string
	Student Student
	File File
	Status string
}

func NewServer(port string) server{
	r := mux.NewRouter()
	r.HandleFunc("/links", Links).Methods("POST","OPTIONS")
	r.HandleFunc("/homeworks", Homeworks).Methods("POST","OPTIONS")
	r.HandleFunc("/homeworks/{uuid}", HomeworksUUID).Methods("PUT","GET","OPTIONS")

	srv := &http.Server{
		Addr:           ":"+port,
		Handler:		r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return server{srv}
}

func (s *server) Serve(){
	log.Printf("Listen at %s\n", s.srv.Addr)
	err := s.srv.ListenAndServe()
	if err != nil{
		log.Fatal(err)
	}
}
