package main

import (
	"encoding/json"
	"github.com/google/uuid"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	HEStatusUnsubmitted = "Unsubmitted"
	HEStatusSubmitted   = "Submitted"
	HEStatusCorrected   = "Corrected"
)

type Student struct {
	Firstname string
	Lastname  string
}

type File struct {
	Text string
}

type HE struct {
	HELinkUuid string
	HeUuid     string
	Student    Student
	File       File
	Status     string
}

func Homeworks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		homeworks_post(w, r)
		return
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
		return
	default:
		MethodNotAllowed(w)
		return
	}
}

func homeworks_post(w http.ResponseWriter, r *http.Request){
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	req := struct {
		HELinkUuid string
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
	indb,err := env.db.ExistsHelink(req.HELinkUuid)
	if !indb {
		log.Printf("HeLink does not exists: %v", req.HELinkUuid)
		http.Error(w, "HeLink does not exists", http.StatusNotFound)
		return
	}
	if err != nil {
		log.Printf("Error: could not connect to db")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	heuuid, err := uuid.NewUUID()
	if err != nil {
		log.Printf("Error creating UUID: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data := HE{
		req.HELinkUuid,
		heuuid.String(),
		Student{
			req.Firstname,
			req.Lastname},
		File{
			"", //TODO: Move to own constants file
		},
		HEStatusUnsubmitted}

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