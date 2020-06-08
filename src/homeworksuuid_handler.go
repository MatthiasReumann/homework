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

	//check if he exists
	indb,err := env.db.ExistsHe(heuuid)
	if !indb {
		log.Printf("HE does not exists: %v", heuuid)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if err != nil {
		log.Printf("Error: could not connect to db")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//get file
	text,_ := env.db.GetFile(heuuid)
	file := File{
		text,
		HEStatusUnsubmitted,
	}

	if err != nil {
		log.Printf("Error: could not connect to db")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
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

	//check if he exists
	indb,err := env.db.ExistsHe(heuuid)
	if !indb {
		log.Printf("HE does not exists: %v", heuuid)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if err != nil {
		log.Printf("Error: could not connect to db")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

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

	//check if file exists
	err = env.db.SetFile(heuuid, req.Text)
	if err != nil {
		log.Printf("Error: could not connect to db")
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if err != nil {
		log.Printf("Error: could not connect to db")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("updated %s", heuuid)

	w.WriteHeader(http.StatusOK)
}