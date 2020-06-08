package main

import (
	"encoding/json"
	uuid2 "github.com/google/uuid"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	HEStatusUnsubmitted = "Unsubmitted"
	HEStatusSubmitted   = "Submitted"
	HEStatusCorrected   = "Corrected"
	EmptyString         = ""
)

type Student struct {
	Firstname string
	Lastname  string
}

type File struct {
	Text   string
	Status string
}

type Submission struct {
	Uuid    string
	Link    Link
	Student Student
	File    File
}

func (s *server) Homeworks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		s.homeworks_post(w, r)
		return
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
		return
	default:
		s.MethodNotAllowed(w)
		return
	}
}

func (s *server) homeworks_post(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	req := struct {
		Uuid string
		Firstname  string
		Lastname   string
	}{}

	err = json.Unmarshal(body, &req)
	if err != nil {
		log.Printf("Error Unmarshal body: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//check if helink in db
	indb,err := env.db.ExistsHelink(req.Uuid)
	if !indb {
		log.Printf("HeLink does not exists: %v", req.Uuid)
		http.Error(w, "HeLink does not exists", http.StatusNotFound)
		return
	}
	if err != nil {
		log.Printf("Error: could not connect to db")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	uuid, err := uuid2.NewUUID()
	if err != nil {
		log.Printf("Error creating UUID: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data := Submission{
		uuid.String(),
		Link{req.Uuid},
		Student{
			req.Firstname,
			req.Lastname},
		File{
			EmptyString,
			HEStatusUnsubmitted}}

	//add he
	err = env.db.AddHe(data)
	if err != nil {
		log.Printf("Error: could not connect to db")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := json.Marshal(data)
	if err != nil {
		log.Printf("Error while marshalling: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write(res)
}
