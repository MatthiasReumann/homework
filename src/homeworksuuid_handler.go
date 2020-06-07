package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

func (s *server) HomeworksUUID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.homeworkuuid_get(w, r)
		return
	case http.MethodPut:
		s.homeworkuuid_put(w, r)
		return
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK);
		return
	default:
		s.MethodNotAllowed(w)
	}
}

func (s *server) homeworkuuid_get(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	heuuid := vars["uuid"]
	log.Print(heuuid)

	//TODO: check if heuuid is in db  //TODO: load text from dataabase

	file := File{
		"text text",
		HEStatusUnsubmitted,
	}

	res, err := json.Marshal(file)
	if err != nil {
		log.Printf("Error while marshalling: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write(res)
}

func (s *server) homeworkuuid_put(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	heuuid := vars["uuid"]

	//TODO: check if heuuid is in db

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var req File

	err = json.Unmarshal(body, &req)
	if err != nil {
		log.Printf("Error Unmarshal body: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//TODO: update text via heuuid

	log.Printf("updated %s", heuuid)

	w.WriteHeader(http.StatusOK)
}